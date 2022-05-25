package helpers

import "fmt"

func BuildWhereQuery(query string, filters map[string]string, allowedFilters map[string]string) string {
	filterString := "WHERE 1=1"

	// filter key is the url name of the filter used as the lookup for the allowed filters list
	for filterKey, filterValList := range filters {
		if realFilterName, ok := allowedFilters[filterKey]; ok {

			filterString = fmt.Sprintf(
				"%s AND %s = %s",
				filterString,
				realFilterName,
				filterValList,
			)
		}
	}
	// template the where clause into the original query
	query = fmt.Sprintf(query, filterString)
	return query
}
