// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "contact": {
            "name": "API Support",
            "email": "io.labs.development@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/signin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Sign Into Account",
                "parameters": [
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "email_validated",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "fk_account_id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "temp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.SignInResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Creates an account on the service",
                "parameters": [
                    {
                        "type": "string",
                        "name": "account_name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "passwordconfirm",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "temp",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/Media/{id}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Creates Media",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for Media account",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Deletes Media by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/account/{id}/pages": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Gets all pages from Account ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Page"
                            }
                        }
                    }
                }
            }
        },
        "/v1/account/{id}/projects": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Gets projects by account ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Project"
                            }
                        }
                    }
                }
            }
        },
        "/v1/content/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Gets content from ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Content"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Updates content by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "account_id",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "active",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "anchor",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "1,073,741,824",
                        "name": "content",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "copywrite",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "created",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "draft",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "page_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "updated",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "weight",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Creates content",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Deletes content by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/media/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Media"
                ],
                "summary": "Gets media by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Media"
                        }
                    }
                }
            }
        },
        "/v1/page/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Returns a page by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Page"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Updates a page by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "absolute_url",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "active",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "created",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "updated",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Creates a page",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "absolute_url",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "active",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "created",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "updated",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Deletes a page by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/project/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Gets project by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Project"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Updates a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Creates a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "Deletes a Project by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/project/{id}/page": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Creates a page and inserts into project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for page",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "absolute_url",
                        "in": "formData"
                    },
                    {
                        "type": "boolean",
                        "name": "active",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "created",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "updated",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/project/{id}/pages": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Page"
                ],
                "summary": "Gets all pages from project ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header as token",
                        "name": "Authentication",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID for project",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Page"
                            }
                        }
                    }
                }
            }
        },
        "/v1/signup/{uid}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "verify"
                ],
                "summary": "Verifies an email with a UID sent to the account email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uid from email",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Content": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "active": {
                    "type": "boolean"
                },
                "anchor": {
                    "type": "string"
                },
                "content": {
                    "description": "1,073,741,824",
                    "type": "string"
                },
                "copywrite": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "draft": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "page_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "model.Media": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "fk_account_id": {
                    "type": "integer"
                },
                "fk_media_wrapper_id": {
                    "type": "integer"
                },
                "fk_record_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "source": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.Page": {
            "type": "object",
            "properties": {
                "absolute_url": {
                    "type": "string"
                },
                "active": {
                    "type": "boolean"
                },
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.Project": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "updated": {
                    "type": "string"
                }
            }
        },
        "model.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
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
	Host:        "api.engin",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Engin",
	Description: "",
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
	swag.Register("swagger", &s{})
}
