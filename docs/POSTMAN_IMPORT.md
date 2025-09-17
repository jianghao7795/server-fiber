# Postman 导入说明

## 导入 Swagger 文档到 Postman

### 方法 1: 直接导入 swagger.json 文件

1. 打开 Postman 应用
2. 点击左上角的 "Import" 按钮
3. 选择 "File" 标签页
4. 点击 "Upload Files" 并选择项目中的 `docs/swagger.json` 文件
5. 点击 "Import" 完成导入

### 方法 2: 导入 swagger.yaml 文件

1. 打开 Postman 应用
2. 点击左上角的 "Import" 按钮
3. 选择 "File" 标签页
4. 点击 "Upload Files" 并选择项目中的 `docs/swagger.yaml` 文件
5. 点击 "Import" 完成导入

### 方法 3: 通过 URL 导入

如果您的项目已经部署到服务器上，可以通过以下 URL 导入：

```
http://your-domain:port/swagger/doc.json
```

## 导入后的配置

### 环境变量设置

在 Postman 中创建环境变量：

1. 点击右上角的齿轮图标
2. 选择 "Add" 创建新环境
3. 添加以下变量：

```
BASE_URL: http://localhost:8888
API_KEY: your-api-key-here
```

### 认证设置

对于需要认证的 API，在请求头中添加：

```
Authorization: Bearer {{API_KEY}}
```

## 主要 API 端点

### 系统管理 API

#### GitHub 管理

- `GET /backend/github/getGithubList` - 获取 GitHub 提交列表
- `GET /backend/github/createGithub` - 创建 GitHub 提交记录

#### 用户管理

- `POST /backend/base/login` - 用户登录
- `GET /backend/user/getUserInfo` - 获取用户信息
- `POST /backend/user/register` - 用户注册

#### 权限管理

- `GET /backend/authority/getAuthorityList` - 获取权限列表
- `POST /backend/authority/createAuthority` - 创建权限

### 应用 API

#### 文章管理

- `POST /backend/article/createArticle` - 创建文章
- `GET /backend/article/getArticleList` - 获取文章列表
- `PUT /backend/article/updateArticle/{id}` - 更新文章
- `DELETE /backend/article/deleteArticle/{id}` - 删除文章

#### 标签管理

- `POST /backend/tag/createTag` - 创建标签
- `GET /backend/tag/getTagList` - 获取标签列表

#### 评论管理

- `POST /backend/comment/createComment` - 创建评论
- `GET /backend/comment/getCommentList` - 获取评论列表

### 前台 API

#### 文章浏览

- `GET /api/getArticleList` - 分页获取文章列表
- `GET /api/getArticle/{id}` - 获取文章详情
- `GET /api/getSearchArticle/{name}/{value}` - 搜索文章

#### 用户功能

- `POST /api/login` - 前台用户登录
- `POST /api/register` - 前台用户注册

### 移动端 API

#### 用户管理

- `POST /backend/mobile/login` - 移动端用户登录
- `GET /backend/mobile/getUserInfo` - 获取移动端用户信息
- `PUT /backend/mobile/updateMobileUser` - 更新移动端用户信息

## 测试建议

### 1. 基础功能测试

- 先测试不需要认证的 API（如登录、注册）
- 获取认证 token 后测试需要认证的 API

### 2. 数据流程测试

- 创建 → 查询 → 更新 → 删除 的完整流程
- 测试分页、搜索、筛选功能

### 3. 错误处理测试

- 测试无效参数
- 测试未授权访问
- 测试资源不存在的情况

## 常见问题

### 1. 导入失败

- 检查 swagger.json 文件是否完整
- 确保文件格式正确（UTF-8 编码）
- 尝试使用 swagger.yaml 文件

### 2. API 调用失败

- 检查 BASE_URL 环境变量是否正确
- 确认服务器是否正在运行
- 检查认证 token 是否有效

### 3. 响应格式问题

- 检查请求头中的 Content-Type
- 确认请求体格式是否正确
- 查看服务器日志获取详细错误信息

## 更新 Swagger 文档

当 API 发生变化时，需要重新生成 Swagger 文档：

```bash
# 安装 swag 命令
go install github.com/swaggo/swag/cmd/swag@latest

# 生成文档
swag init -g cmd/main.go -o docs --parseDependency --parseInternal

# 重新导入到 Postman
```

## 联系支持

如果在导入或使用过程中遇到问题，请：

1. 检查项目文档
2. 查看服务器日志
3. 联系开发团队



