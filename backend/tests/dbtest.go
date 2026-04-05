//go:build integration || stress

package tests

import (
	"testing"
	"time"

	"github.com/uaad/backend/internal/config"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func openTestDB(tb testing.TB) *gorm.DB {
	tb.Helper()
	cfg := config.Load()
	db, err := gorm.Open(gormmysql.Open(cfg.MySQLDSN()), &gorm.Config{})
	if err != nil {
		tb.Fatalf("open database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		tb.Fatalf("sql.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
