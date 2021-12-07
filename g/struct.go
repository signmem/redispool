package g

type GlobalConfig struct {
	Debug         bool              `json:"debug"`
	ROLE 			string 			`json:'role'`
	LogLevel   		string			`json:"loglevel"`
	LogFile 		string			`json:"logfile"`
	Http 			*HttpConfig 	`json:"http"`
	Redis 			*RedisConfig 	`json:"redis"`
	TestLine 		int				`json:"testline"`
}

type HttpConfig struct {
	Enabled 		bool 		`json:"enabled"`
	Listen 			string 		`json:"listen"`
}

type RedisConfig struct {
	Server 			string 		`json:"server"`
	Password 		string 		`json:"password"`
	DefaultDB 		int 		`json:"defaultdb"`
}