package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for an AJAX spider.
//
// Example:
//   const spiderAjax = new SpiderAjax({
//     context: 'MyContext',
//     url: 'https://example.com',
//     maxDuration: 10,
//     inScopeOnly: true,
//     elements: ['a', 'button'],
//     excludedElements: [
//       new ExcludedElement({
//         description: 'Exclude login button',
//         element: 'button',
//         text: 'Login'
//       })
//     ],
//     tests: [
//       new AjaxTest({
//         name: 'Check AJAX requests',
//         type: 'stats',
//         statistic: 'ajax.requests',
//         operator: '>',
//         value: 10,
//         onFail: 'warn'
//       })
//     ]
//   });
//
type SpiderAjax interface {
	ISpiderAjax
	BrowserId() *string
	SetBrowserId(val *string)
	ClickDefaultElems() *bool
	SetClickDefaultElems(val *bool)
	ClickElemsOnce() *bool
	SetClickElemsOnce(val *bool)
	Context() *string
	SetContext(val *string)
	Elements() *[]*string
	SetElements(val *[]*string)
	EventWait() *float64
	SetEventWait(val *float64)
	ExcludedElements() *[]IExcludedElement
	SetExcludedElements(val *[]IExcludedElement)
	InScopeOnly() *bool
	SetInScopeOnly(val *bool)
	MaxCrawlDepth() *float64
	SetMaxCrawlDepth(val *float64)
	MaxCrawlStates() *float64
	SetMaxCrawlStates(val *float64)
	MaxDuration() *float64
	SetMaxDuration(val *float64)
	NumberOfBrowsers() *float64
	SetNumberOfBrowsers(val *float64)
	RandomInputs() *bool
	SetRandomInputs(val *bool)
	ReloadWait() *float64
	SetReloadWait(val *float64)
	RunOnlyIfModern() *bool
	SetRunOnlyIfModern(val *bool)
	ScopeCheck() *string
	SetScopeCheck(val *string)
	Tests() *[]IAjaxTest
	SetTests(val *[]IAjaxTest)
	Url() *string
	SetUrl(val *string)
	User() *string
	SetUser(val *string)
}

// The jsii proxy struct for SpiderAjax
type jsiiProxy_SpiderAjax struct {
	jsiiProxy_ISpiderAjax
}

func (j *jsiiProxy_SpiderAjax) BrowserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) ClickDefaultElems() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"clickDefaultElems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) ClickElemsOnce() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"clickElemsOnce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) Elements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) EventWait() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) ExcludedElements() *[]IExcludedElement {
	var returns *[]IExcludedElement
	_jsii_.Get(
		j,
		"excludedElements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) InScopeOnly() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"inScopeOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) MaxCrawlDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCrawlDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) MaxCrawlStates() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCrawlStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) MaxDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) NumberOfBrowsers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrowsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) RandomInputs() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"randomInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) ReloadWait() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reloadWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) RunOnlyIfModern() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"runOnlyIfModern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) ScopeCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) Tests() *[]IAjaxTest {
	var returns *[]IAjaxTest
	_jsii_.Get(
		j,
		"tests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderAjax) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


// Creates an instance of SpiderAjax.
//
// Example:
//   const spiderAjax = new SpiderAjax({
//     context: 'MyContext',
//     url: 'https://example.com',
//     maxDuration: 10,
//     inScopeOnly: true,
//     elements: ['a', 'button'],
//     excludedElements: [
//       new ExcludedElement({
//         description: 'Exclude login button',
//         element: 'button',
//         text: 'Login'
//       })
//     ],
//     tests: [
//       new AjaxTest({
//         name: 'Check AJAX requests',
//         type: 'stats',
//         statistic: 'ajax.requests',
//         operator: '>',
//         value: 10,
//         onFail: 'warn'
//       })
//     ]
//   });
//
func NewSpiderAjax(options ISpiderAjax) SpiderAjax {
	_init_.Initialize()

	if err := validateNewSpiderAjaxParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpiderAjax{}

	_jsii_.Create(
		"zap-cdk.SpiderAjax",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of SpiderAjax.
//
// Example:
//   const spiderAjax = new SpiderAjax({
//     context: 'MyContext',
//     url: 'https://example.com',
//     maxDuration: 10,
//     inScopeOnly: true,
//     elements: ['a', 'button'],
//     excludedElements: [
//       new ExcludedElement({
//         description: 'Exclude login button',
//         element: 'button',
//         text: 'Login'
//       })
//     ],
//     tests: [
//       new AjaxTest({
//         name: 'Check AJAX requests',
//         type: 'stats',
//         statistic: 'ajax.requests',
//         operator: '>',
//         value: 10,
//         onFail: 'warn'
//       })
//     ]
//   });
//
func NewSpiderAjax_Override(s SpiderAjax, options ISpiderAjax) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SpiderAjax",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SpiderAjax)SetBrowserId(val *string) {
	_jsii_.Set(
		j,
		"browserId",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetClickDefaultElems(val *bool) {
	_jsii_.Set(
		j,
		"clickDefaultElems",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetClickElemsOnce(val *bool) {
	_jsii_.Set(
		j,
		"clickElemsOnce",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetElements(val *[]*string) {
	_jsii_.Set(
		j,
		"elements",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetEventWait(val *float64) {
	_jsii_.Set(
		j,
		"eventWait",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetExcludedElements(val *[]IExcludedElement) {
	_jsii_.Set(
		j,
		"excludedElements",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetInScopeOnly(val *bool) {
	_jsii_.Set(
		j,
		"inScopeOnly",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetMaxCrawlDepth(val *float64) {
	_jsii_.Set(
		j,
		"maxCrawlDepth",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetMaxCrawlStates(val *float64) {
	_jsii_.Set(
		j,
		"maxCrawlStates",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetMaxDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetNumberOfBrowsers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfBrowsers",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetRandomInputs(val *bool) {
	_jsii_.Set(
		j,
		"randomInputs",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetReloadWait(val *float64) {
	_jsii_.Set(
		j,
		"reloadWait",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetRunOnlyIfModern(val *bool) {
	_jsii_.Set(
		j,
		"runOnlyIfModern",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetScopeCheck(val *string) {
	_jsii_.Set(
		j,
		"scopeCheck",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetTests(val *[]IAjaxTest) {
	_jsii_.Set(
		j,
		"tests",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_SpiderAjax)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

