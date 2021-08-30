package bark

import (
	"errors"
	"log"

	"github.com/d0zingcat/seafood/pkg/httpclient"
	"github.com/d0zingcat/seafood/pkg/utils"
)

const (
	LEVEL_ACTIVE         = `active`
	LEVEL_TIME_SENSITIVE = `timeSensitive`
	LEVEL_PASSIVE        = `passive`

	TRUE  = `1`
	FALSE = `0`
)

type BarkEnt struct {
	Key string
	// 标题
	Title string
	// 内容
	Body string
	// 分类 占位
	Category string
	// 铃声
	Sound string
	// 跳转的URL地址
	URL string
	// 是否归档
	IsArchive string
	// 组名称
	Group string
	// 自动拷贝的内容
	Copy string
	// 是否自动拷贝
	AutoCopy string
	// 推送ICONURL
	Icon string
	// 时效性通知
	// active：不设置时的默认值，系统会立即亮屏显示通知。
	// timeSensitive：时效性通知，可在专注状态下显示通知。
	// passive：仅将通知添加到通知列表，不会亮屏提醒
	Level string `valid:"active,timeSensitive,passive"`
}

type Bark interface {
	CheckField(string) error
	SendTitle(string) error
	SendTitleBody(string, string) error
	SendCategoryTitleBody(string, string, string) error
}

func New(key string, params map[string]string) (bark Bark) {
	ent := &BarkEnt{
		Key: key,
	}
	bark = ent
	if params == nil {
		return bark
	}
	for k, v := range params {
		switch k {
		case "title":
			ent.Title = v
		case "body":
			ent.Body = v
		case "sound":
			ent.Sound = v
		case "url":
			ent.URL = v
		case "is_archive":
			ent.IsArchive = v
		case "group":
			ent.Group = v
		case "copy":
			ent.Copy = v
		case "auto_copy":
			ent.AutoCopy = v
		case "icon":
			ent.Icon = v
		case "level":
			ent.Level = v
		}
	}
	return bark
}

func (b *BarkEnt) SendTitle(title string) (err error) {
	err = b.checkKey()
	if err != nil {
		return
	}
	_, err = httpclient.BuildRequst().Get(b.concatKey()+utils.QueryEncode(title), b.buildParams())
	if err != nil {
		log.Printf("fail to push title\n")
		return
	}
	return
}

func (b *BarkEnt) SendTitleBody(title, body string) (err error) {
	err = b.checkKey()
	if err != nil {
		return
	}
	_, err = httpclient.BuildRequst().Get(b.concatKey()+utils.QueryEncode(title)+"/"+utils.QueryEncode(body), b.buildParams())
	if err != nil {
		log.Printf("fail to push title with body\n")
		return
	}
	return
}

func (b *BarkEnt) SendCategoryTitleBody(category, title, body string) (err error) {
	err = b.checkKey()
	if err != nil {
		return
	}
	_, err = httpclient.BuildRequst().Get(b.concatKey()+utils.QueryEncode(category)+"/"+utils.QueryEncode(title)+"/"+utils.QueryEncode(body),
		b.buildParams())
	if err != nil {
		log.Printf("fail to push title, body with category\n")
		return
	}
	return
}

func (b *BarkEnt) checkKey() error {
	if b.CheckField("key") != nil {
		return errors.New("Key Field is mandatory!")
	}
	return nil
}

func (b *BarkEnt) CheckField(field string) error {
	var invalid bool
	switch field {
	case "title":
		if b.Title == "" {
			invalid = true
		}
	case "body":
		if b.Body == "" {
			invalid = true
		}
	case "sound":
		if b.Sound == "" {
			invalid = true
		}
	case "url":
		if b.URL == "" {
			invalid = true
		}
	case "is_archive":
		if b.IsArchive == "" {
			invalid = true
		}
	case "group":
		if b.Group == "" {
			invalid = true
		}
	case "copy":
		if b.Copy == "" {
			invalid = true
		}
	case "auto_copy":
		if b.AutoCopy == "" {
			invalid = true
		}
	case "icon":
		if b.Icon == "" {
			invalid = true
		}
	case "level":
		if b.Level == "" {
			invalid = true
		}
	}
	if invalid {
		return errors.New(field + " is invalid")
	}
	return nil
}

func (b *BarkEnt) buildParams() (params httpclient.Params) {
	return httpclient.Params{
		"sound":      b.Sound,
		"url":        b.URL,
		"is_archive": b.IsArchive,
		"group":      b.Group,
		"copy":       b.Copy,
		"auto_copy":  b.AutoCopy,
		"icon":       b.Icon,
		"level":      b.Level,
	}
}

func (b *BarkEnt) concatKey() string {
	return utils.Conf.BarkURL + b.Key + "/"
}
