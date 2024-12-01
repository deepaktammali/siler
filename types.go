package main

type SiteInfo struct {
	Url      string
	Metadata SiteMetadata
	Links    []string
}

type SiteMetadata struct {
	Title string
}
