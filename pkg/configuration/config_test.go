package configuration

import (
	"reflect"
	"testing"
)

func TestMustReadConfig(t *testing.T) {
	type args struct {
		configFileLocation string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Read config from config/auditmon.yaml with no overrides",
			args: args{configFileLocation: "../../config/auditmon.yaml"},
			want: Config{
				Database: DatabaseConfig{
					Backend:  "sqlite3",
					File:     "auditmon.db",
					Host:     "localhost",
					Port:     5432,
					SSLMode:  "prefer",
					Username: "postgres",
					Password: "postgres",
					Name:     "auditmon",
				},
				Settings: SettingsConfig{
					RetentionDays: 365,
					MaxQueryDays:  30,
				},
				Durability: DurabilityConfig{BufferLocation: "tmp/"},
				Security: SecurityConfig{
					EnforceMFA:             true,
					MinPasswordLength:      7,
					MaxPasswordAgeDays:     90,
					MaxPasswordReuse:       5,
					InitialUser:            "admin@auditmon.net",
					InitialPassword:        "Password1234!",
					SessionMaxSeconds:      86400,
					MaxFailedLogins:        5,
					LockoutDurationSeconds: 300,
				},
				Server: ServerConfig{
					RootURL: "http://localhost:7984",
					Port:    7984,
				},
				Metrics: MetricsConfig{
					SendInflux:     false,
					InfluxURL:      "http://localhost:8086",
					InfluxUser:     "admin",
					InfluxPassword: "admin",
					InfluxDB:       "default",
				},
				Mailer: MailerConfig{
					From:    "noreply@localhost.com",
					Backend: "ses",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustReadConfig(tt.args.configFileLocation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
