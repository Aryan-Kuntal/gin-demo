package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article Body 1"},
	{ID: 2, Title: "Article 2", Content: "Article Body 2"},
}

func getAllArticles() []article {

	return articleList
}

func getArticleById(id int) (*article, error) {

	for _, article := range articleList {
		if article.ID == id {
			return &article, nil

		}
	}
	return nil, errors.New("article not found")
}
