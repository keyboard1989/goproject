package models

type Server struct {
	Addr string `mapstructure:"addr" validate:"required"`
	Port string `mapstructure:"port" validate:"required"`
}

type DB struct {
	DbSource string `mapstructure:"db_source"`
}

type Statis struct {
	ItemLength int    `mapstructure:"item_length" validate:"required"`
	DataLength int    `mapstructure:"data_length" validate:"required"`
	TimeFormat string `mapstructure:"time_format" validate:"required"`
}

type Cache struct {
	RedisSource string `mapstructure:"redis_source"`
}

type HTTPClient struct {
	TimeOut             uint64 `mapstructure:"time_out" validate:"required"`
	MaxIdleConnsPerHost int    `mapstructure:"max_idle_conns_per_host" validate:"required"`
}

type Configuration struct {
	GoVersion  string     `mapstructure:"go_version"`
	Server     Server     `mapstructure:"server" validate:"required"`
	DB         DB         `mapstructure:"db" validate:"required"`
	Statis     Statis     `mapstructure:"statis" validate:"required"`
	Cache      Cache      `mapstructure:"cache" validate:"required"`
	HTTPClient HTTPClient `mapstructure:"http_client" validate:"required"`
}
