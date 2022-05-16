package helper

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

func MarshalInt64(i int64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(i, 10))
	})
}

func UnmarshalInt64(v interface{}) (int64, error) {
	switch data := v.(type) {
	case string:
		return strconv.ParseInt(data, 10, 64)
	case int:
		return int64(data), nil
	case int64:
		return data, nil
	case json.Number:
		return strconv.ParseInt(string(data), 10, 64)
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
