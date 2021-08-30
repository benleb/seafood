package opensea

// Account
type Account struct {
	Data AccountData `json:"data"`
}

type AccountData struct {
	User          UserClass  `json:"user"`
	ProfileImgURL string     `json:"profile_img_url"`
	Address       string     `json:"address"`
	Config        string     `json:"config"`
	Currencies    Currencies `json:"currencies"`
}

type Currencies struct {
}

type UserClass struct {
	Username string `json:"username"`
}

// Asset
type AssetsParams struct {
	Owner                  string `schema:"owner,query_param"`
	TokenIds               string `schema:"token_ids,query_param"`
	AssetContractAddress   string `schema:"asset_contract_address,query_param"`
	AssetContractAddresses string `schema:"asset_contract_addresses,query_param"`
	OrderBy                string `schema:"order_by,query_param"`
	OrderDirection         string `schema:"order_directionz,query_param"`
	Offset                 int    `schema:"offset,query_param"`
	Limit                  int    `schema:"limit,query_param"`
	Collection             string `schema:"collection,query_param"`
}

type AssetsResponse struct {
	Assets []Assets `json:"assets"`
}

type AssetContract struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  interface{} `json:"nft_version"`
	OpenseaVersion              string      `json:"opensea_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                interface{} `json:"external_link"`
	ImageURL                    interface{} `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}
type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}
type User struct {
	Username string `json:"username"`
}

type Owner struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
type Creator struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
type Assets struct {
	ID                      int           `json:"id"`
	TokenID                 string        `json:"token_id"`
	NumSales                int           `json:"num_sales"`
	BackgroundColor         interface{}   `json:"background_color"`
	ImageURL                string        `json:"image_url"`
	ImagePreviewURL         string        `json:"image_preview_url"`
	ImageThumbnailURL       string        `json:"image_thumbnail_url"`
	ImageOriginalURL        interface{}   `json:"image_original_url"`
	AnimationURL            interface{}   `json:"animation_url"`
	AnimationOriginalURL    interface{}   `json:"animation_original_url"`
	Name                    string        `json:"name"`
	Description             string        `json:"description"`
	ExternalLink            interface{}   `json:"external_link"`
	AssetContract           AssetContract `json:"asset_contract"`
	Permalink               string        `json:"permalink"`
	Collection              Collection    `json:"collection"`
	Decimals                interface{}   `json:"decimals"`
	TokenMetadata           interface{}   `json:"token_metadata"`
	Owner                   Owner         `json:"owner"`
	SellOrders              interface{}   `json:"sell_orders"`
	Creator                 Creator       `json:"creator"`
	Traits                  []interface{} `json:"traits"`
	LastSale                interface{}   `json:"last_sale"`
	TopBid                  interface{}   `json:"top_bid"`
	ListingDate             interface{}   `json:"listing_date"`
	IsPresale               bool          `json:"is_presale"`
	TransferFeePaymentToken interface{}   `json:"transfer_fee_payment_token"`
	TransferFee             interface{}   `json:"transfer_fee"`
}

type SingleAssetParams struct {
	AssetContractAddress string `schema:"asset_contract_address,path_param"`
	TokenId              string `schema:"token_id,path_param"`

	AccountAddress string `schema:"account_address,query_param"`
}

type SingleAssetResponse struct {
	ID                      int             `json:"id"`
	TokenID                 string          `json:"token_id"`
	NumSales                int             `json:"num_sales"`
	BackgroundColor         interface{}     `json:"background_color"`
	ImageURL                string          `json:"image_url"`
	ImagePreviewURL         string          `json:"image_preview_url"`
	ImageThumbnailURL       string          `json:"image_thumbnail_url"`
	ImageOriginalURL        string          `json:"image_original_url"`
	AnimationURL            interface{}     `json:"animation_url"`
	AnimationOriginalURL    interface{}     `json:"animation_original_url"`
	Name                    string          `json:"name"`
	Description             interface{}     `json:"description"`
	ExternalLink            string          `json:"external_link"`
	AssetContract           AssetContract   `json:"asset_contract"`
	Permalink               string          `json:"permalink"`
	Collection              Collection      `json:"collection"`
	Decimals                interface{}     `json:"decimals"`
	TokenMetadata           string          `json:"token_metadata"`
	Owner                   Owner           `json:"owner"`
	SellOrders              interface{}     `json:"sell_orders"`
	Creator                 Creator         `json:"creator"`
	Traits                  []Traits        `json:"traits"`
	LastSale                LastSale        `json:"last_sale"`
	TopBid                  interface{}     `json:"top_bid"`
	ListingDate             interface{}     `json:"listing_date"`
	IsPresale               bool            `json:"is_presale"`
	TransferFeePaymentToken interface{}     `json:"transfer_fee_payment_token"`
	TransferFee             interface{}     `json:"transfer_fee"`
	RelatedAssets           []interface{}   `json:"related_assets"`
	Orders                  []interface{}   `json:"orders"`
	Auctions                []interface{}   `json:"auctions"`
	SupportsWyvern          bool            `json:"supports_wyvern"`
	TopOwnerships           []TopOwnerships `json:"top_ownerships"`
	Ownership               Ownership       `json:"ownership"`
	HighestBuyerCommitment  interface{}     `json:"highest_buyer_commitment"`
}
type PaymentTokens struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Address  string  `json:"address"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
	Decimals int     `json:"decimals"`
	EthPrice int     `json:"eth_price"`
	UsdPrice float64 `json:"usd_price"`
}
type PrimaryAssetContracts struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       interface{} `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

type Stats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           int     `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         int     `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        int     `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            int     `json:"total_sales"`
	TotalSupply           int     `json:"total_supply"`
	Count                 int     `json:"count"`
	NumOwners             int     `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            int     `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}
type Collection struct {
	PaymentTokens               []PaymentTokens         `json:"payment_tokens"`
	PrimaryAssetContracts       []PrimaryAssetContracts `json:"primary_asset_contracts"`
	Traits                      Traits                  `json:"traits"`
	Stats                       Stats                   `json:"stats"`
	BannerImageURL              string                  `json:"banner_image_url"`
	ChatURL                     interface{}             `json:"chat_url"`
	CreatedDate                 string                  `json:"created_date"`
	DefaultToFiat               bool                    `json:"default_to_fiat"`
	Description                 string                  `json:"description"`
	DevBuyerFeeBasisPoints      string                  `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string                  `json:"dev_seller_fee_basis_points"`
	DiscordURL                  string                  `json:"discord_url"`
	DisplayData                 DisplayData             `json:"display_data"`
	ExternalURL                 string                  `json:"external_url"`
	Featured                    bool                    `json:"featured"`
	FeaturedImageURL            string                  `json:"featured_image_url"`
	Hidden                      bool                    `json:"hidden"`
	SafelistRequestStatus       string                  `json:"safelist_request_status"`
	ImageURL                    string                  `json:"image_url"`
	IsSubjectToWhitelist        bool                    `json:"is_subject_to_whitelist"`
	LargeImageURL               string                  `json:"large_image_url"`
	MediumUsername              interface{}             `json:"medium_username"`
	Name                        string                  `json:"name"`
	OnlyProxiedTransfers        bool                    `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string                  `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string                  `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               interface{}             `json:"payout_address"`
	RequireEmail                bool                    `json:"require_email"`
	ShortDescription            interface{}             `json:"short_description"`
	Slug                        string                  `json:"slug"`
	TelegramURL                 interface{}             `json:"telegram_url"`
	TwitterUsername             string                  `json:"twitter_username"`
	InstagramUsername           interface{}             `json:"instagram_username"`
	WikiURL                     interface{}             `json:"wiki_url"`
}
type Traits struct {
	TraitType   string      `json:"trait_type"`
	Value       string      `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
}
type Asset struct {
	TokenID  string      `json:"token_id"`
	Decimals interface{} `json:"decimals"`
}
type PaymentToken struct {
	ID       int    `json:"id"`
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	EthPrice string `json:"eth_price"`
	UsdPrice string `json:"usd_price"`
}

type FromAccount struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
type ToAccount struct {
	User          interface{} `json:"user"`
	ProfileImgURL string      `json:"profile_img_url"`
	Address       string      `json:"address"`
	Config        string      `json:"config"`
}
type Transaction struct {
	BlockHash        string      `json:"block_hash"`
	BlockNumber      string      `json:"block_number"`
	FromAccount      FromAccount `json:"from_account"`
	ID               int         `json:"id"`
	Timestamp        string      `json:"timestamp"`
	ToAccount        ToAccount   `json:"to_account"`
	TransactionHash  string      `json:"transaction_hash"`
	TransactionIndex string      `json:"transaction_index"`
}
type LastSale struct {
	Asset          Asset        `json:"asset"`
	AssetBundle    interface{}  `json:"asset_bundle"`
	EventType      string       `json:"event_type"`
	EventTimestamp string       `json:"event_timestamp"`
	AuctionType    interface{}  `json:"auction_type"`
	TotalPrice     string       `json:"total_price"`
	PaymentToken   PaymentToken `json:"payment_token"`
	Transaction    Transaction  `json:"transaction"`
	CreatedDate    string       `json:"created_date"`
	Quantity       string       `json:"quantity"`
}
type TopOwnerships struct {
	Owner    Owner  `json:"owner"`
	Quantity string `json:"quantity"`
}
type Ownership struct {
	Owner    Owner  `json:"owner"`
	Quantity string `json:"quantity"`
}

// contract
type SingleContractParams struct {
	AssetContractAddress string `schema:"asset_contract_address,path_param"`
}

type SingleContractResponse struct {
	Collection                  Collection  `json:"collection"`
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  string      `json:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

// collection
type CollectionsParams struct {
	AssetOwner string `schema:"asset_owner,query_param"`
	Offset     int    `schema:"offset,query_param"`
	Limit      int    `schema:"limit,query_param"`
}
