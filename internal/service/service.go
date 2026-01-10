package service

import commons "proxy-adapter/internal/common"

type Option struct {
	commons.Options
}

type Services struct {
	AzureAD IAzureADService
}
