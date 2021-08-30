package opensea

import "testing"

const (
	MY_ADDR = "0xEaFb86C874A1D848Fe14F86A8d14693238D42B19"
)

func TestGetDefaultAccount(t *testing.T) {
	o := OpenSeaClient{Addr: MY_ADDR}
	account, err := o.GetDefaultAccount()
	if err != nil {
		t.Fatal("fail to get account")
	}
	if account.Data.User.Username == "" {
		t.Fatal("fail to get account info")
	}
	t.Log(account)
}

func TestGetDefaultAccountVerified(t *testing.T) {
	o := OpenSeaClient{Addr: MY_ADDR}
	isVerified, err := o.GetDefaultAccountVerified()
	if err != nil {
		t.Fatal("fail to get verification")
	}
	t.Logf("Account %s verification: %t", MY_ADDR, isVerified)
}

func TestGetAccount(t *testing.T) {
	o := OpenSeaClient{}
	account, err := o.GetAccount(MY_ADDR)
	if err != nil {
		t.Fatal("fail to get account")
	}
	if account.Data.User.Username == "" {
		t.Fatal("fail to get account info")
	}
	t.Log(account)
}

func TestGetAccountVerified(t *testing.T) {
	o := OpenSeaClient{}
	isVerified, err := o.GetAccountVerified(MY_ADDR)
	if err != nil {
		t.Fatal("fail to get verification")
	}
	t.Logf("Account %s verification: %t", MY_ADDR, isVerified)
}

func TestGetSingleContract(t *testing.T) {
	GOOP_ADDR := "0x15a2d6c2b4b9903c27f50cb8b32160ab17f186e2"
	o := OpenSeaClient{}
	singleContractResp, err := o.GetSingleContract(SingleContractParams{
		AssetContractAddress: GOOP_ADDR,
	})
	if err != nil {
		t.Fatal("fail to get single contract")
	}
	t.Logf("%#v\n", singleContractResp)
}

func TestGetAssets(t *testing.T) {
	GOOP_ADDR := "0x15a2d6c2b4b9903c27f50cb8b32160ab17f186e2"
	o := OpenSeaClient{}
	assetsResp, err := o.GetAssets(AssetsParams{
		AssetContractAddress: GOOP_ADDR,
	})
	if err != nil {
		t.Fatal("fail to get assets")
	}
	t.Logf("%#v\n", assetsResp.Assets)
}

func TestGetSingleAsset(t *testing.T) {
	GOOP_ADDR := "0x15a2d6c2b4b9903c27f50cb8b32160ab17f186e2"
	o := OpenSeaClient{}
	singleAssetResp, err := o.GetSingleAsset(SingleAssetParams{
		AssetContractAddress: GOOP_ADDR,
		TokenId:              "1",
	})
	if err != nil {
		t.Fatal("fail to get single asset")
	}
	t.Logf("%#v\n", singleAssetResp)
}

func TestGetCollections(t *testing.T) {
	o := OpenSeaClient{}
	collectionsResp, err := o.GetCollections(CollectionsParams{
		AssetOwner: MY_ADDR,
		Limit:      10,
	})
	if err != nil {
		t.Fatal("fail to get collections")
	}
	t.Logf("%#v\n", collectionsResp)
}
