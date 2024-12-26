package dto

type DslHistoryListReq2 struct {
	IndexName string   `json:"indexName"` // 索引名
	Date      []string `json:"date"`      //开始时间与结束时间（格式：”年-月-日 时:分:秒“ ）
	Page      int      `json:"page"`      //拉取数据当前页
	Limit     int      `json:"limit"`     //拉取条数
}
