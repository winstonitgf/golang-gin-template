package token

type TokenModel struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
