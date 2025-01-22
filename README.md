# SILER

Siler is a web crawler written in Go that explores websites starting from a seed URL and gathers basic metadata and links.


## Prerequisites

- Go 1.23.3 or higher


## Installation

1. Clone the repository:

```bash
git clone https://github.com/deepaktammali/siler.git

cd siler
```

2. Install dependencies:

```bash
go mod tidy
```

## Usage

Currently, the crawler uses a hardcoded starting URL (google.com). Run the crawler:

```bash
go run .
```

## Project Structure

- main.go - Entry point and crawling logic
- types.go - Data structure definitions
- set.go - Set implementation for collecting unique domains
- helpers.go - HTML parsing and page processing utilities
- constants.go - HTML constants



## Current Implementation

### Data Structures

#### SiteInfo

```go
type SiteInfo struct {
    Url      string
    Metadata SiteMetadata
    Links    []string
}
```

#### SiteMetadata

```go
type SiteMetadata struct {
    Title string
}
```

## Current Status

- [x] Fetch site info for a single site
- [ ] Take the site url from the CLI argument
- [ ] Follow the links from the site and collect information about links upto a certain depth
