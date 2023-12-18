package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	// Configure test environment variables
	originalDir, err := os.Getwd()
	assert.NoError(t, err, "Error getting current directory")

	// Change working directory
	err = os.Chdir("/home/fernando/Documents/SMART/server")
	assert.NoError(t, err, "Error changing working directory")

	// Set environment variables
	defer func() {
		os.Unsetenv("MAIN_DB_HOST")
		os.Unsetenv("MAIN_DB_USERR")
		os.Unsetenv("MAIN_DB_PASSWORD")
		os.Unsetenv("MAIN_DB_NAME")
		os.Unsetenv("MAIN_DB_PORT")

		// Restore working directory
		err := os.Chdir(originalDir)
		assert.NoError(t, err, "Error restoring working directory")
	}()

	// Call the function to test
	err = loadEnv()
	assert.NoError(t, err, "Expected no error loading .env file")

	// Verify that the environment variables were set correctly
	assert.Equal(t, "localhost", os.Getenv("MAIN_DB_HOST"), "Variable MAIN_DB_HOST should be set")
	assert.Equal(t, "postgres", os.Getenv("MAIN_DB_USER"), "Variable MAIN_DB_USER should be set")
	assert.Equal(t, "1234", os.Getenv("MAIN_DB_PASSWORD"), "Variable MAIN_DB_PASSWORD should be set")
	assert.Equal(t, "postgres", os.Getenv("MAIN_DB_NAME"), "Variable MAIN_DB_NAME should be set")
	assert.Equal(t, "5432", os.Getenv("MAIN_DB_PORT"), "Variable MAIN_DB_PORT should be set")
}

//! Keep in mind that every time environment variables are added they must be specified!
