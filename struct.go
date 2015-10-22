package go_swagger

type Swagger struct {
	SwaggerVersion      string                     `json:"swagger"`
	Info                map[string]interface{}     `json:"info"`
	Host                string                     `json:"host"`
	BasePath            string                     `json:"basePath"`
	Tags                []Tag                      `json:"tags"`
	Schemes             []string                   `json:"schemes"`
	Paths               map[string]map[string]Path `json:"paths"`
	SecurityDefinitions map[string]interface{}     `json:"securityDefinitions,omitempty"`
	Definitions         map[string]Definition      `json:"definitions,omitempty"`
	ExternalDocs        map[string]interface{}     `json:"externalDocs,omitempty"`
}

type Info struct {
	Description    string            `json:"description"`
	Version        string            `json:"version"`
	Title          string            `json:"title"`
	TermsOfService string            `json:"termsOfService"`
	Contact        map[string]string `json:"contact"`
	License        map[string]string `json:"license"`
}

type Tag struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	ExternalDocs map[string]string `json:"externalDocs,omitempty"`
}

type Properties struct {
	Type        string      `json:"type,omitempty"`
	Format      string      `json:"format,omitempty"`
	Example     interface{} `json:"example,omitempty"`
	Description string      `json:"description,omitempty"`
}

type Definition struct {
	Type       string                `json:"type"`
	Properties map[string]Properties `json:"properties"`
	Xml        map[string]string     `json:"xml"`
}

type Path struct {
	Tags        []string               `json:"tags"`
	Summary     string                 `json:"summary,omitempty"`
	Description string                 `json:"description,omitempty"`
	OperationId string                 `json:"operationId,omitempty"`
	Consumes    []string               `json:"consumes"`
	Produces    []string               `json:"produces"`
	Parameters  []Parameters           `json:"parameters"`
	Responses   map[string]interface{} `json:"responses"`
	Security    interface{}            `json:"security,omitempty"`
}
type Parameters struct {
	In          string                 `json:"in"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Required    bool                   `json:"required"`
	Schema      map[string]interface{} `json:"schema"`
	Type        string                 `json:"type"`
}
