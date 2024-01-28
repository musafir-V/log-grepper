package server

import (
	"github.com/musafir-V/log-grepper/internal/model"
	"testing"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name    string
		req     *model.Request
		wantErr bool
	}{
		{
			name: "from date not parsable",
			req: &model.Request{
				SearchKeyword: "error",
				From:          "2020-01-x01",
				To:            "2020-01-02",
			},
			wantErr: true,
		},
		{
			name: "to date not parsable",
			req: &model.Request{
				SearchKeyword: "error",
				From:          "2020-01-01",
				To:            "2020-01-x02",
			},
			wantErr: true,
		},
		{
			name: "both dates not parsable",
			req: &model.Request{
				SearchKeyword: "error",
				From:          "2020-01-ss01",
				To:            "2020-01-x02",
			},
			wantErr: true,
		},
		{
			name: "from date after to date",
			req: &model.Request{
				SearchKeyword: "error",
				From:          "2020-01-05",
				To:            "2020-01-02",
			},
			wantErr: true,
		},
		{
			name: "from date is more than 7 days before to date",
			req: &model.Request{
				SearchKeyword: "error",
				From:          "2020-01-01",
				To:            "2020-01-09",
			},
			wantErr: true,
		},
		{
			name: "valid request",
			req: &model.Request{
				SearchKeyword: "error",
				From:          "2020-01-01",
				To:            "2020-01-02",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := validateRequest(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateRequest() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
		})
	}
}
