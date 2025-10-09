// zap-cdk
package zapcdk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScan",
		reflect.TypeOf((*ActiveScan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "policyDefinition", GoGetter: "PolicyDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScan{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScan)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanConfig",
		reflect.TypeOf((*ActiveScanConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "excludePaths", GoGetter: "ExcludePaths"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScanConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanConfigConstruct",
		reflect.TypeOf((*ActiveScanConfigConstruct)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanConfigConstruct{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanConfigParameters",
		reflect.TypeOf((*ActiveScanConfigParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultPolicy", GoGetter: "DefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "handleAntiCSRFTokens", GoGetter: "HandleAntiCSRFTokens"},
			_jsii_.MemberProperty{JsiiProperty: "injectPluginIdInHeader", GoGetter: "InjectPluginIdInHeader"},
			_jsii_.MemberProperty{JsiiProperty: "inputVectors", GoGetter: "InputVectors"},
			_jsii_.MemberProperty{JsiiProperty: "maxAlertsPerRule", GoGetter: "MaxAlertsPerRule"},
			_jsii_.MemberProperty{JsiiProperty: "maxRuleDurationInMins", GoGetter: "MaxRuleDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "maxScanDurationInMins", GoGetter: "MaxScanDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "threadPerHost", GoGetter: "ThreadPerHost"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanConfigParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScanConfigParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanJob",
		reflect.TypeOf((*ActiveScanJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "job", GoGetter: "Job"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanJob{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanParameters",
		reflect.TypeOf((*ActiveScanParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addQueryParam", GoGetter: "AddQueryParam"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPolicy", GoGetter: "DefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "delayInMs", GoGetter: "DelayInMs"},
			_jsii_.MemberProperty{JsiiProperty: "handleAntiCSRFTokens", GoGetter: "HandleAntiCSRFTokens"},
			_jsii_.MemberProperty{JsiiProperty: "injectPluginIdInHeader", GoGetter: "InjectPluginIdInHeader"},
			_jsii_.MemberProperty{JsiiProperty: "maxAlertsPerRule", GoGetter: "MaxAlertsPerRule"},
			_jsii_.MemberProperty{JsiiProperty: "maxRuleDurationInMins", GoGetter: "MaxRuleDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "maxScanDurationInMins", GoGetter: "MaxScanDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "scanHeadersAllRequests", GoGetter: "ScanHeadersAllRequests"},
			_jsii_.MemberProperty{JsiiProperty: "tests", GoGetter: "Tests"},
			_jsii_.MemberProperty{JsiiProperty: "threadPerHost", GoGetter: "ThreadPerHost"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScanParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanPolicy",
		reflect.TypeOf((*ActiveScanPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanPolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScanPolicy)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanPolicyConfig",
		reflect.TypeOf((*ActiveScanPolicyConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanPolicyConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanPolicyDefinition",
		reflect.TypeOf((*ActiveScanPolicyDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanPolicyDefinition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScanPolicyDefinition)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ActiveScanPolicyParameters",
		reflect.TypeOf((*ActiveScanPolicyParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "policyDefinition", GoGetter: "PolicyDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_ActiveScanPolicyParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IActiveScanPolicyParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AjaxTest",
		reflect.TypeOf((*AjaxTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_AjaxTest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAjaxTest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AlertFilter",
		reflect.TypeOf((*AlertFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attack", GoGetter: "Attack"},
			_jsii_.MemberProperty{JsiiProperty: "attackRegex", GoGetter: "AttackRegex"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "evidence", GoGetter: "Evidence"},
			_jsii_.MemberProperty{JsiiProperty: "evidenceRegex", GoGetter: "EvidenceRegex"},
			_jsii_.MemberProperty{JsiiProperty: "newRisk", GoGetter: "NewRisk"},
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterRegex", GoGetter: "ParameterRegex"},
			_jsii_.MemberProperty{JsiiProperty: "ruleId", GoGetter: "RuleId"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlRegex", GoGetter: "UrlRegex"},
		},
		func() interface{} {
			j := jsiiProxy_AlertFilter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertFilter)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AlertFilterParameters",
		reflect.TypeOf((*AlertFilterParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertFilters", GoGetter: "AlertFilters"},
			_jsii_.MemberProperty{JsiiProperty: "deleteGlobalAlerts", GoGetter: "DeleteGlobalAlerts"},
		},
		func() interface{} {
			j := jsiiProxy_AlertFilterParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertFilterParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AlertTag",
		reflect.TypeOf((*AlertTag)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			j := jsiiProxy_AlertTag{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertTag)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AlertTags",
		reflect.TypeOf((*AlertTags)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			j := jsiiProxy_AlertTags{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertTags)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AlertTest",
		reflect.TypeOf((*AlertTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "alertName", GoGetter: "AlertName"},
			_jsii_.MemberProperty{JsiiProperty: "attack", GoGetter: "Attack"},
			_jsii_.MemberProperty{JsiiProperty: "confidence", GoGetter: "Confidence"},
			_jsii_.MemberProperty{JsiiProperty: "evidence", GoGetter: "Evidence"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "param", GoGetter: "Param"},
			_jsii_.MemberProperty{JsiiProperty: "risk", GoGetter: "Risk"},
			_jsii_.MemberProperty{JsiiProperty: "scanRuleId", GoGetter: "ScanRuleId"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_AlertTest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAlertTest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.App",
		reflect.TypeOf((*App)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_App{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AuthenticationParameters",
		reflect.TypeOf((*AuthenticationParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "verification", GoGetter: "Verification"},
		},
		func() interface{} {
			j := jsiiProxy_AuthenticationParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAuthenticationParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AuthenticationParametersParameters",
		reflect.TypeOf((*AuthenticationParametersParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "loginPageUrl", GoGetter: "LoginPageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "loginRequestBody", GoGetter: "LoginRequestBody"},
			_jsii_.MemberProperty{JsiiProperty: "loginRequestUrl", GoGetter: "LoginRequestUrl"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "realm", GoGetter: "Realm"},
			_jsii_.MemberProperty{JsiiProperty: "script", GoGetter: "Script"},
			_jsii_.MemberProperty{JsiiProperty: "scriptEngine", GoGetter: "ScriptEngine"},
			_jsii_.MemberProperty{JsiiProperty: "scriptInline", GoGetter: "ScriptInline"},
		},
		func() interface{} {
			j := jsiiProxy_AuthenticationParametersParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAuthenticationParametersParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.AuthenticationParametersVerification",
		reflect.TypeOf((*AuthenticationParametersVerification)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "loggedInRegex", GoGetter: "LoggedInRegex"},
			_jsii_.MemberProperty{JsiiProperty: "loggedOutRegex", GoGetter: "LoggedOutRegex"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "pollAdditionalHeaders", GoGetter: "PollAdditionalHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "pollFrequency", GoGetter: "PollFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "pollPostData", GoGetter: "PollPostData"},
			_jsii_.MemberProperty{JsiiProperty: "pollUnits", GoGetter: "PollUnits"},
			_jsii_.MemberProperty{JsiiProperty: "pollUrl", GoGetter: "PollUrl"},
		},
		func() interface{} {
			j := jsiiProxy_AuthenticationParametersVerification{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAuthenticationParametersVerification)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ContextStructure",
		reflect.TypeOf((*ContextStructure)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataDrivenNodes", GoGetter: "DataDrivenNodes"},
			_jsii_.MemberProperty{JsiiProperty: "structuralParameters", GoGetter: "StructuralParameters"},
		},
		func() interface{} {
			j := jsiiProxy_ContextStructure{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IContextStructure)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ContextUser",
		reflect.TypeOf((*ContextUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			j := jsiiProxy_ContextUser{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IContextUser)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.CookieData",
		reflect.TypeOf((*CookieData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "encodeCookieValues", GoGetter: "EncodeCookieValues"},
		},
		func() interface{} {
			j := jsiiProxy_CookieData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICookieData)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.DataDrivenNode",
		reflect.TypeOf((*DataDrivenNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
		},
		func() interface{} {
			j := jsiiProxy_DataDrivenNode{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDataDrivenNode)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Delay",
		reflect.TypeOf((*Delay)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Delay{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDelay)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.DelayConfig",
		reflect.TypeOf((*DelayConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_DelayConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.DelayParameters",
		reflect.TypeOf((*DelayParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "time", GoGetter: "Time"},
		},
		func() interface{} {
			j := jsiiProxy_DelayParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDelayParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.EnvironmentConfig",
		reflect.TypeOf((*EnvironmentConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_EnvironmentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ExcludedElement",
		reflect.TypeOf((*ExcludedElement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeName", GoGetter: "AttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "element", GoGetter: "Element"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "xpath", GoGetter: "Xpath"},
		},
		func() interface{} {
			j := jsiiProxy_ExcludedElement{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExcludedElement)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ExitStatus",
		reflect.TypeOf((*ExitStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_ExitStatus{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExitStatus)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ExitStatusConfig",
		reflect.TypeOf((*ExitStatusConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ExitStatusConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ExitStatusParameters",
		reflect.TypeOf((*ExitStatusParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "errorExitValue", GoGetter: "ErrorExitValue"},
			_jsii_.MemberProperty{JsiiProperty: "errorLevel", GoGetter: "ErrorLevel"},
			_jsii_.MemberProperty{JsiiProperty: "okExitValue", GoGetter: "OkExitValue"},
			_jsii_.MemberProperty{JsiiProperty: "warnExitValue", GoGetter: "WarnExitValue"},
			_jsii_.MemberProperty{JsiiProperty: "warnLevel", GoGetter: "WarnLevel"},
		},
		func() interface{} {
			j := jsiiProxy_ExitStatusParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExitStatusParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Export",
		reflect.TypeOf((*Export)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Export{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IExport)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ExportConfig",
		reflect.TypeOf((*ExportConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ExportConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.GraphQL",
		reflect.TypeOf((*GraphQL)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "argsType", GoGetter: "ArgsType"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "lenientMaxQueryDepthEnabled", GoGetter: "LenientMaxQueryDepthEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "maxAdditionalQueryDepth", GoGetter: "MaxAdditionalQueryDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxArgsDepth", GoGetter: "MaxArgsDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxQueryDepth", GoGetter: "MaxQueryDepth"},
			_jsii_.MemberProperty{JsiiProperty: "optionalArgsEnabled", GoGetter: "OptionalArgsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "queryGenEnabled", GoGetter: "QueryGenEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "querySplitType", GoGetter: "QuerySplitType"},
			_jsii_.MemberProperty{JsiiProperty: "requestMethod", GoGetter: "RequestMethod"},
			_jsii_.MemberProperty{JsiiProperty: "schemaFile", GoGetter: "SchemaFile"},
			_jsii_.MemberProperty{JsiiProperty: "schemaUrl", GoGetter: "SchemaUrl"},
		},
		func() interface{} {
			j := jsiiProxy_GraphQL{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGraphQL)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.GraphQLConfig",
		reflect.TypeOf((*GraphQLConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_GraphQLConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.HttpHeaders",
		reflect.TypeOf((*HttpHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allRequests", GoGetter: "AllRequests"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			j := jsiiProxy_HttpHeaders{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IHttpHeaders)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScan",
		reflect.TypeOf((*IActiveScan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "policyDefinition", GoGetter: "PolicyDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScan{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanConfig",
		reflect.TypeOf((*IActiveScanConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "excludePaths", GoGetter: "ExcludePaths"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanConfigParameters",
		reflect.TypeOf((*IActiveScanConfigParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "defaultPolicy", GoGetter: "DefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "handleAntiCSRFTokens", GoGetter: "HandleAntiCSRFTokens"},
			_jsii_.MemberProperty{JsiiProperty: "injectPluginIdInHeader", GoGetter: "InjectPluginIdInHeader"},
			_jsii_.MemberProperty{JsiiProperty: "inputVectors", GoGetter: "InputVectors"},
			_jsii_.MemberProperty{JsiiProperty: "maxAlertsPerRule", GoGetter: "MaxAlertsPerRule"},
			_jsii_.MemberProperty{JsiiProperty: "maxRuleDurationInMins", GoGetter: "MaxRuleDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "maxScanDurationInMins", GoGetter: "MaxScanDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "threadPerHost", GoGetter: "ThreadPerHost"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanConfigParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanConfigProps",
		reflect.TypeOf((*IActiveScanConfigProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeScanConfig", GoGetter: "ActiveScanConfig"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanConfigProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanJob",
		reflect.TypeOf((*IActiveScanJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeScan", GoGetter: "ActiveScan"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanJob{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanParameters",
		reflect.TypeOf((*IActiveScanParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addQueryParam", GoGetter: "AddQueryParam"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPolicy", GoGetter: "DefaultPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "delayInMs", GoGetter: "DelayInMs"},
			_jsii_.MemberProperty{JsiiProperty: "handleAntiCSRFTokens", GoGetter: "HandleAntiCSRFTokens"},
			_jsii_.MemberProperty{JsiiProperty: "injectPluginIdInHeader", GoGetter: "InjectPluginIdInHeader"},
			_jsii_.MemberProperty{JsiiProperty: "maxAlertsPerRule", GoGetter: "MaxAlertsPerRule"},
			_jsii_.MemberProperty{JsiiProperty: "maxRuleDurationInMins", GoGetter: "MaxRuleDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "maxScanDurationInMins", GoGetter: "MaxScanDurationInMins"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "scanHeadersAllRequests", GoGetter: "ScanHeadersAllRequests"},
			_jsii_.MemberProperty{JsiiProperty: "tests", GoGetter: "Tests"},
			_jsii_.MemberProperty{JsiiProperty: "threadPerHost", GoGetter: "ThreadPerHost"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanPolicy",
		reflect.TypeOf((*IActiveScanPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanPolicy{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanPolicyDefinition",
		reflect.TypeOf((*IActiveScanPolicyDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanPolicyDefinition{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanPolicyParameters",
		reflect.TypeOf((*IActiveScanPolicyParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "policyDefinition", GoGetter: "PolicyDefinition"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanPolicyParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IActiveScanPolicyProps",
		reflect.TypeOf((*IActiveScanPolicyProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeScanPolicy", GoGetter: "ActiveScanPolicy"},
		},
		func() interface{} {
			return &jsiiProxy_IActiveScanPolicyProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAjaxTest",
		reflect.TypeOf((*IAjaxTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IAjaxTest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAlertFilter",
		reflect.TypeOf((*IAlertFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attack", GoGetter: "Attack"},
			_jsii_.MemberProperty{JsiiProperty: "attackRegex", GoGetter: "AttackRegex"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "evidence", GoGetter: "Evidence"},
			_jsii_.MemberProperty{JsiiProperty: "evidenceRegex", GoGetter: "EvidenceRegex"},
			_jsii_.MemberProperty{JsiiProperty: "newRisk", GoGetter: "NewRisk"},
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterRegex", GoGetter: "ParameterRegex"},
			_jsii_.MemberProperty{JsiiProperty: "ruleId", GoGetter: "RuleId"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlRegex", GoGetter: "UrlRegex"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertFilter{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAlertFilterParameters",
		reflect.TypeOf((*IAlertFilterParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertFilters", GoGetter: "AlertFilters"},
			_jsii_.MemberProperty{JsiiProperty: "deleteGlobalAlerts", GoGetter: "DeleteGlobalAlerts"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertFilterParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAlertTag",
		reflect.TypeOf((*IAlertTag)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertTag{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAlertTags",
		reflect.TypeOf((*IAlertTags)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertTags{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAlertTest",
		reflect.TypeOf((*IAlertTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "action", GoGetter: "Action"},
			_jsii_.MemberProperty{JsiiProperty: "alertName", GoGetter: "AlertName"},
			_jsii_.MemberProperty{JsiiProperty: "attack", GoGetter: "Attack"},
			_jsii_.MemberProperty{JsiiProperty: "confidence", GoGetter: "Confidence"},
			_jsii_.MemberProperty{JsiiProperty: "evidence", GoGetter: "Evidence"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "param", GoGetter: "Param"},
			_jsii_.MemberProperty{JsiiProperty: "risk", GoGetter: "Risk"},
			_jsii_.MemberProperty{JsiiProperty: "scanRuleId", GoGetter: "ScanRuleId"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			return &jsiiProxy_IAlertTest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAuthenticationParameters",
		reflect.TypeOf((*IAuthenticationParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "verification", GoGetter: "Verification"},
		},
		func() interface{} {
			return &jsiiProxy_IAuthenticationParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAuthenticationParametersParameters",
		reflect.TypeOf((*IAuthenticationParametersParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "loginPageUrl", GoGetter: "LoginPageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "loginRequestBody", GoGetter: "LoginRequestBody"},
			_jsii_.MemberProperty{JsiiProperty: "loginRequestUrl", GoGetter: "LoginRequestUrl"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "realm", GoGetter: "Realm"},
			_jsii_.MemberProperty{JsiiProperty: "script", GoGetter: "Script"},
			_jsii_.MemberProperty{JsiiProperty: "scriptEngine", GoGetter: "ScriptEngine"},
			_jsii_.MemberProperty{JsiiProperty: "scriptInline", GoGetter: "ScriptInline"},
		},
		func() interface{} {
			return &jsiiProxy_IAuthenticationParametersParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IAuthenticationParametersVerification",
		reflect.TypeOf((*IAuthenticationParametersVerification)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "loggedInRegex", GoGetter: "LoggedInRegex"},
			_jsii_.MemberProperty{JsiiProperty: "loggedOutRegex", GoGetter: "LoggedOutRegex"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "pollAdditionalHeaders", GoGetter: "PollAdditionalHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "pollFrequency", GoGetter: "PollFrequency"},
			_jsii_.MemberProperty{JsiiProperty: "pollPostData", GoGetter: "PollPostData"},
			_jsii_.MemberProperty{JsiiProperty: "pollUnits", GoGetter: "PollUnits"},
			_jsii_.MemberProperty{JsiiProperty: "pollUrl", GoGetter: "PollUrl"},
		},
		func() interface{} {
			return &jsiiProxy_IAuthenticationParametersVerification{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IContext",
		reflect.TypeOf((*IContext)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "excludePaths", GoGetter: "ExcludePaths"},
			_jsii_.MemberProperty{JsiiProperty: "includePaths", GoGetter: "IncludePaths"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "sessionManagement", GoGetter: "SessionManagement"},
			_jsii_.MemberProperty{JsiiProperty: "structure", GoGetter: "Structure"},
			_jsii_.MemberProperty{JsiiProperty: "technology", GoGetter: "Technology"},
			_jsii_.MemberProperty{JsiiProperty: "urls", GoGetter: "Urls"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
		},
		func() interface{} {
			return &jsiiProxy_IContext{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IContextStructure",
		reflect.TypeOf((*IContextStructure)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "dataDrivenNodes", GoGetter: "DataDrivenNodes"},
			_jsii_.MemberProperty{JsiiProperty: "structuralParameters", GoGetter: "StructuralParameters"},
		},
		func() interface{} {
			return &jsiiProxy_IContextStructure{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IContextUser",
		reflect.TypeOf((*IContextUser)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credentials", GoGetter: "Credentials"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_IContextUser{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ICookieData",
		reflect.TypeOf((*ICookieData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "encodeCookieValues", GoGetter: "EncodeCookieValues"},
		},
		func() interface{} {
			return &jsiiProxy_ICookieData{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IDataDrivenNode",
		reflect.TypeOf((*IDataDrivenNode)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
		},
		func() interface{} {
			return &jsiiProxy_IDataDrivenNode{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IDelay",
		reflect.TypeOf((*IDelay)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IDelay{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IDelayParameters",
		reflect.TypeOf((*IDelayParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "time", GoGetter: "Time"},
		},
		func() interface{} {
			return &jsiiProxy_IDelayParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IDelayProps",
		reflect.TypeOf((*IDelayProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "delay", GoGetter: "Delay"},
		},
		func() interface{} {
			return &jsiiProxy_IDelayProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IEnvironment",
		reflect.TypeOf((*IEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "contexts", GoGetter: "Contexts"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "proxy", GoGetter: "Proxy"},
			_jsii_.MemberProperty{JsiiProperty: "vars", GoGetter: "Vars"},
		},
		func() interface{} {
			return &jsiiProxy_IEnvironment{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IEnvironmentParameters",
		reflect.TypeOf((*IEnvironmentParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailure", GoGetter: "ContinueOnFailure"},
			_jsii_.MemberProperty{JsiiProperty: "failOnError", GoGetter: "FailOnError"},
			_jsii_.MemberProperty{JsiiProperty: "failOnWarning", GoGetter: "FailOnWarning"},
			_jsii_.MemberProperty{JsiiProperty: "progressToStdout", GoGetter: "ProgressToStdout"},
		},
		func() interface{} {
			return &jsiiProxy_IEnvironmentParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IEnvironmentProps",
		reflect.TypeOf((*IEnvironmentProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
		},
		func() interface{} {
			return &jsiiProxy_IEnvironmentProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IEnvironmentProxy",
		reflect.TypeOf((*IEnvironmentProxy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "realm", GoGetter: "Realm"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			return &jsiiProxy_IEnvironmentProxy{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IExcludedElement",
		reflect.TypeOf((*IExcludedElement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeName", GoGetter: "AttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "element", GoGetter: "Element"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "xpath", GoGetter: "Xpath"},
		},
		func() interface{} {
			return &jsiiProxy_IExcludedElement{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IExitStatus",
		reflect.TypeOf((*IExitStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IExitStatus{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IExitStatusParameters",
		reflect.TypeOf((*IExitStatusParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "errorExitValue", GoGetter: "ErrorExitValue"},
			_jsii_.MemberProperty{JsiiProperty: "errorLevel", GoGetter: "ErrorLevel"},
			_jsii_.MemberProperty{JsiiProperty: "okExitValue", GoGetter: "OkExitValue"},
			_jsii_.MemberProperty{JsiiProperty: "warnExitValue", GoGetter: "WarnExitValue"},
			_jsii_.MemberProperty{JsiiProperty: "warnLevel", GoGetter: "WarnLevel"},
		},
		func() interface{} {
			return &jsiiProxy_IExitStatusParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IExitStatusProps",
		reflect.TypeOf((*IExitStatusProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exitStatus", GoGetter: "ExitStatus"},
		},
		func() interface{} {
			return &jsiiProxy_IExitStatusProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IExport",
		reflect.TypeOf((*IExport)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IExport{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IExportProps",
		reflect.TypeOf((*IExportProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "export", GoGetter: "Export"},
		},
		func() interface{} {
			return &jsiiProxy_IExportProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IGraphQL",
		reflect.TypeOf((*IGraphQL)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "argsType", GoGetter: "ArgsType"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "lenientMaxQueryDepthEnabled", GoGetter: "LenientMaxQueryDepthEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "maxAdditionalQueryDepth", GoGetter: "MaxAdditionalQueryDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxArgsDepth", GoGetter: "MaxArgsDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxQueryDepth", GoGetter: "MaxQueryDepth"},
			_jsii_.MemberProperty{JsiiProperty: "optionalArgsEnabled", GoGetter: "OptionalArgsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "queryGenEnabled", GoGetter: "QueryGenEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "querySplitType", GoGetter: "QuerySplitType"},
			_jsii_.MemberProperty{JsiiProperty: "requestMethod", GoGetter: "RequestMethod"},
			_jsii_.MemberProperty{JsiiProperty: "schemaFile", GoGetter: "SchemaFile"},
			_jsii_.MemberProperty{JsiiProperty: "schemaUrl", GoGetter: "SchemaUrl"},
		},
		func() interface{} {
			return &jsiiProxy_IGraphQL{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IGraphQLProps",
		reflect.TypeOf((*IGraphQLProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "graphql", GoGetter: "Graphql"},
		},
		func() interface{} {
			return &jsiiProxy_IGraphQLProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IHttpHeaders",
		reflect.TypeOf((*IHttpHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allRequests", GoGetter: "AllRequests"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
		},
		func() interface{} {
			return &jsiiProxy_IHttpHeaders{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IImport",
		reflect.TypeOf((*IImport)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IImport{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IImportProps",
		reflect.TypeOf((*IImportProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "import", GoGetter: "Import"},
		},
		func() interface{} {
			return &jsiiProxy_IImportProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IInputVectors",
		reflect.TypeOf((*IInputVectors)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookieData", GoGetter: "CookieData"},
			_jsii_.MemberProperty{JsiiProperty: "httpHeaders", GoGetter: "HttpHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "postData", GoGetter: "PostData"},
			_jsii_.MemberProperty{JsiiProperty: "scripts", GoGetter: "Scripts"},
			_jsii_.MemberProperty{JsiiProperty: "urlPath", GoGetter: "UrlPath"},
			_jsii_.MemberProperty{JsiiProperty: "urlQueryStringAndDataDrivenNodes", GoGetter: "UrlQueryStringAndDataDrivenNodes"},
		},
		func() interface{} {
			return &jsiiProxy_IInputVectors{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IJsonPostData",
		reflect.TypeOf((*IJsonPostData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "scanNullValues", GoGetter: "ScanNullValues"},
		},
		func() interface{} {
			return &jsiiProxy_IJsonPostData{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IMonitorTest",
		reflect.TypeOf((*IMonitorTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "site", GoGetter: "Site"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IMonitorTest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.INewType",
		reflect.TypeOf((*INewType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
		},
		func() interface{} {
			return &jsiiProxy_INewType{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IOpenAPI",
		reflect.TypeOf((*IOpenAPI)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiFile", GoGetter: "ApiFile"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrl", GoGetter: "ApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "targetUrl", GoGetter: "TargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			return &jsiiProxy_IOpenAPI{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IOpenAPIProps",
		reflect.TypeOf((*IOpenAPIProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "openapi", GoGetter: "Openapi"},
		},
		func() interface{} {
			return &jsiiProxy_IOpenAPIProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPassiveScanConfig",
		reflect.TypeOf((*IPassiveScanConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IPassiveScanConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPassiveScanConfigProps",
		reflect.TypeOf((*IPassiveScanConfigProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "passiveScanConfig", GoGetter: "PassiveScanConfig"},
		},
		func() interface{} {
			return &jsiiProxy_IPassiveScanConfigProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPassiveScanParameters",
		reflect.TypeOf((*IPassiveScanParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "disableAllRules", GoGetter: "DisableAllRules"},
			_jsii_.MemberProperty{JsiiProperty: "enableTags", GoGetter: "EnableTags"},
			_jsii_.MemberProperty{JsiiProperty: "maxAlertsPerRule", GoGetter: "MaxAlertsPerRule"},
			_jsii_.MemberProperty{JsiiProperty: "maxBodySizeInBytesToScan", GoGetter: "MaxBodySizeInBytesToScan"},
			_jsii_.MemberProperty{JsiiProperty: "scanOnlyInScope", GoGetter: "ScanOnlyInScope"},
		},
		func() interface{} {
			return &jsiiProxy_IPassiveScanParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPassiveScanRule",
		reflect.TypeOf((*IPassiveScanRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			return &jsiiProxy_IPassiveScanRule{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPassiveScanWait",
		reflect.TypeOf((*IPassiveScanWait)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "maxDuration", GoGetter: "MaxDuration"},
		},
		func() interface{} {
			return &jsiiProxy_IPassiveScanWait{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPassiveScanWaitProps",
		reflect.TypeOf((*IPassiveScanWaitProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "passiveScanWait", GoGetter: "PassiveScanWait"},
		},
		func() interface{} {
			return &jsiiProxy_IPassiveScanWaitProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPolicyDefinition",
		reflect.TypeOf((*IPolicyDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertTags", GoGetter: "AlertTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultStrength", GoGetter: "DefaultStrength"},
			_jsii_.MemberProperty{JsiiProperty: "defaultThreshold", GoGetter: "DefaultThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			return &jsiiProxy_IPolicyDefinition{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPollAdditionalHeaders",
		reflect.TypeOf((*IPollAdditionalHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IPollAdditionalHeaders{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPostData",
		reflect.TypeOf((*IPostData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "directWebRemoting", GoGetter: "DirectWebRemoting"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "googleWebToolkit", GoGetter: "GoogleWebToolkit"},
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberProperty{JsiiProperty: "multiPartFormData", GoGetter: "MultiPartFormData"},
			_jsii_.MemberProperty{JsiiProperty: "xml", GoGetter: "Xml"},
		},
		func() interface{} {
			return &jsiiProxy_IPostData{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPostman",
		reflect.TypeOf((*IPostman)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collectionFile", GoGetter: "CollectionFile"},
			_jsii_.MemberProperty{JsiiProperty: "collectionUrl", GoGetter: "CollectionUrl"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			return &jsiiProxy_IPostman{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IPostmanProps",
		reflect.TypeOf((*IPostmanProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "postman", GoGetter: "Postman"},
		},
		func() interface{} {
			return &jsiiProxy_IPostmanProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IReplacer",
		reflect.TypeOf((*IReplacer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deleteAllRules", GoGetter: "DeleteAllRules"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			return &jsiiProxy_IReplacer{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IReplacerProps",
		reflect.TypeOf((*IReplacerProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "replacer", GoGetter: "Replacer"},
		},
		func() interface{} {
			return &jsiiProxy_IReplacerProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IReplacerRule",
		reflect.TypeOf((*IReplacerRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "initiators", GoGetter: "Initiators"},
			_jsii_.MemberProperty{JsiiProperty: "matchRegex", GoGetter: "MatchRegex"},
			_jsii_.MemberProperty{JsiiProperty: "matchString", GoGetter: "MatchString"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "replacementString", GoGetter: "ReplacementString"},
			_jsii_.MemberProperty{JsiiProperty: "tokenProcessing", GoGetter: "TokenProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			return &jsiiProxy_IReplacerRule{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IReport",
		reflect.TypeOf((*IReport)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "confidences", GoGetter: "Confidences"},
			_jsii_.MemberProperty{JsiiProperty: "displayReport", GoGetter: "DisplayReport"},
			_jsii_.MemberProperty{JsiiProperty: "reportDescription", GoGetter: "ReportDescription"},
			_jsii_.MemberProperty{JsiiProperty: "reportDir", GoGetter: "ReportDir"},
			_jsii_.MemberProperty{JsiiProperty: "reportFile", GoGetter: "ReportFile"},
			_jsii_.MemberProperty{JsiiProperty: "reportTitle", GoGetter: "ReportTitle"},
			_jsii_.MemberProperty{JsiiProperty: "risks", GoGetter: "Risks"},
			_jsii_.MemberProperty{JsiiProperty: "sections", GoGetter: "Sections"},
			_jsii_.MemberProperty{JsiiProperty: "sites", GoGetter: "Sites"},
			_jsii_.MemberProperty{JsiiProperty: "template", GoGetter: "Template"},
			_jsii_.MemberProperty{JsiiProperty: "theme", GoGetter: "Theme"},
		},
		func() interface{} {
			return &jsiiProxy_IReport{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IReportProps",
		reflect.TypeOf((*IReportProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "report", GoGetter: "Report"},
		},
		func() interface{} {
			return &jsiiProxy_IReportProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IRequest",
		reflect.TypeOf((*IRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersion", GoGetter: "HttpVersion"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "responseCode", GoGetter: "ResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			return &jsiiProxy_IRequest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IRequestorParameters",
		reflect.TypeOf((*IRequestorParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "requests", GoGetter: "Requests"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			return &jsiiProxy_IRequestorParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IRequestorProps",
		reflect.TypeOf((*IRequestorProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "requestor", GoGetter: "Requestor"},
		},
		func() interface{} {
			return &jsiiProxy_IRequestorProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IRule",
		reflect.TypeOf((*IRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			return &jsiiProxy_IRule{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IRules",
		reflect.TypeOf((*IRules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			return &jsiiProxy_IRules{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISessionManagementParameters",
		reflect.TypeOf((*ISessionManagementParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
		},
		func() interface{} {
			return &jsiiProxy_ISessionManagementParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISessionManagementParametersParameters",
		reflect.TypeOf((*ISessionManagementParametersParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "script", GoGetter: "Script"},
			_jsii_.MemberProperty{JsiiProperty: "scriptEngine", GoGetter: "ScriptEngine"},
		},
		func() interface{} {
			return &jsiiProxy_ISessionManagementParametersParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISoap",
		reflect.TypeOf((*ISoap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "wsdlFile", GoGetter: "WsdlFile"},
			_jsii_.MemberProperty{JsiiProperty: "wsdlUrl", GoGetter: "WsdlUrl"},
		},
		func() interface{} {
			return &jsiiProxy_ISoap{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISoapProps",
		reflect.TypeOf((*ISoapProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "soap", GoGetter: "Soap"},
		},
		func() interface{} {
			return &jsiiProxy_ISoapProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISpider",
		reflect.TypeOf((*ISpider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "tests", GoGetter: "Tests"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_ISpider{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISpiderAjax",
		reflect.TypeOf((*ISpiderAjax)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "browserId", GoGetter: "BrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "clickDefaultElems", GoGetter: "ClickDefaultElems"},
			_jsii_.MemberProperty{JsiiProperty: "clickElemsOnce", GoGetter: "ClickElemsOnce"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "elements", GoGetter: "Elements"},
			_jsii_.MemberProperty{JsiiProperty: "eventWait", GoGetter: "EventWait"},
			_jsii_.MemberProperty{JsiiProperty: "excludedElements", GoGetter: "ExcludedElements"},
			_jsii_.MemberProperty{JsiiProperty: "inScopeOnly", GoGetter: "InScopeOnly"},
			_jsii_.MemberProperty{JsiiProperty: "maxCrawlDepth", GoGetter: "MaxCrawlDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxCrawlStates", GoGetter: "MaxCrawlStates"},
			_jsii_.MemberProperty{JsiiProperty: "maxDuration", GoGetter: "MaxDuration"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfBrowsers", GoGetter: "NumberOfBrowsers"},
			_jsii_.MemberProperty{JsiiProperty: "randomInputs", GoGetter: "RandomInputs"},
			_jsii_.MemberProperty{JsiiProperty: "reloadWait", GoGetter: "ReloadWait"},
			_jsii_.MemberProperty{JsiiProperty: "runOnlyIfModern", GoGetter: "RunOnlyIfModern"},
			_jsii_.MemberProperty{JsiiProperty: "scopeCheck", GoGetter: "ScopeCheck"},
			_jsii_.MemberProperty{JsiiProperty: "tests", GoGetter: "Tests"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			return &jsiiProxy_ISpiderAjax{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISpiderAjaxProps",
		reflect.TypeOf((*ISpiderAjaxProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "spiderAjax", GoGetter: "SpiderAjax"},
		},
		func() interface{} {
			return &jsiiProxy_ISpiderAjaxProps{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISpiderParameters",
		reflect.TypeOf((*ISpiderParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptCookies", GoGetter: "AcceptCookies"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "handleODataParametersVisited", GoGetter: "HandleODataParametersVisited"},
			_jsii_.MemberProperty{JsiiProperty: "handleParameters", GoGetter: "HandleParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logoutAvoidance", GoGetter: "LogoutAvoidance"},
			_jsii_.MemberProperty{JsiiProperty: "maxChildren", GoGetter: "MaxChildren"},
			_jsii_.MemberProperty{JsiiProperty: "maxDepth", GoGetter: "MaxDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxDuration", GoGetter: "MaxDuration"},
			_jsii_.MemberProperty{JsiiProperty: "maxParseSizeBytes", GoGetter: "MaxParseSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "parseComments", GoGetter: "ParseComments"},
			_jsii_.MemberProperty{JsiiProperty: "parseDsStore", GoGetter: "ParseDsStore"},
			_jsii_.MemberProperty{JsiiProperty: "parseGit", GoGetter: "ParseGit"},
			_jsii_.MemberProperty{JsiiProperty: "parseRobotsTxt", GoGetter: "ParseRobotsTxt"},
			_jsii_.MemberProperty{JsiiProperty: "parseSitemapXml", GoGetter: "ParseSitemapXml"},
			_jsii_.MemberProperty{JsiiProperty: "parseSVNEntries", GoGetter: "ParseSVNEntries"},
			_jsii_.MemberProperty{JsiiProperty: "postForm", GoGetter: "PostForm"},
			_jsii_.MemberProperty{JsiiProperty: "processForm", GoGetter: "ProcessForm"},
			_jsii_.MemberProperty{JsiiProperty: "sendRefererHeader", GoGetter: "SendRefererHeader"},
			_jsii_.MemberProperty{JsiiProperty: "threadCount", GoGetter: "ThreadCount"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userAgent", GoGetter: "UserAgent"},
		},
		func() interface{} {
			return &jsiiProxy_ISpiderParameters{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ISpiderTest",
		reflect.TypeOf((*ISpiderTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ISpiderTest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IStatisticsTest",
		reflect.TypeOf((*IStatisticsTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "site", GoGetter: "Site"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IStatisticsTest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ITechnology",
		reflect.TypeOf((*ITechnology)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
		},
		func() interface{} {
			return &jsiiProxy_ITechnology{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.ITotpConfig",
		reflect.TypeOf((*ITotpConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "algorithm", GoGetter: "Algorithm"},
			_jsii_.MemberProperty{JsiiProperty: "digits", GoGetter: "Digits"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
		},
		func() interface{} {
			return &jsiiProxy_ITotpConfig{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IUrlQueryStringAndDataDrivenNodes",
		reflect.TypeOf((*IUrlQueryStringAndDataDrivenNodes)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addParam", GoGetter: "AddParam"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "odata", GoGetter: "Odata"},
		},
		func() interface{} {
			return &jsiiProxy_IUrlQueryStringAndDataDrivenNodes{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IUrlTest",
		reflect.TypeOf((*IUrlTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "requestBodyRegex", GoGetter: "RequestBodyRegex"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeaderRegex", GoGetter: "RequestHeaderRegex"},
			_jsii_.MemberProperty{JsiiProperty: "responseBodyRegex", GoGetter: "ResponseBodyRegex"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeaderRegex", GoGetter: "ResponseHeaderRegex"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			return &jsiiProxy_IUrlTest{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IUserCredentials",
		reflect.TypeOf((*IUserCredentials)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "totp", GoGetter: "Totp"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			return &jsiiProxy_IUserCredentials{}
		},
	)
	_jsii_.RegisterInterface(
		"zap-cdk.IZap",
		reflect.TypeOf((*IZap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobs", GoGetter: "Jobs"},
		},
		func() interface{} {
			return &jsiiProxy_IZap{}
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Import",
		reflect.TypeOf((*Import)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fileName", GoGetter: "FileName"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Import{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImport)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ImportConfig",
		reflect.TypeOf((*ImportConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ImportConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.InputVectors",
		reflect.TypeOf((*InputVectors)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cookieData", GoGetter: "CookieData"},
			_jsii_.MemberProperty{JsiiProperty: "httpHeaders", GoGetter: "HttpHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "postData", GoGetter: "PostData"},
			_jsii_.MemberProperty{JsiiProperty: "scripts", GoGetter: "Scripts"},
			_jsii_.MemberProperty{JsiiProperty: "urlPath", GoGetter: "UrlPath"},
			_jsii_.MemberProperty{JsiiProperty: "urlQueryStringAndDataDrivenNodes", GoGetter: "UrlQueryStringAndDataDrivenNodes"},
		},
		func() interface{} {
			j := jsiiProxy_InputVectors{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInputVectors)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.JsonPostData",
		reflect.TypeOf((*JsonPostData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "scanNullValues", GoGetter: "ScanNullValues"},
		},
		func() interface{} {
			j := jsiiProxy_JsonPostData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJsonPostData)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.MonitorTest",
		reflect.TypeOf((*MonitorTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "site", GoGetter: "Site"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_MonitorTest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IMonitorTest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.OpenAPI",
		reflect.TypeOf((*OpenAPI)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiFile", GoGetter: "ApiFile"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrl", GoGetter: "ApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "targetUrl", GoGetter: "TargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			j := jsiiProxy_OpenAPI{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOpenAPI)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.OpenAPIConfig",
		reflect.TypeOf((*OpenAPIConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_OpenAPIConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PassiveScanConfig",
		reflect.TypeOf((*PassiveScanConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_PassiveScanConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPassiveScanConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PassiveScanConfigConstruct",
		reflect.TypeOf((*PassiveScanConfigConstruct)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_PassiveScanConfigConstruct{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PassiveScanParameters",
		reflect.TypeOf((*PassiveScanParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "disableAllRules", GoGetter: "DisableAllRules"},
			_jsii_.MemberProperty{JsiiProperty: "enableTags", GoGetter: "EnableTags"},
			_jsii_.MemberProperty{JsiiProperty: "maxAlertsPerRule", GoGetter: "MaxAlertsPerRule"},
			_jsii_.MemberProperty{JsiiProperty: "maxBodySizeInBytesToScan", GoGetter: "MaxBodySizeInBytesToScan"},
			_jsii_.MemberProperty{JsiiProperty: "scanOnlyInScope", GoGetter: "ScanOnlyInScope"},
		},
		func() interface{} {
			j := jsiiProxy_PassiveScanParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPassiveScanParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PassiveScanRule",
		reflect.TypeOf((*PassiveScanRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			j := jsiiProxy_PassiveScanRule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPassiveScanRule)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PassiveScanWait",
		reflect.TypeOf((*PassiveScanWait)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "maxDuration", GoGetter: "MaxDuration"},
		},
		func() interface{} {
			j := jsiiProxy_PassiveScanWait{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPassiveScanWait)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PassiveScanWaitConfig",
		reflect.TypeOf((*PassiveScanWaitConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_PassiveScanWaitConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PolicyDefinition",
		reflect.TypeOf((*PolicyDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertTags", GoGetter: "AlertTags"},
			_jsii_.MemberProperty{JsiiProperty: "defaultStrength", GoGetter: "DefaultStrength"},
			_jsii_.MemberProperty{JsiiProperty: "defaultThreshold", GoGetter: "DefaultThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			j := jsiiProxy_PolicyDefinition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPolicyDefinition)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PollAdditionalHeaders",
		reflect.TypeOf((*PollAdditionalHeaders)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "header", GoGetter: "Header"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_PollAdditionalHeaders{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPollAdditionalHeaders)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PostData",
		reflect.TypeOf((*PostData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "directWebRemoting", GoGetter: "DirectWebRemoting"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "googleWebToolkit", GoGetter: "GoogleWebToolkit"},
			_jsii_.MemberProperty{JsiiProperty: "json", GoGetter: "Json"},
			_jsii_.MemberProperty{JsiiProperty: "multiPartFormData", GoGetter: "MultiPartFormData"},
			_jsii_.MemberProperty{JsiiProperty: "xml", GoGetter: "Xml"},
		},
		func() interface{} {
			j := jsiiProxy_PostData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPostData)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Postman",
		reflect.TypeOf((*Postman)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collectionFile", GoGetter: "CollectionFile"},
			_jsii_.MemberProperty{JsiiProperty: "collectionUrl", GoGetter: "CollectionUrl"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_Postman{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPostman)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.PostmanConfig",
		reflect.TypeOf((*PostmanConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_PostmanConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Replacer",
		reflect.TypeOf((*Replacer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "deleteAllRules", GoGetter: "DeleteAllRules"},
			_jsii_.MemberProperty{JsiiProperty: "rules", GoGetter: "Rules"},
		},
		func() interface{} {
			j := jsiiProxy_Replacer{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IReplacer)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ReplacerConfig",
		reflect.TypeOf((*ReplacerConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ReplacerConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ReplacerRule",
		reflect.TypeOf((*ReplacerRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "initiators", GoGetter: "Initiators"},
			_jsii_.MemberProperty{JsiiProperty: "matchRegex", GoGetter: "MatchRegex"},
			_jsii_.MemberProperty{JsiiProperty: "matchString", GoGetter: "MatchString"},
			_jsii_.MemberProperty{JsiiProperty: "matchType", GoGetter: "MatchType"},
			_jsii_.MemberProperty{JsiiProperty: "replacementString", GoGetter: "ReplacementString"},
			_jsii_.MemberProperty{JsiiProperty: "tokenProcessing", GoGetter: "TokenProcessing"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_ReplacerRule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IReplacerRule)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Report",
		reflect.TypeOf((*Report)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "confidences", GoGetter: "Confidences"},
			_jsii_.MemberProperty{JsiiProperty: "displayReport", GoGetter: "DisplayReport"},
			_jsii_.MemberProperty{JsiiProperty: "reportDescription", GoGetter: "ReportDescription"},
			_jsii_.MemberProperty{JsiiProperty: "reportDir", GoGetter: "ReportDir"},
			_jsii_.MemberProperty{JsiiProperty: "reportFile", GoGetter: "ReportFile"},
			_jsii_.MemberProperty{JsiiProperty: "reportTitle", GoGetter: "ReportTitle"},
			_jsii_.MemberProperty{JsiiProperty: "risks", GoGetter: "Risks"},
			_jsii_.MemberProperty{JsiiProperty: "sections", GoGetter: "Sections"},
			_jsii_.MemberProperty{JsiiProperty: "sites", GoGetter: "Sites"},
			_jsii_.MemberProperty{JsiiProperty: "template", GoGetter: "Template"},
			_jsii_.MemberProperty{JsiiProperty: "theme", GoGetter: "Theme"},
		},
		func() interface{} {
			j := jsiiProxy_Report{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IReport)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ReportConfig",
		reflect.TypeOf((*ReportConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ReportConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Request",
		reflect.TypeOf((*Request)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "headers", GoGetter: "Headers"},
			_jsii_.MemberProperty{JsiiProperty: "httpVersion", GoGetter: "HttpVersion"},
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "responseCode", GoGetter: "ResponseCode"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_Request{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRequest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.RequestorConfig",
		reflect.TypeOf((*RequestorConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_RequestorConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.RequestorParameters",
		reflect.TypeOf((*RequestorParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "requests", GoGetter: "Requests"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			j := jsiiProxy_RequestorParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRequestorParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Rule",
		reflect.TypeOf((*Rule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			j := jsiiProxy_Rule{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRule)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Rules",
		reflect.TypeOf((*Rules)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "strength", GoGetter: "Strength"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
		},
		func() interface{} {
			j := jsiiProxy_Rules{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRules)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SOAPConfig",
		reflect.TypeOf((*SOAPConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_SOAPConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SessionManagementParameters",
		reflect.TypeOf((*SessionManagementParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "method", GoGetter: "Method"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
		},
		func() interface{} {
			j := jsiiProxy_SessionManagementParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISessionManagementParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SessionManagementParametersParameters",
		reflect.TypeOf((*SessionManagementParametersParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "script", GoGetter: "Script"},
			_jsii_.MemberProperty{JsiiProperty: "scriptEngine", GoGetter: "ScriptEngine"},
		},
		func() interface{} {
			j := jsiiProxy_SessionManagementParametersParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISessionManagementParametersParameters)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Soap",
		reflect.TypeOf((*Soap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "wsdlFile", GoGetter: "WsdlFile"},
			_jsii_.MemberProperty{JsiiProperty: "wsdlUrl", GoGetter: "WsdlUrl"},
		},
		func() interface{} {
			j := jsiiProxy_Soap{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISoap)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Spider",
		reflect.TypeOf((*Spider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alwaysRun", GoGetter: "AlwaysRun"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "tests", GoGetter: "Tests"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Spider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISpider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SpiderAjax",
		reflect.TypeOf((*SpiderAjax)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "browserId", GoGetter: "BrowserId"},
			_jsii_.MemberProperty{JsiiProperty: "clickDefaultElems", GoGetter: "ClickDefaultElems"},
			_jsii_.MemberProperty{JsiiProperty: "clickElemsOnce", GoGetter: "ClickElemsOnce"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "elements", GoGetter: "Elements"},
			_jsii_.MemberProperty{JsiiProperty: "eventWait", GoGetter: "EventWait"},
			_jsii_.MemberProperty{JsiiProperty: "excludedElements", GoGetter: "ExcludedElements"},
			_jsii_.MemberProperty{JsiiProperty: "inScopeOnly", GoGetter: "InScopeOnly"},
			_jsii_.MemberProperty{JsiiProperty: "maxCrawlDepth", GoGetter: "MaxCrawlDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxCrawlStates", GoGetter: "MaxCrawlStates"},
			_jsii_.MemberProperty{JsiiProperty: "maxDuration", GoGetter: "MaxDuration"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfBrowsers", GoGetter: "NumberOfBrowsers"},
			_jsii_.MemberProperty{JsiiProperty: "randomInputs", GoGetter: "RandomInputs"},
			_jsii_.MemberProperty{JsiiProperty: "reloadWait", GoGetter: "ReloadWait"},
			_jsii_.MemberProperty{JsiiProperty: "runOnlyIfModern", GoGetter: "RunOnlyIfModern"},
			_jsii_.MemberProperty{JsiiProperty: "scopeCheck", GoGetter: "ScopeCheck"},
			_jsii_.MemberProperty{JsiiProperty: "tests", GoGetter: "Tests"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
		},
		func() interface{} {
			j := jsiiProxy_SpiderAjax{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISpiderAjax)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SpiderAjaxConfig",
		reflect.TypeOf((*SpiderAjaxConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_SpiderAjaxConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SpiderConfig",
		reflect.TypeOf((*SpiderConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_SpiderConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.SpiderParameters",
		reflect.TypeOf((*SpiderParameters)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceptCookies", GoGetter: "AcceptCookies"},
			_jsii_.MemberProperty{JsiiProperty: "context", GoGetter: "Context"},
			_jsii_.MemberProperty{JsiiProperty: "handleODataParametersVisited", GoGetter: "HandleODataParametersVisited"},
			_jsii_.MemberProperty{JsiiProperty: "handleParameters", GoGetter: "HandleParameters"},
			_jsii_.MemberProperty{JsiiProperty: "logoutAvoidance", GoGetter: "LogoutAvoidance"},
			_jsii_.MemberProperty{JsiiProperty: "maxChildren", GoGetter: "MaxChildren"},
			_jsii_.MemberProperty{JsiiProperty: "maxDepth", GoGetter: "MaxDepth"},
			_jsii_.MemberProperty{JsiiProperty: "maxDuration", GoGetter: "MaxDuration"},
			_jsii_.MemberProperty{JsiiProperty: "maxParseSizeBytes", GoGetter: "MaxParseSizeBytes"},
			_jsii_.MemberProperty{JsiiProperty: "parseComments", GoGetter: "ParseComments"},
			_jsii_.MemberProperty{JsiiProperty: "parseDsStore", GoGetter: "ParseDsStore"},
			_jsii_.MemberProperty{JsiiProperty: "parseGit", GoGetter: "ParseGit"},
			_jsii_.MemberProperty{JsiiProperty: "parseRobotsTxt", GoGetter: "ParseRobotsTxt"},
			_jsii_.MemberProperty{JsiiProperty: "parseSitemapXml", GoGetter: "ParseSitemapXml"},
			_jsii_.MemberProperty{JsiiProperty: "parseSVNEntries", GoGetter: "ParseSVNEntries"},
			_jsii_.MemberProperty{JsiiProperty: "postForm", GoGetter: "PostForm"},
			_jsii_.MemberProperty{JsiiProperty: "processForm", GoGetter: "ProcessForm"},
			_jsii_.MemberProperty{JsiiProperty: "sendRefererHeader", GoGetter: "SendRefererHeader"},
			_jsii_.MemberProperty{JsiiProperty: "threadCount", GoGetter: "ThreadCount"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userAgent", GoGetter: "UserAgent"},
		},
		func() interface{} {
			j := jsiiProxy_SpiderParameters{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISpiderParameters)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"zap-cdk.SpiderProps",
		reflect.TypeOf((*SpiderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"zap-cdk.SpiderTest",
		reflect.TypeOf((*SpiderTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_SpiderTest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISpiderTest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.StatisticsTest",
		reflect.TypeOf((*StatisticsTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "site", GoGetter: "Site"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_StatisticsTest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IStatisticsTest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Technology",
		reflect.TypeOf((*Technology)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
		},
		func() interface{} {
			j := jsiiProxy_Technology{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITechnology)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.TotpConfig",
		reflect.TypeOf((*TotpConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "algorithm", GoGetter: "Algorithm"},
			_jsii_.MemberProperty{JsiiProperty: "digits", GoGetter: "Digits"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
		},
		func() interface{} {
			j := jsiiProxy_TotpConfig{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ITotpConfig)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.UrlQueryStringAndDataDrivenNodes",
		reflect.TypeOf((*UrlQueryStringAndDataDrivenNodes)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addParam", GoGetter: "AddParam"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "odata", GoGetter: "Odata"},
		},
		func() interface{} {
			j := jsiiProxy_UrlQueryStringAndDataDrivenNodes{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUrlQueryStringAndDataDrivenNodes)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.UrlTest",
		reflect.TypeOf((*UrlTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "onFail", GoGetter: "OnFail"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "requestBodyRegex", GoGetter: "RequestBodyRegex"},
			_jsii_.MemberProperty{JsiiProperty: "requestHeaderRegex", GoGetter: "RequestHeaderRegex"},
			_jsii_.MemberProperty{JsiiProperty: "responseBodyRegex", GoGetter: "ResponseBodyRegex"},
			_jsii_.MemberProperty{JsiiProperty: "responseHeaderRegex", GoGetter: "ResponseHeaderRegex"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_UrlTest{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUrlTest)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.UserCredentials",
		reflect.TypeOf((*UserCredentials)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "totp", GoGetter: "Totp"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
		},
		func() interface{} {
			j := jsiiProxy_UserCredentials{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IUserCredentials)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.Zap",
		reflect.TypeOf((*Zap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "jobs", GoGetter: "Jobs"},
		},
		func() interface{} {
			j := jsiiProxy_Zap{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IZap)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"zap-cdk.ZapConfig",
		reflect.TypeOf((*ZapConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "synth", GoMethod: "Synth"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toYaml", GoMethod: "ToYaml"},
		},
		func() interface{} {
			j := jsiiProxy_ZapConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
}
