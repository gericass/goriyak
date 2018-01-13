package setting

import "github.com/BurntSushi/toml"

type Server struct {
	IP    string
	Name  string
	Mysql MySQL
}

type MySQL struct {
	User     string
	Password string
	Host     string
	Port     string
}

var ServerConfig Server

func Setting() error {
	if _, err := toml.DecodeFile("../config.toml", &ServerConfig); err != nil {
		return err
	}
	return nil
}
