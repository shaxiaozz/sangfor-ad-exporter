package config

type App struct {
	SangforAd SangforAd `mapstructure:"sangfor_ad"`
}

type SangforAd struct {
	DeviceName string `mapstructure:"device_name"` // 设备名称
	Username   string `mapstructure:"username"`    // 账号
	Password   string `mapstructure:"password"`    // 密码
	Url        string `mapstructure:"url"`         // API 地址
}
