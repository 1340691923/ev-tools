package navicat_service

import (
	"ev-plugin/backend/dto"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type esI interface {
	ToSql() elastic.Query
	Append(sqlizer elastic.Query)
}

type Or struct {
	or []elastic.Query
}

func (this *Or) Append(sqlizer elastic.Query) {
	this.or = append(this.or, sqlizer)
}

func (this *Or) ToSql() elastic.Query {
	return elastic.NewBoolQuery().Should(this.or...)
}

type And struct {
	and []elastic.Query
}

func (this *And) Append(sqlizer elastic.Query) {
	this.and = append(this.and, sqlizer)
}

func (this *And) ToSql() elastic.Query {
	return elastic.NewBoolQuery().Must(this.and...)
}

const COMPOUND = "COMPOUND"
const SIMPLE = "SIMPLE"
const AND = "且"
const OR = "或"

func GetWhereSql(anlysisFilter dto.AnalysisFilter) (q elastic.Query, err error) {
	var arrP esI
	colArr := []string{}
	switch anlysisFilter.Relation {
	case AND:
		arrP = &And{}
	case OR:
		arrP = &Or{}
	default:
		return nil, errors.New("错误的连接类型:" + anlysisFilter.Relation)
	}

	for _, v := range anlysisFilter.Filts {
		if v.FilterType == SIMPLE {
			colArr = append(colArr, v.ColumnName)
			expr := getExpr(v.ColumnName, v.Comparator, v.Ftv)
			if expr == nil {
				return nil, errors.New("参数错误")
			}
			arrP.Append(expr)
		} else {
			var arrC esI
			switch v.Relation {
			case AND:
				arrC = &And{}
			case OR:
				arrC = &Or{}
			default:
				return nil, errors.New("错误的连接类型")
			}

			for _, v2 := range v.Filts {
				colArr = append(colArr, v2.ColumnName)
				expr := getExpr(v2.ColumnName, v2.Comparator, v2.Ftv)
				if expr == nil {
					return nil, errors.New("参数错误")
				}
				arrC.Append(expr)
			}
			arrP.Append(arrC.ToSql())
		}
	}

	return arrP.ToSql(), err
}

func getExpr(columnName, comparator string, ftv interface{}) elastic.Query {

	if comparator == "query" {
		return elastic.NewRangeQuery(columnName).Gte(ftv.([]interface{})[0]).Lte(ftv.([]string)[1])
	}

	if comparator == "=" || comparator == "@=" {
		if comparator == "@=" {
			columnName = fmt.Sprintf("%s.keyword", columnName)
		}
		switch ftv.(type) {
		case string:
			return elastic.NewTermsQuery(columnName, ftv.(string))
		case []interface{}:
			return elastic.NewTermsQuery(columnName, ftv.([]interface{})...)
		case []string:

			interfaceList := []interface{}{}
			for _, v := range ftv.([]string) {
				interfaceList = append(interfaceList, v)
			}
			return elastic.NewTermsQuery(columnName, interfaceList...)
		}
	}

	if comparator == "!=" || comparator == "@!=" {
		if comparator == "@!=" {
			columnName = fmt.Sprintf("%s.keyword", columnName)
		}
		switch ftv.(type) {
		case string:
			bq := elastic.NewBoolQuery()
			return bq.MustNot(elastic.NewTermsQuery(columnName, ftv.(string)))
		case []interface{}:
			bq := elastic.NewBoolQuery()
			return bq.MustNot(elastic.NewTermsQuery(columnName, ftv.([]interface{})...))
		case []string:
			bq := elastic.NewBoolQuery()
			interfaceList := []interface{}{}
			for _, v := range ftv.([]string) {
				interfaceList = append(interfaceList, v)
			}
			return bq.MustNot(elastic.NewTermsQuery(columnName, interfaceList...))
		}
	}

	if comparator == ">" {
		switch ftv.(type) {
		case string:
			return elastic.NewRangeQuery(columnName).Gt(ftv.(string))
		}
		return elastic.NewRangeQuery(columnName).Gt(ftv.([]string)[0])
	}

	if comparator == ">=" {
		switch ftv.(type) {
		case string:
			return elastic.NewRangeQuery(columnName).Gte(ftv.(string))
		}
		return elastic.NewRangeQuery(columnName).Gte(ftv.([]string)[0])
	}

	if comparator == "<=" {
		switch ftv.(type) {
		case string:
			return elastic.NewRangeQuery(columnName).Lte(ftv.(string))
		}
		return elastic.NewRangeQuery(columnName).Lte(ftv.([]string)[0])
	}

	if comparator == "<" {
		switch ftv.(type) {
		case string:
			return elastic.NewRangeQuery(columnName).Lt(ftv.(string))
		}
		return elastic.NewRangeQuery(columnName).Lt(ftv.([]string)[0])
	}

	if comparator == "textMatch" {

		switch ftv.(type) {
		case string:
			return elastic.NewMatchQuery(columnName, cast.ToString(ftv))
		case []interface{}:
			return elastic.NewMatchQuery(columnName, cast.ToString(ftv.([]interface{})[0]))
		case []string:
			return elastic.NewMatchQuery(columnName, cast.ToString(ftv.([]string)[0]))
		}
	}

	if comparator == "match" || comparator == "@match" {
		if comparator == "@match" {
			columnName = fmt.Sprintf("%s", columnName)
		}
		switch ftv.(type) {
		case string:
			return elastic.NewWildcardQuery(columnName, "*"+cast.ToString(ftv)+"*")
		case []interface{}:
			return elastic.NewWildcardQuery(columnName, "*"+cast.ToString(ftv.([]interface{})[0])+"*")
		case []string:
			return elastic.NewWildcardQuery(columnName, "*"+cast.ToString(ftv.([]string)[0])+"*")
		}
	}

	if comparator == "isNotNull" {
		bq := elastic.NewBoolQuery()
		return bq.Must(elastic.NewExistsQuery(columnName))
	}

	if comparator == "isNull" {
		bq := elastic.NewBoolQuery()
		return bq.MustNot(elastic.NewExistsQuery(columnName))
	}

	return nil
}
