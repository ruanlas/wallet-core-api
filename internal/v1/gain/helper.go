package gain

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ruanlas/wallet-core-api/internal/v1/gain/gservice"
)

func validateAndGetSearchParams(c *gin.Context) (*gservice.SearchParams, error) {
	month, _ := strconv.ParseUint(c.Query("month"), 10, 32)
	year, _ := strconv.ParseUint(c.Query("year"), 10, 32)
	page, _ := strconv.ParseUint(c.Query("page"), 10, 32)
	pagesize, _ := strconv.ParseUint(c.Query("page_size"), 10, 32)

	if month == uint64(0) || month > 12 {
		return nil, &InvalidArgs{message: fmt.Sprintf("A param month %d is invalid", month)}
	}
	if year == uint64(0) {
		return nil, &InvalidArgs{message: fmt.Sprintf("A param year %d is invalid", year)}
	}
	if page == uint64(0) {
		page = uint64(1)
	}
	if pagesize == uint64(0) {
		pagesize = uint64(10)
	}
	return gservice.NewSearchParamsBuilder().
		AddMonth(uint(month)).
		AddYear(uint(year)).
		AddPage(uint(page)).
		AddPageSize(uint(pagesize)).
		Build(), nil
}
