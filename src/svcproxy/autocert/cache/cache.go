package cache

import (
	"database/sql"
	"fmt"
	"strconv"

	"golang.org/x/crypto/acme/autocert"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"

	// PostgreSQL driver
	_ "github.com/lib/pq"

	sqlcache "svcproxy/autocert/cache/sql"
)

// NewCacheFactory returns Cache instance
func NewCacheFactory(backend string, options map[string]string) (autocert.Cache, error) {
	switch backend {
	case "sql":
		return newSQLCacheBackend(options)
	}

	return nil, fmt.Errorf("Unknown backend specified")
}

func newSQLCacheBackend(options map[string]string) (autocert.Cache, error) {
	driver, ok := options["driver"]
	if !ok {
		return nil, fmt.Errorf("No driver specified")
	}
	dsn, ok := options["dsn"]
	if !ok {
		return nil, fmt.Errorf("dsn option to backend is required")
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("Error establishing database connection: %s", err)
	}
	if e := db.Ping(); e != nil {
		return nil, fmt.Errorf("Error contacting database: %s", e)
	}

	var usePrecaching bool
	usePrecachingString, ok := options["usePrecaching"]
	if ok {
		usePrecaching, _ = strconv.ParseBool(usePrecachingString)
	}

	encryptionKey, ok := options["encryptionKey"]
	if !ok || encryptionKey == "" {
		return sqlcache.NewCache(db, nil, usePrecaching)
	}

	return sqlcache.NewCache(db, []byte(encryptionKey), usePrecaching)
}
