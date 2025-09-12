# Swagger API 文档

## 概述

本项目使用 Swagger 2.0 规范生成 API 文档，支持自动生成和手动维护。

## 文件说明

- `swagger.json` - JSON 格式的 API 文档（推荐用于 Postman 导入）
- `swagger.yaml` - YAML 格式的 API 文档
- `docs.go` - Go 代码格式的文档（用于程序引用）

## 导入到 Postman

### 方法 1: 导入 swagger.json 文件（推荐）

1. 打开 Postman 应用
2. 点击左上角的 "Import" 按钮
3. 选择 "File" 标签页
4. 点击 "Upload Files" 并选择 `docs/swagger.json` 文件
5. 点击 "Import" 完成导入

### 方法 2: 通过 URL 导入

如果项目已部署到服务器，可以通过以下 URL 导入：

```
http://your-domain:port/swagger/doc.json
```

## 主要 API 分类

### 系统管理 API (SysGithub)

- `GET /backend/github/getGithubList` - 获取 GitHub 提交列表
- `GET /backend/github/createGithub` - 创建 GitHub 提交记录

### 用户管理 API

- `POST /backend/base/login` - 用户登录
- `GET /backend/user/getUserInfo` - 获取用户信息
- `POST /backend/user/register` - 用户注册

### 权限管理 API

- `GET /backend/authority/getAuthorityList` - 获取权限列表
- `POST /backend/authority/createAuthority` - 创建权限

### 文章管理 API

- `POST /backend/article/createArticle` - 创建文章
- `GET /backend/article/getArticleList` - 获取文章列表
- `PUT /backend/article/updateArticle/{id}` - 更新文章
- `DELETE /backend/article/deleteArticle/{id}` - 删除文章

### 前台 API

- `GET /api/getArticleList` - 分页获取文章列表
- `GET /api/getArticle/{id}` - 获取文章详情
- `GET /api/getSearchArticle/{name}/{value}` - 搜索文章

### 移动端 API

- `POST /backend/mobile/login` - 移动端用户登录
- `GET /backend/mobile/getUserInfo` - 获取移动端用户信息

## 认证说明

### API Key 认证

大部分 API 需要认证，在请求头中添加：

```
Authorization: Bearer your-api-key-here
```

### 无需认证的 API

- 用户登录
- 用户注册
- 前台文章浏览

## 环境变量配置

在 Postman 中创建环境变量：

```
BASE_URL: http://localhost:8888
API_KEY: your-api-key-here
```

## 更新文档

### 自动生成（推荐）

```bash
# 安装 swag 命令
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init -g cmd/main.go -o docs --parseDependency --parseInternal
```

### 手动维护

如果自动生成有问题，可以手动编辑 `swagger.json` 文件。

## 常见问题

### 1. 导入失败

- 检查文件格式是否正确
- 确保文件编码为 UTF-8
- 尝试使用不同的导入方法

### 2. API 调用失败

- 检查 BASE_URL 环境变量
- 确认服务器是否运行
- 验证认证信息

### 3. 文档不完整

- 检查 API 注释格式
- 重新生成文档
- 手动添加缺失的 API

## 开发规范

### Swagger 注释格式

```go
// @Tags API分类
// @Summary API摘要
// @Description API详细描述
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param 参数名 参数位置 参数类型 是否必须 "参数描述"
// @Success 状态码 {object} 响应类型 "成功描述"
// @Failure 状态码 {object} 响应类型 "失败描述"
// @Router 路由路径 [请求方法]
```

### 示例

```go
// @Tags SysGithub
// @Summary 获取GitHub提交列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param page query int false "页码" default(1)
// @Param pageSize query int false "每页大小" default(10)
// @Success 200 {object} response.Response{data=response.PageResult{list=[]system.SysGithub},msg=string} "获取成功"
// @Failure 400 {object} response.Response{msg=string} "获取失败"
// @Router /github/getGithubList [get]
func (g *SystemGithubApi) GetGithubList(c *fiber.Ctx) error {
    // 实现代码
}
```

## 联系支持

如果在使用过程中遇到问题，请：

1. 检查项目文档
2. 查看服务器日志
3. 联系开发团队


