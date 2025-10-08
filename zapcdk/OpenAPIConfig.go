package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the OpenAPI configuration.
type OpenAPIConfig interface {
	constructs.Construct
	Config() IOpenAPI
	SetConfig(val IOpenAPI)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	ToYaml() IOpenAPI
}

// The jsii proxy struct for OpenAPIConfig
type jsiiProxy_OpenAPIConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_OpenAPIConfig) Config() IOpenAPI {
	var returns IOpenAPI
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenAPIConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of OpenAPIConfig.
func NewOpenAPIConfig(scope constructs.Construct, id *string, props IOpenAPIProps) OpenAPIConfig {
	_init_.Initialize()

	if err := validateNewOpenAPIConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenAPIConfig{}

	_jsii_.Create(
		"zap-cdk.OpenAPIConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of OpenAPIConfig.
func NewOpenAPIConfig_Override(o OpenAPIConfig, scope constructs.Construct, id *string, props IOpenAPIProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.OpenAPIConfig",
		[]interface{}{scope, id, props},
		o,
	)
}

func (j *jsiiProxy_OpenAPIConfig)SetConfig(val IOpenAPI) {
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
func OpenAPIConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenAPIConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.OpenAPIConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenAPIConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenAPIConfig) ToYaml() IOpenAPI {
	var returns IOpenAPI

	_jsii_.Invoke(
		o,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

