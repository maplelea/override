package config

import (
	"github.com/spf13/viper"
)

var Viper *viper.Viper // 导出 Viper 实例

func InitConfig() error {
	Viper = viper.GetViper()      // 获取全局 Viper 实例
	Viper.SetConfigName("config") // 配置文件名称（无扩展名）
	Viper.SetConfigType("yaml")   // 设置配置文件类型为YAML
	Viper.AddConfigPath(".")      // 添加配置文件路径
	return Viper.ReadInConfig()   // 读取配置信息
}
