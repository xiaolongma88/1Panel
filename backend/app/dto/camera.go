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
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"cameraList"`
}

type LabelConfig struct {
	LabelList []string `json:"labelList"`
}

type RtspInfo struct {
	RtspAddr string `json:"rtspAddr"`
	CameraID string `json:"cameraID"`
}
