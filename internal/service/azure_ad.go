package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"proxy-adapter/internal/dto"
	"time"

	"github.com/grokify/go-pkce"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
)

type IAzureADService interface {
	GetAccessToken(ctx context.Context, req *dto.GetAccessTokenRequest) (resp dto.GetAccessTokenResponse, httpStatus int, err error)
}

type azureADService struct {
	opt    Option
	client *resty.Client
}

func NewAzureADService(opt Option) IAzureADService {
	//create http client
	client := resty.New().
		SetTimeout(30 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	return &azureADService{
		opt:    opt,
		client: client,
	}
}

func (s *azureADService) GetAccessToken(ctx context.Context, req *dto.GetAccessTokenRequest) (resp dto.GetAccessTokenResponse, httpStatus int, err error) {
	// Create a code verifier (default 43 bytes)
	verifier, err := pkce.NewCodeVerifier(43)
	if err != nil {
		s.opt.Logger.Error("error create code verifier", zap.Error(err))
		httpStatus = http.StatusInternalServerError
		return
	}
	url := fmt.Sprintf("%s/%s/oauth2/v2.0/token", s.opt.Config.AzureADHost, s.opt.Config.AzureADTenantID)
	requestBody := map[string]string{
		"client_id":             s.opt.Config.AzureADClientID,
		"scope":                 s.opt.Config.AzureADScope,
		"code":                  req.Code,
		"redirect_uri":          s.opt.Config.AzureADRedirectURI,
		"grant_type":            "authorization_code",
		"client_secret":         s.opt.Config.AzureADClientSecret,
		"response_type":         "code",
		"code_challenge":        pkce.CodeChallengeS256(verifier),
		"code_challenge_method": "S256",
	}

	response, err := s.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(requestBody).
		SetResult(&resp).
		Post(url)

	if err != nil {
		s.opt.Logger.Error("error get access token from azure ad", zap.Error(err))
		httpStatus = http.StatusInternalServerError
		return
	}
	err = json.Unmarshal(response.Body(), &resp)
	if err != nil {
		err = errors.New(fmt.Sprintf("[%d] Failed to parse response body", response.StatusCode()))
		return
	}

	if response.StatusCode() != http.StatusOK {
		s.opt.Logger.Error("error get access token from azure ad", zap.Error(err), zap.String("response", string(response.Body())))
		httpStatus = response.StatusCode()
		if resp.Error != "" {
			err = errors.New("[" + resp.Error + "] " + resp.ErrorDescription)
		} else {
			err = errors.New(fmt.Sprintf("[%d] Failed to get access token", response.StatusCode()))
		}
		return
	}
	httpStatus = http.StatusOK
	return
}
