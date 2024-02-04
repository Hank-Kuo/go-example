package database

import (
	"testing"

	"github.com/Hank-Kuo/go-example/config"
	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	t.Run("PKG.Database.sqlite3", func(t *testing.T) {
		sqlite3Cfg := &config.DatabaseConfig{Adapter: "sqlite3", Host: "test.db"}
		db, err := ConnectDB(sqlite3Cfg)
		assert.NoError(t, err)
		assert.NotNil(t, db)
		db.Close()
	})
	t.Run("PKG.Database.mysql", func(t *testing.T) {
		mysqlCfg := &config.DatabaseConfig{Adapter: "mysql", Host: "localhost", Username: "root", Password: "password", Port: 3306, Db: "test"}
		db, err := ConnectDB(mysqlCfg)
		assert.NoError(t, err)
		assert.NotNil(t, db)
		db.Close()
	})
	t.Run("PKG.Database.postgres", func(t *testing.T) {
		postgresCfg := &config.DatabaseConfig{Adapter: "postgres", Host: "localhost", Username: "user", Password: "password", Port: 5432, Db: "test"}
		db, err := ConnectDB(postgresCfg)
		assert.NoError(t, err)
		assert.NotNil(t, db)
		db.Close()
	})
	t.Run("PKG.Database.invalid", func(t *testing.T) {
		invalidCfg := &config.DatabaseConfig{Adapter: "invalid", Host: "localhost"}
		db, err := ConnectDB(invalidCfg)
		assert.Error(t, err)
		assert.Nil(t, db)
	})

}
