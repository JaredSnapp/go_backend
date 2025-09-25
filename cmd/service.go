package cmd

import (
	"github.com/JaredSnapp/go_backend/internal/config"
	HTTPHandler "github.com/JaredSnapp/go_backend/internal/io/inputs/http"
	"github.com/JaredSnapp/go_backend/internal/io/outputs/postgres"
	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/JaredSnapp/go_backend/internal/service/persons"
	"github.com/JaredSnapp/go_backend/internal/service/recruiter"
	"github.com/JaredSnapp/go_backend/internal/service/richText"
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

	recruiterDB := postgres.Repo[models.Recruiter]{Pg: db}
	richTextDB := postgres.Repo[models.RichText]{Pg: db}

	ps := persons.NewService(db)
	r := recruiter.NewService(recruiterDB)
	rt := richText.NewService(richTextDB)

	HTTPHandler := HTTPHandler.NewHandler(conf, ps, r, rt)

	HTTPHandler.Serve()
}
