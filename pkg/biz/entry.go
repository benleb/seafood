package biz

import (
	"fmt"
	"github.com/d0zingcat/seafood/pkg/opensea"
	"github.com/d0zingcat/seafood/pkg/utils"
	"log"
)

func Run(config *utils.Config) {

	client, err := opensea.NewClient()
	if err != nil {
		log.Fatal("fail to init opensea client")
	}

	resp, err := client.GetCollections(opensea.CollectionsParams{
		AssetOwner: "0xEaFb86C874A1D848Fe14F86A8d14693238D42B19",
		Limit:      10,
	})
	fmt.Println(resp)
	// resps, err := client.GetSingleContract(opensea.SingleContractParams{AssetContractAddress: "0xcdb7c1a6fe7e112210ca548c214f656763e13533"})
	// if err == nil {
	// fmt.Println(resps)
	// }

	//url := "https://api.opensea.io/api/v1/events?only_opensea=false&offset=0&limit=20"

	//req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Add("Accept", "application/json")

	//res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()

	//body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

}
