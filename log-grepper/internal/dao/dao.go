package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DAO is the interface for data access object. This can be used to fetch data from any data source.
// For example, in this case we are fetching data from S3. Others can be local, azure blob store, etc.
type DAO interface {
	GetLogs(folder, file string) (*string, error)
}

type s3Dao struct {
	client *s3.S3
}

func NewS3Dao(client *s3.S3) DAO {
	return &s3Dao{
		client: client,
	}
}

func (s s3Dao) GetLogs(folder, file string) (*string, error) {
	fmt.Println(folder)
	input := &s3.GetObjectInput{
		Bucket: aws.String(folder),
		Key:    aws.String(file),
	}

	result, err := s.client.GetObject(input)
	defer result.Body.Close()

	// Read object content
	buf := make([]byte, *result.ContentLength)
	_, err = result.Body.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("error reading object content for folder %s and file %s: %w", folder, file, err)
	}

	ret := string(buf)
	return &ret, nil
}
