package utils

// 验证规则
var (
	IdVerify                    = Rules{"ID": {NotEmpty()}}
	ApiVerify                   = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify                  = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify              = Rules{"Title": {NotEmpty()}}
	LoginVerify                 = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty(), Eq("6")}, "Username": {NotEmpty(), Ge("5")}, "Password": {NotEmpty(), Ge("6")}}
	RegisterVerify              = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify              = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify              = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify              = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify           = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify             = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify           = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify          = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify        = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify      = Rules{"AuthorityId": {NotEmpty()}}
	LoginVerifyFrontend         = Rules{"UserName": {NotEmpty()}, "Password": {NotEmpty()}}                                           // 前台页面登录验证
	RegisterVerifyFrontend      = Rules{"Name": {NotEmpty()}, "Password": {NotEmpty(), Ge("6")}, "RePassword": {NotEmpty(), Ge("6")}} // 前台页面登录验证
	ResetPasswordVerifyFrontend = Rules{"Password": {NotEmpty(), Ge("6")}, "NewPassword": {NotEmpty(), Ge("6")}, "RepeatNewPassword": {NotEmpty(), Ge("6")}}
	UpdateUserVerify            = Rules{"Name": {NotEmpty(), Ge("2")}, "HeadImg": {NotEmpty()}, "Header": {NotEmpty()}}
	MobileLoginVerify           = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty(), Ge("6")}}
	MobileUpdatePasswordVerify  = Rules{"ID": {NotEmpty()}, "Password": {NotEmpty(), Ge("6")}, "NewPassword": {NotEmpty(), Ge("6")}}
	TokenLoginVerify            = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty(), Ge("6")}}
)
