package format

import "github.com/BurntSushi/toml"

// Config is used to represent the list of configured Formatters.
type Config struct {
	Formatters map[string]*Formatter `toml:"formatter"`
}

// ReadConfigFile reads from path and unmarshals toml into a Config instance.
func ReadConfigFile(path string) (cfg *Config, err error) {
	_, err = toml.DecodeFile(path, &cfg)
	return
}
