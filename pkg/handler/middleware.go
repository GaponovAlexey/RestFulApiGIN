package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorization = "Authorization"
	userCtx       = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	handler := c.GetHeader(authorization)
	if handler == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "entry auth header")
		return
	}

	headerParts := strings.Split(handler, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	//pars token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
