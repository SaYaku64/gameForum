// models.article.go

package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func getArticleByID(id int) (*article, error) {
	for _, a := range getArticleFromDB() {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{ID: len(getArticleFromDB()) + 1, Title: title, Content: content}

	insertArticleToDB(a)

	return &a, nil
}
