package spec

// SwaggerAPI ...
type SwaggerAPI struct {
	Swagger             string                        `json:"swagger"`
	Info                Info                          `json:"info"`
	Host                string                        `json:"host"`
	BasePath            string                        `json:"basePath"`
	Tags                []Tag                         `json:"tags"`
	Schemes             []string                      `json:"schemes"`
	Paths               map[string]Path               `json:"paths"`
	SecurityDefinitions map[string]SecurityDefinition `json:"securityDefinitions"`
	Definitions         map[string]Definition         `json:"definitions"`
	ExternalDocs        ExternalDocs                  `json:"externalDocs"`
}

// Info ...
type Info struct {
	Description    string  `json:"description"`
	Version        string  `json:"version"`
	Title          string  `json:"title"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
}

// Contact ...
type Contact struct {
	Email string `json:"email"`
}

// License ...
type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Tag ...
type Tag struct {
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	ExternalDocs ExternalDocs `json:"externalDocs"`
}

// ExternalDocs ...
type ExternalDocs struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Path ...
type Path struct {
	Get    Operation `json:"get"`
	Post   Operation `json:"post"`
	Put    Operation `json:"put"`
	Delete Operation `json:"delete"`
}

// Operation ...
type Operation struct {
	Tags        []string            `json:"tags"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	OperationID string              `json:"operation_id"`
	Produces    string              `json:"produces"`
	Parameters  []Parameters        `json:"parameters"`
	Responses   map[string]Response `json:"responses"`
	Security    []Security          `json:"security"`
	Deprecated  string              `json:"deprecated"`
}

// Parameters ...
type Parameters struct {
	In               string `json:"in"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Required         bool   `json:"required"`
	Type             string `json:"type"`
	Items            []Item `json:"items"`
	Schema           Schema `json:"schema"`
	Format           string `json:"format"`
	CollectionFormat string `json:"collectionFormat"`
}

// Item ...
type Item struct {
	Ref     string   `json:"$ref"`
	Type    string   `json:"type"`
	Enum    []string `json:"enum"`
	Default string   `json:"default"`
}

// Schema ...
type Schema struct {
	Ref   string `json:"$ref"`
	Type  string `json:"type"`
	Items []Item `json:"items"`
}

// Response ...
type Response struct {
	Description string            `json:"description"`
	Schema      Schema            `json:"schema"`
	Headers     map[string]Header `json:"headers"`
}

// Header ...
type Header struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	Description string `json:"description"`
}

// Security ...
type Security struct {
	PetstoreAuth []string `json:"petstore_auth"`
	APIKey       []string `json:"api_key"`
}

// SecurityDefinition ...
type SecurityDefinition struct {
	Type             string            `json:"type"`
	AuthorizationURL string            `json:"authorizationUrl"`
	Flow             string            `json:"flow"`
	Scopes           map[string]string `json:"scopes"`
	In               string            `json:"in"`
}

// Definition ...
type Definition struct {
	Type       string              `json:"type"`
	Required   []string            `json:"required"`
	Properties map[string]Property `json:"properties"`
	XML        map[string]XML      `json:"xml"`
}

// Property ...
type Property struct {
	Ref         string   `json:"$ref"`
	Type        string   `json:"type"`
	Format      string   `json:"format"`
	Description string   `json:"description"`
	XML         XML      `json:"xml"`
	Items       []Item   `json:"items"`
	Example     string   `json:"example"`
	Enum        []string `json:"enum"`
}

// XML ...
type XML struct {
	Name    string `json:"name"`
	Wrapped bool   `json:"wrapped"`
}
