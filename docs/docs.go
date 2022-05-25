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
        "/api/admin/admit": {
            "put": {
                "description": "管理员录取/取消录取面试者到特定组别",
                "tags": [
                    "admin"
                ],
                "summary": "录取/取消录取",
                "parameters": [
                    {
                        "description": "录取",
                        "name": "admit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Admit"
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
            },
            "delete": {
                "description": "管理员录取/取消录取面试者到特定组别",
                "tags": [
                    "admin"
                ],
                "summary": "录取/取消录取",
                "parameters": [
                    {
                        "description": "录取",
                        "name": "admit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Admit"
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
        "/api/admin/applicant/{id}": {
            "get": {
                "description": "管理员获取面试者信息",
                "tags": [
                    "admin"
                ],
                "summary": "面试者信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "面试者ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Applicant"
                        }
                    },
                    "401": {
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
        },
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
                            "$ref": "#/definitions/mid.Admin"
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
        "/api/admin/login/": {
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
        "/api/admin/public/exam/": {
            "get": {
                "description": "获取特定组别的面试题库",
                "tags": [
                    "public"
                ],
                "summary": "获取面试题库",
                "parameters": [
                    {
                        "enum": [
                            "机械",
                            "电控",
                            "视觉"
                        ],
                        "type": "string",
                        "description": "组别",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Question"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/api/admin/public/time/": {
            "get": {
                "description": "获取特定组别的面试时间",
                "tags": [
                    "public"
                ],
                "summary": "获取面试时间",
                "parameters": [
                    {
                        "enum": [
                            "机械",
                            "电控",
                            "视觉"
                        ],
                        "type": "string",
                        "description": "组别",
                        "name": "group",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.OptionalTime"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/api/admin/public/{key}": {
            "get": {
                "description": "可获取\"form\", \"announce\", \"time-frame\"",
                "tags": [
                    "public",
                    "setting"
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
        },
        "/api/admin/question": {
            "post": {
                "description": "管理员添加面试题目",
                "tags": [
                    "admin"
                ],
                "summary": "添加题目",
                "parameters": [
                    {
                        "description": "题目",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Question"
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
        "/api/admin/question/{id}": {
            "put": {
                "description": "管理员修改面试题目",
                "tags": [
                    "admin"
                ],
                "summary": "修改题目",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "题目ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "题目，仅question本身，其他内容将被忽略",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
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
            },
            "delete": {
                "description": "管理员删除面试题目",
                "tags": [
                    "admin"
                ],
                "summary": "删除题目",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "题目ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        "/api/admin/remark/{id}": {
            "get": {
                "description": "管理员获取对特定面试者的评价",
                "tags": [
                    "admin"
                ],
                "summary": "获取评价",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "面试者ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/model.Remark"
                        }
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
            },
            "put": {
                "description": "管理员修改对特定面试者的评价",
                "tags": [
                    "admin"
                ],
                "summary": "修改评价",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "面试者ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "评价，仅remark本身，其他内容将被忽略",
                        "name": "remark",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Remark"
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
        "/api/admin/score": {
            "delete": {
                "description": "管理员删除面试者分数",
                "tags": [
                    "admin"
                ],
                "summary": "删除分数",
                "parameters": [
                    {
                        "description": "分数",
                        "name": "score",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Score"
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
        "/api/admin/score/{id}": {
            "get": {
                "description": "管理员获取面试者分数",
                "tags": [
                    "admin"
                ],
                "summary": "获取分数",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "面试者ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Score"
                        }
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
        "/api/admin/standard": {
            "get": {
                "description": "管理员获取所有评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "所有标准",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Standard"
                            }
                        }
                    },
                    "401": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "put": {
                "description": "设置默认评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "设置默认标准",
                "parameters": [
                    {
                        "description": "仅选取评价标准的ID",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Admin"
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
            },
            "post": {
                "description": "管理员添加评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "添加标准",
                "parameters": [
                    {
                        "description": "标准",
                        "name": "standard",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Standard"
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
        "/api/admin/standard/{id}": {
            "get": {
                "description": "管理员获取特定评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "单个标准",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "评价标准ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Standard"
                        }
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
            },
            "put": {
                "description": "管理员更新特定评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "更新标准",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "评价标准ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "标准",
                        "name": "standard",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Standard"
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
            },
            "delete": {
                "description": "管理员删除特定评价标准",
                "tags": [
                    "admin"
                ],
                "summary": "删除标准",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "评价标准ID",
                        "name": "id",
                        "in": "path",
                        "required": true
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
        "/api/admin/time/{group}": {
            "put": {
                "description": "管理员上传面试时间csv文件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "设置面试时间文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "组",
                        "name": "group",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "csv 文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/api/applicant/info": {
            "get": {
                "description": "面试者获取自己的微信信息",
                "tags": [
                    "applicant"
                ],
                "summary": "面试者微信信息",
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
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/api/applicant/login/": {
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
        "/api/setting/announce": {
            "put": {
                "description": "修改公告设置",
                "tags": [
                    "admin",
                    "setting"
                ],
                "summary": "设置公告",
                "parameters": [
                    {
                        "description": "公告",
                        "name": "announce",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ctrl.announce"
                        }
                    }
                ],
                "responses": {
                    "200": {
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
        "/api/setting/form": {
            "put": {
                "description": "修改报名表设置",
                "tags": [
                    "admin",
                    "setting"
                ],
                "summary": "设置报名表",
                "parameters": [
                    {
                        "description": "报名表",
                        "name": "form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ctrl.form"
                        }
                    }
                ],
                "responses": {
                    "200": {
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
        "/api/setting/time-frame": {
            "put": {
                "description": "修改时间节点设置",
                "tags": [
                    "admin",
                    "setting"
                ],
                "summary": "设置时间节点",
                "parameters": [
                    {
                        "description": "时间节点",
                        "name": "timeframe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ctrl.timeframe"
                        }
                    }
                ],
                "responses": {
                    "200": {
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
        }
    },
    "definitions": {
        "ctrl.announce": {
            "type": "object",
            "properties": {
                "admitted": {
                    "type": "string"
                },
                "failed": {
                    "type": "string"
                },
                "haventAppliedForm": {
                    "type": "string"
                },
                "haventInterview": {
                    "type": "string"
                },
                "haventSelectedTime": {
                    "type": "string"
                },
                "interviewed": {
                    "type": "string"
                }
            }
        },
        "ctrl.field": {
            "type": "object",
            "required": [
                "key",
                "name",
                "regexp",
                "required",
                "type"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "regexp": {
                    "type": "string"
                },
                "required": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "ctrl.form": {
            "type": "object",
            "required": [
                "fields",
                "intent"
            ],
            "properties": {
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ctrl.field"
                    }
                },
                "intent": {
                    "$ref": "#/definitions/ctrl.intent"
                }
            }
        },
        "ctrl.intent": {
            "type": "object",
            "required": [
                "max",
                "min",
                "parallel"
            ],
            "properties": {
                "max": {
                    "type": "integer"
                },
                "min": {
                    "type": "integer"
                },
                "parallel": {
                    "type": "boolean"
                }
            }
        },
        "ctrl.timeframe": {
            "type": "object",
            "required": [
                "done",
                "form-end",
                "form-start",
                "time-end",
                "time-start"
            ],
            "properties": {
                "done": {
                    "type": "string"
                },
                "form-end": {
                    "type": "string"
                },
                "form-start": {
                    "type": "string"
                },
                "time-end": {
                    "type": "string"
                },
                "time-start": {
                    "type": "string"
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
                "standard_id": {
                    "type": "integer"
                }
            }
        },
        "model.Admit": {
            "type": "object",
            "properties": {
                "applicantID": {
                    "type": "integer"
                },
                "group": {
                    "type": "string"
                }
            }
        },
        "model.Applicant": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "form": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "intents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Intent"
                    }
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.Intent": {
            "type": "object",
            "properties": {
                "applicant": {
                    "$ref": "#/definitions/model.Applicant"
                },
                "applicantID": {
                    "type": "integer"
                },
                "group": {
                    "type": "string"
                },
                "intentRank": {
                    "description": "nil 为平行志愿",
                    "type": "integer"
                },
                "optionalTime": {
                    "$ref": "#/definitions/model.OptionalTime"
                },
                "optionalTimeID": {
                    "type": "integer"
                }
            }
        },
        "model.OptionalTime": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "intentRank": {
                    "description": "限定面试轮次",
                    "type": "integer"
                },
                "theDate": {
                    "type": "string"
                },
                "theLocation": {
                    "type": "string"
                },
                "theTime": {
                    "type": "string"
                }
            }
        },
        "model.Question": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                }
            }
        },
        "model.Remark": {
            "type": "object",
            "properties": {
                "applicantID": {
                    "type": "integer"
                },
                "theRemark": {
                    "type": "string"
                }
            }
        },
        "model.Score": {
            "type": "object",
            "properties": {
                "adminID": {
                    "type": "integer"
                },
                "applicantID": {
                    "type": "integer"
                },
                "evaluationDetails": {
                    "type": "string"
                },
                "group": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                },
                "standardID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
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
