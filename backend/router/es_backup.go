package router

// ES备份 路由
func (this *WebServer) runEsBackUp() {

	const AbsolutePath = "/api/backUp"
	group := this.engine.Group("ES备份", AbsolutePath)
	{
		group.POST(false, "快照仓库列表", "/SnapshotRepositoryListAction", this.esBackUpController.SnapshotRepositoryListAction)
		group.POST(false, "删除快照仓库", "/SnapshotDeleteRepositoryAction", this.esBackUpController.SnapshotDeleteRepositoryAction)
		group.POST(false, "快照列表", "/SnapshotListAction", this.esBackUpController.SnapshotListAction)
		group.POST(false, "快照详情", "/SnapshotDetailAction", this.esBackUpController.SnapshotDetailAction)
		group.POST(false, "得到快照状态", "/SnapshotStatusAction", this.esBackUpController.SnapshotStatusAction)
		group.POST(true, "新建快照仓库", "/SnapshotCreateRepositoryAction", this.esBackUpController.SnapshotCreateRepositoryAction)
		group.POST(true, "清理快照仓库", "/CleanupeRepositoryAction", this.esBackUpController.CleanupeRepositoryAction)
		group.POST(true, "创建快照", "/CreateSnapshotAction", this.esBackUpController.CreateSnapshotAction)
		group.POST(true, "删除快照", "/SnapshotDeleteAction", this.esBackUpController.SnapshotDeleteAction)
		group.POST(true, "将索引恢复至快照时状态", "/SnapshotRestoreAction", this.esBackUpController.SnapshotRestoreAction)

	}
}
