package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the Zap configuration.
type ZapConfig interface {
	constructs.Construct
	Config() IZap
	SetConfig(val IZap)
	// The tree node.
	Node() constructs.Node
	// Synthesizes the Zap configuration to a YAML string.
	//
	// Returns: The Zap configuration as a YAML string.
	Synth() *string
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the Zap configuration to YAML format.
	//
	// Returns: The Zap configuration in YAML format.
	ToYaml() IZap
}

// The jsii proxy struct for ZapConfig
type jsiiProxy_ZapConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ZapConfig) Config() IZap {
	var returns IZap
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZapConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of Zap.
func NewZapConfig(scope constructs.Construct, id *string, props IZap) ZapConfig {
	_init_.Initialize()

	if err := validateNewZapConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZapConfig{}

	_jsii_.Create(
		"zap-cdk.ZapConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of Zap.
func NewZapConfig_Override(z ZapConfig, scope constructs.Construct, id *string, props IZap) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ZapConfig",
		[]interface{}{scope, id, props},
		z,
	)
}

func (j *jsiiProxy_ZapConfig)SetConfig(val IZap) {
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
func ZapConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateZapConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ZapConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZapConfig) Synth() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"synth",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZapConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZapConfig) ToYaml() IZap {
	var returns IZap

	_jsii_.Invoke(
		z,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

