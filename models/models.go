package models

import (
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const _DB_NAEM = "root:hadoop@/beeblog?charset=utf8"
const _MYSQL_DRIVER = "mysql"

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index;auto_now_add;type(datetime)"`
	Views           int64     `orm:"index"`
	TupleTime       time.Time `orm:"index;auto_now_add;type(datetime)"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Category        string `orm:"index"`
	Attachment      string
	Created         time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"index;auto_now;type(datetime)"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index;auto_now_add;type(datetime)"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index;auto_now_add;type(datetime)"`
}

func RegisterDB() {
	//if !com.IsExist(_DB_NAEM) {
	//	// a/b/c/d/e/1f
	//	os.MkdirAll(path.Dir(_DB_NAEM),os.ModePerm)
	//}
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DB_NAEM, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	category := &Category{Title: name}

	error := o.QueryTable("Category").Filter("title", name).One(category)
	if error == nil {

		return error
	}

	_, e := o.Insert(category)
	if e != nil {
		return e
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	categories := make([]*Category, 0)

	queryTable := o.QueryTable("category")

	_, e := queryTable.All(&categories)
	return categories, e

}

func DeleteCategory(id string) error {
	_, error := strconv.ParseInt(id, 10, 64)
	if error != nil {
		return error
	}
	o := orm.NewOrm()
	//category := &Category{Id: cid}

	// 句柄
	queryTable := o.QueryTable("category")
	//_, e := o.Delete(category)
	_, e := queryTable.Filter("id", id).Delete()
	if e != nil {
		beego.Error(e)
	}
	//o.Commit()
	return nil
}

func AddTopic(title, content, category, attachment string) error {
	o := orm.NewOrm()

	// 检查是否有相同的文章标题
	name := &Topic{Title: title}
	// 我们要插入的标题+内容
	topic := &Topic{Title: title, Content: content, Category: category, Attachment: attachment}

	queryTable := o.QueryTable("topic")
	error := queryTable.Filter("title", title).One(name)
	if error == nil {
		return error
	}

	_, e := o.Insert(topic)
	if e != nil {
		return e
	}

	// 更新分类统计
	cate := new(Category)
	queryT := o.QueryTable("category")
	err := queryT.Filter("Title", category).One(cate)
	if err == nil {
		// 如果不存在简单的忽略
		cate.TopicCount++
		_, err = o.Update(cate)

	}

	return err
}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	queryTable := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			queryTable = queryTable.Filter("category", cate)
		}
		_, err = queryTable.OrderBy("-created").All(&topics)
	} else {
		if len(cate) > 0 {
			queryTable = queryTable.Filter("category", cate)
		}
		_, err = queryTable.All(&topics)
	}
	return topics, err
}

func GetOneTopic(tid string) (*Topic, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != err {
		return nil, err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: id}
	queryTable := o.QueryTable("topic")
	err = queryTable.Filter("id", id).One(topic)
	if err != nil {
		beego.Error(err)
	}

	topic.Views++
	_, err = o.Update(topic)

	return topic, err

}

func UpdateOneTopic(topic *Topic) error {
	o := orm.NewOrm()

	//queryTable := o.QueryTable("topic")
	// 检索topic表，获得旧的category
	tpc := new(Topic)
	err := o.QueryTable("topic").Filter("id", &topic.Id).One(tpc)
	beego.Trace(tpc)
	if err != nil {
		return err
	}

	// 删除旧的附件
	oldAttachment := tpc.Attachment
	if len(oldAttachment) > 0 {
		os.Remove(path.Join("attachment", oldAttachment))
	}

	// 检索category表，获得topicCount值
	category := new(Category)
	oldCate := tpc.Category
	err = o.QueryTable("category").Filter("title", oldCate).One(category)
	beego.Trace(category)
	if err != nil {
		return err
	}
	category.TopicCount--
	_, err = o.Update(category, "TopicCount")

	// =------------------------------=
	// =
	// =
	// =------------------------------=

	// 更新category表，使得更新之後的topicCOunt增加
	newCate := new(Category)
	//newCate := &Category{Title: topic.Category}
	err = o.QueryTable("category").Filter("title", &topic.Category).One(newCate)
	beego.Trace(newCate)
	if err == nil {
		newCate.TopicCount++
		_, err = o.Update(newCate, "TopicCount")
	}

	// 更新topic表
	_, err = o.Update(topic, "Title", "Content", "Category", "Updated", "Attachment")
	return err

}

func DeleteOneTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: id}
	err = o.QueryTable("topic").Filter("id", id).One(topic)
	if err != nil {
		return err
	}
	// 查询并更新category表
	category := new(Category)
	cate := topic.Category
	queryTable := o.QueryTable("category")
	err = queryTable.Filter("Title", cate).One(category)
	if err == nil {
		// 更新
		category.TopicCount--
		_, err = o.Update(category)
	}
	_, e := o.Delete(topic)
	if e != nil {
		return e
	}

	return err

}

func DeleteOneCategory(cid string) error {
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: id}

	_, e := o.Delete(cate)
	if e != nil {
		return e
	}
	return nil

}

func AddReply(tid, nickname, content string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	comment := &Comment{
		Tid:     id,
		Name:    nickname,
		Created: time.Now(),
		Content: content,
	}
	_, err = o.Insert(comment)
	return err
}

func GetAllReplies(tid string) ([]*Comment, error) {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	queryTable := o.QueryTable("comment")

	replies := make([]*Comment, 0)

	_, e := queryTable.Filter("tid", id).All(&replies)
	return replies, e
}

func DeleteComment(id string) error {
	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	comment := &Comment{Id: iid}
	_, e := o.Delete(comment)
	if e != nil {
		return e
	}
	return nil
}

func UpdateTopicReply(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	topic := new(Topic)
	o := orm.NewOrm()
	err = o.QueryTable("topic").Filter("Id", id).One(topic)
	if err != nil {
		return err
	}

	topic.ReplyCount++
	//topic.ReplyLastUserId =
	//ReplyTime       time.Time `orm:"index;auto_now_add;type(datetime)"`
	//ReplyCount      int64
	//ReplyLastUserId int64
	_, err = o.Update(topic, "ReplyCount")
	return err

}
