package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the Postman configuration.
type PostmanConfig interface {
	constructs.Construct
	Config() IPostman
	SetConfig(val IPostman)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the Postman configuration to YAML format.
	//
	// Returns: The Postman configuration in YAML format.
	ToYaml() IPostman
}

// The jsii proxy struct for PostmanConfig
type jsiiProxy_PostmanConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PostmanConfig) Config() IPostman {
	var returns IPostman
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostmanConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of PostmanConfig.
func NewPostmanConfig(scope constructs.Construct, id *string, props IPostmanProps) PostmanConfig {
	_init_.Initialize()

	if err := validateNewPostmanConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PostmanConfig{}

	_jsii_.Create(
		"zap-cdk.PostmanConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of PostmanConfig.
func NewPostmanConfig_Override(p PostmanConfig, scope constructs.Construct, id *string, props IPostmanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PostmanConfig",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_PostmanConfig)SetConfig(val IPostman) {
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
func PostmanConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePostmanConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.PostmanConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostmanConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PostmanConfig) ToYaml() IPostman {
	var returns IPostman

	_jsii_.Invoke(
		p,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

