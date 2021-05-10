package signhost

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		t      bool
		auth   string
		appKey string
	}

	tests := []struct{
		name string
		args args
		want *Config
	} {
		{
			"config set to testing and with API token and app token",
			args{
				t: true,
				auth: APITokenEnv,
				appKey: AppKeyEnv,
			},
			&Config{
				testing: true,
				auth: "SIGNHOST_API_TOKEN",
				appKey: "SIGNHOST_APP_KEY",
			},
		},
		{
			"config set to testing with develop and only app token",
			args{
				t:      true,
				appKey: AppKeyEnv,
			},
			&Config{
				testing: true,
				appKey:  "SIGNHOST_APP_KEY",
			},
		},
		{
			"config set to testing with develop and only API token",
			args{
				t:      true,
				auth: APITokenEnv,
			},
			&Config{
				testing: true,
				auth:  "SIGNHOST_API_TOKEN",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.t, tt.args.auth, tt.args.appKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}