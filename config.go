package vkwordle

type Config struct {
	DBConfig *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Database string `env:"DB_NAME,required"`
}
