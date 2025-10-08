package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the delay configuration.
type DelayConfig interface {
	constructs.Construct
	Config() IDelay
	SetConfig(val IDelay)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the delay configuration to YAML format.
	//
	// Returns: The delay configuration in YAML format.
	ToYaml() IDelay
}

// The jsii proxy struct for DelayConfig
type jsiiProxy_DelayConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DelayConfig) Config() IDelay {
	var returns IDelay
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DelayConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of DelayConfig.
func NewDelayConfig(scope constructs.Construct, id *string, props IDelayProps) DelayConfig {
	_init_.Initialize()

	if err := validateNewDelayConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DelayConfig{}

	_jsii_.Create(
		"zap-cdk.DelayConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of DelayConfig.
func NewDelayConfig_Override(d DelayConfig, scope constructs.Construct, id *string, props IDelayProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.DelayConfig",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DelayConfig)SetConfig(val IDelay) {
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
func DelayConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDelayConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.DelayConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DelayConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DelayConfig) ToYaml() IDelay {
	var returns IDelay

	_jsii_.Invoke(
		d,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

