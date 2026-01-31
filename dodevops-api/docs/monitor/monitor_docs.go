package docsmonitor

const MonitorPaths = `
        "/api/v1/monitor/agent/delete/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除指定的agent数据，用于服务器离线无法正常卸载的情况",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "删除agent数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Agent ID",
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
        "/api/v1/monitor/agent/deploy": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "自动编译agent二进制文件，拷贝到目标主机并启动服务，单个主机传[hostId]，多个主机传[hostId1,hostId2,hostId3]",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "部署agent到指定主机(支持单个或多个)",
                "parameters": [
                    {
                        "description": "部署参数",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BatchDeployAgentDto"
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
        "/api/v1/monitor/agent/heartbeat": {
            "post": {
                "description": "Agent主动上报心跳信息，通过IP自动识别主机",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "更新Agent心跳",
                "parameters": [
                    {
                        "description": "心跳数据",
                        "name": "heartbeat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AgentHeartbeatDto"
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
        "/api/v1/monitor/agent/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取所有Agent的列表信息，支持分页和筛选",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取Agent列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
                        "name": "hostId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页大小",
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
        "/api/v1/monitor/agent/restart/{id}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "重启指定主机上的agent服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "重启agent",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
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
        "/api/v1/monitor/agent/statistics": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取Agent的统计信息，包括各状态数量、平台分布等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取Agent统计信息",
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
        "/api/v1/monitor/agent/status/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定主机上agent的运行状态、版本信息等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取agent状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
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
        "/api/v1/monitor/agent/uninstall": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "停止agent服务并删除相关文件，单个主机传[hostId]，多个主机传[hostId1,hostId2,hostId3]",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "卸载指定主机的agent(支持单个或多个)",
                "parameters": [
                    {
                        "description": "卸载参数(只需hostIds字段)",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BatchDeployAgentDto"
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
        "/api/v1/monitor/host/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取主机的CPU、内存、磁盘使用率",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取主机监控数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
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
        "/api/v1/monitor/hosts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "批量获取主机的CPU、内存、磁盘使用率",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "批量获取主机监控数据,主机ID列表，逗号分隔，如：1,2,3",
                "parameters": [
                    {
                        "type": "string",
                        "description": "主机ID列表，逗号分隔，如：1,2,3",
                        "name": "ids",
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
        "/api/v1/monitor/hosts/{id}/all-metrics": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取主机的CPU、内存、磁盘、网络等所有指标的历史数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取主机所有指标历史数据",
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
                        "example": "\"2025-08-02 15:00:00\"",
                        "description": "开始时间(格式: 2025-08-02 15:00:00)",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"2025-08-02 16:00:00\"",
                        "description": "结束时间(格式: 2025-08-02 16:00:00)",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"1h\"",
                        "description": "时间范围(30m/1h/3h/6h/12h/24h)",
                        "name": "duration",
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
        "/api/v1/monitor/hosts/{id}/history": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取主机的CPU、内存、磁盘使用率历史数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取主机指标历史数据-CPU、内存、磁盘",
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
                        "description": "指标类型(cpu/memory/disk)",
                        "name": "metric",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"2025-08-02 15:00:00\"",
                        "description": "开始时间(格式: 2025-08-02 15:00:00)",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"2025-08-02 16:00:00\"",
                        "description": "结束时间(格式: 2025-08-02 16:00:00)",
                        "name": "end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"1h\"",
                        "description": "时间范围(30m/1h/3h/6h/12h/24h)",
                        "name": "duration",
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
        "/api/v1/monitor/hosts/{id}/ports": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取主机所有TCP端口的监听状态、服务名称和进程信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取主机端口信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
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
        "/api/v1/monitor/hosts/{id}/top-processes": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取主机CPU和内存使用率前5名的进程信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "监控"
                ],
                "summary": "获取主机TOP进程使用率",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主机ID",
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

const MonitorDefinitions = `
        "model.MonitoringInfo": {
            "type": "object",
            "properties": {
                "cpu": {
                    "description": "CPU监控",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ClusterResourceMetrics"
                        }
                    ]
                },
                "memory": {
                    "description": "内存监控",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ClusterResourceMetrics"
                        }
                    ]
                },
                "network": {
                    "description": "网络监控",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NetworkMetrics"
                        }
                    ]
                },
                "storage": {
                    "description": "存储监控",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.StorageMetrics"
                        }
                    ]
                }
            }
        }`
