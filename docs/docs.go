// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2021-12-07 03:55:33.479709 +0300 +03 m=+17.676862660
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/currencies": {
            "get": {
                "description": "List all cryptocurrencies portfolio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "List All Cryptocurrencies",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/common.ApiResponse"
                        }
                    }
                }
            }
        },
        "/currency": {
            "put": {
                "description": "Create cryptocurrency portfolio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "Create Cryptocurrency",
                "parameters": [
                    {
                        "description": "Example Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateAndUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.ApiResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.CreateAndUpdate"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/currency/{id}": {
            "get": {
                "description": "Read cryptocurrency portfolio by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "Read Cryptocurrency",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cryptocurrency id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/common.ApiResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete cryptocurrency portfolio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "Delete Cryptocurrency",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cryptocurrency id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/common.ApiResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update cryptocurrency portfolio by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "Update Cryptocurrency",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cryptocurrency id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Example Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateAndUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/common.ApiResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "description": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "request.CreateAndUpdate": {
            "type": "object",
            "required": [
                "amount",
                "code"
            ],
            "properties": {
                "amount": {
                    "description": "Amount of the code",
                    "type": "integer",
                    "x-order-2": true,
                    "example": 3
                },
                "code": {
                    "description": "Symbol of the cryptocurrency",
                    "type": "string",
                    "x-order-1": true,
                    "example": "BTC"
                }
            }
        },
        "response.CreateAndUpdate": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "Amount of the code",
                    "type": "integer",
                    "x-order-3": true,
                    "example": 4
                },
                "code": {
                    "description": "Code of the cryptocurrency portfolio",
                    "type": "string",
                    "x-order-2": true,
                    "example": "BTC"
                },
                "id": {
                    "description": "ID of the cryptocurrency portfolio",
                    "type": "string",
                    "x-order-1": true,
                    "example": "61ae85f3b45c25aa9cdaba99"
                },
                "price": {
                    "description": "Price of the code and multiplied amount",
                    "type": "number",
                    "x-order-4": true,
                    "example": 4900.01
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "CRYPTOCURRENCY PORTFOLIO",
	Description: "This is a sample CRUD operations on currency system.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
