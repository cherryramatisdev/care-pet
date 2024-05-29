package main

import (
	"context"
	"log"

	"github.com/spf13/pflag"

	"github.com/scylladb/care-pet/go/config"
	"github.com/scylladb/care-pet/go/db"
	"github.com/scylladb/gocqlx/v2/migrate"
)

var verbose = pflag.Bool("verbose", false, "output more info")

func main() {
	pflag.Parse()

	log.Println("Bootstrap database...")

	if *verbose {
		log.Printf("Configuration = %+v\n", config.Config())
	}

	createKeyspace()
	log.Println("Keyspace Created!")
	migrateKeyspace()
	log.Println("Migrations settled!")
	printKeyspaceMetadata()
	log.Println("Now you're ready to use the application! Make sure to test the whole app: sensor, server and loadtest.")
}

func createKeyspace() {
	ses, err := config.Session()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := ses.Query(db.KeySpaceCQL).Exec(); err != nil {
		log.Fatalln("ensure keyspace exists: ", err)
	}
}

func migrateKeyspace() {
	ses, err := config.Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	if err := migrate.Migrate(context.Background(), ses, "db/cql"); err != nil {
		log.Fatalln("migrate: ", err)
	}
}

func printKeyspaceMetadata() {
	ses, err := config.Keyspace()
	if err != nil {
		log.Fatalln("session: ", err)
	}
	defer ses.Close()

	m, err := ses.KeyspaceMetadata(db.KeySpace)
	if err != nil {
		log.Fatalln("keyspace metadata: ", err)
	}

	log.Printf("Keyspace metadata = %+v\n", *m)
}
