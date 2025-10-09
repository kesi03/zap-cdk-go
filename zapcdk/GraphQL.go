package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type GraphQL interface {
	IGraphQL
	ArgsType() *string
	SetArgsType(val *string)
	Endpoint() *string
	SetEndpoint(val *string)
	LenientMaxQueryDepthEnabled() *bool
	SetLenientMaxQueryDepthEnabled(val *bool)
	MaxAdditionalQueryDepth() *float64
	SetMaxAdditionalQueryDepth(val *float64)
	MaxArgsDepth() *float64
	SetMaxArgsDepth(val *float64)
	MaxQueryDepth() *float64
	SetMaxQueryDepth(val *float64)
	OptionalArgsEnabled() *bool
	SetOptionalArgsEnabled(val *bool)
	QueryGenEnabled() *bool
	SetQueryGenEnabled(val *bool)
	QuerySplitType() *string
	SetQuerySplitType(val *string)
	RequestMethod() *string
	SetRequestMethod(val *string)
	SchemaFile() *string
	SetSchemaFile(val *string)
	SchemaUrl() *string
	SetSchemaUrl(val *string)
}

// The jsii proxy struct for GraphQL
type jsiiProxy_GraphQL struct {
	jsiiProxy_IGraphQL
}

func (j *jsiiProxy_GraphQL) ArgsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"argsType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) LenientMaxQueryDepthEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"lenientMaxQueryDepthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) MaxAdditionalQueryDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAdditionalQueryDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) MaxArgsDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxArgsDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) MaxQueryDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxQueryDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) OptionalArgsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"optionalArgsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) QueryGenEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"queryGenEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) QuerySplitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"querySplitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) RequestMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) SchemaFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphQL) SchemaUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaUrl",
		&returns,
	)
	return returns
}


// Creates an instance of GraphQL.
func NewGraphQL(options IGraphQL) GraphQL {
	_init_.Initialize()

	if err := validateNewGraphQLParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_GraphQL{}

	_jsii_.Create(
		"zap-cdk.GraphQL",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of GraphQL.
func NewGraphQL_Override(g GraphQL, options IGraphQL) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.GraphQL",
		[]interface{}{options},
		g,
	)
}

func (j *jsiiProxy_GraphQL)SetArgsType(val *string) {
	_jsii_.Set(
		j,
		"argsType",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetEndpoint(val *string) {
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetLenientMaxQueryDepthEnabled(val *bool) {
	_jsii_.Set(
		j,
		"lenientMaxQueryDepthEnabled",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetMaxAdditionalQueryDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxAdditionalQueryDepth",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetMaxArgsDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxArgsDepth",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetMaxQueryDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxQueryDepth",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetOptionalArgsEnabled(val *bool) {
	_jsii_.Set(
		j,
		"optionalArgsEnabled",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetQueryGenEnabled(val *bool) {
	_jsii_.Set(
		j,
		"queryGenEnabled",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetQuerySplitType(val *string) {
	_jsii_.Set(
		j,
		"querySplitType",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetRequestMethod(val *string) {
	_jsii_.Set(
		j,
		"requestMethod",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetSchemaFile(val *string) {
	_jsii_.Set(
		j,
		"schemaFile",
		val,
	)
}

func (j *jsiiProxy_GraphQL)SetSchemaUrl(val *string) {
	_jsii_.Set(
		j,
		"schemaUrl",
		val,
	)
}

