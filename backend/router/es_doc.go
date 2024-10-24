package router

// ES文档 路由
func (this *WebServer) runEsDoc() {

	const AbsolutePath = "/api/es_doc"
	group := this.engine.Group("ES文档操作", AbsolutePath)
	{
		group.POST(true, "删除文档数据", "/DeleteRowByIDAction", this.esDocController.DeleteRowByIDAction)
		group.POST(true, "修改文档", "/UpdateByIDAction", this.esDocController.UpdateByIDAction)
		group.POST(true, "新增文档", "/InsertAction", this.esDocController.InsertAction)

	}
}
