package main

import(
	"strconv"
	"gitee.com/under-my-umbrella/cloud/router"
	_ "gitee.com/under-my-umbrella/cloud/db"
	Utils "gitee.com/under-my-umbrella/cloud/utils"
)

func main(){
	Port := Utils.ReadConfig().Port

	r := router.SetupRouter()
	r.Run(":" + strconv.Itoa(Port))
}