package mux

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"bnk.to/core/tools/auth"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/postgres"
	"github.com/xo/dburl"
	"google.golang.org/grpc"
)

// ContextKey is a type that is used as a context key for associating data to a request.
type ContextKey uint8

const (
	// StorageKey associates a db.Storage to the request.
	StorageKey ContextKey = iota
)

// Storage returns the db.Storage associated with the context and nil if there isn't one.
func Storage(ctx context.Context) db.Storage {
	val := ctx.Value(StorageKey)
	storage, ok := val.(db.Storage)
	if !ok {
		return nil
	}
	return storage
}

// OpenDB opens the database specified by the URL.
func OpenDB(database string, authBackend auth.Auth) (db.Storage, error) {
	parsedURL, err := dburl.Parse(database)
	if err != nil {
		return nil, err
	}
	db, err := dburl.Open(database)
	if err != nil {
		return nil, err
	}
	switch parsedURL.Driver {
	case "postgres":
		return postgres.New(db), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", parsedURL.Driver)
	}
}

// Interceptor returns a unary server interceptor that associates a db.Storage to the request.
func Interceptor(databaseURL string, authBackend auth.Auth) grpc.UnaryServerInterceptor {
	connCache := make(map[string]db.Storage)
	mutex := new(sync.RWMutex)

	dbForCtx := func(ctx context.Context) (db.Storage, error) {
		instanceName, err := authBackend.InstanceName(ctx)
		if err != nil {
			return nil, err
		}
		instanceURL := strings.ReplaceAll(databaseURL, "{instance}", instanceName)
		// Fast path. Check if a connection is already established using only a
		// read lock.
		mutex.RLock()
		if storage, ok := connCache[instanceURL]; ok {
			mutex.RUnlock()
			return storage, nil
		}
		mutex.RUnlock()
		// Slow path. Check again if a connection has been established by
		// another lock holder before this goroutine successfully acquire the
		// lock. If there is still no connections, establish the connection.
		mutex.Lock()
		defer mutex.Unlock()
		if storage, ok := connCache[instanceURL]; ok {
			return storage, nil
		}
		storage, err := OpenDB(instanceURL, authBackend)
		if err != nil {
			return nil, err
		}
		connCache[instanceURL] = storage
		return storage, nil
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		storage, err := dbForCtx(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, StorageKey, storage)
		return handler(ctx, req)
	}
}
