package util

import (
	"encoding/json"
	"math"
)

type JsonFloat64 float64

func (j *JsonFloat64) UnmarshalJSON(data []byte) error {
	var s float64
	var inf string
	err := json.Unmarshal(data, &s)
	if err != nil {
		err = json.Unmarshal(data, &inf)
		if err != nil {
			return err
		}
		*j = JsonFloat64(math.Inf(1))
		return nil
	}
	*j = JsonFloat64(s)
	return nil
}

func (j JsonFloat64) MarshalJSON() ([]byte, error) {
	v := float64(j)

	if math.IsInf(v, 0) {
		s := `"Infinity"`
		return []byte(s), nil
	}
	return json.Marshal(v)
}
