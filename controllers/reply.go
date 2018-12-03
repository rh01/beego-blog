package controllers

import (
	"github.com/astaxie/beego"
	"github.com/rh01/beego-blog/models"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	tid := this.Input().Get("tid")

	nickname := this.Input().Get("nickname")
	content := this.Input().Get("content")
	err := models.AddReply(tid, nickname, content)
	err = models.UpdateTopicReply(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 301)
	return
}

func (this *ReplyController) Del() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	// 評論id
	id := this.Input().Get("id")
	// topic id
	tid := this.Input().Get("tid")
	err := models.DeleteComment(id)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Redirect("/topic/view/"+tid, 302)
	return
}
