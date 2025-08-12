/*
 * @Author: 12866449444136360 liangsz@aliyun.com
 * @Date: 2025-08-12 14:00:46
 * @LastEditors: 12866449444136360 liangsz@aliyun.com
 * @LastEditTime: 2025-08-12 14:02:41
 * @FilePath: \自学项目\go-backend\internal\config\config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string `mapstructure:"ENV"`
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// 设置默认值
	if AppConfig.ServerAddr == "" {
		AppConfig.ServerAddr = ":8080"
	}

	if AppConfig.Env == "" {
		AppConfig.Env = "development"
	}
}
