### Metadata Scraper
[![CI](https://github.com/bilalcaliskan/metadata-scraper/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/metadata-scraper/actions?query=workflow%3ACI)

This tool searches Bing and Google search engines for specific `domain` and docx files and fetch information from these docx files. 

#### Running
##### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/metadata-scraper/releases) page.

After extracted the tarball, run the below command:
```shell
$ ./metadata-scraper --domain ${YOUR_PREFERRED_DOMAIN} 
```

##### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/metadata-scraper:latest
```