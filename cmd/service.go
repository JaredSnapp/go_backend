package cmd

import (
	"fmt"
	//  "log"

	// "github.com/JaredSnapp/go_backend/io/inputs/HTTPHandler"
	HTTPHandler "github.com/JaredSnapp/go_backend/internal/io/inputs"
	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Main backend service",

	Run: main,
}

func main(cmd *cobra.Command, args []string) {
	// person := models.Person{}
	person := models.Person{Name: "george", Age: 24}
	printPerson(&person)

	HTTPHandler.Serve()

	// Serve()

	//  http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}

func printPerson(P *models.Person) {
	fmt.Println("Person:")
	fmt.Println(P.Name)
	fmt.Println(P.Age)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, World!")
// }
