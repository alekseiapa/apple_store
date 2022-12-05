package api

import (
	"os"
	"testing"
	"time"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/stretchr/testify/require"

	"github.com/alekseiapa/apple_store/util"
	"github.com/gin-gonic/gin"

	// postgres driver for Go's database/sql package
	_ "github.com/lib/pq"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
