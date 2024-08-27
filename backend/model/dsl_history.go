// 数据库实体层
package model

import (
	"context"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api"
	"github.com/1340691923/eve-plugin-sdk-go/sql_builder"
	"time"
)

/*
	http://sql2struct.atotoa.com 根据表结构生成 go结构体

GmDslHistoryModel DSL历史记录
*/
type GmDslHistoryModel struct {
	ID         int      `gorm:"column:id" json:"id" form:"id" db:"id"`
	Uid        int      `gorm:"column:uid" json:"uid" form:"uid" db:"uid"`
	Method     string   `gorm:"column:method" json:"method" form:"method" db:"method"`
	Path       string   `gorm:"column:path" json:"path" form:"path" db:"path"`
	Body       string   `gorm:"column:body" json:"body" form:"body" db:"body"`
	Created    string   `gorm:"column:created" json:"created" form:"created" db:"created"`
	FilterDate []string `json:"date"`
	IndexName  string   `json:"indexName"`
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
}

// 表名
func (this *GmDslHistoryModel) TableName() string {
	return "gm_dsl_history"
}

// Insert
func (this *GmDslHistoryModel) Insert() (err error) {
	sql, args := sql_builder.SqlBuilder.
		Insert(this.TableName()).
		SetMap(map[string]interface{}{
			"uid":     this.Uid,
			"method":  this.Method,
			"path":    this.Path,
			"body":    this.Body,
			"created": time.Now().Format("2006-01-02 15:04:05"),
		}).MustSql()
	_, err = ev_api.GetEvApi().StoreExec(context.Background(), sql, args...)
	if err != nil {
		return
	}
	return
}

// Clean
func (this *GmDslHistoryModel) Clean() (err error) {
	sql, args := sql_builder.SqlBuilder.
		Delete(this.TableName()).
		Where(sql_builder.Eq{"uid": this.Uid}).
		MustSql()
	_, err = ev_api.GetEvApi().StoreExec(context.Background(), sql, args...)
	if err != nil {
		return
	}
	return
}

// List
func (this *GmDslHistoryModel) List() (gmDslHistoryModelList []GmDslHistoryModel, err error) {
	builder := sql_builder.SqlBuilder.
		Select("*").
		From(this.TableName()).
		Where(sql_builder.Eq{"uid": this.Uid}).
		OrderBy("id desc").
		Limit(uint64(this.Limit)).Offset(sql_builder.CreatePage(this.Page, this.Limit))

	if this.IndexName != "" {
		builder = builder.Where(sql_builder.Like{"path": sql_builder.CreateLike(this.IndexName)})
	}

	if len(this.FilterDate) > 0 {
		builder = builder.Where(sql_builder.Gte{"created": this.FilterDate[0]}).Where(sql_builder.Lte{"created": this.FilterDate[1]})
	}

	SQL, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}
	err = ev_api.GetEvApi().StoreSelect(context.Background(), &gmDslHistoryModelList, SQL, args...)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	return gmDslHistoryModelList, nil
}

// Count
func (this *GmDslHistoryModel) Count() (count int, err error) {
	builder := sql_builder.SqlBuilder.
		Select("count(*) c").
		From(this.TableName()).
		Where(sql_builder.Eq{"uid": this.Uid})

	if this.IndexName != "" {
		builder = builder.Where(sql_builder.Like{"path": sql_builder.CreateLike(this.IndexName)})
	}

	if len(this.FilterDate) > 0 {
		builder = builder.Where(sql_builder.Gte{"created": this.FilterDate[0]}).Where(sql_builder.Lte{"created": this.FilterDate[1]})
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return
	}

	type Cnt struct {
		C int `json:"c"`
	}
	var cnt Cnt

	err = ev_api.GetEvApi().StoreFirst(context.Background(), &cnt, sql, args...)

	if err != nil {
		return
	}

	count = cnt.C

	return
}
