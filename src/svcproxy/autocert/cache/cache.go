package cache

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/acme/autocert"

	sqlcache "svcproxy/autocert/cache/sql"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// NewCacheFactory returns Cache instance
func NewCacheFactory(backend string, options map[string]string) (autocert.Cache, error) {
	switch backend {
	case "sql":
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
			return nil, fmt.Errorf("Error establising database connection: %s", err)
		}
		if e := db.Ping(); e != nil {
			return nil, fmt.Errorf("Error contacting database: %s", e)
		}

		return sqlcache.NewCache(db, []byte(options["encryptionKey"]))
	}

	return nil, fmt.Errorf("Unknown backend specified")
}