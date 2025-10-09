package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the ZAP configuration.
//
// Example:
//   const zapConfig = new Zap({
//     env: new Environment({ /* environment config options *\/ }),
//     jobs: [
//       new Job({ /* job config options *\/ }),
//       new Job({ /* another job config options *\/ })
//     ]
//   });
//
type Zap interface {
	IZap
	Env() IEnvironment
	SetEnv(val IEnvironment)
	Jobs() *[]interface{}
	SetJobs(val *[]interface{})
}

// The jsii proxy struct for Zap
type jsiiProxy_Zap struct {
	jsiiProxy_IZap
}

func (j *jsiiProxy_Zap) Env() IEnvironment {
	var returns IEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Zap) Jobs() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"jobs",
		&returns,
	)
	return returns
}


// Creates an instance of Zap.
func NewZap(options IZap) Zap {
	_init_.Initialize()

	if err := validateNewZapParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Zap{}

	_jsii_.Create(
		"zap-cdk.Zap",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Zap.
func NewZap_Override(z Zap, options IZap) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Zap",
		[]interface{}{options},
		z,
	)
}

func (j *jsiiProxy_Zap)SetEnv(val IEnvironment) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_Zap)SetJobs(val *[]interface{}) {
	if err := j.validateSetJobsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobs",
		val,
	)
}

