package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	url_aws   = "https://ip-ranges.amazonaws.com/ip-ranges.json"
	url_azure = "https://download.microsoft.com/download/7/1/D/71D86715-5596-4529-9B13-DA13A5DE5B63/ServiceTags_Public_20230116.json"
	url_gcp   = "https://www.gstatic.com/ipranges/cloud.json"
)

func GetJsonFromUrl(url string, cloud interface{}) {
	// Get
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Decode
	err = json.NewDecoder(resp.Body).Decode(&cloud)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var aws AwsJson
	GetJsonFromUrl(url_aws, &aws)
	for _, v := range aws.Prefixes {
		fmt.Printf("%v,aws,%v,%v\n", v.Ip_prefix, v.Region, v.Service)
	}

	var gcp GcpJson
	GetJsonFromUrl(url_gcp, &gcp)
	for _, v := range gcp.Prefixes {
		if v.Ipv4Prefix == "" {
			fmt.Printf("%v,gcp,%v,%v\n", v.Ipv6Prefix, v.Scope, v.Service)
		} else if v.Ipv6Prefix == "" {
			fmt.Printf("%v,gcp,%v,%v\n", v.Ipv4Prefix, v.Scope, v.Service)
		}
	}

	var azure AzureJson
	GetJsonFromUrl(url_azure, &azure)
	for _, v := range azure.Values {
		for index := range v.Properties.AddressPrefixes {
			if v.Properties.SystemService == "" {
				fmt.Printf("%v,%v,UnknownService\n", v.Properties.AddressPrefixes[index], v.Properties.Platform)
			} else {
				fmt.Printf("%v,%v,%v\n", v.Properties.AddressPrefixes[index], v.Properties.Platform, v.Properties.SystemService)
			}
		}
	}
}

