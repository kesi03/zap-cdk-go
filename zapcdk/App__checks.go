//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func validateApp_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

