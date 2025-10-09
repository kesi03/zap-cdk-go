package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the OpenAPI import configuration.
//
// Example:
//   const openApiConfig = new OpenAPI({
//     apiFile: 'api-definition.yaml',
//     context: 'MyContext',
//     user: 'apiUser',
//     targetUrl: 'https://api.example.com'
//   });
//
type OpenAPI interface {
	IOpenAPI
	ApiFile() *string
	SetApiFile(val *string)
	ApiUrl() *string
	SetApiUrl(val *string)
	Context() *string
	SetContext(val *string)
	TargetUrl() *string
	SetTargetUrl(val *string)
	User() *string
	SetUser(val *string)
}

// The jsii proxy struct for OpenAPI
type jsiiProxy_OpenAPI struct {
	jsiiProxy_IOpenAPI
}

func (j *jsiiProxy_OpenAPI) ApiFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenAPI) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenAPI) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenAPI) TargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenAPI) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


// Creates an instance of OpenAPI.
func NewOpenAPI(options IOpenAPI) OpenAPI {
	_init_.Initialize()

	if err := validateNewOpenAPIParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenAPI{}

	_jsii_.Create(
		"zap-cdk.OpenAPI",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of OpenAPI.
func NewOpenAPI_Override(o OpenAPI, options IOpenAPI) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.OpenAPI",
		[]interface{}{options},
		o,
	)
}

func (j *jsiiProxy_OpenAPI)SetApiFile(val *string) {
	_jsii_.Set(
		j,
		"apiFile",
		val,
	)
}

func (j *jsiiProxy_OpenAPI)SetApiUrl(val *string) {
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_OpenAPI)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_OpenAPI)SetTargetUrl(val *string) {
	_jsii_.Set(
		j,
		"targetUrl",
		val,
	)
}

func (j *jsiiProxy_OpenAPI)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

