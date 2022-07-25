package shodan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type QueryInfo struct {
	City     string   `json:"city"`
	Host     string   `json:"ip_str"`
	IP       int      `json:"ip"`
	Ports    []int    `json:"ports"`
	Domains  []string `json:"domains"`
	Hostname []string `json:"hostnames"`
}

const shodanAPIINFO = "https://api.shodan.io/api-info"

func jungle(r *http.Response) {
	if r.StatusCode == 401 {
		log.Fatalf("[-]Maximum number of queries\n\n")
	}
}

func printBody(r *http.Response) []byte {
	defer r.Body.Close()
	context, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return context
}

func TestREQ(token string) {
	r, err := http.NewRequest(http.MethodGet, shodanAPIINFO, nil)
	if err != nil {
		log.Fatalln(err)
	}
	params := make(url.Values)
	params.Add("key", token)
	r.URL.RawQuery = params.Encode()
	req, err := http.DefaultClient.Do(r)
	fmt.Println("[*]url : ", r.URL)
	defer req.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(printBody(req))
}

func (QueryInfo) QueryByIP(token, ip string) QueryInfo {
	QueryURL := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ip, token)
	r, err := http.Get(QueryURL)
	jungle(r)
	if err != nil {
		log.Fatalln(err)
	}
	var info QueryInfo
	err = json.Unmarshal(printBody(r), &info)
	if err != nil {
		log.Fatalln(err)
	}
	return info
}

func (QueryInfo) QuerySubdomainsByDomain(token, domain string) []string {
	QueryURL := fmt.Sprintf("https://api.shodan.io/dns/domain/%s?key=%s", domain, token)
	r, err := http.Get(QueryURL)
	jungle(r)
	if err != nil {
		log.Fatalln(err)
	}
	var result []string
	err = json.Unmarshal(printBody(r), &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func (QueryInfo) OutputQueryByIP(test QueryInfo) {
	fmt.Println("[+]City:", test.City)
	fmt.Println("[+]Host:", test.Host)
	//fmt.Println("[+]IP:", test.IP)
	fmt.Println("[+]Ports:", test.Ports)
	fmt.Println("[+]Domains:", test.Domains)
	fmt.Println("[+]Hostname:", test.Hostname)
}

func (QueryInfo) OutputSubdomains(test []string) {
	fmt.Printf("[+]Subdomains:\n%s", test)
}
