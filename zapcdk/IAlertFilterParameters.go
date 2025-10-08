package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAlertFilterParameters interface {
	AlertFilters() *[]IAlertFilter
	SetAlertFilters(a *[]IAlertFilter)
	DeleteGlobalAlerts() *bool
	SetDeleteGlobalAlerts(d *bool)
}

// The jsii proxy for IAlertFilterParameters
type jsiiProxy_IAlertFilterParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IAlertFilterParameters) AlertFilters() *[]IAlertFilter {
	var returns *[]IAlertFilter
	_jsii_.Get(
		j,
		"alertFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilterParameters)SetAlertFilters(val *[]IAlertFilter) {
	if err := j.validateSetAlertFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertFilters",
		val,
	)
}

func (j *jsiiProxy_IAlertFilterParameters) DeleteGlobalAlerts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deleteGlobalAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilterParameters)SetDeleteGlobalAlerts(val *bool) {
	_jsii_.Set(
		j,
		"deleteGlobalAlerts",
		val,
	)
}

