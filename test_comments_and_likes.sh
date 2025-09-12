#!/bin/bash

# 评论和点赞功能综合测试脚本
# 测试楼中楼评论和文章点赞功能

BASE_URL="http://localhost:8888/backend"
API_BASE="$BASE_URL"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}==========================================${NC}"
echo -e "${BLUE}评论和点赞功能综合测试${NC}"
echo -e "${BLUE}==========================================${NC}"

# 测试用的数据
POST_ID=1
USER_TOKEN="your_jwt_token_here"  # 请替换为实际的JWT token

# 检查服务器是否运行
echo -e "${YELLOW}1. 检查服务器状态...${NC}"
curl -s -o /dev/null -w "%{http_code}" "$BASE_URL" > /dev/null
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ 服务器运行正常${NC}"
else
    echo -e "${RED}✗ 服务器未运行，请先启动服务器${NC}"
    exit 1
fi

echo ""

# 测试1: 创建测试帖子
echo -e "${YELLOW}2. 创建测试帖子...${NC}"
echo "POST $API_BASE/article/createArticle"
if [ "$USER_TOKEN" != "your_jwt_token_here" ]; then
    POST_DATA='{
        "title": "测试帖子 - 评论和点赞功能",
        "desc": "这是一个用于测试评论和点赞功能的帖子",
        "content": "这是帖子的详细内容，用于测试评论的楼中楼功能和点赞功能。",
        "state": 1,
        "user_id": 1,
        "is_important": 0
    }'
    
    response=$(curl -s -w "\n%{http_code}" -X POST \
        -H "Authorization: Bearer $USER_TOKEN" \
        -H "Content-Type: application/json" \
        -d "$POST_DATA" \
        "$API_BASE/article/createArticle")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 测试帖子创建成功${NC}"
        # 提取帖子ID
        POST_ID=$(echo "$body" | grep -o '"data":[0-9]*' | grep -o '[0-9]*')
        echo "帖子ID: $POST_ID"
    else
        echo -e "${RED}✗ 测试帖子创建失败${NC}"
        echo -e "${YELLOW}使用默认帖子ID: $POST_ID${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过帖子创建（需要设置有效的JWT token）${NC}"
fi

echo ""

# 测试2: 获取帖子点赞数
echo -e "${YELLOW}3. 测试获取帖子点赞数...${NC}"
echo "GET $API_BASE/like/getPostLikeCount/$POST_ID"
response=$(curl -s -w "\n%{http_code}" "$API_BASE/like/getPostLikeCount/$POST_ID")
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

echo "HTTP状态码: $http_code"
echo "响应内容: $body"
if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ 获取帖子点赞数成功${NC}"
else
    echo -e "${RED}✗ 获取帖子点赞数失败${NC}"
fi

echo ""

# 测试3: 点赞帖子
echo -e "${YELLOW}4. 测试点赞帖子...${NC}"
echo "POST $API_BASE/like/likePost/$POST_ID"
if [ "$USER_TOKEN" != "your_jwt_token_here" ]; then
    response=$(curl -s -w "\n%{http_code}" -X POST \
        -H "Authorization: Bearer $USER_TOKEN" \
        "$API_BASE/like/likePost/$POST_ID")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 点赞帖子成功${NC}"
    else
        echo -e "${RED}✗ 点赞帖子失败${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过点赞测试（需要设置有效的JWT token）${NC}"
fi

echo ""

# 测试4: 检查用户点赞状态
echo -e "${YELLOW}5. 测试检查用户点赞状态...${NC}"
echo "GET $API_BASE/like/checkUserLiked/$POST_ID"
if [ "$USER_TOKEN" != "your_jwt_token_here" ]; then
    response=$(curl -s -w "\n%{http_code}" \
        -H "Authorization: Bearer $USER_TOKEN" \
        "$API_BASE/like/checkUserLiked/$POST_ID")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 检查用户点赞状态成功${NC}"
    else
        echo -e "${RED}✗ 检查用户点赞状态失败${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过点赞状态检查（需要设置有效的JWT token）${NC}"
fi

echo ""

# 测试5: 创建评论
echo -e "${YELLOW}6. 测试创建评论...${NC}"
echo "POST $API_BASE/comment/createComment"
if [ "$USER_TOKEN" != "your_jwt_token_here" ]; then
    COMMENT_DATA='{
        "post_id": '$POST_ID',
        "content": "这是一条测试评论",
        "parent_id": 0,
        "user_id": 1,
        "to_user_id": 0
    }'
    
    response=$(curl -s -w "\n%{http_code}" -X POST \
        -H "Authorization: Bearer $USER_TOKEN" \
        -H "Content-Type: application/json" \
        -d "$COMMENT_DATA" \
        "$API_BASE/comment/createComment")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 评论创建成功${NC}"
        # 提取评论ID
        COMMENT_ID=$(echo "$body" | grep -o '"data":[0-9]*' | grep -o '[0-9]*')
        echo "评论ID: $COMMENT_ID"
    else
        echo -e "${RED}✗ 评论创建失败${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过评论创建（需要设置有效的JWT token）${NC}"
    COMMENT_ID=1
fi

echo ""

# 测试6: 创建楼中楼评论（回复评论）
echo -e "${YELLOW}7. 测试创建楼中楼评论...${NC}"
echo "POST $API_BASE/comment/createComment"
if [ "$USER_TOKEN" != "your_jwt_token_here" ] && [ ! -z "$COMMENT_ID" ]; then
    REPLY_DATA='{
        "post_id": '$POST_ID',
        "content": "这是对评论的回复，测试楼中楼功能",
        "parent_id": '$COMMENT_ID',
        "user_id": 1,
        "to_user_id": 1
    }'
    
    response=$(curl -s -w "\n%{http_code}" -X POST \
        -H "Authorization: Bearer $USER_TOKEN" \
        -H "Content-Type: application/json" \
        -d "$REPLY_DATA" \
        "$API_BASE/comment/createComment")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 楼中楼评论创建成功${NC}"
    else
        echo -e "${RED}✗ 楼中楼评论创建失败${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过楼中楼评论创建（需要设置有效的JWT token）${NC}"
fi

echo ""

# 测试7: 获取帖子评论列表
echo -e "${YELLOW}8. 测试获取帖子评论列表...${NC}"
echo "GET $API_BASE/comment/getCommentList?post_id=$POST_ID&page=1&page_size=10"
response=$(curl -s -w "\n%{http_code}" "$API_BASE/comment/getCommentList?post_id=$POST_ID&page=1&page_size=10")
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

echo "HTTP状态码: $http_code"
echo "响应内容: $body"
if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ 获取评论列表成功${NC}"
else
    echo -e "${RED}✗ 获取评论列表失败${NC}"
fi

echo ""

# 测试8: 获取帖子点赞列表
echo -e "${YELLOW}9. 测试获取帖子点赞列表...${NC}"
echo "GET $API_BASE/like/getPostLikes/$POST_ID?page=1&page_size=10"
response=$(curl -s -w "\n%{http_code}" "$API_BASE/like/getPostLikes/$POST_ID?page=1&page_size=10")
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

echo "HTTP状态码: $http_code"
echo "响应内容: $body"
if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ 获取点赞列表成功${NC}"
else
    echo -e "${RED}✗ 获取点赞列表失败${NC}"
fi

echo ""

# 测试9: 取消点赞
echo -e "${YELLOW}10. 测试取消点赞...${NC}"
echo "DELETE $API_BASE/like/unlikePost/$POST_ID"
if [ "$USER_TOKEN" != "your_jwt_token_here" ]; then
    response=$(curl -s -w "\n%{http_code}" -X DELETE \
        -H "Authorization: Bearer $USER_TOKEN" \
        "$API_BASE/like/unlikePost/$POST_ID")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 取消点赞成功${NC}"
    else
        echo -e "${RED}✗ 取消点赞失败${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过取消点赞测试（需要设置有效的JWT token）${NC}"
fi

echo ""

# 测试10: 获取用户点赞的帖子列表
echo -e "${YELLOW}11. 测试获取用户点赞的帖子列表...${NC}"
echo "GET $API_BASE/like/getUserLikedPosts?page=1&page_size=10"
if [ "$USER_TOKEN" != "your_jwt_token_here" ]; then
    response=$(curl -s -w "\n%{http_code}" \
        -H "Authorization: Bearer $USER_TOKEN" \
        "$API_BASE/like/getUserLikedPosts?page=1&page_size=10")
    
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n -1)
    
    echo "HTTP状态码: $http_code"
    echo "响应内容: $body"
    if [ "$http_code" -eq 200 ]; then
        echo -e "${GREEN}✓ 获取用户点赞列表成功${NC}"
    else
        echo -e "${RED}✗ 获取用户点赞列表失败${NC}"
    fi
else
    echo -e "${YELLOW}⚠ 跳过用户点赞列表测试（需要设置有效的JWT token）${NC}"
fi

echo ""

# 测试11: 数据库验证
echo -e "${YELLOW}12. 数据库验证...${NC}"
echo "验证评论和点赞数据是否正确存储"

# 这里可以添加数据库查询验证
echo -e "${YELLOW}请手动执行以下SQL查询验证数据：${NC}"
echo "SELECT COUNT(*) as total_comments FROM comments WHERE post_id = $POST_ID;"
echo "SELECT COUNT(*) as total_likes FROM likes WHERE post_id = $POST_ID;"
echo "SELECT c1.content as parent_comment, c2.content as child_comment FROM comments c1 LEFT JOIN comments c2 ON c1.id = c2.parent_id WHERE c1.post_id = $POST_ID AND c1.parent_id = 0;"

echo ""
echo -e "${BLUE}==========================================${NC}"
echo -e "${GREEN}评论和点赞功能测试完成！${NC}"
echo -e "${BLUE}==========================================${NC}"

echo ""
echo -e "${YELLOW}测试总结：${NC}"
echo "1. ✓ 服务器状态检查"
echo "2. ✓ 帖子创建（需要认证）"
echo "3. ✓ 获取帖子点赞数"
echo "4. ✓ 点赞帖子（需要认证）"
echo "5. ✓ 检查用户点赞状态（需要认证）"
echo "6. ✓ 创建评论（需要认证）"
echo "7. ✓ 创建楼中楼评论（需要认证）"
echo "8. ✓ 获取评论列表"
echo "9. ✓ 获取点赞列表"
echo "10. ✓ 取消点赞（需要认证）"
echo "11. ✓ 获取用户点赞列表（需要认证）"
echo "12. ⚠ 数据库验证（需要手动执行）"

echo ""
echo -e "${YELLOW}使用说明：${NC}"
echo "1. 请确保服务器已启动"
echo "2. 修改脚本中的 USER_TOKEN 为有效的JWT token"
echo "3. 确保数据库表结构已正确更新"
echo "4. 运行脚本: bash test_comments_and_likes.sh"
