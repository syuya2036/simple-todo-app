package main

import (
	"os"
)

func init() {
	env := "docker"
	switch env {
	case "dev":
		// DBの接続情報
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_NAME", "postgres")
		os.Setenv("DB_PASSWORD", "postgres")
		os.Setenv("DB_SSLMODE", "disable")
		os.Setenv("PORT", "8080")
	case "docker":
		// DBの接続情報
		os.Setenv("DB_HOST", "postgres")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_NAME", "postgres")
		os.Setenv("DB_PASSWORD", "postgres")
		os.Setenv("DB_SSLMODE", "disable")
		os.Setenv("PORT", "8080")
	}
}