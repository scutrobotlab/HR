// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/admin/info": {
            "get": {
                "description": "获取管理员信息",
                "tags": [
                    "admin"
                ],
                "summary": "管理员信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Admin"
                        }
                    },
                    "401": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/api/admin/login/{code}": {
            "post": {
                "description": "管理员登录",
                "tags": [
                    "admin"
                ],
                "summary": "管理员登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oauth2 code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/mid.Admin"
                        }
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/api/admin/standard": {
            "put": {
                "description": "设置默认评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "设置默认标准",
                "parameters": [
                    {
                        "description": "评价标准的ID",
                        "name": "Standard",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ctrl.standard"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/api/applicant/login/{token}": {
            "post": {
                "description": "面试者登录",
                "tags": [
                    "applicant"
                ],
                "summary": "面试者登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "oauth2 token",
                        "name": "token",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/mid.Applicant"
                        }
                    },
                    "401": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/api/public/{key}": {
            "get": {
                "description": "可获取\"form\", \"announce\", \"time-frame\"",
                "tags": [
                    "public"
                ],
                "summary": "获取设置",
                "parameters": [
                    {
                        "enum": [
                            "form",
                            "announce",
                            "time-frame"
                        ],
                        "type": "string",
                        "description": "获取设置的键",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "ctrl.standard": {
            "type": "object",
            "required": [
                "standard_id"
            ],
            "properties": {
                "standard_id": {
                    "type": "integer"
                }
            }
        },
        "mid.Admin": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/mid.Group"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "mid.Applicant": {
            "type": "object",
            "properties": {
                "headimgurl": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "openid": {
                    "type": "string"
                }
            }
        },
        "mid.Group": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Admin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "standard": {
                    "$ref": "#/definitions/model.Standard"
                },
                "standardID": {
                    "type": "integer"
                }
            }
        },
        "model.Standard": {
            "type": "object",
            "properties": {
                "adminID": {
                    "description": "上次修改的管理员",
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "name": {
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
