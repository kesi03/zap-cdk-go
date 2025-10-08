package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ITotpConfig interface {
	Algorithm() *string
	SetAlgorithm(a *string)
	Digits() *float64
	SetDigits(d *float64)
	Period() *float64
	SetPeriod(p *float64)
	Secret() *string
	SetSecret(s *string)
}

// The jsii proxy for ITotpConfig
type jsiiProxy_ITotpConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_ITotpConfig) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITotpConfig)SetAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_ITotpConfig) Digits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITotpConfig)SetDigits(val *float64) {
	_jsii_.Set(
		j,
		"digits",
		val,
	)
}

func (j *jsiiProxy_ITotpConfig) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITotpConfig)SetPeriod(val *float64) {
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_ITotpConfig) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITotpConfig)SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

