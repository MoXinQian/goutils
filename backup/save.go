package backup

import (
	"godatabackup/client"
	"godatabackup/entity"
	"godatabackup/global"
	"godatabackup/util"
	"log"
	"strings"
)

// 开始备份任务
func StartBackupTask(confReq entity.Config) {
	oldConf, _ := entity.GetConfigCache()
	conf := &entity.Config{}
	conf.EncryptKey = oldConf.EncryptKey
	if conf.EncryptKey == "" {
		encryptKey, err := util.GenerateEncryptKey()
		if err != nil {
			log.Println("生成Key失败")
			return
		}
		conf.EncryptKey = encryptKey
	}
	global.OSS_UPLOAD_FUNC = confReq.OssUploadFunc
	global.OSS_REMOVE_FUNC = confReq.OssRemoveFunc
	for _, item := range confReq.BackupConfig {
		conf.BackupConfig = append(
			conf.BackupConfig,
			entity.BackupConfig{
				ProjectName: item.ProjectName,
				Command:     item.Command,
				SaveDays:    item.SaveDays,
				SaveDaysS3:  item.SaveDaysS3,
				StartTime:   item.StartTime,
				Period:      item.Period,
				Pwd:         item.Pwd,
				BackupType:  item.BackupType,
				Enabled:     item.Enabled,
			},
		)
	}
	for i := 0; i < len(conf.BackupConfig); i++ {
		if conf.BackupConfig[i].Pwd != "" &&
			(len(oldConf.BackupConfig) == 0 || conf.BackupConfig[i].Pwd != oldConf.BackupConfig[i].Pwd) {
			encryptPwd, err := util.EncryptByEncryptKey(conf.EncryptKey, conf.BackupConfig[i].Pwd)
			if err != nil {
				log.Println("加密失败")
				return
			}
			conf.BackupConfig[i].Pwd = encryptPwd
		}
	}
	// Webhook
	conf.WebhookURL = strings.TrimSpace(conf.WebhookURL)
	conf.WebhookRequestBody = strings.TrimSpace(conf.WebhookRequestBody)
	err := conf.SaveConfig()
	if err == nil {
		if conf.BackupNow {
			client.RunOnce()
		}
		//删除任务
		go client.DeleteOldBackup()
		// 重新进行循环
		client.StopRunLoop()
		go client.RunLoop()
	}
}
