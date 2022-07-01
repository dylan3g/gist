package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dylan3g/gist/api"
)

func usage() {
	fmt.Println("gist v0.1.0")
    fmt.Println("Usage: gist <id> [--info|--content]")
    fmt.Println("Flags:")
    fmt.Println("   --info: Fetch information abount a gist.")
    fmt.Println("   --content: Print the gist content to stdout.")
    fmt.Println("   --help: Display this message.")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

    if args[1] == "--help" {
        usage()
        os.Exit(1)
    }
	gist, err := api.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}
	owner := *gist.Owner.Login
	description := *gist.Description
	gist_id := *gist.Id
	created := *gist.CreatedAt
	updated := *gist.UpdatedAt

	switch args[2] {
	case "--info":

		fmt.Printf("Owner: %s\n", owner)
		fmt.Printf("Description: %s\n", description)
		fmt.Printf("Gist id: %s\n", gist_id)
		fmt.Printf("Created at: %s\n", created)
		fmt.Printf("Updated at: %s\n", updated)

		fmt.Println("Files:")
		for _, f := range gist.Files {
			fmt.Printf("\t%s\n", *f.Filename)
		}
	case "--content":
		for _, f := range gist.Files {
			fmt.Println(*f.Content)
		}
	}
}
