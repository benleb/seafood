package bark

import (
	"testing"

	"github.com/d0zingcat/seafood/pkg/utils"
)

func TestNewWithoutMap(t *testing.T) {
	conf, err := utils.LoadConfig("../../../..")
	if err != nil {
		t.Fatal("fail to read config")
	}
	utils.Conf = &conf
	b := New(utils.Conf.BarkKey, nil)
	if b == nil {
		t.Fatal("fail to new a bark")
	}
}

func TestNewWithMap(t *testing.T) {
	conf, err := utils.LoadConfig("../../../..")
	if err != nil {
		t.Fatal("fail to read config")
	}
	utils.Conf = &conf
	b := New(utils.Conf.BarkKey, map[string]string{
		"level": "active",
	})
	if b == nil {
		t.Fatal("fail to new a bark")
	}
	b.CheckField("level")
}

func TestSendTitle(t *testing.T) {
	conf, err := utils.LoadConfig("../../../..")
	if err != nil {
		t.Fatal("fail to read config")
	}
	utils.Conf = &conf
	b := New(utils.Conf.BarkKey, nil)
	if b == nil {
		t.Fatal("fail to new a bark")
	}
	err = b.SendTitle("this is a unit test for bark")
	if err != nil {
		t.Fatalf("fail to send title: %v\n", err)
	}

}

func TestSendTitleBody(t *testing.T) {
	conf, err := utils.LoadConfig("../../../..")
	if err != nil {
		t.Fatal("fail to read config")
	}
	utils.Conf = &conf
	b := New(utils.Conf.BarkKey, nil)
	if b == nil {
		t.Fatal("fail to new a bark")
	}
	err = b.SendTitleBody("this is a unit test for bark", "body for bark")
	if err != nil {
		t.Fatalf("fail to send title/body: %v\n", err)
	}

}

func TestSendCategoryTitleBody(t *testing.T) {
	conf, err := utils.LoadConfig("../../../..")
	if err != nil {
		t.Fatal("fail to read config")
	}
	utils.Conf = &conf
	b := New(utils.Conf.BarkKey, map[string]string{"url": "https://www.google.com/"})
	if b == nil {
		t.Fatal("fail to new a bark")
	}
	err = b.SendCategoryTitleBody("category", "this is a unit test for bark with category", "body for bark")
	if err != nil {
		t.Fatalf("fail to send title/body/category: %v\n", err)
	}

}
