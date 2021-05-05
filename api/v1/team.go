package v1

import (
	"danci-api/model"
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"danci-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// 获取团队列表
func GetTeamList(context *gin.Context) {
	claims, exists := context.Get("claims")
	if exists {
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
	}
	responses, _ := services.GetTeamList()
	response.OkWithDetailed(responses, "获取成功", context)
}

// 创建团队
func CreateTeam(context *gin.Context) {
	claims, exists := context.Get("claims")
	if exists {
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		name, isExist := context.GetPostForm("name")
		if isExist {
			admins, err := services.FindAdmins(customClaims.ID)
			if err == nil {
				team := &model.Team{
					Name:     name,
					AdminId:  customClaims.ID,
					NickName: customClaims.NickName,
					Admins:   admins,
				}
				if err := services.CreateTeam(*team); err != nil {
					response.FailWithMessage("创建团队失败！！", context)
				} else {
					response.OkWithMessage("创建团队成功！", context)
				}
			}
		}
	}
}

// 根据团队去创建项目
func AddTeamProject(context *gin.Context) {
	claims, exists := context.Get("claims")
	if exists {
		var addTeamProjectParams request.AddTeamProjectParams
		var customClaims request.CustomClaims
		utils.InterfaceToJsonToStruct(claims, &customClaims)
		_ = context.ShouldBind(&addTeamProjectParams)
		team, err := services.FindTeam(addTeamProjectParams.TeamId)
		if err != nil {
			response.FailWithMessage("没有查询到团队！", context)
		} else {
			projectModel := model.Project{
				ProjectName: addTeamProjectParams.ProjectName,
				ProjectType: addTeamProjectParams.ProjectType,
				Logo:        addTeamProjectParams.Logo,
				MonitorId:   addTeamProjectParams.MonitorId,
				AdminID:     customClaims.ID,
				TeamID:      team.ID,
			}
			project, err := services.CreateProject(projectModel)
			if err != nil {
				response.FailWithMessage("项目创建出错！", context)
			} else {
				response.OkWithDetailed(project, "项目创建成功！", context)
			}
		}
	}
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// 给团队绑定管理员
func BindTeamAdmins(context *gin.Context) {
	var teamAdminsParams request.BindTeamAdminsParams
	_ = context.ShouldBind(&teamAdminsParams)
	team, err := services.FindTeam(teamAdminsParams.TeamId)
	if err != nil {
		response.FailWithMessage("没有查询到团队！", context)
	} else {
		// 接受需要被绑定的用户ids
		admins, err := services.FindAdmins(teamAdminIds(&team, strings.Split(teamAdminsParams.AdminIds, ","))...)
		if err == nil {
			team.Admins = admins
			if err := services.UpdateTeamAdminIds(&team); err != nil {
				response.FailWithMessage("绑定团队成员失败！", context)
			} else {
				response.OkWithMessage("绑定团队成员成功！", context)
			}
		}
	}
}

func teamAdminIds(team *model.Team, paramAdminIds []string) (adminIds []uint) {
	adminIds = []uint{team.AdminId}
	for _, value := range paramAdminIds {
		adminIds = append(adminIds, StrToUInt(value))
	}
	return
}
