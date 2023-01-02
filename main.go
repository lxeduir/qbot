package main

import (
	"edulx/web/db"
	"edulx/web/public"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		u := make(map[string]interface{})
		if err := c.BindJSON(&u); err != nil {
			panic(err)
			return
		}
		user := public.User{
			MessageType: u["message_type"].(string),
			MessageId:   int64(u["message_id"].(float64)),
			RawMessage:  u["raw_message"].(string),
			//GroupID:     int64(u["group_id"].(float64)),
			Sender: public.Sender{
				Age:      int(u["sender"].(map[string]interface{})["age"].(float64)),
				NickName: u["sender"].(map[string]interface{})["nickname"].(string),
				Sex:      u["sender"].(map[string]interface{})["sex"].(string),
				UserID:   int64(u["sender"].(map[string]interface{})["user_id"].(float64)),
			},
		}
		if user.MessageType == "group" {
			user.GroupID = int64(u["group_id"].(float64))
		}
		fmt.Println(user)
		user.Start()
	})
	ErrorApi(r)
	err := r.Run(":5701")
	if err != nil {
		return
	}
}
func ErrorApi(rack *gin.Engine) {
	rack.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg":   "来到了荒漠",
			"code":  200,
			"api":   "error",
			"error": "请检测api路径是否正确",
		})
	})
}
func init() {
	err := db.InitClirnt()
	if err != nil {
		return
	}
}
