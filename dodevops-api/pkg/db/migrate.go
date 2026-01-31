// pkg/db/migrate.go
package db

import (
	appmodel "dodevops-api/api/app/model"
	cmdbmodel "dodevops-api/api/cmdb/model"
	ccmodel "dodevops-api/api/configcenter/model"
	k8smodel "dodevops-api/api/k8s/model"
	monitormodel "dodevops-api/api/monitor/model"
	systemmodel "dodevops-api/api/system/model"
	taskmodel "dodevops-api/api/task/model"
	toolmodel "dodevops-api/api/tool/model"

	"gorm.io/gorm"
)

// 注册所有需要自动建表的 model
var models = []interface{}{
	&cmdbmodel.CmdbGroup{},
	&ccmodel.EcsAuth{},
	&ccmodel.KeyManage{},
	&ccmodel.SyncSchedule{},
	&cmdbmodel.CmdbHost{},
	&cmdbmodel.CmdbSQLRecord{},
	&cmdbmodel.CmdbSQL{},
	&ccmodel.AccountAuth{},
	&taskmodel.TaskTemplate{},
	&taskmodel.Task{},
	&taskmodel.TaskWork{},
	&taskmodel.TaskAnsible{},
	&taskmodel.TaskAnsibleWork{},
	&taskmodel.TaskAnsibleHistory{},
	&taskmodel.TaskAnsibleworkHistory{},
	&taskmodel.ConfigAnsible{},
	&monitormodel.Agent{},
	&k8smodel.KubeCluster{},
	&appmodel.Application{},
	&appmodel.JenkinsEnv{},
	&appmodel.QuickDeployment{},
	&appmodel.QuickDeploymentTask{},
	&systemmodel.SysOperationLog{},
	&toolmodel.Tool{},
	&toolmodel.ServiceDeploy{},
	// 可以继续添加其他模型...
}

// 自动迁移所有模型
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(models...)
}
