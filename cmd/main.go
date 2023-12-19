package main

import (
	"github.com/Edwinfpirajan/server-smart.git/config"
	"github.com/Edwinfpirajan/server-smart.git/internal/src"
)

func main() {
	src.MainDbConnection()
	config.Environments()
}
