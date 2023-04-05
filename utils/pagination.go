package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination -> struct for Pagination
type Pagination struct {
	Page     int
	Sort     string
	PageSize int
	Offset   int
	All      bool
	Keyword  string
}

type CursorPagination struct {
	Limit  int
	Cursor int
}

// BuildPagination -> builds the pagination
func BuildPagination(c *gin.Context) Pagination {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	sort := c.Query("sort")
	keyword := c.Query("keyword")

	var all bool
	if pageSizeStr == "Infinity" {
		all = true
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	return Pagination{
		Page:     page,
		Sort:     sort,
		PageSize: pageSize,
		Offset:   (page - 1) * pageSize,
		All:      all,
		Keyword:  keyword,
	}
}

func BuildCursorPagination(c *gin.Context) (CursorPagination, error) {
	limitStr := c.Query("limit")
	cursorStr := c.Query("cursor")
	
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 1
	}

	cursor, err := strconv.Atoi(cursorStr)
	if err != nil {
		return CursorPagination{
			Limit:  limit,
			Cursor: 0,
		}, err
	}

	return CursorPagination{
		Limit:  limit,
		Cursor: cursor,
	}, nil

}
