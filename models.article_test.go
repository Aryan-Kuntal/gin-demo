package main

import (
	testing "testing"
)

func TestGetAllArticles(t *testing.T) {
	aList := getAllArticles()

	if len(aList) != len(articleList) {
		t.Fail()
	}

	for i, v := range aList {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
