package example

import (
	"server-fiber/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.MODEL
	FileName     string         `json:"file_name" form:"file_name" gorm:"comment:文件名"`
	FileMd5      string         `json:"file_md5" form:"file_md5" gorm:"comment:文件的md5"`
	FilePath     string         `json:"file_path" form:"file_path" gorm:"comment:文件路径"`
	ExaFileChunk []ExaFileChunk `json:"exa_file_chunk" form:"exa_file_chunk"` // 切片结构体
	ChunkTotal   int            `json:"chunk_total" form:"chunk_total" gorm:"comment:切片总数"`
	IsFinish     bool           `json:"is_finish" form:"is_finish" gorm:"comment:是否完成"`
}

func (ExaFile) TableName() string {
	return "exa_files"
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}

func (ExaFileChunk) TableName() string {
	return "exa_file_chunks"
}
