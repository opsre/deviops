package docsdashboard

const DashboardPaths = `
        "/dashboard/assets": {
            "get": {
                "description": "获取系统资产统计数据，包括主机(按云平台)、数据库(按类型)、K8s集群等资产分布",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "看板管理"
                ],
                "summary": "获取资产统计数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
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
                                            "$ref": "#/definitions/model.AssetStats"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/dashboard/business-distribution": {
            "get": {
                "description": "获取各业务线的服务数量分布，包括总服务数量和各业务线的服务占比",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "看板管理"
                ],
                "summary": "获取业务分布统计数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
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
                                            "$ref": "#/definitions/model.BusinessDistributionStats"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/dashboard/stats": {
            "get": {
                "description": "获取系统看板的各项统计数据，包括主机、K8s集群、发布、任务、服务和数据库统计",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "看板管理"
                ],
                "summary": "获取看板统计数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
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
                                            "$ref": "#/definitions/model.DashboardStats"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }`

const DashboardDefinitions = `
        "model.DashboardStats": {
            "type": "object",
            "properties": {
                "databaseStats": {
                    "description": "数据库统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.DatabaseStats"
                        }
                    ]
                },
                "deploymentStats": {
                    "description": "发布统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.DeploymentStats"
                        }
                    ]
                },
                "hostStats": {
                    "description": "主机统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.HostStats"
                        }
                    ]
                },
                "k8sClusterStats": {
                    "description": "K8s集群统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.K8sClusterStats"
                        }
                    ]
                },
                "serviceStats": {
                    "description": "服务统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ServiceStats"
                        }
                    ]
                },
                "taskStats": {
                    "description": "任务统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.TaskStats"
                        }
                    ]
                }
            }
        }`
