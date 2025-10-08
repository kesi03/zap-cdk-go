package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the replacer configuration.
type ReplacerConfig interface {
	constructs.Construct
	Config() IReplacer
	SetConfig(val IReplacer)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the replacer configuration to YAML format.
	//
	// Returns: The replacer configuration in YAML format.
	ToYaml() IReplacer
}

// The jsii proxy struct for ReplacerConfig
type jsiiProxy_ReplacerConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ReplacerConfig) Config() IReplacer {
	var returns IReplacer
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ReplacerConfig.
func NewReplacerConfig(scope constructs.Construct, id *string, props IReplacerProps) ReplacerConfig {
	_init_.Initialize()

	if err := validateNewReplacerConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplacerConfig{}

	_jsii_.Create(
		"zap-cdk.ReplacerConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ReplacerConfig.
func NewReplacerConfig_Override(r ReplacerConfig, scope constructs.Construct, id *string, props IReplacerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ReplacerConfig",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_ReplacerConfig)SetConfig(val IReplacer) {
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
func ReplacerConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReplacerConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ReplacerConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplacerConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplacerConfig) ToYaml() IReplacer {
	var returns IReplacer

	_jsii_.Invoke(
		r,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

