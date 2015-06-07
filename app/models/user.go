package models

import (
	"fmt"
	_ "github.com/qiniu-share/utils"
	_ "github.com/revel/config"
	"github.com/revel/revel"
	_ "html/template"
	_ "os"
	"regexp"
	_ "strconv"
	_ "strings"
	"time"
)

type User struct {
	Id         int64  `xorm:"pk autoincr"`
	Email      string `xorm:"varchar(32)"`
	Username   string `xorm:"varchar(20)"`
	Encrypt    string `xorm:"varchar(6)"`
	Password   string `xorm:"varchar(32)"`
	Nickname   string `xorm:"varchar(20)"`
	Regdate    string `xorm:"DateTime"`
	Lastdate   string `xorm:"DateTime"`
	Loginnum   int64  `xorm:"smallint(5)"`
	Islock     int64  `xorm:"tinyint(1)"`
	Status     int64  `xorm:"tinyint(1)"`
	Createtime string `xorm:"DateTime"`
}

func MockUser() *User {
	user := new(User)
	user.Id = 1
	user.Email = "niuxiaoqi@qiniu.com"
	user.Username = "niuxiaoqi"
	user.Encrypt = "Test Encrypt"
	user.Password = "123456"
	user.Nickname = "qiniu"
	user.Regdate = time.Now().Format("2006-01-02 15:04:04")
	user.Createtime = time.Now().Format("2006-01-02 15:04:04")
	user.Lastdate = "0000-00-00 00:00:00"
	user.Loginnum = 1
	user.Islock = 0
	user.Status = 1

	return user
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{50},
		revel.MinSize{6},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}

func (a *User) Create() bool {
	/**/
	user := new(User)

	//user.Id = a.Id
	user.Email = a.Email
	user.Username = a.Username
	user.Encrypt = a.Encrypt
	user.Password = a.Password
	user.Nickname = a.Nickname
	user.Regdate = a.Regdate
	user.Createtime = a.Createtime
	user.Lastdate = a.Lastdate
	user.Loginnum = a.Loginnum
	user.Islock = a.Islock
	user.Status = a.Status

	has, err := DB_Write.Insert(user)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	pk := DB_Read.IdOf(user)
	id, _ := pk[0].(int64)
	a.Id = id

	return true
}

//删除内容
func (a *User) Delete(Id int64) bool {
	/**/
	user := new(User)

	has, err := DB_Write.Id(Id).Delete(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//编辑内容
func (a *User) Update(Id int64) bool {
	/**/
	user := new(User)

	user.Email = a.Email
	user.Username = a.Username
	user.Encrypt = a.Encrypt
	user.Password = a.Password
	user.Nickname = a.Nickname
	user.Regdate = a.Regdate
	user.Createtime = a.Createtime
	user.Lastdate = a.Lastdate
	user.Loginnum = a.Loginnum
	user.Islock = a.Islock
	user.Status = a.Status

	has, err := DB_Write.Id(Id).Cols("Email", "Username", "Encrypt", "Password", "Nickname", "Regdate", "Createtime", "Lastdate", "Loginnum", "Islock", "Status").Update(user)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}

	return true
}

//根据Id获取内容信息
func (a *User) GetById(Id int64) *User {
	/**/
	user := new(User)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(user)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return user
}

func (a *User) GetByName(name string) (re bool) {
	/**/
	re = true
	a.Username = name
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Get(a)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		re = true
	}

	if !has {
		re = true
	}

	return re
}

func (a *User) GetByMail(mail string) *User {
	user := new(User)
	user.Email = mail

	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Get(user)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	}

	return user
}

func (a *User) FindForPage(num int, start int) []*User {
	users := make([]*User, 0)

	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	/**/
	err := DB_Read.Limit(num, start).Find(&users)

	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	return users
}
