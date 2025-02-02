package controllers

import (
	"fmt"
	"log"
	"net/http"
	"ssorc3/verkurzen/internal/data"
	"ssorc3/verkurzen/internal/generate"

	"github.com/gin-gonic/gin"
)

type ShortenController struct {
    repo data.ShortenRepo
    logger *log.Logger
}

func NewShortenController(repo data.ShortenRepo, logger *log.Logger) ShortenController {
    return ShortenController{
        repo: repo,
        logger: logger,
    }
}

func (controller ShortenController) handleGet(c *gin.Context) {
    linkId := c.Param("linkId")

    fullUrl, err := controller.repo.GetFullUrl(linkId)
    if err != nil {
        c.Status(http.StatusInternalServerError)
    }

    c.Redirect(http.StatusTemporaryRedirect, fullUrl)
}

type shortenUrlBody struct {
    FullUrl string `json:"fullUrl"`
}

func (controller ShortenController) handlePost(c *gin.Context) {
    var body shortenUrlBody
    err := c.ShouldBindJSON(&body)
    if err != nil {
        c.Error(err)
    }

    linkId := generate.NewLinkId()
    err = controller.repo.StoreLink(linkId, body.FullUrl)

    if err != nil {
        c.Error(err)
    }

    c.JSON(http.StatusOK, gin.H{
        "linkId": linkId,
        "link": fmt.Sprintf("http://localhost:8081/%s", linkId),
    })
}

func (c ShortenController) RegisterRoutes(router *gin.Engine) {
    router.GET("/:linkId", c.handleGet)
    router.POST("/", c.handlePost)
}
