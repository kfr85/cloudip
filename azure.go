package main

type AzureJson struct {
	ChangeNumber int          `json:"changeNumber"`
	Cloud        string       `json:"cloud"`
	Values       []AzureIPSet `json:"values"`
}

type AzureIPSet struct {
	Name       string        `json:"name"`
	Id         string        `json:"id"`
	Properties AzureProperty `json:"properties"`
}

type AzureProperty struct {
	ChangeNumber    int      `json:"changeNumber"`
	Region          string   `json:"region"`
	RegionId        int      `json:"regionId"`
	Platform        string   `json:"platform"`
	SystemService   string   `json:"systemService"`
	AddressPrefixes []string `json:"addressPrefixes"`
	NetworkFeatures []string `json:"networkFeatures"`
}
