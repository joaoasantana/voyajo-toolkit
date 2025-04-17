package configs

// Database represents the database configuration.
// It includes details required to connect to a database, such as driver, host, port, name, username, and password.
type Database struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
