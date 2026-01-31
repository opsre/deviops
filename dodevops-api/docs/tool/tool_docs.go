package docstool

const ToolPaths = `
        "/api/v1/tool": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新导航工具接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool导航工具"
                ],
                "summary": "更新导航工具",
                "parameters": [
                    {
                        "description": "导航工具信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateToolDto"
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
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建导航工具接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool导航工具"
                ],
                "summary": "创建导航工具",
                "parameters": [
                    {
                        "description": "导航工具信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddToolDto"
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
        "/api/v1/tool/all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取所有导航工具（不分页）",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool导航工具"
                ],
                "summary": "获取所有导航工具",
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
        "/api/v1/tool/deploy": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建一个服务部署任务，异步执行部署",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool运维工具箱"
                ],
                "summary": "创建部署任务",
                "parameters": [
                    {
                        "description": "部署参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateDeployDto"
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
        "/api/v1/tool/deploy/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取服务部署历史记录列表（分页）",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool运维工具箱"
                ],
                "summary": "获取部署历史列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务名称（模糊查询）",
                        "name": "serviceName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "主机ID",
                        "name": "hostId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "状态: 0=部署中, 1=运行中, 2=已停止, 3=部署失败",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
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
        "/api/v1/tool/deploy/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "停止并删除已部署的服务",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool运维工具箱"
                ],
                "summary": "卸载服务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "部署ID",
                        "name": "id",
                        "in": "path",
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
        "/api/v1/tool/deploy/{id}/status": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据部署ID获取部署状态和日志",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool运维工具箱"
                ],
                "summary": "获取部署状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "部署ID",
                        "name": "id",
                        "in": "path",
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
        "/api/v1/tool/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取导航工具列表（分页）",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool导航工具"
                ],
                "summary": "获取导航工具列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标题（模糊查询）",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "pageSize",
                        "in": "query"
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
        "/api/v1/tool/services": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取所有可部署的服务列表及分类",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool运维工具箱"
                ],
                "summary": "获取可部署服务列表",
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
        "/api/v1/tool/services/{serviceId}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据服务ID获取服务的详细信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool运维工具箱"
                ],
                "summary": "获取服务详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "服务ID",
                        "name": "serviceId",
                        "in": "path",
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
        "/api/v1/tool/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID获取导航工具详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool导航工具"
                ],
                "summary": "根据ID获取导航工具",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "工具ID",
                        "name": "id",
                        "in": "path",
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
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除导航工具接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tool导航工具"
                ],
                "summary": "删除导航工具",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "工具ID",
                        "name": "id",
                        "in": "path",
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

const ToolDefinitions = `
        "model.AddToolDto": {
            "type": "object",
            "required": [
                "link",
                "title"
            ],
            "properties": {
                "icon": {
                    "description": "导航图标",
                    "type": "string"
                },
                "link": {
                    "description": "链接地址",
                    "type": "string",
                    "maxLength": 500
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "title": {
                    "description": "导航标题",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        },
        "model.UpdateToolDto": {
            "type": "object",
            "required": [
                "id",
                "link",
                "title"
            ],
            "properties": {
                "icon": {
                    "description": "导航图标",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "link": {
                    "description": "链接地址",
                    "type": "string",
                    "maxLength": 500
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "title": {
                    "description": "导航标题",
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 1
                }
            }
        }`
