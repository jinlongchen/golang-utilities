package pg

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JsonbMap map[string]interface{}

func (p JsonbMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}
func (p *JsonbMap) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i JsonbMap
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*p = i

	return nil
}
func (p *JsonbMap) From(src interface{}) error {
	if src == nil {
		return nil
	}

	source, err := json.Marshal(src)
	if err != nil {
		return errors.New("type assertion .([]byte) failed")
	}

	var i JsonbMap
	err = json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*p = i

	return nil
}
func (p *JsonbMap) To(src interface{}) error {
	if src == nil {
		return nil
	}
	source, err := json.Marshal(p)
	if err != nil {
		return errors.New("type assertion .([]byte) failed")
	}
	//log.Debugf("marshal source:%s,%v", string(source), p)
	return json.Unmarshal(source, src)
}

type JsonbMapArray []map[string]interface{}

func (p JsonbMapArray) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}
func (p *JsonbMapArray) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i JsonbMapArray
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p = i

	return nil
}
func (p *JsonbMapArray) From(src interface{}) error {
	if src == nil {
		return nil
	}

	source, err := json.Marshal(src)
	if err != nil {
		return errors.New("type assertion .([]byte) failed")
	}

	var i []map[string]interface{}
	err = json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*p = i
	//var ok bool
	//*p, ok = i.([]map[string]interface{})
	//if !ok {
	//    return errors.New("Type assertion .([]map[string]interface{}) failed.")
	//}

	return nil
}
func (p *JsonbMapArray) To(src interface{}) error {
	if src == nil {
		return nil
	}

	source, err := json.Marshal(p)
	if err != nil {
		return errors.New("type assertion .([]byte) failed")
	}

	return json.Unmarshal(source, src)
}

type Array []interface{}

func (p Array) Value() (driver.Value, error) {
	j, err := json.Marshal(&p)
	return j, err
}
func (p *Array) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i Array
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p = i

	return nil
}
func (p *Array) From(src interface{}) error {
	if src == nil {
		return nil
	}

	source, err := json.Marshal(src)
	if err != nil {
		return errors.New("type assertion .([]byte) failed")
	}

	var i []interface{}
	err = json.Unmarshal(source, &i)
	if err != nil {
		return err
	}
	*p = i

	return nil
}
func (p Array) To(src interface{}) error {
	if src == nil {
		return nil
	}
	source, err := json.Marshal(&p)
	if err != nil {
		return errors.New("type assertion .([]byte) failed")
	}

	return json.Unmarshal(source, src)
}
