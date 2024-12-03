package frontend

import (
	global "server-fiber/model"
	"server-fiber/model/frontend"
)

type Images struct{}

func (s *Images) GetImagesList() (list []frontend.ExaFile, err error) {
	err = global.DB.Model(&frontend.ExaFile{}).Where("proportion > 1.6").Order("id desc").Find(&list).Error
	return
}
