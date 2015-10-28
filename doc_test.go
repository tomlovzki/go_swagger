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
	h.AddDefinition("courseDetail", definition).AddDefinition("genByJsonItem", Definition{
		Type: "object",
		Properties: GenByJson(`{
    "bankId": {
        "type": "integer",
        "example": "44655",
        "description": "课程对应题库id"
    },
    "cid": {
        "type": "integer",
        "example": 29786,
        "description": "课程id"
    },
    "courseScore": {
        "type": "integer",
        "example": 0,
        "description": "课程得分"
    },
    "courseType": {
        "type": "string",
        "example": "10",
        "description": "课程类型 10课程 20题库"
    },
    "credit": {
        "type": "integer",
        "example": 2.5,
        "description": "课程学分"
    },
    "curSec": {
        "type": "object",
        "example": {
            "rows": 0,
            "sid": 13058
        },
        "description": "当前正在学习的章节，如果没有此字段，则没有阅读记录，rows:行数，sid章节id"
    },
    "description": {
        "type": "string",
        "example": "description",
        "description": "课程description"
    },
    "detail": {
        "type": "object",
        "example": {
            "cType": 1,
            "chapterCnt": 6,
            "credit": 0,
            "credit2": 0,
            "joinCount": 13,
            "paperCnt": 0,
            "sectionCnt": 6,
            "startTime": "0000-00-00 00:00:00"
        },
        "description": "课程附加属性:{\n      \"cType\": 1,\n      \"chapterCnt\": 6,章数量\n      \"credit\": 0,课程学分\n      \"credit2\": 0,获得学分\n      \"joinCount\": 13,参与人数\n      \"paperCnt\": 0,试卷数量\n      \"sectionCnt\": 6,节数量\n      \"startTime\": \"0000-00-00 00:00:00\" 开始时间\n    }"
    },
    "id": {
        "type": "integer",
        "example": 2982,
        "description": "课程购买id"
    },
    "imgs": {
        "type": "string",
        "example": "http://u.dev.jxzy.com/zT=q0x==",
        "description": "课程封面"
    },
    "name": {
        "type": "string",
        "example": "课程名称",
        "description": "课程名称"
    },
    "process": {
        "type": "string",
        "example": 0.83,
        "description": "课程学习进度"
    },
    "status": {
        "type": "string",
        "example": "ONSELL",
        "description": "课程状态"
    },
    "uid": {
        "type": "integer",
        "example": 5,
        "description": "课程创建者id"
    },
    "uname": {
        "type": "string",
        "example": "liuqg",
        "description": "课程创建者名称"
    }
}`),
		Xml: map[string]string{"name": "genByJsonItem"},
	})

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
