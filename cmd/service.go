package cmd

import (
	"github.com/JaredSnapp/go_backend/internal/config"
	HTTPHandler "github.com/JaredSnapp/go_backend/internal/io/inputs/http"
	"github.com/JaredSnapp/go_backend/internal/io/outputs/postgres"
	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/JaredSnapp/go_backend/internal/service/goals"
	"github.com/JaredSnapp/go_backend/internal/service/persons"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Main backend service",

	Run: main,
}

func main(cmd *cobra.Command, args []string) {
	conf := config.Get()

	db := postgres.NewService(conf)
	db.Connect()
	err := db.Migrate()
	if err != nil {
		// log.Fatal(err)
		panic("Error Migrating DB")
	}

	// goalDB := postgres.GenericGen[*models.GoalMetaData](db)
	goalDB := postgres.Repo[models.GoalMetaData]{Pg: db}

	ps := persons.NewService(db)
	g := goals.NewService(goalDB)

	HTTPHandler := HTTPHandler.NewHandler(conf, ps, g)

	HTTPHandler.Serve()
}
