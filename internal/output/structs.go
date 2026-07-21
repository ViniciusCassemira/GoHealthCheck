package output

type CheckResult struct {
	URL             string
	StatusCode      int
	ProtocolVersion string
}

type CheckOptions struct {
	CheckStatusCode      bool
	CheckProtocolversion bool
}
