## Opinionated URL Shortner Client
This is a simple URL shortner Go Lang client that uses the [short.io](https://short.io)

## TL;DR
- One liner (for my use case)
```
./short.io -key $short_io_key -customized-link-path pushpa -original-link https://www.imdb.com/title/tt9389998
```
## How to use 
- Create a free account on [short.io](https://short.io)
- Get the API Key
- Register your custome domain name with short.io and update the CNAME record to point to your domain name

- Short the URL
```
./short.io -key xxxxx -allow-duplicates true -customized-link-path pushpa -domain short.ksingh.in -original-link https://www.imdb.com/title/tt9389998-tag RHD -title 'mai jhukaga nahi' -utm-source ksingh-blogs -debug false
```
- All Options
```
Usage of ./short.io :
  -allow-duplicates string
    	(optional) allow duplicates. If empty - it will be set to false (default "false")
  -customized-link-path string
    	(optional) path part of newly created link. If empty - it will be generated automatically
  -debug string
    	(optional) debug mode. If empty - it will be set to false (default "false")
  -domain string
    	domain to shorten the link to (default "short.ksingh.in")
  -key string
    	API key
  -original-link string
    	link that needs to be shortend
  -tag string
    	(optional) tag to be added to created URL (default "RHD")
  -title string
    	(optional) title of created URL to be shown in short.io admin panel, use  ' '  (default "hello test")
  -utm-source string
    	(optional) utm source to be added to created URL (default "ksingh-blogs")
```
## How to build 
```
git clone https://github.com/ksingh7/url-shortner.git
cd url-shortner
go build -o short.io main.go
```