package code

func init() {
	register(ErrBind, 400, "Error bind request body")
	register(ErrSuccess, 200, "OK")
	register(ErrUnknown, 500, "Unknown")
	register(ErrValidation, 400, "Validation failed")
	register(ErrDatabase, 500, "Database Error")
	register(ErrSignatureInvalid, 400, "Authorization failed")
	register(ErrPageNotFound, 404, "Page not found")
	register(ErrCephInstall, 500, "Ceph install failed")
	register(ErrHostConnectionByPass, 200, "无法与目标主机建立链接,请检查密码、端口等配置")
	register(ErrHostConnectionByKey, 200, "无法与目标主机建立链接,请检查ssh免密配置,并将秘钥文件命名为id_rsa存放到用户家目录下的.ssh下")
	register(ErrLogFileNotFound, 200, "无法找到安装日志,请先初始化存储")
}
