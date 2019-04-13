package controller

import "identity/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string //用户登录名称
	Password	string //登录密码
	IsAdmin	string //是否是管理员
}


var users []User

func init() {

	admin := User{LoginName:"suxuefeng", Password:"123456", IsAdmin:"T"}
	user1 := User{LoginName:"xiaoming", Password:"123456", IsAdmin:"F"}
	user2 := User{LoginName:"xiaohong", Password:"123456", IsAdmin:"F"}
	user3 := User{LoginName:"xiaohei", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)

}

//判断是否是管理员
func isAdmin(cuser User) bool {
	if cuser.IsAdmin == "T" {
		return true
	}
	return false
}