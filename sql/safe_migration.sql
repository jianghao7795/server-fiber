-- 安全的数据库迁移脚本
-- 用于更新comments表结构以支持新的点赞功能

-- 开始事务
START TRANSACTION;

-- 1. 备份现有数据（可选）
-- CREATE TABLE comments_backup AS SELECT * FROM comments;

-- 2. 检查当前表结构
SELECT 'Current comments table structure:' as info;
DESCRIBE comments;

-- 3. 检查现有索引
SELECT 'Current indexes:' as info;
SHOW INDEX FROM comments;

-- 4. 安全地更新表结构
-- 注意：以下操作会逐步执行，如果任何一步失败，整个事务会回滚

-- 4.1 如果存在article_id字段，重命名为post_id
SET @sql = (SELECT IF(
  EXISTS(
    SELECT 1 FROM information_schema.COLUMNS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = 'comments' 
    AND COLUMN_NAME = 'article_id'
  ),
  'ALTER TABLE comments CHANGE COLUMN article_id post_id bigint unsigned NOT NULL COMMENT ''帖子ID''',
  'SELECT ''article_id column does not exist, skipping rename'' as message'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 4.2 确保post_id字段存在且类型正确
ALTER TABLE comments 
  MODIFY COLUMN post_id bigint unsigned NOT NULL COMMENT '帖子ID';

-- 4.3 更新parent_id字段
ALTER TABLE comments 
  MODIFY COLUMN parent_id bigint unsigned DEFAULT '0' COMMENT '父评论ID';

-- 4.4 更新user_id字段
ALTER TABLE comments 
  MODIFY COLUMN user_id bigint unsigned NOT NULL COMMENT '用户ID';

-- 4.5 更新to_user_id字段
ALTER TABLE comments 
  MODIFY COLUMN to_user_id bigint unsigned DEFAULT '0' COMMENT '回复用户ID';

-- 4.6 更新content字段
ALTER TABLE comments 
  MODIFY COLUMN content text COMMENT '评论内容';

-- 5. 添加必要的索引（如果不存在）
-- 5.1 添加post_id索引
SET @sql = (SELECT IF(
  NOT EXISTS(
    SELECT 1 FROM information_schema.STATISTICS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = 'comments' 
    AND INDEX_NAME = 'idx_comments_post_id'
  ),
  'ALTER TABLE comments ADD KEY idx_comments_post_id (post_id)',
  'SELECT ''idx_comments_post_id already exists'' as message'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 5.2 添加parent_id索引
SET @sql = (SELECT IF(
  NOT EXISTS(
    SELECT 1 FROM information_schema.STATISTICS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = 'comments' 
    AND INDEX_NAME = 'idx_comments_parent_id'
  ),
  'ALTER TABLE comments ADD KEY idx_comments_parent_id (parent_id)',
  'SELECT ''idx_comments_parent_id already exists'' as message'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 5.3 添加user_id索引
SET @sql = (SELECT IF(
  NOT EXISTS(
    SELECT 1 FROM information_schema.STATISTICS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = 'comments' 
    AND INDEX_NAME = 'idx_comments_user_id'
  ),
  'ALTER TABLE comments ADD KEY idx_comments_user_id (user_id)',
  'SELECT ''idx_comments_user_id already exists'' as message'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 5.4 添加to_user_id索引
SET @sql = (SELECT IF(
  NOT EXISTS(
    SELECT 1 FROM information_schema.STATISTICS 
    WHERE TABLE_SCHEMA = DATABASE() 
    AND TABLE_NAME = 'comments' 
    AND INDEX_NAME = 'idx_comments_to_user_id'
  ),
  'ALTER TABLE comments ADD KEY idx_comments_to_user_id (to_user_id)',
  'SELECT ''idx_comments_to_user_id already exists'' as message'
));
PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- 6. 验证更新后的表结构
SELECT 'Updated comments table structure:' as info;
DESCRIBE comments;

-- 7. 验证索引
SELECT 'Updated indexes:' as info;
SHOW INDEX FROM comments;

-- 8. 检查数据完整性
SELECT 'Data integrity check:' as info;
SELECT 
  COUNT(*) as total_comments,
  COUNT(DISTINCT post_id) as unique_posts,
  COUNT(DISTINCT user_id) as unique_users,
  COUNT(CASE WHEN parent_id > 0 THEN 1 END) as reply_comments
FROM comments;

-- 9. 提交事务
COMMIT;

-- 10. 显示迁移完成信息
SELECT 'Migration completed successfully!' as status;
