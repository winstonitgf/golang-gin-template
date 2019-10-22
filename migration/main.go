package main

import (
	"fmt"
	"os"
	"strings"

	orm "template/main/databases"
	"template/main/models/template"

	"github.com/spf13/viper"
)

// Migration 文件
// http://gorm.io/docs/migration.html
func LoadConfig() {
	if _, err := os.Stat("../config.yml"); os.IsNotExist(err) {
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("../.")
		err := viper.ReadInConfig()
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func main() {

	// 讀取設定檔
	LoadConfig()

	// 開啟連線
	orm.GormOpen()
	defer orm.GormClose()

	// 開始migrate
	orm.Eloquent.AutoMigrate(&template.Template{})

	// 塞資料
	var template template.Template
	_ = orm.Eloquent.Raw("delete from templates").Scan(&template).Error
	_ = orm.Eloquent.Raw(`insert into templates(flex1) values("APP")`).Scan(&template).Error

	return
}
