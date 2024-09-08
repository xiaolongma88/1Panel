package configs

type ServerConfig struct {
	System    System    `mapstructure:"system"`
	LogConfig LogConfig `mapstructure:"log"`
	DirConfig DirConfig `mapstructure:"dir"`
}
