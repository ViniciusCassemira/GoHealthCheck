package output

type UrlCheckResult struct {
	URL             string `json:"url"`
	ExecutionTimeInSeconds float64 `json:"execution_time_in_seconds"`
	StatusCode      int `json:"status_code"`
	ProtocolVersion string `json:"protocol_version"`
}

type UrlCheckResultList struct {
	TotalUrls int `json:"total_results"`
	ExecutionTimeInSeconds float64 `json:"execution_time_in_seconds"`
	Results []UrlCheckResult `json:"results"`
}

type CheckOptions struct {
	CheckStatusCode      bool `json:"status_code"`
	CheckProtocolversion bool `json:"protocol_version"`
}

type FlagConfig struct {
	UrlFilePath string
	JsonPath string
}
