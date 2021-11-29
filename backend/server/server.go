package main

import (
	"context"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/hc100/nobolist/backend"
	"github.com/hc100/nobolist/backend/ent"
	"github.com/hc100/nobolist/backend/ent/migrate"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/hc100/nobolist/backend/ent/runtime"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load dotenv")
	}

	dataSource := os.Getenv("DSN")
	if dataSource == "" {
		panic("failed to load datasource")
	}

	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	log, _ := zap.NewDevelopment()
	client, err := ent.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	srv := handler.NewDefaultServer(backend.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	http.Handle("/",
		playground.Handler("GraphQL playground", "/query"),
	)
	http.Handle("/query", srv)

	log.Info("listening on", zap.String("address", cli.Addr))
	if err := http.ListenAndServe(cli.Addr, nil); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
