package router

// ES mapping 路由
func (this *WebServer) runEsMap() {

	group := this.engine.Group("/api/es_map")
	{
		group.POST("/ListAction", this.esMappingController.ListAction)
		group.POST("/UpdateMappingAction", this.esMappingController.UpdateMappingAction)
	}
}
