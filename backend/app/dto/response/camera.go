package response

type KafkaConfig struct {
	Enable   bool   `json:"enable"`
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type CameraConfig struct {
	CameraList []struct {
		CamID    string `json:"camID"`
		RtspPath string `json:"rtspPath"`
		Width    string `json:"width"`
		Height   string `json:"height"`
	} `json:"cameraList"`
}

type LabelConfig struct {
	LabelList []string `json:"labelList"`
}

type Config struct {
	KafkaConfig  KafkaConfig  `json:"KafkaConfig"`
	CamereConfig CameraConfig `json:"CamereConfig"`
	LabelConfig  LabelConfig  `json:"LabelConfig"`
}
