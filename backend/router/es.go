package router

// ES基础操作 路由
func (this *WebServer) runEs() {

	const AbsolutePath = "/api/es"
	group := this.engine.Group("ES基本操作", AbsolutePath)
	{
		group.POST(true, "将集群恢复成可读写状态", "/RecoverCanWrite", this.esController.RecoverCanWrite)
		group.POST(false, "cat操作", "/CatAction", this.esController.CatAction)
		group.POST(true, "执行DSL请求（POST）", "/RunDslPOSTAction", this.esController.RunDslPOSTAction)
		group.POST(true, "执行DSL请求（GET）", "/RunDslGETAction", this.esController.RunDslGETAction)
		group.POST(true, "执行DSL请求（PUT）", "/RunDslPUTAction", this.esController.RunDslPUTAction)
		group.POST(true, "执行DSL请求（DELETE）", "/RunDslDELETEAction", this.esController.RunDslDELETEAction)
		group.POST(true, "执行DSL请求（HEAD）", "/RunDslHEADAction", this.esController.RunDslHEADAction)
		group.POST(true, "进行SQL转DSL", "/SqlToDslAction", this.esController.SqlToDslAction)

	}
}
