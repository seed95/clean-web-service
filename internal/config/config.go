package config

type (
	Config struct {
		Logger Logger `yaml:"logger"`
	}

	Logger struct {
		Logrus Logrus `yaml:"logrus"`
	}

	Logrus struct {
		Path         string `yaml:"internal_path"`
		Pattern      string `yaml:"filename_pattern"`
		RotationSize string `yaml:"max_size"`
		RotationTime string `yaml:"rotation_time"`
		MaxAge       string `yaml:"max_age"`
	}
)
