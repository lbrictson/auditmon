package configuration

import (
	"os"
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
					SSLMode:  "disable",
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
					MinPasswordLength:      7,
					MaxPasswordAgeDays:     90,
					MaxPasswordReuse:       5,
					InitialUser:            "admin",
					InitialPassword:        "Password1234!",
					SessionMaxSeconds:      86400,
					MaxFailedLogins:        5,
					LockoutDurationSeconds: 300,
					SessionSecret:          "thisIsNotSafePleaseChangeIt",
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
				Logging: LoggingConfig{Level: "info"},
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

func TestMustReadConfigWithEnvOverrides(t *testing.T) {
	// Setup env overrides to validate they work as expected
	os.Setenv("AUDITMON_DATABASE_BACKEND", "postgres")
	os.Setenv("AUDITMON_SECURITY_ENFORCE_MFA", "false")
	os.Setenv("AUDITMON_SERVER_PORT", "5555")
	os.Setenv("AUDITMON_LOGGING_LEVEL", "error")
	type args struct {
		configFileLocation string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Read config from config/auditmon.yaml with some overrides",
			args: args{configFileLocation: "../../config/auditmon.yaml"},
			want: Config{
				Database: DatabaseConfig{
					Backend:  "postgres",
					File:     "auditmon.db",
					Host:     "localhost",
					Port:     5432,
					SSLMode:  "disable",
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
					MinPasswordLength:      7,
					MaxPasswordAgeDays:     90,
					MaxPasswordReuse:       5,
					InitialUser:            "admin",
					InitialPassword:        "Password1234!",
					SessionMaxSeconds:      86400,
					MaxFailedLogins:        5,
					LockoutDurationSeconds: 300,
					SessionSecret:          "thisIsNotSafePleaseChangeIt",
				},
				Server: ServerConfig{
					RootURL: "http://localhost:7984",
					Port:    5555,
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
				Logging: LoggingConfig{Level: "error"},
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
