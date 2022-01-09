package qiita

import (
	"net/http"
	"testing"
)

func TestItemsLikesGetCall_Do(t *testing.T) {
	client := &http.Client{}

	srv, _ := NewService(client)
	_, _ = srv.Items.Likes.Get("7ccbc949458bd0af22bd").Do()
}
