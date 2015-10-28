package go_swagger

import (
	"encoding/json"
	"fmt"
)

type H struct {
}

var VSwagger Swagger = Swagger{}

func (s Swagger) ToString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
func (h H) AddTag(tag Tag) H {
	if VSwagger.Tags == nil {
		VSwagger.Tags = []Tag{}
	}
	VSwagger.Tags = append(VSwagger.Tags, tag)
	return h
}

func (h H) AddDefinition(key string, definition Definition) H {
	if VSwagger.Definitions == nil {
		VSwagger.Definitions = map[string]Definition{}
	}
	VSwagger.Definitions[key] = definition
	return h
}

func (h H) AddPath(key string, path map[string]Path) H {
	fmt.Println("add path,key:", key)
	if VSwagger.Paths == nil {
		VSwagger.Paths = map[string]map[string]Path{}
	}
	VSwagger.Paths[key] = path
	return h
}

func (h H) InitSwagger(s Swagger) H {
	VSwagger.SwaggerVersion = s.SwaggerVersion
	VSwagger.Info = s.Info
	VSwagger.Host = s.Host
	VSwagger.BasePath = s.BasePath
	VSwagger.Schemes = s.Schemes
	VSwagger.ExternalDocs = s.ExternalDocs
	return h
}

func GenByJson(jsonStr string) (m map[string]Properties) {
	json.Unmarshal([]byte(jsonStr), &m)
	return
}

func NewH() H {
	return H{}
}
