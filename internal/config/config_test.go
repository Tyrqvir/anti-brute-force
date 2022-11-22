package config

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr error
	}{
		{
			name: "No such config file return error",
			args: args{
				configFile: "notExist.toml",
			},
			want:    &Config{},
			wantErr: errors.New("failed to read config file: open notExist.toml: no such file or directory"),
		},
		{
			name: "valid config file",
			args: args{
				configFile: "../../configs/config.toml",
			},
			want: &Config{
				GRPC: GRPCConf{
					Address: ":8889",
				},
				Logger: LoggerConf{
					Level: "debug",
				},
				Redis: RedisConf{
					DSN: "127.0.0.1:6379",
				},
				RateLimit: RateLimitConf{
					LoginPerMinute:    10,
					PasswordPerMinute: 100,
					IPPerMinute:       1000,
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.args.configFile)
			if tt.wantErr != nil {
				assert.EqualErrorf(t, err, tt.wantErr.Error(), "Error should be: %v, got: %v", tt.wantErr, err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
