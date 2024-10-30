package example

import (
	"server-fiber/global"
)

type ExcelInfo struct {
	FileName string             `json:"fileName" form:"fileName"` // 文件名
	InfoList []FielUploadImport `json:"infoList" form:"infoList"`
}

type FielUploadImport struct {
	global.MODEL
	FileName    string `json:"filename" form:"filename" gorm:"column:filename;"`             // 文件名
	FileNameMd5 string `json:"filename_md5" form:"filename_md5" gorm:"column:filename_md5;"` // 文件名md5
	State       int    `json:"state" form:"state" gorm:"column:state;"`                      // 状态
	FileSize    int64  `json:"file_size" form:"file_size" gorm:"column:file_size;"`          // 文件大小
	FilePath    string `json:"file_path" form:"file_path" gorm:"column:file_path;"`          // 文件路径
	FileType    string `json:"file_type" form:"file_type" gorm:"column:file_type;"`          // 文件类型
}

func (FielUploadImport) TableName() string {
	return "file_upload_import"
}
