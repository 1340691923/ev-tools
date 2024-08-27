package router

// ES索引 路由
func (this *WebServer) runEsIndex() {

	const AbsolutePath = "/api/es_index"
	group := this.engine.Group(AbsolutePath)
	{
		group.POST("/DeleteAction", this.esIndexController.DeleteAction)
		group.POST("/CreateAction", this.esIndexController.CreateAction)
		group.POST("/GetSettingsAction", this.esIndexController.GetSettingsAction)
		group.POST("/IndexNamesAction", this.esIndexController.IndexNamesAction)
		group.POST("/ReindexAction", this.esIndexController.ReindexAction)
		group.POST("/GetAliasAction", this.esIndexController.GetAliasAction)
		group.POST("/MoveAliasToIndex", this.esIndexController.MoveAliasToIndex)
		group.POST("/AddAliasToIndex", this.esIndexController.AddAliasToIndex)
		group.POST("/BatchAddAliasToIndex", this.esIndexController.BatchAddAliasToIndex)
		group.POST("/RemoveAlias", this.esIndexController.RemoveAlias)
		group.POST("/GetSettingsInfoAction", this.esIndexController.GetSettingsInfoAction)
		group.POST("/StatsAction", this.esIndexController.StatsAction)
		group.POST("/IndexsCountAction", this.esIndexController.IndexsCountAction)
		group.POST("/Refresh", this.esIndexController.Refresh)
		group.POST("/Close", this.esIndexController.Close)
		group.POST("/Open", this.esIndexController.Open)
		group.POST("/CacheClear", this.esIndexController.CacheClear)
		group.POST("/Flush", this.esIndexController.Flush)
		group.POST("/Empty", this.esIndexController.Empty)
		group.POST("/Forcemerge", this.esIndexController.Forcemerge)
	}
}
