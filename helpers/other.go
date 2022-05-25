package helpers

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse("2006-01-02", s)
	//t3, err := dateparse.ParseAny(s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

func (t *CustomTime) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	////date, _ := time.Parse("2006-01-02", (value))
	//switch v := value.(type) {
	////case []byte:
	////fmt.Println(v)
	////	//return t.UnmarshalText(string(v))
	//case string:
	//	fmt.Println(v)
	//	return t.UnmarshalText(v)
	//default:
	//	//fmt.Errorf("cannot sql.Scan() MyTime from: %#v", v)
	//}
	return nil
}

func (t CustomTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

//UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}
