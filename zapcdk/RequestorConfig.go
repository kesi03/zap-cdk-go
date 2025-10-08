package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the requestor configuration.
type RequestorConfig interface {
	constructs.Construct
	Config() IRequestorParameters
	SetConfig(val IRequestorParameters)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the requestor configuration to YAML format.
	//
	// Returns: The requestor configuration in YAML format.
	ToYaml() IRequestorParameters
}

// The jsii proxy struct for RequestorConfig
type jsiiProxy_RequestorConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_RequestorConfig) Config() IRequestorParameters {
	var returns IRequestorParameters
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestorConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of RequestorConfig.
func NewRequestorConfig(scope constructs.Construct, id *string, props IRequestorProps) RequestorConfig {
	_init_.Initialize()

	if err := validateNewRequestorConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RequestorConfig{}

	_jsii_.Create(
		"zap-cdk.RequestorConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of RequestorConfig.
func NewRequestorConfig_Override(r RequestorConfig, scope constructs.Construct, id *string, props IRequestorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.RequestorConfig",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RequestorConfig)SetConfig(val IRequestorParameters) {
	if err := j.validateSetConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func RequestorConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRequestorConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.RequestorConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestorConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RequestorConfig) ToYaml() IRequestorParameters {
	var returns IRequestorParameters

	_jsii_.Invoke(
		r,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

