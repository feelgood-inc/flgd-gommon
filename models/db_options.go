package models

type DBOptions struct {
	Host              *string `json:"host"`
	Port              *int    `json:"port"`
	User              *string `json:"user"`
	Password          *string `json:"password"`
	DBName            string  `json:"db_name"`
	URI               *string `json:"uri"`
	WithMonitor       *bool   `json:"with_monitor"`
	ConnectionTimeout *int    `json:"connection_timeout"`
	SocketTimeout     *int    `json:"socket_timeout"`
	MaxPoolSize       *uint64 `json:"max_pool_size"`
	MinPoolSize       *uint64 `json:"min_pool_size"`
	Timeout           int     `json:"timeout"`
	RetryReads        *bool   `json:"retry_reads"`
	RetryWrites       *bool   `json:"retry_writes"`
	TimeZone          *string `json:"time_zone"`
}
