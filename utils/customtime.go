package utils

import (
	"database/sql/driver"
	"log"
	"strings"
	"time"
)
//CustomTime allows to Unmarshal json with custom layout
//CustomTime can be used to pass NULL timestamp to database
type CustomTime struct {
    Time time.Time
    Valid bool
}

const dateLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON (b []byte) (err error) {
    s := strings.Trim(string(b), "\"")
    log.Println(s)
    if s == "null" {
        ct.Time = time.Time{}
        return
    }
    ct.Time, err = time.Parse(dateLayout, s)
    return
}

// Scan implements the Scanner interface.
func (nt *CustomTime) Scan(value interface{}) error {
    nt.Time, nt.Valid = value.(time.Time)
    return nil
}

// Value implements the driver Valuer interface.
func (nt CustomTime) Value() (driver.Value, error) {
    if !nt.Valid {
        return nil, nil
    }
    return nt.Time, nil
}
