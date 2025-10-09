//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (j *jsiiProxy_Zap) validateSetEnvParameters(val IEnvironment) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Zap) validateSetJobsParameters(val *[]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case IActiveScan:
			// ok
		case IActiveScanPolicy:
			// ok
		case IActiveScanConfig:
			// ok
		case ISpider:
			// ok
		case IDelay:
			// ok
		case IExitStatus:
			// ok
		case IExport:
			// ok
		case IGraphQL:
			// ok
		case IImport:
			// ok
		case IOpenAPI:
			// ok
		case IPassiveScanConfig:
			// ok
		case IPassiveScanWait:
			// ok
		case IPostman:
			// ok
		case IReplacer:
			// ok
		case IReport:
			// ok
		case IRequest:
			// ok
		case ISoap:
			// ok
		case ISpiderAjax:
			// ok
		case INewType:
			// ok
		default:
			if !_jsii_.IsAnonymousProxy(v) {
				return fmt.Errorf("parameter val[%#v] must be one of the allowed types: IActiveScan, IActiveScanPolicy, IActiveScanConfig, ISpider, IDelay, IExitStatus, IExport, IGraphQL, IImport, IOpenAPI, IPassiveScanConfig, IPassiveScanWait, IPostman, IReplacer, IReport, IRequest, ISoap, ISpiderAjax, INewType; received %#v (a %T)", idx_97dfc6, v, v)
			}
		}
	}

	return nil
}

func validateNewZapParameters(options IZap) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

