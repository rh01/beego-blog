package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"strconv"
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
	Content         string    `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index;auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"index;auto_now;type(datetime)"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index;auto_now_add;type(datetime)"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	//if !com.IsExist(_DB_NAEM) {
	//	// a/b/c/d/e/1f
	//	os.MkdirAll(path.Dir(_DB_NAEM),os.ModePerm)
	//}
	orm.RegisterModel(new(Category), new(Topic))
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
	cid, error := strconv.ParseInt(id, 10, 64)
	if error != nil {
		return error
	}
	o := orm.NewOrm()
	//category := &Category{Id: id}

	// 句柄
	queryTable := o.QueryTable("category")
	_, e := queryTable.Filter("Id", cid).Delete()
	if e != nil {
		beego.Error(e)
	}
	o.Commit()
	return nil
}

func AddTopic(title, content string) error {
	o := orm.NewOrm()

	// 检查是否有相同的文章标题
	name := &Topic{Title: title}
	// 我们要插入的标题+内容
	topic := &Topic{Title: title, Content: content}

	queryTable := o.QueryTable("topic")
	error := queryTable.Filter("title", title).One(name)
	if error == nil {
		return error
	}

	_, e := o.Insert(topic)
	if e != nil {
		return e
	}

	return nil

	return nil
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)

	queryTable := o.QueryTable("topic")

	var err error
	if isDesc {
		_, err = queryTable.OrderBy("-created").All(&topics)
	} else {
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

	topic.Views ++
	_, err = o.Update(topic)

	return topic, err

}

func UpdateOneTopic(topic *Topic) error {
	o := orm.NewOrm()

	//queryTable := o.QueryTable("topic")

	_, e := o.Update(topic, "Title", "Content", "Updated")
	if e != nil {
		return e
	}
	return nil

}

func DeleteOneTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: id}

	_, e := o.Delete(topic)
	if e!= nil {
		return e
	}
	return nil

}
