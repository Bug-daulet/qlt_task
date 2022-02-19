package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Bug-daulet/qlt_task/pkg/models/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)

type app struct {
	errorLog 		*log.Logger
	infoLog  		*log.Logger
	dbPoolPayment   *postgres.PaymentRepository
	dbPoolCategory	*postgres.CategoryRepository
}


func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	connString := "postgres://postgres:admin@localhost:5432/qlt_task"
	dsn := flag.String("dsn", connString, "PostgreSQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	var checkDBConn string
	err = conn.QueryRow(context.Background(), "select 'Connected to database'").Scan(&checkDBConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	infoLog.Print(checkDBConn)


	app := &app{
		errorLog: 			errorLog,
		infoLog:  			infoLog,
		dbPoolPayment:  	&postgres.PaymentRepository{Pool: conn},
		dbPoolCategory:		&postgres.CategoryRepository{Pool: conn},
	}


	mux := app.routes()

	srv := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: app.errorLog,
	}

	app.infoLog.Printf("Starting a server on %s", *addr)
	err = srv.ListenAndServe()
	app.errorLog.Fatal(err)

}

func openDB(dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
