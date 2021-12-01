package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
	"github.com/go-chi/chi"
	"github.com/hc100/nobolist/backend"
	"github.com/hc100/nobolist/backend/ent"
	"github.com/hc100/nobolist/backend/ent/migrate"
	"github.com/hc100/nobolist/backend/internal/auth"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/hc100/nobolist/backend/ent/runtime"
)

const envDevelopment = "development"
const defaultAllowedOrigins = "http://localhost:3000"
const defaultEnv = "production"
const defaultDebugCors = false

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load dotenv")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = defaultEnv
	}

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = defaultAllowedOrigins
	}

	debugCors := defaultDebugCors
	if env == envDevelopment {
		debugCors = true
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

	router := chi.NewRouter()

	router.Use(auth.Middleware(client))

	allowedOriginsList := strings.Split(allowedOrigins, "|")
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   allowedOriginsList,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            debugCors,
	}).Handler)

	if env == envDevelopment {
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Info("connect to http://localhost:%s/ for GraphQL playground", zap.String("address", cli.Addr))
	}
	router.Handle("/query", srv)

	log.Info("listening on", zap.String("address", cli.Addr))
	if err := http.ListenAndServe(cli.Addr, router); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
