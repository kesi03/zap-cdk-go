package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the environment configuration.
type EnvironmentConfig interface {
	constructs.Construct
	Config() IEnvironment
	SetConfig(val IEnvironment)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the environment configuration to YAML format.
	//
	// Returns: The environment configuration in YAML format.
	ToYaml() IEnvironment
}

// The jsii proxy struct for EnvironmentConfig
type jsiiProxy_EnvironmentConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EnvironmentConfig) Config() IEnvironment {
	var returns IEnvironment
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvironmentConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of EnvironmentConfig.
func NewEnvironmentConfig(scope constructs.Construct, id *string, props IEnvironmentProps) EnvironmentConfig {
	_init_.Initialize()

	if err := validateNewEnvironmentConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnvironmentConfig{}

	_jsii_.Create(
		"zap-cdk.EnvironmentConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of EnvironmentConfig.
func NewEnvironmentConfig_Override(e EnvironmentConfig, scope constructs.Construct, id *string, props IEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.EnvironmentConfig",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EnvironmentConfig)SetConfig(val IEnvironment) {
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
func EnvironmentConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnvironmentConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.EnvironmentConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnvironmentConfig) ToYaml() IEnvironment {
	var returns IEnvironment

	_jsii_.Invoke(
		e,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

