package message

// OK 请求正确 code 为 0 =============================================
var OK = ErrMsg{Code: 0, Message: "OK"}

// SErrSystem 系统错误, 前缀为 100  ===================================
type SErrSystem struct {
	Internal ErrMsg
}

// ErrSystem -
var ErrSystem = SErrSystem{
	Internal: ErrMsg{Code: 10001, Message: "内部服务器错误"},
}

// SErrDB 数据库错误, 前缀为 101  ===================================
type SErrDB struct {
}

// SErrToken 认证错误, 前缀是 201 ===================================
type SErrToken struct {
	GenerateFail   ErrMsg
	UpdateFail     ErrMsg
	HeaderIllegal  ErrMsg
	NoToken        ErrMsg
	NoAccess       ErrMsg
	NoManageAccess ErrMsg
	TokenExpired   ErrMsg
	Other          ErrMsg
}

// ErrToken -
var ErrToken = SErrToken{
	GenerateFail:   ErrMsg{Code: 20101, Message: "令牌生成异常"},
	UpdateFail:     ErrMsg{Code: 20102, Message: "令牌更新异常"},
	HeaderIllegal:  ErrMsg{Code: 20103, Message: "请求头 Authorization 不合法"},
	NoToken:        ErrMsg{Code: 20104, Message: "请求未携带令牌，无权限访问"},
	NoAccess:       ErrMsg{Code: 20105, Message: "当前令牌无权限"},
	NoManageAccess: ErrMsg{Code: 20106, Message: "当前令牌无管理员权限"},
	TokenExpired:   ErrMsg{Code: 20107, Message: "令牌已过期"},
	Other:          ErrMsg{Code: 20108, Message: "令牌异常"},
}

// SErrHTTPData 查询数据格式错误 前缀是 301  =========================
type SErrHTTPData struct {
	BindFail ErrMsg
	Illegal  ErrMsg
	AddFail  ErrMsg
	DelFail  ErrMsg
}

// ErrHTTPData -
var ErrHTTPData = SErrHTTPData{
	BindFail: ErrMsg{Code: 30101, Message: "请求数据格式异常"},
	Illegal:  ErrMsg{Code: 30102, Message: "请求数据不合法"},
	AddFail:  ErrMsg{Code: 30103, Message: "添加失败"},
	DelFail:  ErrMsg{Code: 30104, Message: "删除失败"},
}

// SErrUser 用户错误, 前缀为 501  ====================================
type SErrUser struct {
	AddFail           ErrMsg
	DelFail           ErrMsg
	ModFail           ErrMsg
	NotFound          ErrMsg
	LoginFail         ErrMsg
	PasswordIncorrect ErrMsg
	IDIllegal         ErrMsg
	NameExist         ErrMsg
	EmailExist        ErrMsg
	UUIDIllegal       ErrMsg
	ModPasswordFail   ErrMsg
}

// ErrUser -
var ErrUser = SErrUser{
	AddFail:           ErrMsg{Code: 50101, Message: "用户添加异常"},
	DelFail:           ErrMsg{Code: 50102, Message: "用户删除异常"},
	ModFail:           ErrMsg{Code: 50103, Message: "用户修改异常"},
	NotFound:          ErrMsg{Code: 50104, Message: "用户不存在"},
	LoginFail:         ErrMsg{Code: 50105, Message: "用户登录异常"},
	PasswordIncorrect: ErrMsg{Code: 50106, Message: "密码不正确"},
	IDIllegal:         ErrMsg{Code: 50107, Message: "用户ID不合法"},
	NameExist:         ErrMsg{Code: 50120, Message: "用户名已注册"},
	EmailExist:        ErrMsg{Code: 50121, Message: "邮箱已注册"},
	UUIDIllegal:       ErrMsg{Code: 50122, Message: "UUID不合法"},
	ModPasswordFail:   ErrMsg{Code: 50123, Message: "修改密码失败"},
}

// SErrCard 卡片错误, 前缀为 502 ===================================
type SErrCard struct {
	AddFail   ErrMsg
	DelFail   ErrMsg
	ModFail   ErrMsg
	NotFound  ErrMsg
	IDIllegal ErrMsg
}

// ErrCard -
var ErrCard = SErrCard{
	AddFail:   ErrMsg{Code: 50201, Message: "卡片添加异常"},
	DelFail:   ErrMsg{Code: 50202, Message: "卡片删除异常"},
	ModFail:   ErrMsg{Code: 50203, Message: "卡片修改异常"},
	NotFound:  ErrMsg{Code: 50204, Message: "卡片不存在"},
	IDIllegal: ErrMsg{Code: 50207, Message: "卡片ID不合法"},
}

// SErrCategory 卡片错误, 前缀为 503 ===================================
type SErrCategory struct {
	AddFail   ErrMsg
	DelFail   ErrMsg
	ModFail   ErrMsg
	NotFound  ErrMsg
	IDIllegal ErrMsg
}

// ErrCategory -
var ErrCategory = SErrCategory{
	AddFail:   ErrMsg{Code: 50301, Message: "分类添加异常"},
	DelFail:   ErrMsg{Code: 50302, Message: "分类删除异常"},
	ModFail:   ErrMsg{Code: 50303, Message: "分类修改异常"},
	NotFound:  ErrMsg{Code: 50304, Message: "分类不存在"},
	IDIllegal: ErrMsg{Code: 50307, Message: "分类ID不合法"},
}

// SErrTag 卡片错误, 前缀为 504 ===================================
type SErrTag struct {
	AddFail   ErrMsg
	DelFail   ErrMsg
	ModFail   ErrMsg
	NotFound  ErrMsg
	IDIllegal ErrMsg
}

// ErrTag -
var ErrTag = SErrTag{
	AddFail:   ErrMsg{Code: 50401, Message: "标签添加异常"},
	DelFail:   ErrMsg{Code: 50402, Message: "标签删除异常"},
	ModFail:   ErrMsg{Code: 50403, Message: "标签修改异常"},
	NotFound:  ErrMsg{Code: 50404, Message: "标签不存在"},
	IDIllegal: ErrMsg{Code: 50407, Message: "标签ID不合法"},
}

// SErrCardCategory 卡片错误, 前缀为 504 ===================================
type SErrCardCategory struct {
	HasAdd ErrMsg
	HasDel ErrMsg
}

// ErrCardCategory -
var ErrCardCategory = SErrCardCategory{
	HasAdd: ErrMsg{Code: 51101, Message: "已添加"},
	HasDel: ErrMsg{Code: 51102, Message: "已删除"},
}
