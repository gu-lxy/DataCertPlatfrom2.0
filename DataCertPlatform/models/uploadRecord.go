package models

import "DataCertPlatform/db_mysql"

/**
 * 上传文件的记录
 */
type UploadRecord struct {
	Id        int
	UserId    int
	FileName  string
	FileSize  int
	FileCert  string
	FileTitle string
	CertTime  int
}

func (u UploadRecord) SaveRecord(){
	db_mysql.Db.Exec("insert into ...")
}
