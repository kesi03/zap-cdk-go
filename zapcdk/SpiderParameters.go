package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type SpiderParameters interface {
	ISpiderParameters
	AcceptCookies() *bool
	SetAcceptCookies(val *bool)
	Context() *string
	SetContext(val *string)
	HandleODataParametersVisited() *bool
	SetHandleODataParametersVisited(val *bool)
	HandleParameters() *string
	SetHandleParameters(val *string)
	LogoutAvoidance() *bool
	SetLogoutAvoidance(val *bool)
	MaxChildren() *float64
	SetMaxChildren(val *float64)
	MaxDepth() *float64
	SetMaxDepth(val *float64)
	MaxDuration() *float64
	SetMaxDuration(val *float64)
	MaxParseSizeBytes() *float64
	SetMaxParseSizeBytes(val *float64)
	ParseComments() *bool
	SetParseComments(val *bool)
	ParseDsStore() *bool
	SetParseDsStore(val *bool)
	ParseGit() *bool
	SetParseGit(val *bool)
	ParseRobotsTxt() *bool
	SetParseRobotsTxt(val *bool)
	ParseSitemapXml() *bool
	SetParseSitemapXml(val *bool)
	ParseSVNEntries() *bool
	SetParseSVNEntries(val *bool)
	PostForm() *bool
	SetPostForm(val *bool)
	ProcessForm() *bool
	SetProcessForm(val *bool)
	SendRefererHeader() *bool
	SetSendRefererHeader(val *bool)
	ThreadCount() *float64
	SetThreadCount(val *float64)
	Url() *string
	SetUrl(val *string)
	User() *string
	SetUser(val *string)
	UserAgent() *string
	SetUserAgent(val *string)
}

// The jsii proxy struct for SpiderParameters
type jsiiProxy_SpiderParameters struct {
	jsiiProxy_ISpiderParameters
}

func (j *jsiiProxy_SpiderParameters) AcceptCookies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"acceptCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) HandleODataParametersVisited() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"handleODataParametersVisited",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) HandleParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) LogoutAvoidance() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logoutAvoidance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) MaxChildren() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxChildren",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) MaxDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) MaxDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) MaxParseSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParseSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ParseComments() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseComments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ParseDsStore() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseDsStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ParseGit() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseGit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ParseRobotsTxt() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseRobotsTxt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ParseSitemapXml() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseSitemapXml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ParseSVNEntries() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseSVNEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) PostForm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"postForm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ProcessForm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"processForm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) SendRefererHeader() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sendRefererHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) ThreadCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderParameters) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}


// Creates an instance of SpiderParameters.
func NewSpiderParameters(options ISpiderParameters) SpiderParameters {
	_init_.Initialize()

	if err := validateNewSpiderParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpiderParameters{}

	_jsii_.Create(
		"zap-cdk.SpiderParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of SpiderParameters.
func NewSpiderParameters_Override(s SpiderParameters, options ISpiderParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SpiderParameters",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SpiderParameters)SetAcceptCookies(val *bool) {
	_jsii_.Set(
		j,
		"acceptCookies",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetHandleODataParametersVisited(val *bool) {
	_jsii_.Set(
		j,
		"handleODataParametersVisited",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetHandleParameters(val *string) {
	_jsii_.Set(
		j,
		"handleParameters",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetLogoutAvoidance(val *bool) {
	_jsii_.Set(
		j,
		"logoutAvoidance",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetMaxChildren(val *float64) {
	_jsii_.Set(
		j,
		"maxChildren",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetMaxDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxDepth",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetMaxDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetMaxParseSizeBytes(val *float64) {
	_jsii_.Set(
		j,
		"maxParseSizeBytes",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetParseComments(val *bool) {
	_jsii_.Set(
		j,
		"parseComments",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetParseDsStore(val *bool) {
	_jsii_.Set(
		j,
		"parseDsStore",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetParseGit(val *bool) {
	_jsii_.Set(
		j,
		"parseGit",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetParseRobotsTxt(val *bool) {
	_jsii_.Set(
		j,
		"parseRobotsTxt",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetParseSitemapXml(val *bool) {
	_jsii_.Set(
		j,
		"parseSitemapXml",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetParseSVNEntries(val *bool) {
	_jsii_.Set(
		j,
		"parseSVNEntries",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetPostForm(val *bool) {
	_jsii_.Set(
		j,
		"postForm",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetProcessForm(val *bool) {
	_jsii_.Set(
		j,
		"processForm",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetSendRefererHeader(val *bool) {
	_jsii_.Set(
		j,
		"sendRefererHeader",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetThreadCount(val *float64) {
	_jsii_.Set(
		j,
		"threadCount",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_SpiderParameters)SetUserAgent(val *string) {
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

