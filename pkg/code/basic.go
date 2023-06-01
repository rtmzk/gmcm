package code

const (
	ErrSuccess int = iota + 100001
	ErrUnknown
	ErrBind
	ErrValidation
	ErrPageNotFound
	ErrLogFileNotFound
)

const (
	ErrDatabase int = iota + 100100
	ErrSignatureInvalid
	ErrHostConnectionByKey
	ErrHostConnectionByPass
	ErrCephInstall
)
