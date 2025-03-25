package config

type HTTPServer struct {
	Address string
}

// struct tags
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yml:"http_server"`
}