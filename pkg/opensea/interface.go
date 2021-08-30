package opensea

import "github.com/d0zingcat/seafood/pkg/httpclient"

type OpenSea interface {
	OpenSeaAccount
	OpenSeaAsset
	OpenSeaAccount
	OpenSeaContract
	OpenSeaCollection
}

type OpenSeaAccount interface {
	GetDefaultAccount() (Account, error)
	GetDefaultAccountVerified() (bool, error)
	GetAccount(addr string) (Account, error)
	GetAccountVerified(addr string) (bool, error)
}

type OpenSeaAsset interface {
	GetAssets(params AssetsParams) (*AssetsResponse, error)
}

type OpenSeaContract interface {
	GetSingleContract(params SingleContractParams) (*SingleContractResponse, error)
}

type OpenSeaCollection interface {
	GetCollections(params CollectionsParams) ([]Collection, error)
}

func NewClient() (OpenSea, error) {
	client := OpenSeaClient{
		HttpClient: httpclient.BuildRequst(),
	}
	return &client, nil
}
