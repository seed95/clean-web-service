package config

type (
	Config struct {
		Logger     Logger     `yaml:"logger"`
		Translator Translator `yaml:"translator"`
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

	Translator struct {
		I18N I18N `yaml:"i18n"`
	}

	I18N struct {
		MessagePath string `yaml:"message_path"`
	}
)
