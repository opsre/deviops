package docscmdb

const CmdbPaths = `
        "/api/v1/cmdb/database": {
            "put": {
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
                    "CMDB数据库"
                ],
                "summary": "修改数据库记录",
                "parameters": [
                    {
                        "description": "数据库信息(必须包含ID)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbSQL"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "创建数据库",
                "parameters": [
                    {
                        "description": "数据库信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbSQL"
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
            "delete": {
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
                    "CMDB数据库"
                ],
                "summary": "删除数据库记录",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据库ID（query参数）",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "description": "请求体（包含id字段）",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "type": "object"
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
        "/api/v1/cmdb/database/byname": {
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
                    "CMDB数据库"
                ],
                "summary": "根据名称查询数据库",
                "parameters": [
                    {
                        "type": "string",
                        "description": "数据库名称",
                        "name": "name",
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
        "/api/v1/cmdb/database/bytype": {
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
                    "CMDB数据库"
                ],
                "summary": "根据类型查询数据库",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据库类型(1=MySQL 2=PostgreSQL 3=Redis 4=MongoDB 5=Elasticsearch)",
                        "name": "type",
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
        "/api/v1/cmdb/database/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID获取数据库详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "根据ID获取数据库详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "数据库ID",
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
        "/api/v1/cmdb/databaselist": {
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
                    "CMDB数据库"
                ],
                "summary": "获取数据库列表[分页]",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
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
        "/api/v1/cmdb/groupadd": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新增资产分组接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "新增资产分组接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbGroup"
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
        "/api/v1/cmdb/groupbyname": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据名称查询资产分组",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "根据名称查询资产分组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分组名称",
                        "name": "name",
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
        "/api/v1/cmdb/groupdelete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除资产分组接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "删除资产分组接口",
                "parameters": [
                    {
                        "description": "分组ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbGroupIdDto"
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
        "/api/v1/cmdb/grouplist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询所有资产分组，并以树形结构返回",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "查询所有资产分组（树形结构）",
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
        "/api/v1/cmdb/grouplistwithhosts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询所有资产分组及关联主机，并以树形结构返回",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "查询所有资产分组及关联主机（树形结构）",
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
        "/api/v1/cmdb/groupupdate": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新资产分组接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "更新资产分组接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbGroup"
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
        "/api/v1/cmdb/hostbyip": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据IP查询主机(匹配内网IP、公网IP或SSH IP)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "根据IP查询主机",
                "parameters": [
                    {
                        "type": "string",
                        "description": "IP地址",
                        "name": "ip",
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
        "/api/v1/cmdb/hostbyname": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据主机名称模糊查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "根据主机名称模糊查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "主机名称(模糊匹配)",
                        "name": "name",
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
        "/api/v1/cmdb/hostbystatus": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据状态查询主机(1-\u003e认证成功,2-\u003e未认证,3-\u003e认证失败)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "根据状态查询主机",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "状态(1/2/3)",
                        "name": "status",
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
        "/api/v1/cmdb/hostcloudcreatealiyun": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建阿里云主机(通过阿里云API获取主机信息)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "创建阿里云主机",
                "parameters": [
                    {
                        "description": "阿里云主机信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCmdbHostCloudDto"
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
        "/api/v1/cmdb/hostcloudcreatebaidu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建百度云主机(通过百度云API自动扫描所有区域并获取主机信息)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "创建百度云主机",
                "parameters": [
                    {
                        "description": "百度云主机信息(AccessKey和SecretKey)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCmdbHostCloudDto"
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
        "/api/v1/cmdb/hostcloudcreatetencent": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建腾讯云主机(通过腾讯云API获取主机信息)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "创建腾讯云主机",
                "parameters": [
                    {
                        "description": "腾讯云主机信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCmdbHostCloudDto"
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
        "/api/v1/cmdb/hostcreate": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "创建主机",
                "parameters": [
                    {
                        "description": "主机信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCmdbHostDto"
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
        "/api/v1/cmdb/hostdelete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "删除主机",
                "parameters": [
                    {
                        "description": "主机ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbHostIdDto"
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
        "/api/v1/cmdb/hostgroup": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据分组ID获取主机列表（包括所有子分组的主机）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "根据分组ID获取主机列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分组ID",
                        "name": "groupId",
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
        "/api/v1/cmdb/hostimport": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "通过上传Excel模板批量导入主机（Excel列顺序：主机别名、SSH地址、SSH端口、SSH用户、备注）",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "从Excel导入主机",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Excel文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "分组ID",
                        "name": "groupId",
                        "in": "formData",
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
        "/api/v1/cmdb/hostinfo": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID获取主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "根据ID获取主机",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
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
        "/api/v1/cmdb/hostlist": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取主机列表(分页)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "获取主机列表(分页)",
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
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/cmdb/hostssh/command/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "在SSH终端执行命令",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB主机SSH"
                ],
                "summary": "执行SSH命令",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命令",
                        "name": "command",
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
        "/api/v1/cmdb/hostssh/upload/{id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "上传本地文件到远程SSH服务器",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB主机SSH"
                ],
                "summary": "上传文件到SSH服务器",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "要上传的文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "远程服务器目标路径",
                        "name": "destPath",
                        "in": "formData",
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
        "/api/v1/cmdb/hostsync": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据主机ID自动同步获取目标主机的基本信息(主机名称、操作系统、CPU、内存、磁盘、内网IP、公网IP)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "同步主机基本信息",
                "parameters": [
                    {
                        "description": "主机ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbHostIdDto"
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
        "/api/v1/cmdb/hosttemplate": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "下载主机导入Excel模板",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "下载主机导入模板",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        },
        "/api/v1/cmdb/hostupdate": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB资产管理"
                ],
                "summary": "更新主机",
                "parameters": [
                    {
                        "description": "主机信息(包含ID)",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateCmdbHostDto"
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
        "/api/v1/cmdb/sql": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "执行更新语句(通过数据库ID/名称)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "执行更新语句",
                "parameters": [
                    {
                        "description": "SQL更新请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
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
                "description": "执行插入语句(通过数据库ID/名称)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "执行插入语句",
                "parameters": [
                    {
                        "description": "SQL插入请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
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
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "执行删除语句(通过数据库ID/名称)\n执行原生SQL语句(通过数据库ID/名称)",
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库",
                    "CMDB数据库"
                ],
                "summary": "执行原生SQL语句",
                "parameters": [
                    {
                        "description": "SQL删除请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
                        }
                    },
                    {
                        "description": "SQL执行请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
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
        "/api/v1/cmdb/sql/databaselist": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定数据库实例的数据库列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "获取数据库列表",
                "parameters": [
                    {
                        "description": "数据库查询请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
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
        "/api/v1/cmdb/sql/execute": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "执行删除语句(通过数据库ID/名称)\n执行原生SQL语句(通过数据库ID/名称)",
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库",
                    "CMDB数据库"
                ],
                "summary": "执行原生SQL语句",
                "parameters": [
                    {
                        "description": "SQL删除请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
                        }
                    },
                    {
                        "description": "SQL执行请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
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
        "/api/v1/cmdb/sql/select": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "执行查询语句(通过数据库ID/名称)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "执行查询语句",
                "parameters": [
                    {
                        "description": "SQL查询请求",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SQLRequest"
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
        "/api/v1/cmdb/sqlLog/clean": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "清空SQL操作日志接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "清空SQL操作日志接口",
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
        "/api/v1/cmdb/sqlLog/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除SQL操作日志",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "根据id删除SQL操作日志",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CmdbSqlLogIdDto"
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
        "/api/v1/cmdb/sqlLog/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页获取SQL操作日志列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CMDB数据库"
                ],
                "summary": "分页获取SQL操作日志列表接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "执行用户",
                        "name": "execUser",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "开始时间",
                        "name": "beginTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "endTime",
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
        }`

const CmdbDefinitions = `
        "model.ClusterResourceMetrics": {
            "type": "object",
            "properties": {
                "available": {
                    "description": "可用量",
                    "type": "string"
                },
                "requestRate": {
                    "description": "请求率 (0-100)",
                    "type": "number"
                },
                "total": {
                    "description": "总量",
                    "type": "string"
                },
                "usageRate": {
                    "description": "使用率 (0-100)",
                    "type": "number"
                },
                "used": {
                    "description": "已使用",
                    "type": "string"
                }
            }
        },
        "model.CmdbGroup": {
            "type": "object",
            "properties": {
                "children": {
                    "description": "子分组（虚拟字段，用于树形展示）",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CmdbGroup"
                    }
                },
                "createTime": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "hostCount": {
                    "description": "主机数量（虚拟字段，包含所有子分组的主机数量）",
                    "type": "integer"
                },
                "hosts": {
                    "description": "关联的主机列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CmdbHost"
                    }
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "name": {
                    "description": "分组名称",
                    "type": "string"
                },
                "parentId": {
                    "description": "父级分组ID（0 表示根分组）",
                    "type": "integer"
                }
            }
        },
        "model.CmdbGroupIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.CmdbHost": {
            "type": "object",
            "properties": {
                "billingType": {
                    "type": "string"
                },
                "cpu": {
                    "type": "string"
                },
                "createTime": {
                    "$ref": "#/definitions/util.HTime"
                },
                "disk": {
                    "type": "string"
                },
                "expireTime": {
                    "$ref": "#/definitions/util.HTime"
                },
                "group": {
                    "$ref": "#/definitions/model.CmdbGroup"
                },
                "groupId": {
                    "type": "integer"
                },
                "hostName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "instanceId": {
                    "type": "string"
                },
                "memory": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "privateIp": {
                    "type": "string"
                },
                "publicIp": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "sshIp": {
                    "type": "string"
                },
                "sshKeyId": {
                    "type": "integer"
                },
                "sshName": {
                    "type": "string"
                },
                "sshPort": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "updateTime": {
                    "$ref": "#/definitions/util.HTime"
                },
                "vendor": {
                    "type": "integer"
                }
            }
        },
        "model.CmdbHostIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.CmdbSQL": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "所属账号ID",
                    "type": "integer"
                },
                "createdAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "description": {
                    "description": "描述/备注",
                    "type": "string"
                },
                "groupId": {
                    "description": "所属业务组ID",
                    "type": "integer"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "name": {
                    "description": "数据库名称",
                    "type": "string"
                },
                "tags": {
                    "description": "标签(多个标签用逗号分隔)",
                    "type": "string"
                },
                "type": {
                    "description": "数据库类型(1=MySQL 2=PostgreSQL 3=Redis 4=MongoDB 5=Elasticsearch)",
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
        "model.CmdbSqlLogIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.CreateCmdbHostCloudDto": {
            "type": "object",
            "required": [
                "accessKey",
                "groupId",
                "region",
                "secretKey",
                "vendor"
            ],
            "properties": {
                "accessKey": {
                    "description": "AK",
                    "type": "string"
                },
                "groupId": {
                    "description": "分组ID",
                    "type": "integer"
                },
                "instanceId": {
                    "description": "实例ID(可选)",
                    "type": "string"
                },
                "region": {
                    "description": "区域",
                    "type": "string"
                },
                "secretKey": {
                    "description": "SK",
                    "type": "string"
                },
                "vendor": {
                    "description": "云厂商:2-\u003e阿里云,3-\u003e腾讯云",
                    "type": "integer"
                }
            }
        },
        "model.CreateCmdbHostDto": {
            "type": "object",
            "required": [
                "groupId",
                "hostName",
                "sshIp",
                "sshKeyId",
                "sshName"
            ],
            "properties": {
                "groupId": {
                    "description": "主机分组ID",
                    "type": "integer"
                },
                "hostName": {
                    "description": "主机名称(唯一标识)",
                    "type": "string"
                },
                "remark": {
                    "description": "备注信息(可选)",
                    "type": "string"
                },
                "sshIp": {
                    "description": "SSH连接IP(公网或私网IP)",
                    "type": "string"
                },
                "sshKeyId": {
                    "description": "SSH凭据ID(从ecsAuth表获取)",
                    "type": "integer"
                },
                "sshName": {
                    "description": "SSH登录用户名",
                    "type": "string"
                },
                "sshPort": {
                    "description": "SSH端口(默认22)",
                    "type": "integer"
                }
            }
        },
        "model.CreateResourceQuotaRequest": {
            "type": "object",
            "required": [
                "hard",
                "name"
            ],
            "properties": {
                "hard": {
                    "description": "硬限制",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "ResourceQuota名称",
                    "type": "string"
                },
                "scopeSelector": {
                    "description": "作用域选择器",
                    "type": "object",
                    "additionalProperties": true
                },
                "scopes": {
                    "description": "作用域",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.OtherResources": {
            "type": "object",
            "properties": {
                "kafka": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "other": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "rabbitmq": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "redis": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "zookeeper": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.ResourceInfo": {
            "type": "object",
            "properties": {
                "allocatable": {
                    "description": "可分配量",
                    "type": "string"
                },
                "capacity": {
                    "description": "总容量",
                    "type": "string"
                },
                "requests": {
                    "description": "请求量",
                    "type": "string"
                },
                "usage": {
                    "description": "使用量",
                    "type": "string"
                }
            }
        },
        "model.ResourceQuotaDetail": {
            "type": "object",
            "properties": {
                "cpuQuota": {
                    "description": "CPU配额详情",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.QuotaInfo"
                        }
                    ]
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "hard": {
                    "description": "硬限制",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "memoryQuota": {
                    "description": "内存配额详情",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.QuotaInfo"
                        }
                    ]
                },
                "name": {
                    "description": "ResourceQuota名称",
                    "type": "string"
                },
                "storageQuota": {
                    "description": "存储配额详情",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.QuotaInfo"
                        }
                    ]
                },
                "used": {
                    "description": "已使用",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.ResourceQuotaListResponse": {
            "type": "object",
            "properties": {
                "resourceQuotas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ResourceQuotaDetail"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.ResourceSpec": {
            "type": "object",
            "properties": {
                "cpu": {
                    "description": "CPU",
                    "type": "string"
                },
                "memory": {
                    "description": "内存",
                    "type": "string"
                }
            }
        },
        "model.ResourceUsage": {
            "type": "object",
            "properties": {
                "cpu": {
                    "description": "CPU使用量 (如: \"100m\", \"1.5\")",
                    "type": "string"
                },
                "memory": {
                    "description": "内存使用量 (如: \"128Mi\", \"1Gi\")",
                    "type": "string"
                }
            }
        },
        "model.ResourceUsageRate": {
            "type": "object",
            "properties": {
                "cpuRate": {
                    "description": "CPU使用率 (百分比: 0-100)",
                    "type": "number"
                },
                "memoryRate": {
                    "description": "内存使用率 (百分比: 0-100)",
                    "type": "number"
                }
            }
        },
        "model.UpdateCmdbHostDto": {
            "type": "object",
            "required": [
                "groupId",
                "hostName",
                "sshIp",
                "sshKeyId",
                "sshName"
            ],
            "properties": {
                "groupId": {
                    "description": "主机分组ID",
                    "type": "integer"
                },
                "hostName": {
                    "description": "主机名称(唯一标识)",
                    "type": "string"
                },
                "id": {
                    "description": "主机ID",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注信息(可选)",
                    "type": "string"
                },
                "sshIp": {
                    "description": "SSH连接IP(公网或私网IP)",
                    "type": "string"
                },
                "sshKeyId": {
                    "description": "SSH凭据ID(从ecsAuth表获取)",
                    "type": "integer"
                },
                "sshName": {
                    "description": "SSH登录用户名",
                    "type": "string"
                },
                "sshPort": {
                    "description": "SSH端口(默认22)",
                    "type": "integer"
                },
                "vendor": {
                    "description": "厂商类型:1-\u003e自建,2-\u003e阿里云,3-\u003e腾讯云",
                    "type": "integer"
                }
            }
        },
        "model.UpdateResourceQuotaRequest": {
            "type": "object",
            "required": [
                "hard"
            ],
            "properties": {
                "hard": {
                    "description": "硬限制",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "scopeSelector": {
                    "description": "作用域选择器",
                    "type": "object",
                    "additionalProperties": true
                },
                "scopes": {
                    "description": "作用域",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.WorkloadResources": {
            "type": "object",
            "properties": {
                "limits": {
                    "description": "资源限制",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceSpec"
                        }
                    ]
                },
                "requests": {
                    "description": "资源请求",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceSpec"
                        }
                    ]
                }
            }
        }`
