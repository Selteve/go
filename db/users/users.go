package UserDb

import (
	Config "gitee.com/under-my-umbrella/cloud/config"
	"gitee.com/under-my-umbrella/cloud/db"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
    "fmt"
)

var database *sqlx.DB

func init() {
    database = db.GetDB()
}

// CheckUser 查询用户是否存在
func CheckUser(username string) bool {
    var user Config.UserConfig
    query := "SELECT * FROM users WHERE username = ?"
    err := database.Get(&user, query, username)
    if err != nil {
        fmt.Println("查询发生错误:", err)
        return false
    } else {
       return true    
    }
}

// GetUserByUsername 查询用户id
func GetUserIdByUsername(username string) (int, error) {
    var user Config.UserConfig
    query := "SELECT * FROM users WHERE username = ?"
    err := database.Get(&user, query, username)
    if err != nil {
        return 0, err
    }
    return user.Id, nil
}

// CreateUser 保存用户到数据库
func CreateUser(username string, password []byte, userImg string) bool {
    query := "INSERT INTO users (username, password, user_img) VALUES (?, ?, ?)"
    _, err := database.Exec(query, username, password, userImg)
    return err == nil
}


// GetUser 查询用户信息并验证密码
func GetUser(username string, password []byte) bool {
    var user Config.UserConfig
    query := "SELECT * FROM users WHERE username = ?"
    err := database.Get(&user, query, username)
    if err != nil {
        return false
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil
}