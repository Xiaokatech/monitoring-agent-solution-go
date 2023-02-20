package agentHttp

import (
	"sync"

	"github.com/JieanYang/HelloWorldGoAgent/src/agentHttp/agentHttpController"
	"github.com/gin-gonic/gin"
)

func StartHttp() {
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		router := gin.Default()

		router.GET("/", agentHttpController.HomeGetController)
		router.POST("/", agentHttpController.HomePostController)

		// RunCommand - with session key
		router.POST("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		router.POST("/RunCommandWithUrl", agentHttpController.RunCommandByUrl)

		// Run http web service
		router.Run(":9001")
		wg.Done()

		// http.HandleFunc("/", agentHttpController.HomeController)
		// // http.HandleFunc("/RunCommandWithFormData", agentHttpController.RunCommandWithFormData)
		// // http.HandleFunc("/RunCommandWithBody", agentHttpController.RunCommandWithBody)

		// // Auth
		// http.HandleFunc("/auth/authenticateByAuthKey", agentHttpController.HomeController)
		// http.HandleFunc("/auth/generateTransferKeyByAuthKey", agentHttpController.HomeController)
		// http.HandleFunc("/auth/generateSessionKeyByTransferKey", agentHttpController.HomeController)

		// // RunCommand - with session key
		// http.HandleFunc("/RunCommandByScriptContent", agentHttpController.RunCommandByScriptContent)
		// http.HandleFunc("/RunCommandWithUrl", agentHttpController.RunCommandByUrl)

		// err := http.ListenAndServe(":9001", nil) // Block code
		// if err != nil {
		// 	panic(err)
		// }
	}()

	wg.Wait()

}
