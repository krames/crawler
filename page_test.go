package main

import (
	"testing"
)

func TestNewPage(t *testing.T) {
	source := "http://testurl.com"
	page := NewPage(source)

	if page.Source == source {
		t.Logf("Source was correctly set to %v", page.Source)
	} else {
		t.Errorf("Expected page.Source to equal %v, but got %v instead", source, page.Source)
	}
}

func TestAddLinks(t *testing.T) {
	page := NewPage("http://testurl.com")
	test_link := "http://somehost.com/path"
	page.AddLink(test_link)
	t.Log(page.Links())

	if len(page.Links()) != 1 {
		t.Errorf("Page should only contain 1 link, but it contained %v", len(page.Links()))
	}

	if page.Links()[0] != test_link {
		t.Error("Link was not added")
	}
}
