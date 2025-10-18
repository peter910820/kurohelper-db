package models

type InitConfig struct {
	DBOwner    string
	DBPassword string
	DBPort     int
	DBConfig   []DBConfig
}

type DBConfig struct {
	DBName string
}
