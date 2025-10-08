package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the GraphQLConfig construct.
type IGraphQLProps interface {
	Graphql() IGraphQL
	SetGraphql(g IGraphQL)
}

// The jsii proxy for IGraphQLProps
type jsiiProxy_IGraphQLProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IGraphQLProps) Graphql() IGraphQL {
	var returns IGraphQL
	_jsii_.Get(
		j,
		"graphql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphQLProps)SetGraphql(val IGraphQL) {
	if err := j.validateSetGraphqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"graphql",
		val,
	)
}

