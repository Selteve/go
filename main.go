package main

import(
	"gitee.com/under-my-umbrella/cloud/router"
	_ "gitee.com/under-my-umbrella/cloud/db"
)

func main(){
	r := router.SetupRouter()
	r.Run(":8000")
}