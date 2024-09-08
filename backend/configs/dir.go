package configs

type DirConfig struct {
	AppDir    string `mapstructure:"app"`
	AppConfig string `mapstructure:"app_config"`
	ResultDir string `mapstructure:"result"`
	AppLogDir string `mapstructure:"app_log"`
	TestVideo string `mapstructure:"test_video"`
}
