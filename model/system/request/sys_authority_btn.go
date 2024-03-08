package request

type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menuID" form:"menuID"`
	AuthorityId string `json:"authorityId" form:"authorityId"`
	Selected    []uint `json:"selected" form:"selected"`
}
