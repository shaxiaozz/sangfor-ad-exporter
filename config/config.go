package config

import (
	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/spf13/viper"
	"log"
)

var configFilePaths = []string{
	"/etc/sangfor-ad-exporter", "./config", "../config",
}

func InitConfig() *App {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	log.Printf("将自动从以下路径查找 config.yaml: %v", configFilePaths)

	found := false

	for _, dir := range configFilePaths {
		configFile := dir + "/config.yaml"
		if !fileutil.Exist(configFile) {
			log.Printf("%s 不存在，跳过", configFile)
			continue
		}

		v.AddConfigPath(dir)
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("读取配置文件失败 %s: %v", configFile, err)
		}

		log.Printf("成功加载配置文件: %s", v.ConfigFileUsed())
		found = true
		break
	}

	if !found {
		log.Fatal("未找到任何可用的 config.yaml，程序退出")
	}

	var cfg App
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatal("配置文件解析失败:", err)
	}

	return &cfg
}
