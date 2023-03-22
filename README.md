# 数据备份组件

该组件可以备份文件以及数据库。

## 如何使用

引入包
```bash
go get -u 192.168.90.202/ec/deepssd-godatabackup@v1.0.0
```

代码示例：
```go
func main() {
	config := entity.Config{
		//webhook 为post 请求
		Webhook: entity.Webhook{
			WebhookURL:         "localhost:8080/backup/webhook",
			WebhookRequestBody: "",
		},
		//备份配置
		BackupConfig: []entity.BackupConfig{
			{
				ProjectName: "test",
				//备份执行命令
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
		//对象存储配置
		OssConfig: entity.OssConfig{
			Path:      "your_path",
			Bucket:    "your_bucket",
			EndPoint:  "your_endpoint",
			AccessKey: "your_accessKey",
			SecretKey: "your_secretKey",
		},
		//删除对象存储回调
		OssRemoveFunc: func(config entity.OssConfig, backupConfig entity.BackupConfig) bool {
			fmt.Println("todo 删除对象存储数据，执行结果")
			return false
		},
		//上传对象存储回调
		OssUploadFunc: func(config entity.OssConfig, backupConfig entity.BackupConfig) bool {
			fmt.Println("todo 上传对象存储")
			return false
		},
	}
    //新开协程执行
	go backup.StartBackupTask(config)
}
```

## 数据库备份

- postgres

  | 说明     | 备份脚本                                                     |
  | -------- | ------------------------------------------------------------ |
  | 备份单个 | PGPASSWORD="#{PWD}" pg_dump --host 192.168.1.11 --port 5432 --dbname db-name --user postgres --clean --create --file #{DATE}.sql |
  | 备份全部 | PGPASSWORD="#{PWD}" pg_dumpall --host 192.168.1.11 --port 5432 --user postgres --clean --file #{DATE}.sql |
  | 还原     | psql -U postgres -f 2021-11-12_10_29.sql                     |
  | 还原指定 | psql -U postgres -d dn-name -f 2021-11-12_10_29.sql          |

- mysql/mariadb

  | 说明     | 备份脚本                                                     |
  | -------- | ------------------------------------------------------------ |
  | 备份单个 | mysqldump -h192.168.1.11 -uroot -p#{PWD} db-name > #{DATE}.sql |
  | 备份全部 | mysqldump -h192.168.1.11 -uroot -p#{PWD} --all-databases > #{DATE}.sql |
  | 还原     | mysql -uroot -p123456 db-name <2021-11-12_10_29.sql          |

- 变量说明

    | 变量名  | 说明           |
    | ------- | -------------- |
    | #{DATE} | 年-月-日_时_分 |
    | #{PWD}  | 下方的密码变量 |
