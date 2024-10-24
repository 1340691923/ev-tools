package router

// ES mapping 路由
func (this *WebServer) runEsMap() {

	group := this.engine.Group("ES映射操作", "/api/es_map")
	{
		group.POST(false, "查询映射列表", "/ListAction", this.esMappingController.ListAction)
		group.POST(true, "修改Es映射", "/UpdateMappingAction", this.esMappingController.UpdateMappingAction)
	}
}
