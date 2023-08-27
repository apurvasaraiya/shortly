package inmemory

import (
	"reflect"
	"testing"
)

func TestInMemoryDBSave(t *testing.T) {
	url := "https://google.com"
	id := "id1"

	db := NewInMemoryDB()
	err := db.SaveURLAndId(url, id)
	if err != nil {
		t.Fatalf("error occured while saving url %s and id %s %v", url, id, err)
	}
}

func TestInMemoryDBFetchByID(t *testing.T) {
	url := "https://google.com"
	id := "id1"

	db := NewInMemoryDB()
	err := db.SaveURLAndId(url, id)
	if err != nil {
		t.Fatalf("error occured while saving url %s and id %s %v", url, id, err)
	}

	urlFromID, err := db.FetchURLFromID(id)
	if err != nil {
		t.Fatalf("error occured while fetching url from id %s %v", id, err)
	}

	if url != urlFromID {
		t.Fatalf("fetched wrong url %s for id %s", url, id)
	}
}

func TestInMemoryDBFetchByURL(t *testing.T) {
	url := "https://google.com"
	id := "id1"

	db := NewInMemoryDB()
	err := db.SaveURLAndId(url, id)
	if err != nil {
		t.Fatalf("error occured while saving url %s and id %s %v", url, id, err)
	}

	idFromURL, err := db.FetchIDFromURL(url)
	if err != nil {
		t.Fatalf("error occured while fetching id from url %s %v", url, err)
	}

	if id != idFromURL {
		t.Fatalf("fetched wrong id %s for url %s", id, url)
	}
}

func TestInMemoryDBMetrics(t *testing.T) {
	db := NewInMemoryDB()
	db.SaveURLAndId("https://google.com", "id1")
	db.SaveURLAndId("https://gmail.com", "id2")

	db.IncrementCountForHostname("google.com")
	db.IncrementCountForHostname("gmail.com")
	db.IncrementCountForHostname("google.com")

	expectedMetrics := map[string]uint{
		"google.com": 2,
		"gmail.com":  1,
	}

	metrics, err := db.FetchMetrics(3)
	if err != nil {
		t.Fatalf("error occured while fetching metrics %v", err)
	}

	if !reflect.DeepEqual(expectedMetrics, metrics) {
		t.Fatalf("expected: %v,\ngot: %v", expectedMetrics, metrics)
	}
}
