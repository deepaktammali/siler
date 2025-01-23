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

Run the crawler:

```bash
go run *.go --website=https://www.google.com --depth=3
```

This runs in headless mode to see the browser. add -rod=show.

```bash
go run *.go -rod=show --website=https://www.google.com --depth=3
```

Output is generated under output folder

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
