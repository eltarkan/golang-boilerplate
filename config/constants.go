package config

type Environment struct {
	DBName    string `mapstructure:"DB_NAME"`
	RedisAddr string `mapstructure:"REDIS_ADDR"`
}
