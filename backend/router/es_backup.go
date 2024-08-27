package router

// ES备份 路由
func (this *WebServer) runEsBackUp() {

	const AbsolutePath = "/api/backUp"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/SnapshotRepositoryListAction", this.esBackUpController.SnapshotRepositoryListAction)
		group.POST("/SnapshotCreateRepositoryAction", this.esBackUpController.SnapshotCreateRepositoryAction)
		group.POST("/SnapshotDeleteRepositoryAction", this.esBackUpController.SnapshotDeleteRepositoryAction)
		group.POST("/SnapshotListAction", this.esBackUpController.SnapshotListAction)
		group.POST("/CleanupeRepositoryAction", this.esBackUpController.CleanupeRepositoryAction)
		group.POST("/CreateSnapshotAction", this.esBackUpController.CreateSnapshotAction)
		group.POST("/SnapshotDeleteAction", this.esBackUpController.SnapshotDeleteAction)
		group.POST("/SnapshotDetailAction", this.esBackUpController.SnapshotDetailAction)
		group.POST("/SnapshotRestoreAction", this.esBackUpController.SnapshotRestoreAction)
		group.POST("/SnapshotStatusAction", this.esBackUpController.SnapshotStatusAction)

	}
}
