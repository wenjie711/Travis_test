package controllers

import (
	"github.com/qiniu-share/app/models"
	"github.com/qiniu-share/utils"
	"github.com/qiniu/api/rs"
	"github.com/revel/revel"
	"math/rand"
	"strconv"
	"time"
)

type Sharing struct {
	*revel.Controller
}

var attachmentContext map[string][]*models.Attachment

func init() {
	attachmentContext = make(map[string][]*models.Attachment)
}

// var attachments []models.Attachment

//首页
func (c Sharing) Index() revel.Result {
	// 每次访问首页都将attachments重置
	_, has := c.Session["uid"]
	if !has {
		c.Session["uid"] = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
		revel.INFO.Printf("Create a new uid %s\n", c.Session["uid"])
	} else {
		uid, _ := c.Session["uid"]
		revel.INFO.Printf("Use exist uid %s\n", uid)
	}
	attachments := make([]*models.Attachment, 0, 5)
	attachmentContext[c.Session["uid"]] = attachments
	return c.RenderTemplate("App/Publish.html")
}

func (c Sharing) Publish(title, content string) revel.Result {
	s := new(models.Sharing)
	s.Aid = 0
	s.Cid = 0
	s.Title = title
	s.Content = content
	s.Thumb = ""
	s.IsPublish = 1
	revel.INFO.Printf("Save article : %v [%v]\n", s, s.Create())
	attachments, _ := attachmentContext[c.Session["uid"]]
	for _, att := range attachments {
		att.Create(*s)
	}
	return c.RenderJson(interface{}(map[string]string{"ok": "success"}))
}

func (c Sharing) UpToken() revel.Result {
	putPolicy := rs.PutPolicy{Scope: "qiniu-share"}
	uptoken := putPolicy.Token(nil)
	return c.RenderJson(interface{}(map[string]string{"uptoken": uptoken}))
}

func (c *Sharing) Upload() revel.Result {
	data := utils.Phase(c.Controller)
	revel.INFO.Printf("Upload finish! : hash->%s , key->%s\n", data["hash"], data["key"])
	a := new(models.Attachment)
	a.Key = data["key"].(string)
	a.Name = data["key"].(string)
	attachments, _ := attachmentContext[c.Session["uid"]]
	attachments = append(attachments, a)
	revel.INFO.Printf("%v\n", attachments)
	return c.RenderJson(interface{}(map[string]string{"ok": "success"}))
}

func (c *Sharing) DownloadUrl(id string) revel.Result {
	a := new(models.Attachment)
	a = a.GetById(4)
	baseUrl := rs.MakeBaseUrl(Domain, a.Key)
	policy := rs.GetPolicy{}
	url := policy.MakeRequest(baseUrl, nil)
	return c.RenderJson(interface{}(map[string]string{"url": url}))
	// return c.Redirect(url)
}

//################
func (c *Sharing) List() revel.Result {
	return c.Render()
}

func (c *Sharing) ListAPI() revel.Result {
	//profile1 string = "七牛于2011年创立，致力于提供最适合开发者的数据在线托管、传输加速以及云端处理等服务。区别于国内外其他云存储，七牛自行研发的全分布式架构解决了其他云存储单一数据中心架构可能存在的风险，同时首创双向加速特性对数据上传下载均加速，使得数据访问速度较传统方案平均提升 50%以上。此外，镜像存储、`客户端直传、断点续上传、云端富媒体处理等七牛云存储独有的功能也大大提升了开发效率，最大程度减少了服务器资源浪费。";
	profile1 := "hello world"
	return c.RenderJson(interface{}(map[string]interface{}{"nextPage": "2", "items": []interface{}{map[string]interface{}{"id": "4", "cid": "5", "aid": "2", "title": "后端服务设计模式", "author": "韩拓", "createTime": "2015.06.04", "brief": profile1}}}))
}

func (c *Sharing) Show() revel.Result {
	sharingDao := new(models.Sharing)
	sharing := sharingDao.GetById(1)
	return c.Render(sharing)
}
