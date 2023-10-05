package gorm_jsonb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSON map[string]interface{}

type JSONArray []map[string]interface{}

func (p JSON) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *JSON) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	err := json.Unmarshal(source, &p)
	if err != nil {
		return err
	}

	return nil
}

func (ja JSONArray) Value() (driver.Value, error) {
	j, err := json.Marshal(ja)
	return j, err
}

func (ja *JSONArray) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}
	*ja = make([]map[string]interface{}, 0)
	err := json.Unmarshal(source, &ja)
	if err != nil {
		return err
	}

	return nil
}
