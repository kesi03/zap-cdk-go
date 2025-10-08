package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the Spider configuration.
type SpiderConfig interface {
	constructs.Construct
	Config() ISpider
	SetConfig(val ISpider)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the spider configuration to YAML format.
	//
	// Returns: The spider configuration in YAML format.
	ToYaml() ISpider
}

// The jsii proxy struct for SpiderConfig
type jsiiProxy_SpiderConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SpiderConfig) Config() ISpider {
	var returns ISpider
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of SpiderConfig.
func NewSpiderConfig(scope constructs.Construct, id *string, props *SpiderProps) SpiderConfig {
	_init_.Initialize()

	if err := validateNewSpiderConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpiderConfig{}

	_jsii_.Create(
		"zap-cdk.SpiderConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of SpiderConfig.
func NewSpiderConfig_Override(s SpiderConfig, scope constructs.Construct, id *string, props *SpiderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SpiderConfig",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SpiderConfig)SetConfig(val ISpider) {
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
func SpiderConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSpiderConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.SpiderConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiderConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpiderConfig) ToYaml() ISpider {
	var returns ISpider

	_jsii_.Invoke(
		s,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

