package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainDbConnection(t *testing.T) {
	// Configurar variables de entorno de prueba
	os.Setenv("MAIN_DB_HOST", "localhost")
	os.Setenv("MAIN_DB_USER", "postgres")
	os.Setenv("MAIN_DB_PASSWORD", "1234")
	os.Setenv("MAIN_DB_NAME", "postgres")
	os.Setenv("MAIN_DB_PORT", "5432")

	// Asegurarse de que las variables de entorno se limpien después de la prueba
	defer func() {
		os.Unsetenv("MAIN_DB_HOST")
		os.Unsetenv("MAIN_DB_USER")
		os.Unsetenv("MAIN_DB_PASSWORD")
		os.Unsetenv("MAIN_DB_NAME")
		os.Unsetenv("MAIN_DB_PORT")
	}()

	// Llamar a la función que se va a probar
	MainDbConnection()

	// Verificar que la conexión a la base de datos se estableció correctamente
	assert.NotNil(t, DB, "DB should not be nil after connecting to the database")
}
