package database

import (
	"github.com/stretchr/testify/require"
	"simple-ecommerce/cmd/internal/config"
	"testing"
)

func init() {
	filename := "../../api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
}

func TestConnectPostgres(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})
}
