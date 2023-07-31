package main

import "testing"

// Test for function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check length of list of returned articles equals length of global
	// variable holding the list - articleList
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
