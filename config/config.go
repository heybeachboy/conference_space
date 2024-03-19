package config

type Mode string

const (
	DevMode    Mode = `develop`
	TestMode   Mode = `test`
	ProdMode   Mode = `prod`
	SqliteName      = `ImDb`
)

type config struct {
	Env        string `yaml:"Env"`
	Debug      bool   `yaml:"Debug"`
	ServerAddr string `yaml:"ServerAddr"`
	ServerPort int    `yaml:"ServerPort"`
	MysqlDb    struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		DbName   string `yaml:"DbName"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		SetLog   bool   `yaml:"SetLog"`
	} `yaml:"Mysql"`
	Token struct {
		ShareSecret string `yaml:"ShareSecret"`
		LifeExpire  int64  `yaml:"LifeExpire"`
	} `yaml:"Token"`
	AgoraRtc struct {
		AppId          string `yaml:"AppId"`
		AppCertificate string `yaml:"AppCertificate"`
	} `yaml:"AgoraRtc"`
}

func (c *config) IsDebug() bool {
	return c.Debug
}

func (c *config) IsDevelopEnv() bool {
	return Mode(c.Env) == DevMode
}

func (c *config) IsTestEnv() bool {
	return Mode(c.Env) == TestMode
}

func (c *config) IsProdEnv() bool {
	return Mode(c.Env) == ProdMode
}
func (c *config) GetCurrentEnv() Mode {
	return Mode(c.Env)
}
