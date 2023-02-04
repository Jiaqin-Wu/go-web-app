package settings

func Init() (err error) {
	// 指定配置文件名（不需要后缀）
	viper.setConfigName("config")
	// 指定配置文件类型
	viper.setConfigType("yaml")
	// 指定查找配置文件路径（这里使用相对路径）

}
