package settings

type Config struct {
	Server ServerSettings `mapstructure:"server"`
	Mysql  MySQLSettings  `mapstructure:"mysql"`
	Logger LoggerSetting  `mapstructure:"logger"`
	Redis  RedisSetting   `mapstructure:"redis"`
	Smtp   SMTPConfig     `mapstructure:"smtp"`
	JWT    JWTConfig      `mapstructure:"jwt"`
}

type ServerSettings struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MySQLSettings struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	DbName          string `mapstructure:"dbName"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	TokenHourLifespan string `mapstructure:"token_hour_lifespan"`
	JwtExpiration     string `mapstructure:"jwt_expiration"`
	ApiSecret         string `mapstructure:"api_secret"`
}
