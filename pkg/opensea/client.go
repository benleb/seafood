package opensea

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/d0zingcat/seafood/pkg/consts"
	"github.com/d0zingcat/seafood/pkg/httpclient"
	"github.com/d0zingcat/seafood/pkg/utils"
)

type OpenSeaClient struct {
	Addr       string
	Alias      string
	HttpClient *httpclient.Request
}

func (o *OpenSeaClient) GetDefaultAccount() (account Account, err error) {
	resp, err := httpclient.Get(utils.AsembleURL(consts.OPENSEA_API_ENDPOINT_ACCOUNT) + o.Addr)
	if err != nil {
		return
	}
	// defer resp.Body.Close()
	// err = json.NewDecoder(resp.Body).Decode(&account)
	err = resp.Json(&account)
	if err != nil {
		return
	}
	return
}

func (o *OpenSeaClient) GetDefaultAccountVerified() (isVerified bool, err error) {
	account, err := o.GetDefaultAccount()
	if err != nil {
		return
	}
	return account.Data.Config == "verified", nil

}

func (o *OpenSeaClient) GetAccount(addr string) (account Account, err error) {
	resp, err := httpclient.Get(utils.AsembleURL(consts.OPENSEA_API_ENDPOINT_ACCOUNT) + addr)
	if err != nil {
		return
	}
	err = resp.Json(&account)
	if err != nil {
		return
	}
	return
}

func (o *OpenSeaClient) GetAccountVerified(addr string) (isVerified bool, err error) {
	account, err := o.GetAccount(addr)
	if err != nil {
		return
	}
	return account.Data.Config == "verified", nil

}

func (o *OpenSeaClient) GetAssets(params AssetsParams) (assets *AssetsResponse, err error) {
	path_params, query_params := o.buildHTTPGetParams(params)
	log.Println(path_params, query_params)
	resp, err := httpclient.Get(utils.AsembleURL(consts.OPENSEA_API_ENDPOINT_ASSETS, path_params...), query_params)
	if err != nil {
		return nil, err
	}
	assets = new(AssetsResponse)
	resp.Json(assets)
	return
}

func (o *OpenSeaClient) GetSingleContract(params SingleContractParams) (singleContract *SingleContractResponse, err error) {
	path_params, query_params := o.buildHTTPGetParams(params)
	resp, err := httpclient.Get(utils.AsembleURL(consts.OPENSEA_API_ENDPOINT_SINGLE_CONTRACT, path_params...), query_params)
	if err != nil {
		return nil, err
	}
	singleContract = new(SingleContractResponse)
	resp.Json(singleContract)
	return
}

func (o *OpenSeaClient) GetSingleAsset(params SingleAssetParams) (singleAsset *SingleAssetResponse, err error) {
	path_params, query_params := o.buildHTTPGetParams(params)
	resp, err := httpclient.Get(utils.AsembleURL(consts.OPENSEA_API_ENDPOINT_ASSET, path_params...), query_params)
	if err != nil {
		return nil, err
	}
	singleAsset = new(SingleAssetResponse)
	resp.Json(singleAsset)
	return
}

func (o *OpenSeaClient) GetCollections(params CollectionsParams) (collections []Collection, err error) {
	path_params, query_params := o.buildHTTPGetParams(params)
	resp, err := httpclient.Get(utils.AsembleURL(consts.OPENSEA_API_ENDPOINT_COLLECTIONS, path_params...), query_params)
	if err != nil {
		return nil, err
	}
	resp.Json(&collections)
	return
}

// as path params should be defined in struct in order like `param:"param_name,path_param"` flag, the value of field
// would be append into one single array
func (o *OpenSeaClient) buildHTTPGetParams(entity interface{}) ([]string, httpclient.Params) {
	query_params := make(httpclient.Params)
	path_params := []string{}
	v := reflect.ValueOf(entity)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		tag := string(t.Field(i).Tag.Get("schema"))
		//Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		// TODO: check if param is required
		if strings.Contains(tag, "required") {

		}
		// TODO: could use a default tag
		args := strings.Split(tag, ",")
		param := args[0]
		if strings.Contains(tag, "path_param") {
			path_params = append(path_params, o.fieldValueToString(v.Field(i).Interface()))
		} else {
			query_params[param] = o.fieldValueToString(v.Field(i).Interface())
		}
	}
	return path_params, query_params
}

func (o *OpenSeaClient) fieldValueToString(val interface{}) string {
	res := ""
	switch val.(type) {
	case string:
		res = val.(string)
	case int32:
		res = strconv.Itoa(val.(int))
	case int64:
		res = strconv.Itoa(val.(int))
	case int:
		res = strconv.Itoa(val.(int))
	}
	return res
}
