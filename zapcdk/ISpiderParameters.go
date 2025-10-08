package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing the parameters for a spider configuration.
type ISpiderParameters interface {
	AcceptCookies() *bool
	SetAcceptCookies(a *bool)
	Context() *string
	SetContext(c *string)
	HandleODataParametersVisited() *bool
	SetHandleODataParametersVisited(h *bool)
	HandleParameters() *string
	SetHandleParameters(h *string)
	LogoutAvoidance() *bool
	SetLogoutAvoidance(l *bool)
	MaxChildren() *float64
	SetMaxChildren(m *float64)
	MaxDepth() *float64
	SetMaxDepth(m *float64)
	MaxDuration() *float64
	SetMaxDuration(m *float64)
	MaxParseSizeBytes() *float64
	SetMaxParseSizeBytes(m *float64)
	ParseComments() *bool
	SetParseComments(p *bool)
	ParseDsStore() *bool
	SetParseDsStore(p *bool)
	ParseGit() *bool
	SetParseGit(p *bool)
	ParseRobotsTxt() *bool
	SetParseRobotsTxt(p *bool)
	ParseSitemapXml() *bool
	SetParseSitemapXml(p *bool)
	ParseSVNEntries() *bool
	SetParseSVNEntries(p *bool)
	PostForm() *bool
	SetPostForm(p *bool)
	ProcessForm() *bool
	SetProcessForm(p *bool)
	SendRefererHeader() *bool
	SetSendRefererHeader(s *bool)
	ThreadCount() *float64
	SetThreadCount(t *float64)
	Url() *string
	SetUrl(u *string)
	User() *string
	SetUser(u *string)
	UserAgent() *string
	SetUserAgent(u *string)
}

// The jsii proxy for ISpiderParameters
type jsiiProxy_ISpiderParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_ISpiderParameters) AcceptCookies() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"acceptCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetAcceptCookies(val *bool) {
	_jsii_.Set(
		j,
		"acceptCookies",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) HandleODataParametersVisited() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"handleODataParametersVisited",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetHandleODataParametersVisited(val *bool) {
	_jsii_.Set(
		j,
		"handleODataParametersVisited",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) HandleParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handleParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetHandleParameters(val *string) {
	_jsii_.Set(
		j,
		"handleParameters",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) LogoutAvoidance() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"logoutAvoidance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetLogoutAvoidance(val *bool) {
	_jsii_.Set(
		j,
		"logoutAvoidance",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) MaxChildren() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxChildren",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetMaxChildren(val *float64) {
	_jsii_.Set(
		j,
		"maxChildren",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) MaxDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetMaxDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxDepth",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) MaxDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetMaxDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) MaxParseSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParseSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetMaxParseSizeBytes(val *float64) {
	_jsii_.Set(
		j,
		"maxParseSizeBytes",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ParseComments() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseComments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetParseComments(val *bool) {
	_jsii_.Set(
		j,
		"parseComments",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ParseDsStore() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseDsStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetParseDsStore(val *bool) {
	_jsii_.Set(
		j,
		"parseDsStore",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ParseGit() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseGit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetParseGit(val *bool) {
	_jsii_.Set(
		j,
		"parseGit",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ParseRobotsTxt() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseRobotsTxt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetParseRobotsTxt(val *bool) {
	_jsii_.Set(
		j,
		"parseRobotsTxt",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ParseSitemapXml() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseSitemapXml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetParseSitemapXml(val *bool) {
	_jsii_.Set(
		j,
		"parseSitemapXml",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ParseSVNEntries() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parseSVNEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetParseSVNEntries(val *bool) {
	_jsii_.Set(
		j,
		"parseSVNEntries",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) PostForm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"postForm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetPostForm(val *bool) {
	_jsii_.Set(
		j,
		"postForm",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ProcessForm() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"processForm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetProcessForm(val *bool) {
	_jsii_.Set(
		j,
		"processForm",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) SendRefererHeader() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"sendRefererHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetSendRefererHeader(val *bool) {
	_jsii_.Set(
		j,
		"sendRefererHeader",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) ThreadCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetThreadCount(val *float64) {
	_jsii_.Set(
		j,
		"threadCount",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (j *jsiiProxy_ISpiderParameters) UserAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderParameters)SetUserAgent(val *string) {
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

