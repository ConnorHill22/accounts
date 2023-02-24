package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

var model_to_prefix_map = map[string]string{
	"User": "u-",
}

func Generate(model_name string) string {
	return fmt.Sprintf("%v-%v", model_to_prefix_map[model_name], uuid.New())
}
