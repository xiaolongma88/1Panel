package dto

type CameraContent struct {
	KafkaConfig  `json:"KafkaConfig"`
	CamereConfig `json:"CamereConfig"`
	LabelConfig  `json:"LabelConfig"`
}

type KafkaConfig struct {
	Enable   bool   `json:"enable"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type CamereConfig struct {
	CameraList []struct {
		CamID    string `json:"camID"`
		RtspPath string `json:"rtspPath"`
	} `json:"cameraList"`
}

type LabelConfig struct {
	LabelList []string `json:"labelList"`
}
