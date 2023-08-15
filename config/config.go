package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Mssql  Mssql           `mapstructure:"mssql" json:"mssql" yaml:"mssql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Oracle Oracle          `mapstructure:"oracle" json:"oracle" yaml:"oracle"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	Redis  Redis           `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT    JWT             `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Local  Local           `mapstructure:"local" json:"local" yaml:"local"`
	Email  Email           `mapstructure:"email" json:"email" yaml:"email"`
	// 跨域配置
	Cors  CORS  `mapstructure:"cors" json:"cors" yaml:"cors"`
	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
}
