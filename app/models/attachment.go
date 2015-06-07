package models

import (
	"github.com/revel/revel"
	"time"
)

type Attachment struct {
	Id         int64  `xorm:"pk autoincr"`
	Sid        int64  `xorm:"int(10)"`
	Key        string `xorm:"varchar(45)"`
	Name       string `xorm:"varchar(45)"`
	Type       string `xorm:"varchar(45)"`
	Downloads  int64  `xorm:"int(10)"`
	Createtime string `xorm:"DateTime"`
	Updatetime string `xorm:"DateTime"`
}

//Mock a Attachment Entity
func MockAttachment() *Attachment {
	attachment := new(Attachment)
	attachment.Sid = 1
	attachment.Key = "Test Key"
	attachment.Name = "Test Name"
	attachment.Type = "Test Type"
	attachment.Downloads = 0
	attachment.Createtime = time.Now().Format("2006-01-02 15:04:04")
	attachment.Updatetime = "0000-00-00 00:00:00"

	return attachment
}

//添加附件
func (a *Attachment) Create(sharing Sharing) bool {

	attachment := new(Attachment)
	attachment.Sid = sharing.Id
	attachment.Key = a.Key
	attachment.Name = a.Name
	attachment.Type = a.Type
	attachment.Downloads = 0
	attachment.Createtime = time.Now().Format("2006-01-02 15:04:04")
	attachment.Updatetime = "0000-00-00 00:00:00"

	has, err := DB_Write.Insert(attachment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	revel.INFO.Printf("Create : %v", attachment)
	return true
}

//修改附件
func (a *Attachment) Update(Id int64) bool {
	attachment := new(Attachment)
	attachment.Key = a.Key
	attachment.Name = a.Name
	attachment.Type = a.Type
	attachment.Downloads = a.Downloads
	attachment.Updatetime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Id(Id).Cols("key", "name", "type", "download", "updatetime").Update(attachment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

func (a *Attachment) Delete(Id int64) bool {
	attachment := new(Attachment)

	has, err := DB_Write.Id(Id).Delete(attachment)
	revel.TRACE.Printf("Delete : %v", attachment)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

func (a *Attachment) GetById(Id int64) *Attachment {
	attachment := new(Attachment)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(attachment)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return attachment
}
