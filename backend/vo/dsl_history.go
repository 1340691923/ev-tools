package vo

import "time"

type DisHistoryListRes struct {
	Count int64              `json:"count"` //查询数据总条数
	List  []GmDslHistoryData `json:"list"`  //数据列表
}

type GmDslHistoryData struct {
	Id      int       `gorm:"column:id;primary_key;NOT NULL" json:"id"`                                        //id
	Uid     int       `gorm:"column:uid;default:0" json:"uid"`                                                 //用户id
	Method  string    `gorm:"column:method;default:" json:"method"`                                            //请求方法
	Path    string    `gorm:"column:path;default:" json:"path"`                                                //请求path
	Body    *string   `gorm:"column:body" json:"body"`                                                         //请求body
	Created time.Time `gorm:"column:created;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"created"` //操作时间
}
