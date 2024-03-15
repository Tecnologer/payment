package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ActivityLogDetail json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *ActivityLogDetail) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = ActivityLogDetail(result)

	return err
}

// Value return json value, implement driver.Valuer interface
func (j ActivityLogDetail) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}

	return json.RawMessage(j).MarshalJSON()
}

func (j ActivityLogDetail) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

func (j *ActivityLogDetail) UnmarshalJSON(data []byte) error {
	var result json.RawMessage
	err := json.Unmarshal(data, &result)
	*j = ActivityLogDetail(result)

	return err
}
