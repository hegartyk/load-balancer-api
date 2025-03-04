//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	"go.infratographer.com/load-balancer-api/internal/ent/generated/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	_ "ariga.io/atlas/sql/postgres"
	_ "ariga.io/atlas/sql/postgres/postgrescheck"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("db/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod db/create_migration.go <name>'")
	}
	dbURI, ok := os.LookupEnv("ATLAS_DB_URI")
	if !ok {
		log.Fatalln("failed to load the ATLAS_DB_URI env var")
	}

	// Generate migrations using Atlas support for postgres (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dbURI, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
