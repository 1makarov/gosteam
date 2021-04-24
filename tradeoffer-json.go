package gosteam

type EconItemDesc struct {
	ClassID         uint64        `json:"classid,string"`    // for matching with EconItem
	InstanceID      uint64        `json:"instanceid,string"` // for matching with EconItem
	Tradable        bool          `json:"tradable"`
	BackgroundColor string        `json:"background_color"`
	IconURL         string        `json:"icon_url"`
	IconLargeURL    string        `json:"icon_url_large"`
	IconDragURL     string        `json:"icon_drag_url"`
	Name            string        `json:"name"`
	NameColor       string        `json:"name_color"`
	MarketName      string        `json:"market_name"`
	MarketHashName  string        `json:"market_hash_name"`
	MarketFeeApp    uint32        `json:"market_fee_app"`
	Comodity        bool          `json:"comodity"`
	Actions         []*EconAction `json:"actions"`
	Tags            []*EconTag    `json:"tags"`
	Descriptions    []*EconDesc   `json:"descriptions"`
}

type TradeOffer struct {
	ID                 uint64      `json:"tradeofferid,string"`
	Partner            uint32      `json:"accountid_other"`
	ReceiptID          uint64      `json:"tradeid,string"`
	RecvItems          []*EconItem `json:"items_to_receive"`
	SendItems          []*EconItem `json:"items_to_give"`
	Message            string      `json:"message"`
	State              uint8       `json:"trade_offer_state"`
	ConfirmationMethod uint8       `json:"confirmation_method"`
	Created            int64       `json:"time_created"`
	Updated            int64       `json:"time_updated"`
	Expires            int64       `json:"expiration_time"`
	EscrowEndDate      int64       `json:"escrow_end_date"`
	RealTime           bool        `json:"from_real_time_trade"`
	IsOurOffer         bool        `json:"is_our_offer"`
}

type TradeOfferResponse struct {
	Offer          *TradeOffer     `json:"offer"`                 // GetTradeOffer
	SentOffers     []*TradeOffer   `json:"trade_offers_sent"`     // GetTradeOffers
	ReceivedOffers []*TradeOffer   `json:"trade_offers_received"` // GetTradeOffers
	Descriptions   []*EconItemDesc `json:"descriptions"`          // GetTradeOffers
}

type APIResponse struct {
	Inner *TradeOfferResponse `json:"response"`
}

type EconItem struct {
	AssetID    uint64 `json:"assetid,string,omitempty"`
	InstanceID uint64 `json:"instanceid,string,omitempty"`
	ClassID    uint64 `json:"classid,string,omitempty"`
	AppID      uint32 `json:"appid"`
	ContextID  uint64 `json:"contextid,string"`
	Amount     uint32 `json:"amount,string"`
	Missing    bool   `json:"missing,omitempty"`
	EstUSD     uint32 `json:"est_usd,string"`
}

type EconDesc struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Color string `json:"color"`
}

type EconTag struct {
	InternalName          string `json:"internal_name"`
	Category              string `json:"category"`
	LocalizedCategoryName string `json:"localized_category_name"`
	LocalizedTagName      string `json:"localized_tag_name"`
}

type EconAction struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

// https://partner.steamgames.com/doc/webapi/IEconService
type GetTradeOffers struct {
	Apikey               string
	GetSentOffers        bool
	GetReceivedOffers    bool
	GetDescriptions      bool
	ActiveOnly           bool
	HistoricalOnly       bool
	Language             string
	TimeHistoricalCutoff uint32
}
