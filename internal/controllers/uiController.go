package controllers

import (
	"fmt"
	"ssorc3/verkurzen/internal/config"
	"ssorc3/verkurzen/internal/services"

	"github.com/gin-gonic/gin"
)

type UIController struct {
    service services.ShortenService
}

func NewUIController(service services.ShortenService) UIController {
    return UIController{
        service: service,
    }
}

func (controller UIController) Index(c *gin.Context) {
    c.HTML(200, "index", nil)
}

func (controller UIController) HandleFormPost(c *gin.Context) {
    type formData struct {
        Url string `form:"url"`
    }

    type shortenLinkData struct {
        Link string
    }

    var data formData
    err := c.Bind(&data)
    if err != nil {
        c.Error(err)
    }

    linkId, err := controller.service.StoreUrl(data.Url)
    if err != nil {
        c.Error(err)
    }

    c.HTML(200, "shortenedLink", shortenLinkData{ Link:
        fmt.Sprintf("http://%s:%d/%s", config.Default.Server.Host, config.Default.Server.Port, linkId) })
}

func (controller UIController) RegisterRoutes(r *gin.RouterGroup) {
    r.GET("/", controller.Index)
    r.POST("/shorten", controller.HandleFormPost)
}
