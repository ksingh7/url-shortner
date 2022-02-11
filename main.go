package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"flag"
	"os"
)
	type Payload struct {
		AllowDuplicates string	`json:"allowDuplicates"`
		Domain string		`json:"domain"`
		Tags []string		`json:"tags"`
		OriginalURL string	`json:"originalURL"`
		UtmSource string	`json:"utmSource"`
		UtmMedium string	`json:"utmMedium"`
		UtmCampaign string	`json:"utmCampaign"`
		Title string	`json:"title"`
		Path string	`json:"path"`
		}

func main() {
	key := flag.String("key", "", "API key")
	originalURL := flag.String("original-link", "", "link that needs to be shortend")
	domain := flag.String("domain", "short.ksingh.in", "domain to shorten the link to (example : short.ksingh.in)")
	path := flag.String("customized-link-path", "", "(optional) path part of newly created link. If empty - it will be generated automatically")
	title := flag.String("title", "hello test", "(optional) title of created URL to be shown in short.io admin panel, use  ' ' ")
	allowDuplicates := flag.String("allow-duplicates", "false", "(optional) allow duplicates. If empty - it will be set to false")
	tag := flag.String("tag", "RHD", "(optional) tag to be added to created URL")
	utmSource := flag.String("utm-source", "ksingh-blogs", "(optional) utm source to be added to created URL")
	debug := flag.String("debug", "false", "(optional) debug mode. If empty - it will be set to false")
	flag.Parse()
	url := "https://api.short.io/links"
	// key := "sk_QsTrbgdzY8Sg8NdJ"

	payload := &Payload{
	AllowDuplicates: *allowDuplicates,
	Domain: *domain,
	Tags: []string{*tag},
	OriginalURL: *originalURL,
	Path: *path,
	Title: *title,
	UtmSource: *utmSource,
	UtmMedium: *tag,
	UtmCampaign: *tag,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)

	if *debug == "true" {
	fmt.Println(payloadBuf)
	os.Exit(0)
	}

	req, _ := http.NewRequest("POST", url, payloadBuf)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", *key)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))

	if res.StatusCode == 200 {
		var data map[string]interface{}
		json.Unmarshal(body, &data)
		fmt.Println(data["shortURL"])
	} else {
		fmt.Println(res.StatusCode)
		fmt.Println(string(body))
	}

}
