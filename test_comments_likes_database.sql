-- 评论和点赞功能数据库测试脚本
-- 用于验证评论的楼中楼功能和点赞功能

-- 1. 清理测试数据
DELETE FROM likes WHERE post_id IN (SELECT id FROM posts WHERE title LIKE '测试%');
DELETE FROM comments WHERE post_id IN (SELECT id FROM posts WHERE title LIKE '测试%');
DELETE FROM posts WHERE title LIKE '测试%';
DELETE FROM users WHERE name LIKE 'test_user%';

-- 2. 创建测试用户
INSERT INTO users (id, name, password, created_at, updated_at) VALUES 
(1001, 'test_user_1', 'password123', NOW(), NOW()),
(1002, 'test_user_2', 'password123', NOW(), NOW()),
(1003, 'test_user_3', 'password123', NOW(), NOW()),
(1004, 'test_user_4', 'password123', NOW(), NOW());

-- 3. 创建测试帖子
INSERT INTO posts (id, title, text, user_id, like_count, created_at, updated_at) VALUES 
(2001, '测试帖子1 - 评论和点赞功能', '这是用于测试评论和点赞功能的帖子内容', 1001, 0, NOW(), NOW()),
(2002, '测试帖子2 - 楼中楼评论', '这是用于测试楼中楼评论功能的帖子内容', 1001, 0, NOW(), NOW()),
(2003, '测试帖子3 - 并发点赞', '这是用于测试并发点赞功能的帖子内容', 1002, 0, NOW(), NOW());

-- 4. 创建点赞数据
INSERT INTO likes (user_id, post_id, created_at, updated_at) VALUES 
-- 帖子1的点赞
(1001, 2001, NOW(), NOW()),
(1002, 2001, NOW(), NOW()),
(1003, 2001, NOW(), NOW()),
-- 帖子2的点赞
(1001, 2002, NOW(), NOW()),
(1004, 2002, NOW(), NOW()),
-- 帖子3的点赞（模拟并发）
(1001, 2003, NOW(), NOW()),
(1002, 2003, NOW(), NOW()),
(1003, 2003, NOW(), NOW()),
(1004, 2003, NOW(), NOW());

-- 5. 更新帖子点赞数
UPDATE posts SET like_count = 3 WHERE id = 2001;
UPDATE posts SET like_count = 2 WHERE id = 2002;
UPDATE posts SET like_count = 4 WHERE id = 2003;

-- 6. 创建评论数据（包括楼中楼）
INSERT INTO comments (post_id, user_id, content, parent_id, to_user_id, created_at, updated_at) VALUES 
-- 帖子1的评论
(2001, 1001, '这是第一条评论', 0, 0, NOW(), NOW()),
(2001, 1002, '这是第二条评论', 0, 0, NOW(), NOW()),
(2001, 1003, '这是第三条评论', 0, 0, NOW(), NOW()),

-- 帖子2的楼中楼评论
(2002, 1001, '这是父评论', 0, 0, NOW(), NOW()),
(2002, 1002, '这是对父评论的回复1', 1, 1001, NOW(), NOW()),
(2002, 1003, '这是对父评论的回复2', 1, 1001, NOW(), NOW()),
(2002, 1004, '这是对回复1的回复', 2, 1002, NOW(), NOW()),

-- 帖子3的评论
(2003, 1001, '并发测试评论1', 0, 0, NOW(), NOW()),
(2003, 1002, '并发测试评论2', 0, 0, NOW(), NOW());

-- 7. 验证数据完整性
SELECT '=== 数据完整性验证 ===' as info;

-- 验证用户数据
SELECT '用户数据' as test_name, COUNT(*) as count FROM users WHERE name LIKE 'test_user%';

-- 验证帖子数据
SELECT '帖子数据' as test_name, COUNT(*) as count FROM posts WHERE title LIKE '测试%';

-- 验证点赞数据
SELECT '点赞数据' as test_name, COUNT(*) as count FROM likes WHERE post_id IN (2001, 2002, 2003);

-- 验证评论数据
SELECT '评论数据' as test_name, COUNT(*) as count FROM comments WHERE post_id IN (2001, 2002, 2003);

-- 8. 测试点赞功能
SELECT '=== 点赞功能测试 ===' as info;

-- 测试帖子点赞统计
SELECT '帖子点赞统计' as test_name,
       p.id as post_id,
       p.title as post_title,
       p.like_count as post_like_count,
       COUNT(l.id) as actual_like_count,
       CASE WHEN p.like_count = COUNT(l.id) THEN '✓' ELSE '✗' END as status
FROM posts p 
LEFT JOIN likes l ON p.id = l.post_id AND l.deleted_at IS NULL
WHERE p.id IN (2001, 2002, 2003)
GROUP BY p.id, p.title, p.like_count;

-- 测试用户点赞的帖子
SELECT '用户点赞的帖子' as test_name,
       u.name as user_name,
       p.title as post_title,
       l.created_at as like_time
FROM users u
JOIN likes l ON u.id = l.user_id AND l.deleted_at IS NULL
JOIN posts p ON l.post_id = p.id AND p.deleted_at IS NULL
WHERE u.name LIKE 'test_user%'
ORDER BY u.id, l.created_at DESC;

-- 9. 测试评论功能
SELECT '=== 评论功能测试 ===' as info;

-- 测试帖子评论统计
SELECT '帖子评论统计' as test_name,
       p.id as post_id,
       p.title as post_title,
       COUNT(c.id) as total_comments,
       COUNT(CASE WHEN c.parent_id = 0 THEN 1 END) as parent_comments,
       COUNT(CASE WHEN c.parent_id > 0 THEN 1 END) as child_comments
FROM posts p
LEFT JOIN comments c ON p.id = c.post_id AND c.deleted_at IS NULL
WHERE p.id IN (2001, 2002, 2003)
GROUP BY p.id, p.title;

-- 测试楼中楼评论结构
SELECT '楼中楼评论结构' as test_name,
       p.title as post_title,
       c1.content as parent_comment,
       c1.user_id as parent_user_id,
       c2.content as child_comment,
       c2.user_id as child_user_id,
       c2.to_user_id as reply_to_user_id
FROM posts p
JOIN comments c1 ON p.id = c1.post_id AND c1.parent_id = 0 AND c1.deleted_at IS NULL
LEFT JOIN comments c2 ON c1.id = c2.parent_id AND c2.deleted_at IS NULL
WHERE p.id = 2002
ORDER BY c1.id, c2.id;

-- 测试评论层级深度
WITH RECURSIVE comment_hierarchy AS (
    -- 基础查询：所有父评论
    SELECT 
        id,
        post_id,
        content,
        parent_id,
        user_id,
        to_user_id,
        1 as level,
        CAST(id AS CHAR(1000)) as path
    FROM comments 
    WHERE parent_id = 0 AND post_id = 2002 AND deleted_at IS NULL
    
    UNION ALL
    
    -- 递归查询：子评论
    SELECT 
        c.id,
        c.post_id,
        c.content,
        c.parent_id,
        c.user_id,
        c.to_user_id,
        ch.level + 1,
        CONCAT(ch.path, '->', c.id)
    FROM comments c
    JOIN comment_hierarchy ch ON c.parent_id = ch.id
    WHERE c.deleted_at IS NULL
)
SELECT '评论层级深度' as test_name,
       level,
       COUNT(*) as comment_count,
       GROUP_CONCAT(content ORDER BY id) as comments
FROM comment_hierarchy
GROUP BY level
ORDER BY level;

-- 10. 性能测试
SELECT '=== 性能测试 ===' as info;

-- 测试点赞查询性能
EXPLAIN SELECT 
    p.id, p.title, p.like_count,
    COUNT(l.id) as actual_likes
FROM posts p
LEFT JOIN likes l ON p.id = l.post_id
WHERE p.id IN (2001, 2002, 2003)
GROUP BY p.id, p.title, p.like_count;

-- 测试评论查询性能
EXPLAIN SELECT 
    p.id, p.title,
    c.id as comment_id, c.content, c.parent_id,
    u.name as user_name
FROM posts p
JOIN comments c ON p.id = c.post_id
JOIN users u ON c.user_id = u.id
WHERE p.id IN (2001, 2002, 2003)
ORDER BY p.id, c.parent_id, c.id;

-- 11. 数据一致性检查
SELECT '=== 数据一致性检查 ===' as info;

-- 检查点赞数一致性
SELECT '点赞数一致性检查' as test_name,
       p.id as post_id,
       p.like_count as stored_count,
       COUNT(l.id) as actual_count,
       CASE WHEN p.like_count = COUNT(l.id) THEN '✓ 一致' ELSE '✗ 不一致' END as status
FROM posts p
LEFT JOIN likes l ON p.id = l.post_id AND l.deleted_at IS NULL
WHERE p.id IN (2001, 2002, 2003)
GROUP BY p.id, p.like_count;

-- 检查评论的post_id引用完整性
SELECT '评论post_id引用完整性' as test_name,
       c.id as comment_id,
       c.post_id,
       p.title as post_title,
       CASE WHEN p.id IS NOT NULL THEN '✓ 有效' ELSE '✗ 无效' END as status
FROM comments c
LEFT JOIN posts p ON c.post_id = p.id
WHERE c.post_id IN (2001, 2002, 2003);

-- 检查评论的parent_id引用完整性
SELECT '评论parent_id引用完整性' as test_name,
       c.id as comment_id,
       c.parent_id,
       pc.content as parent_content,
       CASE WHEN pc.id IS NOT NULL OR c.parent_id = 0 THEN '✓ 有效' ELSE '✗ 无效' END as status
FROM comments c
LEFT JOIN comments pc ON c.parent_id = pc.id
WHERE c.post_id IN (2001, 2002, 2003) AND c.parent_id > 0;

-- 12. 清理测试数据（可选）
-- 取消注释以下行来清理测试数据
-- DELETE FROM likes WHERE post_id IN (2001, 2002, 2003);
-- DELETE FROM comments WHERE post_id IN (2001, 2002, 2003);
-- DELETE FROM posts WHERE id IN (2001, 2002, 2003);
-- DELETE FROM users WHERE id IN (1001, 1002, 1003, 1004);

SELECT '=== 测试完成 ===' as info;
