package service

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/musafir-V/log-grepper/internal/constant"
	"github.com/musafir-V/log-grepper/internal/dao"
)

type LogGrepperService interface {
	GrepLogs(to, from time.Time, match string) []*string
}

type logGrepperService struct {
	dao dao.DAO
}

func NewLogGrepperService(dao dao.DAO) LogGrepperService {
	return &logGrepperService{
		dao: dao,
	}
}

func (l logGrepperService) GrepLogs(to, from time.Time, match string) []*string {
	folders := getFolderNames(from, to)
	var wg sync.WaitGroup
	ch := make(chan []*string, len(folders)*len(constant.FileNames))
	for _, folder := range folders {
		for _, file := range constant.FileNames {
			fd := folder
			fl := file
			wg.Add(1)
			go func(folder, file, match string) {
				defer wg.Done()
				res, _ := l.processFile(folder, file, match)
				ch <- res
			}(fd, fl, match)
		}
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	var ret []*string
	for res := range ch {
		ret = append(ret, res...)
	}
	return ret
}

func (l logGrepperService) processFile(folder, file, match string) ([]*string, error) {
	content, err := l.dao.GetLogs(folder, file)
	if err != nil {
		fmt.Printf("error getting logs for folder %s and file %s: %v ignoring file\n", folder, file, err)
		return nil, fmt.Errorf("error getting logs for folder %s and file %s: %w", folder, file, err)
	}
	ret := make([]*string, 0)
	// cut the content into lines
	lines := strings.Split(*content, "\n")
	for _, line := range lines {
		if strings.Contains(line, match) {
			ret = append(ret, &line)
		}
	}
	return ret, nil
}

func getFolderNames(from time.Time, to time.Time) []string {
	var buckets []string
	for t := from; t.Before(to) || t.Equal(to); t = t.AddDate(0, 0, 1) {
		buckets = append(buckets, t.Format(constant.LogFolderTimeLayout))
	}
	return buckets
}
