package database

import (
    "database/sql"
    "log"

     _ "github.com/lib/pq"
    "e-library/backend/internal/config"
)

func InitDB(cfg config.Config) *sql.DB {
    db, err := sql.Open("postgres", cfg.DBUrl)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal("Database ping failed:", err)
    }

    return db
}
