package main

import (
	"github.com/musafir-V/log-grepper/internal/model"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/musafir-V/log-grepper/internal/dao"
	"github.com/musafir-V/log-grepper/internal/server"
	"github.com/musafir-V/log-grepper/internal/service"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(model.Config.Region),
		Region:      aws.String(model.Config.Region),
		Credentials: credentials.NewStaticCredentials(model.Config.AccessKeyID, model.Config.SecretAccessKey, ""),
	})
	if err != nil {
		panic(err)
	}
	client := s3.New(sess)
	s3dao := dao.NewS3Dao(client)
	grepper := service.NewLogGrepperService(s3dao)
	http.HandleFunc("/search", server.NewSearchHandler(grepper).GetMatchingLogs)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
