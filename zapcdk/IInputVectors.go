package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the configuration for input vectors used in an active scan.
type IInputVectors interface {
	// Configuration for cookie data scanning.
	CookieData() ICookieData
	SetCookieData(c ICookieData)
	// Configuration for HTTP header scanning.
	HttpHeaders() IHttpHeaders
	SetHttpHeaders(h IHttpHeaders)
	// Configuration for POST data scanning.
	PostData() IPostData
	SetPostData(p IPostData)
	// If Input Vector scripts should be used.
	//
	// Default: true.
	Scripts() *bool
	SetScripts(s *bool)
	// If URL path segments should be scanned.
	//
	// Default: false.
	UrlPath() *bool
	SetUrlPath(u *bool)
	// Configuration for query parameters and data-driven nodes.
	UrlQueryStringAndDataDrivenNodes() IUrlQueryStringAndDataDrivenNodes
	SetUrlQueryStringAndDataDrivenNodes(u IUrlQueryStringAndDataDrivenNodes)
}

// The jsii proxy for IInputVectors
type jsiiProxy_IInputVectors struct {
	_ byte // padding
}

func (j *jsiiProxy_IInputVectors) CookieData() ICookieData {
	var returns ICookieData
	_jsii_.Get(
		j,
		"cookieData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputVectors)SetCookieData(val ICookieData) {
	if err := j.validateSetCookieDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieData",
		val,
	)
}

func (j *jsiiProxy_IInputVectors) HttpHeaders() IHttpHeaders {
	var returns IHttpHeaders
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputVectors)SetHttpHeaders(val IHttpHeaders) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_IInputVectors) PostData() IPostData {
	var returns IPostData
	_jsii_.Get(
		j,
		"postData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputVectors)SetPostData(val IPostData) {
	if err := j.validateSetPostDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postData",
		val,
	)
}

func (j *jsiiProxy_IInputVectors) Scripts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scripts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputVectors)SetScripts(val *bool) {
	_jsii_.Set(
		j,
		"scripts",
		val,
	)
}

func (j *jsiiProxy_IInputVectors) UrlPath() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"urlPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputVectors)SetUrlPath(val *bool) {
	_jsii_.Set(
		j,
		"urlPath",
		val,
	)
}

func (j *jsiiProxy_IInputVectors) UrlQueryStringAndDataDrivenNodes() IUrlQueryStringAndDataDrivenNodes {
	var returns IUrlQueryStringAndDataDrivenNodes
	_jsii_.Get(
		j,
		"urlQueryStringAndDataDrivenNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputVectors)SetUrlQueryStringAndDataDrivenNodes(val IUrlQueryStringAndDataDrivenNodes) {
	if err := j.validateSetUrlQueryStringAndDataDrivenNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlQueryStringAndDataDrivenNodes",
		val,
	)
}

