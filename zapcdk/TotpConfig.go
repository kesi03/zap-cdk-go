package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the TOTP (Time-based One-Time Password) configuration for a user.
type TotpConfig interface {
	ITotpConfig
	Algorithm() *string
	SetAlgorithm(val *string)
	Digits() *float64
	SetDigits(val *float64)
	Period() *float64
	SetPeriod(val *float64)
	Secret() *string
	SetSecret(val *string)
}

// The jsii proxy struct for TotpConfig
type jsiiProxy_TotpConfig struct {
	jsiiProxy_ITotpConfig
}

func (j *jsiiProxy_TotpConfig) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TotpConfig) Digits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TotpConfig) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TotpConfig) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}


// Creates an instance of TotpConfig.
func NewTotpConfig(options ITotpConfig) TotpConfig {
	_init_.Initialize()

	if err := validateNewTotpConfigParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_TotpConfig{}

	_jsii_.Create(
		"zap-cdk.TotpConfig",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of TotpConfig.
func NewTotpConfig_Override(t TotpConfig, options ITotpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.TotpConfig",
		[]interface{}{options},
		t,
	)
}

func (j *jsiiProxy_TotpConfig)SetAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_TotpConfig)SetDigits(val *float64) {
	_jsii_.Set(
		j,
		"digits",
		val,
	)
}

func (j *jsiiProxy_TotpConfig)SetPeriod(val *float64) {
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_TotpConfig)SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

