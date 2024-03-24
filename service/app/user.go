package app

import (
	"errors"

	"server-fiber/global"
	"server-fiber/model/app"
	appReq "server-fiber/model/app/request"
	"server-fiber/model/common/request"

	"gorm.io/gorm"
)

type UserService struct {
}

// CreateUser 创建User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user *app.User) (err error) {
	var userCurrent app.User
	if !errors.Is(global.DB.Where("name = ?", user.Name).First(&userCurrent).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}
	// user.Password = utils.MD5V([]byte(user.Password))
	err = global.DB.Create(user).Error
	return err
}

// DeleteUser 删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUser(ID int) (err error) {
	err = global.DB.Where("id = ?", ID).Delete(&app.User{}).Error
	return err
}

// DeleteUserByIds 批量删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUserByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]app.User{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUser 更新User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) UpdateUser(user *app.User) (err error) {
	err = global.DB.Save(user).Error
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUser(id uint) (user app.User, err error) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUserInfoList(info *appReq.UserSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&app.User{})

	var users []app.User
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if info.Name != "" {
		db = db.Where("name like ?", "%"+info.Name+"%")
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&users).Error
	return users, total, err
}

func (userService *UserService) FindIsUser(id uint) bool {
	var user app.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	return errors.Is(err, gorm.ErrRecordNotFound)
}
