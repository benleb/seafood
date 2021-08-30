package biz

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"text/template"

	"github.com/d0zingcat/seafood/pkg/opensea"
	"github.com/d0zingcat/seafood/pkg/utils"
)

type BizEnt struct {
}

func New() Biz {
	var biz Biz
	bizEnt := &BizEnt{}
	biz = bizEnt
	return biz
}

type Biz interface {
	CheckMyAssetsFloorPrice() (string, string, map[string]interface{})
}

func (b *BizEnt) CheckMyAssetsFloorPrice() (title, body string, extra map[string]interface{}) {
	extra = make(map[string]interface{})
	client, err := opensea.NewClient()
	if err != nil {
		log.Printf("fail to init an opensea client %v\n", err)
		return
	}
	collections, err := client.GetCollections(opensea.CollectionsParams{
		AssetOwner: utils.Conf.OpenSeaMyAddr,
		Limit:      300,
		Offset:     0,
	})
	if err != nil {
		log.Printf("fail to get opensea collections: %v\n", err)
		return
	}
	highestFloorPriceName := ""
	highestFloorPrice := 0.00
	highestFloorPriceImage := ""
	pushTemplate := `
	{{range $key, $val := .}}
	Name: {{index . "name"}}
	FloorPrice: {{index . "floor_price"}}
	{{end}}
	`
	var params []map[string]interface{}
	for _, collection := range collections {
		if len(collection.PrimaryAssetContracts) == 0 {
			continue
		}
		name := collection.PrimaryAssetContracts[0].Name
		floorPrice := collection.Stats.FloorPrice
		image := collection.ImageURL
		params = append(params, map[string]interface{}{
			"name":        name,
			"floor_price": math.Round(floorPrice*10000) / 10000,
		})
		log.Println(floorPrice, highestFloorPrice)
		if floorPrice >= highestFloorPrice {
			highestFloorPriceName = name
			highestFloorPrice = floorPrice
			highestFloorPriceImage = image
		}
	}
	log.Println(params)
	tmpl, err := template.New("push").Parse(pushTemplate)
	if err != nil {
		log.Printf("fail to parse a template %v\n", err)
		return
	}
	var output bytes.Buffer
	if err := tmpl.Execute(&output, params); err != nil {
		log.Printf("fail to substitute params %v\n", err)
		return
	}
	title = fmt.Sprintf("%s FP: %v", highestFloorPriceName, highestFloorPrice)
	body = output.String()
	extra["image_url"] = highestFloorPriceImage
	return
}
