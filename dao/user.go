package dao

import (
	"go-blog/models"
)

func GetUserNameById(userId string) {

}

func GetUserIdByName(username string, password string) *models.User {
	//row := DB.Query("select * from goblog_user where username=? and password = ? limit 1", username, password)
	//if row.Err() != nil {
	//	log.Printf(row.Err().Error())
	//}
	var user = &models.User{}
	//err := row.Scan(&user.Uid, &user.Username, &user.Password, &user.Avatar)
	//if err != nil {
	//	return nil
	//}
	return user
}
