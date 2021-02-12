package model

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RequestQuery struct {
	Limit   int64  `json:"limit"`
	Page    int64  `json:"page"`
	Sort    string `json:"sort"`
	SortDir string `json:"sortDir"`
	Search  string `json:"search"`
}

func newRequestQuery() RequestQuery {
	requestQuery := RequestQuery{
		Limit:   10,
		Page:    1,
		Sort:    "createdAt",
		SortDir: "asc",
		Search:  "",
	}
	return requestQuery
}

func AssignRequestQuery(
	c *fiber.Ctx,
) (RequestQuery, error) {
	var err error
	requestQuery := newRequestQuery()

	limit := c.Query("limit")
	page := c.Query("page")
	sort := c.Query("sort")
	sortDir := c.Query("sortDir")
	search := c.Query("search")

	if limit != "" {
		requestQuery.Limit, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return requestQuery, err
		}
	}
	if page != "" {
		requestQuery.Page, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			return requestQuery, err
		}
	}
	if sort != "" {
		requestQuery.Sort = sort
	}
	if sortDir != "" {
		requestQuery.SortDir = sortDir
	}
	if search != "" {
		requestQuery.Search = strings.TrimSpace(search)
	}

	return requestQuery, nil
}
