package config

type Config struct {
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
	Repo   Repo   `yaml:"repo"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Repo struct {
	MongoDB MongoDB `yaml:"mongodb"`
	Redis   Redis   `yaml:"redis"`
}

type MongoDB struct {
	URL string `yaml:"url"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
}
