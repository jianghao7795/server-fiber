package system

import (
	"errors"
	"fmt"
	"path/filepath"
	"server-fiber/model/common/request"
	"server-fiber/model/system"
	systemReq "server-fiber/model/system/request"
	"server-fiber/model/system/response"
	"server-fiber/utils"
	"strings"
	"time"

	global "server-fiber/model"

	"go.uber.org/zap"
)

var ErrRepeat = errors.New("重复创建")

// CreateAutoCodeHistory 创建代码生成器历史记录
// RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
func (autoCodeHistoryService *AutoCodeHistoryService) CreateAutoCodeHistory(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string, Package string) error {
	return global.DB.Create(&system.SysAutoCodeHistory{
		Package:       Package,
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
	}).Error
}

// First 根据id获取代码生成器历史的数据
func (autoCodeHistoryService *AutoCodeHistoryService) First(info *request.GetById) (string, error) {
	var meta string
	return meta, global.DB.Model(system.SysAutoCodeHistory{}).Select("request_meta").Where("id = ?", info.Uint()).First(&meta).Error
}

// Repeat 检测重复
func (autoCodeHistoryService *AutoCodeHistoryService) Repeat(structName string, Package string) bool {
	var count int64
	global.DB.Model(&system.SysAutoCodeHistory{}).Where("struct_name = ? and package = ? and flag = 0", structName, Package).Count(&count)
	return count > 0
}

// RollBack 回滚
func (autoCodeHistoryService *AutoCodeHistoryService) RollBack(info *systemReq.RollBack) error {
	md := system.SysAutoCodeHistory{}
	if err := global.DB.Where("id = ?", info.ID).First(&md).Error; err != nil {
		return err
	}
	// 清除API表
	err := ApiServiceApp.DeleteApiByIds(strings.Split(md.ApiIDs, ";"))
	if err != nil {
		global.LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}
	// 删除表
	if info.DeleteTable {
		if err = AutoCodeServiceApp.DropTable(md.TableName); err != nil {
			global.LOG.Error("ClearTag DropTable:", zap.Error(err))
		}
	}
	// 删除文件

	for _, path := range strings.Split(md.AutoCodePath, ";") {
		// 增加安全判断补丁:
		_path, err := filepath.Abs(path)
		if err != nil || _path != path {
			continue
		}
		// 迁移
		nPath := filepath.Join(global.CONFIG.AutoCode.Root,
			"rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		// 判断目标文件是否存在
		for utils.FileExist(nPath) {
			nPath += fmt.Sprintf("_%d", time.Now().Nanosecond())
		}
		err = utils.FileMove(path, nPath)
		if err != nil {
			global.LOG.Error("文件迁移失败: ", zap.Error(err))
		}
	}
	// 清除注入
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.AutoClearCode(meta[0], meta[2])
		}
	}
	md.Flag = 1
	return global.DB.Save(&md).Error
}

// Delete 删除历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) Delete(info *request.GetById) error {
	return global.DB.Where("id = ?", info.Uint()).Delete(&system.SysAutoCodeHistory{}).Error
}

// GetList 获取系统历史数据
func (autoCodeHistoryService *AutoCodeHistoryService) GetList(info request.PageInfo) (list []response.AutoCodeHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.SysAutoCodeHistory{})
	var entities []response.AutoCodeHistory
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}
