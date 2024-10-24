package router

// ES 任务 路由
func (this *WebServer) runEsTask() {

	const AbsolutePath = "/api/es_task"
	group := this.engine.Group("ES异步任务", AbsolutePath)
	{
		group.POST(false, "Es任务列表", "/ListAction", this.esTaskController.ListAction)
		group.POST(true, "取消ES任务", "/CancelAction", this.esTaskController.CancelAction)
	}
}
