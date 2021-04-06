package parser

type SwageSpec struct {
	API []SwageAPI
}

type SwageAPI struct {
	Header   APIHeader
	Request  APIRequest
	Response APIResponse
}

type APIHeader struct {
	Tag         string
	ID          string
	Path        string
	Method      string
	Consumes    string
	Produces    string
	Summary     string
	Description string
}

type APIRequest struct {
	Required      string
	Schema        string
	ParameterType string // path, body, query, header, ...
	DataType      string // string, integer, boolean, object, ...
	Enum          string
	Example       string
	Description   string
}

type APIResponse struct {
	StatusCode   string
	Schema       string
	ResponseType string // body
	DataType     string // string, integer, boolean, object, ...
	Enum         string
	Example      string
	Description  string
}
