package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for input vectors used in an active scan.
//
// Example:
//   const inputVectorsConfig = new InputVectors({
//     urlQueryStringAndDataDrivenNodes: new UrlQueryStringAndDataDrivenNodes(),
//     postData: new PostData(),
//
type InputVectors interface {
	IInputVectors
	// Configuration for cookie data scanning.
	CookieData() ICookieData
	SetCookieData(val ICookieData)
	// Configuration for HTTP header scanning.
	HttpHeaders() IHttpHeaders
	SetHttpHeaders(val IHttpHeaders)
	// Configuration for POST data scanning.
	PostData() IPostData
	SetPostData(val IPostData)
	// If Input Vector scripts should be used.
	//
	// Default: true.
	Scripts() *bool
	SetScripts(val *bool)
	// If URL path segments should be scanned.
	//
	// Default: false.
	UrlPath() *bool
	SetUrlPath(val *bool)
	// Configuration for query parameters and data-driven nodes.
	UrlQueryStringAndDataDrivenNodes() IUrlQueryStringAndDataDrivenNodes
	SetUrlQueryStringAndDataDrivenNodes(val IUrlQueryStringAndDataDrivenNodes)
}

// The jsii proxy struct for InputVectors
type jsiiProxy_InputVectors struct {
	jsiiProxy_IInputVectors
}

func (j *jsiiProxy_InputVectors) CookieData() ICookieData {
	var returns ICookieData
	_jsii_.Get(
		j,
		"cookieData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputVectors) HttpHeaders() IHttpHeaders {
	var returns IHttpHeaders
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputVectors) PostData() IPostData {
	var returns IPostData
	_jsii_.Get(
		j,
		"postData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputVectors) Scripts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputVectors) UrlPath() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"urlPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputVectors) UrlQueryStringAndDataDrivenNodes() IUrlQueryStringAndDataDrivenNodes {
	var returns IUrlQueryStringAndDataDrivenNodes
	_jsii_.Get(
		j,
		"urlQueryStringAndDataDrivenNodes",
		&returns,
	)
	return returns
}


// Creates an instance of InputVectors.
//
// Example:
//   const inputVectorsConfig = new InputVectors({
//     urlQueryStringAndDataDrivenNodes: new UrlQueryStringAndDataDrivenNodes(),
//     postData: new PostData(),
//     urlPath: false,
//     httpHeaders: new HttpHeaders(),
//     cookieData: new CookieData(),
//     scripts: true
//   });
//   console.log(inputVectorsConfig.postData.enabled); // true
//
func NewInputVectors(options IInputVectors) InputVectors {
	_init_.Initialize()

	j := jsiiProxy_InputVectors{}

	_jsii_.Create(
		"zap-cdk.InputVectors",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of InputVectors.
//
// Example:
//   const inputVectorsConfig = new InputVectors({
//     urlQueryStringAndDataDrivenNodes: new UrlQueryStringAndDataDrivenNodes(),
//     postData: new PostData(),
//     urlPath: false,
//     httpHeaders: new HttpHeaders(),
//     cookieData: new CookieData(),
//     scripts: true
//   });
//   console.log(inputVectorsConfig.postData.enabled); // true
//
func NewInputVectors_Override(i InputVectors, options IInputVectors) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.InputVectors",
		[]interface{}{options},
		i,
	)
}

func (j *jsiiProxy_InputVectors)SetCookieData(val ICookieData) {
	if err := j.validateSetCookieDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieData",
		val,
	)
}

func (j *jsiiProxy_InputVectors)SetHttpHeaders(val IHttpHeaders) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_InputVectors)SetPostData(val IPostData) {
	if err := j.validateSetPostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postData",
		val,
	)
}

func (j *jsiiProxy_InputVectors)SetScripts(val *bool) {
	_jsii_.Set(
		j,
		"scripts",
		val,
	)
}

func (j *jsiiProxy_InputVectors)SetUrlPath(val *bool) {
	_jsii_.Set(
		j,
		"urlPath",
		val,
	)
}

func (j *jsiiProxy_InputVectors)SetUrlQueryStringAndDataDrivenNodes(val IUrlQueryStringAndDataDrivenNodes) {
	if err := j.validateSetUrlQueryStringAndDataDrivenNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlQueryStringAndDataDrivenNodes",
		val,
	)
}

