package main

/**
	* @author: XiongHaiying
	* @date: 2024/8/16
	* @time: 2024年08月16日 15:40:40
	* @desc: 本模块用于自动生成代码,自动生成数据库表结构,自动生成代码，数据库名称需要手动创建*（store）
	*
	* @param:
	* @return:
**/

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	db "gitee.com/under-my-umbrella/cloud/db"
	_ "github.com/go-sql-driver/mysql"
)
// 全局变量
var database *sqlx.DB
// 初始化数据库
func init() {
	database = db.GetDB()
}

func main() {
	// 自动创建users表
	CreateUsersTable()
}

// 自动创建users表,字段为id,username,password,user_img, 字段类型为int,varchar(255),varchar(255),varchar(255),id为主键,自增,其他字段可为空
func CreateUsersTable() {
	sql := `CREATE TABLE IF NOT EXISTS users (
				id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
				username varchar(255) DEFAULT '' COMMENT '用户名',
				password varchar(255) DEFAULT '' COMMENT '<PASSWORD>',
				user_img varchar(255) DEFAULT '' COMMENT '用户头像',
				PRIMARY KEY (id)
		);`
	_, err := database.Exec(sql)
	if err != nil {
		fmt.Println("create table error", err)
	} else {
		fmt.Println("create table success")
	}
}


// 建立一个files表,
/*
 * @author: XiongHaiying
 * @date: 2024/8/16
 * @time: 2024年08月16日 15:43:37
 * @desc: 
	* @param:
	* @return:
 * @throws: id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文件ID',
 * @throws: file_name varchar(255) DEFAULT '' COMMENT '文件名称',
 * @throws: file_path varchar(255) DEFAULT '' COMMENT '文件路径',
 * @throws: PRIMARY KEY (id)
 */


func CreateFilesTable() {

}


