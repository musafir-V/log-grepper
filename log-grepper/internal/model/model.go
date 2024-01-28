package model

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
	Region:          "us-east-1",
	AccessKeyID:     "abc",
	SecretAccessKey: "xyz",
}
