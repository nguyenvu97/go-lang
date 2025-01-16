package setting

type Config struct {
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"log"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Server ServerSetting `mapstructure:"server"`
	Kafka  KafkaSetting  `mapstructure:"kafka"`
	JwtData 	JwtSetting 	`mapstructure:"jwt"`
}

type MysqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DbName          string `mapstructure:"dbname"`
	Password        string `mapstructure:"password"`
	UserName        string `mapstructure:"username"`
	MaxIdleCons     int    `mapstructure:"maxIdleCons"`
	MaxOpenCons     int    `mapstructure:"maxOpenCons"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Pool_size int    `mapstructure:"pool_size"`
}
type ServerSetting struct {
	Port  int    `mapstructure:"port"`
	Model string `mapstructure:"model"`
}

type KafkaSetting struct {
	Url      string `mapstructure:"url"`
	Port     int    `mapstructure:"port"`
	Topic    string `mapstructure:"topic"`
	Consumer string `mapstructure:"consumer"`
}
type JwtSetting struct {
	PrivateKey                string `mapstructure:"private_key"`
	PublicKey                 string `mapstructure:"public_key"`
	Expiration                int64  `mapstructure:"expiration"`
	RefreshTokenExpiration    int64  `mapstructure:"refresh_token_expiration"`
}


