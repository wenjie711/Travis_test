package models

//内容管理
import (
	"fmt"
	"github.com/qiniu-share/utils"
	"github.com/revel/config"
	"github.com/revel/revel"
	"html/template"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sharing struct {
	Id         int64  `xorm:"pk autoincr"`
	Cid        int64  `xorm:"index"`
	Aid        int64  `xorm:"index"`
	Title      string `xorm:"varchar(80)"`
	Thumb      string `xorm:"varchar(100)"`
	Content    string `xorm:"text"`
	IsPublish  int64  `xorm:"tinyint(1)"`
	Hits       int64  `xorm:"int(10)"`
	Createtime string `xorm:"DateTime"`
	Updatetime string `xorm:"DateTime"`
	// Category   *Category `xorm:"- <- ->"`
	// Admin      *Admin    `xorm:"- <- ->"`
}

func MockSharing() *Sharing {
	sharing := new(Sharing)
	sharing.Cid = 0
	sharing.Aid = 0
	sharing.Title = "Test Title"
	sharing.Thumb = "Test Thumb"
	sharing.Content = "Test Content"
	sharing.IsPublish = 1
	sharing.Hits = 0
	sharing.Createtime = time.Now().Format("2006-01-02 15:04:04")
	sharing.Updatetime = "0000-00-00 00:00:00"

	return sharing
}

//添加内容
func (a *Sharing) Create() bool {

	sharing := new(Sharing)
	sharing.Cid = a.Cid
	sharing.Aid = a.Aid
	sharing.Title = a.Title
	sharing.Thumb = a.Thumb
	sharing.Content = a.Content
	sharing.IsPublish = a.IsPublish
	sharing.Hits = 0
	sharing.Createtime = time.Now().Format("2006-01-02 15:04:04")
	sharing.Updatetime = "0000-00-00 00:00:00"

	has, err := DB_Write.Insert(sharing)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	pk := DB_Read.IdOf(sharing)
	id, _ := pk[0].(int64)
	a.Id = id
	return true
}

//删除内容
func (a *Sharing) Delete(Id int64) bool {
	sharing := new(Sharing)

	has, err := DB_Write.Id(Id).Delete(sharing)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//编辑内容
func (a *Sharing) Update(Id int64) bool {

	sharing := new(Sharing)
	sharing.Title = a.Title
	sharing.Thumb = a.Thumb
	sharing.Content = a.Content
	sharing.IsPublish = a.IsPublish
	sharing.Updatetime = time.Now().Format("2006-01-02 15:04:04")

	has, err := DB_Write.Id(Id).Cols("title", "thumb", "content", "status", "updatetime").Update(sharing)
	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
		return false
	}
	return true
}

//根据Id获取内容信息
func (a *Sharing) GetById(Id int64) *Sharing {
	sharing := new(Sharing)
	//返回的结果为两个参数，一个has为该条记录是否存在，
	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
	has, err := DB_Read.Id(Id).Get(sharing)

	if err != nil {
		revel.WARN.Println(has)
		revel.WARN.Printf("错误: %v", err)
	} else {
		//目前版本暂时不需要
		// admin := new(Admin)
		// sharing.Admin = admin.GetById(sharing.Aid)

		// //所属栏目
		// category := new(Category)
		// sharing.Category = category.GetById(sharing.Cid)

		if sharing.Thumb != "" {
			//判断是否是系统的分隔符
			separator := "/"
			if os.IsPathSeparator('\\') {
				separator = "\\"
			} else {
				separator = "/"
			}

			config_file := (revel.BasePath + "/conf/config.conf")
			config_file = strings.Replace(config_file, "/", separator, -1)
			config_conf, _ := config.ReadDefault(config_file)

			//前台网站地址
			sitedomain, _ := config_conf.String("website", "website.sitedomain")
			sharing.Thumb = sitedomain + sharing.Thumb
		}

	}

	return sharing
}

//根据Cid获取栏目单页面内容
// func (a *Sharing) GetByCid(Cid int64) *Sharing {
// 	sharing := new(Sharing)
// 	//返回的结果为两个参数，一个has为该条记录是否存在，
// 	//第二个参数err为是否有错误。不管err是否为nil，has都有可能为true或者false。
// 	has, err := DB_Read.Where("`cid`=" + strconv.FormatInt(Cid, 10)).Get(sharing)

// 	if err != nil {
// 		revel.WARN.Println(has)
// 		revel.WARN.Printf("错误: %v", err)
// 	} else {
// 		admin := new(Admin)
// 		sharing.Admin = admin.GetById(sharing.Aid)

// 		//所属栏目
// 		category := new(Category)
// 		sharing.Category = category.GetById(sharing.Cid)
// 	}

// 	return sharing
// }

//获取内容列表
func (a *Sharing) GetByAll(search string, Cid int64, Page int64, Perpage int64) (sharing_arr []*Sharing, html template.HTML, where map[string]interface{}) {
	sharing_list := []*Sharing{}

	//查询条件
	var WhereStr string = " 1 AND `cid`=" + strconv.FormatInt(Cid, 10) + " AND "

	if len(search) > 0 {

		//解码
		where = utils.DecodeSegment(search)

		revel.WARN.Println(where)

		if where["start_time"] != "" {
			WhereStr += " `createtime`='" + fmt.Sprintf("%s", where["start_time"]) + "' AND "
		}

		if where["end_time"] != "" {
			WhereStr += " `updatetime`='" + fmt.Sprintf("%s", where["end_time"]) + "' AND "
		}

		if where["searchtype"] != "" && where["keyword"] != "" {

			if where["searchtype"] == "1" {
				//标题
				WhereStr += " `title` like '%" + fmt.Sprintf("%s", where["keyword"]) + "%' AND "
			} else if where["searchtype"] == "2" {
				//简介
				WhereStr += " `description` like '%" + fmt.Sprintf("%s", where["keyword"]) + "%' AND "
				// } else if where["searchtype"] == "3" {

				// 	//用户名
				// 	keyword := fmt.Sprintf("%s", where["keyword"])
				// 	Keyword, err := strconv.Atoi(keyword)

				// 	revel.WARN.Println(Keyword)

				// 	if err != nil {
				// 		admin := new(Admin)
				// 		admin_info := admin.GetByRealName(keyword)

				// 		if admin_info.Id > 0 {
				// 			WhereStr += " `aid`='" + strconv.FormatInt(admin_info.Id, 10) + "' AND "
				// 		}
				// 	} else {
				// 		revel.WARN.Println(keyword)
				// 		WhereStr += " `aid`='" + keyword + "' AND "
				// 	}

				// } else {
				// 	//ID
				// 	WhereStr += " `id`='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
			}
		}
	}

	WhereStr += " 1 "

	//查询总数
	sharing := new(Sharing)
	Total, err := DB_Read.Where(WhereStr).Count(sharing)
	if err != nil {
		revel.WARN.Printf("错误: %v", err)
	}

	//分页
	Pager := new(utils.Page)
	Pager.SubPage_link = "/Content/list/" + strconv.FormatInt(Cid, 10) + "/"
	Pager.Nums = Total
	Pager.Perpage = Perpage
	Pager.Current_page = Page
	Pager.SubPage_type = 2
	pages := Pager.Show()

	//查询数据
	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&sharing_list)

	//暂时不需要
	// if len(sharing_list) > 0 {
	// 	admin := new(Admin)
	// 	category := new(Category)

	// 	for i, v := range sharing_list {
	// 		sharing_list[i].Admin = admin.GetById(v.Aid)

	// 		//所属栏目
	// 		sharing_list[i].Category = category.GetById(v.Cid)
	// 	}
	// }

	return sharing_list, pages, where
}

//获取内容列表
// func (a *Sharing) GetByList(Cid int64, Search string, Page int64, Perpage int64) (sharing_arr []*Sharing, html template.HTML, where map[string]interface{}) {
// 	sharing_list := []*Sharing{}

// 	//查询条件
// 	var WhereStr string = " 1 AND "

// 	if len(Search) > 0 {

// 		//解码
// 		where = utils.DecodeSegment(Search)

// 		revel.WARN.Println(where)

// 		if where["cid"] != "0" {
// 			WhereStr += " `cid`='" + fmt.Sprintf("%s", where["cid"]) + "' AND "
// 		}

// 		if where["field"] != "" && where["keyword"] != "" {

// 			if where["field"] == "title" {
// 				//标题
// 				WhereStr += " `title` like %'" + fmt.Sprintf("%s", where["keyword"]) + "'% AND "

// 			} else if where["field"] == "keywords" {
// 				//关键词
// 				WhereStr += " `keywords` like %'" + fmt.Sprintf("%s", where["keyword"]) + "'% AND "
// 			} else if where["field"] == "description" {
// 				//描述
// 				WhereStr += " `description` like %'" + fmt.Sprintf("%s", where["keyword"]) + "'% AND "
// 			} else if where["field"] == "id" {
// 				//ID
// 				WhereStr += " `id`='" + fmt.Sprintf("%s", where["keyword"]) + "' AND "
// 			}

// 		}
// 	}

// 	WhereStr += " 1 "

// 	//查询总数
// 	sharing := new(Sharing)
// 	Total, err := DB_Read.Where(WhereStr).Count(sharing)
// 	if err != nil {
// 		revel.WARN.Printf("错误: %v", err)
// 	}

// 	//分页
// 	Pager := new(utils.Page)
// 	Pager.SubPage_link = "/Content/relationlist/" + strconv.FormatInt(Cid, 10) + "/"
// 	Pager.Nums = Total
// 	Pager.Perpage = Perpage
// 	Pager.Current_page = Page
// 	Pager.SubPage_type = 2
// 	pages := Pager.Show()

// 	//查询数据
// 	DB_Read.Where(WhereStr).Limit(int(Perpage), int((Page-1)*Pager.Perpage)).Desc("id").Find(&sharing_list)

// 	if len(sharing_list) > 0 {

// 		admin := new(Admin)
// 		category := new(Category)

// 		for i, v := range sharing_list {
// 			sharing_list[i].Admin = admin.GetById(v.Aid)

// 			//所属栏目
// 			sharing_list[i].Category = category.GetById(v.Cid)
// 		}
// 	}

// 	return sharing_list, pages, where
// }

//批量移动
// func (a *Sharing) Remove(Cid int64, ids string) bool {
// 	sharing := new(Sharing)

// 	sharing.Cid = Cid

// 	has, err := DB_Write.Where("id in (" + ids + ")").Cols("cid").Update(sharing)
// 	if err != nil {
// 		revel.WARN.Println(has)
// 		revel.WARN.Printf("错误: %v", err)
// 		return false
// 	}
// 	return true
// }

//标题是否存在
// func (a *Sharing) HasTitle() bool {

// 	sharing := new(Sharing)
// 	has, err := DB_Read.Where("title=?", a.Title).Get(sharing)
// 	if err != nil {
// 		revel.WARN.Printf("错误: %v", has)
// 		revel.WARN.Printf("错误: %v", err)
// 		return false
// 	}

// 	if sharing.Id > 0 {
// 		return true
// 	}
// 	return false
// }
