package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Edwinfpirajan/server-smart.git/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvironments_whenEnvAllSet(t *testing.T) {

	config.Environments()

	assert.Equal(t, 8080, config.Cfg.Server.Port)

}

func TestEnvironments_whenFileReadError(t *testing.T) {

	os.Rename("../.env", "../.env.bkp")
	defer os.Rename("../.env.bkp", "../.env")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			assert.Contains(
				t, r, "Error can't loaded .env file",
			)
		}
	}()
	config.Environments()

}
