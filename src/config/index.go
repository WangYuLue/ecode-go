package config

// SMysql struct
type SMysql struct {
	UserName string
	Password string
	IP       string
	Port     string
	DBName   string
}

// SRedis struct
type SRedis struct {
	Password string
	IP       string
	Port     string
	DB       int
}

// SEmail struct
type SEmail struct {
	User     string
	Password string // 这里密码不是邮箱密码，是你设置的smtp授权码
	Host     string
}

// SEmailConfirmUser struct
type SEmailConfirmUser struct {
	Title      string // 邮件标题
	SuccessURL string // 验证成功后的跳转
	FailURL    string // 验证失败后的跳转
}

// SEmailResetPassword -
type SEmailResetPassword struct {
	ResetURL string
}
