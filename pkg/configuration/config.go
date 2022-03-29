package configuration

import (
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database   DatabaseConfig   `yaml:"database"`
	Settings   SettingsConfig   `yaml:"settings"`
	Durability DurabilityConfig `yaml:"durability"`
	Security   SecurityConfig   `yaml:"security"`
	Server     ServerConfig     `yaml:"server"`
	Metrics    MetricsConfig    `yaml:"metrics"`
	Mailer     MailerConfig     `yaml:"mailer"`
	Logging    LoggingConfig    `yaml:"logging"`
}

type DatabaseConfig struct {
	Backend  string `yaml:"backend" envconfig:"DATABASE_BACKEND"`
	File     string `yaml:"file" envconfig:"DATABASE_FILE"`
	Host     string `yaml:"host" envconfig:"DATABASE_HOST"`
	Port     int    `yaml:"port" envconfig:"DATABASE_PORT"`
	SSLMode  string `yaml:"ssl_mode" envconfig:"DATABASE_SSL_MODE"`
	Username string `yaml:"username" envconfig:"DATABASE_USERNAME"`
	Password string `yaml:"password" envconfig:"DATABASE_PASSWORD"`
	Name     string `yaml:"name" envconfig:"DATABASE_NAME"`
}

type SettingsConfig struct {
	RetentionDays int `yaml:"retention_days" envconfig:"SETTINGS_RETENTION_DAYS"`
	MaxQueryDays  int `yaml:"max_query_days" envconfig:"SETTINGS_MAX_QUERY_DAYS"`
	MaxResults    int `yaml:"max_results" envconfig:"SETTINGS_MAX_RESULTS"`
}

type DurabilityConfig struct {
	BufferLocation string `yaml:"buffer_location" envconfig:"DURABILITY_BUFFER_LOCATION"`
}

type SecurityConfig struct {
	MinPasswordLength      int    `yaml:"min_password_length" envconfig:"SECURITY_MIN_PASSWORD_LENGTH"`
	MaxPasswordAgeDays     int    `yaml:"max_password_age_days" envconfig:"SECURITY_MAX_PASSWORD_AGE_DAYS"`
	MaxPasswordReuse       int    `yaml:"max_password_reuse" envconfig:"SECURITY_MAX_PASSWORD_REUSE"`
	InitialUser            string `yaml:"initial_user" envconfig:"SECURITY_INITIAL_USER"`
	InitialPassword        string `yaml:"initial_password" envconfig:"SECURITY_INITIAL_PASSWORD"`
	SessionMaxSeconds      int    `yaml:"session_max_seconds" envconfig:"SECURITY_MAX_SESSION_SECONDS"`
	MaxFailedLogins        int    `yaml:"max_failed_logins" envconfig:"SECURITY_MAX_FAILED_LOGINS"`
	LockoutDurationSeconds int    `yaml:"lockout_duration_seconds" envconfig:"SECURITY_LOCKOUT_DURATION_SECONDS"`
	SessionSecret          string `yaml:"session_secret" envconfig:"SECURITY_SESSION_SECRET"`
}

type ServerConfig struct {
	RootURL string `yaml:"root_url" envconfig:"SERVER_ROOT_URL"`
	Port    int    `yaml:"port" envconfig:"SERVER_PORT"`
}

type MetricsConfig struct {
	SendInflux     bool   `yaml:"send_influx" envconfig:"METRICS_SEND_INFLUX"`
	InfluxURL      string `yaml:"influx_url" envconfig:"METRICS_INFLUX_URL"`
	InfluxUser     string `yaml:"influx_user" envconfig:"METRICS_INFLUX_USER"`
	InfluxPassword string `yaml:"influx_password" envconfig:"METRICS_INFLUX_PASSWORD"`
	InfluxDB       string `yaml:"influx_db" envconfig:"METRICS_INFLUX_DB"`
}

type MailerConfig struct {
	From    string `yaml:"from" envconfig:"MAILER_FROM"`
	Backend string `yaml:"backend" envconfig:"MAILER_BACKEND"`
}

type LoggingConfig struct {
	Level string `yaml:"level" envconfig:"LOGGING_LEVEL"`
}

// MustReadConfig reads all configurations from the provided config file and env vars
func MustReadConfig(configFileLocation string) Config {
	c := Config{}
	buf, err := ioutil.ReadFile(configFileLocation)
	if err != nil {
		panic(fmt.Sprintf("unable to open configuration file %v %v", configFileLocation, err.Error()))
	}
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		panic(fmt.Sprintf("unable to parse configuration file %v %v", configFileLocation, err.Error()))

	}
	err = envconfig.Process("auditmon", &c.Database)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Server)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Security)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Mailer)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Metrics)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Settings)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Durability)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	err = envconfig.Process("auditmon", &c.Logging)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	return c
}
