package config_test

import (
	"testing"

	"github.com/Edwinfpirajan/server-smart.git/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvironments(t *testing.T) {
	config.Environments()

	assert.Equal(t, 8080, config.Cfg.Server.Port)
}
