package system

type ApiService struct{}

var ApiServiceApp = new(ApiService)

type JwtService struct{}

var JwtServiceApp = new(JwtService)

type AuthorityBtnService struct{}

var AuthorityBtnServiceApp = new(AuthorityBtnService)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

type AutoCodeService struct{}

var AutoCodeServiceApp = new(AutoCodeService)

type autoCodeMysql struct{}

var AutoCodeMysql = new(autoCodeMysql)

type autoCodePgsql struct{}

var AutoCodePgsql = new(autoCodePgsql)

type AutoCodeHistoryService struct{}

var AutoCodeHistoryServiceApp = new(AutoCodeHistoryService)

type BaseMenuService struct{}

var BaseMenuServiceApp = new(BaseMenuService)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

type DictionaryDetailService struct{}

var DictionaryDetailServiceApp = new(DictionaryDetailService)

type DictionaryService struct{}

var DictionaryServiceApp = new(DictionaryService)

type GithubService struct{}

var GithubServiceApp = new(GithubService)

type MysqlInitHandler struct{}
type PgsqlInitHandler struct{}
type MenuService struct{}

var MenuServiceApp = new(MenuService)

type OperationRecordService struct{}

var OperationRecordServiceApp = new(OperationRecordService)

type SystemConfigService struct{}

var SystemConfigServiceApp = new(SystemConfigService)

type Problem struct{}

var ProblemApp = new(Problem)

type UserService struct{}

var UserServiceApp = new(UserService)
