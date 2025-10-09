//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func validateNewTechnologyParameters(options ITechnology) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

