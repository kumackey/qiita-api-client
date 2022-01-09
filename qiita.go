package qiita

import (
	"encoding/json"
	"errors"
	"net/http"
)

const apiVersion = "v2"
const basePath = "https://qiita.com/api/"

type ItemsLikesService struct {
	s *Service
}

type ItemsService struct {
	s *Service

	Likes *ItemsLikesService
}

type Service struct {
	client   *http.Client
	BasePath string
	Items    *ItemsService
}

func NewService(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}

	s := &Service{client: client, BasePath: basePath}
	s.Items = NewItemsService(s)

	return s, nil
}

func NewItemsService(s *Service) *ItemsService {
	rs := &ItemsService{s: s}
	rs.Likes = NewItemsLikesService(s)

	return rs
}

func NewItemsLikesService(s *Service) *ItemsLikesService {
	rs := &ItemsLikesService{s: s}

	return rs
}

type ItemsLikesGetCall struct {
	s *Service
}

func (r *ItemsLikesService) Get(_ string) *ItemsLikesGetCall {
	c := &ItemsLikesGetCall{s: r.s}

	return c
}

func (c ItemsLikesGetCall) Do() (interface{}, error) {
	res, err := c.doRequest()

	target := &User{}
	if res.StatusCode == http.StatusNoContent {
		return target, nil
	}
	err = json.NewDecoder(res.Body).Decode(target)
	_ = err

	return target, nil
}

func (c ItemsLikesGetCall) doRequest() (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, c.s.BasePath+apiVersion+"/users/kumackey", nil)

	// 要らないかも
	//req.Header.Set("Authorization", ``)

	return c.s.client.Do(req)
}
