package output

type UrlCheckResult struct {
	URL             string `json:"url"`
	ExecutionTime float64 `json:"execution_time"`
	HttpInfo        HttpInfo `json:"http_info"`
	DnsInfo         DnsInfo `json:"dns_info"`
}

type UrlCheckResultList struct {
	TotalUrls int `json:"total_results"`
	ExecutionTime float64 `json:"execution_time"`
	CheckOptions CheckOptions `json:"search_options"`
	Results []UrlCheckResult `json:"results"`
}

type CheckOptions struct {
	CheckHttpInfo        bool `json:"http_info"`
	CheckDnsInfo         bool `json:"dns_info"`
}

type FlagConfig struct {
	UrlFilePath string
	JsonPath string
}

type DnsInfo struct {
	ARecords []string `json:"a_records"`
	AAAARecords []string `json:"aaaa_records"`
	CnameRecords []string `json:"cname_records"`
	MxRecords []string `json:"mx_records"`
	NsRecords []string `json:"ns_records"`
	TxtRecords []string `json:"txt_records"`
}

type HttpInfo struct {
	ProtocolVersion string `json:"protocol_version"`
	StatusCode      int    `json:"status_code"`
}
