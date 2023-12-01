package deviceno

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用gin框架实现一个restful接口
func StartHttp(path, listenAddr, prefix string) {
	//1.创建一个gin的实例
	//2.创建一个路由
	//3.创建一个handler
	//4.把handler注册到路由上
	//5.启动http服务
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/deviceno", func(c *gin.Context) {
		GenerateDeviceNoWithPrefix(prefix, path)
		dno, err := ReadFile(path)
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		dnos := DeviceNo{
			DeviceNo: dno,
			Error:    errMsg,
		}
		c.JSON(http.StatusOK, dnos)
	})
	r.GET("/netinfo", func(c *gin.Context) {
		netinfo := GetNetInfo()
		c.JSON(http.StatusOK, netinfo)
	})
	r.Run(listenAddr)
}

type DeviceNo struct {
	DeviceNo string `json:"device_no"`
	Error    string `json:"error"`
}
