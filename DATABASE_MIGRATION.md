# 数据库迁移指南

## 概述

本文档说明如何安全地更新数据库表结构，特别是 comments 表，以支持新的文章点赞功能。

## 迁移文件说明

### 1. SQL 迁移文件

- **`sql/create_like_tables.sql`** - 创建所有新表的 SQL 脚本
- **`sql/update_comments_table.sql`** - 更新 comments 表结构的 SQL 脚本
- **`sql/safe_migration.sql`** - 安全的数据库迁移脚本（推荐使用）

### 2. 迁移工具

- **`migrate_database.sh`** - 自动化的数据库迁移工具
- **`verify_database.sh`** - 数据库结构验证工具

## 迁移步骤

### 方法 1：使用自动化工具（推荐）

1. **配置数据库连接**：

   ```bash
   export DB_HOST="localhost"
   export DB_PORT="3306"
   export DB_NAME="server_fiber"
   export DB_USER="root"
   export DB_PASSWORD="your_password"
   ```

2. **运行迁移工具**：

   ```bash
   ./migrate_database.sh
   ```

3. **验证迁移结果**：
   ```bash
   ./verify_database.sh
   ```

### 方法 2：手动执行 SQL

1. **备份数据库**：

   ```bash
   mysqldump -u root -p server_fiber > backup_$(date +%Y%m%d_%H%M%S).sql
   ```

2. **执行迁移脚本**：

   ```bash
   mysql -u root -p server_fiber < sql/safe_migration.sql
   ```

3. **验证结果**：
   ```bash
   mysql -u root -p server_fiber -e "DESCRIBE comments;"
   ```

## 表结构变更说明

### comments 表结构更新

| 字段名     | 原类型  | 新类型                  | 说明                   |
| ---------- | ------- | ----------------------- | ---------------------- |
| article_id | int     | post_id bigint unsigned | 重命名并更新类型       |
| parent_id  | int     | bigint unsigned         | 更新类型，支持楼中楼   |
| user_id    | int     | bigint unsigned         | 更新类型               |
| to_user_id | int     | bigint unsigned         | 更新类型，支持回复用户 |
| content    | varchar | text                    | 更新类型，支持长文本   |

### 新增索引

- `idx_comments_post_id` - post_id 字段索引
- `idx_comments_parent_id` - parent_id 字段索引
- `idx_comments_user_id` - user_id 字段索引
- `idx_comments_to_user_id` - to_user_id 字段索引

## 安全注意事项

1. **备份数据**：迁移前务必备份数据库
2. **测试环境**：先在测试环境验证迁移脚本
3. **事务安全**：使用事务确保数据一致性
4. **回滚计划**：准备回滚方案

## 故障排除

### 常见问题

1. **字段已存在**：

   - 脚本会自动检查并跳过已存在的字段
   - 不会覆盖现有数据

2. **索引冲突**：

   - 脚本会检查索引是否存在
   - 避免重复创建索引

3. **数据类型不匹配**：
   - 脚本会安全地转换数据类型
   - 确保数据完整性

### 回滚操作

如果迁移失败，可以恢复备份：

```bash
mysql -u root -p server_fiber < backup_YYYYMMDD_HHMMSS.sql
```

## 验证清单

迁移完成后，请验证以下项目：

- [ ] comments 表存在
- [ ] 所有必需字段存在且类型正确
- [ ] 所有索引已创建
- [ ] 数据完整性检查通过
- [ ] 应用程序正常运行

## 性能优化

迁移后的性能优化建议：

1. **索引优化**：

   - 为常用查询字段添加索引
   - 定期分析查询性能

2. **数据清理**：

   - 清理无效的评论数据
   - 优化表结构

3. **监控**：
   - 监控数据库性能
   - 设置告警机制

## 联系支持

如果遇到问题，请：

1. 检查错误日志
2. 验证数据库连接
3. 确认权限设置
4. 联系技术支持

---

**重要提醒**：在生产环境执行迁移前，请务必在测试环境充分验证！
