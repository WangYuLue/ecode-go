package message

// OK 请求正确 code 为 0 =============================================
var OK = &Errno{Code: 0, Message: "OK"}

// SErrSystem 系统错误, 前缀为 100  ===================================
type SErrSystem struct {
	Internal *Errno
}

// ErrSystem -
var ErrSystem = SErrSystem{
	Internal: &Errno{Code: 10001, Message: "内部服务器错误"},
}

// SErrDB 数据库错误, 前缀为 101  ===================================
type SErrDB struct {
}

// SErrToken 认证错误, 前缀是 201 ===================================
type SErrToken struct {
	GenerateFail   *Errno
	UpdateFail     *Errno
	HeaderIllegal  *Errno
	NoToken        *Errno
	NoAccess       *Errno
	NoManageAccess *Errno
	TokenExpired   *Errno
	Other          *Errno
}

// ErrToken -
var ErrToken = SErrToken{
	GenerateFail:   &Errno{Code: 20101, Message: "令牌生成异常"},
	UpdateFail:     &Errno{Code: 20102, Message: "令牌更新异常"},
	HeaderIllegal:  &Errno{Code: 20103, Message: "请求头 Authorization 不合法"},
	NoToken:        &Errno{Code: 20104, Message: "请求未携带令牌，无权限访问"},
	NoAccess:       &Errno{Code: 20105, Message: "当前令牌无权限"},
	NoManageAccess: &Errno{Code: 20106, Message: "当前令牌无管理员权限"},
	TokenExpired:   &Errno{Code: 20107, Message: "令牌已过期"},
	Other:          &Errno{Code: 20108, Message: "令牌异常"},
}

// SErrHTTPData 查询数据格式错误 前缀是 301  =========================
type SErrHTTPData struct {
	BindFail *Errno
	Illegal  *Errno
	AddFail  *Errno
	DelFail  *Errno
}

// ErrHTTPData -
var ErrHTTPData = SErrHTTPData{
	BindFail: &Errno{Code: 30101, Message: "请求数据格式异常"},
	Illegal:  &Errno{Code: 30102, Message: "请求数据不合法"},
	AddFail:  &Errno{Code: 30103, Message: "添加失败"},
	DelFail:  &Errno{Code: 30104, Message: "删除失败"},
}

// SErrUser 用户错误, 前缀为 501  ====================================
type SErrUser struct {
	AddFail           *Errno
	DelFail           *Errno
	ModFail           *Errno
	NotFound          *Errno
	LoginFail         *Errno
	PasswordIncorrect *Errno
	IDIllegal         *Errno
	NameExist         *Errno
	EmailExist        *Errno
	UUIDIllegal       *Errno
	ModPasswordFail   *Errno
}

// ErrUser -
var ErrUser = SErrUser{
	AddFail:           &Errno{Code: 50101, Message: "用户添加异常"},
	DelFail:           &Errno{Code: 50102, Message: "用户删除异常"},
	ModFail:           &Errno{Code: 50103, Message: "用户修改异常"},
	NotFound:          &Errno{Code: 50104, Message: "用户不存在"},
	LoginFail:         &Errno{Code: 50105, Message: "用户登录异常"},
	PasswordIncorrect: &Errno{Code: 50106, Message: "密码不正确"},
	IDIllegal:         &Errno{Code: 50107, Message: "用户ID不合法"},
	NameExist:         &Errno{Code: 50120, Message: "用户名已注册"},
	EmailExist:        &Errno{Code: 50121, Message: "邮箱已注册"},
	UUIDIllegal:       &Errno{Code: 50122, Message: "UUID不合法"},
	ModPasswordFail:   &Errno{Code: 50123, Message: "修改密码失败"},
}

// SErrCard 卡片错误, 前缀为 502 ===================================
type SErrCard struct {
	AddFail   *Errno
	DelFail   *Errno
	ModFail   *Errno
	NotFound  *Errno
	IDIllegal *Errno
}

// ErrCard -
var ErrCard = SErrCard{
	AddFail:   &Errno{Code: 50201, Message: "卡片添加异常"},
	DelFail:   &Errno{Code: 50202, Message: "卡片删除异常"},
	ModFail:   &Errno{Code: 50203, Message: "卡片修改异常"},
	NotFound:  &Errno{Code: 50204, Message: "卡片不存在"},
	IDIllegal: &Errno{Code: 50207, Message: "卡片ID不合法"},
}

// SErrCategory 卡片错误, 前缀为 503 ===================================
type SErrCategory struct {
	AddFail   *Errno
	DelFail   *Errno
	ModFail   *Errno
	NotFound  *Errno
	IDIllegal *Errno
}

// ErrCategory -
var ErrCategory = SErrCategory{
	AddFail:   &Errno{Code: 50301, Message: "分类添加异常"},
	DelFail:   &Errno{Code: 50302, Message: "分类删除异常"},
	ModFail:   &Errno{Code: 50303, Message: "分类修改异常"},
	NotFound:  &Errno{Code: 50304, Message: "分类不存在"},
	IDIllegal: &Errno{Code: 50307, Message: "分类ID不合法"},
}

// SErrTag 卡片错误, 前缀为 504 ===================================
type SErrTag struct {
	AddFail   *Errno
	DelFail   *Errno
	ModFail   *Errno
	NotFound  *Errno
	IDIllegal *Errno
}

// ErrTag -
var ErrTag = SErrTag{
	AddFail:   &Errno{Code: 50401, Message: "标签添加异常"},
	DelFail:   &Errno{Code: 50402, Message: "标签删除异常"},
	ModFail:   &Errno{Code: 50403, Message: "标签修改异常"},
	NotFound:  &Errno{Code: 50404, Message: "标签不存在"},
	IDIllegal: &Errno{Code: 50407, Message: "标签ID不合法"},
}

// SErrCardCategory 卡片错误, 前缀为 504 ===================================
type SErrCardCategory struct {
	HasAdd *Errno
	HasDel *Errno
}

// ErrCardCategory -
var ErrCardCategory = SErrCardCategory{
	HasAdd: &Errno{Code: 51101, Message: "已添加"},
	HasDel: &Errno{Code: 51102, Message: "已删除"},
}
