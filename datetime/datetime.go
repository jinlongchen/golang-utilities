package datetime

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type DateTime float64

func (dt DateTime) MarshalJSON() ([]byte, error) {
	uType := reflect.TypeOf(dt).Elem()

	return []byte(fmt.Sprintf("%0.2f", n)), nil
}

func (n *DateTime) UnmarshalJSON(b []byte) error {
	var f float64
	err := json.Unmarshal(b, &f)
	*n = DateTime(f)
	return err
}

func (n DateTime) String() string {
	return fmt.Sprintf("%0.2f", n)
}
