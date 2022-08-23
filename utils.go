package kavenegar

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//ToString ...
func ToString(i interface{}) string {
	return strings.Trim(strings.Replace(fmt.Sprint(i), " ", ",", -1), "[]")
}

//ToJson ...
func ToJson(i interface{}) string {
	_json, _ := json.Marshal(i)
	return string(_json)
}

//ToUnix ...
func ToUnix(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

//structToUrlValues ...
func structToURLValues(i interface{}) url.Values {
	v := url.Values{}
	if reflect.ValueOf(i).IsNil() {
		return v
	}
	m := structToMapString(i)
	for k, s := range m {
		switch {
		case len(s) == 1:
			v.Set(k, s[0])
		case len(s) > 1:
			for i := range s {
				v.Add(k, s[i])
			}
		}
	}

	return v
}

// structToMapString converts struct as map string
func structToMapString(i interface{}) map[string][]string {
	ms := map[string][]string{}
	iv := reflect.ValueOf(i).Elem()
	tp := iv.Type()

	for i := 0; i < iv.NumField(); i++ {
		if isMap(iv.Field(i)) {
			m := iv.Field(i).Interface()
			for key, value := range m.(map[string]string) {
				ms[key] = []string{value}
			}
		} else {
			k := tp.Field(i).Name
			f := iv.Field(i)
			ms[k] = valueToString(f)
		}
	}

	return ms
}

func isMap(f reflect.Value) bool {
	return reflect.TypeOf(f.Interface()).Kind() == reflect.Map
}

// valueToString converts supported type of f as slice string
func valueToString(f reflect.Value) []string {
	var v []string

	switch reflect.TypeOf(f.Interface()).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = []string{strconv.FormatInt(f.Int(), 10)}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = []string{strconv.FormatUint(f.Uint(), 10)}
	case reflect.Float32:
		v = []string{strconv.FormatFloat(f.Float(), 'f', 4, 32)}
	case reflect.Float64:
		v = []string{strconv.FormatFloat(f.Float(), 'f', 4, 64)}
	case reflect.Bool:
		v = []string{strconv.FormatBool(f.Bool())}
	case reflect.Slice:
		for i := 0; i < f.Len(); i++ {
			if s := valueToString(f.Index(i)); len(s) == 1 {
				v = append(v, s[0])
			}
		}
	case reflect.String:
		v = []string{f.String()}
	}

	return v
}
