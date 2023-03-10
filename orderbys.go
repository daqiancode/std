package std

import "strings"

var OrderBysSep = ","
var OrderBySep = " "

type OrderBy struct {
	Field string
	Asc   bool
}

func ParseOrderBy(str string) OrderBy {
	parts := strings.Split(str, OrderBySep)
	if len(parts) == 0 {
		return OrderBy{}
	}
	asc := true
	if len(parts) >= 2 {
		order := strings.ToLower(parts[1])
		asc = order == "asc" || order == "true"
	}
	return OrderBy{
		Field: parts[0],
		Asc:   asc,
	}
}
func (s OrderBy) String() string {
	if s.Asc {
		return s.Field + " ASC"
	} else {
		return s.Field + " DESC"
	}
}

type OrderBys []OrderBy

func ParseOrderBys(orderByStr string) OrderBys {
	parts := strings.Split(orderByStr, OrderBysSep)
	r := make([]OrderBy, len(parts))
	for i, v := range parts {
		r[i] = ParseOrderBy(v)
	}
	return r
}

func (s OrderBys) String() string {
	r := ""
	for _, v := range s {
		r += v.String() + ","
	}
	if len(s) > 0 {
		return r[0 : len(r)-1]
	}
	return r
}

func (s OrderBys) Pick(fields ...string) OrderBys {
	var r []OrderBy
	for _, v := range s {
		if indexStr(fields, v.Field) != -1 {
			r = append(r, v)
		}
	}
	return r
}
func (s OrderBys) AddPrefix(prefixes map[string]string) OrderBys {
	var r []OrderBy
	for _, v := range s {
		if p, ok := prefixes[v.Field]; ok {
			v.Field = p + "." + v.Field
		}
		r = append(r, v)

	}
	return r
}

func indexStr(a []string, e string) int {
	for i, v := range a {
		if v == e {
			return i
		}
	}
	return -1
}
