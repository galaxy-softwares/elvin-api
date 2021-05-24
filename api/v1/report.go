package v1

import (
	"danci-api/global"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var handles *utils.Handles

func init() {
	handles = utils.NewHandles()
	routerHandles := map[string]utils.RouterFunc{
		"PAGE_LOAD": func(context *gin.Context) {
			var performanceBody request.PerformanceBody
			err := context.ShouldBindJSON(&performanceBody)
			performanceBody.IP = context.ClientIP()
			reportProducer(context, performanceBody, err)
			return
		},
		"HTTP_LOG": func(context *gin.Context) {
			var httpBody request.HttpBody
			err := context.ShouldBindJSON(&httpBody)
			httpBody.IP = context.ClientIP()
			reportProducer(context, httpBody, err)
			return
		},
		"PAGE_VIEW": func(context *gin.Context) {
			var pageViewBody request.PageViewBody
			err := context.ShouldBindJSON(&pageViewBody)
			pageViewBody.IP = context.ClientIP()
			reportProducer(context, pageViewBody, err)
			return
		},
		"OPERATION": func(context *gin.Context) {
			var operationBody request.OperationBody
			err := context.ShouldBindJSON(&operationBody)
			operationBody.IP = context.ClientIP()
			reportProducer(context, operationBody, err)
			return
		},
		"RESOURCE": func(context *gin.Context) {
			var resourceBody request.ResourceErrorBody
			err := context.ShouldBindJSON(&resourceBody)
			resourceBody.IP = context.ClientIP()
			reportProducer(context, resourceBody, err)
			return
		},
		"JS_ERROR": func(context *gin.Context) {
			var issuesBody request.IssuesBody
			err := context.ShouldBindJSON(&issuesBody)
			issuesBody.IP = context.ClientIP()
			reportProducer(context, issuesBody, err)
			return
		},
	}
	handles.RoutersHandlerRegister(routerHandles)
}
func Report(context *gin.Context) {
	reportBody := &request.ReportBody{}
	reportBody.ActionType = context.Query("action_type")
	handles.RouterHandlers[reportBody.ActionType](context)
}

func reportProducer(context *gin.Context, body interface{}, err error) {
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(body)
	global.GVA_REDIS.LPush("reportData", res)
	response.Ok(context)
	return
}

// 改造前

func CreatePagePerformance(context *gin.Context) {
	var performanceBody request.PerformanceBody
	err := context.ShouldBindJSON(&performanceBody)
	performanceBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(performanceBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := performanceBody.HappenDay + performanceBody.MonitorId + performanceBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreateHttpInfo(context *gin.Context) {
	var httpBody request.HttpBody
	err := context.BindJSON(&httpBody)
	httpBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(httpBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := httpBody.HappenDay + httpBody.MonitorId
	if httpBody.Status >= 400 {
		incrKey = "HTTP_ERROR_LOG"
	} else {
		incrKey = "HTTP_LOG"
	}
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreateResourcesError(context *gin.Context) {
	var resourceErrorBody request.ResourceErrorBody
	err := context.BindJSON(&resourceErrorBody)
	resourceErrorBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(resourceErrorBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := resourceErrorBody.HappenDay + resourceErrorBody.MonitorId + resourceErrorBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreatePageOperation(context *gin.Context) {
	var operationBody request.OperationBody
	err := context.BindJSON(&operationBody)
	operationBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(operationBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := operationBody.HappenDay + operationBody.MonitorId + operationBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreatePageIssues(context *gin.Context) {
	var issuesBody request.IssuesBody
	err := context.BindJSON(&issuesBody)
	issuesBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(issuesBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := issuesBody.HappenDay + issuesBody.MonitorId + issuesBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}

func CreatePageView(context *gin.Context) {
	var pageViewBody request.PageViewBody
	err := context.BindJSON(&pageViewBody)
	pageViewBody.IP = context.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	res, _ := json.Marshal(pageViewBody)
	global.GVA_REDIS.LPush("reportData", res)
	incrKey := pageViewBody.HappenDay + pageViewBody.MonitorId + pageViewBody.ActionType
	global.GVA_REDIS.Incr(incrKey)
	response.Ok(context)
}
