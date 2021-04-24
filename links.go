package gosteam

const (
	steamDefault   = "https://steamcommunity.com"
	steamLogin     = steamDefault + "/login"
	steamGetRSAkey = steamDefault + "/login/getrsakey/"
	steamDoLogin   = steamDefault + "/login/dologin/"

	// tradeoffer
	apiGetTradeOffer     = "https://api.steampowered.com/IEconService/GetTradeOffer/v1/?"
	apiGetTradeOffers    = "https://api.steampowered.com/IEconService/GetTradeOffers/v1/?"
	apiDeclineTradeOffer = "https://api.steampowered.com/IEconService/DeclineTradeOffer/v1/"
	apiCancelTradeOffer  = "https://api.steampowered.com/IEconService/CancelTradeOffer/v1/"
	steamTradeoffers = "https://steamcommunity.com/my/tradeoffers/privacy"
)
