package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the SpiderAjax configuration.
type SpiderAjaxConfig interface {
	constructs.Construct
	Config() ISpiderAjax
	SetConfig(val ISpiderAjax)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the SpiderAjax configuration to YAML format.
	//
	// Returns: The SpiderAjax configuration in YAML format.
	ToYaml() ISpiderAjax
}

// The jsii proxy struct for SpiderAjaxConfig
type jsiiProxy_SpiderAjaxConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SpiderAjaxConfig) Config() ISpiderAjax {
	var returns ISpiderAjax
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjaxConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of SpiderAjaxConfig.
func NewSpiderAjaxConfig(scope constructs.Construct, id *string, props ISpiderAjaxProps) SpiderAjaxConfig {
	_init_.Initialize()

	if err := validateNewSpiderAjaxConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpiderAjaxConfig{}

	_jsii_.Create(
		"zap-cdk.SpiderAjaxConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of SpiderAjaxConfig.
func NewSpiderAjaxConfig_Override(s SpiderAjaxConfig, scope constructs.Construct, id *string, props ISpiderAjaxProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SpiderAjaxConfig",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SpiderAjaxConfig)SetConfig(val ISpiderAjax) {
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
func SpiderAjaxConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpiderAjaxConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.SpiderAjaxConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiderAjaxConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiderAjaxConfig) ToYaml() ISpiderAjax {
	var returns ISpiderAjax

	_jsii_.Invoke(
		s,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

