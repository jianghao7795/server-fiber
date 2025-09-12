-- 数据库测试脚本
-- 用于测试文章点赞功能的数据完整性

-- 1. 创建测试数据
INSERT INTO users (id, name, password, created_at, updated_at) VALUES 
(1, 'test_user_1', 'password123', NOW(), NOW()),
(2, 'test_user_2', 'password123', NOW(), NOW()),
(3, 'test_user_3', 'password123', NOW(), NOW());

-- 2. 创建测试帖子
INSERT INTO posts (id, title, text, user_id, like_count, created_at, updated_at) VALUES 
(1, '测试帖子1', '这是第一个测试帖子的内容', 1, 0, NOW(), NOW()),
(2, '测试帖子2', '这是第二个测试帖子的内容', 1, 0, NOW(), NOW()),
(3, '测试帖子3', '这是第三个测试帖子的内容', 2, 0, NOW(), NOW());

-- 3. 创建测试分类
INSERT INTO categories (id, name, sort, created_at, updated_at) VALUES 
(1, '技术', 1, NOW(), NOW()),
(2, '生活', 2, NOW(), NOW()),
(3, '娱乐', 3, NOW(), NOW());

-- 4. 创建分类帖子关联
INSERT INTO category_posts (category_id, post_id, created_at, updated_at) VALUES 
(1, 1, NOW(), NOW()),
(1, 2, NOW(), NOW()),
(2, 3, NOW(), NOW());

-- 5. 创建测试点赞数据
INSERT INTO likes (user_id, post_id, created_at, updated_at) VALUES 
(2, 1, NOW(), NOW()),
(3, 1, NOW(), NOW()),
(1, 2, NOW(), NOW()),
(2, 3, NOW(), NOW()),
(3, 3, NOW(), NOW());

-- 6. 更新帖子点赞数
UPDATE posts SET like_count = 2 WHERE id = 1;
UPDATE posts SET like_count = 1 WHERE id = 2;
UPDATE posts SET like_count = 2 WHERE id = 3;

-- 7. 创建测试评论（支持楼中楼）
INSERT INTO comments (post_id, user_id, content, parent_id, created_at, updated_at) VALUES 
(1, 2, '这是一条评论', 0, NOW(), NOW()),
(1, 3, '这是对评论的回复', 1, NOW(), NOW()),
(2, 1, '这是另一条评论', 0, NOW(), NOW());

-- 8. 验证数据完整性
SELECT '用户数据' as test_name, COUNT(*) as count FROM users WHERE deleted_at IS NULL;
SELECT '帖子数据' as test_name, COUNT(*) as count FROM posts WHERE deleted_at IS NULL;
SELECT '分类数据' as test_name, COUNT(*) as count FROM categories WHERE deleted_at IS NULL;
SELECT '点赞数据' as test_name, COUNT(*) as count FROM likes WHERE deleted_at IS NULL;
SELECT '分类帖子关联' as test_name, COUNT(*) as count FROM category_posts WHERE deleted_at IS NULL;
SELECT '评论数据' as test_name, COUNT(*) as count FROM comments WHERE deleted_at IS NULL;

-- 9. 测试查询
SELECT '帖子点赞统计' as test_name, 
       p.id, 
       p.title, 
       p.like_count as post_like_count,
       COUNT(l.id) as actual_like_count
FROM posts p 
LEFT JOIN likes l ON p.id = l.post_id AND l.deleted_at IS NULL
WHERE p.deleted_at IS NULL
GROUP BY p.id, p.title, p.like_count;

-- 10. 测试用户点赞的帖子
SELECT '用户点赞的帖子' as test_name,
       u.name as user_name,
       p.title as post_title,
       l.created_at as like_time
FROM users u
JOIN likes l ON u.id = l.user_id AND l.deleted_at IS NULL
JOIN posts p ON l.post_id = p.id AND p.deleted_at IS NULL
WHERE u.deleted_at IS NULL
ORDER BY u.id, l.created_at DESC;

-- 11. 测试楼中楼评论
SELECT '楼中楼评论' as test_name,
       p.title as post_title,
       c1.content as parent_comment,
       c2.content as child_comment,
       u1.name as parent_user,
       u2.name as child_user
FROM posts p
JOIN comments c1 ON p.id = c1.post_id AND c1.parent_id = 0 AND c1.deleted_at IS NULL
JOIN comments c2 ON c1.id = c2.parent_id AND c2.deleted_at IS NULL
JOIN users u1 ON c1.user_id = u1.id
JOIN users u2 ON c2.user_id = u2.id
WHERE p.deleted_at IS NULL;

-- 12. 测试分类帖子关联
SELECT '分类帖子关联' as test_name,
       c.name as category_name,
       p.title as post_title,
       u.name as author_name
FROM categories c
JOIN category_posts cp ON c.id = cp.category_id AND cp.deleted_at IS NULL
JOIN posts p ON cp.post_id = p.id AND p.deleted_at IS NULL
JOIN users u ON p.user_id = u.id
WHERE c.deleted_at IS NULL
ORDER BY c.name, p.title;

-- 13. 清理测试数据（可选）
-- DELETE FROM likes WHERE user_id IN (1, 2, 3);
-- DELETE FROM comments WHERE user_id IN (1, 2, 3);
-- DELETE FROM category_posts WHERE post_id IN (1, 2, 3);
-- DELETE FROM posts WHERE id IN (1, 2, 3);
-- DELETE FROM categories WHERE id IN (1, 2, 3);
-- DELETE FROM users WHERE id IN (1, 2, 3);
