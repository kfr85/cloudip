package main

type AwsJson struct {
	Sync_token  string     `json:"syncToken"`
	Create_date string     `json:"createDate"`
	Prefixes    []AwsIPSet `json:"prefixes"`
}

type AwsIPSet struct {
	Ip_prefix            string `json:"ip_prefix"`
	Region               string `json:"region"`
	Service              string `json:"service"`
	Network_border_group string `json:"network_border_group"`
}
