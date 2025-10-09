//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_ActiveScanParameters) validateSetTestsParameters(val *[]interface{}) error {
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case IAlertTest:
			// ok
		case IMonitorTest:
			// ok
		case IStatisticsTest:
			// ok
		case IUrlTest:
			// ok
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter val[%#v] must be one of the allowed types: IAlertTest, IMonitorTest, IStatisticsTest, IUrlTest; received %#v (a %T)", idx_97dfc6, v, v)
			}
		}
	}

	return nil
}

