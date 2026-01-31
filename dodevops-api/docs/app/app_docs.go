package docsapp

const AppPaths = `
`

const AppDefinitions = `
        "model.AppEnvironmentResponse": {
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
                "business_dept_id": {
                    "description": "业务部门ID",
                    "type": "integer"
                },
                "business_group_id": {
                    "description": "业务组ID",
                    "type": "integer"
                },
                "environment": {
                    "description": "环境名称",
                    "type": "string"
                },
                "is_configured": {
                    "description": "是否已配置",
                    "type": "boolean"
                },
                "jenkins_job_url": {
                    "description": "Jenkins任务URL",
                    "type": "string"
                },
                "jenkins_server_id": {
                    "description": "Jenkins服务器ID",
                    "type": "integer"
                },
                "jenkins_server_name": {
                    "description": "Jenkins服务器名称",
                    "type": "string"
                },
                "job_name": {
                    "description": "Jenkins任务名称",
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
                    "description": "应用状态文本",
                    "type": "string"
                }
            }
        },
        "model.Application": {
            "type": "object",
            "properties": {
                "business_dept_id": {
                    "description": "业务部门ID(关联sys_dept)",
                    "type": "integer"
                },
                "business_group_id": {
                    "description": "基本信息",
                    "type": "integer"
                },
                "code": {
                    "description": "应用编码",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "databases": {
                    "description": "关联数据库(cmdb_sql表ID)",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "description": "应用介绍",
                    "type": "string"
                },
                "dev_owners": {
                    "description": "负责人信息 (多个用户ID，关联sys_admin表)",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "domains": {
                    "description": "关联资源 (存储资源ID)",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "health_api": {
                    "description": "健康检查接口",
                    "type": "string"
                },
                "hosts": {
                    "description": "关联主机(cmdb_host表ID)",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "jenkins_envs": {
                    "description": "关联的Jenkins环境配置（级联删除）",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsEnv"
                    }
                },
                "name": {
                    "description": "应用名称",
                    "type": "string"
                },
                "ops_owners": {
                    "description": "运维负责人",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "other_res": {
                    "description": "关联其他资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.OtherResources"
                        }
                    ]
                },
                "programming_lang": {
                    "description": "技术信息",
                    "type": "string"
                },
                "repo_url": {
                    "description": "仓库地址",
                    "type": "string"
                },
                "start_command": {
                    "description": "启动命令",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "stop_command": {
                    "description": "停止命令",
                    "type": "string"
                },
                "test_owners": {
                    "description": "测试负责人",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.ApplicationListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Application"
                    }
                },
                "total": {
                    "description": "总数",
                    "type": "integer"
                }
            }
        },
        "model.CreateApplicationRequest": {
            "type": "object",
            "required": [
                "business_dept_id",
                "business_group_id",
                "name"
            ],
            "properties": {
                "business_dept_id": {
                    "description": "业务部门ID",
                    "type": "integer"
                },
                "business_group_id": {
                    "description": "业务组ID",
                    "type": "integer"
                },
                "code": {
                    "description": "应用编码(可选，不提供则根据名称自动生成)",
                    "type": "string"
                },
                "databases": {
                    "description": "关联数据库ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "description": "应用介绍",
                    "type": "string"
                },
                "dev_owners": {
                    "description": "负责人信息",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "domains": {
                    "description": "关联资源",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "health_api": {
                    "description": "健康检查接口",
                    "type": "string"
                },
                "hosts": {
                    "description": "关联主机ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "jenkins_envs": {
                    "description": "Jenkins环境配置(可选，如果不提供则创建默认的3套环境：prod, test, dev)",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CreateJenkinsEnvRequest"
                    }
                },
                "name": {
                    "description": "应用名称",
                    "type": "string"
                },
                "ops_owners": {
                    "description": "运维负责人ID数组",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "other_res": {
                    "description": "其他资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.OtherResources"
                        }
                    ]
                },
                "programming_lang": {
                    "description": "技术信息",
                    "type": "string"
                },
                "repo_url": {
                    "description": "仓库地址",
                    "type": "string"
                },
                "start_command": {
                    "description": "启动命令",
                    "type": "string"
                },
                "stop_command": {
                    "description": "停止命令",
                    "type": "string"
                },
                "test_owners": {
                    "description": "测试负责人ID数组",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.UpdateApplicationRequest": {
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
                "databases": {
                    "description": "关联数据库ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "description": {
                    "description": "应用介绍",
                    "type": "string"
                },
                "dev_owners": {
                    "description": "负责人信息",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "domains": {
                    "description": "关联资源",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "health_api": {
                    "description": "健康检查接口",
                    "type": "string"
                },
                "hosts": {
                    "description": "关联主机ID",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "jenkins_envs": {
                    "description": "Jenkins环境配置(可选，如果提供则完全替换现有配置)",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UpdateJenkinsEnvRequest"
                    }
                },
                "name": {
                    "description": "应用名称",
                    "type": "string"
                },
                "ops_owners": {
                    "description": "运维负责人ID数组",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "other_res": {
                    "description": "其他资源",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.OtherResources"
                        }
                    ]
                },
                "programming_lang": {
                    "description": "技术信息",
                    "type": "string"
                },
                "repo_url": {
                    "description": "仓库地址",
                    "type": "string"
                },
                "start_command": {
                    "description": "启动命令",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "stop_command": {
                    "description": "停止命令",
                    "type": "string"
                },
                "test_owners": {
                    "description": "测试负责人ID数组",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }`
