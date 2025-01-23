package main

import "sync"

type SiteInfoCache struct {
	mu          sync.Mutex
	siteInfoMap map[string]*SiteInfo
}

type SiteInfo struct {
	Url      string
	Metadata SiteMetadata
	Links    []string
}

type SiteMetadata struct {
	Title string
}

func (c *SiteInfoCache) GetSiteInfo(site string) (*SiteInfo, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	siteInfo, ok := c.siteInfoMap[site]
	return siteInfo, ok
}

func (c *SiteInfoCache) SetSiteInfo(site string, siteInfo *SiteInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.siteInfoMap[site] = siteInfo
}
