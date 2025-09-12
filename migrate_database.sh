#!/bin/bash

# 数据库迁移工具脚本
# 用于安全地更新comments表结构

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

# 数据库配置
DB_HOST=${DB_HOST:-"localhost"}
DB_PORT=${DB_PORT:-"3306"}
DB_NAME=${DB_NAME:-"server_fiber"}
DB_USER=${DB_USER:-"root"}
DB_PASSWORD=${DB_PASSWORD:-""}

echo -e "${BLUE}==========================================${NC}"
echo -e "${BLUE}数据库迁移工具 - 更新comments表结构${NC}"
echo -e "${BLUE}==========================================${NC}"

# 检查MySQL客户端
if ! command -v mysql &> /dev/null; then
    echo -e "${RED}✗ MySQL客户端未安装${NC}"
    exit 1
fi

# 检查数据库连接
echo -e "${YELLOW}1. 检查数据库连接...${NC}"
mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" -e "USE $DB_NAME; SELECT 1;" 2>/dev/null
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 数据库连接成功${NC}"
else
    echo -e "${RED}✗ 数据库连接失败${NC}"
    echo -e "${YELLOW}请检查数据库配置：${NC}"
    echo "  DB_HOST: $DB_HOST"
    echo "  DB_PORT: $DB_PORT"
    echo "  DB_NAME: $DB_NAME"
    echo "  DB_USER: $DB_USER"
    exit 1
fi

# 备份数据库
echo -e "${YELLOW}2. 备份数据库...${NC}"
BACKUP_FILE="backup_$(date +%Y%m%d_%H%M%S).sql"
mysqldump -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" > "$BACKUP_FILE"
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 数据库备份完成: $BACKUP_FILE${NC}"
else
    echo -e "${RED}✗ 数据库备份失败${NC}"
    exit 1
fi

# 检查comments表是否存在
echo -e "${YELLOW}3. 检查comments表...${NC}"
TABLE_EXISTS=$(mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" -e "USE $DB_NAME; SHOW TABLES LIKE 'comments';" 2>/dev/null | wc -l)
if [ "$TABLE_EXISTS" -gt 1 ]; then
    echo -e "${GREEN}✓ comments表存在${NC}"
    
    # 显示当前表结构
    echo -e "${YELLOW}当前comments表结构：${NC}"
    mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" -e "USE $DB_NAME; DESCRIBE comments;"
    
    # 询问是否继续
    echo -e "${YELLOW}是否继续更新comments表结构？(y/N)${NC}"
    read -r response
    if [[ ! "$response" =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}操作已取消${NC}"
        exit 0
    fi
else
    echo -e "${YELLOW}comments表不存在，将创建新表${NC}"
fi

# 执行迁移
echo -e "${YELLOW}4. 执行数据库迁移...${NC}"
mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < sql/safe_migration.sql
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 数据库迁移成功${NC}"
else
    echo -e "${RED}✗ 数据库迁移失败${NC}"
    echo -e "${YELLOW}正在恢复备份...${NC}"
    mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" "$DB_NAME" < "$BACKUP_FILE"
    echo -e "${YELLOW}备份已恢复${NC}"
    exit 1
fi

# 验证迁移结果
echo -e "${YELLOW}5. 验证迁移结果...${NC}"
echo -e "${YELLOW}更新后的comments表结构：${NC}"
mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" -e "USE $DB_NAME; DESCRIBE comments;"

echo -e "${YELLOW}comments表索引：${NC}"
mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASSWORD" -e "USE $DB_NAME; SHOW INDEX FROM comments;"

# 清理备份文件（可选）
echo -e "${YELLOW}6. 清理备份文件...${NC}"
echo -e "${YELLOW}是否删除备份文件 $BACKUP_FILE？(y/N)${NC}"
read -r response
if [[ "$response" =~ ^[Yy]$ ]]; then
    rm "$BACKUP_FILE"
    echo -e "${GREEN}✓ 备份文件已删除${NC}"
else
    echo -e "${YELLOW}备份文件保留: $BACKUP_FILE${NC}"
fi

echo -e "${GREEN}==========================================${NC}"
echo -e "${GREEN}数据库迁移完成！${NC}"
echo -e "${GREEN}==========================================${NC}"
