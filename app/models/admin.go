package models

import (
	"fmt"
	_ "github.com/qiniu-share/utils"
	_ "github.com/revel/config"
	"github.com/revel/revel"
	_ "html/template"
	_ "os"
	_ "regexp"
	_ "strconv"
	_ "strings"
	"time"
)

type Admin struct {
	Id            int64  `xorm:"pk autoincr"`
	Username      string `xorm:"varchar(255)"`
	Password      string `xorm:"varchar(32)"`
	Roleid        int64  `xorm:"index"`
	Email         string `xorm:"varchar(40)"`
	Realname      string `xorm:"varchar(50)"`
	Status        int64  `xorm:"tinyint(1)"`
	Createtime    string `xorm:"DateTime"`
	Lastlogintime string `xorm:"DateTime"`
}

func MockAdmin() *Admin {
	admin := new(Admin)
	admin.Id = 1
	admin.Email = "admin@qiniu.com"
	admin.Username = "admin"
	admin.Roleid = 1
	admin.Password = "123456"
	admin.Realname = "qiniu"
	admin.Lastlogintime = time.Now().Format("2006-01-02 15:04:04")
	admin.Createtime = time.Now().Format("2006-01-02 15:04:04")
	admin.Status = 1

	return admin
}

func (u *Admin) String() string {
	return fmt.Sprintf("Admin(%s)", u.Username)
}

func (admin *Admin) Validate(v *revel.Validation) {
	v.Check(admin.Username,
		revel.Required{},
		revel.MaxSize{50},
		revel.MinSize{6},
		revel.Match{userRegex},
	)

	ValidatePassword(v, admin.Password).
		Key("admin.Password")

}

func (a *Admin) Create() bool {
	/**/
	admin := new(Admin)

	//admin.Id = a.Id
	admin.Email = a.Email
	admin.Username = a.Username
	admin.Roleid = a.Roleid
	admin.Password = a.Password
	admin.Realname = a.Realname
	admin.Lastlogintime = a.Lastlogintime
	admin.Createtime = a.Createtime
	admin.Status = a.Status

	has, err := DB_Write.Insert(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	pk := DB_Read.IdOf(admin)
	id, _ := pk[0].(int64)
	a.Id = id

	return true
}

//删除内容
func (a *Admin) Delete(Id int64) bool {
	/**/
	admin := new(Admin)

	has, err := DB_Write.Id(Id).Delete(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//编辑内容
func (a *Admin) Update(Id int64) bool {
	/**/
	admin := new(Admin)

	admin.Email = a.Email
	admin.Username = a.Username
	admin.Roleid = a.Roleid
	admin.Password = a.Password
	admin.Realname = a.Realname
	admin.Lastlogintime = a.Lastlogintime
	admin.Createtime = a.Createtime
	admin.Status = a.Status

	has, err := DB_Write.Id(Id).Cols("Email", "Username", "Roleid", "Password", "Realname", "Lastlogtime", "Createtime", "Status").Update(admin)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//根据Id获取内容信息
func (a *Admin) GetById(Id int64) (re bool) {
	/**/
	re = true
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(a)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		re = false
	}

	if !has {
		re = false
	}

	return re
}

func (a *Admin) GetByName(name string) (re bool) {
	/**/
	re = true
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	a.Username = name
	has, err := DB_Read.Get(a)

	//has, err := DB_Read.Where("username=?", name).Get(a)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		re = false
	}

	if !has {
		re = false
	}

	return re
}

func (a *Admin) GetByMail(mail string) *Admin {
	admin := new(Admin)
	admin.Email = mail

	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Get(admin)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return admin
}

func (a *Admin) FindForPage(num int, start int) []*Admin {
	admins := make([]*Admin, 0)

	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	/**/
	err := DB_Read.Limit(num, start).Find(&admins)

	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	return admins
}
