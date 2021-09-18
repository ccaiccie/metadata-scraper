# Metadata Scraper
[![CI](https://github.com/bilalcaliskan/metadata-scraper/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/metadata-scraper/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/metadata-scraper)](https://hub.docker.com/r/bilalcaliskan/metadata-scraper/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/metadata-scraper)](https://goreportcard.com/report/github.com/bilalcaliskan/metadata-scraper)
[![codecov](https://codecov.io/gh/bilalcaliskan/metadata-scraper/branch/master/graph/badge.svg)](https://codecov.io/gh/bilalcaliskan/metadata-scraper)
[![Release](https://img.shields.io/github/release/bilalcaliskan/metadata-scraper.svg)](https://github.com/bilalcaliskan/metadata-scraper/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/metadata-scraper)](https://github.com/bilalcaliskan/metadata-scraper)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This tool searches Bing and Google search engines for specific `domain` and docx files and fetch information from these docx files.

## Configuration
metadata-scraper can be customized with several command line arguments:
```
--domain            target domain to search on search engines, defaults to nytimes.com
--fileType          file type to search for on the specified domain, defaults to docx
```

## Download
### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/metadata-scraper/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ ./metadata-scraper --randomLength 12 --emailDomains jentrix.com,geekale.com
```

### Docker
You can simply run docker image with default configuration:
```shell
$ docker run bilalcaliskan/metadata-scraper:latest
```

## Development
This project requires below tools while developing:
- [Golang 1.16](https://golang.org/doc/go1.16)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)

## License
Apache License 2.0
