package global

type ErrCode int32

const (
	ErrInvalidReleaseParam ErrCode = 10001
	ErrConfigFileNotExists ErrCode = 10002
	ErrInitZapLoggerFailed ErrCode = 10003
)

func (e ErrCode)String() string {
	switch e {
	case ErrInvalidReleaseParam:
		return "Error: Invalid Release Param"
	case ErrConfigFileNotExists:
		return "Error: Config File Not Exists"
	case ErrInitZapLoggerFailed:
		return "Error: Init Zap Logger Instance Failed"
	default:
		return "Unknown"
	}
}
