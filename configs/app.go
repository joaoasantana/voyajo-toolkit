package configs

// App represents the application configuration.
// It includes metadata about the application such as its name, version, description, author, and environment.
type App struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
	Environment string `yaml:"environment"`
}
