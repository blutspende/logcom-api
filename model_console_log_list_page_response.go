/*
 * LogCom API
 *
 * LogCom Swagger documentation
 *
 * API version: 1.0.48
 * Contact: laborit@blutspende.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logcomapi

type ConsoleLogListPageResponse struct {
	// The actual page number
	CurrentPage int32 `json:"currentPage,omitempty"`
	// The items
	Items []ConsoleLogDto `json:"items,omitempty"`
	// The number of items per page
	PageSize int32 `json:"pageSize,omitempty"`
	// The total count of items
	TotalCount int32 `json:"totalCount,omitempty"`
	// The total pages
	TotalPages int32 `json:"totalPages,omitempty"`
}
