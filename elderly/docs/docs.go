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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "liu",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/dht22": {
            "get": {
                "description": "查询温湿度列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "温湿度列表信息接口"
                ],
                "summary": "查询温湿度列表信息接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseDhtList"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户登录接口"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "登录参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseLoginList"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户注册接口"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "注册参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseSignUpList"
                        }
                    }
                }
            }
        },
        "/smooth_led": {
            "get": {
                "description": "查询烟雾灯光列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "烟雾灯光列表信息接口"
                ],
                "summary": "查询烟雾灯光列表信息接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseLedList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResLogin": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "controller._ResponseDhtList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Dht"
                    }
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponseLedList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Led"
                    }
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponseLoginList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.ResLogin"
                    }
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "controller._ResponseSignUpList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.Dht": {
            "type": "object",
            "properties": {
                "humidity": {
                    "type": "integer"
                },
                "sense_time": {
                    "type": "string"
                },
                "temperature": {
                    "type": "integer"
                }
            }
        },
        "models.Led": {
            "type": "object",
            "properties": {
                "sense_time": {
                    "type": "string"
                },
                "states": {
                    "type": "string"
                },
                "vlight": {
                    "type": "integer"
                },
                "vsmooth": {
                    "type": "integer"
                }
            }
        },
        "models.ParamLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.ParamSignUp": {
            "type": "object",
            "required": [
                "password",
                "re_password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "re_password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "8.130.39.242:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "老人智能陪伴系统接口文档",
	Description:      "老人智能陪伴系统",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}