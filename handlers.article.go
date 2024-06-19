package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(ctx *gin.Context) {

	articlesList := getAllArticles()

	render(ctx, gin.H{
		"title":   "Home Page",
		"payload": articlesList,
	}, "index.html")
}

func showArticle(ctx *gin.Context) {

	article_id, err := strconv.Atoi(ctx.Param("article_id"))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	article, err := getArticleById(article_id)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	render(ctx, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "index.html")

}
