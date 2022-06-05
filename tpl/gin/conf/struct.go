package conf

type ServerCfg struct {
	Mode         string `ini:"mode"`
	Port         string `ini:"port"`
	StaticDir    string `ini:"static_dir"`
	ReadTimeout  int    `ini:"read_timeout"`
	WriteTimeout int    `ini:"write_timeout"`
	SwaggerHost  string `ini:"swagger_host"`
}

type DatabaseCfg struct {
	User          string `ini:"user"`
	Password      string `ini:"password"`
	DBName        string `ini:"dbname"`
	Host          string `ini:"host"`
	Port          int    `ini:"port"`
	MaxIdleConn   int    `ini:"max_idle_conn"`
	MaxOpenConn   int    `ini:"max_open_conn"`
	LogLevel      int    `ini:"log_level"`
	SlowThreshold int    `ini:"threshold"`
	Debug         bool   `ini:"debug"`
	Type          string `ini:"type"`
	URL           string `ini:"url"`
}

type LoggerCfg struct {
	MaxSize   int    `ini:"maxsize"`
	MaxAge    int    `ini:"maxage"`
	MaxBackup int    `ini:"backup"`
	Level     string `ini:"level"`
	Filepath  string `ini:"filepath"`
}
