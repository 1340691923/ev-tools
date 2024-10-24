package router

// ES基础操作 路由
func (this *WebServer) runEsCrud() {

	const AbsolutePath = "/api/es_crud"
	group := this.engine.Group("ES可视化操作", AbsolutePath)
	{
		group.POST(false, "可视化筛选获取数据", "/GetList", this.esCrudController.GetList)
		group.POST(false, "可视化GetDSL", "/GetDSL", this.esCrudController.GetDSL)
		group.POST(false, "下载excel", "/Download", this.esCrudController.Download)
	}
}
