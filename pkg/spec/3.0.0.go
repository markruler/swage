package spec

// OpenAPI 3.0.0
type OpenAPI struct {
	Openapi      string                          `json:"openapi"`
	Info         V3Info                          `json:"info"`
	Servers      []Server                        `json:"servers"`
	Paths        map[string]map[string]Operation `json:"paths"`
	Components   Component                       `json:"components"`
	Security     []Security                      `json:"security"`
	Tags         []Tag                           `json:"tags"`
	ExternalDocs ExternalDocs                    `json:"externalDocs"`
}

type V3Info struct {
}

type Server struct {
}

type Component struct {
}
