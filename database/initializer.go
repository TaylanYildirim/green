package database

import (
	"green/config"
)

func InitDB() {
	config.Load()
	Init()
}
