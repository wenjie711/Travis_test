QShare API
==========

#1. 用户管理
- 尹小委部分

##1.1 登录

请求包：

```
POST /Account?name=<Name>&pwd=<Pwd> HTTP/1.1
Content-Type: text/plain
Content-Length: <ctxListSize>
```

返回包：

```
HTTP/1.1 200 OK
Content-Type: application/json

{
		“role”: <Role>,
		“id”: <Id>
}

```

##1.2 新增用户

请求包：

```
POST /User/Additon?user=<User> HTTP/1.1
Content-Type: text/plain
Content-Length: <ctxListSize>
```

返回包：

```
HTTP/1.1 200 OK
Content-Type:application/json

{
	“error”: <SrrMsg string>
}

```

##1.3 删除用户

请求包：

```
POST /User/Removal/id HTTP/1.1
Content-Type: text/plain
Content-Length: <ctxListSize>
```

返回包：

```
HTTP/1.1 200 OK
Content-Type:application/json

{
	“error”: <SrrMsg string>
}

```

##1.4 编辑用户

请求包：

```
POST /User/Editor?user=<User> HTTP/1.1
Content-Type: text/plain
Content-Length: <ctxListSize>
```

返回包：

```
HTTP/1.1 200 OK
Content-Type:application/json

{
	“error”: <SrrMsg string>
}

```

##1.5 查询用户

请求包：

```
POST /User/Inquiry/id HTTP/1.1
Content-Type: text/plain
Content-Length: <ctxListSize>
```

返回包：

```
HTTP/1.1 200 OK
Content-Type:application/json

{
	“items”:
	[
		{
			“id” : <Id>
			“name” : <Name>,
			“email” : <Email>,
			“realname” : <Realname>,
			“status” : <Status>,
			“roleid” : <Roleid>,
			“createtime” : <Createtime>,
			“lastlogintime” : <Lastlogintime>
		}
	]
}

```

#2. 分享管理

##2.1 获取分享

请求包：

```
GET /sharing?page=<Pagenum>&pageSize=<PageSize> HTTP/1.1
Authorization QShare<Token>
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "nextPage": <NextPageNum>,
    "items": 
    [
        {               
            "id": <SharingId>,   
            "cid": <CategeoryId>,
            "aid": <AdminId>,
            "title" : <TitleContent>",
            "CreateTime": <CreateTime>
        },
        ...
    ]
}
```

##2.2 获取分享详细内容

请求包：

```
GET /sharing/id HTTP/1.1
Authorization QShare<Token>
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "item": {         
            "id": <SharingId>,   
            "cid": <CategeoryId>,
            "aid": <AdminId>,
            "title" : <TitleContent>",
            "CreateTime": <CreateTime>
   	}
}
```

================



##2.3 新增分享

请求包:
```
POST /sharing HTTP/1.1

{
	"cid": <CategoryID	Integer>,
	"aid": <AdminID	Integer>,
	"title": <Title	String>,
	"thumb": <ThumbURL	String>,
	"content": <Content	String>,
	"ispublish": <IsPublish	Integer>,
}
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
	"msg":"success"
}
```

##2.4 修改分享

请求包：
```
PUT /sharing/<ID> HTTP/1.1

{
	"cid": <CategoryID	Integer>,
	"aid": <AdminID	Integer>,
	"title": <Title	String>,
	"thumb": <ThumbURL	String>,
	"content": <Content	String>,
	"ispublish": <IsPublish	Integer>,
}
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
	"msg":"success"
}
```

##2.5 删除分享

请求包：
```
DELETE /sharing/<ID> HTTP/1.1
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
	"msg":"success"
}
```

#3. 附件管理
- 何璧辉部分

##3.1 请求附件URL

请求包：

```
GET /sharing/<Id>/attachmnets/<AttachmnetId>?Action=<Type> HTTP/1.1
Accept application/json
```

返回包：

```
200 OK
Content-Type: application/json

{
	"attachmentName": <Name>，
	"attachmentType": <Type>，
	"attachmentUrl": <Url>
}
```
##3.2 请求附件下载

请求包：

```
GET /sharing/<Id>/attachmnets/<AttachmnetId>?Action=<Type> HTTP/1.1
```

返回包：

```
302 Found
```

##3.3 附件上传
`调用qiniu-jsSDK完成上传，上传完成后通过本API提交上传信息`

请求包:
```
POST /attachment HTTP/1.1

{
	"key": <Key	String>,
	"name": <FileName	Integer>,
	"type": <Type	String>
}
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
	"msg":"success"
}
```

##3.4 附件修改
`调用qiniu-jsSDK完成上传，上传完成后通过本API提交上传信息`

请求包：
```
PUT /attachment/<ID> HTTP/1.1

{
	"key": <Key	String>,
	"name": <FileName	Integer>,
	"type": <Type	String>
}
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
	"msg":"success"
}
```

##3.5 附件删除

请求包：
```
DELETE /attachment/<ID> HTTP/1.1
```

返回包：
```
HTTP/1.1 200 OK
Content-Type: application/json

{
	"msg":"success"
}
```
#4. 标签管理 TODO

##4.1 二级标题

请求包：

```
填内容
```

返回包：

```
填内容
```
