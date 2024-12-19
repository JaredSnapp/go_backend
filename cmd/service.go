package cmd

import (
	"github.com/JaredSnapp/go_backend/internal/config"
	HTTPHandler "github.com/JaredSnapp/go_backend/internal/io/inputs/http"
	"github.com/JaredSnapp/go_backend/internal/service/persons"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Main backend service",

	Run: main,
}

func main(cmd *cobra.Command, args []string) {

	ps := persons.NewService()

	conf := config.Get()
	HTTPHandler := HTTPHandler.NewHandler(conf, ps)
	HTTPHandler.Serve()

	// Serve()

	//  http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
