package config

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type App struct {
	Name  string `mapstructure:"name" json:"name" yaml:"name"`
	Port  string `mapstructure:"port" json:"port" yaml:"port"`
	Token Token  `mapstructure:"token" json:"token" yaml:"token"`
}

type Token struct {
	Expiration         int64  `mapstructure:"expiration" json:"expiration" yaml:"expiration"`
	Secret             string `mapstructure:"secret" json:"secret" yaml:"secret"`
	RefreshGracePeriod int64  `mapstructure:"refresh_grace_period" json:"refresh_grace_period" yaml:"refresh_grace_period"` // token 自动刷新宽限时间（秒）

}

type Database struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	DbName       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns int    `json:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Redis struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}
