package database

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "123"
// 	dbname   = "postgres"
// 	sslmode  = "disable"
// )

var (
	DB  *sql.DB
	err error
)

// Embed semua file SQL migration
//
//go:embed sql_migrations/*.sql
var dbMigrations embed.FS

var DbConnection *sql.DB

func Connect() *sql.DB {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using system env")
	}

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Successfully connected to database!")
	DB = db
	return db
}

func DBMigrate(dbParam *sql.DB) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Println("🚀 Migration success, applied", n, "migrations!")
}
