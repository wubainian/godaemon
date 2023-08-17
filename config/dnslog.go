package config

type Dnslog struct {
	Domains []string `mapstructure:"domains" json:"domains" yaml:"domains"` //dns域名列表
	Cname   string   `mapstructure:"cname" json:"cname" yaml:"cname"`       //cname地址
}
