package router

// ES文档 路由
func (this *WebServer) runEsDoc() {

	const AbsolutePath = "/api/es_doc"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/DeleteRowByIDAction", this.esDocController.DeleteRowByIDAction)
		group.POST("/UpdateByIDAction", this.esDocController.UpdateByIDAction)
		group.POST("/InsertAction", this.esDocController.InsertAction)

	}
}
