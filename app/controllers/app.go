package controllers

import (
	. "github.com/qiniu/api/conf"
	"github.com/revel/revel"
)

func init() {
	ACCESS_KEY = "-OKRw1ScfA8yFr-r1T3pGMm9LKJSXu9YnOt0vzbm"
	SECRET_KEY = "EUU5E2jSqQ2aetcu4R9IAAm6H1LFrh4IL8lpt8qp"
}

const (
	Domain = "7xjf2x.com1.z0.glb.clouddn.com"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
