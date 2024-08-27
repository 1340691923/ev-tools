package router

// ES基础操作 路由
func (this *WebServer) runEsCrud() {

	const AbsolutePath = "/api/es_crud"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/GetList", this.esCrudController.GetList)
		group.POST("/GetDSL", this.esCrudController.GetDSL)
		group.POST("/Download", this.esCrudController.Download)
	}
}
