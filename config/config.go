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
	Port              string
	Development       bool
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxConnectionIdle time.Duration
	MaxConnectionAge  time.Duration
	Kafka             Kafka
}

type Http struct {
	Port              string
	PprofPort         string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
}

type Lightstep struct {
	AccessToken string
}

type Sentry struct {
	SentryDSN string
}

// Logger config
type Logger struct {
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type HTTPClient struct {
	InternalURL    string
	XApplicationID *string
	RetryCount     int
	RetryWaitTime  time.Duration
}

type MongoDB struct {
	URI         string
	User        string
	Password    string
	DB          string
	MaxPoolSize uint64
	MinPoolSize uint64
	RetryWrites bool
}

type Postgres struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     uint64
	TimeZone *string
}

type Kafka struct {
	Brokers []string
}

type Redis struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDB string
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

func exportConfig() error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if os.Getenv("MODE") == "DOCKER" {
		viper.SetConfigName("config-docker.yml")
	} else {
		viper.SetConfigName("config.yaml")
	}

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
