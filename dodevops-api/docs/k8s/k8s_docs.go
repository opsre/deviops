package docsk8s

const K8sPaths = `
        "/api/v1/k8s/cluster": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页获取K8s集群列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "获取K8s集群列表",
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
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KubeClusterListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "创建K8s集群，可选择是否自动部署",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "创建K8s集群",
                "parameters": [
                    {
                        "description": "集群创建参数",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateKubeClusterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KubeCluster"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据集群ID获取集群详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "获取K8s集群详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KubeCluster"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
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
                "description": "更新集群的基本信息（名称、描述、kubeconfig等）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "更新K8s集群信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "集群更新参数",
                        "name": "cluster",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateKubeClusterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KubeCluster"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "删除指定的K8s集群（只能删除已停止的集群）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "删除K8s集群",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误或集群状态不允许删除",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/detail": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取集群的完整详细信息，包括节点、工作负载、组件、网络配置、监控信息等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "获取K8s集群详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ClusterDetailResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取整个集群的事件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s事件管理"
                ],
                "summary": "获取集群事件列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "限制返回数量",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.EventListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定集群的所有命名空间信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "获取K8s命名空间列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.NamespaceListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "在指定集群中创建新的命名空间",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "创建K8s命名空间",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "命名空间信息",
                        "name": "namespace",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateNamespaceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.K8sNamespace"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces/{namespaceName}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定命名空间的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "获取K8s命名空间详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.K8sNamespace"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
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
                "description": "更新指定命名空间的标签和注释",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "更新K8s命名空间",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新信息",
                        "name": "namespace",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateNamespaceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.K8sNamespace"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "删除指定的命名空间（会同时删除其中的所有资源）",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "删除K8s命名空间",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces/{namespaceName}/events": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定命名空间的事件列表，支持按资源类型和名称过滤",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s事件管理"
                ],
                "summary": "获取K8s事件列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "Pod",
                            "Deployment",
                            "StatefulSet",
                            "DaemonSet",
                            "Service"
                        ],
                        "type": "string",
                        "description": "资源类型",
                        "name": "kind",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "资源名称",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 100,
                        "description": "限制返回数量",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.EventListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces/{namespaceName}/limitranges": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定命名空间的所有默认资源限制",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "获取默认资源限制列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.LimitRangeListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "为指定命名空间创建默认资源限制",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "创建默认资源限制",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "默认资源限制信息",
                        "name": "limitrange",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateLimitRangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.LimitRangeDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces/{namespaceName}/limitranges/{limitRangeName}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新指定的默认资源限制",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "更新默认资源限制",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "LimitRange名称",
                        "name": "limitRangeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新信息",
                        "name": "limitrange",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateLimitRangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.LimitRangeDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "默认资源限制不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "删除指定的默认资源限制",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "删除默认资源限制",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "LimitRange名称",
                        "name": "limitRangeName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "默认资源限制不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces/{namespaceName}/resourcequotas": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定命名空间的所有资源配额",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "获取资源配额列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResourceQuotaListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "为指定命名空间创建资源配额",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "创建资源配额",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "资源配额信息",
                        "name": "quota",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateResourceQuotaRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResourceQuotaDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "命名空间不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/namespaces/{namespaceName}/resourcequotas/{quotaName}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新指定的资源配额",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "更新资源配额",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ResourceQuota名称",
                        "name": "quotaName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新信息",
                        "name": "quota",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateResourceQuotaRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ResourceQuotaDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "资源配额不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "删除指定的资源配额",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s命名空间管理"
                ],
                "summary": "删除资源配额",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "命名空间名称",
                        "name": "namespaceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ResourceQuota名称",
                        "name": "quotaName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "资源配额不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定集群的所有节点详细信息，包括名称/IP地址、状态、配置、资源使用情况等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "获取K8s节点信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
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
                                                "$ref": "#/definitions/model.K8sNode"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定节点的详细信息，包括容器组、资源使用详情等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "获取单个节点详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.K8sNodeDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}/cordon": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "设置节点的可调度状态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "封锁/解封节点",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "封锁信息",
                        "name": "cordon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CordonNodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}/drain": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "驱逐节点上的所有Pod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "驱逐节点",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "驱逐配置",
                        "name": "drain",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DrainNodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "驱逐成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}/enhanced": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取节点的完整详细信息，包括基本信息、系统信息、K8s组件版本、资源使用情况、监控信息、Pod列表等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "获取增强的节点详细信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.NodeDetailResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}/labels": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "为指定节点添加标签",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "为节点添加标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "标签信息",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddLabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "添加成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "移除指定节点的标签",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "移除节点标签",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "标签信息",
                        "name": "label",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RemoveLabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "移除成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}/resources": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取节点的资源容量、分配情况和Pod资源使用详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "获取节点资源分配详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.NodeResourceAllocation"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/nodes/{nodeName}/taints": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "为指定节点添加污点，控制Pod调度",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "为节点添加污点",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "污点信息",
                        "name": "taint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddTaintRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "添加成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
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
                "description": "移除指定节点的污点",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s节点管理"
                ],
                "summary": "移除节点污点",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "节点名称",
                        "name": "nodeName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "污点信息",
                        "name": "taint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RemoveTaintRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "移除成功",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "节点不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/status": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取集群运行状态和节点信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "获取K8s集群状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
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
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/k8s/cluster/{id}/sync": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "通过K8s API同步集群版本、节点数量、集群状态等信息并更新数据库",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s集群管理"
                ],
                "summary": "同步K8s集群信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "同步成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.KubeCluster"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "404": {
                        "description": "集群不存在",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }`

const K8sDefinitions = `
        "model.ApplicationForDeployment": {
            "type": "object",
            "properties": {
                "can_deploy": {
                    "description": "是否可以部署",
                    "type": "boolean"
                },
                "code": {
                    "description": "应用编码",
                    "type": "string"
                },
                "environment": {
                    "description": "环境名称",
                    "type": "string"
                },
                "id": {
                    "description": "应用ID",
                    "type": "integer"
                },
                "jenkins_env_id": {
                    "description": "Jenkins环境配置ID",
                    "type": "integer"
                },
                "job_name": {
                    "description": "Jenkins任务名称",
                    "type": "string"
                },
                "name": {
                    "description": "应用名称",
                    "type": "string"
                },
                "reason": {
                    "description": "不能部署的原因",
                    "type": "string"
                }
            }
        },
        "model.BusinessLineServiceTree": {
            "type": "object",
            "properties": {
                "business_group_id": {
                    "description": "业务组ID",
                    "type": "integer"
                },
                "business_group_name": {
                    "description": "业务组名称",
                    "type": "string"
                },
                "service_count": {
                    "description": "服务数量",
                    "type": "integer"
                },
                "services": {
                    "description": "服务列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServiceTreeNode"
                    }
                }
            }
        },
        "model.ContainerInfo": {
            "type": "object",
            "properties": {
                "env": {
                    "description": "环境变量",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EnvVar"
                    }
                },
                "image": {
                    "description": "镜像",
                    "type": "string"
                },
                "name": {
                    "description": "容器名称",
                    "type": "string"
                },
                "ports": {
                    "description": "端口配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerPort"
                    }
                },
                "ready": {
                    "description": "就绪状态",
                    "type": "boolean"
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadResources"
                        }
                    ]
                },
                "restartCount": {
                    "description": "重启次数",
                    "type": "integer"
                },
                "state": {
                    "description": "状态",
                    "type": "string"
                }
            }
        },
        "model.ContainerMetricsInfo": {
            "type": "object",
            "properties": {
                "limits": {
                    "description": "资源限制量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "name": {
                    "description": "容器名称",
                    "type": "string"
                },
                "requests": {
                    "description": "资源请求量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "restartCount": {
                    "description": "重启次数",
                    "type": "integer"
                },
                "state": {
                    "description": "容器状态",
                    "type": "string"
                },
                "usage": {
                    "description": "资源使用量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "usageRate": {
                    "description": "使用率",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsageRate"
                        }
                    ]
                }
            }
        },
        "model.ContainerPort": {
            "type": "object",
            "properties": {
                "containerPort": {
                    "description": "容器端口",
                    "type": "integer"
                },
                "name": {
                    "description": "端口名称",
                    "type": "string"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                }
            }
        },
        "model.ContainerSpec": {
            "type": "object",
            "required": [
                "image",
                "name"
            ],
            "properties": {
                "args": {
                    "description": "启动参数",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "command": {
                    "description": "启动命令",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "env": {
                    "description": "环境变量",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EnvVar"
                    }
                },
                "image": {
                    "description": "镜像",
                    "type": "string"
                },
                "name": {
                    "description": "容器名称",
                    "type": "string"
                },
                "ports": {
                    "description": "端口配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerPort"
                    }
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadResources"
                        }
                    ]
                },
                "volumeMounts": {
                    "description": "存储卷挂载",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VolumeMount"
                    }
                }
            }
        },
        "model.ContainerStatus": {
            "type": "object",
            "properties": {
                "image": {
                    "description": "镜像",
                    "type": "string"
                },
                "name": {
                    "description": "容器名称",
                    "type": "string"
                },
                "ready": {
                    "description": "就绪状态",
                    "type": "boolean"
                },
                "restartCount": {
                    "description": "重启次数",
                    "type": "integer"
                },
                "state": {
                    "description": "状态",
                    "type": "string"
                }
            }
        },
        "model.CordonNodeRequest": {
            "type": "object",
            "properties": {
                "reason": {
                    "type": "string"
                },
                "unschedulable": {
                    "type": "boolean"
                }
            }
        },
        "model.CreateDeploymentRequest": {
            "type": "object",
            "required": [
                "name",
                "template"
            ],
            "properties": {
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Deployment名称",
                    "type": "string"
                },
                "replicas": {
                    "description": "副本数，默认1",
                    "type": "integer"
                },
                "strategy": {
                    "description": "部署策略",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.DeploymentStrategy"
                        }
                    ]
                },
                "template": {
                    "description": "Pod模板",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PodTemplateSpec"
                        }
                    ]
                }
            }
        },
        "model.CreateIngressRequest": {
            "type": "object",
            "required": [
                "name",
                "rules"
            ],
            "properties": {
                "annotations": {
                    "description": "注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "class": {
                    "description": "Ingress类",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "Ingress名称",
                    "type": "string"
                },
                "rules": {
                    "description": "路由规则",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressRuleSpec"
                    }
                },
                "tls": {
                    "description": "TLS配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressTLSSpec"
                    }
                }
            }
        },
        "model.CreateNamespaceRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "annotations": {
                    "description": "注释",
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
                },
                "name": {
                    "description": "命名空间名称",
                    "type": "string"
                }
            }
        },
        "model.CreatePVCRequest": {
            "type": "object",
            "required": [
                "accessModes",
                "name",
                "resources"
            ],
            "properties": {
                "accessModes": {
                    "description": "访问模式",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "PVC名称",
                    "type": "string"
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVCResourcesSpec"
                        }
                    ]
                },
                "selector": {
                    "description": "标签选择器",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVCSelectorSpec"
                        }
                    ]
                },
                "storageClass": {
                    "description": "存储类名称",
                    "type": "string"
                },
                "volumeMode": {
                    "description": "卷模式",
                    "type": "string"
                }
            }
        },
        "model.CreatePVRequest": {
            "type": "object",
            "required": [
                "accessModes",
                "capacity",
                "name",
                "persistentVolumeSource"
            ],
            "properties": {
                "accessModes": {
                    "description": "访问模式",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "容量",
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
                },
                "mountOptions": {
                    "description": "挂载选项",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "PV名称",
                    "type": "string"
                },
                "nodeAffinity": {
                    "description": "节点亲和性",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVNodeAffinity"
                        }
                    ]
                },
                "persistentVolumeSource": {
                    "description": "存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVSourceSpec"
                        }
                    ]
                },
                "reclaimPolicy": {
                    "description": "回收策略",
                    "type": "string"
                },
                "storageClassName": {
                    "description": "存储类名称",
                    "type": "string"
                },
                "volumeMode": {
                    "description": "卷模式",
                    "type": "string"
                }
            }
        },
        "model.CreatePodFromYAMLRequest": {
            "type": "object",
            "required": [
                "yamlContent"
            ],
            "properties": {
                "dryRun": {
                    "description": "是否只进行校验不实际创建",
                    "type": "boolean"
                },
                "validateOnly": {
                    "description": "是否只校验YAML格式",
                    "type": "boolean"
                },
                "yamlContent": {
                    "description": "YAML内容",
                    "type": "string"
                }
            }
        },
        "model.CreatePodFromYAMLResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "Pod所在的命名空间",
                    "type": "string"
                },
                "parsedObject": {
                    "description": "解析的对象信息",
                    "type": "object",
                    "additionalProperties": true
                },
                "podName": {
                    "description": "创建的Pod名称",
                    "type": "string"
                },
                "success": {
                    "description": "是否创建成功",
                    "type": "boolean"
                },
                "validationResult": {
                    "description": "校验结果（DryRun时返回）",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ValidateYAMLResponse"
                        }
                    ]
                }
            }
        },
        "model.CreateQuickDeploymentRequest": {
            "type": "object",
            "required": [
                "applications",
                "business_dept_id",
                "business_group_id",
                "title"
            ],
            "properties": {
                "applications": {
                    "description": "应用列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.QuickDeploymentAppRequest"
                    }
                },
                "business_dept_id": {
                    "description": "业务部门ID",
                    "type": "integer"
                },
                "business_group_id": {
                    "description": "业务组ID",
                    "type": "integer"
                },
                "description": {
                    "description": "发布描述",
                    "type": "string"
                },
                "title": {
                    "description": "发布标题",
                    "type": "string"
                }
            }
        },
        "model.CreateServiceRequest": {
            "type": "object",
            "required": [
                "name",
                "ports",
                "selector",
                "type"
            ],
            "properties": {
                "externalIPs": {
                    "description": "外部IP列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "服务名称",
                    "type": "string"
                },
                "ports": {
                    "description": "端口配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServicePortSpec"
                    }
                },
                "selector": {
                    "description": "选择器",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "type": {
                    "description": "服务类型",
                    "type": "string"
                }
            }
        },
        "model.DeploymentRevision": {
            "type": "object",
            "properties": {
                "annotations": {
                    "description": "注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "changeReason": {
                    "description": "变更原因",
                    "type": "string"
                },
                "creationTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "images": {
                    "description": "镜像列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "replicasSummary": {
                    "description": "副本统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ReplicasSummary"
                        }
                    ]
                },
                "revision": {
                    "description": "版本号",
                    "type": "integer"
                },
                "status": {
                    "description": "版本状态 (current/historical)",
                    "type": "string"
                }
            }
        },
        "model.DeploymentRevisionDetail": {
            "type": "object",
            "properties": {
                "annotations": {
                    "description": "注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "changeReason": {
                    "description": "变更原因",
                    "type": "string"
                },
                "conditions": {
                    "description": "状态条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.WorkloadCondition"
                    }
                },
                "creationTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "images": {
                    "description": "镜像列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "podTemplate": {
                    "description": "Pod模板",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PodTemplateSpec"
                        }
                    ]
                },
                "replicaSets": {
                    "description": "关联的ReplicaSet信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ReplicaSetInfo"
                    }
                },
                "replicasSummary": {
                    "description": "副本统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ReplicasSummary"
                        }
                    ]
                },
                "revision": {
                    "description": "版本号",
                    "type": "integer"
                },
                "status": {
                    "description": "版本状态 (current/historical)",
                    "type": "string"
                },
                "strategy": {
                    "description": "部署策略",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.DeploymentStrategy"
                        }
                    ]
                }
            }
        },
        "model.DeploymentRolloutHistoryResponse": {
            "type": "object",
            "properties": {
                "currentRevision": {
                    "description": "当前版本号",
                    "type": "integer"
                },
                "deploymentName": {
                    "description": "Deployment名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "revisions": {
                    "description": "版本列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.DeploymentRevision"
                    }
                },
                "totalRevisions": {
                    "description": "总版本数",
                    "type": "integer"
                }
            }
        },
        "model.DeploymentRolloutStatusResponse": {
            "type": "object",
            "properties": {
                "availableReplicas": {
                    "description": "可用副本数",
                    "type": "integer"
                },
                "conditions": {
                    "description": "状态条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.WorkloadCondition"
                    }
                },
                "currentRevision": {
                    "description": "当前版本号",
                    "type": "integer"
                },
                "deploymentName": {
                    "description": "Deployment名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "observedGeneration": {
                    "description": "观察到的代数",
                    "type": "integer"
                },
                "paused": {
                    "description": "是否已暂停",
                    "type": "boolean"
                },
                "progressDeadline": {
                    "description": "进度截止时间",
                    "type": "integer"
                },
                "readyReplicas": {
                    "description": "就绪副本数",
                    "type": "integer"
                },
                "rolloutComplete": {
                    "description": "是否滚动发布完成",
                    "type": "boolean"
                },
                "status": {
                    "description": "总体状态 (Progressing/Complete/Failed/Paused)",
                    "type": "string"
                },
                "strategy": {
                    "description": "部署策略",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.DeploymentStrategy"
                        }
                    ]
                },
                "updatedReplicas": {
                    "description": "已更新副本数",
                    "type": "integer"
                }
            }
        },
        "model.DeploymentStats": {
            "type": "object",
            "properties": {
                "failed": {
                    "description": "失败次数",
                    "type": "integer"
                },
                "success": {
                    "description": "成功次数",
                    "type": "integer"
                },
                "successRate": {
                    "description": "成功率",
                    "type": "number"
                },
                "total": {
                    "description": "发布总次数",
                    "type": "integer"
                }
            }
        },
        "model.DeploymentStrategy": {
            "type": "object",
            "properties": {
                "rollingUpdate": {
                    "description": "滚动更新配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RollingUpdateDeployment"
                        }
                    ]
                },
                "type": {
                    "description": "策略类型",
                    "type": "string"
                }
            }
        },
        "model.DrainNodeRequest": {
            "type": "object",
            "properties": {
                "deleteLocalData": {
                    "description": "删除本地数据",
                    "type": "boolean"
                },
                "force": {
                    "description": "强制驱逐",
                    "type": "boolean"
                },
                "gracePeriodSeconds": {
                    "description": "优雅终止时间",
                    "type": "integer"
                },
                "ignoreDaemonSets": {
                    "description": "忽略DaemonSet",
                    "type": "boolean"
                }
            }
        },
        "model.ExecuteQuickDeploymentRequest": {
            "type": "object",
            "required": [
                "deployment_id"
            ],
            "properties": {
                "deployment_id": {
                    "description": "发布ID",
                    "type": "integer"
                },
                "execution_mode": {
                    "description": "执行模式: 1=并行(默认) 2=串行",
                    "type": "integer",
                    "default": 1,
                    "enum": [
                        1,
                        2
                    ]
                }
            }
        },
        "model.IngressBackend": {
            "type": "object",
            "properties": {
                "service": {
                    "description": "服务后端",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.IngressServiceBackend"
                        }
                    ]
                }
            }
        },
        "model.IngressDetail": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Ingress类/控制器类型",
                    "type": "string"
                },
                "controllerName": {
                    "description": "Controller名称",
                    "type": "string"
                },
                "controllerVersion": {
                    "description": "Controller版本",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "endpoints": {
                    "description": "访问端点",
                    "type": "array",
                    "items": {
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
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "loadBalancer": {
                    "description": "负载均衡器信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.IngressLoadBalancer"
                        }
                    ]
                },
                "name": {
                    "description": "Ingress名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "rules": {
                    "description": "路由规则",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressRule"
                    }
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "tls": {
                    "description": "TLS配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressTLS"
                    }
                },
                "type": {
                    "description": "Ingress类型：公网Nginx/内网Nginx等",
                    "type": "string"
                }
            }
        },
        "model.IngressListResponse": {
            "type": "object",
            "properties": {
                "ingresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sIngress"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.IngressLoadBalancer": {
            "type": "object",
            "properties": {
                "ingress": {
                    "description": "入口信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressLoadBalancerIngress"
                    }
                }
            }
        },
        "model.IngressLoadBalancerIngress": {
            "type": "object",
            "properties": {
                "hostname": {
                    "description": "主机名",
                    "type": "string"
                },
                "ip": {
                    "description": "IP地址",
                    "type": "string"
                },
                "ports": {
                    "description": "端口信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressLoadBalancerPort"
                    }
                }
            }
        },
        "model.IngressLoadBalancerPort": {
            "type": "object",
            "properties": {
                "port": {
                    "description": "端口号",
                    "type": "integer"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                }
            }
        },
        "model.IngressPath": {
            "type": "object",
            "properties": {
                "backend": {
                    "description": "后端服务",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.IngressBackend"
                        }
                    ]
                },
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "pathType": {
                    "description": "路径类型 Exact/Prefix/ImplementationSpecific",
                    "type": "string"
                }
            }
        },
        "model.IngressPathSpec": {
            "type": "object",
            "required": [
                "path",
                "pathType",
                "serviceName",
                "servicePort"
            ],
            "properties": {
                "path": {
                    "description": "路径",
                    "type": "string"
                },
                "pathType": {
                    "description": "路径类型",
                    "type": "string"
                },
                "serviceName": {
                    "description": "服务名称",
                    "type": "string"
                },
                "servicePort": {
                    "description": "服务端口",
                    "type": "integer"
                }
            }
        },
        "model.IngressRule": {
            "type": "object",
            "properties": {
                "host": {
                    "description": "主机名",
                    "type": "string"
                },
                "http": {
                    "description": "HTTP规则",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.IngressRuleValue"
                        }
                    ]
                }
            }
        },
        "model.IngressRuleSpec": {
            "type": "object",
            "required": [
                "host",
                "paths"
            ],
            "properties": {
                "host": {
                    "description": "主机名",
                    "type": "string"
                },
                "paths": {
                    "description": "路径规则",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressPathSpec"
                    }
                }
            }
        },
        "model.IngressRuleValue": {
            "type": "object",
            "properties": {
                "paths": {
                    "description": "路径规则",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressPath"
                    }
                }
            }
        },
        "model.IngressServiceBackend": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "服务名称",
                    "type": "string"
                },
                "port": {
                    "description": "服务端口",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.IngressServicePort"
                        }
                    ]
                }
            }
        },
        "model.IngressServicePort": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "端口名称",
                    "type": "string"
                },
                "number": {
                    "description": "端口号",
                    "type": "integer"
                }
            }
        },
        "model.IngressTLS": {
            "type": "object",
            "properties": {
                "hosts": {
                    "description": "主机列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "secretName": {
                    "description": "证书Secret名称",
                    "type": "string"
                }
            }
        },
        "model.IngressTLSSpec": {
            "type": "object",
            "required": [
                "hosts",
                "secretName"
            ],
            "properties": {
                "hosts": {
                    "description": "主机列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "secretName": {
                    "description": "证书Secret名称",
                    "type": "string"
                }
            }
        },
        "model.K8sClusterStats": {
            "type": "object",
            "properties": {
                "healthy": {
                    "description": "健康数量",
                    "type": "integer"
                },
                "offline": {
                    "description": "离线数量",
                    "type": "integer"
                },
                "total": {
                    "description": "集群总数",
                    "type": "integer"
                }
            }
        },
        "model.K8sConfigMap": {
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
                }
            }
        },
        "model.K8sEvent": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "发生次数",
                    "type": "integer"
                },
                "firstTime": {
                    "description": "首次时间",
                    "type": "string"
                },
                "lastTime": {
                    "description": "最后时间",
                    "type": "string"
                },
                "message": {
                    "description": "消息",
                    "type": "string"
                },
                "reason": {
                    "description": "原因",
                    "type": "string"
                },
                "source": {
                    "description": "事件源",
                    "type": "string"
                },
                "type": {
                    "description": "事件类型",
                    "type": "string"
                }
            }
        },
        "model.K8sIngress": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Ingress类/控制器类型",
                    "type": "string"
                },
                "controllerName": {
                    "description": "Controller名称",
                    "type": "string"
                },
                "controllerVersion": {
                    "description": "Controller版本",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "endpoints": {
                    "description": "访问端点",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "loadBalancer": {
                    "description": "负载均衡器信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.IngressLoadBalancer"
                        }
                    ]
                },
                "name": {
                    "description": "Ingress名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "rules": {
                    "description": "路由规则",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressRule"
                    }
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "tls": {
                    "description": "TLS配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressTLS"
                    }
                },
                "type": {
                    "description": "Ingress类型：公网Nginx/内网Nginx等",
                    "type": "string"
                }
            }
        },
        "model.K8sNamespace": {
            "type": "object",
            "properties": {
                "annotations": {
                    "description": "注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "limitRanges": {
                    "description": "默认资源限制列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LimitRangeDetail"
                    }
                },
                "name": {
                    "description": "命名空间名称",
                    "type": "string"
                },
                "resourceCount": {
                    "description": "资源统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NamespaceResourceCount"
                        }
                    ]
                },
                "resourceQuotas": {
                    "description": "资源配额列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ResourceQuotaDetail"
                    }
                },
                "status": {
                    "description": "状态 Active/Terminating",
                    "type": "string"
                }
            }
        },
        "model.K8sNode": {
            "type": "object",
            "properties": {
                "conditions": {
                    "description": "节点状态详细条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeCondition"
                    }
                },
                "configuration": {
                    "description": "节点配置信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeConfiguration"
                        }
                    ]
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "externalIP": {
                    "description": "外部IP地址",
                    "type": "string"
                },
                "internalIP": {
                    "description": "内部IP地址",
                    "type": "string"
                },
                "name": {
                    "description": "节点名称",
                    "type": "string"
                },
                "podMetrics": {
                    "description": "容器组统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PodMetrics"
                        }
                    ]
                },
                "resources": {
                    "description": "CPU和内存资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeResources"
                        }
                    ]
                },
                "roles": {
                    "description": "节点角色 control-plane,master 或 worker",
                    "type": "string"
                },
                "runtime": {
                    "description": "运行时和版本信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RuntimeInfo"
                        }
                    ]
                },
                "scheduling": {
                    "description": "调度相关信息（污点、是否可调度等）",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeSchedulingInfo"
                        }
                    ]
                },
                "status": {
                    "description": "节点状态 Ready/NotReady",
                    "type": "string"
                }
            }
        },
        "model.K8sNodeDetail": {
            "type": "object",
            "properties": {
                "conditions": {
                    "description": "节点状态详细条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeCondition"
                    }
                },
                "configuration": {
                    "description": "节点配置信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeConfiguration"
                        }
                    ]
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EventInfo"
                    }
                },
                "externalIP": {
                    "description": "外部IP地址",
                    "type": "string"
                },
                "internalIP": {
                    "description": "内部IP地址",
                    "type": "string"
                },
                "metrics": {
                    "description": "详细监控指标",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeMetrics"
                        }
                    ]
                },
                "name": {
                    "description": "节点名称",
                    "type": "string"
                },
                "podMetrics": {
                    "description": "容器组统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PodMetrics"
                        }
                    ]
                },
                "pods": {
                    "description": "节点上运行的Pod列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PodInfo"
                    }
                },
                "resources": {
                    "description": "CPU和内存资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeResources"
                        }
                    ]
                },
                "roles": {
                    "description": "节点角色 control-plane,master 或 worker",
                    "type": "string"
                },
                "runtime": {
                    "description": "运行时和版本信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RuntimeInfo"
                        }
                    ]
                },
                "scheduling": {
                    "description": "调度相关信息（污点、是否可调度等）",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeSchedulingInfo"
                        }
                    ]
                },
                "status": {
                    "description": "节点状态 Ready/NotReady",
                    "type": "string"
                }
            }
        },
        "model.K8sPersistentVolume": {
            "type": "object",
            "properties": {
                "accessModes": {
                    "description": "访问模式",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "总量",
                    "type": "string"
                },
                "claimRef": {
                    "description": "绑定存储声明",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVClaimRef"
                        }
                    ]
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "mountOptions": {
                    "description": "挂载选项",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "PV名称",
                    "type": "string"
                },
                "nodeAffinity": {
                    "description": "节点亲和性",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVNodeAffinity"
                        }
                    ]
                },
                "persistentVolumeSource": {
                    "description": "存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVSource"
                        }
                    ]
                },
                "reclaimPolicy": {
                    "description": "回收策略 Retain/Delete/Recycle",
                    "type": "string"
                },
                "status": {
                    "description": "状态 Available/Bound/Released/Failed",
                    "type": "string"
                },
                "storageClass": {
                    "description": "存储类型",
                    "type": "string"
                },
                "volumeMode": {
                    "description": "卷模式",
                    "type": "string"
                }
            }
        },
        "model.K8sPersistentVolumeClaim": {
            "type": "object",
            "properties": {
                "accessModes": {
                    "description": "访问模式",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "总量",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "PVC名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "status": {
                    "description": "状态 Pending/Bound/Lost",
                    "type": "string"
                },
                "storageClass": {
                    "description": "存储类型",
                    "type": "string"
                },
                "volumeMode": {
                    "description": "卷模式 Filesystem/Block",
                    "type": "string"
                },
                "volumeName": {
                    "description": "关联的存储卷",
                    "type": "string"
                }
            }
        },
        "model.K8sPodDetail": {
            "type": "object",
            "properties": {
                "conditions": {
                    "description": "Pod状态条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PodCondition"
                    }
                },
                "containers": {
                    "description": "容器信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerInfo"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "hostIP": {
                    "description": "主机IP",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "实例名称",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "phase": {
                    "description": "阶段",
                    "type": "string"
                },
                "podIP": {
                    "description": "Pod IP",
                    "type": "string"
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadResources"
                        }
                    ]
                },
                "restartCount": {
                    "description": "重启次数",
                    "type": "integer"
                },
                "runningTime": {
                    "description": "运行时间",
                    "type": "string"
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "volumes": {
                    "description": "挂载卷信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VolumeInfo"
                    }
                }
            }
        },
        "model.K8sPodInfo": {
            "type": "object",
            "properties": {
                "containers": {
                    "description": "容器信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerInfo"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "hostIP": {
                    "description": "主机IP",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "实例名称",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "phase": {
                    "description": "阶段",
                    "type": "string"
                },
                "podIP": {
                    "description": "Pod IP",
                    "type": "string"
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadResources"
                        }
                    ]
                },
                "restartCount": {
                    "description": "重启次数",
                    "type": "integer"
                },
                "runningTime": {
                    "description": "运行时间",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                }
            }
        },
        "model.K8sSecret": {
            "type": "object",
            "properties": {
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "data": {
                    "description": "数据(base64编码)",
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer",
                            "format": "int32"
                        }
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
                    "description": "Secret名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "stringData": {
                    "description": "字符串数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "type": {
                    "description": "Secret类型",
                    "type": "string"
                }
            }
        },
        "model.K8sService": {
            "type": "object",
            "properties": {
                "clusterIP": {
                    "description": "集群IP",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "endpoints": {
                    "description": "端点信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServiceEndpoint"
                    }
                },
                "externalIPs": {
                    "description": "外部IP列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "服务名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "ports": {
                    "description": "端口配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServicePort"
                    }
                },
                "selector": {
                    "description": "选择器",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "服务类型 ClusterIP/NodePort/LoadBalancer/ExternalName",
                    "type": "string"
                }
            }
        },
        "model.K8sStorageClass": {
            "type": "object",
            "properties": {
                "allowVolumeExpansion": {
                    "description": "允许卷扩展",
                    "type": "boolean"
                },
                "allowedTopologies": {
                    "description": "允许的拓扑",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StorageClassTopology"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "mountOptions": {
                    "description": "挂载选项",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "存储类名称",
                    "type": "string"
                },
                "parameters": {
                    "description": "参数",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "provisioner": {
                    "description": "提供者",
                    "type": "string"
                },
                "reclaimPolicy": {
                    "description": "回收策略",
                    "type": "string"
                },
                "volumeBindingMode": {
                    "description": "卷绑定模式",
                    "type": "string"
                }
            }
        },
        "model.K8sWorkload": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "images": {
                    "description": "镜像列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "readyReplicas": {
                    "description": "就绪副本数",
                    "type": "integer"
                },
                "replicas": {
                    "description": "副本数",
                    "type": "integer"
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadResources"
                        }
                    ]
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "类型",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadType"
                        }
                    ]
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "model.K8sWorkloadDetail": {
            "type": "object",
            "properties": {
                "conditions": {
                    "description": "状态条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.WorkloadCondition"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "images": {
                    "description": "镜像列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "pods": {
                    "description": "Pod列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sPodInfo"
                    }
                },
                "readyReplicas": {
                    "description": "就绪副本数",
                    "type": "integer"
                },
                "replicas": {
                    "description": "副本数",
                    "type": "integer"
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadResources"
                        }
                    ]
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "类型",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadType"
                        }
                    ]
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "model.NamespaceListResponse": {
            "type": "object",
            "properties": {
                "namespaces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sNamespace"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.NamespaceMetricsInfo": {
            "type": "object",
            "properties": {
                "namespace": {
                    "description": "命名空间名称",
                    "type": "string"
                },
                "podCount": {
                    "description": "Pod总数",
                    "type": "integer"
                },
                "podMetrics": {
                    "description": "Pod监控列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PodMetricsSummary"
                    }
                },
                "resourceQuota": {
                    "description": "命名空间资源配额",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NamespaceResourceQuota"
                        }
                    ]
                },
                "runningPods": {
                    "description": "运行中的Pod数",
                    "type": "integer"
                },
                "timestamp": {
                    "description": "采集时间",
                    "type": "string"
                },
                "totalUsage": {
                    "description": "总资源使用量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "usageRate": {
                    "description": "资源使用率",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsageRate"
                        }
                    ]
                }
            }
        },
        "model.NamespaceResourceCount": {
            "type": "object",
            "properties": {
                "configMapCount": {
                    "description": "ConfigMap数量",
                    "type": "integer"
                },
                "podCount": {
                    "description": "Pod数量",
                    "type": "integer"
                },
                "secretCount": {
                    "description": "Secret数量",
                    "type": "integer"
                },
                "serviceCount": {
                    "description": "Service数量",
                    "type": "integer"
                }
            }
        },
        "model.NamespaceResourceQuota": {
            "type": "object",
            "properties": {
                "hard": {
                    "description": "硬限制",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "used": {
                    "description": "已使用",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                }
            }
        },
        "model.NodeCondition": {
            "type": "object",
            "properties": {
                "reason": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.NodeConfig": {
            "type": "object",
            "required": [
                "etcdHostIds",
                "masterHostIds"
            ],
            "properties": {
                "etcdHostIds": {
                    "description": "ETCD节点主机ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "masterHostIds": {
                    "description": "Master节点主机ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "workerHostIds": {
                    "description": "Worker节点主机ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.NodeConfiguration": {
            "type": "object",
            "properties": {
                "annotations": {
                    "description": "节点注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "architecture": {
                    "description": "系统架构",
                    "type": "string"
                },
                "kernelVersion": {
                    "description": "内核版本",
                    "type": "string"
                },
                "labels": {
                    "description": "节点标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "osImage": {
                    "description": "操作系统镜像",
                    "type": "string"
                },
                "role": {
                    "description": "节点角色 master/worker",
                    "type": "string"
                }
            }
        },
        "model.NodeDetailResponse": {
            "type": "object",
            "properties": {
                "allocatable": {
                    "description": "可分配资源",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "annotations": {
                    "description": "注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "architecture": {
                    "description": "架构",
                    "type": "string"
                },
                "bootID": {
                    "description": "启动ID",
                    "type": "string"
                },
                "capacity": {
                    "description": "资源信息",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "conditions": {
                    "description": "状态条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeCondition"
                    }
                },
                "containerRuntimeVersion": {
                    "description": "容器运行时版本",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "externalIP": {
                    "description": "外部IP",
                    "type": "string"
                },
                "hostname": {
                    "description": "主机名",
                    "type": "string"
                },
                "internalIP": {
                    "description": "IP地址信息",
                    "type": "string"
                },
                "kernelVersion": {
                    "description": "内核版本",
                    "type": "string"
                },
                "kubeProxyVersion": {
                    "description": "Kube-Proxy版本",
                    "type": "string"
                },
                "kubeletVersion": {
                    "description": "K8s组件版本",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "machineID": {
                    "description": "机器ID",
                    "type": "string"
                },
                "monitoring": {
                    "description": "监控信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeMonitoringInfo"
                        }
                    ]
                },
                "name": {
                    "description": "基本信息",
                    "type": "string"
                },
                "operatingSystem": {
                    "description": "操作系统",
                    "type": "string"
                },
                "osImage": {
                    "description": "系统信息",
                    "type": "string"
                },
                "podCIDR": {
                    "description": "CIDR信息",
                    "type": "string"
                },
                "podCIDRs": {
                    "description": "容器组CIDR列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "podInfo": {
                    "description": "Pod信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodePodInfo"
                        }
                    ]
                },
                "podList": {
                    "description": "节点上的Pod列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodePodDetail"
                    }
                },
                "providerID": {
                    "description": "提供者ID",
                    "type": "string"
                },
                "status": {
                    "description": "状态信息",
                    "type": "string"
                },
                "systemUUID": {
                    "description": "系统UUID",
                    "type": "string"
                },
                "taints": {
                    "description": "污点列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeTaint"
                    }
                },
                "uid": {
                    "description": "UID",
                    "type": "string"
                },
                "unschedulable": {
                    "description": "调度信息",
                    "type": "boolean"
                }
            }
        },
        "model.NodeInfo": {
            "type": "object",
            "properties": {
                "allocatable": {
                    "description": "可分配资源",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "CPU、内存等资源容量",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "conditions": {
                    "description": "节点状态条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeCondition"
                    }
                },
                "externalIP": {
                    "type": "string"
                },
                "internalIP": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "role": {
                    "description": "master/worker/etcd",
                    "type": "string"
                },
                "status": {
                    "description": "Ready/NotReady",
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "model.NodeMetrics": {
            "type": "object",
            "properties": {
                "cpuUsagePercentage": {
                    "type": "number"
                },
                "diskUsagePercentage": {
                    "type": "number"
                },
                "memoryUsagePercentage": {
                    "type": "number"
                },
                "networkInBytes": {
                    "type": "integer"
                },
                "networkOutBytes": {
                    "type": "integer"
                }
            }
        },
        "model.NodeMetricsInfo": {
            "type": "object",
            "properties": {
                "allocatable": {
                    "description": "可分配量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "capacity": {
                    "description": "总容量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "podCount": {
                    "description": "Pod数量",
                    "type": "integer"
                },
                "podMetrics": {
                    "description": "Pod监控摘要",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PodMetricsSummary"
                    }
                },
                "systemInfo": {
                    "description": "系统信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeSystemInfo"
                        }
                    ]
                },
                "timestamp": {
                    "description": "采集时间",
                    "type": "string"
                },
                "usage": {
                    "description": "资源使用量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "usageRate": {
                    "description": "使用率",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsageRate"
                        }
                    ]
                }
            }
        },
        "model.NodeMonitoringInfo": {
            "type": "object",
            "properties": {
                "cpu": {
                    "description": "CPU使用情况",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeResourceUsage"
                        }
                    ]
                },
                "memory": {
                    "description": "内存使用情况",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeResourceUsage"
                        }
                    ]
                },
                "network": {
                    "description": "网络使用情况",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeNetworkUsage"
                        }
                    ]
                },
                "storage": {
                    "description": "存储使用情况",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeResourceUsage"
                        }
                    ]
                }
            }
        },
        "model.NodeNetworkUsage": {
            "type": "object",
            "properties": {
                "inboundBytes": {
                    "description": "入站字节数",
                    "type": "integer"
                },
                "inboundPackets": {
                    "description": "入站包数",
                    "type": "integer"
                },
                "outboundBytes": {
                    "description": "出站字节数",
                    "type": "integer"
                },
                "outboundPackets": {
                    "description": "出站包数",
                    "type": "integer"
                }
            }
        },
        "model.NodePodDetail": {
            "type": "object",
            "properties": {
                "containers": {
                    "description": "容器状态",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerStatus"
                    }
                },
                "cpuLimits": {
                    "description": "CPU限制",
                    "type": "string"
                },
                "cpuRequests": {
                    "description": "CPU请求",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "memoryLimits": {
                    "description": "内存限制",
                    "type": "string"
                },
                "memoryRequests": {
                    "description": "内存请求",
                    "type": "string"
                },
                "name": {
                    "description": "Pod名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "phase": {
                    "description": "阶段",
                    "type": "string"
                },
                "restartCount": {
                    "description": "重启次数",
                    "type": "integer"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                }
            }
        },
        "model.NodePodInfo": {
            "type": "object",
            "properties": {
                "failedPods": {
                    "description": "失败的Pod数",
                    "type": "integer"
                },
                "pendingPods": {
                    "description": "等待中的Pod数",
                    "type": "integer"
                },
                "runningPods": {
                    "description": "运行中的Pod数",
                    "type": "integer"
                },
                "succeededPods": {
                    "description": "成功的Pod数",
                    "type": "integer"
                },
                "totalPods": {
                    "description": "Pod总数",
                    "type": "integer"
                }
            }
        },
        "model.NodeResourceAllocation": {
            "type": "object",
            "properties": {
                "allocatable": {
                    "description": "可分配资源",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "allocated": {
                    "description": "已分配资源",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "节点总容量",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "nodeName": {
                    "type": "string"
                },
                "podList": {
                    "description": "Pod资源使用详情",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PodResourceInfo"
                    }
                }
            }
        },
        "model.NodeResourceUsage": {
            "type": "object",
            "properties": {
                "available": {
                    "description": "可用量",
                    "type": "string"
                },
                "limits": {
                    "description": "限制量",
                    "type": "string"
                },
                "requestRate": {
                    "description": "请求率 (0-100)",
                    "type": "number"
                },
                "requests": {
                    "description": "请求量",
                    "type": "string"
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
        "model.NodeResources": {
            "type": "object",
            "properties": {
                "cpu": {
                    "description": "CPU资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceInfo"
                        }
                    ]
                },
                "memory": {
                    "description": "内存资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceInfo"
                        }
                    ]
                }
            }
        },
        "model.NodeSchedulingInfo": {
            "type": "object",
            "properties": {
                "taints": {
                    "description": "节点污点",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeTaint"
                    }
                },
                "unschedulable": {
                    "description": "是否不可调度",
                    "type": "boolean"
                }
            }
        },
        "model.NodeSystemInfo": {
            "type": "object",
            "properties": {
                "architecture": {
                    "description": "系统架构",
                    "type": "string"
                },
                "containerRuntimeVersion": {
                    "description": "容器运行时版本",
                    "type": "string"
                },
                "kernelVersion": {
                    "description": "内核版本",
                    "type": "string"
                },
                "kubeProxyVersion": {
                    "description": "KubeProxy版本",
                    "type": "string"
                },
                "kubeletVersion": {
                    "description": "Kubelet版本",
                    "type": "string"
                },
                "osImage": {
                    "description": "操作系统镜像",
                    "type": "string"
                }
            }
        },
        "model.NodeTaint": {
            "type": "object",
            "properties": {
                "effect": {
                    "description": "NoSchedule, PreferNoSchedule, NoExecute",
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.PVAWSElasticBlockStoreVolumeSource": {
            "type": "object",
            "properties": {
                "fsType": {
                    "description": "文件系统类型",
                    "type": "string"
                },
                "partition": {
                    "description": "分区",
                    "type": "integer"
                },
                "readOnly": {
                    "description": "只读",
                    "type": "boolean"
                },
                "volumeID": {
                    "description": "卷ID",
                    "type": "string"
                }
            }
        },
        "model.PVCDetail": {
            "type": "object",
            "properties": {
                "accessModes": {
                    "description": "访问模式",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "总量",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "PVC名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "status": {
                    "description": "状态 Pending/Bound/Lost",
                    "type": "string"
                },
                "storageClass": {
                    "description": "存储类型",
                    "type": "string"
                },
                "volumeMode": {
                    "description": "卷模式 Filesystem/Block",
                    "type": "string"
                },
                "volumeName": {
                    "description": "关联的存储卷",
                    "type": "string"
                }
            }
        },
        "model.PVCListResponse": {
            "type": "object",
            "properties": {
                "pvcs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sPersistentVolumeClaim"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.PVCMatchExp": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "键",
                    "type": "string"
                },
                "operator": {
                    "description": "操作符",
                    "type": "string"
                },
                "values": {
                    "description": "值",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.PVCResourcesSpec": {
            "type": "object",
            "required": [
                "requests"
            ],
            "properties": {
                "limits": {
                    "description": "资源限制",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "requests": {
                    "description": "资源请求",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.PVCSIVolumeSource": {
            "type": "object",
            "properties": {
                "driver": {
                    "description": "驱动名称",
                    "type": "string"
                },
                "fsType": {
                    "description": "文件系统类型",
                    "type": "string"
                },
                "readOnly": {
                    "description": "只读",
                    "type": "boolean"
                },
                "volumeAttributes": {
                    "description": "卷属性",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "volumeHandle": {
                    "description": "卷句柄",
                    "type": "string"
                }
            }
        },
        "model.PVCSelectorSpec": {
            "type": "object",
            "properties": {
                "matchExpressions": {
                    "description": "匹配表达式",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PVCMatchExp"
                    }
                },
                "matchLabels": {
                    "description": "匹配标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.PVClaimRef": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "description": "API版本",
                    "type": "string"
                },
                "kind": {
                    "description": "类型",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "uid": {
                    "description": "UID",
                    "type": "string"
                }
            }
        },
        "model.PVDetail": {
            "type": "object",
            "properties": {
                "accessModes": {
                    "description": "访问模式",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "capacity": {
                    "description": "总量",
                    "type": "string"
                },
                "claimRef": {
                    "description": "绑定存储声明",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVClaimRef"
                        }
                    ]
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "mountOptions": {
                    "description": "挂载选项",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "PV名称",
                    "type": "string"
                },
                "nodeAffinity": {
                    "description": "节点亲和性",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVNodeAffinity"
                        }
                    ]
                },
                "persistentVolumeSource": {
                    "description": "存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVSource"
                        }
                    ]
                },
                "reclaimPolicy": {
                    "description": "回收策略 Retain/Delete/Recycle",
                    "type": "string"
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "status": {
                    "description": "状态 Available/Bound/Released/Failed",
                    "type": "string"
                },
                "storageClass": {
                    "description": "存储类型",
                    "type": "string"
                },
                "volumeMode": {
                    "description": "卷模式",
                    "type": "string"
                }
            }
        },
        "model.PVHostPathVolumeSource": {
            "type": "object",
            "properties": {
                "path": {
                    "description": "主机路径",
                    "type": "string"
                },
                "type": {
                    "description": "类型",
                    "type": "string"
                }
            }
        },
        "model.PVISCSIVolumeSource": {
            "type": "object",
            "properties": {
                "fsType": {
                    "description": "文件系统类型",
                    "type": "string"
                },
                "iqn": {
                    "description": "IQN",
                    "type": "string"
                },
                "iscsiInterface": {
                    "description": "iSCSI接口",
                    "type": "string"
                },
                "lun": {
                    "description": "LUN",
                    "type": "integer"
                },
                "portals": {
                    "description": "门户列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "readOnly": {
                    "description": "只读",
                    "type": "boolean"
                },
                "targetPortal": {
                    "description": "目标门户",
                    "type": "string"
                }
            }
        },
        "model.PVListResponse": {
            "type": "object",
            "properties": {
                "pvs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sPersistentVolume"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.PVLocalVolumeSource": {
            "type": "object",
            "properties": {
                "fsType": {
                    "description": "文件系统类型",
                    "type": "string"
                },
                "path": {
                    "description": "本地路径",
                    "type": "string"
                }
            }
        },
        "model.PVNFSVolumeSource": {
            "type": "object",
            "properties": {
                "path": {
                    "description": "NFS路径",
                    "type": "string"
                },
                "readOnly": {
                    "description": "只读",
                    "type": "boolean"
                },
                "server": {
                    "description": "NFS服务器",
                    "type": "string"
                }
            }
        },
        "model.PVNodeAffinity": {
            "type": "object",
            "properties": {
                "required": {
                    "description": "必需的节点选择器",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVNodeSelector"
                        }
                    ]
                }
            }
        },
        "model.PVNodeSelector": {
            "type": "object",
            "properties": {
                "nodeSelectorTerms": {
                    "description": "节点选择器条件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PVNodeSelectorTerm"
                    }
                }
            }
        },
        "model.PVNodeSelectorRequirement": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "键",
                    "type": "string"
                },
                "operator": {
                    "description": "操作符",
                    "type": "string"
                },
                "values": {
                    "description": "值",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.PVNodeSelectorTerm": {
            "type": "object",
            "properties": {
                "matchExpressions": {
                    "description": "匹配表达式",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PVNodeSelectorRequirement"
                    }
                },
                "matchFields": {
                    "description": "匹配字段",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.PVNodeSelectorRequirement"
                    }
                }
            }
        },
        "model.PVSource": {
            "type": "object",
            "properties": {
                "awsElasticBlockStore": {
                    "description": "AWS EBS",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVAWSElasticBlockStoreVolumeSource"
                        }
                    ]
                },
                "csi": {
                    "description": "CSI存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVCSIVolumeSource"
                        }
                    ]
                },
                "hostPath": {
                    "description": "HostPath存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVHostPathVolumeSource"
                        }
                    ]
                },
                "iscsi": {
                    "description": "iSCSI存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVISCSIVolumeSource"
                        }
                    ]
                },
                "local": {
                    "description": "Local存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVLocalVolumeSource"
                        }
                    ]
                },
                "nfs": {
                    "description": "NFS存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVNFSVolumeSource"
                        }
                    ]
                }
            }
        },
        "model.PVSourceSpec": {
            "type": "object",
            "properties": {
                "awsElasticBlockStore": {
                    "description": "AWS EBS",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVAWSElasticBlockStoreVolumeSource"
                        }
                    ]
                },
                "csi": {
                    "description": "CSI存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVCSIVolumeSource"
                        }
                    ]
                },
                "hostPath": {
                    "description": "HostPath存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVHostPathVolumeSource"
                        }
                    ]
                },
                "iscsi": {
                    "description": "iSCSI存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVISCSIVolumeSource"
                        }
                    ]
                },
                "local": {
                    "description": "Local存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVLocalVolumeSource"
                        }
                    ]
                },
                "nfs": {
                    "description": "NFS存储源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVNFSVolumeSource"
                        }
                    ]
                }
            }
        },
        "model.PauseDeploymentResponse": {
            "type": "object",
            "properties": {
                "deploymentName": {
                    "description": "Deployment名称",
                    "type": "string"
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "status": {
                    "description": "当前状态",
                    "type": "string"
                },
                "success": {
                    "description": "是否暂停成功",
                    "type": "boolean"
                }
            }
        },
        "model.PodCondition": {
            "type": "object",
            "properties": {
                "lastTransitionTime": {
                    "description": "最后转换时间",
                    "type": "string"
                },
                "message": {
                    "description": "消息",
                    "type": "string"
                },
                "reason": {
                    "description": "原因",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "条件类型",
                    "type": "string"
                }
            }
        },
        "model.PodInfo": {
            "type": "object",
            "properties": {
                "cpuUsage": {
                    "type": "string"
                },
                "memUsage": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.PodMetrics": {
            "type": "object",
            "properties": {
                "allocated": {
                    "description": "已分配的Pod数量",
                    "type": "integer"
                },
                "total": {
                    "description": "总的Pod容量",
                    "type": "integer"
                }
            }
        },
        "model.PodMetricsInfo": {
            "type": "object",
            "properties": {
                "containers": {
                    "description": "容器监控信息列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerMetricsInfo"
                    }
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "podName": {
                    "description": "Pod名称",
                    "type": "string"
                },
                "resourceQuota": {
                    "description": "资源配额信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PodResourceQuota"
                        }
                    ]
                },
                "timestamp": {
                    "description": "采集时间",
                    "type": "string"
                },
                "totalUsage": {
                    "description": "总使用量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "usageRate": {
                    "description": "使用率信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsageRate"
                        }
                    ]
                }
            }
        },
        "model.PodMetricsSummary": {
            "type": "object",
            "properties": {
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "podName": {
                    "description": "Pod名称",
                    "type": "string"
                },
                "usage": {
                    "description": "资源使用量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "usageRate": {
                    "description": "使用率",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsageRate"
                        }
                    ]
                }
            }
        },
        "model.PodResourceInfo": {
            "type": "object",
            "properties": {
                "limits": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "requests": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.PodResourceQuota": {
            "type": "object",
            "properties": {
                "limits": {
                    "description": "资源限制量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                },
                "requests": {
                    "description": "资源请求量",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ResourceUsage"
                        }
                    ]
                }
            }
        },
        "model.PodTemplateSpec": {
            "type": "object",
            "required": [
                "containers"
            ],
            "properties": {
                "containers": {
                    "description": "容器规格",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ContainerSpec"
                    }
                },
                "labels": {
                    "description": "Pod标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "nodeSelector": {
                    "description": "节点选择器",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "tolerations": {
                    "description": "容忍度",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Toleration"
                    }
                },
                "volumes": {
                    "description": "存储卷规格",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VolumeSpec"
                    }
                }
            }
        },
        "model.QuickDeployment": {
            "type": "object",
            "properties": {
                "business_dept_id": {
                    "description": "业务部门ID",
                    "type": "integer"
                },
                "business_group_id": {
                    "description": "业务组ID",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "creator_id": {
                    "description": "创建人ID",
                    "type": "integer"
                },
                "creator_name": {
                    "description": "创建人姓名",
                    "type": "string"
                },
                "description": {
                    "description": "发布描述",
                    "type": "string"
                },
                "duration": {
                    "description": "发布耗时(秒)",
                    "type": "integer"
                },
                "end_time": {
                    "description": "结束发布时间",
                    "type": "string"
                },
                "execution_mode": {
                    "description": "执行模式: 1=并行 2=串行",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "start_time": {
                    "description": "开始发布时间",
                    "type": "string"
                },
                "status": {
                    "description": "发布状态: 1=待发布 2=发布中 3=发布成功 4=发布失败 5=已取消",
                    "type": "integer"
                },
                "task_count": {
                    "description": "任务数量，记录用户提交的发布任务数量",
                    "type": "integer"
                },
                "tasks": {
                    "description": "关联的发布任务",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.QuickDeploymentTask"
                    }
                },
                "title": {
                    "description": "发布标题",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.QuickDeploymentAppRequest": {
            "type": "object",
            "required": [
                "app_id",
                "environment"
            ],
            "properties": {
                "app_id": {
                    "description": "应用ID（按数组顺序执行）",
                    "type": "integer"
                },
                "environment": {
                    "description": "应用发布环境",
                    "type": "string"
                }
            }
        },
        "model.QuickDeploymentListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.QuickDeployment"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "model.QuickDeploymentTask": {
            "type": "object",
            "properties": {
                "app_code": {
                    "description": "应用编码",
                    "type": "string"
                },
                "app_id": {
                    "description": "应用ID",
                    "type": "integer"
                },
                "app_name": {
                    "description": "应用名称",
                    "type": "string"
                },
                "application": {
                    "description": "关联",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Application"
                        }
                    ]
                },
                "build_number": {
                    "description": "构建编号",
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "deployment_id": {
                    "description": "发布ID",
                    "type": "integer"
                },
                "duration": {
                    "description": "任务耗时(秒)",
                    "type": "integer"
                },
                "end_time": {
                    "description": "任务结束时间",
                    "type": "string"
                },
                "environment": {
                    "description": "环境名称",
                    "type": "string"
                },
                "error_message": {
                    "description": "错误信息",
                    "type": "string"
                },
                "execute_order": {
                    "description": "执行顺序",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "jenkins_env": {
                    "$ref": "#/definitions/model.JenkinsEnv"
                },
                "jenkins_env_id": {
                    "description": "Jenkins环境配置ID",
                    "type": "integer"
                },
                "jenkins_job_url": {
                    "description": "Jenkins任务URL",
                    "type": "string"
                },
                "log_url": {
                    "description": "日志URL",
                    "type": "string"
                },
                "start_time": {
                    "description": "任务开始时间",
                    "type": "string"
                },
                "status": {
                    "description": "任务状态: 1=未部署 2=部署中 3=成功 4=异常",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.ResumeDeploymentResponse": {
            "type": "object",
            "properties": {
                "deploymentName": {
                    "description": "Deployment名称",
                    "type": "string"
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "status": {
                    "description": "当前状态",
                    "type": "string"
                },
                "success": {
                    "description": "是否恢复成功",
                    "type": "boolean"
                }
            }
        },
        "model.RollbackDeploymentRequest": {
            "type": "object",
            "properties": {
                "toRevision": {
                    "description": "目标版本号，0表示回滚到上一版本",
                    "type": "integer"
                }
            }
        },
        "model.RollbackDeploymentResponse": {
            "type": "object",
            "properties": {
                "deploymentName": {
                    "description": "Deployment名称",
                    "type": "string"
                },
                "fromRevision": {
                    "description": "回滚前版本号",
                    "type": "integer"
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "rolloutStatus": {
                    "description": "滚动发布状态",
                    "type": "string"
                },
                "success": {
                    "description": "是否回滚成功",
                    "type": "boolean"
                },
                "toRevision": {
                    "description": "回滚后版本号",
                    "type": "integer"
                }
            }
        },
        "model.RollingUpdateDeployment": {
            "type": "object",
            "properties": {
                "maxSurge": {
                    "description": "最大激增数量",
                    "type": "string"
                },
                "maxUnavailable": {
                    "description": "最大不可用数量",
                    "type": "string"
                }
            }
        },
        "model.ServiceDetail": {
            "type": "object",
            "properties": {
                "clusterIP": {
                    "description": "集群IP",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "endpoints": {
                    "description": "端点信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServiceEndpoint"
                    }
                },
                "events": {
                    "description": "相关事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "externalIPs": {
                    "description": "外部IP列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "description": "服务名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "pods": {
                    "description": "关联的Pod列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sPodInfo"
                    }
                },
                "ports": {
                    "description": "端口配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServicePort"
                    }
                },
                "selector": {
                    "description": "选择器",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "spec": {
                    "description": "完整规格配置"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "服务类型 ClusterIP/NodePort/LoadBalancer/ExternalName",
                    "type": "string"
                }
            }
        },
        "model.ServiceEndpoint": {
            "type": "object",
            "properties": {
                "hostname": {
                    "description": "主机名",
                    "type": "string"
                },
                "ip": {
                    "description": "端点IP",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "ports": {
                    "description": "端口信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.EndpointPort"
                    }
                },
                "ready": {
                    "description": "就绪状态",
                    "type": "boolean"
                }
            }
        },
        "model.ServiceJenkinsEnv": {
            "type": "object",
            "properties": {
                "env_name": {
                    "description": "环境名称",
                    "type": "string"
                },
                "id": {
                    "description": "环境配置ID",
                    "type": "integer"
                },
                "is_configured": {
                    "description": "是否已配置完整",
                    "type": "boolean"
                },
                "jenkins_server_id": {
                    "description": "Jenkins服务器ID",
                    "type": "integer"
                },
                "job_name": {
                    "description": "Jenkins任务名称",
                    "type": "string"
                }
            }
        },
        "model.ServiceListResponse": {
            "type": "object",
            "properties": {
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sService"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.ServicePort": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "端口名称",
                    "type": "string"
                },
                "nodePort": {
                    "description": "节点端口(NodePort类型时使用)",
                    "type": "integer"
                },
                "port": {
                    "description": "服务端口",
                    "type": "integer"
                },
                "protocol": {
                    "description": "协议 TCP/UDP",
                    "type": "string"
                },
                "targetPort": {
                    "description": "目标端口",
                    "type": "string"
                }
            }
        },
        "model.ServicePortSpec": {
            "type": "object",
            "required": [
                "port",
                "protocol"
            ],
            "properties": {
                "name": {
                    "description": "端口名称",
                    "type": "string"
                },
                "nodePort": {
                    "description": "节点端口(NodePort类型时使用)",
                    "type": "integer"
                },
                "port": {
                    "description": "服务端口",
                    "type": "integer"
                },
                "protocol": {
                    "description": "协议",
                    "type": "string"
                },
                "targetPort": {
                    "description": "目标端口",
                    "type": "string"
                }
            }
        },
        "model.ServiceStats": {
            "type": "object",
            "properties": {
                "businessLines": {
                    "description": "业务线数量",
                    "type": "integer"
                },
                "total": {
                    "description": "服务总数",
                    "type": "integer"
                }
            }
        },
        "model.ServiceTreeNode": {
            "type": "object",
            "properties": {
                "business_dept_id": {
                    "description": "业务部门ID",
                    "type": "integer"
                },
                "business_dept_name": {
                    "description": "业务部门名称",
                    "type": "string"
                },
                "code": {
                    "description": "应用编码",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "应用ID",
                    "type": "integer"
                },
                "jenkins_envs": {
                    "description": "Jenkins环境配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServiceJenkinsEnv"
                    }
                },
                "name": {
                    "description": "应用名称",
                    "type": "string"
                },
                "programming_lang": {
                    "description": "编程语言",
                    "type": "string"
                },
                "status": {
                    "description": "应用状态",
                    "type": "integer"
                },
                "status_text": {
                    "description": "状态文本",
                    "type": "string"
                }
            }
        },
        "model.UpdateIngressRequest": {
            "type": "object",
            "properties": {
                "annotations": {
                    "description": "注释",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "class": {
                    "description": "Ingress类",
                    "type": "string"
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "rules": {
                    "description": "路由规则",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressRuleSpec"
                    }
                },
                "tls": {
                    "description": "TLS配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IngressTLSSpec"
                    }
                }
            }
        },
        "model.UpdateNamespaceRequest": {
            "type": "object",
            "properties": {
                "annotations": {
                    "description": "注释",
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
        },
        "model.UpdatePVCRequest": {
            "type": "object",
            "properties": {
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "resources": {
                    "description": "资源配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.PVCResourcesSpec"
                        }
                    ]
                }
            }
        },
        "model.UpdatePVRequest": {
            "type": "object",
            "properties": {
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "mountOptions": {
                    "description": "挂载选项",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "reclaimPolicy": {
                    "description": "回收策略",
                    "type": "string"
                }
            }
        },
        "model.UpdatePodYAMLRequest": {
            "type": "object",
            "required": [
                "yamlContent"
            ],
            "properties": {
                "dryRun": {
                    "description": "是否只进行校验不实际更新",
                    "type": "boolean"
                },
                "force": {
                    "description": "是否强制更新（删除重建）",
                    "type": "boolean"
                },
                "validateOnly": {
                    "description": "是否只校验YAML格式",
                    "type": "boolean"
                },
                "yamlContent": {
                    "description": "YAML内容",
                    "type": "string"
                }
            }
        },
        "model.UpdatePodYAMLResponse": {
            "type": "object",
            "properties": {
                "changes": {
                    "description": "变更说明",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "Pod所在的命名空间",
                    "type": "string"
                },
                "podName": {
                    "description": "更新的Pod名称",
                    "type": "string"
                },
                "success": {
                    "description": "是否更新成功",
                    "type": "boolean"
                },
                "updateStrategy": {
                    "description": "更新策略 (patch/recreate)",
                    "type": "string"
                },
                "validationResult": {
                    "description": "校验结果（DryRun时返回）",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ValidateYAMLResponse"
                        }
                    ]
                },
                "warnings": {
                    "description": "警告信息",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.UpdateServiceRequest": {
            "type": "object",
            "properties": {
                "externalIPs": {
                    "description": "外部IP列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "ports": {
                    "description": "端口配置",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ServicePortSpec"
                    }
                },
                "selector": {
                    "description": "选择器",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "type": {
                    "description": "服务类型",
                    "type": "string"
                }
            }
        }`
