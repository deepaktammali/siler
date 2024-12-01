package main

import (
	"strings"

	"golang.org/x/net/html"
)

func (app *Config) GetPageHTML(url string) (string, error) {
	page := app.Browser.MustPage(url)
	page.MustWindowFullscreen()
	content, error := page.MustWaitStable().HTML()
	return content, error
}

func (app *Config) GetSiteInfo(url string) (*SiteInfo, error) {
	page, _ := app.GetPageHTML(url)
	metadata, _ := ParseMetadata(page)
	links, _ := ParseLinks(page)

	siteInfo := &SiteInfo{
		Metadata: metadata,
		Links:    links,
	}

	return siteInfo, nil
}

func ParseMetadata(htmlContent string) (SiteMetadata, error) {
	siteMetadata := SiteMetadata{}

	var headNode *html.Node
	doc, _ := html.Parse(strings.NewReader(htmlContent))

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == HtmlTagHead {
			headNode = n
			break
		}
	}

	for n := range headNode.Descendants() {
		if n.Type == html.ElementNode && n.Data == HtmlTagTitle {
			siteMetadata.Title = n.FirstChild.Data
		}
	}

	return siteMetadata, nil
}

func ParseLinks(htmlContent string) ([]string, error) {
	links := []string{}
	doc, _ := html.Parse(strings.NewReader(htmlContent))

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == HtmlTagLink {
			for _, attribute := range n.Attr {
				if attribute.Key == HtmlAttributeLinkHref {
					link := attribute.Val
					links = append(links, link)
				}
			}
		}
	}

	return links, nil
}
