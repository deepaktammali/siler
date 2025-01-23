package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-rod/rod"
)

type Config struct {
	Browser *rod.Browser
}

var maxDepth int

var siteInfoCache = &SiteInfoCache{
	siteInfoMap: make(map[string]*SiteInfo),
}

func (app *Config) crawlWebsite(site string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	_, ok := siteInfoCache.GetSiteInfo(site)

	// If already visited and stored return
	if ok == true {
		fmt.Printf("Skipped exploring site=%s as it is already explored\n", site)
		return
	}

	siteInfo, err := app.GetSiteInfo(site)

	if err != nil {
		fmt.Printf("Error getting site info for %s: %v\n", site, err)
		return
	}

	siteInfoCache.SetSiteInfo(siteInfo.Url, siteInfo)

	var domainsToExplore = Set{}

	for _, link := range siteInfo.Links {
		urlHome, err := normalizeSiteURL(link)

		if err != nil {
			continue
		}

		domainsToExplore.Add(urlHome)
	}

	filteredDomainsList := domainsToExplore.Keys()

	var subWg sync.WaitGroup

	if depth < maxDepth {
		for _, link := range filteredDomainsList {
			fmt.Printf("Exploring site=%s\n", link)

			subWg.Add(1)
			go app.crawlWebsite(link, depth+1, &subWg)
		}
	}

	subWg.Wait()
}

func main() {
	defaultRootWebsite := "https://www.google.com"
	rootWebsite := flag.String("website", defaultRootWebsite, "The root website to explore from")
	flag.IntVar(&maxDepth, "depth", 3, "Max depth to explore")

	flag.Parse()

	browser := rod.New().NoDefaultDevice().MustConnect()
	defer browser.MustClose()

	app := Config{
		Browser: browser,
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go app.crawlWebsite(*rootWebsite, 1, &wg)

	wg.Wait()

	domainsExplored := make([]string, 0, len(siteInfoCache.siteInfoMap))

	for k := range siteInfoCache.siteInfoMap {
		domainsExplored = append(domainsExplored, k)
	}

	fmt.Printf("Explored %d hosts\n", len(domainsExplored))
	jsonString, err := json.Marshal(siteInfoCache.siteInfoMap)

	if err != nil {
		fmt.Println("Error converting the site data to json string")
		fmt.Println(err)
	} else {
		filename := fmt.Sprintf("output/site_info_%s.json", time.Now().Format("2006-01-02_15-04-05"))
		siteInfoFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
		_, err = siteInfoFile.Write(jsonString)

		if err != nil {
			fmt.Println("Error writing json string to file", err)
		}

		if err := siteInfoFile.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
