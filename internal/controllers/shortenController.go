package controllers

import (
	"fmt"
	"net/http"
	"ssorc3/verkurzen/internal/config"
	"ssorc3/verkurzen/internal/services"

	"github.com/gin-gonic/gin"
)

type ShortenController struct {
    service services.ShortenService
}

func NewShortenController(service services.ShortenService) ShortenController {
    return ShortenController{
        service: service,
    }
}

// @Param linkId query string true "Id of the link to redirect to"
// @Success 307   "Redirects to the specified link"
// @Router /:linkId [get]
func (controller ShortenController) handleGet(c *gin.Context) {
    linkId := c.Param("linkId")

    fullUrl, err := controller.service.GetFullUrl(linkId)
    if err != nil {
        c.Status(http.StatusInternalServerError)
    }

    c.Redirect(http.StatusTemporaryRedirect, fullUrl)
}

type shortenUrlBody struct {
    FullUrl string `json:"fullUrl"`
}

type shortenUrlResponse struct {
    LinkId string `json:"linkId"`
    Link string `json:"link"`
}

// @Param request body controllers.shortenUrlBody true "Expected body"
// @Success 200 {object} controllers.shortenUrlResponse
// @Router / [post]
func (controller ShortenController) handlePost(c *gin.Context) {
    var body shortenUrlBody
    err := c.ShouldBindJSON(&body)
    if err != nil {
        c.Error(err)
    }

    linkId, err := controller.service.StoreUrl(body.FullUrl)

    if err != nil {
        c.Error(err)
    }

    c.JSON(http.StatusOK, shortenUrlResponse{
        LinkId: linkId,
        Link: fmt.Sprintf("http://%s:%d/%s", config.Default.Server.Host, config.Default.Server.Port, linkId),
    })
}

func (c ShortenController) RegisterRoutes(router *gin.RouterGroup) {
    router.GET("/:linkId", c.handleGet)
    router.POST("/", c.handlePost)
}
