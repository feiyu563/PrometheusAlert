package controllers

// Issue: https://github.com/feiyu563/PrometheusAlert/issues/181

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// GitlabController is a beego controller for gitlab webhook events
type GitlabController struct {
	beego.Controller
}

// GitlabProject is a project section in gitlab event
type GitlabProject struct {
	Name     string `json:"name"`
	Homepage string `json:"homepage"`
}

// GitlabRepository is a repository section in gitlab event
type GitlabRepository struct {
	Name     string `json:"name"`
	Homepage string `json:"homepage"`
}

// GitlabUser is a user section in gitlab event
// event: Issue, Note, Merge Request, Wiki, Job, Deployment, Feature Flag
type GitlabUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// GitlabAssignee is a assignee section in gitlab event
type GitlabAssignee struct {
	Username string `json:"username"`
}

// GitlabCommit is a commit section in gitlab event
type GitlabCommit struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

// GitlabWiki is a wiki section in gitlab event
type GitlabWiki struct {
	WebUrl string `json:"web_url"`
}

// GitlabObjectAttributes is a object_attribute section in gitlab event
// event: Issue, Merge Request, Pipeline, Comment, Wiki
type GitlabObjectAttributes struct {
	// Issue, Merge Request, Wiki
	Action string `json:"action"`
	Title  string `json:"title"`
	// Issue, Merge Request
	State       string `json:"state"`
	Description string `json:"description"`
	// Issue, Comment, Wiki
	Url string `json:"url"`
	// Merge Request
	TargetBranch string `json:"target_branch"`
	SourceBranch string `json:"source_branch"`
	MergeStatus  string `json:"merge_status"`
	// Pipeline
	Ref    string `json:"ref"`
	Sha    string `json:"sha"`
	Source string `json:"source"`
	Status string `json:"status"`
	// Comment
	Note         string `json:"note"`
	NoteableType string `json:"noteable_type"`
	// Wiki
	Message string `json:"message"`
	// Feature Flag
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// GitlabRunner is a runner section in gitlab event
type GitlabRunner struct {
	Active      bool   `json:"active"`
	RunnerType  string `json:"runner_type"`
	IsShared    bool   `json:"is_shared"`
	Description string `json:"description"`
}

// GitlabEvent contains all of gitalb events
type GitlabEvent struct {
	ObjectKind       string                 `json:"object_kind"`
	Ref              string                 `json:"ref"`
	CheckoutSha      string                 `json:"checkout_sha"`
	Message          string                 `json:"message"`
	Username         string                 `json:"user_name"`
	UserUsername     string                 `json:"user_username"`
	Sha              string                 `json:"sha"`
	BuildName        string                 `json:"build_name"`
	BuildStage       string                 `json:"build_stage"`
	BuildStatus      string                 `json:"build_status"`
	Status           string                 `json:"status"`
	DeploymentUrl    string                 `json:"deployable_url"`
	Environment      string                 `json:"environment"`
	Name             string                 `json:"name"`
	Url              string                 `json:"url"`
	Description      string                 `json:"description"`
	Tag              string                 `json:"tag"`
	Action           string                 `json:"action"`
	Commits          []GitlabCommit         `json:"commits"`
	User             GitlabUser             `json:"user"`
	Project          GitlabProject          `json:"project"`
	Repository       GitlabRepository       `json:"repository"`
	ObjectAttributes GitlabObjectAttributes `json:"object_attributes"`
	Assignees        []GitlabAssignee       `json:"assignees"`
	Wiki             GitlabWiki             `json:"wiki"`
	Runner           GitlabRunner           `json:"runner"`
}

// GitlabWeixin sends gitlab webhook events to wx robot
func (c *GitlabController) GitlabWeixin() {
	event := GitlabEvent{}
	eventType := c.Ctx.Request.Header.Get("X-Gitlab-Event")
	wxurl := c.GetString("wxurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &event)
	c.Data["json"] = sendGitlabEvent(3, event, eventType, logsign, wxurl)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

// GitlabDingding sends gitlab webhook events to dingtalk robot
func (c *GitlabController) GitlabDingding() {
	event := GitlabEvent{}
	eventType := c.Ctx.Request.Header.Get("X-Gitlab-Event")
	ddurl := c.GetString("ddurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &event)
	c.Data["json"] = sendGitlabEvent(2, event, eventType, logsign, ddurl)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

// GitlabFeishu sends gitlab webhook events to feishu v2 robot
func (c *GitlabController) GitlabFeishu() {
	event := GitlabEvent{}
	eventType := c.Ctx.Request.Header.Get("X-Gitlab-Event")
	fsurl := c.GetString("fsurl")
	logsign := "[" + LogsSign() + "]"
	logs.Info(logsign, string(c.Ctx.Input.RequestBody))
	json.Unmarshal(c.Ctx.Input.RequestBody, &event)
	c.Data["json"] = sendGitlabEvent(4, event, eventType, logsign, fsurl)
	logs.Info(logsign, c.Data["json"])
	c.ServeJSON()
}

func genWXtext(event GitlabEvent, eventType string) string {
	var WXtext, WXbasetext, WXothertext string

	// 有些payload中不同时包含project和repository信息，因此需要判断下
	var name, homepage string
	if event.Project.Name != "" {
		name = event.Project.Name
		homepage = event.Project.Homepage
	} else {
		name = event.Repository.Name
		homepage = event.Repository.Homepage
	}

	WXbasetext = "[Gitlab事件通知](" + homepage + ")\n" +
		"> `事件类型`: " + eventType + "\n" +
		"> `仓库链接`: [" + name + "](" + homepage + ")\n"

	switch eventType {
	case "Push Hook":
		WXothertext = "> `提交用户`: " + event.Username + "(@" + event.UserUsername + ")\n" +
			"> `当前Ref`: " + event.Ref + "\n" +
			"> `当前提交`: " + event.CheckoutSha + "\n"
		if len(event.Commits) != 0 {
			WXothertext = WXothertext + "> `提交信息`: \n" + "\n" + event.Commits[len(event.Commits)-1].Message
		}
	case "Tag Push Hook":
		WXothertext = "> `提交用户`: " + event.Username + "(@" + event.UserUsername + ")\n" +
			"> `当前Ref`: " + event.Ref + "\n" +
			"> `当前提交`: " + event.CheckoutSha + "\n"
		if len(event.Commits) != 0 {
			WXothertext = WXothertext + "> `提交信息`: \n" + "\n" + event.Commits[0].Message
		}
	case "Merge Request Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"> `源分支`: " + event.ObjectAttributes.SourceBranch + "\n" +
			"> `目标分支`: " + event.ObjectAttributes.TargetBranch + "\n" +
			"> `合并请求链接`: [" + event.ObjectAttributes.Title + "](" + event.ObjectAttributes.Url + ")\n" +
			"> `合并请求状态`: " + event.ObjectAttributes.Action + "\n"
		if len(event.Assignees) != 0 {
			WXothertext = WXothertext + "> `分配给`: @" + event.Assignees[0].Username + "\n"
		}
		// 描述内容可能有多行，放最后
		WXothertext = WXothertext + "> `合并请求描述`:\n" + "\n" + event.ObjectAttributes.Description
	case "Issue Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"> `议题链接`: [" + event.ObjectAttributes.Title + "](" + event.ObjectAttributes.Url + ")\n" +
			"> `议题状态`: " + event.ObjectAttributes.Action + "\n"
		if len(event.Assignees) != 0 {
			WXothertext = WXothertext + "> `分配给`: @" + event.Assignees[0].Username + "\n"
		}
		WXothertext = WXothertext + "> `议题描述`:\n" + "\n" + event.ObjectAttributes.Description
	case "Pipeline Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"> `Ref`: " + event.ObjectAttributes.Ref + "\n" +
			"> `Sha`: " + event.ObjectAttributes.Sha + "\n" +
			"> `源`: " + event.ObjectAttributes.Source + "\n" +
			"> `状态`: " + event.ObjectAttributes.Status + "\n"
	case "Job Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "\n" +
			"> `当前Ref`: " + event.Ref + "\n" +
			"> `当前Sha`: " + event.Sha + "\n" +
			"> `当前环境`: " + event.Environment + "\n" +
			"> `构建名称`: " + event.BuildName + "\n" +
			"> `构建阶段`: " + event.BuildStage + "\n" +
			"> `构建状态`: " + event.BuildStatus + "\n" +
			"> `Runner状态`: " + strconv.FormatBool(event.Runner.Active) + "\n" +
			"> `Runner类型`: " + event.Runner.RunnerType + "\n" +
			"> `Runner是否共享`: " + strconv.FormatBool(event.Runner.IsShared) + "\n" +
			"> `Runner描述`: " + event.Runner.Description + "\n"
	case "Note Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"> `评论地址`: [" + event.ObjectAttributes.Url + "](" + event.ObjectAttributes.Url + ")\n" +
			"> `评论类型`: " + event.ObjectAttributes.NoteableType + "\n" +
			"> `评论内容`:\n" + "\n" + event.ObjectAttributes.Note
	case "Wiki Page Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"> `Wiki链接`: [" + event.ObjectAttributes.Title + "](" + event.Wiki.WebUrl + ")\n" +
			"> `Wiki状态`: " + event.ObjectAttributes.Action + "\n" +
			"> `提交消息`:\n" + "\n" + event.ObjectAttributes.Message
	case "Deployment Hook":
		WXothertext = "> `提交用户`: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"> `部署状态`: " + event.Status + "\n" +
			"> `部署环境`: " + event.Environment + "\n" +
			"> `部署地址`: " + event.DeploymentUrl + "\n"
	case "Feature Flag Hook":
		WXothertext = "> `提交用户`: " + event.User.Username + "(@" + event.User.Name + ")\n" +
			"> `功能标志名称`: " + event.ObjectAttributes.Name + "\n" +
			"> `功能标志状态`: " + strconv.FormatBool(event.ObjectAttributes.Active) + "\n" +
			"> `功能标志描述`:\n" + "\n" + event.ObjectAttributes.Description
	case "Release Hook":
		WXothertext = "> `发布链接`: [" + event.Name + "](" + event.Url + ")\n" +
			"> `发布状态`: " + event.Action + "\n" +
			"> `发布版本`: " + event.Tag + "\n" +
			"> `发布描述`:\n" + "\n" + event.Description
	default:
		WXothertext = "**程序暂未处理的事件**"
	}

	WXtext = WXbasetext + WXothertext
	return WXtext
}

func genDDtext(event GitlabEvent, eventType string) string {
	var DDtext, DDbasetext, DDothertext string

	// 有些payload中不同时包含project和repository信息，因此需要判断下
	var name, homepage string
	if event.Project.Name != "" {
		name = event.Project.Name
		homepage = event.Project.Homepage
	} else {
		name = event.Repository.Name
		homepage = event.Repository.Homepage
	}

	// dingding换行 \n\n
	DDbasetext = "[Gitlab事件通知](" + homepage + ")\n\n" +
		"> <font color='#FF0000'>事件类型</font>: " + eventType + "\n\n" +
		"> <font color='#FF0000'>仓库链接</font>: [" + name + "](" + homepage + ")\n\n"

	switch eventType {
	case "Push Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.Username + "(@" + event.UserUsername + ")\n\n" +
			">  <font color='#FF0000'>当前Ref</font>: " + event.Ref + "\n\n" +
			">  <font color='#FF0000'>当前提交</font>: " + event.CheckoutSha + "\n\n"
		if len(event.Commits) != 0 {
			DDothertext = DDothertext + ">  <font color='#FF0000'>提交信息</font>: \n\n" + "\n\n" + event.Commits[len(event.Commits)-1].Message
		}
	case "Tag Push Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.Username + "(@" + event.UserUsername + ")\n\n" +
			">  <font color='#FF0000'>当前Ref</font>: " + event.Ref + "\n\n" +
			">  <font color='#FF0000'>当前提交</font>: " + event.CheckoutSha + "\n\n"
		if len(event.Commits) != 0 {
			DDothertext = DDothertext + ">  <font color='#FF0000'>提交信息</font>: \n\n" + "\n\n" + event.Commits[0].Message
		}
	case "Merge Request Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "(@" + event.User.Username + ")\n\n" +
			">  <font color='#FF0000'>源分支</font>: " + event.ObjectAttributes.SourceBranch + "\n\n" +
			">  <font color='#FF0000'>目标分支</font>: " + event.ObjectAttributes.TargetBranch + "\n\n" +
			">  <font color='#FF0000'>合并请求链接</font>: [" + event.ObjectAttributes.Title + "](" + event.ObjectAttributes.Url + ")\n\n" +
			">  <font color='#FF0000'>合并请求状态</font>: " + event.ObjectAttributes.Action + "\n\n"
		if len(event.Assignees) != 0 {
			DDothertext = DDothertext + "> <font color='#FF0000'>分配给</font>: @" + event.Assignees[0].Username + "\n\n"
		}
		// 描述内容可能有多行，放最后
		DDothertext = DDothertext + ">  <font color='#FF0000'>合并请求描述</font>:\n\n" + "\n\n" + event.ObjectAttributes.Description
	case "Issue Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "(@" + event.User.Username + ")\n\n" +
			">  <font color='#FF0000'>议题链接</font>: [" + event.ObjectAttributes.Title + "](" + event.ObjectAttributes.Url + ")\n\n" +
			">  <font color='#FF0000'>议题状态</font>: " + event.ObjectAttributes.Action + "\n\n"
		if len(event.Assignees) != 0 {
			DDothertext = DDothertext + ">  <font color='#FF0000'>分配给</font>: @" + event.Assignees[0].Username + "\n\n"
		}
		DDothertext = DDothertext + ">  <font color='#FF0000'>议题描述</font>:\n\n" + "\n\n" + event.ObjectAttributes.Description
	case "Pipeline Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "(@" + event.User.Username + ")\n\n" +
			">  <font color='#FF0000'>Ref</font>: " + event.ObjectAttributes.Ref + "\n\n" +
			">  <font color='#FF0000'>Sha</font>: " + event.ObjectAttributes.Sha + "\n\n" +
			">  <font color='#FF0000'>源</font>: " + event.ObjectAttributes.Source + "\n\n" +
			">  <font color='#FF0000'>状态</font>: " + event.ObjectAttributes.Status + "\n\n"
	case "Job Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "\n\n" +
			">  <font color='#FF0000'>当前Ref</font>: " + event.Ref + "\n\n" +
			">  <font color='#FF0000'>当前Sha</font>: " + event.Sha + "\n\n" +
			">  <font color='#FF0000'>当前环境</font>: " + event.Environment + "\n\n" +
			">  <font color='#FF0000'>构建名称</font>: " + event.BuildName + "\n\n" +
			">  <font color='#FF0000'>构建阶段</font>: " + event.BuildStage + "\n\n" +
			">  <font color='#FF0000'>构建状态</font>: " + event.BuildStatus + "\n\n" +
			">  <font color='#FF0000'>Runner状态</font>: " + strconv.FormatBool(event.Runner.Active) + "\n\n" +
			">  <font color='#FF0000'>Runner类型</font>: " + event.Runner.RunnerType + "\n\n" +
			">  <font color='#FF0000'>Runner是否共享</font>: " + strconv.FormatBool(event.Runner.IsShared) + "\n\n" +
			">  <font color='#FF0000'>Runner描述</font>: " + event.Runner.Description + "\n\n"
	case "Note Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "(@" + event.User.Username + ")\n\n" +
			">  <font color='#FF0000'>评论地址</font>: [" + event.ObjectAttributes.Url + "](" + event.ObjectAttributes.Url + ")\n\n" +
			">  <font color='#FF0000'>评论类型</font>: " + event.ObjectAttributes.NoteableType + "\n\n" +
			">  <font color='#FF0000'>评论内容</font>:\n\n" + "\n\n" + event.ObjectAttributes.Note
	case "Wiki Page Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "(@" + event.User.Username + ")\n\n" +
			">  <font color='#FF0000'>Wiki链接</font>: [" + event.ObjectAttributes.Title + "](" + event.Wiki.WebUrl + ")\n\n" +
			">  <font color='#FF0000'>Wiki状态</font>: " + event.ObjectAttributes.Action + "\n\n" +
			">  <font color='#FF0000'>提交消息</font>:\n\n" + "\n\n" + event.ObjectAttributes.Message
	case "Deployment Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Name + "(@" + event.User.Username + ")\n\n" +
			">  <font color='#FF0000'>部署状态</font>: " + event.Status + "\n\n" +
			">  <font color='#FF0000'>部署环境</font>: " + event.Environment + "\n\n" +
			">  <font color='#FF0000'>部署地址</font>: " + event.DeploymentUrl + "\n\n"
	case "Feature Flag Hook":
		DDothertext = ">  <font color='#FF0000'>提交用户</font>: " + event.User.Username + "(@" + event.User.Name + ")\n\n" +
			">  <font color='#FF0000'>功能标志名称</font>: " + event.ObjectAttributes.Name + "\n\n" +
			">  <font color='#FF0000'>功能标志状态</font>: " + strconv.FormatBool(event.ObjectAttributes.Active) + "\n\n" +
			">  <font color='#FF0000'>功能标志描述</font>:\n\n" + "\n\n" + event.ObjectAttributes.Description
	case "Release Hook":
		DDothertext = "> <font color='#FF0000'>发布链接</font>: [" + event.Name + "](" + event.Url + ")\n\n" +
			">  <font color='#FF0000'>发布状态</font>: " + event.Action + "\n\n" +
			">  <font color='#FF0000'>发布版本</font>: " + event.Tag + "\n\n" +
			">  <font color='#FF0000'>发布描述</font>:\n\n" + "\n\n" + event.Description
	default:
		DDothertext = "**程序暂未处理的事件**"
	}

	DDtext = DDbasetext + DDothertext
	return DDtext
}

func genFStext(event GitlabEvent, eventType string) string {
	var FStext, FSbasetext, FSothertext string

	// 有些payload中不同时包含project和repository信息，因此需要判断下
	var name, homepage string
	if event.Project.Name != "" {
		name = event.Project.Name
		homepage = event.Project.Homepage
	} else {
		name = event.Repository.Name
		homepage = event.Repository.Homepage
	}

	FSbasetext = "[Gitlab事件通知](" + homepage + ")\n" +
		"**事件类型**: " + eventType + "\n" +
		"**仓库链接**: [" + name + "](" + homepage + ")\n"

	switch eventType {
	case "Push Hook":
		FSothertext = "**提交用户**: " + event.Username + "(@" + event.UserUsername + ")\n" +
			"**当前Ref**: " + event.Ref + "\n" +
			"**当前提交**: " + event.CheckoutSha + "\n"
		if len(event.Commits) != 0 {
			FSothertext = FSothertext + "**提交信息**: \n" + "\n" + event.Commits[len(event.Commits)-1].Message
		}
	case "Tag Push Hook":
		FSothertext = "**提交用户**: " + event.Username + "(@" + event.UserUsername + ")\n" +
			"**当前Ref**: " + event.Ref + "\n" +
			"**当前提交**: " + event.CheckoutSha + "\n"
		if len(event.Commits) != 0 {
			FSothertext = FSothertext + "**提交信息**: \n" + "\n" + event.Commits[0].Message
		}
	case "Merge Request Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"**源分支**: " + event.ObjectAttributes.SourceBranch + "\n" +
			"**目标分支**: " + event.ObjectAttributes.TargetBranch + "\n" +
			"**合并请求链接**: [" + event.ObjectAttributes.Title + "](" + event.ObjectAttributes.Url + ")\n" +
			"**合并请求状态**: " + event.ObjectAttributes.Action + "\n"
		if len(event.Assignees) != 0 {
			FSothertext = FSothertext + "**分配给**: @" + event.Assignees[0].Username + "\n"
		}
		// 描述内容可能有多行，放最后
		FSothertext = FSothertext + "**合并请求描述**:\n" + "\n" + event.ObjectAttributes.Description
	case "Issue Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"**议题链接**: [" + event.ObjectAttributes.Title + "](" + event.ObjectAttributes.Url + ")\n" +
			"**议题状态**: " + event.ObjectAttributes.Action + "\n"
		if len(event.Assignees) != 0 {
			FSothertext = FSothertext + "**分配给**: @" + event.Assignees[0].Username + "\n"
		}
		FSothertext = FSothertext + "**议题描述**:\n" + "\n" + event.ObjectAttributes.Description
	case "Pipeline Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"**Ref**: " + event.ObjectAttributes.Ref + "\n" +
			"**Sha**: " + event.ObjectAttributes.Sha + "\n" +
			"**源**: " + event.ObjectAttributes.Source + "\n" +
			"**状态**: " + event.ObjectAttributes.Status + "\n"
	case "Job Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "\n" +
			"**当前Ref**: " + event.Ref + "\n" +
			"**当前Sha**: " + event.Sha + "\n" +
			"**当前环境**: " + event.Environment + "\n" +
			"**构建名称**: " + event.BuildName + "\n" +
			"**构建阶段**: " + event.BuildStage + "\n" +
			"**构建状态**: " + event.BuildStatus + "\n" +
			"**Runner状态**: " + strconv.FormatBool(event.Runner.Active) + "\n" +
			"**Runner类型**: " + event.Runner.RunnerType + "\n" +
			"**Runner是否共享**: " + strconv.FormatBool(event.Runner.IsShared) + "\n" +
			"**Runner描述**: " + event.Runner.Description + "\n"
	case "Note Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"**评论地址**: [" + event.ObjectAttributes.Url + "](" + event.ObjectAttributes.Url + ")\n" +
			"**评论类型**: " + event.ObjectAttributes.NoteableType + "\n" +
			"**评论内容**:\n" + "\n" + event.ObjectAttributes.Note
	case "Wiki Page Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"**Wiki链接**: [" + event.ObjectAttributes.Title + "](" + event.Wiki.WebUrl + ")\n" +
			"**Wiki状态**: " + event.ObjectAttributes.Action + "\n" +
			"**提交消息**:\n" + "\n" + event.ObjectAttributes.Message
	case "Deployment Hook":
		FSothertext = "**提交用户**: " + event.User.Name + "(@" + event.User.Username + ")\n" +
			"**部署状态**: " + event.Status + "\n" +
			"**部署环境**: " + event.Environment + "\n" +
			"**部署地址**: " + event.DeploymentUrl + "\n"
	case "Feature Flag Hook":
		FSothertext = "**提交用户**: " + event.User.Username + "(@" + event.User.Name + ")\n" +
			"**功能标志名称**: " + event.ObjectAttributes.Name + "\n" +
			"**功能标志状态**: " + strconv.FormatBool(event.ObjectAttributes.Active) + "\n" +
			"**功能标志描述**:\n" + "\n" + event.ObjectAttributes.Description
	case "Release Hook":
		FSothertext = "**发布链接**: [" + event.Name + "](" + event.Url + ")\n" +
			"**发布状态**: " + event.Action + "\n" +
			"**发布版本**: " + event.Tag + "\n" +
			"**发布描述**:\n" + "\n" + event.Description
	default:
		FSothertext = "**程序暂未处理的事件**"
	}

	FStext = FSbasetext + FSothertext
	return FStext
}

func sendGitlabEvent(typeid int, event GitlabEvent, eventType, logsign, sendURL string) string {
	switch typeid {
	// 1 email
	case 1:
		EmailMessage := ""
		if sendURL == "" {
			sendURL = beego.AppConfig.String("Default_emails")
		}
		SendEmail(EmailMessage, sendURL, logsign)
	// 2 dingding robot
	case 2:
		DDtext := genDDtext(event, eventType)
		if sendURL == "" {
			sendURL = beego.AppConfig.String("ddurl")
		}
		PostToDingDing("Gitlab", DDtext, sendURL, "", logsign)
	// 3 weixin robot
	case 3:
		WXtext := genWXtext(event, eventType)
		if sendURL == "" {
			sendURL = beego.AppConfig.String("wxurl")
		}
		// Todo
		// @somebody
		PostToWeiXin(WXtext, sendURL, "", logsign)
	case 4:
		FStext := genFStext(event, eventType)
		if sendURL == "" {
			sendURL = beego.AppConfig.String("fsurl")
		}
		// Todo
		// @somebody
		PostToFeiShuv2(eventType, FStext, sendURL, "", logsign)
	default:
		return "未知的发送端，消息没有发送。"
	}

	return "消息发送完成!"
}
