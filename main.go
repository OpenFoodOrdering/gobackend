package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb mongo.Client

func main() {
	app := app()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Function Passed As Action to Cli Structm
func run(cCtx *cli.Context) error {
	// New CHI Router
	r := chi.NewRouter()

	ServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	ClientOptions := options.Client().ApplyURI(cCtx.String("mongodb_url")).SetServerAPIOptions(ServerAPI)

	// Initialize MongoDb Client Connection Pool Using ClientOptions
	MongoDb, err := mongo.NewClient(ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Get Port
	port := fmt.Sprint(":", cCtx.Int("port"))

	// Serve Using Router
	http.ListenAndServe(port, r)

	return nil
}

// Function that Generates the Cli struct
func app() cli.App {
	return cli.App{
		Name:   "FoodOrderingBackendGo",
		Usage:  "Serve and Respond To Food Order Requests",
		Action: run,
		// Flags:
		Flags: []cli.Flag{
			// Port Where the App Would Run
			&cli.IntFlag{
				Name:    "port",
				Value:   3000,
				Aliases: []string{"p"},
				Usage:   "Port Where the App Will Run",
				EnvVars: []string{"SERVER_PORT"},
			},
			&cli.StringFlag{
				Name:    "mongodb_url",
				Aliases: []string{"db"},
				Usage:   "Where the Database Will Exist",
				EnvVars: []string{"DATABASE_URL"},
			},
		},
	}
}
