//go:build integration || stress

package tests

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("../.env")
	_ = godotenv.Load(".env")
	os.Exit(m.Run())
}
