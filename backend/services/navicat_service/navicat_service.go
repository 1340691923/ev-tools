package navicat_service

import (
	"context"
	"encoding/csv"
	"ev-tools/backend/global"
	"ev-tools/backend/util"
	"fmt"
	"github.com/1340691923/eve-plugin-sdk-go/backend/logger"
	"github.com/1340691923/eve-plugin-sdk-go/ev_api/dto"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	proto2 "github.com/1340691923/eve-plugin-sdk-go/ev_api/proto"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"sort"
)

type NavicatService struct{}

func NewNavicatService() *NavicatService {
	return &NavicatService{}
}

func (this *NavicatService) CrudGetList(ctx context.Context, esClient pkg.ClientInterface, crudFilter *dto.CrudFilter) (res json.RawMessage, count int64, err error) {
	q, err := GetWhereSql(crudFilter.Relation)
	if err != nil {
		return
	}

	queryBody, err := q.Source()
	if err != nil {
		return
	}
	req := proto2.SearchRequest{
		Index: []string{crudFilter.IndexName},
	}

	sortArr := []map[string]interface{}{}
	for _, tmp := range crudFilter.SortList {
		sortArr = append(sortArr, map[string]interface{}{tmp.Col: tmp.SortRule})
	}

	searchBody := Search{
		Query: queryBody,
		From:  int(CreatePage(crudFilter.Page, crudFilter.Limit)),
		Size:  crudFilter.Limit,
		Sort:  sortArr,
	}

	resp, err := esClient.EsSearch(
		ctx, req, searchBody,
	)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	res = resp.JsonRawMessage()
	version, err := esClient.EsVersion()
	if err != nil {
		return
	}

	switch version {
	case 6:
		count = gjson.GetBytes(resp.ResByte(), "hits.total").Int()
	default:
		count = gjson.GetBytes(resp.ResByte(), "hits.total.value").Int()
	}

	return
}

func (this *NavicatService) CrudGetDSL(ctx context.Context, crudFilter *dto.CrudFilter) (res Search, err error) {
	q, err := GetWhereSql(crudFilter.Relation)
	if err != nil {
		return
	}

	queryBody, err := q.Source()
	if err != nil {
		return
	}
	sortArr := []map[string]interface{}{}
	for _, tmp := range crudFilter.SortList {
		sortArr = append(sortArr, map[string]interface{}{tmp.Col: tmp.SortRule})
	}

	res = Search{
		Query: queryBody,
		From:  int(CreatePage(crudFilter.Page, crudFilter.Limit)),
		Size:  crudFilter.Limit,
		Sort:  sortArr,
	}

	return
}

func (this *NavicatService) CrudDownload(ctx context.Context, esClient pkg.ClientInterface, filter *dto.CrudFilter) (downloadFileName string, titleList []string, searchData [][]string, err error) {

	mappingRes, err := esClient.EsGetMapping(ctx, []string{filter.IndexName})
	if err != nil {
		logger.DefaultLogger.Error("err", err)
		return
	}
	if mappingRes.StatusErr() != nil {
		err = mappingRes.StatusErr()
		return
	}

	fields := map[string]interface{}{}

	err = json.Unmarshal(mappingRes.ResByte(), &fields)

	if err != nil {
		logger.DefaultLogger.Error("err", err)
		return
	}

	fieldsArr := []string{"_index", "_type", "_id"}
	data, ok := fields[filter.IndexName].(map[string]interface{})

	if !ok {
		err = errors.New("该索引没有映射结构")
		return
	}
	propertiesArr := []string{}
	properties := map[string]interface{}{}
	mappings, ok := data["mappings"].(map[string]interface{})
	if !ok {
		err = errors.New("该索引没有映射结构")
		return
	}
	version, err := esClient.EsVersion()
	if err != nil {
		logger.DefaultLogger.Error("err", err)
		return
	}

	switch version {
	case 6:

		typeName := ""

		for key := range mappings {
			typeName = key
		}

		typeObj := mappings[typeName].(map[string]interface{})

		properties, ok = typeObj["properties"].(map[string]interface{})
		if !ok {
			err = errors.New("该索引没有映射结构")
			return
		}

	default:

		properties, ok = mappings["properties"].(map[string]interface{})
		if !ok {
			err = errors.New("该索引没有映射结构")
			return
		}

	}

	for key := range properties {
		propertiesArr = append(propertiesArr, key)
	}

	sort.Strings(propertiesArr)
	fieldsArr = append(fieldsArr, propertiesArr...)
	q, err := GetWhereSql(filter.Relation)
	if err != nil {
		logger.DefaultLogger.Error("err", err)
		return
	}

	querySource, err := q.Source()

	if err != nil {
		logger.DefaultLogger.Error("err", err)
		return
	}

	searchBody := Search{
		Query: querySource,
		Size:  8000,
		Sort: []map[string]interface{}{
			{
				"_id": "desc",
			},
		},
	}

	resp, err := esClient.EsSearch(
		ctx, proto2.SearchRequest{Index: []string{filter.IndexName}}, searchBody,
	)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}

	histsArr := gjson.GetBytes(resp.ResByte(), "hits.hits").Array()

	lastIdArr := histsArr[len(histsArr)-1].Get("sort").Array()

	llist := [][]string{}

	flushHitsDataFn := func(hits []gjson.Result) {

		for _, data := range hits {

			list := []string{data.Get("_index").String(), data.Get("_type").String(), data.Get("_id").String()}

			m := data.Get("_source").Map()

			for _, field := range fieldsArr {
				if field == "_index" || field == "_type" || field == "_id" {
					continue
				}
				if value, ok := m[field]; ok {
					list = append(list, ToExcelData(value.Value()))
				} else {
					list = append(list, "")
				}
			}

			llist = append(llist, list)
		}
	}

	flushHitsDataFn(histsArr)
	haveData := true

	for haveData {
		searchAfter := []interface{}{}
		for _, v := range lastIdArr {
			searchAfter = append(searchAfter, v.Value())
		}

		searchBody := Search{
			Query: querySource,
			Size:  8000,
			Sort: []map[string]interface{}{
				{
					"_id": "desc",
				},
			},
			SearchAfter: &searchAfter,
		}
		var searchResp *proto2.Response

		searchResp, err = esClient.EsSearch(ctx, proto2.SearchRequest{Index: []string{filter.IndexName}}, searchBody)

		if err != nil {
			logger.DefaultLogger.Error("err", err)
			return
		}

		if searchResp.StatusErr() != nil {
			err = searchResp.StatusErr()
			return
		}

		hitsArr := gjson.GetBytes(searchResp.ResByte(), "hits.hits").Array()

		if len(hitsArr) == 0 {
			break
		}

		lastIdArr = hitsArr[len(hitsArr)-1].Get("sort").Array()
		flushHitsDataFn(hitsArr)
	}
	downloadFileName = "test"
	titleList = fieldsArr
	searchData = llist

	return

}

func (this *NavicatService) DownloadExcel(downloadFileName string, titleList []string, data [][]string, ctx *gin.Context) (err error) {

	csvFile := fmt.Sprintf("%s.csv", time.Now().Format("20060102150405"))

	var downloadUrl = filepath.Join(global.GetTmpFileStorePath(), csvFile)

	if !util.CheckFileIsExist(downloadUrl) {
		os.Create(downloadUrl)
	}

	file, err := os.OpenFile(downloadUrl, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
		return
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write(titleList)

	for _, d := range data {
		w.Write(d)
	}
	w.Flush()

	f, err := os.Open(downloadUrl)
	if err != nil {
		logger.DefaultLogger.Error("os.Open failed:", err)
		return
	}
	defer f.Close()

	// 将文件读取出来
	filedata, err := io.ReadAll(f)
	if err != nil {
		logger.DefaultLogger.Error("io.ReadAll failed:", err)
		return
	}
	ctx.Header("Content-Disposition", `attachment; filename="`+downloadFileName+`.xlsx"`)
	ctx.Writer.Write(filedata)

	return
}

func CreatePage(page, limit int) int {
	tmp := (page - 1) * limit
	return int(tmp)
}

func ToExcelData(any interface{}) string {

	if any == nil {
		return ""
	}
	switch value := any.(type) {
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.Itoa(int(value))
	case int16:
		return strconv.Itoa(int(value))
	case int32:
		return strconv.Itoa(int(value))
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		if value.IsZero() {
			return ""
		}
		return value.String()
	case *time.Time:
		if value == nil {
			return ""
		}
		return value.String()
	case []int:

		b, _ := json.Marshal(value)
		return string(b)
	case []int32:

		b, _ := json.Marshal(value)
		return string(b)
	case []int16:

		b, _ := json.Marshal(value)
		return string(b)
	case []int8:

		b, _ := json.Marshal(value)
		return string(b)
	case []int64:

		b, _ := json.Marshal(value)
		return string(b)
	case []float64:

		b, _ := json.Marshal(value)
		return string(b)
	case []float32:

		b, _ := json.Marshal(value)
		return string(b)
	case []uint64:

		b, _ := json.Marshal(value)
		return string(b)
	case []uint16:

		b, _ := json.Marshal(value)
		return string(b)
	case []string:

		b, _ := json.Marshal(value)
		return string(b)
	case []interface{}:

		b, _ := json.Marshal(value)
		return string(b)
	case map[string]interface{}:

		b, _ := json.Marshal(value)
		return string(b)
	default:
		return ""
	}
}

func StringPtr(v string) *string {
	return &v
}

func IntPtr(v int) *int {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}
