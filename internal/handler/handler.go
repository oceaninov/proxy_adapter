package handler

import (
	commons "proxy-adapter/internal/common"
	"proxy-adapter/internal/service"
)

type HandlerOption struct {
	commons.Options
	*service.Services
}

type Handlers struct {
	AzureAD AzureADHandler
}
