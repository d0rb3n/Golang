package modules

import "time"

type MySQLConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBName      string
	ExecTimeout time.Duration
}
