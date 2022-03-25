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
	RetentionDays int `yaml:"retention_days" envconfig:"DATABASE_RETENTION_DAYS"`
	MaxQueryDays  int `yaml:"max_query_days" envconfig:"DATABASE_MAX_QUERY_DAYS"`
}

type DurabilityConfig struct {
	BufferLocation string `yaml:"buffer_location" envconfig:"DURABILITY_BUFFER_LOCATION"`
}

type SecurityConfig struct {
	EnforceMFA             bool   `yaml:"enforce_mfa" envconfig:"SETTINGS_ENFORCE_MFA"`
	MinPasswordLength      int    `yaml:"min_password_length" envconfig:"SETTINGS_MIN_PASSWORD_LENGTH"`
	MaxPasswordAgeDays     int    `yaml:"max_password_age_days" envconfig:"SETTINGS_MAX_PASSWORD_AGE_DAYS"`
	MaxPasswordReuse       int    `yaml:"max_password_reuse" envconfig:"SETTINGS_MAX_PASSWORD_REUSE"`
	InitialUser            string `yaml:"initial_user" envconfig:"SETTINGS_INITIAL_USER"`
	InitialPassword        string `yaml:"initial_password" envconfig:"SETTINGS_INITIAL_PASSWORD"`
	SessionMaxSeconds      int    `yaml:"session_max_seconds" envconfig:"SETTINGS_MAX_SESSION_SECONDS"`
	MaxFailedLogins        int    `yaml:"max_failed_logins" envconfig:"SETTINGS_MAX_FAILED_LOGINS"`
	LockoutDurationSeconds int    `yaml:"lockout_duration_seconds" envconfig:"SETTINGS_LOCKOUT_DURATION_SECONDS"`
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
	err = envconfig.Process("auditmon", &c)
	if err != nil {
		panic(fmt.Sprintf("unable to read env vars to configure auditmon %v", err.Error()))
	}
	return c
}
