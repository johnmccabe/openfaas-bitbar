package config

// Config TODO
type Config struct {
	Auths []Auth `yaml:"auths"`
}

// Auth TODO
type Auth struct {
	Gateway string `yaml:"gateway"`
	Auth    string `yaml:"auth,omitempty"`
	Token   string `yaml:"token,omitempty"`
}
