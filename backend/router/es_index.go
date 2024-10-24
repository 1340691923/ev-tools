package router

// ES索引 路由
func (this *WebServer) runEsIndex() {

	const AbsolutePath = "/api/es_index"
	group := this.engine.Group("ES索引操作", AbsolutePath)
	{
		group.POST(false, "获取索引名列表", "/IndexNamesAction", this.esIndexController.IndexNamesAction)
		group.POST(false, "获取索引配置", "/GetSettingsAction", this.esIndexController.GetSettingsAction)
		group.POST(false, "获取所有索引配置", "/GetSettingsInfoAction", this.esIndexController.GetSettingsInfoAction)
		group.POST(false, "获取索引的Stats", "/StatsAction", this.esIndexController.StatsAction)
		group.POST(false, "得到所有的索引数量", "/IndexsCountAction", this.esIndexController.IndexsCountAction)

		group.POST(true, "删除索引", "/DeleteAction", this.esIndexController.DeleteAction)
		group.POST(true, "创建索引", "/CreateAction", this.esIndexController.CreateAction)
		group.POST(true, "进行重建索引", "/ReindexAction", this.esIndexController.ReindexAction)
		group.POST(true, "创建索引别名", "/GetAliasAction", this.esIndexController.GetAliasAction)
		group.POST(true, "迁移别名到索引", "/MoveAliasToIndex", this.esIndexController.MoveAliasToIndex)
		group.POST(true, "新增别名到索引", "/AddAliasToIndex", this.esIndexController.AddAliasToIndex)
		group.POST(true, "批量新增别名到索引", "/BatchAddAliasToIndex", this.esIndexController.BatchAddAliasToIndex)
		group.POST(true, "删除别名", "/RemoveAlias", this.esIndexController.RemoveAlias)
		group.POST(true, "刷新索引", "/Refresh", this.esIndexController.Refresh)
		group.POST(true, "关闭索引", "/Close", this.esIndexController.Close)
		group.POST(true, "开启索引", "/Open", this.esIndexController.Open)
		group.POST(true, "清理索引缓存", "/CacheClear", this.esIndexController.CacheClear)
		group.POST(true, "将所有索引刷新到磁盘", "/Flush", this.esIndexController.Flush)
		group.POST(true, "清空索引", "/Empty", this.esIndexController.Empty)
		group.POST(true, "强制合并索引", "/Forcemerge", this.esIndexController.Forcemerge)
	}
}
