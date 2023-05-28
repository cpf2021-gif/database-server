package util

import (
	"os"
	"os/exec"
	"time"

	"server/global"

	"github.com/spf13/viper"
)

func Backup(viper *viper.Viper) error {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	nowtime := time.Now().In(loc).Format("2006-01-02")

	global.GL_VIPER.Set("database.backuptime", nowtime)
	global.GL_CONFIG.App.BackupTime = nowtime

	if err := global.GL_VIPER.WriteConfig(); err != nil {
		return err
	}

	// 备份数据库
	//docker exec -i some-postgres bash -c "pg_dump -U postgres -d inventory_management_system > backup.sql"
	cmd := exec.Command("docker", "exec", "-i", "some-postgres", "bash", "-c", "pg_dump -U postgres -d inventory_management_system > backup.sql")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
