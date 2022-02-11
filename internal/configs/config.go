package configs

type ConfigImp interface {
	GetServerHost() string
	GetDatabaseHost() string
	GetServerPort() string
	GetDatabasePort() string
	GetDatabaseName() string
	GetDatabaseUser() string
	GetDatabasePassword() string
}
