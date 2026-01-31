package docsconfigcenter

const ConfigcenterPaths = `
        "/api/v1/config/accountauth": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据ID获取账号详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "账号ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.AccountAuth"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "更新账号认证信息",
                "parameters": [
                    {
                        "description": "账号认证信息",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AccountAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.AccountAuth"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "创建账号认证信息",
                "parameters": [
                    {
                        "description": "账号认证信息",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AccountAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.AccountAuth"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "删除账号认证信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "账号ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/accountauth/alias": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据别名查询账号",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号别名",
                        "name": "alias",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.AccountAuth"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/accountauth/decrypt": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "解密密码",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "账号ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/accountauth/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取账号列表（分页）",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/result.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.AccountAuth"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/accountauth/type": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据类型查询账号",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.AccountAuth"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/ecsauthadd": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "创建凭据",
                "parameters": [
                    {
                        "description": "凭据信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateEcsPasswordAuthDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/ecsauthdelete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "删除凭据",
                "parameters": [
                    {
                        "description": "凭据ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EcsAuthIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/ecsauthdetail": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据ID获取凭据详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "凭据ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.EcsAuthVo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/ecsauthinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据名称获取凭据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "凭据名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.EcsAuthVo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/ecsauthlist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取所有凭据（分页）",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/result.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.EcsAuthVo"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/ecsauthupdate": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "更新凭据",
                "parameters": [
                    {
                        "description": "凭据信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateEcsAuthDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/keymanage": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据ID获取密钥详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "密钥ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KeyManage"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "更新密钥管理信息",
                "parameters": [
                    {
                        "description": "密钥管理信息",
                        "name": "keyManage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateKeyManageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KeyManage"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "创建密钥管理信息",
                "parameters": [
                    {
                        "description": "密钥管理信息",
                        "name": "keyManage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateKeyManageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KeyManage"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "删除密钥管理信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "密钥ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/keymanage/decrypt": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "解密密钥信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "密钥ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/keymanage/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取密钥列表（分页）",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/result.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.KeyManage"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/keymanage/sync": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据密钥类型自动同步对应云厂商的主机",
                "parameters": [
                    {
                        "description": "同步参数，keyType: 1=阿里云,2=腾讯云,3=百度云",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "properties": {
                                "keyId": {
                                    "type": "integer"
                                },
                                "keyType": {
                                    "type": "integer"
                                }
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/keymanage/type": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据云厂商类型查询密钥",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "云厂商类型",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.KeyManage"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "根据ID获取定时同步配置详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "配置ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.SyncSchedule"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "更新定时同步配置",
                "parameters": [
                    {
                        "description": "定时同步配置信息",
                        "name": "syncSchedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSyncScheduleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.SyncSchedule"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "创建定时同步配置",
                "parameters": [
                    {
                        "description": "定时同步配置信息",
                        "name": "syncSchedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateSyncScheduleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.SyncSchedule"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "删除定时同步配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "配置ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule/active": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取所有启用的定时同步配置",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.SyncSchedule"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取定时同步配置列表（分页）",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/result.PageResult"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "list": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/model.SyncSchedule"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule/log": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取同步日志",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "配置ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "properties": {
                                                "lastRunTime": {
                                                    "type": "string"
                                                },
                                                "syncLog": {
                                                    "type": "string"
                                                }
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule/scheduler-stats": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "获取调度器状态信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule/toggle-status": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "切换配置状态（启用/禁用）",
                "parameters": [
                    {
                        "description": "配置ID和状态",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "properties": {
                                "id": {
                                    "type": "integer"
                                },
                                "status": {
                                    "type": "integer"
                                }
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/config/sync-schedule/trigger": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Config配置中心"
                ],
                "summary": "手动触发定时同步（测试用）",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "配置ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }`

const ConfigcenterDefinitions = `
        "model.AccountAuth": {
            "type": "object",
            "properties": {
                "alias": {
                    "description": "别名",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "host": {
                    "description": "IP地址",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "password": {
                    "description": "密码(加密存储)",
                    "type": "string"
                },
                "port": {
                    "description": "端口",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "type": {
                    "description": "类型",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                }
            }
        },
        "model.ConfigMapDetail": {
            "type": "object",
            "properties": {
                "binaryData": {
                    "description": "二进制数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer",
                            "format": "int32"
                        }
                    }
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "data": {
                    "description": "数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "immutable": {
                    "description": "是否不可变",
                    "type": "boolean"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "ConfigMap名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "usage": {
                    "description": "使用情况（哪些Pod在使用）",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.ConfigMapListResponse": {
            "type": "object",
            "properties": {
                "configMaps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sConfigMap"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.CreateConfigMapRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "binaryData": {
                    "description": "二进制数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer",
                            "format": "int32"
                        }
                    }
                },
                "data": {
                    "description": "数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "immutable": {
                    "description": "是否不可变",
                    "type": "boolean"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "ConfigMap名称",
                    "type": "string"
                }
            }
        },
        "model.RegistryConfig": {
            "type": "object",
            "properties": {
                "privateRegistry": {
                    "description": "私有镜像仓库地址",
                    "type": "string"
                },
                "registryPassword": {
                    "description": "镜像仓库密码",
                    "type": "string"
                },
                "registryUsername": {
                    "description": "镜像仓库用户名",
                    "type": "string"
                },
                "usePrivateRegistry": {
                    "description": "是否使用私有仓库",
                    "type": "boolean"
                }
            }
        },
        "model.UpdateConfigMapRequest": {
            "type": "object",
            "properties": {
                "binaryData": {
                    "description": "二进制数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer",
                            "format": "int32"
                        }
                    }
                },
                "data": {
                    "description": "数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        }`
