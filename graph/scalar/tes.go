package scalar

import (
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/oxwazz/twitter/helpers"
)

func MarshalNullString(ns helpers.NullString) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalString(ns.String)
}

// UnmarshalNullString is a custom unmarshaller.
func UnmarshalNullString(v interface{}) (helpers.NullString, error) {
	fmt.Println(3333, v)
	if v == nil {
		return helpers.NullString{struct {
			String string
			Valid  bool
		}{String: "", Valid: false}}, nil
	}
	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalString(v)
	return helpers.NullString{struct {
		String string
		Valid  bool
	}{String: s, Valid: true}}, err
}

func MarshalCustomTime(t helpers.CustomTime) graphql.Marshaler {
	date := t.Time.Format("2006-01-02")
	//date = fmt.Sprintf(`"%s"`, date)
	return graphql.MarshalString(date)
}

func UnmarshalCustomTime(v interface{}) (helpers.CustomTime, error) {
	return helpers.CustomTime{}, errors.New("time should be RFC3339Nano formatted string")
}
