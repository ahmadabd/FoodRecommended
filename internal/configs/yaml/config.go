package yaml

func (c *Config) GetDatabaseHost() string {
	return c.Database.Host
}

func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func (c *Config) GetDatabasePort() string {
	return c.Database.Port
}

func (c *Config) GetDatabaseName() string {
	return c.Database.Name
}

func (c *Config) GetDatabaseUser() string {
	return c.Database.User
}

func (c *Config) GetDatabasePassword() string {
	return c.Database.Password
}
