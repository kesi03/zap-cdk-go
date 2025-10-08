package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the exit status configuration.
type ExitStatusConfig interface {
	constructs.Construct
	Config() IExitStatus
	SetConfig(val IExitStatus)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the exit status configuration to YAML format.
	//
	// Returns: The exit status configuration in YAML format.
	ToYaml() IExitStatus
}

// The jsii proxy struct for ExitStatusConfig
type jsiiProxy_ExitStatusConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ExitStatusConfig) Config() IExitStatus {
	var returns IExitStatus
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatusConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ExitStatusConfig.
func NewExitStatusConfig(scope constructs.Construct, id *string, props IExitStatusProps) ExitStatusConfig {
	_init_.Initialize()

	if err := validateNewExitStatusConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExitStatusConfig{}

	_jsii_.Create(
		"zap-cdk.ExitStatusConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ExitStatusConfig.
func NewExitStatusConfig_Override(e ExitStatusConfig, scope constructs.Construct, id *string, props IExitStatusProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ExitStatusConfig",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_ExitStatusConfig)SetConfig(val IExitStatus) {
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
func ExitStatusConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExitStatusConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ExitStatusConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExitStatusConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExitStatusConfig) ToYaml() IExitStatus {
	var returns IExitStatus

	_jsii_.Invoke(
		e,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

