package router

// DslHistory 路由
func (this *WebServer) runDslHistory() {

	const AbsolutePath = "/api/dslHistory"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/ListAction", this.esController.DslHistoryAction)
		group.POST("/CleanAction", this.esController.CleanDslHistoryAction)
	}
}
