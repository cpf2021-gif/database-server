package config

type App struct {
	Name       string `mapstructure:"name" json:"name" yaml:"name"`
	DBtype     string `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"`
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	BackupTime string `mapstructure:"backuptime" json:"backuptime" yaml:"backuptime"`
}
