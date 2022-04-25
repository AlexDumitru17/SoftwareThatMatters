package main

import (
	"fmt"
	"runtime"

	"github.com/AJMBrands/SoftwareThatMatters/ingest"
)

const limited_discovery_query string = "https://libraries.io/api/search?api_key=3dc75447d3681ffc2d17517265765d23&page=1&per_page=2&platforms=NPM"

const discovery_query string = "https://libraries.io/api/search?api_key=3dc75447d3681ffc2d17517265765d23&platforms=NPM"

var packagesAndVersions []ingest.DiscoveryResponse

var m1, m2 runtime.MemStats

func main() {
	runtime.ReadMemStats(&m1)
	packagesAndVersions = ingest.Ingest(discovery_query)
	runtime.ReadMemStats(&m2)
	fmt.Println(packagesAndVersions)
}
