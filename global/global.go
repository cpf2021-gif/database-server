package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"server/config"
)

var (
	GL_VIPER  *viper.Viper
	GL_CONFIG config.Configuration
	GL_DB     *gorm.DB
)
