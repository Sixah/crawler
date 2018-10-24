package controller

import (
	"context"
	"crawler/engine"
	"crawler/frontend/model"
	"crawler/frontend/view"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("sixah").Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	return result, nil
}
