package email

// UserConfirmTitle -
func UserConfirmTitle() string {
	return "ECode 激活邮件"
}

// UserConfirmTemplete 用户激活模版
func UserConfirmTemplete(id, name, uuid string) string {
	return `
	<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Document</title>
		</head>
		<body>
			<h3> ，欢迎来到 ECode!</h3>
			<h3>点击如下链接来激活你的账号</h3>
			<div><a href=""></a></div>
			<h3>代码很有趣，愿你在这里有美妙的体验</h3>
		</body>
	</html>
	`
}
