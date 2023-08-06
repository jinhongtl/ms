package internal

//总配置
type Config struct {
	Database *Database
	Server   *Server
	Log      *Log
}

//数据库配置
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

//服务器配置
type Server struct {
	Port string
}

//日志
type Log struct {
	Path  string
	Level uint8
}
