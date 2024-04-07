package system

import systemServer "server-fiber/service/system"

// service 端的服务

var apiService = new(systemServer.ApiService)
var authorityBtnService = new(systemServer.AuthorityBtnService)
var authorityService = new(systemServer.AuthorityService)
var menuService = new(systemServer.MenuService)
var casbinService = new(systemServer.CasbinService)
var autoCodeHistoryService = new(systemServer.AutoCodeHistoryService)
var autoCodeService = new(systemServer.AutoCodeService)
var dictionaryDetailService = new(systemServer.DictionaryDetailService)
var dictionaryService = new(systemServer.DictionaryService)
var githubService = new(systemServer.GithubService)
var initDBService = new(systemServer.InitDBService)
var jwtService = new(systemServer.JwtService)
var baseMenuService = new(systemServer.BaseMenuService)
var operationRecordService = new(systemServer.OperationRecordService)
var systemConfigService = new(systemServer.SystemConfigService)
var userProblem = new(systemServer.Problem)
var userService = new(systemServer.UserService)
