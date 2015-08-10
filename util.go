package openfec

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

func ToValues(in interface{}) (url.Values, error) {
	result := url.Values{}
	asjson, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	var asmap map[string]interface{}
	if err := json.Unmarshal(asjson, &asmap); err != nil {
		return nil, err
	}
	for key, val := range asmap {
		switch reflect.TypeOf(val).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				result.Add(key, fmt.Sprint(s.Index(i)))
			}

		default:
			result.Add(key, fmt.Sprint(val))
		}
	}
	return result, nil
}
