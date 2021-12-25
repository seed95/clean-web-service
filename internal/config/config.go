package config

type (
	Config struct {
		Logger     Logger     `yaml:"logger"`
		Translator Translator `yaml:"translator"`
		Database   Database   `yaml:"database"`
		Validation Validation `yaml:"validation"`
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
		I18n I18n `yaml:"i18n"`
	}

	I18n struct {
		MessagePath string `yaml:"message_path"`
	}

	Database struct {
		Postgres Postgres `yaml:"postgres"`
	}

	Postgres struct {
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		DBName    string `yaml:"db_name"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		SSLMode   string `yaml:"ssl_mode"`
		TimeZone  string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
		Migration bool   `yaml:"migration"`
	}

	Validation struct {
		UsernameMinLength  int `yaml:"username_min_length"`
		UsernameMaxLength  int `yaml:"username_max_length"`
		PasswordMinLetters int `yaml:"password_min_letters"`
	}
)
