# 文章点赞功能 API 使用说明

## 概述

根据 ERD 图设计，实现了完整的文章点赞功能，包括：

- **用户 User**：每个用户有多篇帖子和多条评论和多个点赞
- **分类 Category**：帖子所属分类类目支持多个分类一个帖子或者多个帖子公用同一个分类
- **帖子 Post**：每篇帖子有多个分类并可获得多个赞
- **评论 Comment**：每条评论属于一个用户并关联一篇帖子，且评论支持楼中楼
- **点赞 Like**：每个赞关联一篇帖子，多个点赞可以关联同一篇帖子
- **分类帖子关联 CategoryPost**：帖子和分类的关联关系表

## API 接口

### 1. 点赞帖子

**POST** `/backend/like/likePost/{post_id}`

**描述**：用户点赞指定帖子

**请求参数**：

- `post_id` (路径参数): 帖子 ID

**响应示例**：

```json
{
  "code": 200,
  "msg": "点赞成功",
  "data": null
}
```

### 2. 取消点赞帖子

**DELETE** `/backend/like/unlikePost/{post_id}`

**描述**：用户取消点赞指定帖子

**请求参数**：

- `post_id` (路径参数): 帖子 ID

**响应示例**：

```json
{
  "code": 200,
  "msg": "取消点赞成功",
  "data": null
}
```

### 3. 获取帖子点赞列表

**GET** `/backend/like/getPostLikes/{post_id}`

**描述**：分页获取指定帖子的点赞用户列表

**请求参数**：

- `post_id` (路径参数): 帖子 ID
- `page` (查询参数): 页码，默认 1
- `page_size` (查询参数): 每页数量，默认 10

**响应示例**：

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "created_at": "2024-01-01T10:00:00Z",
        "user_id": 1,
        "post_id": 1,
        "user": {
          "id": 1,
          "name": "用户名",
          "header": "头像URL"
        }
      }
    ],
    "total": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

### 4. 检查用户是否点赞了帖子

**GET** `/backend/like/checkUserLiked/{post_id}`

**描述**：检查当前用户是否点赞了指定帖子

**请求参数**：

- `post_id` (路径参数): 帖子 ID

**响应示例**：

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": {
    "liked": true
  }
}
```

### 5. 获取用户点赞的帖子列表

**GET** `/backend/like/getUserLikedPosts`

**描述**：分页获取当前用户点赞的帖子列表

**请求参数**：

- `page` (查询参数): 页码，默认 1
- `page_size` (查询参数): 每页数量，默认 10

**响应示例**：

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": {
    "list": [
      {
        "id": 1,
        "created_at": "2024-01-01T10:00:00Z",
        "user_id": 1,
        "post_id": 1,
        "post": {
          "id": 1,
          "title": "帖子标题",
          "text": "帖子内容",
          "like_count": 5,
          "user": {
            "id": 1,
            "name": "作者名"
          },
          "categories": [
            {
              "id": 1,
              "name": "技术"
            }
          ]
        }
      }
    ],
    "total": 5,
    "page": 1,
    "pageSize": 10
  }
}
```

### 6. 获取帖子点赞数

**GET** `/backend/like/getPostLikeCount/{post_id}`

**描述**：获取指定帖子的点赞总数

**请求参数**：

- `post_id` (路径参数): 帖子 ID

**响应示例**：

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": {
    "like_count": 25
  }
}
```

## 数据库表结构

### 1. likes 表

```sql
CREATE TABLE `likes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `post_id` bigint unsigned NOT NULL COMMENT '帖子ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_likes_user_post` (`user_id`, `post_id`)
);
```

### 2. posts 表

```sql
CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) DEFAULT NULL COMMENT '帖子标题',
  `text` longtext COMMENT '帖子内容',
  `user_id` bigint unsigned NOT NULL COMMENT '帖子作者ID',
  `like_count` int DEFAULT '0' COMMENT '点赞数',
  -- 其他字段...
  PRIMARY KEY (`id`)
);
```

### 3. categories 表

```sql
CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT '分类名称',
  `sort` int DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`)
);
```

### 4. category_posts 表

```sql
CREATE TABLE `category_posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `category_id` bigint unsigned NOT NULL COMMENT '分类ID',
  `post_id` bigint unsigned NOT NULL COMMENT '帖子ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_category_posts_category_post` (`category_id`, `post_id`)
);
```

## 使用说明

1. **认证**：所有 API 都需要 JWT 认证，在请求头中携带 `Authorization: Bearer <token>`
2. **权限**：用户只能操作自己的点赞记录
3. **防重复**：系统会自动防止用户重复点赞同一篇帖子
4. **事务安全**：点赞和取消点赞操作都使用数据库事务，确保数据一致性
5. **性能优化**：使用索引优化查询性能，支持分页查询

## 错误处理

- `400`：参数错误
- `401`：用户未登录
- `500`：服务器内部错误

常见错误信息：

- "您已经点赞过这篇帖子了"
- "您还没有点赞过这篇帖子"
- "帖子不存在"
- "用户未登录"
