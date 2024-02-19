package initialize

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
	"watchAlert/controllers/repo"
	"watchAlert/globals"
	"watchAlert/models"
	"watchAlert/utils/cmd"
)

var perms []models.UserPermissions

func InitPermissionsSQL() {

	permissions := []models.UserPermissions{
		{
			Key: "UserRegister",
			API: "/api/system/register",
		},
		{
			Key: "UserList",
			API: "/api/w8t/user/userList",
		},
		{
			Key: "UserUpdate",
			API: "/api/w8t/user/userUpdate",
		},
		{
			Key: "UserDelete",
			API: "/api/w8t/user/userDelete",
		},
		{
			Key: "UserChangePass",
			API: "/api/w8t/user/userChangePass",
		},
		{
			Key: "SearchDutyUser",
			API: "/api/w8t/user/searchDutyUser",
		},
		{
			Key: "RoleCreate",
			API: "/api/w8t/role/roleCreate",
		},
		{
			Key: "RoleUpdate",
			API: "/api/w8t/role/roleUpdate",
		},
		{
			Key: "RoleDelete",
			API: "/api/w8t/role/roleDelete",
		},
		{
			Key: "RoleList",
			API: "/api/w8t/role/roleList",
		},
		{
			Key: "SilenceCreate",
			API: "/api/w8t/silence/silenceCreate",
		},
		{
			Key: "SilenceUpdate",
			API: "/api/w8t/silence/silenceUpdate",
		},
		{
			Key: "SilenceDelete",
			API: "/api/w8t/silence/silenceDelete",
		},
		{
			Key: "SilenceList",
			API: "/api/w8t/silence/silenceList",
		},
		{
			Key: "RuleCreate",
			API: "/api/w8t/rule/ruleCreate",
		},
		{
			Key: "RuleUpdate",
			API: "/api/w8t/rule/ruleUpdate",
		},
		{
			Key: "RuleDelete",
			API: "/api/w8t/rule/ruleDelete",
		},
		{
			Key: "RuleList",
			API: "/api/w8t/rule/ruleList",
		},
		{
			Key: "RuleSearch",
			API: "/api/w8t/rule/RuleSearch",
		},
		{
			Key: "DutyManageCreate",
			API: "/api/w8t/dutyManage/dutyManageCreate",
		},
		{
			Key: "DutyManageUpdate",
			API: "/api/w8t/dutyManage/dutyManageUpdate",
		},
		{
			Key: "DutyManageDelete",
			API: "/api/w8t/dutyManage/dutyManageUelete",
		},
		{
			Key: "DutyManageList",
			API: "/api/w8t/dutyManage/dutyManageList",
		},
		{
			Key: "DutyManageSearch",
			API: "/api/w8t/dutyManage/dutyManageSearch",
		},
		{
			Key: "DutyScheduleCreate",
			API: "/api/w8t/calendar/calendarCreate",
		},
		{
			Key: "DutyScheduleUpdate",
			API: "/api/w8t/calendar/calendarUpdate",
		},
		{
			Key: "DutyScheduleSearch",
			API: "/api/w8t/calendar/calendarSearch",
		},
		{
			Key: "NoticeCreate",
			API: "/api/w8t/notice/noticeCreate",
		},
		{
			Key: "NoticeUpdate",
			API: "/api/w8t/notice/noticeUpdate",
		},
		{
			Key: "NoticeDelete",
			API: "/api/w8t/notice/noticeDelete",
		},
		{
			Key: "NoticeList",
			API: "/api/w8t/notice/noticeList",
		},
		{
			Key: "NoticeSearch",
			API: "/api/w8t/notice/noticeSearch",
		},
		{
			Key: "DataSourceCreate",
			API: "/api/w8t/datasource/dataSourceCreate",
		},
		{
			Key: "DataSourceUpdate",
			API: "/api/w8t/datasource/dataSourceUpdate",
		},
		{
			Key: "DataSourceDelete",
			API: "/api/w8t/datasource/dataSourceDelete",
		},
		{
			Key: "DataSourceList",
			API: "/api/w8t/datasource/dataSourceList",
		},
		{
			Key: "DataSourceSearch",
			API: "/api/w8t/datasource/dataSourceSearch",
		},
		{
			Key: "CurrentEventList",
			API: "/api/w8t/event/curEvent",
		},
		{
			Key: "HistoryEventList",
			API: "/api/w8t/event/hisEvent",
		},
		{
			Key: "PermissionsList",
			API: "/api/w8t/permissions/permsList",
		},
		{
			Key: "NoticeTemplateList",
			API: "/api/w8t/noticeTemplate/noticeTemplateList",
		},
		{
			Key: "NoticeTemplateCreate",
			API: "/api/w8t/noticeTemplate/noticeTemplateCreate",
		},
		{
			Key: "NoticeTemplateUpdate",
			API: "/api/w8t/noticeTemplate/noticeTemplateUpdate",
		},
		{
			Key: "NoticeTemplateDelete",
			API: "/api/w8t/noticeTemplate/noticeTemplateDelete",
		},
	}

	perms = permissions

	globals.DBCli.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.UserPermissions{})
	repo.DBCli.Create(&models.UserPermissions{}, &permissions)

}

func InitUserRolesSQL() {

	permsString, _ := json.Marshal(perms)

	roles := models.UserRole{
		ID:          "ur-" + cmd.RandId(),
		Name:        "admin",
		Description: "system",
		Permissions: string(permsString),
		CreateAt:    time.Now().Unix(),
	}

	var adminRole models.UserRole
	globals.DBCli.Model(&models.UserRole{}).Where("name = ?", "admin").First(&adminRole)

	if adminRole.Name != "" {
		return
	}
	repo.DBCli.Create(&models.UserRole{}, &roles)

}
