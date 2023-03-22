package global

import "godatabackup/entity"

//@Author: morris
//@Date: 2023/3/21
//@Desc: <Your code description>
var (
	OSS_UPLOAD_FUNC = func(config entity.OssConfig, backupConfig entity.BackupConfig) bool { return true }
	OSS_REMOVE_FUNC = func(config entity.OssConfig, backupConfig entity.BackupConfig) bool { return true }
)
