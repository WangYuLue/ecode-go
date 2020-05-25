package config

// SMysql struct
type SMysql struct {
	UserName string
	Password string
	IP       string
	Port     string
	DBName   string
}

// SEmail struct
type SEmail struct {
	User     string
	Password string // 这里密码不是邮箱密码，是你设置的smtp授权码
	Host     string
}
