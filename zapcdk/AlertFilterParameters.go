package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type AlertFilterParameters interface {
	IAlertFilterParameters
	AlertFilters() *[]IAlertFilter
	SetAlertFilters(val *[]IAlertFilter)
	DeleteGlobalAlerts() *bool
	SetDeleteGlobalAlerts(val *bool)
}

// The jsii proxy struct for AlertFilterParameters
type jsiiProxy_AlertFilterParameters struct {
	jsiiProxy_IAlertFilterParameters
}

func (j *jsiiProxy_AlertFilterParameters) AlertFilters() *[]IAlertFilter {
	var returns *[]IAlertFilter
	_jsii_.Get(
		j,
		"alertFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilterParameters) DeleteGlobalAlerts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deleteGlobalAlerts",
		&returns,
	)
	return returns
}


// Creates an instance of AlertFilterParameters.
//
// Example:
//   const alertFilterParams = new AlertFilterParameters({
//     deleteGlobalAlerts: true,
//     alertFilters: [
//       new AlertFilter({
//         ruleId: 10010,
//         newRisk: 'Low',
//         context: 'MyContext',
//         url: '.*example.*',
//         urlRegex: true,
//         parameter: 'sessionid',
//         parameterRegex: false,
//         attack: 'SQL Injection',
//         attackRegex: false,
//         evidence: 'SELECT',
//         evidenceRegex: true
//       })
//     ]
//   });
//
func NewAlertFilterParameters(options IAlertFilterParameters) AlertFilterParameters {
	_init_.Initialize()

	if err := validateNewAlertFilterParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertFilterParameters{}

	_jsii_.Create(
		"zap-cdk.AlertFilterParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AlertFilterParameters.
//
// Example:
//   const alertFilterParams = new AlertFilterParameters({
//     deleteGlobalAlerts: true,
//     alertFilters: [
//       new AlertFilter({
//         ruleId: 10010,
//         newRisk: 'Low',
//         context: 'MyContext',
//         url: '.*example.*',
//         urlRegex: true,
//         parameter: 'sessionid',
//         parameterRegex: false,
//         attack: 'SQL Injection',
//         attackRegex: false,
//         evidence: 'SELECT',
//         evidenceRegex: true
//       })
//     ]
//   });
//
func NewAlertFilterParameters_Override(a AlertFilterParameters, options IAlertFilterParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AlertFilterParameters",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AlertFilterParameters)SetAlertFilters(val *[]IAlertFilter) {
	if err := j.validateSetAlertFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertFilters",
		val,
	)
}

func (j *jsiiProxy_AlertFilterParameters)SetDeleteGlobalAlerts(val *bool) {
	_jsii_.Set(
		j,
		"deleteGlobalAlerts",
		val,
	)
}

