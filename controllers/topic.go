package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"time"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	var err error
	this.Data["Topics"], err = models.GetAllTopics(true)
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

	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)

}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	tid := this.Input().Get("id")

	//var e error
	topic, e := models.GetOneTopic(tid)
	if e != nil {
		beego.Error(e)
	}
	this.Data["Topic"] = topic

	this.TplName = "topic_view.html"
}

func (this *TopicController) Modify() {
	tid := this.Input().Get("tid")
	oneTopic, e := models.GetOneTopic(tid)
	if e != nil {
		beego.Error(e)
	}

	this.Data["Topic"] = oneTopic
	this.TplName = "topic_modify.html"

}

func (this *TopicController) Update() {
	tid := this.Input().Get("tid")
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	topic := &models.Topic{
		Id:      id,
		Title:   title,
		Content: content,
		Updated: time.Now(),
	}
	err = models.UpdateOneTopic(topic)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return

}

func (this *TopicController) Delete() {
	tid := this.Input().Get("tid")
	err := models.DeleteOneTopic(tid)
	if err!=nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
	return
}
