package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"golang.org/x/net/html"
)

func (app *Config) GetPageHTML(url string) (string, error) {
	page := app.Browser.MustPage()
	page.MustWindowFullscreen()
	page.MustWaitNavigation()
	page.Navigate(url)

	_, _ = CatchPanic(func() *rod.Page {
		return page.Timeout(10 * time.Second).MustWaitStable()
	})

	content, error := page.HTML()
	page.MustClose()
	return content, error
}

func (app *Config) GetSiteInfo(url string) (*SiteInfo, error) {
	page, _ := app.GetPageHTML(url)
	metadata, _ := ParseMetadata(page)
	links, _ := ParseLinks(page)

	siteInfo := &SiteInfo{
		Url:      url,
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

func normalizeSiteURL(site string) (string, error) {
	u, err := url.Parse(site)

	if err != nil {
		return "", err
	}

	if u.Scheme == "" || u.Host == "" {
		return "", errors.New("The scheme or the host value is empty")
	}

	normalizedURL := fmt.Sprintf("%s://%s", u.Scheme, u.Host)

	return normalizedURL, err
}

func CatchPanic[T interface{}](f func() T) (ret T, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = errors.New(fmt.Sprint(rec))
		}
	}()
	return f(), nil
}
