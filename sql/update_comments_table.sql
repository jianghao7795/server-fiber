-- 更新comments表结构的迁移脚本
-- 执行前请务必备份数据库！

-- 1. 检查comments表是否存在
SELECT TABLE_NAME 
FROM information_schema.TABLES 
WHERE TABLE_SCHEMA = DATABASE() 
AND TABLE_NAME = 'comments';

-- 2. 如果comments表不存在，创建新表
CREATE TABLE IF NOT EXISTS `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `post_id` bigint unsigned NOT NULL COMMENT '帖子ID',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '父评论ID',
  `content` text COMMENT '评论内容',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `to_user_id` bigint unsigned DEFAULT '0' COMMENT '回复用户ID',
  PRIMARY KEY (`id`),
  KEY `idx_comments_deleted_at` (`deleted_at`),
  KEY `idx_comments_post_id` (`post_id`),
  KEY `idx_comments_parent_id` (`parent_id`),
  KEY `idx_comments_user_id` (`user_id`),
  KEY `idx_comments_to_user_id` (`to_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='评论表';

-- 3. 如果comments表已存在，安全地更新字段结构
-- 注意：以下操作会修改现有表结构，请确保已备份数据

-- 检查现有字段
SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT
FROM information_schema.COLUMNS 
WHERE TABLE_SCHEMA = DATABASE() 
AND TABLE_NAME = 'comments'
ORDER BY ORDINAL_POSITION;

-- 更新post_id字段（从article_id改为post_id）
-- 如果存在article_id字段，先重命名
ALTER TABLE `comments` 
  CHANGE COLUMN `article_id` `post_id` bigint unsigned NOT NULL COMMENT '帖子ID';

-- 添加post_id索引
ALTER TABLE `comments` 
  ADD KEY `idx_comments_post_id` (`post_id`);

-- 更新parent_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `parent_id` bigint unsigned DEFAULT '0' COMMENT '父评论ID';

-- 添加parent_id索引
ALTER TABLE `comments` 
  ADD KEY `idx_comments_parent_id` (`parent_id`);

-- 更新user_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `user_id` bigint unsigned NOT NULL COMMENT '用户ID';

-- 添加user_id索引
ALTER TABLE `comments` 
  ADD KEY `idx_comments_user_id` (`user_id`);

-- 更新to_user_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `to_user_id` bigint unsigned DEFAULT '0' COMMENT '回复用户ID';

-- 添加to_user_id索引
ALTER TABLE `comments` 
  ADD KEY `idx_comments_to_user_id` (`to_user_id`);

-- 更新content字段
ALTER TABLE `comments` 
  MODIFY COLUMN `content` text COMMENT '评论内容';

-- 4. 验证表结构
DESCRIBE `comments`;

-- 5. 验证索引
SHOW INDEX FROM `comments`;

-- 6. 检查外键约束（如果需要）
SELECT 
  CONSTRAINT_NAME,
  COLUMN_NAME,
  REFERENCED_TABLE_NAME,
  REFERENCED_COLUMN_NAME
FROM information_schema.KEY_COLUMN_USAGE 
WHERE TABLE_SCHEMA = DATABASE() 
AND TABLE_NAME = 'comments' 
AND REFERENCED_TABLE_NAME IS NOT NULL;
