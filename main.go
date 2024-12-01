package main

import (
	"fmt"
	"net/url"

	"github.com/go-rod/rod"
)

type Config struct {
	Browser *rod.Browser
}

const maxDepth = 1

var siteInfoMap = map[string]*SiteInfo{}

func (app *Config) crawlWebsite(site string, depth int) {
	siteInfo, _ := app.GetSiteInfo(site)
	siteInfoMap[site] = siteInfo

	var domainsToExplore = Set{}

	for _, link := range siteInfo.Links {
		u, err := url.Parse(link)
		if err != nil {
			continue
		}

		if u.Scheme == "" || u.Host == "" {
			continue
		}

		urlHome := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
		fmt.Println(link, urlHome)
		domainsToExplore.Add(urlHome)
	}

	filteredDomainsList := domainsToExplore.Keys()

	if depth < maxDepth {
		for _, link := range filteredDomainsList {
			app.crawlWebsite(link, depth+1)
		}
	}
}

func main() {
	browser := rod.New().NoDefaultDevice().MustConnect()
	defer browser.MustClose()

	app := Config{
		Browser: browser,
	}

	app.crawlWebsite("https://www.google.com", 1)
	fmt.Println(siteInfoMap)
}
