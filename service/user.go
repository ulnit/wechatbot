/*
 * @Description:
 * @Version: 1.0
 * @Autor: Sean
 * @Date: 2023-02-20 14:48:32
 * @LastEditors: Sean
 * @LastEditTime: 2023-07-20 14:16:15
 */
package service

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/869413421/wechatbot/config"
	"github.com/patrickmn/go-cache"
)

// UserServiceInterface 用户业务接口
type UserServiceInterface interface {
	GetUserSessionContext(userId string) string
	SetUserSessionContext(userId string, question, reply string)
	ClearUserSessionContext(userId string, msg string) bool
}

var _ UserServiceInterface = (*UserService)(nil)

// UserService 用戶业务
type UserService struct {
	// 缓存
	cache *cache.Cache
}

// ClearUserSessionContext 清空GPT上下文，接收文本中包含`我要问下一个问题`，并且Unicode 字符数量不超过20就清空
func (s *UserService) ClearUserSessionContext(userId string, msg string) bool {
	if strings.Contains(msg, "我要问下一个问题") && utf8.RuneCountInString(msg) < 20 {
		s.cache.Delete(userId)
		return true
	}
	return false
}

// NewUserService 创建新的业务层
func NewUserService() UserServiceInterface {
	return &UserService{cache: cache.New(time.Second*config.LoadConfig().SessionTimeout, time.Minute*10)}
}

// GetUserSessionContext 获取用户会话上下文文本
func (s *UserService) GetUserSessionContext(userId string) string {
	sessionContext, ok := s.cache.Get(userId)
	if !ok {
		return ""
	}
	return sessionContext.(string)
}

// SetUserSessionContext 设置用户会话上下文文本，question用户提问内容，GPT回复内容
func (s *UserService) SetUserSessionContext(userId string, question, reply string) {
	value := question + "\n" + reply
	s.cache.Set(userId, value, time.Second*config.LoadConfig().SessionTimeout)
}
