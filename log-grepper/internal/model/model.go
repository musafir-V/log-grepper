package model

import "os"

type Request struct {
	SearchKeyword string `json:"searchKeyword"`
	From          string `json:"from"`
	To            string `json:"to"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AwsConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
}

var Config = AwsConfig{
	Region:          getEnvAsString("AWS_REGION", "us-east-1"),
	AccessKeyID:     getEnvAsString("AWS_ACCESS_KEY_ID", "0"),
	SecretAccessKey: getEnvAsString("AWS_SECRET_ACCESS_KEY", "0"),
}

func getEnvAsString(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
