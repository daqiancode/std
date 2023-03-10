package std

// type PageParam struct {
// 	PageIndex int
// 	PageSize  int
// }

// func (s PageParam) Offset() int {
// 	return s.PageIndex * s.PageSize
// }

type PageData struct {
	Items      interface{} `json:"items"`
	TotalCount int         `json:"totalCount"`
	PageSize   int         `json:"pageSize"`
	PageIndex  int         `json:"pageIndex"`
	TotalPages int         `json:"totalPages"`
}

func NewPageData(items interface{}, totalCount, pageSize, pageIndex int) PageData {
	var totalPages int
	if pageSize != 0 {
		totalPages = totalCount / pageSize
		if totalCount != pageSize*totalCount {
			totalPages += 1
		}
	}
	return PageData{
		Items:      items,
		TotalCount: totalCount,
		PageSize:   pageSize,
		PageIndex:  pageIndex,
		TotalPages: totalPages,
	}
}
