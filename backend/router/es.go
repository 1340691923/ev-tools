package router

// ES基础操作 路由
func (this *WebServer) runEs() {

	const AbsolutePath = "/api/es"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/RecoverCanWrite", this.esController.RecoverCanWrite)
		group.POST("/CatAction", this.esController.CatAction)
		group.POST("/RunDslPOSTAction", this.esController.RunDslPOSTAction)
		group.POST("/RunDslGETAction", this.esController.RunDslGETAction)
		group.POST("/RunDslPUTAction", this.esController.RunDslPUTAction)
		group.POST("/RunDslDELETEAction", this.esController.RunDslDELETEAction)
		group.POST("/RunDslHEADAction", this.esController.RunDslHEADAction)

		group.POST("/SqlToDslAction", this.esController.SqlToDslAction)

	}
}
