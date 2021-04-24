package gosteam

import cookiejar "github.com/mihakyr/go-jar"

type session struct {
	cookieClient *cookiejar.CookieJar
	oauth        OAuth
	sessionID    string
	apiKey       string
	deviceID     string
	umqID        string
	chatMessage  int
	language     string
}

type LoginResponse struct {
	Success      bool   `json:"success"`
	PublicKeyMod string `json:"publickey_mod"`
	PublicKeyExp string `json:"publickey_exp"`
	Timestamp    string
	TokenGID     string
}

type LoginSession struct {
	Success           bool   `json:"success"`
	LoginComplete     bool   `json:"login_complete"`
	RequiresTwoFactor bool   `json:"requires_twofactor"`
	Message           string `json:"message"`
	RedirectURI       string `json:"redirect_uri"`
	OAuthInfo         string `json:"oauth"`
}

type OAuth struct {
	SteamID       SteamID `json:"steamid,string"`
	Token         string  `json:"oauth_token"`
	WGToken       string  `json:"wgtoken"`
	WGTokenSecure string  `json:"wgtoken_secure"`
	WebCookie     string  `json:"webcookie"`
}