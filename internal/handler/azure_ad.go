package handler

import (
	"net/http"
	"proxy-adapter/internal/constant"
	"proxy-adapter/internal/dto"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AzureADHandler struct {
	HandlerOption
}

func (h *AzureADHandler) GetAccessToken(c echo.Context) (status int, resp dto.HttpResponse) {
	req := new(dto.GetAccessTokenRequest)
	if err := c.Bind(req); err != nil {
		h.HandlerOption.Options.Logger.Error("Error bind request",
			zap.Error(err),
		)
		status = http.StatusBadRequest
		resp = dto.FailedHttpResponse("", constant.ErrBindRequest, nil)
		return
	}

	err := req.Validate()
	if err != nil {
		status = http.StatusBadRequest
		resp = dto.FailedHttpResponse("", err.Error(), nil)
		return
	}

	data, status, err := h.Services.AzureAD.GetAccessToken(c.Request().Context(), req)
	if err != nil {
		resp = dto.FailedHttpResponse("", err.Error(), nil)
		return
	}

	resp = dto.SuccessHttpResponse("", "Get azure access token succeed", data)
	return
}
