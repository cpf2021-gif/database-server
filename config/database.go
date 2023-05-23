package config

type Database struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	SSLMode  string `mapstructure:"sslmode" json:"sslmode" yaml:"sslmode"`
	Max_idle_conns int `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	Max_open_conns int `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
}