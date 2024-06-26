package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page       int    `json:"page" form:"page"`             // 页码
	PageSize   int    `json:"pageSize" form:"pageSize"`     // 每页大小
	Keyword    string `json:"keyword" form:"keyword"`       //关键字
	IsCropper  string `json:"is_cropper" form:"is_cropper"` //是否为截图 1, 或 2
	Proportion string `json:"proportion" form:"proportion"`
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
