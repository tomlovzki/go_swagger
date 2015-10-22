package go_swagger

import (
	"encoding/json"
	"testing"
)

func TestInitSwagger(t *testing.T) {
	info := `{"description": "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key special-key to test the authorization filters.",
    "version": "1.0.0",
    "title": "Swagger Petstore",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    }}`
	infoParse := map[string]interface{}{}
	json.Unmarshal([]byte(info), &infoParse)
	s := Swagger{
		SwaggerVersion: "2.0",
		Info:           infoParse,
		Host:           "rcp.dev.jxzy.com",
		BasePath:       "",
		Schemes:        []string{"http"},
	}
	h := NewH()
	h.InitSwagger(s)

	tag1 := Tag{Name: "course", Description: "course Description"}
	tag2 := Tag{Name: "tag2Name", Description: "tag2Name Description"}
	h.AddTag(tag1).AddTag(tag2)
	definition := Definition{
		Type: "object",
		Properties: map[string]Properties{
			"id": Properties{
				Type:   "integer",
				Format: "int64",
			}, "name": Properties{
				Type:        "string",
				Example:     "name exp",
				Description: "name desc",
			},
		},
		Xml: map[string]string{"name": "courseDetail"},
	}
	h.AddDefinition("courseDetail", definition)

	path := Path{
		Tags:        []string{"course"},
		Summary:     "course summary",
		Description: "course desc",
		OperationId: "course OperationId",
		Consumes:    []string{"application/json"},
		Produces:    []string{"application/json"},
		Parameters: []Parameters{
			Parameters{
				In:          "query",
				Name:        "id",
				Description: "courseId",
				Required:    true,
				Type:        "integer",
			},
		},
		Responses: map[string]interface{}{
			"200": map[string]interface{}{
				"description": "successful operation",
				"schema": map[string]string{
					"$ref": "#/definitions/courseDetail",
				},
			},
			"400": map[string]string{
				"description": "Invalid Order",
			},
		},
	}
	h.AddPath("/get-course", map[string]Path{"get": path})
	t.Log(VSwagger.ToString())
}
