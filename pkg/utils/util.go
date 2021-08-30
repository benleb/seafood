package utils

import (
	"github.com/d0zingcat/seafood/pkg/consts"
	"strings"
)

func AsembleURL(endpoint int, paths ...string) string {
	baseURL := joinURL(consts.OPENSEA_API_URL, consts.OPENSEA_API_ROOT, consts.OPENSEA_API_VERSION)
	switch endpoint {
	case consts.OPENSEA_API_ENDPOINT_ACCOUNT:
		return joinURL(consts.OPENSEA_API_URL, consts.OPENSEA_API_URL_ACCOUNT, paths...)
	case consts.OPENSEA_API_ENDPOINT_ASSET:
		return joinURL(baseURL, consts.OPENSEA_API_URL_ASSETS, paths...)
	case consts.OPENSEA_API_ENDPOINT_ASSETS:
		return joinURL(baseURL, consts.OPENSEA_API_URL_ASSETS, paths...)
	case consts.OPENSEA_API_ENDPOINT_SINGLE_CONTRACT:
		return joinURL(baseURL, consts.OPENSEA_API_URL_SINGLE_CONTRACT, paths...)
	case consts.OPENSEA_API_ENDPOINT_COLLECTIONS:
		return joinURL(baseURL, consts.OPENSEA_API_URL_COLLECTIONS, paths...)
	default:
		return ""
	}

}

func joinURL(baseURL, path string, parts ...string) string {
	return strings.Join(append([]string{baseURL, path}, parts...), "")
}
