package router

// DslHistory 路由
func (this *WebServer) runDslHistory() {

	const AbsolutePath = "/api/dslHistory"
	group := this.engine.Group("DSL历史记录", AbsolutePath)
	{
		group.POST(false, "查询自己的DSL历史记录", "/ListAction", this.esController.DslHistoryAction)
		group.POST(false, "清理自己的DSL历史记录", "/CleanAction", this.esController.CleanDslHistoryAction)
	}
}
