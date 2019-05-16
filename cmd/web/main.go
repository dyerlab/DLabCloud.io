package main

import (
	"os"
	"fmt"
	"flag"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/dyerlab/DLabCloud.io/pkg/models/postgres"
)

// var db *gorm.DB
// var err error

//var dbName = "postgres"
//var dbConnect = "host=localhost port=5432 user=rodney dbname=dlab password=bob sslmode=disable"

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	manuscripts *postgres.ManuscriptModel
	snippets    *postgres.SnippetModel
}

func main() {
	fmt.Println("Starting dlabcloud")
	
	addr := flag.String("addr", ":8082", "HTTP Network Address")
	dbName := flag.String("db", "postgres", "Database name (only postgres supported at this time")
	dbConnect := flag.String("dbConnect", "host=localhost port=5432 user=rodney dbname=dlab password=bob sslmode=disable", "Default connection string with host, port, user, etc.")
	flag.Parse()

	// Custom logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New( os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	// Open the database object
	db, err := gorm.Open(*dbName, *dbConnect)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()



	// Make appliation object
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		manuscripts: &postgres.ManuscriptModel{DB: db},
		snippets: &postgres.SnippetModel{DB: db},
	}

	// Migrate the GORM models
	app.manuscripts.AutoMigrate()
	app.snippets.AutoMigrate()


	// Custom server to grab server errors with our errorLog framework.
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}


