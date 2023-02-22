// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
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
        "/api/members": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Members"
                ],
                "summary": "Index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Member"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Error"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee Number",
                        "name": "no",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Full Name",
                        "name": "full_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Profile Image",
                        "name": "profile_img",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kana Name",
                        "name": "kana_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Employment Start Date",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Emoployment End Date",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Employment Status",
                        "name": "employment_status_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Status",
                        "name": "status_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Member"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Error"
                        }
                    }
                }
            }
        },
        "/api/members/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Show",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Member"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Error"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Member"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee Number",
                        "name": "no",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Full Name",
                        "name": "full_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Profile Image",
                        "name": "profile_img",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kana Name",
                        "name": "kana_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Employment Start Date",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Emoployment End Date",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Employment Status",
                        "name": "employment_status_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Status",
                        "name": "status_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Member"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "domain.EmploymentStatus": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "employment_status": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Member": {
            "type": "object",
            "properties": {
                "biography": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "employment_status": {
                    "$ref": "#/definitions/domain.EmploymentStatus"
                },
                "employment_status_id": {
                    "type": "integer"
                },
                "end_date": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "kana_name": {
                    "type": "string"
                },
                "motto": {
                    "type": "string"
                },
                "no": {
                    "type": "string"
                },
                "profile_img": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "status_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Status": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
