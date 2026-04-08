package config

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type AliPay struct {
	PrivateKey string
	AppId      string
	NotifyURL  string
	ReturnURL  string
}
type AppConfig struct {
	Mysql
	AliPay
}
