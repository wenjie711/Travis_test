package controllers

import (
	"fmt"
	"github.com/qiniu-share/app/models"
	_ "github.com/qiniu-share/utils"
	_ "github.com/qiniu/api/rs"
	"github.com/revel/revel"
	_ "golang.org/x/crypto/bcrypt"
)

type Account struct {
	*revel.Controller
}

//首页
func (c Account) Index() revel.Result {
	//
	return c.Render()
}

//登录
func (c Account) Login(name string, pwd string, remember bool) revel.Result {
	//
	tl := "account"
	c.Validation.Required(name)
	c.Validation.Required(pwd)

	if !c.Validation.HasErrors() {
		admin := new(models.Admin)
		if admin.GetByName(name) {

			if admin.Password == pwd {
				c.Session["user"] = name
				c.Session["userrole"] = "admin"
				c.Session["uid"] = fmt.Sprintf("%d", admin.Id)

				if remember {
					c.Session.SetDefaultExpiration()
				} else {
					c.Session.SetNoExpiration()
				}

				tl = "account/admin"
				return c.Redirect(tl)
			}
		}

		//user := models.MockUser()
		user := new(models.User)
		if user.GetByName(name) {

			if user.Password == pwd {
				c.Session["user"] = name
				c.Session["userrole"] = "user"
				c.Session["uid"] = fmt.Sprintf("%d", user.Id)

				if remember {
					c.Session.SetDefaultExpiration()
				} else {
					c.Session.SetNoExpiration()
				}

				tl = "sharing/list"
				return c.Redirect(tl)
			}
		}
	}

	c.Validation.Keep()
	c.FlashParams()
	c.Flash.Out["username"] = name
	c.Flash.Error("用户名或密码错误")
	//c.Flash.Error(admin.Password)
	//return c.RenderTemplate(tl)
	return c.Redirect(tl)
}

func (c Account) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	tl := "account"
	return c.Redirect(tl)
}

func (c Account) Administration() revel.Result {

	return c.Render()
}

/***********************************************

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
*/
