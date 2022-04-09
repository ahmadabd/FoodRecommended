package configs

type ConfigImp interface {
	GetDatabaseHost() string
	GetServerPort() string
	GetDatabasePort() string
	GetDatabaseName() string
	GetDatabaseUser() string
	GetDatabasePassword() string
}
