package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/OpenFoodOrdering/gobackend/data"
	"github.com/OpenFoodOrdering/gobackend/db"
	"github.com/go-chi/chi"
	"github.com/urfave/cli/v2"
)

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

	// Initialize MongoDb
	db.Init(cCtx.String("mongodb_url"))

	// Get Port
	port := fmt.Sprint(":", cCtx.Int("port"))

	// Get Specific Item
	r.Get("/menus/{id}", data.GetOneMenu)

	// Get All Menus
	r.Get("/menus/", data.GetMenus)

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
