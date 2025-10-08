package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing the parameters for an AJAX spider configuration.
type ISpiderAjax interface {
	BrowserId() *string
	SetBrowserId(b *string)
	ClickDefaultElems() *bool
	SetClickDefaultElems(c *bool)
	ClickElemsOnce() *bool
	SetClickElemsOnce(c *bool)
	Context() *string
	SetContext(c *string)
	Elements() *[]*string
	SetElements(e *[]*string)
	EventWait() *float64
	SetEventWait(e *float64)
	ExcludedElements() *[]IExcludedElement
	SetExcludedElements(e *[]IExcludedElement)
	InScopeOnly() *bool
	SetInScopeOnly(i *bool)
	MaxCrawlDepth() *float64
	SetMaxCrawlDepth(m *float64)
	MaxCrawlStates() *float64
	SetMaxCrawlStates(m *float64)
	MaxDuration() *float64
	SetMaxDuration(m *float64)
	NumberOfBrowsers() *float64
	SetNumberOfBrowsers(n *float64)
	RandomInputs() *bool
	SetRandomInputs(r *bool)
	ReloadWait() *float64
	SetReloadWait(r *float64)
	RunOnlyIfModern() *bool
	SetRunOnlyIfModern(r *bool)
	ScopeCheck() *string
	SetScopeCheck(s *string)
	Tests() *[]IAjaxTest
	SetTests(t *[]IAjaxTest)
	Url() *string
	SetUrl(u *string)
	User() *string
	SetUser(u *string)
}

// The jsii proxy for ISpiderAjax
type jsiiProxy_ISpiderAjax struct {
	_ byte // padding
}

func (j *jsiiProxy_ISpiderAjax) BrowserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetBrowserId(val *string) {
	_jsii_.Set(
		j,
		"browserId",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) ClickDefaultElems() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"clickDefaultElems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetClickDefaultElems(val *bool) {
	_jsii_.Set(
		j,
		"clickDefaultElems",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) ClickElemsOnce() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"clickElemsOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetClickElemsOnce(val *bool) {
	_jsii_.Set(
		j,
		"clickElemsOnce",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) Elements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetElements(val *[]*string) {
	_jsii_.Set(
		j,
		"elements",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) EventWait() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetEventWait(val *float64) {
	_jsii_.Set(
		j,
		"eventWait",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) ExcludedElements() *[]IExcludedElement {
	var returns *[]IExcludedElement
	_jsii_.Get(
		j,
		"excludedElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetExcludedElements(val *[]IExcludedElement) {
	_jsii_.Set(
		j,
		"excludedElements",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) InScopeOnly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"inScopeOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetInScopeOnly(val *bool) {
	_jsii_.Set(
		j,
		"inScopeOnly",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) MaxCrawlDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCrawlDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetMaxCrawlDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxCrawlDepth",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) MaxCrawlStates() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCrawlStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetMaxCrawlStates(val *float64) {
	_jsii_.Set(
		j,
		"maxCrawlStates",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) MaxDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetMaxDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) NumberOfBrowsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrowsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetNumberOfBrowsers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBrowsers",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) RandomInputs() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"randomInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetRandomInputs(val *bool) {
	_jsii_.Set(
		j,
		"randomInputs",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) ReloadWait() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reloadWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetReloadWait(val *float64) {
	_jsii_.Set(
		j,
		"reloadWait",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) RunOnlyIfModern() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"runOnlyIfModern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetRunOnlyIfModern(val *bool) {
	_jsii_.Set(
		j,
		"runOnlyIfModern",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) ScopeCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetScopeCheck(val *string) {
	_jsii_.Set(
		j,
		"scopeCheck",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) Tests() *[]IAjaxTest {
	var returns *[]IAjaxTest
	_jsii_.Get(
		j,
		"tests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetTests(val *[]IAjaxTest) {
	_jsii_.Set(
		j,
		"tests",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_ISpiderAjax) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderAjax)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

