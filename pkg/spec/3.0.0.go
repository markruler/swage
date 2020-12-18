package spec

// OpenAPI is OAS v3.0.x
// http://spec.openapis.org/oas/v3.0.3
type OpenAPI struct {
	Openapi      string                          `json:"openapi"`
	Info         V3Info                          `json:"info"`
	Servers      []Server                        `json:"servers"`
	Paths        map[string]map[string]Operation `json:"paths"`
	Components   Component                       `json:"components"`
	Tags         []Tag                           `json:"tags"`
	ExternalDocs ExternalDocs                    `json:"externalDocs"`
	// Security     []Security                      `json:"security"`
}

// V3Info ...
type V3Info struct {
}

// Server ...
type Server struct {
}

// Component ...
type Component struct {
}
