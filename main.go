package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/d0zingcat/seafood/pkg/biz"
	"github.com/d0zingcat/seafood/pkg/biz/plugins/msgpush/bark"
	"github.com/d0zingcat/seafood/pkg/utils"
	"github.com/robfig/cron/v3"

	_ "github.com/lib/pq"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	utils.Conf = &config

	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresDatabase))
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}
	fmt.Println(db)

	c := cron.New()

	c.AddFunc("*/30 * * * *", func() {
		var bi biz.BizEnt
		title, body, extra := bi.CheckMyAssetsFloorPrice()
		icon := extra["image_url"].(string)

		bark := bark.New(utils.Conf.BarkKey, map[string]string{
			"group": "FP WatchDog",
			"url":   fmt.Sprintf("https://opensea.io/%s", utils.Conf.OpenSeaUsername),
			"icon":  icon,
		})
		bark.SendTitleBody(title, body)
	})
	c.Start()
	defer c.Stop()
	select {}
}
