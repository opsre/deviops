package docssystem

const SystemPaths = `

        "/api/v1/admin/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新增用户接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "新增用户接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddSysAdminDto"
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
        "/api/v1/admin/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id删除用户接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysAdminIdDto"
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
        "/api/v1/admin/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id查询用户接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id查询用户接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
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
        "/api/v1/admin/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页获取用户列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "分页获取用户列表接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "帐号启用状态：1-\u003e启用,2-\u003e禁用",
                        "name": "status",
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
        },
        "/api/v1/admin/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改用户接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改用户接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysAdminDto"
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
        "/api/v1/admin/updatePassword": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "重置密码接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "重置密码接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ResetSysAdminPasswordDto"
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
        "/api/v1/admin/updatePersonal": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改个人信息接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改个人信息接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePersonalDto"
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
        "/api/v1/admin/updatePersonalPassword": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改用户密码接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改用户密码接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePersonalPasswordDto"
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
        "/api/v1/admin/updateStatus": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户状态启用/停用接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "用户状态启用/停用接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysAdminStatusDto"
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
        "/api/v1/captcha": {
            "get": {
                "description": "验证码接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "验证码接口",
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
        "/api/v1/dept/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新增部门接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "新增部门接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysDept"
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
        "/api/v1/dept/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除部门接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id删除部门接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysDeptIdDto"
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
        "/api/v1/dept/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id查询部门接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id查询部门接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
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
        "/api/v1/dept/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询部门列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "查询部门列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "部门名称",
                        "name": "deptName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "部门状态",
                        "name": "deptStatus",
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
        "/api/v1/dept/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改部门接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改部门接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysDept"
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
        "/api/v1/dept/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取某部门下的所有用户",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "获取某部门下的所有用户接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "部门ID",
                        "name": "deptId",
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
        "/api/v1/dept/vo/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "部门下拉列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "部门下拉列表接口",
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
        "/api/v1/login": {
            "post": {
                "description": "用户登录接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginDto"
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
        "/api/v1/menu/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新增菜单接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "新增菜单接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenu"
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
        "/api/v1/menu/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除菜单接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id删除菜单接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenuIdDto"
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
        "/api/v1/menu/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id查询菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id查询菜单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
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
        "/api/v1/menu/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询菜单列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "查询菜单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜单名称",
                        "name": "menuName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "菜单状态",
                        "name": "menuStatus",
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
        "/api/v1/menu/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改菜单接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改菜单接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenu"
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
        "/api/v1/menu/vo/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询新增选项列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "查询新增选项列表接口",
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
        "/api/v1/post/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新增岗位接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "新增岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysPost"
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
        "/api/v1/post/batch/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "批量删除岗位接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "批量删除岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelSysPostDto"
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
        "/api/v1/post/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除岗位接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id删除岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysPostIdDto"
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
        "/api/v1/post/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id查询岗位",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id查询岗位",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
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
        "/api/v1/post/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页查询岗位列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "分页查询岗位列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "岗位名称",
                        "name": "postName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "状态：1-\u003e启用,2-\u003e禁用",
                        "name": "postStatus",
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
        },
        "/api/v1/post/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改岗位接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysPost"
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
        "/api/v1/post/updateStatus": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "岗位状态修改接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "岗位状态修改接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysPostStatusDto"
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
        "/api/v1/post/vo/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "岗位下拉列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "岗位下拉列表",
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
        "/api/v1/role/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "新增角色接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "新增角色接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddSysRoleDto"
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
        "/api/v1/role/assignPermissions": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分配权限接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "分配权限接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RoleMenu"
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
        "/api/v1/role/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除角色接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id删除角色接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysRoleIdDto"
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
        "/api/v1/role/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id查询角色接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id查询角色接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
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
        "/api/v1/role/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页查询角色列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "分页查询角色列表接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "roleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "帐号启用状态：1-\u003e启用,2-\u003e禁用",
                        "name": "status",
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
        },
        "/api/v1/role/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "修改角色",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysRoleDto"
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
        "/api/v1/role/updateStatus": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "角色状态启用/停用接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "角色状态启用/停用接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysRoleStatusDto"
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
        "/api/v1/role/vo/idList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据角色id查询菜单数据接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据角色id查询菜单数据接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
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
        "/api/v1/role/vo/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "角色下拉列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "角色下拉列表",
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
        "/api/v1/sysLoginInfo/batch/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "批量删除登录日志接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "批量删除登录日志接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelSysLoginInfoDto"
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
        "/api/v1/sysLoginInfo/clean": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "清空登录日志接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "清空登录日志接口",
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
        "/api/v1/sysLoginInfo/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据ID删除登录日志接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据ID删除登录日志接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysLoginInfoIdDto"
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
        "/api/v1/sysLoginInfo/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页获取登录日志列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "分页获取登录日志列表接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "登录状态（1-成功 2-失败）",
                        "name": "loginStatus",
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
        },
        "/api/v1/sysOperationLog/batch/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "批量删除操作日志接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "批量删除操作日志接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.BatchDeleteSysOperationLogDto"
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
        "/api/v1/sysOperationLog/clean": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "清空操作日志接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "清空操作日志接口",
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
        "/api/v1/sysOperationLog/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据id删除操作日志",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "根据id删除操作日志",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysOperationLogIdDto"
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
        "/api/v1/sysOperationLog/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "分页获取操作日志列表接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "分页获取操作日志列表接口",
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
                        "description": "用户名",
                        "name": "username",
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
        },
        "/api/v1/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "单图片上传接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System系统管理"
                ],
                "summary": "单图片上传接口",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
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
        "/apps": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取应用列表，支持分页和筛选",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "获取应用列表",
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
                    },
                    {
                        "type": "string",
                        "description": "应用名称(模糊查询)",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "应用编码(模糊查询)",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "应用类型",
                        "name": "app_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "status",
                        "in": "query"
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
                                            "$ref": "#/definitions/model.ApplicationListResponse"
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
                "description": "创建新的应用，应用编码(code)为可选参数，不提供则根据应用名称自动生成",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "创建应用",
                "parameters": [
                    {
                        "description": "创建应用请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateApplicationRequest"
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
                                            "$ref": "#/definitions/model.Application"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/business-group-options": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取业务组和业务部门的树形选择器数据，支持二级分组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ServiceTree"
                ],
                "summary": "获取业务组选项",
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
                                                "type": "object",
                                                "additionalProperties": true
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
        "/apps/deployment/applications": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据业务组、部门和环境获取可发布的应用列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "获取可发布的应用列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "业务组ID",
                        "name": "business_group_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "业务部门ID",
                        "name": "business_dept_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "环境名称",
                        "name": "environment",
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
                                                "$ref": "#/definitions/model.ApplicationForDeployment"
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
        "/apps/deployment/execute": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "启动快速发布流程，支持串行或并行执行模式",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "执行快速发布",
                "parameters": [
                    {
                        "description": "执行快速发布请求，支持选择执行模式",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ExecuteQuickDeploymentRequest"
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
        "/apps/deployment/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取快速发布列表，支持分页和筛选",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "获取快速发布列表",
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
                    },
                    {
                        "type": "integer",
                        "description": "业务组ID",
                        "name": "business_group_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "业务部门ID",
                        "name": "business_dept_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "环境名称",
                        "name": "environment",
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
                        "description": "创建人ID",
                        "name": "creator_id",
                        "in": "query"
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
                                            "$ref": "#/definitions/model.QuickDeploymentListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/deployment/quick": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建快速发布流程，包含多个应用的发布任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "创建快速发布",
                "parameters": [
                    {
                        "description": "创建快速发布请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateQuickDeploymentRequest"
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
                                            "$ref": "#/definitions/model.QuickDeployment"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/deployment/tasks/{task_id}/log": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取快速发布任务的Jenkins构建日志",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "获取任务构建日志",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "日志起始位置",
                        "name": "start",
                        "in": "query"
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
        "/apps/deployment/tasks/{task_id}/status": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取快速发布任务的实时状态信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "获取任务状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "task_id",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.TaskStatusResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/deployment/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取快速发布详情，包含所有任务信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "QuickDeployment"
                ],
                "summary": "获取快速发布详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "发布ID",
                        "name": "id",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.QuickDeployment"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/environment": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定应用在指定环境的详细配置信息，包括Jenkins配置状态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "获取应用环境配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "app_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "环境名称",
                        "name": "environment",
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
                                            "$ref": "#/definitions/model.AppEnvironmentResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/jenkins-job/validate": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "验证指定Jenkins服务器中是否存在指定的任务名称",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "验证Jenkins任务是否存在",
                "parameters": [
                    {
                        "description": "验证Jenkins任务请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ValidateJenkinsJobRequest"
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
                                            "$ref": "#/definitions/model.ValidateJenkinsJobResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/jenkins-servers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取所有类型为Jenkins(type=4)的服务器配置信息，用于Jenkins环境配置选择",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "获取Jenkins服务器列表",
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
                                                "$ref": "#/definitions/model.JenkinsServerOption"
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
        "/apps/service-tree": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据业务线查询服务，按业务线重新组装排序服务，类似于服务树结构",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ServiceTree"
                ],
                "summary": "获取业务线服务树",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "业务组ID列表，为空则查询所有",
                        "name": "business_group_ids",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "应用状态筛选，为空则查询所有状态",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "环境名称筛选，为空则不筛选环境配置",
                        "name": "environment",
                        "in": "query"
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
                                                "$ref": "#/definitions/model.BusinessLineServiceTree"
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
        "/apps/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据应用ID获取应用详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "获取应用详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.Application"
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
                "description": "更新应用信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "更新应用",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新应用请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateApplicationRequest"
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
                                            "$ref": "#/definitions/model.Application"
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
                "description": "删除指定的应用",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "删除应用",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
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
        "/apps/{id}/jenkins-envs": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据应用ID获取所有Jenkins环境配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "JenkinsEnv"
                ],
                "summary": "获取应用的所有Jenkins环境配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
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
                                                "$ref": "#/definitions/model.JenkinsEnv"
                                            }
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
                "description": "为指定应用添加新的Jenkins环境配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "为应用添加Jenkins环境配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Jenkins环境配置请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateJenkinsEnvRequest"
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
                                            "$ref": "#/definitions/model.JenkinsEnv"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/apps/{id}/jenkins-envs/{env_id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新指定应用的Jenkins环境配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "更新应用的Jenkins环境配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Jenkins环境配置ID",
                        "name": "env_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新Jenkins环境配置请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateJenkinsEnvRequest"
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
                                            "$ref": "#/definitions/model.JenkinsEnv"
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
                "description": "删除指定应用的Jenkins环境配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Application"
                ],
                "summary": "删除应用的Jenkins环境配置",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Jenkins环境配置ID",
                        "name": "env_id",
                        "in": "path",
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
        "/jenkins/servers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取所有配置的Jenkins服务器",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins服务器列表",
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
                            "allOf": [
                                {
                                    "$ref": "#/definitions/result.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.JenkinsServerListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/servers/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据服务器ID获取Jenkins服务器详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins服务器详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "id",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.JenkinsServerInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/test-connection": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "测试Jenkins服务器连接是否正常",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "测试Jenkins连接",
                "parameters": [
                    {
                        "description": "连接测试请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TestJenkinsConnectionRequest"
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
                                            "$ref": "#/definitions/model.TestJenkinsConnectionResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定Jenkins服务器的所有任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins任务列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.JenkinsJobListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs/search": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据关键词模糊搜索指定Jenkins服务器的任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "搜索Jenkins任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "搜索关键词",
                        "name": "keyword",
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
                                            "$ref": "#/definitions/model.JenkinsJobListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs/{jobName}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定任务的详细信息和构建历史",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins任务详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务名称",
                        "name": "jobName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.JenkinsJobDetailResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs/{jobName}/builds/{buildNumber}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定构建的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins构建详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务名称",
                        "name": "jobName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "构建编号",
                        "name": "buildNumber",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.JenkinsBuildDetailResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs/{jobName}/builds/{buildNumber}/log": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定构建的日志信息，支持分页获取",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins构建日志",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务名称",
                        "name": "jobName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "构建编号",
                        "name": "buildNumber",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "开始位置",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "是否返回HTML格式",
                        "name": "html",
                        "in": "query"
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
                                            "$ref": "#/definitions/model.GetBuildLogResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs/{jobName}/builds/{buildNumber}/stop": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "停止指定的Jenkins构建任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "停止Jenkins构建",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务名称",
                        "name": "jobName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "构建编号",
                        "name": "buildNumber",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "停止构建请求",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.StopBuildRequest"
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
                                            "$ref": "#/definitions/model.StopBuildResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/jobs/{jobName}/start": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "启动指定的Jenkins任务，支持带参数构建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "启动Jenkins任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务名称",
                        "name": "jobName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "启动任务请求",
                        "name": "request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.StartJobRequest"
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
                                            "$ref": "#/definitions/model.StartJobResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/queue": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定Jenkins服务器的构建队列信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins队列信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.JenkinsQueue"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/jenkins/{serverId}/system-info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取指定Jenkins服务器的系统信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Jenkins"
                ],
                "summary": "获取Jenkins系统信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "服务器ID",
                        "name": "serverId",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.JenkinsSystemInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/k8s/cluster/{id}/namespaces/{namespaceName}/configmaps": {
            "get": {
                "description": "获取指定集群和命名空间下的ConfigMap列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "获取ConfigMap列表",
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
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页大小",
                        "name": "pageSize",
                        "in": "query"
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
                                            "$ref": "#/definitions/model.ConfigMapListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建新的ConfigMap",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "创建ConfigMap",
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
                        "description": "创建ConfigMap请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateConfigMapRequest"
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
                                            "$ref": "#/definitions/model.K8sConfigMap"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/k8s/cluster/{id}/namespaces/{namespaceName}/configmaps/{configMapName}": {
            "get": {
                "description": "获取指定ConfigMap的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "获取ConfigMap详情",
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
                        "description": "ConfigMap名称",
                        "name": "configMapName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.ConfigMapDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "更新指定的ConfigMap",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "更新ConfigMap",
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
                        "description": "ConfigMap名称",
                        "name": "configMapName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新ConfigMap请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateConfigMapRequest"
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
                                            "$ref": "#/definitions/model.K8sConfigMap"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "删除指定的ConfigMap",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "删除ConfigMap",
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
                        "description": "ConfigMap名称",
                        "name": "configMapName",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/configmaps/{configMapName}/yaml": {
            "get": {
                "description": "获取指定ConfigMap的YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "获取ConfigMap YAML",
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
                        "description": "ConfigMap名称",
                        "name": "configMapName",
                        "in": "path",
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
            },
            "put": {
                "description": "通过YAML更新ConfigMap配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "更新ConfigMap YAML",
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
                        "description": "ConfigMap名称",
                        "name": "configMapName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "$ref": "#/definitions/model.K8sConfigMap"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments": {
            "post": {
                "description": "在指定命名空间中创建新的Deployment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment管理"
                ],
                "summary": "创建Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment配置",
                        "name": "deployment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateDeploymentRequest"
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
                                            "$ref": "#/definitions/model.K8sWorkload"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}": {
            "put": {
                "description": "更新指定的Deployment配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment管理"
                ],
                "summary": "更新Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新配置",
                        "name": "deployment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateWorkloadRequest"
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
                                            "$ref": "#/definitions/model.K8sWorkload"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的Deployment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment管理"
                ],
                "summary": "删除Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/history": {
            "get": {
                "description": "获取指定Deployment的所有版本历史信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment版本管理"
                ],
                "summary": "获取Deployment版本历史",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.DeploymentRolloutHistoryResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/pause": {
            "post": {
                "description": "暂停正在进行的Deployment滚动更新过程",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment版本管理"
                ],
                "summary": "暂停Deployment滚动更新",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.PauseDeploymentResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/restart": {
            "post": {
                "description": "通过更新Pod模板来重启Deployment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment管理"
                ],
                "summary": "重启Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/resume": {
            "post": {
                "description": "恢复被暂停的Deployment滚动更新过程",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment版本管理"
                ],
                "summary": "恢复Deployment滚动更新",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.ResumeDeploymentResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/revisions/{revision}": {
            "get": {
                "description": "获取指定Deployment特定版本的详细配置信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment版本管理"
                ],
                "summary": "获取Deployment指定版本详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "版本号",
                        "name": "revision",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.DeploymentRevisionDetail"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/rollback": {
            "post": {
                "description": "将Deployment回滚到指定的历史版本",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment版本管理"
                ],
                "summary": "回滚Deployment到指定版本",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "回滚请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RollbackDeploymentRequest"
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
                                            "$ref": "#/definitions/model.RollbackDeploymentResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/rollout-status": {
            "get": {
                "description": "获取指定Deployment的当前滚动发布状态和进度信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment版本管理"
                ],
                "summary": "获取Deployment滚动发布状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.DeploymentRolloutStatusResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/deployments/{deploymentName}/scale": {
            "post": {
                "description": "调整Deployment的副本数",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment管理"
                ],
                "summary": "伸缩Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Deployment名称",
                        "name": "deploymentName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "伸缩配置",
                        "name": "scale",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ScaleWorkloadRequest"
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
                                            "type": "string"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/ingresses": {
            "get": {
                "description": "获取指定命名空间的Ingress列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "获取Ingress列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.IngressListResponse"
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
            },
            "post": {
                "description": "在指定命名空间中创建新的Ingress",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "创建Ingress",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress配置",
                        "name": "ingress",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateIngressRequest"
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
                                            "$ref": "#/definitions/model.K8sIngress"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/ingresses/{ingressName}": {
            "get": {
                "description": "获取指定Ingress的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "获取Ingress详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.IngressDetail"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "更新指定的Ingress配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "更新Ingress",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新配置",
                        "name": "ingress",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateIngressRequest"
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
                                            "$ref": "#/definitions/model.K8sIngress"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的Ingress",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "删除Ingress",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/ingresses/{ingressName}/events": {
            "get": {
                "description": "获取指定Ingress的相关事件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "获取Ingress事件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
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
                                                "$ref": "#/definitions/model.K8sEvent"
                                            }
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/ingresses/{ingressName}/monitoring": {
            "get": {
                "description": "获取指定Ingress的监控指标和状态信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "获取Ingress监控信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/ingresses/{ingressName}/yaml": {
            "get": {
                "description": "获取指定Ingress的完整YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "获取Ingress的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "通过提供的YAML内容更新Ingress配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Ingress管理"
                ],
                "summary": "通过YAML更新Ingress",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Ingress名称",
                        "name": "ingressName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "yaml",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "type": "string"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/metrics": {
            "get": {
                "description": "获取指定命名空间下所有Pod的CPU、内存等监控指标汇总",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s监控"
                ],
                "summary": "获取命名空间监控指标",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.NamespaceMetricsInfo"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods": {
            "get": {
                "description": "获取指定命名空间的Pod列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "获取Pod列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/yaml": {
            "post": {
                "description": "通过提供的YAML内容创建Pod，支持校验模式和DryRun模式",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "通过YAML创建Pod",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "创建请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePodFromYAMLRequest"
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
                                            "$ref": "#/definitions/model.CreatePodFromYAMLResponse"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}": {
            "get": {
                "description": "获取指定Pod的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "获取Pod详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.K8sPodDetail"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的Pod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "删除Pod",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}/containers": {
            "get": {
                "description": "获取指定Pod中所有容器的名称列表，用于终端连接时选择容器",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s容器终端"
                ],
                "summary": "获取Pod中的容器列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
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
                                                "type": "string"
                                            }
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}/events": {
            "get": {
                "description": "获取指定Pod的相关事件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "获取Pod事件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}/logs": {
            "get": {
                "description": "获取指定Pod容器的日志",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "获取Pod日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "容器名称",
                        "name": "container",
                        "in": "query"
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}/metrics": {
            "get": {
                "description": "获取指定Pod的CPU、内存等监控指标和使用率",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s监控"
                ],
                "summary": "获取Pod监控指标",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.PodMetricsInfo"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}/terminal": {
            "get": {
                "description": "通过WebSocket连接到指定Pod的终端",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s容器终端"
                ],
                "summary": "连接到Pod终端",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "容器名称（默认为Pod中第一个容器）",
                        "name": "containerName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "执行命令（默认为/bin/bash）",
                        "name": "command",
                        "in": "query"
                    }
                ],
                "responses": {
                    "101": {
                        "description": "Switching Protocols"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pods/{podName}/yaml": {
            "get": {
                "description": "获取指定Pod的完整YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "获取Pod的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "通过YAML内容更新指定的Pod配置，支持校验模式和DryRun模式",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "更新Pod的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Pod名称",
                        "name": "podName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePodYAMLRequest"
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
                                            "$ref": "#/definitions/model.UpdatePodYAMLResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pvcs": {
            "get": {
                "description": "获取指定命名空间的PVC列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "获取PVC列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.PVCListResponse"
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
            },
            "post": {
                "description": "在指定命名空间中创建新的PVC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "创建PVC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "PVC配置",
                        "name": "pvc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePVCRequest"
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
                                            "$ref": "#/definitions/model.K8sPersistentVolumeClaim"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pvcs/{pvcName}": {
            "get": {
                "description": "获取指定PVC的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "获取PVC详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "PVC名称",
                        "name": "pvcName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.PVCDetail"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "更新指定的PVC配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "更新PVC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "PVC名称",
                        "name": "pvcName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新配置",
                        "name": "pvc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePVCRequest"
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
                                            "$ref": "#/definitions/model.K8sPersistentVolumeClaim"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的PVC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "删除PVC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "PVC名称",
                        "name": "pvcName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/pvcs/{pvcName}/yaml": {
            "get": {
                "description": "获取指定PVC的完整YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "获取PVC的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "PVC名称",
                        "name": "pvcName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "通过提供的YAML内容更新PVC配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PVC"
                ],
                "summary": "通过YAML更新PVC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "PVC名称",
                        "name": "pvcName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "yaml",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "type": "string"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/secrets": {
            "get": {
                "description": "获取指定集群和命名空间下的Secret列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "获取Secret列表",
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
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每页大小",
                        "name": "pageSize",
                        "in": "query"
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
                                            "$ref": "#/definitions/model.SecretListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "创建新的Secret",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "创建Secret",
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
                        "description": "创建Secret请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateSecretRequest"
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
                                            "$ref": "#/definitions/model.K8sSecret"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/k8s/cluster/{id}/namespaces/{namespaceName}/secrets/{secretName}": {
            "get": {
                "description": "获取指定Secret的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "获取Secret详情",
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
                        "description": "Secret名称",
                        "name": "secretName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.SecretDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "更新指定的Secret",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "更新Secret",
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
                        "description": "Secret名称",
                        "name": "secretName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新Secret请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSecretRequest"
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
                                            "$ref": "#/definitions/model.K8sSecret"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "删除指定的Secret",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "删除Secret",
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
                        "description": "Secret名称",
                        "name": "secretName",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/secrets/{secretName}/yaml": {
            "get": {
                "description": "获取指定Secret的YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "获取Secret YAML",
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
                        "description": "Secret名称",
                        "name": "secretName",
                        "in": "path",
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
            },
            "put": {
                "description": "通过YAML更新Secret配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s配置管理"
                ],
                "summary": "更新Secret YAML",
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
                        "description": "Secret名称",
                        "name": "secretName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "$ref": "#/definitions/model.K8sSecret"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/k8s/cluster/{id}/namespaces/{namespaceName}/services": {
            "get": {
                "description": "获取指定命名空间的Service列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "获取Service列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.ServiceListResponse"
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
            },
            "post": {
                "description": "在指定命名空间中创建新的Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "创建Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service配置",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateServiceRequest"
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
                                            "$ref": "#/definitions/model.K8sService"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/services/{serviceName}": {
            "get": {
                "description": "获取指定Service的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "获取Service详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service名称",
                        "name": "serviceName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.ServiceDetail"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "更新指定的Service配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "更新Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service名称",
                        "name": "serviceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新配置",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateServiceRequest"
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
                                            "$ref": "#/definitions/model.K8sService"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "删除Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service名称",
                        "name": "serviceName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/services/{serviceName}/events": {
            "get": {
                "description": "获取指定Service的相关事件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "获取Service事件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service名称",
                        "name": "serviceName",
                        "in": "path",
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
                                                "$ref": "#/definitions/model.K8sEvent"
                                            }
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/services/{serviceName}/yaml": {
            "get": {
                "description": "获取指定Service的完整YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "获取Service的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service名称",
                        "name": "serviceName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "通过提供的YAML内容更新Service配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s Service管理"
                ],
                "summary": "通过YAML更新Service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "Service名称",
                        "name": "serviceName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "yaml",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "type": "string"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/workloads": {
            "get": {
                "description": "获取指定命名空间的工作负载列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s工作负载管理"
                ],
                "summary": "获取工作负载列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                            "deployment",
                            "statefulset",
                            "daemonset",
                            "job",
                            "cronjob",
                            "all"
                        ],
                        "type": "string",
                        "description": "工作负载类型",
                        "name": "type",
                        "in": "query"
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
                                            "$ref": "#/definitions/model.WorkloadListResponse"
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/workloads/yaml": {
            "put": {
                "description": "通过YAML内容更新指定的工作负载配置，支持deployment,statefulset,daemonset,job,cronjob",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作负载YAML管理"
                ],
                "summary": "更新工作负载的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                        "description": "更新请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateWorkloadYAMLRequest"
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
                                            "$ref": "#/definitions/model.UpdateWorkloadYAMLResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/workloads/{type}/{workloadName}": {
            "get": {
                "description": "获取指定工作负载的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s工作负载管理"
                ],
                "summary": "获取工作负载详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                            "deployment",
                            "statefulset",
                            "daemonset",
                            "job",
                            "cronjob"
                        ],
                        "type": "string",
                        "description": "工作负载类型",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "工作负载名称",
                        "name": "workloadName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.K8sWorkloadDetail"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/workloads/{type}/{workloadName}/pods": {
            "get": {
                "description": "获取指定工作负载下的所有Pod信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod管理"
                ],
                "summary": "获取工作负载下的Pod列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                            "deployment",
                            "statefulset",
                            "daemonset",
                            "job",
                            "cronjob"
                        ],
                        "type": "string",
                        "description": "工作负载类型",
                        "name": "type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "工作负载名称",
                        "name": "workloadName",
                        "in": "path",
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
                                                "$ref": "#/definitions/model.K8sPodInfo"
                                            }
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
        "/k8s/cluster/{id}/namespaces/{namespaceName}/workloads/{workloadType}/{workloadName}/yaml": {
            "get": {
                "description": "获取指定工作负载的完整YAML配置，支持deployment,statefulset,daemonset,job,cronjob",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作负载YAML管理"
                ],
                "summary": "获取工作负载的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                            "deployment",
                            "statefulset",
                            "daemonset",
                            "job",
                            "cronjob"
                        ],
                        "type": "string",
                        "description": "工作负载类型",
                        "name": "workloadType",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "工作负载名称",
                        "name": "workloadName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.GetWorkloadYAMLResponse"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/nodes/{nodeName}/metrics": {
            "get": {
                "description": "获取指定节点的CPU、内存等监控指标和使用率",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s监控"
                ],
                "summary": "获取节点监控指标",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.NodeMetricsInfo"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/pvs": {
            "get": {
                "description": "获取集群中的PV列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "获取PV列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.PVListResponse"
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
            },
            "post": {
                "description": "在集群中创建新的PV",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "创建PV",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PV配置",
                        "name": "pv",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePVRequest"
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
                                            "$ref": "#/definitions/model.K8sPersistentVolume"
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
        "/k8s/cluster/{id}/pvs/{pvName}": {
            "get": {
                "description": "获取指定PV的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "获取PV详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PV名称",
                        "name": "pvName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.PVDetail"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "更新指定的PV配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "更新PV",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PV名称",
                        "name": "pvName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新配置",
                        "name": "pv",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePVRequest"
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
                                            "$ref": "#/definitions/model.K8sPersistentVolume"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的PV",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "删除PV",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PV名称",
                        "name": "pvName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/pvs/{pvName}/yaml": {
            "get": {
                "description": "获取指定PV的完整YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "获取PV的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PV名称",
                        "name": "pvName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "通过提供的YAML内容更新PV配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-PV"
                ],
                "summary": "通过YAML更新PV",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PV名称",
                        "name": "pvName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "yaml",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "type": "string"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/storageclasses": {
            "get": {
                "description": "获取集群中的存储类列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "获取存储类列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
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
                                            "$ref": "#/definitions/model.StorageClassListResponse"
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
            },
            "post": {
                "description": "在集群中创建新的存储类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "创建存储类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "存储类配置",
                        "name": "storageClass",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateStorageClassRequest"
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
                                            "$ref": "#/definitions/model.K8sStorageClass"
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
        "/k8s/cluster/{id}/storageclasses/{storageClassName}": {
            "get": {
                "description": "获取指定存储类的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "获取存储类详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "存储类名称",
                        "name": "storageClassName",
                        "in": "path",
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
                                            "$ref": "#/definitions/model.StorageClassDetail"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "更新指定的存储类配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "更新存储类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "存储类名称",
                        "name": "storageClassName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新配置",
                        "name": "storageClass",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateStorageClassRequest"
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
                                            "$ref": "#/definitions/model.K8sStorageClass"
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
                    "404": {
                        "description": "Not Found",
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
            },
            "delete": {
                "description": "删除指定的存储类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "删除存储类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "存储类名称",
                        "name": "storageClassName",
                        "in": "path",
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/storageclasses/{storageClassName}/yaml": {
            "get": {
                "description": "获取指定存储类的完整YAML配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "获取存储类的YAML配置",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "存储类名称",
                        "name": "storageClassName",
                        "in": "path",
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
                                            "additionalProperties": true
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
                    "404": {
                        "description": "Not Found",
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
            },
            "put": {
                "description": "通过提供的YAML内容更新存储类配置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "K8s存储管理-StorageClass"
                ],
                "summary": "通过YAML更新存储类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "存储类名称",
                        "name": "storageClassName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "YAML内容",
                        "name": "yaml",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
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
                                            "type": "string"
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
                    "404": {
                        "description": "Not Found",
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
        "/k8s/cluster/{id}/yaml/validate": {
            "post": {
                "description": "校验提供的YAML内容是否符合Kubernetes资源规范",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "YAML校验"
                ],
                "summary": "校验YAML格式",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "集群ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "校验请求",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ValidateYAMLRequest"
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
                                            "$ref": "#/definitions/model.ValidateYAMLResponse"
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
        "/task/monitor/queue/clear-failed": {
            "post": {
                "description": "清空失败任务队列中的所有任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "清空失败队列",
                "responses": {
                    "200": {
                        "description": "OK",
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
        "/task/monitor/queue/details": {
            "get": {
                "description": "获取各个优先级队列的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "获取队列详细信息",
                "responses": {
                    "200": {
                        "description": "OK",
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
        "/task/monitor/queue/metrics": {
            "get": {
                "description": "获取任务队列的运行指标，包括队列长度、处理统计等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "获取任务队列指标",
                "responses": {
                    "200": {
                        "description": "OK",
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
        "/task/monitor/queue/retry-failed": {
            "post": {
                "description": "将失败队列中的任务重新提交到正常队列",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "重试失败队列中的任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "重试任务数量限制，默认10，最大100",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
        "/task/monitor/scheduled/pause": {
            "post": {
                "description": "暂停正在运行的定时任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "暂停定时任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "task_id",
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
                    },
                    "400": {
                        "description": "Bad Request",
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
        "/task/monitor/scheduled/reset": {
            "post": {
                "description": "将定时任务的所有子任务状态重置为等待中(1)，用于修复状态异常的情况",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "重置定时任务子任务状态",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "task_id",
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
                    },
                    "400": {
                        "description": "Bad Request",
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
        "/task/monitor/scheduled/resume": {
            "post": {
                "description": "恢复已暂停的定时任务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "恢复定时任务",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "task_id",
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
                    },
                    "400": {
                        "description": "Bad Request",
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
        "/task/monitor/scheduler/stats": {
            "get": {
                "description": "获取全局调度器的统计信息，包括活跃任务数、下次运行时间等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "获取调度器统计信息",
                "responses": {
                    "200": {
                        "description": "OK",
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
        "/task/monitor/system/status": {
            "get": {
                "description": "获取任务队列和调度器的整体运行状态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "获取任务系统整体状态",
                "responses": {
                    "200": {
                        "description": "OK",
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
        "/task/monitor/task/status": {
            "get": {
                "description": "获取任务的详细状态信息，包括可执行的操作",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "任务监控"
                ],
                "summary": "获取任务状态详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "task_id",
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
                    },
                    "400": {
                        "description": "Bad Request",
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

const SystemDefinitions = `
        "controller.ListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TaskAnsible"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "controller.SQLRequest": {
            "type": "object",
            "properties": {
                "databaseId": {
                    "description": "数据库ID",
                    "type": "integer"
                },
                "databaseName": {
                    "description": "数据库名称",
                    "type": "string"
                },
                "sql": {
                    "description": "SQL语句",
                    "type": "string"
                }
            }
        },
        "model.AddLabelRequest": {
            "type": "object",
            "required": [
                "key",
                "value"
            ],
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.AddSysAdminDto": {
            "type": "object",
            "required": [
                "deptId",
                "email",
                "nickname",
                "password",
                "phone",
                "postId",
                "roleId",
                "status",
                "username"
            ],
            "properties": {
                "deptId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "note": {
                    "description": "备注",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "postId": {
                    "description": "岗位id",
                    "type": "integer"
                },
                "roleId": {
                    "description": "角色id",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.AddSysRoleDto": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "roleKey": {
                    "description": "角色key",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "model.AddTaintRequest": {
            "type": "object",
            "required": [
                "effect",
                "key"
            ],
            "properties": {
                "effect": {
                    "type": "string",
                    "enum": [
                        "NoSchedule",
                        "PreferNoSchedule",
                        "NoExecute"
                    ]
                },
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.AgentHeartbeatDto": {
            "type": "object",
            "properties": {
                "hostname": {
                    "description": "Agent所在主机的hostname",
                    "type": "string"
                },
                "ip": {
                    "description": "Agent所在主机的IP地址 (内网IP)",
                    "type": "string"
                },
                "pid": {
                    "description": "Agent进程ID",
                    "type": "integer"
                },
                "port": {
                    "description": "Agent监听端口",
                    "type": "integer"
                },
                "token": {
                    "description": "认证token",
                    "type": "string"
                }
            }
        },
        "model.AssetCategoryStats": {
            "type": "object",
            "properties": {
                "category": {
                    "description": "资产分类名称 (主机/数据库/K8s集群)",
                    "type": "string"
                },
                "items": {
                    "description": "具体资产项统计",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AssetItemStats"
                    }
                },
                "total": {
                    "description": "该分类总数",
                    "type": "integer"
                }
            }
        },
        "model.AssetItemStats": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "数量",
                    "type": "integer"
                },
                "name": {
                    "description": "资产项名称 (自建主机/阿里云/MySQL等)",
                    "type": "string"
                }
            }
        },
        "model.AssetStats": {
            "type": "object",
            "properties": {
                "categories": {
                    "description": "资产分类统计",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AssetCategoryStats"
                    }
                },
                "totalAssets": {
                    "description": "总资产数量",
                    "type": "integer"
                }
            }
        },
        "model.BatchDeleteSysOperationLogDto": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "id列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.BatchDeployAgentDto": {
            "type": "object",
            "required": [
                "hostIds"
            ],
            "properties": {
                "hostIds": {
                    "description": "主机ID列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "version": {
                    "description": "版本",
                    "type": "string"
                }
            }
        },
        "model.BuildAction": {
            "type": "object",
            "properties": {
                "_class": {
                    "description": "操作类型",
                    "type": "string"
                }
            }
        },
        "model.BuildExecutor": {
            "type": "object",
            "properties": {
                "node": {
                    "description": "节点名称",
                    "type": "string"
                },
                "number": {
                    "description": "执行器编号",
                    "type": "integer"
                }
            }
        },
        "model.BusinessDistributionStats": {
            "type": "object",
            "properties": {
                "businessLines": {
                    "description": "业务线列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BusinessLineStats"
                    }
                },
                "totalServices": {
                    "description": "总服务数量",
                    "type": "integer"
                }
            }
        },
        "model.BusinessLineStats": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "业务组ID",
                    "type": "integer"
                },
                "name": {
                    "description": "业务线名称",
                    "type": "string"
                },
                "percentage": {
                    "description": "占比",
                    "type": "number"
                },
                "serviceCount": {
                    "description": "服务数量",
                    "type": "integer"
                }
            }
        },
        "model.ChangeSet": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "变更项目",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ChangeSetItem"
                    }
                },
                "kind": {
                    "description": "变更类型",
                    "type": "string"
                }
            }
        },
        "model.ChangeSetItem": {
            "type": "object",
            "properties": {
                "affectedPaths": {
                    "description": "影响路径",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "author": {
                    "description": "作者",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.User"
                        }
                    ]
                },
                "comment": {
                    "description": "提交注释",
                    "type": "string"
                },
                "date": {
                    "description": "提交日期",
                    "type": "string"
                },
                "id": {
                    "description": "提交ID",
                    "type": "string"
                },
                "msg": {
                    "description": "提交消息",
                    "type": "string"
                },
                "timestamp": {
                    "description": "时间戳",
                    "type": "integer"
                }
            }
        },
        "model.ClusterDetailResponse": {
            "type": "object",
            "properties": {
                "cluster": {
                    "$ref": "#/definitions/model.KubeCluster"
                },
                "components": {
                    "description": "安装的组件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ComponentInfo"
                    }
                },
                "events": {
                    "description": "集群事件",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ClusterEvent"
                    }
                },
                "monitoring": {
                    "description": "监控信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.MonitoringInfo"
                        }
                    ]
                },
                "network": {
                    "description": "网络配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NetworkInfo"
                        }
                    ]
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.NodeInfo"
                    }
                },
                "runtime": {
                    "description": "运行时信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RuntimeSummary"
                        }
                    ]
                },
                "summary": {
                    "$ref": "#/definitions/model.ClusterSummary"
                },
                "workloads": {
                    "description": "工作负载统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.WorkloadSummary"
                        }
                    ]
                }
            }
        },
        "model.ClusterEvent": {
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
                "involvedObject": {
                    "description": "相关对象",
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
        "model.ClusterSummary": {
            "type": "object",
            "properties": {
                "masterNodes": {
                    "type": "integer"
                },
                "readyNodes": {
                    "type": "integer"
                },
                "totalNodes": {
                    "type": "integer"
                },
                "workerNodes": {
                    "type": "integer"
                }
            }
        },
        "model.ComponentInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "组件名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                },
                "type": {
                    "description": "组件类型 (system/addon)",
                    "type": "string"
                },
                "version": {
                    "description": "版本",
                    "type": "string"
                }
            }
        },
        "model.CreateDeployDto": {
            "type": "object",
            "required": [
                "hostId",
                "installDir",
                "serviceId",
                "version"
            ],
            "properties": {
                "autoStart": {
                    "description": "是否自动启动",
                    "type": "boolean"
                },
                "envVars": {
                    "description": "环境变量",
                    "type": "object",
                    "additionalProperties": true
                },
                "hostId": {
                    "description": "主机ID",
                    "type": "integer"
                },
                "installDir": {
                    "description": "安装目录",
                    "type": "string"
                },
                "serviceId": {
                    "description": "服务ID (如: mysql)",
                    "type": "string"
                },
                "version": {
                    "description": "版本 (如: 5.7)",
                    "type": "string"
                }
            }
        },
        "model.CreateEcsPasswordAuthDto": {
            "type": "object",
            "required": [
                "name",
                "password",
                "port",
                "type",
                "username"
            ],
            "properties": {
                "name": {
                    "description": "凭证名称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "port": {
                    "description": "端口号",
                    "type": "integer"
                },
                "publicKey": {
                    "description": "公钥",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "type": {
                    "description": "认证类型:1-\u003e密码",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.CreateJenkinsEnvRequest": {
            "type": "object",
            "required": [
                "env_name"
            ],
            "properties": {
                "app_id": {
                    "description": "应用ID(由控制器自动设置)",
                    "type": "integer"
                },
                "env_name": {
                    "description": "环境名称",
                    "type": "string"
                },
                "jenkins_server_id": {
                    "description": "Jenkins服务器ID(关联account_auth表)",
                    "type": "integer"
                },
                "job_name": {
                    "description": "Jenkins任务名称",
                    "type": "string"
                }
            }
        },
        "model.CreateKeyManageDto": {
            "type": "object",
            "required": [
                "keyId",
                "keySecret",
                "keyType"
            ],
            "properties": {
                "keyId": {
                    "description": "密钥ID",
                    "type": "string"
                },
                "keySecret": {
                    "description": "密钥Secret",
                    "type": "string"
                },
                "keyType": {
                    "description": "云厂商类型：1=阿里云，2=腾讯云，3=百度云，4=华为云，5=AWS云",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注信息",
                    "type": "string"
                }
            }
        },
        "model.CreateKubeClusterRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "autoDeploy": {
                    "description": "是否自动部署",
                    "type": "boolean"
                },
                "clusterType": {
                    "description": "集群类型:1-自建,2-导入(默认为自建)",
                    "type": "integer"
                },
                "deploymentMode": {
                    "description": "部署模式:1-单Master,2-多Master",
                    "type": "integer"
                },
                "description": {
                    "description": "集群描述",
                    "type": "string"
                },
                "enabledComponents": {
                    "description": "启用组件",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "kubeconfig": {
                    "description": "导入集群参数",
                    "type": "string"
                },
                "name": {
                    "description": "集群名称",
                    "type": "string"
                },
                "nodeConfig": {
                    "description": "节点配置",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.NodeConfig"
                        }
                    ]
                },
                "privateRegistry": {
                    "description": "私有镜像仓库地址（兼容旧版本）",
                    "type": "string"
                },
                "registryConfig": {
                    "description": "镜像仓库配置（新版本）",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RegistryConfig"
                        }
                    ]
                },
                "registryPassword": {
                    "description": "镜像仓库密码（兼容旧版本）",
                    "type": "string"
                },
                "registryUsername": {
                    "description": "镜像仓库用户名（兼容旧版本）",
                    "type": "string"
                },
                "taskDescription": {
                    "description": "任务描述",
                    "type": "string"
                },
                "taskName": {
                    "description": "任务名称",
                    "type": "string"
                },
                "version": {
                    "description": "自建集群参数",
                    "type": "string"
                }
            }
        },
        "model.CreateLimitRangeRequest": {
            "type": "object",
            "required": [
                "name",
                "spec"
            ],
            "properties": {
                "name": {
                    "description": "LimitRange名称",
                    "type": "string"
                },
                "spec": {
                    "description": "LimitRange规格",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.LimitRangeRequestSpec"
                        }
                    ]
                }
            }
        },
        "model.CreateSecretRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
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
        "model.CreateStorageClassRequest": {
            "type": "object",
            "required": [
                "name",
                "provisioner"
            ],
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
        "model.CreateSyncScheduleDto": {
            "type": "object",
            "required": [
                "cronExpr",
                "keyTypes",
                "name"
            ],
            "properties": {
                "cronExpr": {
                    "description": "cron表达式",
                    "type": "string"
                },
                "keyTypes": {
                    "description": "要同步的云厂商类型（JSON数组格式：[1,2,3]）",
                    "type": "string"
                },
                "name": {
                    "description": "配置名称",
                    "type": "string"
                },
                "remark": {
                    "description": "备注信息",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1=启用，0=禁用",
                    "type": "integer"
                }
            }
        },
        "model.DatabaseStats": {
            "type": "object",
            "properties": {
                "byType": {
                    "description": "按类型统计",
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "total": {
                    "description": "数据库总数",
                    "type": "integer"
                }
            }
        },
        "model.DelSysLoginInfoDto": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "Id列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.DelSysPostDto": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "Id列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.EcsAuthIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.EcsAuthVo": {
            "type": "object",
            "properties": {
                "createTime": {
                    "$ref": "#/definitions/util.HTime"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "publicKey": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.EndpointPort": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "端口名称",
                    "type": "string"
                },
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
        "model.EnvVar": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "变量名",
                    "type": "string"
                },
                "value": {
                    "description": "变量值",
                    "type": "string"
                }
            }
        },
        "model.EventInfo": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "firstTime": {
                    "type": "string"
                },
                "lastTime": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.EventListResponse": {
            "type": "object",
            "properties": {
                "events": {
                    "description": "事件列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sEvent"
                    }
                },
                "filter": {
                    "description": "过滤条件",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "namespace": {
                    "description": "命名空间（如果是命名空间级别的查询）",
                    "type": "string"
                },
                "total": {
                    "description": "事件总数",
                    "type": "integer"
                }
            }
        },
        "model.GetBuildLogResponse": {
            "type": "object",
            "properties": {
                "buildNumber": {
                    "description": "构建编号",
                    "type": "integer"
                },
                "hasMore": {
                    "description": "是否有更多日志",
                    "type": "boolean"
                },
                "jobName": {
                    "description": "任务名称",
                    "type": "string"
                },
                "log": {
                    "description": "日志内容",
                    "type": "string"
                },
                "moreData": {
                    "description": "是否有更多数据",
                    "type": "boolean"
                },
                "server": {
                    "description": "服务器名称",
                    "type": "string"
                },
                "textSize": {
                    "description": "文本大小",
                    "type": "integer"
                }
            }
        },
        "model.GetWorkloadYAMLResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "success": {
                    "description": "是否获取成功",
                    "type": "boolean"
                },
                "workloadName": {
                    "description": "工作负载名称",
                    "type": "string"
                },
                "workloadType": {
                    "description": "工作负载类型",
                    "type": "string"
                },
                "yamlContent": {
                    "description": "YAML内容",
                    "type": "string"
                }
            }
        },
        "model.HostStats": {
            "type": "object",
            "properties": {
                "offline": {
                    "description": "离线数量",
                    "type": "integer"
                },
                "online": {
                    "description": "在线数量",
                    "type": "integer"
                },
                "total": {
                    "description": "主机总数",
                    "type": "integer"
                }
            }
        },
        "model.JenkinsBuild": {
            "type": "object",
            "properties": {
                "actions": {
                    "description": "构建操作",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.BuildAction"
                    }
                },
                "building": {
                    "description": "是否正在构建",
                    "type": "boolean"
                },
                "changeSet": {
                    "description": "变更集",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.ChangeSet"
                        }
                    ]
                },
                "culprits": {
                    "description": "责任人",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "description": {
                    "description": "构建描述",
                    "type": "string"
                },
                "displayName": {
                    "description": "显示名称",
                    "type": "string"
                },
                "duration": {
                    "description": "构建时长(毫秒)",
                    "type": "integer"
                },
                "estimatedDuration": {
                    "description": "预计时长(毫秒)",
                    "type": "integer"
                },
                "executor": {
                    "description": "执行器信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.BuildExecutor"
                        }
                    ]
                },
                "fullDisplayName": {
                    "description": "完整显示名称",
                    "type": "string"
                },
                "keepLog": {
                    "description": "是否保留日志",
                    "type": "boolean"
                },
                "number": {
                    "description": "构建编号",
                    "type": "integer"
                },
                "queueId": {
                    "description": "队列ID",
                    "type": "integer"
                },
                "result": {
                    "description": "构建结果 SUCCESS/FAILURE/UNSTABLE/ABORTED",
                    "type": "string"
                },
                "timestamp": {
                    "description": "开始时间戳",
                    "type": "integer"
                },
                "url": {
                    "description": "构建URL",
                    "type": "string"
                }
            }
        },
        "model.JenkinsBuildDetailResponse": {
            "type": "object",
            "properties": {
                "build": {
                    "$ref": "#/definitions/model.JenkinsBuild"
                },
                "log": {
                    "description": "构建日志",
                    "type": "string"
                },
                "server": {
                    "description": "服务器名称",
                    "type": "string"
                }
            }
        },
        "model.JenkinsComputer": {
            "type": "object",
            "properties": {
                "displayName": {
                    "description": "显示名称",
                    "type": "string"
                },
                "executors": {
                    "description": "执行器列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsExecutor"
                    }
                },
                "icon": {
                    "description": "图标",
                    "type": "string"
                },
                "iconClassName": {
                    "description": "图标类名",
                    "type": "string"
                },
                "idle": {
                    "description": "是否空闲",
                    "type": "boolean"
                },
                "jnlpAgent": {
                    "description": "是否JNLP代理",
                    "type": "boolean"
                },
                "launchSupported": {
                    "description": "是否支持启动",
                    "type": "boolean"
                },
                "loadStatistics": {
                    "description": "负载统计",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsLoadStatistics"
                        }
                    ]
                },
                "manualLaunchAllowed": {
                    "description": "是否允许手动启动",
                    "type": "boolean"
                },
                "monitorData": {
                    "description": "监控数据",
                    "type": "object",
                    "additionalProperties": true
                },
                "numExecutors": {
                    "description": "执行器数量",
                    "type": "integer"
                },
                "offline": {
                    "description": "是否离线",
                    "type": "boolean"
                },
                "offlineCause": {
                    "description": "离线原因"
                },
                "oneOffExecutors": {
                    "description": "一次性执行器",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsExecutor"
                    }
                },
                "temporarilyOffline": {
                    "description": "是否临时离线",
                    "type": "boolean"
                }
            }
        },
        "model.JenkinsEnv": {
            "type": "object",
            "properties": {
                "app_id": {
                    "description": "应用ID",
                    "type": "integer"
                },
                "application": {
                    "description": "关联",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Application"
                        }
                    ]
                },
                "created_at": {
                    "type": "string"
                },
                "env_name": {
                    "description": "环境名称",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "jenkins_server_id": {
                    "description": "Jenkins服务器ID(关联account_auth)",
                    "type": "integer"
                },
                "job_name": {
                    "description": "Jenkins任务名称",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.JenkinsExecutor": {
            "type": "object",
            "properties": {
                "currentExecutable": {
                    "description": "当前执行的任务"
                },
                "currentWorkUnit": {
                    "description": "当前工作单元"
                },
                "idle": {
                    "description": "是否空闲",
                    "type": "boolean"
                },
                "likelyStuck": {
                    "description": "是否可能卡住",
                    "type": "boolean"
                },
                "number": {
                    "description": "执行器编号",
                    "type": "integer"
                },
                "progress": {
                    "description": "进度",
                    "type": "integer"
                }
            }
        },
        "model.JenkinsJob": {
            "type": "object",
            "properties": {
                "_class": {
                    "description": "任务类型",
                    "type": "string"
                },
                "actions": {
                    "description": "任务操作",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JobAction"
                    }
                },
                "buildable": {
                    "description": "是否可构建",
                    "type": "boolean"
                },
                "color": {
                    "description": "状态颜色(blue/red/yellow等)",
                    "type": "string"
                },
                "description": {
                    "description": "任务描述",
                    "type": "string"
                },
                "displayName": {
                    "description": "显示名称",
                    "type": "string"
                },
                "lastBuild": {
                    "description": "最后一次构建",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsBuild"
                        }
                    ]
                },
                "lastFailedBuild": {
                    "description": "最后一次失败构建",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsBuild"
                        }
                    ]
                },
                "lastStableBuild": {
                    "description": "最后一次稳定构建",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsBuild"
                        }
                    ]
                },
                "lastSuccessfulBuild": {
                    "description": "最后一次成功构建",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsBuild"
                        }
                    ]
                },
                "name": {
                    "description": "任务名称",
                    "type": "string"
                },
                "property": {
                    "description": "任务属性",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JobProperty"
                    }
                },
                "url": {
                    "description": "任务URL",
                    "type": "string"
                }
            }
        },
        "model.JenkinsJobDetailResponse": {
            "type": "object",
            "properties": {
                "builds": {
                    "description": "构建历史",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsBuild"
                    }
                },
                "job": {
                    "$ref": "#/definitions/model.JenkinsJob"
                },
                "server": {
                    "description": "服务器名称",
                    "type": "string"
                }
            }
        },
        "model.JenkinsJobListResponse": {
            "type": "object",
            "properties": {
                "jobs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsJob"
                    }
                },
                "server": {
                    "description": "服务器名称",
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.JenkinsLabel": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "标签名称",
                    "type": "string"
                }
            }
        },
        "model.JenkinsLoadStatistics": {
            "type": "object",
            "properties": {
                "busyExecutors": {
                    "description": "忙碌执行器数",
                    "type": "integer"
                },
                "idleExecutors": {
                    "description": "空闲执行器数",
                    "type": "integer"
                },
                "queueLength": {
                    "description": "队列长度",
                    "type": "integer"
                },
                "totalExecutors": {
                    "description": "总执行器数",
                    "type": "integer"
                }
            }
        },
        "model.JenkinsQueue": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "队列项目",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsQueueItem"
                    }
                }
            }
        },
        "model.JenkinsQueueItem": {
            "type": "object",
            "properties": {
                "actions": {
                    "description": "操作",
                    "type": "array",
                    "items": {}
                },
                "blocked": {
                    "description": "是否阻塞",
                    "type": "boolean"
                },
                "buildable": {
                    "description": "是否可构建",
                    "type": "boolean"
                },
                "buildableStartMilliseconds": {
                    "description": "可构建开始时间",
                    "type": "integer"
                },
                "id": {
                    "description": "队列项目ID",
                    "type": "integer"
                },
                "inQueueSince": {
                    "description": "入队时间",
                    "type": "integer"
                },
                "params": {
                    "description": "参数",
                    "type": "string"
                },
                "stuck": {
                    "description": "是否卡住",
                    "type": "boolean"
                },
                "task": {
                    "description": "任务信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsTask"
                        }
                    ]
                },
                "url": {
                    "description": "URL",
                    "type": "string"
                },
                "why": {
                    "description": "等待原因",
                    "type": "string"
                }
            }
        },
        "model.JenkinsServerInfo": {
            "type": "object",
            "properties": {
                "alias": {
                    "description": "别名(服务器名称)",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "description": {
                    "description": "描述(备注)",
                    "type": "string"
                },
                "host": {
                    "description": "Jenkins服务器地址",
                    "type": "string"
                },
                "id": {
                    "description": "账号ID",
                    "type": "integer"
                },
                "port": {
                    "description": "端口",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.JenkinsServerListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsServerInfo"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.JenkinsServerOption": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "服务器ID",
                    "type": "integer"
                },
                "name": {
                    "description": "服务器名称(别名)",
                    "type": "string"
                }
            }
        },
        "model.JenkinsSystemInfo": {
            "type": "object",
            "properties": {
                "assignedLabels": {
                    "description": "分配的标签",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsLabel"
                    }
                },
                "computers": {
                    "description": "计算机列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsComputer"
                    }
                },
                "mode": {
                    "description": "运行模式",
                    "type": "string"
                },
                "nodeDescription": {
                    "description": "节点描述",
                    "type": "string"
                },
                "nodeName": {
                    "description": "节点名称",
                    "type": "string"
                },
                "numExecutors": {
                    "description": "执行器数量",
                    "type": "integer"
                },
                "overallLoad": {
                    "description": "总体负载",
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "primaryView": {
                    "description": "主视图",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsView"
                        }
                    ]
                },
                "unlabeledLoad": {
                    "description": "未标记负载",
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                },
                "useCrumbs": {
                    "description": "是否使用CSRF保护",
                    "type": "boolean"
                },
                "useSecurity": {
                    "description": "是否使用安全",
                    "type": "boolean"
                },
                "version": {
                    "description": "Jenkins版本",
                    "type": "string"
                },
                "views": {
                    "description": "视图列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsView"
                    }
                }
            }
        },
        "model.JenkinsView": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "视图描述",
                    "type": "string"
                },
                "jobs": {
                    "description": "视图中的任务",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.JenkinsJob"
                    }
                },
                "name": {
                    "description": "视图名称",
                    "type": "string"
                },
                "url": {
                    "description": "视图URL",
                    "type": "string"
                }
            }
        },
        "model.JobAction": {
            "type": "object",
            "properties": {
                "_class": {
                    "description": "操作类型",
                    "type": "string"
                }
            }
        },
        "model.JobProperty": {
            "type": "object",
            "properties": {
                "_class": {
                    "description": "属性类型",
                    "type": "string"
                }
            }
        },
        "model.KeyManage": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "id": {
                    "type": "integer"
                },
                "keyId": {
                    "description": "密钥ID(加密存储)",
                    "type": "string"
                },
                "keySecret": {
                    "description": "密钥Secret(加密存储)",
                    "type": "string"
                },
                "keyType": {
                    "description": "云厂商类型：1=阿里云，2=腾讯云，3=百度云，4=华为云，5=AWS云",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注信息",
                    "type": "string"
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
        "model.KubeCluster": {
            "type": "object",
            "properties": {
                "clusterType": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "credential": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastSyncAt": {
                    "type": "string"
                },
                "masterNodes": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nodeCount": {
                    "type": "integer"
                },
                "readyNodes": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                },
                "workerNodes": {
                    "type": "integer"
                }
            }
        },
        "model.KubeClusterListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.KubeCluster"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.LimitRangeDetail": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "limits": {
                    "description": "限制项列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LimitRangeItem"
                    }
                },
                "name": {
                    "description": "LimitRange名称",
                    "type": "string"
                }
            }
        },
        "model.LimitRangeItem": {
            "type": "object",
            "properties": {
                "default": {
                    "description": "默认值",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "defaultRequest": {
                    "description": "默认请求值",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "max": {
                    "description": "最大限制",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "maxLimitRequestRatio": {
                    "description": "最大限制与请求比率",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "min": {
                    "description": "最小限制",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "type": {
                    "description": "限制类型 Container/Pod/PersistentVolumeClaim",
                    "type": "string"
                }
            }
        },
        "model.LimitRangeListResponse": {
            "type": "object",
            "properties": {
                "limitRanges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LimitRangeDetail"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.LimitRangeRequestSpec": {
            "type": "object",
            "required": [
                "limits"
            ],
            "properties": {
                "limits": {
                    "description": "限制项列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.LimitRangeItem"
                    }
                }
            }
        },
        "model.LoginDto": {
            "type": "object",
            "required": [
                "idKey",
                "image",
                "password",
                "username"
            ],
            "properties": {
                "idKey": {
                    "description": "uuid",
                    "type": "string"
                },
                "image": {
                    "description": "验证码",
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 4
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.NetworkInfo": {
            "type": "object",
            "properties": {
                "apiServerEndpoint": {
                    "description": "API Server内网端点",
                    "type": "string"
                },
                "dnsService": {
                    "description": "DNS服务",
                    "type": "string"
                },
                "networkPlugin": {
                    "description": "网络插件",
                    "type": "string"
                },
                "podCIDR": {
                    "description": "Pod CIDR",
                    "type": "string"
                },
                "proxyMode": {
                    "description": "服务转发模式",
                    "type": "string"
                },
                "serviceCIDR": {
                    "description": "Service CIDR",
                    "type": "string"
                }
            }
        },
        "model.NetworkMetrics": {
            "type": "object",
            "properties": {
                "inboundTraffic": {
                    "description": "入站流量",
                    "type": "string"
                },
                "outboundTraffic": {
                    "description": "出站流量",
                    "type": "string"
                },
                "packetsIn": {
                    "description": "入站包数",
                    "type": "integer"
                },
                "packetsOut": {
                    "description": "出站包数",
                    "type": "integer"
                }
            }
        },
        "model.QuotaInfo": {
            "type": "object",
            "properties": {
                "hard": {
                    "description": "限制值",
                    "type": "string"
                },
                "used": {
                    "description": "已使用值",
                    "type": "string"
                }
            }
        },
        "model.RemoveLabelRequest": {
            "type": "object",
            "required": [
                "key"
            ],
            "properties": {
                "key": {
                    "type": "string"
                }
            }
        },
        "model.RemoveTaintRequest": {
            "type": "object",
            "required": [
                "key"
            ],
            "properties": {
                "effect": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "model.ReplicaSetInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "name": {
                    "description": "ReplicaSet名称",
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
                "revision": {
                    "description": "版本号",
                    "type": "integer"
                },
                "status": {
                    "description": "状态",
                    "type": "string"
                }
            }
        },
        "model.ReplicasSummary": {
            "type": "object",
            "properties": {
                "available": {
                    "description": "可用副本数",
                    "type": "integer"
                },
                "current": {
                    "description": "当前副本数",
                    "type": "integer"
                },
                "desired": {
                    "description": "期望副本数",
                    "type": "integer"
                },
                "ready": {
                    "description": "就绪副本数",
                    "type": "integer"
                },
                "updated": {
                    "description": "已更新副本数",
                    "type": "integer"
                }
            }
        },
        "model.ResetSysAdminPasswordDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "model.RoleMenu": {
            "type": "object",
            "required": [
                "id",
                "menuIds"
            ],
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "menuIds": {
                    "description": "菜单id列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.RuntimeInfo": {
            "type": "object",
            "properties": {
                "containerRuntimeVersion": {
                    "description": "容器运行时版本",
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
                "operatingSystem": {
                    "description": "操作系统",
                    "type": "string"
                },
                "osImage": {
                    "description": "操作系统镜像",
                    "type": "string"
                }
            }
        },
        "model.RuntimeSummary": {
            "type": "object",
            "properties": {
                "apiServerVersion": {
                    "description": "API Server版本",
                    "type": "string"
                },
                "containerRuntime": {
                    "description": "容器运行时",
                    "type": "string"
                },
                "coreDNSVersion": {
                    "description": "CoreDNS版本",
                    "type": "string"
                },
                "etcdVersion": {
                    "description": "etcd版本",
                    "type": "string"
                },
                "kubeProxyVersion": {
                    "description": "kube-proxy版本",
                    "type": "string"
                },
                "kubernetesVersion": {
                    "description": "Kubernetes版本",
                    "type": "string"
                },
                "upTime": {
                    "description": "集群运行时间",
                    "type": "string"
                }
            }
        },
        "model.ScaleWorkloadRequest": {
            "type": "object",
            "required": [
                "replicas"
            ],
            "properties": {
                "replicas": {
                    "description": "目标副本数",
                    "type": "integer"
                }
            }
        },
        "model.SecretDetail": {
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
                    "description": "Secret名称",
                    "type": "string"
                },
                "namespace": {
                    "description": "命名空间",
                    "type": "string"
                },
                "spec": {
                    "description": "完整规格配置"
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
        "model.SecretListResponse": {
            "type": "object",
            "properties": {
                "secrets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sSecret"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.StartJobRequest": {
            "type": "object",
            "properties": {
                "parameters": {
                    "description": "构建参数",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "reason": {
                    "description": "构建原因",
                    "type": "string"
                }
            }
        },
        "model.StartJobResponse": {
            "type": "object",
            "properties": {
                "buildNumber": {
                    "description": "构建编号(如果已知)",
                    "type": "integer"
                },
                "jobName": {
                    "description": "任务名称",
                    "type": "string"
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "queueId": {
                    "description": "队列ID",
                    "type": "integer"
                },
                "server": {
                    "description": "服务器名称",
                    "type": "string"
                },
                "success": {
                    "description": "是否启动成功",
                    "type": "boolean"
                }
            }
        },
        "model.StopBuildRequest": {
            "type": "object",
            "properties": {
                "reason": {
                    "description": "停止原因",
                    "type": "string"
                }
            }
        },
        "model.StopBuildResponse": {
            "type": "object",
            "properties": {
                "buildNumber": {
                    "description": "构建编号",
                    "type": "integer"
                },
                "jobName": {
                    "description": "任务名称",
                    "type": "string"
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "server": {
                    "description": "服务器名称",
                    "type": "string"
                },
                "success": {
                    "description": "是否停止成功",
                    "type": "boolean"
                }
            }
        },
        "model.StorageClassDetail": {
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
                "spec": {
                    "description": "完整规格配置"
                },
                "volumeBindingMode": {
                    "description": "卷绑定模式",
                    "type": "string"
                }
            }
        },
        "model.StorageClassListResponse": {
            "type": "object",
            "properties": {
                "storageClasses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sStorageClass"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.StorageClassTopology": {
            "type": "object",
            "properties": {
                "matchLabelExpressions": {
                    "description": "匹配标签表达式",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StorageClassTopologyExp"
                    }
                }
            }
        },
        "model.StorageClassTopologyExp": {
            "type": "object",
            "properties": {
                "key": {
                    "description": "键",
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
        "model.StorageMetrics": {
            "type": "object",
            "properties": {
                "boundPVs": {
                    "description": "已绑定PV数",
                    "type": "integer"
                },
                "storageClasses": {
                    "description": "存储类列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "totalPVCs": {
                    "description": "PVC总数",
                    "type": "integer"
                },
                "totalPVs": {
                    "description": "PV总数",
                    "type": "integer"
                }
            }
        },
        "model.SyncSchedule": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "cronExpr": {
                    "description": "cron表达式",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "keyTypes": {
                    "description": "要同步的云厂商类型（JSON数组格式：[1,2,3]）",
                    "type": "string"
                },
                "lastRunTime": {
                    "description": "上次执行时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "name": {
                    "description": "配置名称",
                    "type": "string"
                },
                "nextRunTime": {
                    "description": "下次执行时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "remark": {
                    "description": "备注信息",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1=启用，0=禁用",
                    "type": "integer"
                },
                "syncLog": {
                    "description": "最近一次同步日志",
                    "type": "string"
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
        "model.SysAdminIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysDept": {
            "type": "object",
            "properties": {
                "children": {
                    "description": "子集",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysDept"
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
                "deptName": {
                    "description": "部门名称",
                    "type": "string"
                },
                "deptStatus": {
                    "description": "部门状态（1-\u003e正常 2-\u003e停用）",
                    "type": "integer"
                },
                "deptType": {
                    "description": "部门类型（1-\u003e公司, 2-\u003e中心，3-\u003e部门）",
                    "type": "integer"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "parentId": {
                    "description": "父id",
                    "type": "integer"
                }
            }
        },
        "model.SysDeptIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysLoginInfoIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysMenu": {
            "type": "object",
            "properties": {
                "children": {
                    "description": "子集",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysMenu"
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
                "icon": {
                    "description": "菜单图标",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "menuName": {
                    "description": "菜单名称",
                    "type": "string"
                },
                "menuStatus": {
                    "description": "启用状态；1-\u003e禁用；2-\u003e启用",
                    "type": "integer"
                },
                "menuType": {
                    "description": "菜单类型：1-\u003e目录；2-\u003e菜单；3-\u003e按钮",
                    "type": "integer"
                },
                "parentId": {
                    "description": "父菜单id",
                    "type": "integer"
                },
                "sort": {
                    "description": "排序",
                    "type": "integer"
                },
                "url": {
                    "description": "菜单url",
                    "type": "string"
                },
                "value": {
                    "description": "权限值",
                    "type": "string"
                }
            }
        },
        "model.SysMenuIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysOperationLogIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysPost": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "postCode": {
                    "description": "岗位编码",
                    "type": "string"
                },
                "postName": {
                    "description": "岗位名称",
                    "type": "string"
                },
                "postStatus": {
                    "description": "状态（1-\u003e正常 2-\u003e停用）",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                }
            }
        },
        "model.SysPostIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysRoleIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.TestJenkinsConnectionRequest": {
            "type": "object",
            "required": [
                "password",
                "url",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码或API Token",
                    "type": "string"
                },
                "url": {
                    "description": "Jenkins服务器地址",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.TestJenkinsConnectionResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "错误信息",
                    "type": "string"
                },
                "message": {
                    "description": "响应消息",
                    "type": "string"
                },
                "success": {
                    "description": "是否连接成功",
                    "type": "boolean"
                },
                "systemInfo": {
                    "description": "系统信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.JenkinsSystemInfo"
                        }
                    ]
                }
            }
        },
        "model.Toleration": {
            "type": "object",
            "properties": {
                "effect": {
                    "description": "效果",
                    "type": "string"
                },
                "key": {
                    "description": "键",
                    "type": "string"
                },
                "operator": {
                    "description": "操作符",
                    "type": "string"
                },
                "value": {
                    "description": "值",
                    "type": "string"
                }
            }
        },
        "model.UpdateEcsAuthDto": {
            "type": "object",
            "required": [
                "name",
                "password",
                "port",
                "type",
                "username"
            ],
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "name": {
                    "description": "凭证名称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "port": {
                    "description": "端口号",
                    "type": "integer"
                },
                "publicKey": {
                    "description": "公钥",
                    "type": "string"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                },
                "type": {
                    "description": "认证类型:1-\u003e密码",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.UpdateJenkinsEnvRequest": {
            "type": "object",
            "properties": {
                "env_name": {
                    "description": "环境名称",
                    "type": "string"
                },
                "id": {
                    "description": "环境配置ID(可选，不提供则创建新的)",
                    "type": "integer"
                },
                "jenkins_server_id": {
                    "description": "Jenkins服务器ID(关联account_auth表)",
                    "type": "integer"
                },
                "job_name": {
                    "description": "Jenkins任务名称",
                    "type": "string"
                }
            }
        },
        "model.UpdateKeyManageDto": {
            "type": "object",
            "required": [
                "id",
                "keyId",
                "keySecret",
                "keyType"
            ],
            "properties": {
                "id": {
                    "description": "密钥ID",
                    "type": "integer"
                },
                "keyId": {
                    "description": "密钥ID",
                    "type": "string"
                },
                "keySecret": {
                    "description": "密钥Secret",
                    "type": "string"
                },
                "keyType": {
                    "description": "云厂商类型：1=阿里云，2=腾讯云，3=百度云，4=华为云，5=AWS云",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注信息",
                    "type": "string"
                }
            }
        },
        "model.UpdateKubeClusterRequest": {
            "type": "object",
            "properties": {
                "credential": {
                    "description": "K8s凭证(kubeconfig内容)",
                    "type": "string"
                },
                "description": {
                    "description": "集群描述",
                    "type": "string"
                },
                "name": {
                    "description": "集群名称",
                    "type": "string"
                },
                "version": {
                    "description": "集群版本(可选，同步时会自动更新)",
                    "type": "string"
                }
            }
        },
        "model.UpdateLimitRangeRequest": {
            "type": "object",
            "required": [
                "spec"
            ],
            "properties": {
                "spec": {
                    "description": "LimitRange规格",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.LimitRangeRequestSpec"
                        }
                    ]
                }
            }
        },
        "model.UpdatePersonalDto": {
            "type": "object",
            "required": [
                "email",
                "nickname",
                "note",
                "phone",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "icon": {
                    "description": "头像",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "note": {
                    "description": "备注",
                    "type": "string"
                },
                "phone": {
                    "description": "电话",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.UpdatePersonalPasswordDto": {
            "type": "object",
            "required": [
                "newPassword",
                "password",
                "resetPassword"
            ],
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "newPassword": {
                    "description": "新密码",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "resetPassword": {
                    "description": "重复密码",
                    "type": "string"
                }
            }
        },
        "model.UpdateSecretRequest": {
            "type": "object",
            "properties": {
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
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "stringData": {
                    "description": "字符串数据",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.UpdateStorageClassRequest": {
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
                "parameters": {
                    "description": "参数",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.UpdateSyncScheduleDto": {
            "type": "object",
            "required": [
                "cronExpr",
                "id",
                "keyTypes",
                "name"
            ],
            "properties": {
                "cronExpr": {
                    "description": "cron表达式",
                    "type": "string"
                },
                "id": {
                    "description": "配置ID",
                    "type": "integer"
                },
                "keyTypes": {
                    "description": "要同步的云厂商类型（JSON数组格式：[1,2,3]）",
                    "type": "string"
                },
                "name": {
                    "description": "配置名称",
                    "type": "string"
                },
                "remark": {
                    "description": "备注信息",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1=启用，0=禁用",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysAdminDto": {
            "type": "object",
            "properties": {
                "deptId": {
                    "description": "部门id",
                    "type": "integer"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "note": {
                    "description": "备注",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "postId": {
                    "description": "岗位id",
                    "type": "integer"
                },
                "roleId": {
                    "description": "角色id",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "model.UpdateSysAdminStatusDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysPostStatusDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "postStatus": {
                    "description": "状态（1-\u003e正常 2-\u003e停用）",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysRoleDto": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "Id",
                    "type": "integer"
                },
                "roleKey": {
                    "description": "角色key",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysRoleStatusDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "model.UpdateWorkloadRequest": {
            "type": "object",
            "properties": {
                "labels": {
                    "description": "标签",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "strategy": {
                    "description": "部署策略"
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
        "model.UpdateWorkloadYAMLRequest": {
            "type": "object",
            "required": [
                "workloadName",
                "workloadType",
                "yamlContent"
            ],
            "properties": {
                "dryRun": {
                    "description": "是否只进行校验不实际更新",
                    "type": "boolean"
                },
                "force": {
                    "description": "是否强制更新",
                    "type": "boolean"
                },
                "validateOnly": {
                    "description": "是否只校验YAML格式",
                    "type": "boolean"
                },
                "workloadName": {
                    "description": "工作负载名称",
                    "type": "string"
                },
                "workloadType": {
                    "description": "工作负载类型: deployment,statefulset,daemonset,job,cronjob",
                    "type": "string"
                },
                "yamlContent": {
                    "description": "YAML内容",
                    "type": "string"
                }
            }
        },
        "model.UpdateWorkloadYAMLResponse": {
            "type": "object",
            "properties": {
                "appliedAt": {
                    "description": "应用时间",
                    "type": "string"
                },
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
                    "description": "命名空间",
                    "type": "string"
                },
                "success": {
                    "description": "是否更新成功",
                    "type": "boolean"
                },
                "updateStrategy": {
                    "description": "更新策略 (patch/update/rolling)",
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
                },
                "workloadName": {
                    "description": "工作负载名称",
                    "type": "string"
                },
                "workloadType": {
                    "description": "工作负载类型",
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "absoluteUrl": {
                    "description": "绝对URL",
                    "type": "string"
                },
                "fullName": {
                    "description": "全名",
                    "type": "string"
                }
            }
        },
        "model.ValidateJenkinsJobRequest": {
            "type": "object",
            "required": [
                "jenkins_server_id",
                "job_name"
            ],
            "properties": {
                "jenkins_server_id": {
                    "description": "Jenkins服务器ID",
                    "type": "integer"
                },
                "job_name": {
                    "description": "任务名称",
                    "type": "string"
                }
            }
        },
        "model.ValidateJenkinsJobResponse": {
            "type": "object",
            "properties": {
                "exists": {
                    "description": "任务是否存在",
                    "type": "boolean"
                },
                "job_name": {
                    "description": "任务名称",
                    "type": "string"
                },
                "job_url": {
                    "description": "任务URL(如果存在)",
                    "type": "string"
                },
                "message": {
                    "description": "验证消息",
                    "type": "string"
                },
                "server_id": {
                    "description": "服务器ID",
                    "type": "integer"
                }
            }
        },
        "model.ValidateYAMLRequest": {
            "type": "object",
            "required": [
                "yamlContent"
            ],
            "properties": {
                "resourceType": {
                    "description": "资源类型，如pod、deployment等",
                    "type": "string"
                },
                "yamlContent": {
                    "description": "YAML内容",
                    "type": "string"
                }
            }
        },
        "model.ValidateYAMLResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "description": "错误列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "parsedObject": {
                    "description": "解析后的对象",
                    "type": "object",
                    "additionalProperties": true
                },
                "suggestions": {
                    "description": "建议列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "valid": {
                    "description": "是否有效",
                    "type": "boolean"
                },
                "warnings": {
                    "description": "警告列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.VolumeInfo": {
            "type": "object",
            "properties": {
                "mountPath": {
                    "description": "挂载路径",
                    "type": "string"
                },
                "name": {
                    "description": "卷名称",
                    "type": "string"
                },
                "readOnly": {
                    "description": "只读状态",
                    "type": "boolean"
                },
                "type": {
                    "description": "卷类型",
                    "type": "string"
                }
            }
        },
        "model.VolumeMount": {
            "type": "object",
            "required": [
                "mountPath",
                "name"
            ],
            "properties": {
                "mountPath": {
                    "description": "挂载路径",
                    "type": "string"
                },
                "name": {
                    "description": "卷名称",
                    "type": "string"
                },
                "readOnly": {
                    "description": "只读模式",
                    "type": "boolean"
                }
            }
        },
        "model.VolumeSpec": {
            "type": "object",
            "required": [
                "name",
                "type"
            ],
            "properties": {
                "config": {
                    "description": "卷配置",
                    "type": "object",
                    "additionalProperties": true
                },
                "name": {
                    "description": "卷名称",
                    "type": "string"
                },
                "type": {
                    "description": "卷类型",
                    "type": "string"
                }
            }
        },
        "model.WorkloadCondition": {
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
        "model.WorkloadListResponse": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "workloads": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.K8sWorkload"
                    }
                }
            }
        },
        "model.WorkloadSummary": {
            "type": "object",
            "properties": {
                "runningPods": {
                    "description": "运行中的Pod数",
                    "type": "integer"
                },
                "totalCronJobs": {
                    "description": "CronJob总数",
                    "type": "integer"
                },
                "totalDaemonSets": {
                    "description": "DaemonSet总数",
                    "type": "integer"
                },
                "totalDeployments": {
                    "description": "Deployment总数",
                    "type": "integer"
                },
                "totalJobs": {
                    "description": "Job总数",
                    "type": "integer"
                },
                "totalPods": {
                    "description": "Pod总数",
                    "type": "integer"
                },
                "totalStatefulSets": {
                    "description": "StatefulSet总数",
                    "type": "integer"
                }
            }
        },
        "model.WorkloadType": {
            "type": "string",
            "enum": [
                "Deployment",
                "StatefulSet",
                "DaemonSet",
                "Job",
                "CronJob",
                "Pod"
            ],
            "x-enum-varnames": [
                "WorkloadTypeDeployment",
                "WorkloadTypeStatefulSet",
                "WorkloadTypeDaemonSet",
                "WorkloadTypeJob",
                "WorkloadTypeCronJob",
                "WorkloadTypePod"
            ]
        },
        "result.PageResult": {
            "type": "object",
            "properties": {
                "list": {
                    "description": "数据列表"
                },
                "page": {
                    "description": "当前页码",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页数量",
                    "type": "integer"
                },
                "total": {
                    "description": "总记录数",
                    "type": "integer"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "返回的数据"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "util.HTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "dao.ListResponse": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ConfigAnsible"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        }`
