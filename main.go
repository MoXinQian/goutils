package main

import (
	"fmt"
	"godatabackup/backup"
	"godatabackup/entity"
)

func main() {
	config := entity.Config{
		//webhook 为post 请求
		Webhook: entity.Webhook{
			WebhookURL:         "localhost:8080/backup/webhook",
			WebhookRequestBody: "",
		},
		BackupConfig: []entity.BackupConfig{
			{
				ProjectName: "test",
				Command:     "copy D:\\k8s_install.7z k8s_install.7z-#{DATE}.7z",
				SaveDays:    7,
				SaveDaysS3:  19,
				StartTime:   14,
				Period:      1440, //原则上一天备份一次
				Pwd:         "123",
				BackupType:  1, //0 备份文件 1备份数据库
				Enabled:     0,
			},
		},
		OssConfig: entity.OssConfig{
			Path:      "your_path",
			Bucket:    "your_bucket",
			EndPoint:  "your_endpoint",
			AccessKey: "your_accessKey",
			SecretKey: "your_secretKey",
		},
		OssRemoveFunc: func(config entity.OssConfig, backupConfig entity.BackupConfig) bool {
			fmt.Println("todo 删除对象存储数据，执行结果")
			return false
		},
		OssUploadFunc: func(config entity.OssConfig, backupConfig entity.BackupConfig) bool {
			fmt.Println("todo 上传对象存储")
			return false
		},
	}
	backup.StartBackupTask(config)
}
