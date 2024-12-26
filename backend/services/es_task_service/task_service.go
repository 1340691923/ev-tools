package es_task_service

import (
	"context"
	"ev-plugin/backend/dto"
	"ev-plugin/backend/vo"
	"github.com/goccy/go-json"

	"github.com/1340691923/eve-plugin-sdk-go/ev_api/pkg"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type EsTaskService struct{}

func NewEsTaskService() *EsTaskService {
	return &EsTaskService{}
}

func (this *EsTaskService) TaskList(ctx context.Context, esClient pkg.ClientInterface) (res map[string]vo.TaskInfo, err error) {

	resp, err := esClient.EsTaskList(ctx)
	if err != nil {
		return
	}
	if resp.StatusErr() != nil {
		err = resp.StatusErr()
		return
	}
	var tasksListResponse vo.Tasks
	err = json.Unmarshal(resp.ResByte(), &tasksListResponse)
	if err != nil {
		return
	}

	res = map[string]vo.TaskInfo{}

	for _, node := range tasksListResponse.Nodes {
		for taskId, taskInfo := range node.Tasks {
			res[taskId] = taskInfo
		}
	}

	return
}

func (this *EsTaskService) Cancel(ctx context.Context, esClient pkg.ClientInterface, cancelTask *dto.CancelTask) (err error) {
	res, err := esClient.EsTasksCancel(ctx, cancelTask.TaskID)
	if err != nil {
		return
	}

	if res.StatusErr() != nil {
		err = res.StatusErr()
		return
	}

	if gjson.GetBytes(res.ResByte(), "node_failures").Exists() {
		err = errors.New(string(res.ResByte()))
		return
	}

	return
}
