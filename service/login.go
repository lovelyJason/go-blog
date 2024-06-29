package service

import (
	"errors"
	"go-blog/models"
	"go-blog/utils"
	"log"
)

func Login(username, password string) (*models.LoginResponse, error) {
	// 业务上应该是获取到uid，生成token
	uid := 101
	/*
	   注意：jwt a.b.c:a -> 加密规则， b -> 数据包括过期时间，c -> 对a和b进行加密
	*/
	token, err := utils.Award(&uid)
	log.Printf(token, err)
	if err != nil {
		// 注意：error是类型，errors是标准模块
		return nil, errors.New("token生成失败")
	}
	var userInfo models.UserInfo
	userInfo.Uid = "1"
	userInfo.Username = "jason"
	userInfo.Avatar = ""
	var res = &models.LoginResponse{
		Token:    token,
		UserInfo: userInfo,
	}
	return res, nil
}
