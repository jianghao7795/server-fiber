# 测试说明

## 测试结构

```
tests/
├── api/                    # API 层测试
│   └── system/            # 系统模块 API 测试
├── integration/            # 集成测试
│   └── github_integration_test.go
├── service/                # 服务层测试
│   └── system/            # 系统模块服务测试
└── README.md              # 本文件
```

## 运行测试

### 运行所有测试
```bash
go test ./tests/... -v
```

### 运行短测试（跳过集成测试）
```bash
go test ./tests/... -v -short
```

### 运行特定测试文件
```bash
go test ./tests/integration/github_integration_test.go -v
```

### 运行特定测试函数
```bash
go test ./tests/integration/... -v -run TestApplicationStartup
```

## 测试类型

### 1. 单元测试 (Unit Tests)
- 位置：`tests/api/` 和 `tests/service/`
- 特点：测试单个函数或方法，使用 mock 对象
- 运行：快速，不需要外部依赖

### 2. 集成测试 (Integration Tests)
- 位置：`tests/integration/`
- 特点：测试多个组件之间的交互
- 注意：某些测试需要完整的应用环境（数据库、配置等）

### 3. 服务测试 (Service Tests)
- 位置：`tests/service/`
- 特点：测试业务逻辑层
- 使用：mock 数据库和其他外部依赖

## 测试环境要求

### 开发环境
- Go 1.22+
- 测试依赖：`github.com/stretchr/testify`

### CI/CD 环境
- 测试数据库（MySQL 或 PostgreSQL）
- 完整的配置文件
- 网络连接（用于 GitHub API 测试）

## 跳过测试

### 使用 -short 标志
```bash
go test -short ./tests/...
```

### 在代码中使用 t.Skip()
```go
if testing.Short() {
    t.Skip("跳过集成测试")
}
```

## 测试最佳实践

1. **命名规范**：测试函数以 `Test` 开头
2. **错误处理**：使用 `assert` 包进行断言
3. **Mock 使用**：外部依赖使用 mock 对象
4. **测试隔离**：每个测试应该是独立的
5. **清理资源**：测试完成后清理资源

## 常见问题

### 数据库连接错误
- 原因：测试环境没有初始化数据库
- 解决：使用 mock 或跳过需要数据库的测试

### 导入错误
- 原因：模块路径不正确
- 解决：检查 `go.mod` 文件中的模块名

### 网络依赖错误
- 原因：测试需要外部网络连接
- 解决：使用 mock 或跳过网络相关测试

## 添加新测试

1. 在相应的目录创建测试文件
2. 导入必要的包和依赖
3. 编写测试函数
4. 运行测试验证
5. 更新本文档

## 持续集成

在 CI/CD 流程中：
1. 运行单元测试：`go test -short ./tests/...`
2. 运行集成测试：`go test ./tests/...`
3. 生成测试报告
4. 设置测试覆盖率阈值
