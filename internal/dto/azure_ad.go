package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type (
	GetAccessTokenRequest struct {
		Code        string `json:"code"`
		RedirectUri string `json:"redirect_uri"`
	}
	GetAccessTokenResponse struct {
		AccessToken      string `json:"access_token"`
		TokenType        string `json:"token_type"`
		ExpiresIn        int    `json:"expires_in"`
		Scope            string `json:"scope"`
		IDToken          string `json:"id_token"`
		RefreshToken     string `json:"refresh_token"`
		Error            string `json:"error,omitempty"`
		ErrorDescription string `json:"error_description,omitempty"`
	}
)

func (r GetAccessTokenRequest) Validate() error {
	validation.ErrorTag = "tag"
	return validation.ValidateStruct(&r,
		validation.Field(&r.Code,
			validation.Required),
		validation.Field(&r.RedirectUri,
			validation.Required),
	)
}
