-- 创建点赞相关表

-- 创建分类表
CREATE TABLE IF NOT EXISTS `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类表';

-- 创建帖子表（如果不存在）
CREATE TABLE IF NOT EXISTS `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) DEFAULT NULL COMMENT '帖子标题',
  `text` longtext COMMENT '帖子内容',
  `publish_at` datetime(3) DEFAULT NULL COMMENT '发布时间',
  `user_id` bigint unsigned NOT NULL COMMENT '帖子作者ID',
  `state` int DEFAULT '1' COMMENT '帖子状态',
  `is_important` int DEFAULT '0' COMMENT '首页是否显示',
  `reading_quantity` int DEFAULT '0' COMMENT '阅读量',
  `like_count` int DEFAULT '0' COMMENT '点赞数',
  PRIMARY KEY (`id`),
  KEY `idx_posts_deleted_at` (`deleted_at`),
  KEY `idx_posts_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='帖子表';

-- 创建点赞表
CREATE TABLE IF NOT EXISTS `likes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `post_id` bigint unsigned NOT NULL COMMENT '帖子ID',
  PRIMARY KEY (`id`),
  KEY `idx_likes_deleted_at` (`deleted_at`),
  KEY `idx_likes_user_id` (`user_id`),
  KEY `idx_likes_post_id` (`post_id`),
  UNIQUE KEY `idx_likes_user_post` (`user_id`, `post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='点赞表';

-- 创建分类帖子关联表
CREATE TABLE IF NOT EXISTS `category_posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `category_id` bigint unsigned NOT NULL COMMENT '分类ID',
  `post_id` bigint unsigned NOT NULL COMMENT '帖子ID',
  PRIMARY KEY (`id`),
  KEY `idx_category_posts_deleted_at` (`deleted_at`),
  KEY `idx_category_posts_category_id` (`category_id`),
  KEY `idx_category_posts_post_id` (`post_id`),
  UNIQUE KEY `idx_category_posts_category_post` (`category_id`, `post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='分类帖子关联表';

-- 更新评论表结构（如果存在）
-- 注意：这里假设comments表已经存在，需要根据实际情况调整

-- 检查并更新comments表结构
-- 如果comments表不存在，先创建
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

-- 如果comments表已存在，更新字段结构
-- 注意：执行前请备份数据！

-- 更新post_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `post_id` bigint unsigned NOT NULL COMMENT '帖子ID',
  ADD KEY `idx_comments_post_id` (`post_id`);

-- 更新parent_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `parent_id` bigint unsigned DEFAULT '0' COMMENT '父评论ID',
  ADD KEY `idx_comments_parent_id` (`parent_id`);

-- 更新user_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  ADD KEY `idx_comments_user_id` (`user_id`);

-- 更新to_user_id字段
ALTER TABLE `comments` 
  MODIFY COLUMN `to_user_id` bigint unsigned DEFAULT '0' COMMENT '回复用户ID',
  ADD KEY `idx_comments_to_user_id` (`to_user_id`);

-- 更新content字段
ALTER TABLE `comments` 
  MODIFY COLUMN `content` text COMMENT '评论内容';
