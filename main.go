package main

import (
	"fmt"
	"love-recycling-go/bootstrap"
	"love-recycling-go/global"
	"love-recycling-go/models"
)

func main() {
	bootstrap.InitializeConfig()
	global.App.Log = bootstrap.InitializeLog()
	global.App.DB = bootstrap.InitializeDB()
	global.App.Redis = bootstrap.InitializeRedis()

	var sysUser models.SysUser

	global.App.DB.First(&sysUser)

	fmt.Println(sysUser)
	global.App.Log.Info("log init success!")

	bootstrap.RunServer()

}
