package system

import (
	"server-fiber/model/common/request"
	"strconv"
	"strings"

	global "server-fiber/model"
)

// SysAutoCodeHistory 自动迁移代码记录,用于回滚,重放使用
type SysAutoCodeHistory struct {
	global.MODEL
	Package       string `json:"package"`
	TableName     string `json:"tableName"`
	RequestMeta   string `gorm:"type:text" json:"requestMeta,omitempty"`   // 前端传入的结构化信息
	AutoCodePath  string `gorm:"type:text" json:"autoCodePath,omitempty"`  // 其他meta信息 path;path
	InjectionMeta string `gorm:"type:text" json:"injectionMeta,omitempty"` // 注入的内容 RouterPath@functionName@RouterString;
	StructName    string `json:"structName"`
	StructCNName  string `json:"structCNName"`
	ApiIDs        string `json:"apiIDs,omitempty"` // api表注册内容
	Flag          int    `json:"flag"`             // 表示对应状态 0 代表创建, 1 代表回滚 ...
}

// ToRequestIds ApiIDs 转换 request.IdsReq
// Author wuhao
func (m *SysAutoCodeHistory) ToRequestIds() request.IdsReq {
	if m.ApiIDs == "" {
		return request.IdsReq{}
	}
	slice := strings.Split(m.ApiIDs, ";")
	ids := make([]int, 0, len(slice))
	length := len(slice)
	for i := range length {
		id, _ := strconv.ParseInt(slice[i], 10, 32)
		ids = append(ids, int(id))
	}
	return request.IdsReq{Ids: ids}
}
