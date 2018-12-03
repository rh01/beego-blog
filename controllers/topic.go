package controllers

import (
	"github.com/astaxie/beego"
	"github.com/rh01/beego-blog/models"
	"time"
	"strconv"
	"path"
)

type TopicController struct {
	beego.Controller
}

// 注解路由使用
func (this *TopicController) URLMapping() {
	this.Mapping("View", this.View)
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Topics"], err = models.GetAllTopics("", true)
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "topic.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	cate := this.Input().Get("category")

	// 文件上传功能实现-获取附件
	_, header, e := this.GetFile("attachment")
	if e != nil {
		beego.Error(e)
	}

	var attachment string
	if header!=nil {
		// 保存附件
		attachment = header.Filename
		beego.Trace(attachment)
		err := this.SaveToFile("attachment", path.Join("attachment", attachment))
		// filename: tmp.go
		// attachment/tmp.go
		if err!=nil {
			beego.Error(err)
		}

	}


	err := models.AddTopic(title, content, cate, attachment)
	//err = models.AddOrUpdateCategory(cate)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)

}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
}

// @router /topic/view/:id [get]
func (this *TopicController) View() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	tid := this.Ctx.Input.Param(":id")
	//var e error
	topic, err := models.GetOneTopic(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topic"] = topic

	this.TplName = "topic_view.html"

	replies, e := models.GetAllReplies(tid)
	this.Data["Replies"] = replies
	if e != nil {
		beego.Error(e)
		return
	}

}

func (this *TopicController) Modify() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")

	oneTopic, e := models.GetOneTopic(tid)
	if e != nil {
		beego.Error(e)
	}

	this.Data["Topic"] = oneTopic
	this.TplName = "topic_modify.html"

}

func (this *TopicController) Update() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	cate := this.Input().Get("category")
	attachment := this.Input().Get("attachment")

	topic := &models.Topic{
		Id:       id,
		Title:    title,
		Content:  content,
		Category: cate,
		Updated:  time.Now(),
		Attachment:attachment,
	}
	err = models.UpdateOneTopic(topic)
	//err = models.AddOrUpdateCategory(cate)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return

}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")
	err := models.DeleteOneTopic(tid)

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
	return
}
