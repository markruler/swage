package spec

// SwaggerAPI is OAS v2.0 (formerly Swagger Specification)
// http://spec.openapis.org/oas/v2.0
type SwaggerAPI struct {
	Swagger             string                          `json:"swagger"`
	Info                Info                            `json:"info"`
	Host                string                          `json:"host"`
	BasePath            string                          `json:"basePath"`
	Tags                []Tag                           `json:"tags"`
	Schemes             []string                        `json:"schemes"`
	Consumes            []string                        `json:"consumes"`
	Produces            []string                        `json:"produces"`
	Paths               map[string]map[string]Operation `json:"paths"`
	Definitions         map[string]Definition           `json:"definitions"`
	SecurityDefinitions map[string]SecurityDefinition   `json:"securityDefinitions"`
	ExternalDocs        ExternalDocs                    `json:"externalDocs"`
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
	Name  string `json:"name"`
	URL   string `json:"url"`
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

// Operation ...
type Operation struct {
	Tags         []string              `json:"tags"`
	Summary      string                `json:"summary"`
	Description  string                `json:"description"`
	ExternalDocs ExternalDocs          `json:"externalDocs"`
	OperationID  string                `json:"operationId"`
	Consumes     []string              `json:"consumes"`
	Produces     []string              `json:"produces"`
	Parameters   []Parameters          `json:"parameters"`
	Responses    map[string]Response   `json:"responses"`
	Schemes      []string              `json:"schemes"`
	Deprecated   string                `json:"deprecated"`
	Security     []map[string][]string `json:"security"`
}

// Parameters ...
type Parameters struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	// in: "body"
	Schema           Schema        `json:"schema"`
	Format           string        `json:"format"`
	AllowEmptyValue  bool          `json:"allowEmptyValue"`
	Items            Items         `json:"items"`            // Required if type is “array”.
	CollectionFormat string        `json:"collectionFormat"` // ["csv", "ssv", "tsv", "pipes", "multi"]
	Default          interface{}   `json:"default"`
	Maximum          float32       `json:"maximum"`
	ExclusiveMaximum bool          `json:"exclusiveMaximum"`
	Minimum          float32       `json:"minimum"`
	ExclusiveMinimum bool          `json:"exclusiveMinimum"`
	MaxLength        uint32        `json:"maxLength"`
	MinLength        uint32        `json:"minLength"`
	Pattern          string        `json:"pattern"`
	MaxItems         uint32        `json:"maxItems"`
	MinItems         uint32        `json:"minItems"`
	UniqueItems      bool          `json:"uniqueItems"`
	Enum             []interface{} `json:"enum"`
	MultipleOf       float32       `json:"multipleOf"`
}

// Items ...
type Items struct {
	Ref     string   `json:"$ref"`
	Type    string   `json:"type"`
	Enum    []string `json:"enum"`
	Default string   `json:"default"`
}

// Schema ...
type Schema struct {
	Discriminator        string              `json:"discriminator"`
	ReadOnly             bool                `json:"readOnly"`
	XML                  XML                 `json:"xml"`
	ExternalDocs         ExternalDocs        `json:"externalDocs"`
	Example              interface{}         `json:"example"` // TODO: any
	Type                 string              `json:"type"`
	Format               string              `json:"format"`
	Required             []string            `json:"required"`
	Ref                  string              `json:"$ref"`
	Items                Items               `json:"items"`
	Properties           map[string]Property `json:"properties"`
	AdditionalProperties Property            `json:"additionalProperties"`
	AllOf                []interface{}       `json:"allOf"` // TODO: any
}

// Response ...
type Response struct {
	Description string                       `json:"description"`
	Schema      Schema                       `json:"schema"`
	Headers     map[string]map[string]string `json:"headers"`
	Examples    map[string]map[string]string `json:"examples"` // TODO: any
}

// Header ...
type Header struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	Description string `json:"description"`
}

// SecurityDefinition ...
type SecurityDefinition struct {
	// type: [ "basic", "apiKey", "oauth2" ]
	Type             string            `json:"type"`
	In               string            `json:"in,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	Flow             string            `json:"flow,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

// Definition ...
type Definition struct {
	Type       string              `json:"type"`
	Required   []string            `json:"required"`
	Properties map[string]Property `json:"properties"`
	XML        XML                 `json:"xml"`
}

// Property ...
type Property struct {
	Ref         string   `json:"$ref"`
	Type        string   `json:"type"`
	Format      string   `json:"format"`
	Description string   `json:"description"`
	XML         XML      `json:"xml"`
	Items       Items    `json:"items"`
	Example     string   `json:"example"`
	Enum        []string `json:"enum"`
	Maximum     float32  `json:"maximum"`
	Minimum     float32  `json:"minimum"`
}

// XML ...
type XML struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Prefix    string `json:"prefix"`
	Attribute bool   `json:"attribute"`
	Wrapped   bool   `json:"wrapped"`
}
