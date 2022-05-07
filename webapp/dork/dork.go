//NOt Completed

package main

import (
	"context"
	"log"
	"os"

	googlesearch "github.com/rocketlaunchr/google-search"
)

func main() {
	ctx := context.Background()
	search_dork := "inurl:/+CSCOE+/logon.html?"

	output, err := googlesearch.Search(ctx, search_dork)
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println(output[1])
	f, err := os.Create("url.txt")
	{
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err := f.WriteString(output)
		if err != nil {
			log.Fatal(err)
		}
	}
}
