package models

import (
	"encoding/json"

	"gorm.io/datatypes"
)

func MarshalJSON(v interface{}) (datatypes.JSON, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(b), nil
}
