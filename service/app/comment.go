package app

import (
	"server-fiber/global"
	comment "server-fiber/model/app"
	commentReq "server-fiber/model/app/request"
	"server-fiber/model/common/request"
	"strings"

	"gorm.io/gorm"
)

type CommentService struct{}

// CreateComment 创建Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) CreateComment(comment *comment.Comment) (err error) {
	err = global.DB.Create(comment).Error
	return err
}

// DeleteComment 删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteComment(comment comment.Comment) (err error) {
	err = global.DB.Delete(&comment).Error
	return err
}

// DeleteCommentByIds 批量删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteCommentByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]comment.Comment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateComment 更新Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) UpdateComment(comment *comment.Comment) (err error) {
	err = global.DB.Save(comment).Error
	return err
}

// GetComment 根据id获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetComment(id uint) (comment comment.Comment, err error) {
	err = global.DB.Preload("Article").Where("id = ?", id).First(&comment).Error
	return
}

// GetCommentInfoList 分页获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetCommentInfoList(info *commentReq.CommentSearch) (list []comment.Comment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&comment.Comment{}).Preload("Article", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User")
	}).Preload("Praise")
	if info.ArticleId != 0 {
		db = db.Where("article_id = ?", info.ArticleId)
	}
	if info.Content != "" {
		db = db.Where("content like ?", strings.Join([]string{"%", info.Content, "%"}, ""))
		// db = db.Where("MATCH(content) AGAINST('+" + info.Content + "')")
	}
	var comments []comment.Comment
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&comments).Error
	return comments, total, err
}

// GetCommentTreeList 分页获取Treelist

func (commentService *CommentService) GetCommentTreeList(info *commentReq.CommentSearch) (list []comment.Comment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&comment.Comment{})
	if info.ArticleId != 0 {
		db = db.Where("article_id = ?", info.ArticleId)
	}
	if info.Content != "" {
		db = db.Where("content like ?", strings.Join([]string{"%", info.Content, "%"}, ""))
	}
	err = db.Where("parent_id = ?", 0).Count(&total).Error
	if err != nil {
		return
	}

	var commentList []comment.Comment

	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", 0).Preload("Article").Preload("User").Order("id desc").Find(&commentList).Error
	// err = db.Limit(limit).Offset(offset).Where("parent_id = ?", 0).Find(&commentList).Error
	if len(commentList) > 0 {
		for comment := range commentList {
			err = commentService.findChildrenComment(&commentList[comment])
		}
	}

	// log.Println("err1: ", err)

	return commentList, total, err
}

func (commentService *CommentService) findChildrenComment(comment *comment.Comment) (err error) {
	err = global.DB.Where("parent_id = ?", comment.ID).Preload("User").Find(&comment.Children).Error
	if len(comment.Children) > 0 {
		for k := range comment.Children {
			err = commentService.findChildrenComment(&comment.Children[k])
		}
	}
	return err
}

// LikeIt 点赞一条记录
func (*CommentService) PutLikeItOrDislike(info *comment.Praise) (err error) {
	db := global.DB.Model(&comment.Praise{})

	if info.ID == 0 {
		var praise comment.Praise
		err = db.Raw("Select id, comment_id, user_id, created_at, updated_at from praise where user_id = ? and comment_id = ? limit 1", info.UserId, info.CommentId).Scan(&praise).Error
		if err != nil {
			return err
		}
		if praise.ID == 0 {
			err = db.Create(info).Error
		} else {
			info.ID = praise.ID
			info.CreatedAt = praise.CreatedAt
			info.UpdatedAt = praise.UpdatedAt
			err = db.Exec("UPDATE `praise` SET `deleted_at`=NULL,`comment_id`=?,`user_id`=? where id = ? ORDER BY `praise`.`id` LIMIT 1", info.CommentId, info.UserId, info.ID).Error

		}
	} else {
		err = db.Where("id = ?", info.ID).Delete(info).Error
	}

	// if err != nil {
	// 	return err
	// }
	return err
}
