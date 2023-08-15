package system

type ServiceGroup struct {
	InitDBService
	CasbinService
	JwtService
	UserService
	OperationRecordService
}
