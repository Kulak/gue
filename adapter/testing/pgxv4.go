package testing

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"

	"github.com/kulak/gue/v2/adapter"
	"github.com/kulak/gue/v2/adapter/pgxv4"
)

// OpenTestPoolMaxConnsPGXv4 opens connections pool user in testing
func OpenTestPoolMaxConnsPGXv4(t testing.TB, maxConnections int32, schema string) adapter.ConnPool {
	t.Helper()

	applyMigrations(schema).Do(func() {
		doApplyMigrations(t, schema)
	})

	connPoolConfig, err := pgxpool.ParseConfig(testConnDSN(t))
	require.NoError(t, err)

	connPoolConfig.MaxConns = maxConnections

	poolPGXv4, err := pgxpool.ConnectConfig(context.Background(), connPoolConfig)
	require.NoError(t, err)

	pool := pgxv4.NewConnPool(poolPGXv4)

	t.Cleanup(func() {
		truncateAndClose(t, pool, schema)
	})

	return pool
}

// OpenTestPoolPGXv4 opens connections pool user in testing
func OpenTestPoolPGXv4(t testing.TB) adapter.ConnPool {
	t.Helper()

	return OpenTestPoolMaxConnsPGXv4(t, defaultPoolConns, defaultSchema)
}
