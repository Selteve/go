package config

/**
	* @author: XiongHaiying
	* @date: 2024/8/16
	* @time: 2024年08月16日 15:40:40
	* @desc: 本模块用于整个项目的配置文件
	*
	* @param:
	* @return:
**/

// 定义用户配置结构体
type UserConfig struct {
	Id       int    `form:"id" db:"id" json:"id"`
	Username string `form:"username" db:"username" json:"username"`
	Password string `form:"password" db:"password" json:"password"`
	UserImg  string `form:"user_img" db:"user_img" json:"user_img"`
}


