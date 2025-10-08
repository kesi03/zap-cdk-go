package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the GraphQL configuration.
type GraphQLConfig interface {
	constructs.Construct
	Config() IGraphQL
	SetConfig(val IGraphQL)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	ToYaml() IGraphQL
}

// The jsii proxy struct for GraphQLConfig
type jsiiProxy_GraphQLConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_GraphQLConfig) Config() IGraphQL {
	var returns IGraphQL
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQLConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of GraphQLConfig.
func NewGraphQLConfig(scope constructs.Construct, id *string, props IGraphQLProps) GraphQLConfig {
	_init_.Initialize()

	if err := validateNewGraphQLConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GraphQLConfig{}

	_jsii_.Create(
		"zap-cdk.GraphQLConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of GraphQLConfig.
func NewGraphQLConfig_Override(g GraphQLConfig, scope constructs.Construct, id *string, props IGraphQLProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.GraphQLConfig",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GraphQLConfig)SetConfig(val IGraphQL) {
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
func GraphQLConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGraphQLConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.GraphQLConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphQLConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GraphQLConfig) ToYaml() IGraphQL {
	var returns IGraphQL

	_jsii_.Invoke(
		g,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

