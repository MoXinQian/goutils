package entity

//@Author: morris
//@Date: 2023/3/20
//@Desc: 对象存储配置
type OssConfig struct {
	Path      string `yaml:"path"`
	Bucket    string `yaml:"bucket"`
	EndPoint  string `yaml:"endPoint"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}
