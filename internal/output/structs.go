package output

type CheckResult struct {
	URL             string `json:"url"`
	StatusCode      int `json:"status_code"`
	ProtocolVersion string `json:"protocol_version"`
}

type CheckResultList struct {
	TotalUrls int `json:"total_results"`
	Results []CheckResult `json:"results"`
}

type CheckOptions struct {
	CheckStatusCode      bool
	CheckProtocolversion bool
}

type FlagConfig struct {
	UrlFilePath string
	JsonPath string
}
