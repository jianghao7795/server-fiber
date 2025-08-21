#!/bin/bash

# 测试运行脚本
echo "🚀 开始运行测试..."

# 设置环境变量
export GO_ENV=test

# 运行所有测试
echo "📋 运行单元测试..."
go test ./... -v -short

echo ""
echo "🔗 运行集成测试..."
go test ./... -v -run Integration

echo ""
echo "⚡ 运行性能测试..."
go test ./... -v -bench=.

echo ""
echo "📊 生成测试覆盖率报告..."
go test ./... -v -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

echo ""
echo "✅ 测试完成！"
echo "📈 覆盖率报告已生成: coverage.html"
echo "📄 覆盖率数据: coverage.out"
