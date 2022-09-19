package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	HttpPort = "HTTP_PORT"
)

// Config of application
type Config struct {
	AppVersion      string
	Env             string
	ServiceName     string
	ServiceMainPath string
	Server          Server
	Logger          Logger
	MongoDB         MongoDB
	Kafka           Kafka
	Http            Http
	Redis           Redis
	Sentry          Sentry
	Lightstep       Lightstep
	Postgres        Postgres
	HTTPClient      HTTPClient
}

// Server config
type Server struct {
	Port              string `mapstructure:"PORT"`
	Development       bool   `mapstructure:"DEVELOPMENT"`
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
	Kafka             Kafka
}

type Http struct {
	Port              string `mapstructure:"HTTP_PORT"`
	PprofPort         string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
}

type Lightstep struct {
	AccessToken string `mapstructure:"LIGHTSTEP_ACCESS_TOKEN"`
}

type Sentry struct {
	SentryDSN        string  `mapstructure:"SENTRY_DSN"`
	TracesSampleRate float64 `mapstructure:"SENTRY_TRACES_SAMPLE_RATE"`
}

// Logger config
type Logger struct {
	DisableCaller     bool   `mapstructure:"LOGGER_DISABLE_CALLER"`
	DisableStacktrace bool   `mapstructure:"LOGGER_DISABLE_STACKTRACE"`
	Encoding          string `mapstructure:"LOGGER_ENCODING"`
	Level             string `mapstructure:"LOGGER_LEVEL"`
}

type HTTPClient struct {
	InternalURL    string  `mapstructure:"INTERNAL_URL"`
	XApplicationID *string `mapstructure:"X_APPLICATION_ID"`
	RetryCount     int     `mapstructure:"RETRY_COUNT"`
	RetryWaitTime  time.Duration
}

type MongoDB struct {
	URI         string `mapstructure:"MONGODB_URI"`
	User        string `mapstructure:"MONGODB_USER"`
	Password    string `mapstructure:"MONGODB_PASSWORD"`
	DB          string `mapstructure:"MONGODB_DB"`
	MaxPoolSize uint64
	MinPoolSize uint64
	RetryWrites bool
}

type Postgres struct {
	Host     string  `mapstructure:"POSTGRES_HOST"`
	User     string  `mapstructure:"POSTGRES_USER"`
	Password string  `mapstructure:"POSTGRES_PASSWORD"`
	DB       string  `mapstructure:"POSTGRES_DB"`
	Port     uint64  `mapstructure:"POSTGRES_PORT"`
	TimeZone *string `mapstructure:"POSTGRES_TIMEZONE"`
}

type Kafka struct {
	Brokers  []string `mapstructure:"KAFKA_BROKERS"`
	Username string   `mapstructure:"KAFKA_USERNAME"`
}

type Redis struct {
	RedisAddr      string  `mapstructure:"REDIS_ADDR"`
	RedisPassword  *string `mapstructure:"REDIS_PASSWORD"`
	RedisDB        string  `mapstructure:"REDIS_DB"`
	RedisDefaultDB string  `mapstructure:"REDIS_DEFAULT_DB"`
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    time.Duration
	DB             int
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	switch os.Getenv("MODE") {
	case "LOCAL":
		viper.SetConfigName("local")
	case "DEV":
		viper.SetConfigName("dev")
	case "prod":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("local")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// ParseConfig Parse config file
func ParseConfig() (*Config, error) {
	if err := exportConfig(); err != nil {
		return nil, err
	}

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	httpPort := os.Getenv(HttpPort)
	if httpPort != "" {
		c.Http.Port = httpPort
	}

	return &c, nil
}
