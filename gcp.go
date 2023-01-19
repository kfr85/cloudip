package main

type GcpJson struct {
	Sync_token    string     `json:"syncToken"`
	Creation_date string     `json:"creationTime"`
	Prefixes      []GcpIPSet `json:"prefixes"`
}

type GcpIPSet struct {
	Ipv4Prefix string `json:"ipv4Prefix"`
	Ipv6Prefix string `json:"ipv6Prefix"`
	Service    string `json:"service"`
	Scope      string `json:"scope"`
}
