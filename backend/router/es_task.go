package router

// ES 任务 路由
func (this *WebServer) runEsTask() {

	const AbsolutePath = "/api/es_task"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/ListAction", this.esTaskController.ListAction)
		group.POST("/CancelAction", this.esTaskController.CancelAction)
	}
}
