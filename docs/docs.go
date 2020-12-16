// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/all": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "获取所有的Api 不分页",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "创建api",
                "parameters": [
                    {
                        "description": "api路径, api中文描述, api组, 方法",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "删除api",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/detail": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "根据id获取api",
                "parameters": [
                    {
                        "description": "根据id获取api",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DetailApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/list": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "分页获取API列表",
                "parameters": [
                    {
                        "description": "分页获取API列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SearchApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api"
                ],
                "summary": "创建基础api",
                "parameters": [
                    {
                        "description": "api路径, api中文描述, api组, 方法",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/captcha": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"验证码获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名, 密码, 验证码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"登陆成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/all": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "获取所有的菜单 不分页",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "创建菜单",
                "parameters": [
                    {
                        "description": "菜单名",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/delete": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "description": "ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/detail": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "根据id获取菜单信息",
                "parameters": [
                    {
                        "description": "id",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DetailMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/tree": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "分页获取菜单列表",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "更新菜单",
                "parameters": [
                    {
                        "description": "菜单名",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/bindApi": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "给角色绑定接口权限",
                "parameters": [
                    {
                        "description": "角色id, apiId",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BindRoleApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/bindMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "给角色绑定菜单权限",
                "parameters": [
                    {
                        "description": "角色id, 菜单id",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BindRoleMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "角色名，角色key",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/role/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "更新角色",
                "parameters": [
                    {
                        "description": "角色名，角色key",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/bindRole": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "给用户绑定角色权限",
                "parameters": [
                    {
                        "description": "用户id, 角色id",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BindUserRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"绑定成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "用户名, 昵称, 密码, 角色ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BindRoleApi": {
            "type": "object",
            "required": [
                "apiIds",
                "roleId"
            ],
            "properties": {
                "apiIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "roleId": {
                    "type": "integer"
                }
            }
        },
        "dto.BindRoleMenu": {
            "type": "object",
            "required": [
                "menuIds",
                "roleId"
            ],
            "properties": {
                "menuIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "roleId": {
                    "type": "integer"
                }
            }
        },
        "dto.BindUserRole": {
            "type": "object",
            "required": [
                "roleIds",
                "userId"
            ],
            "properties": {
                "roleIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateApi": {
            "type": "object",
            "properties": {
                "apiGroup": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "dto.CreateMenu": {
            "type": "object",
            "properties": {
                "component": {
                    "type": "string"
                },
                "defaultMenu": {
                    "type": "boolean"
                },
                "hidden": {
                    "type": "boolean"
                },
                "icon": {
                    "type": "string"
                },
                "keepAlive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.CreateRole": {
            "type": "object",
            "required": [
                "key",
                "name"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUser": {
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
        "dto.DeleteApi": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "dto.DeleteMenu": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "dto.DetailApi": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.DetailMenu": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "required": [
                "captcha",
                "captchaId",
                "password",
                "username"
            ],
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captchaId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.SearchApi": {
            "type": "object",
            "properties": {
                "apiGroup": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateApi": {
            "type": "object",
            "properties": {
                "apiGroup": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateRole": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
