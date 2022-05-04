package main

import (
	"github.com/AJMBrands/SoftwareThatMatters/ingest"
	"runtime"
)

const limited_discovery_query string = "https://libraries.io/api/search?api_key=3dc75447d3681ffc2d17517265765d23&page=1&per_page=1&platforms=MAVEN"

const discovery_query string = "https://libraries.io/api/search?api_key=3dc75447d3681ffc2d17517265765d23&platforms=NPM&per_page=20"

const offline_file string = "data/100packages.json"
const outPathTemplate string = "data/out/parsed_data_%d.csv"

var m1, m2 runtime.MemStats

//TODO: Make ingest process and file writing scalable
func main() {
	runtime.ReadMemStats(&m1) // Reading memory stats for debug purposes
	//ingest.IngestFile(offline_file, outPath)
	ingest.Ingest(limited_discovery_query, outPathTemplate)
	runtime.ReadMemStats(&m2)
}
