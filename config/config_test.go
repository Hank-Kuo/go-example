package config

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestGetConf(t *testing.T) {
	t.Run("Config.getConf", func(t *testing.T) {
		os.Setenv("HOST", "testhost")
		os.Setenv("PORT", "9000")
		os.Setenv("DEBUG", "true")
		os.Setenv("POSTGRES_HOST", "testhost")
		os.Setenv("POSTGRES_USER", "testpostgres")
		os.Setenv("POSTGRES_PASSWORD", "testpostgres")
		os.Setenv("POSTGRES_DATABASE", "testdb")
		os.Setenv("SSL_MODE", "true")
		os.Setenv("JAEGER_HOST", "testhost:14268/api/traces")

		cfg, err := GetConf()

		assert.NoError(t, err)

		assert.Equal(t, "testhost", cfg.Server.Host)
		assert.Equal(t, "testhost", cfg.Database.Host)
		assert.Equal(t, "testdb", cfg.Database.Db)
	})

}
