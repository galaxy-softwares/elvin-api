package v1

import (
	"dancin-api/model/request"
	"dancin-api/model/response"
	"dancin-api/services"

	"github.com/gin-gonic/gin"
)

func GetIssues(context *gin.Context) {
	var params request.RequestParams
	_ = context.BindQuery(&params)
	params.StartTime = params.StartTime + " 00:00:00"
	params.EndTime = params.EndTime + " 23:59:59"
	responses, err := services.GetIssues(params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}

func GetIssuesDetail(context *gin.Context) {
	var jsErrorParams request.JsErrorParams
	_ = context.BindQuery(&jsErrorParams)

	responses, err := services.GetIssuesDetail(jsErrorParams.IssueId, jsErrorParams.ErrorId, jsErrorParams.MonitorId)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OkWithDetailed(responses, "获取成功", context)
}
