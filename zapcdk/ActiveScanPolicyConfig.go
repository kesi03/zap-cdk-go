package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the active scan policy configuration.
type ActiveScanPolicyConfig interface {
	constructs.Construct
	Config() IActiveScanPolicy
	SetConfig(val IActiveScanPolicy)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the active scan policy configuration to YAML format.
	//
	// Returns: The active scan policy configuration in YAML format.
	ToYaml() IActiveScanPolicy
}

// The jsii proxy struct for ActiveScanPolicyConfig
type jsiiProxy_ActiveScanPolicyConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ActiveScanPolicyConfig) Config() IActiveScanPolicy {
	var returns IActiveScanPolicy
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicyConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanPolicyConfig.
func NewActiveScanPolicyConfig(scope constructs.Construct, id *string, props IActiveScanPolicyProps) ActiveScanPolicyConfig {
	_init_.Initialize()

	if err := validateNewActiveScanPolicyConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanPolicyConfig{}

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicyConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanPolicyConfig.
func NewActiveScanPolicyConfig_Override(a ActiveScanPolicyConfig, scope constructs.Construct, id *string, props IActiveScanPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicyConfig",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ActiveScanPolicyConfig)SetConfig(val IActiveScanPolicy) {
	if err := j.validateSetConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ActiveScanPolicyConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveScanPolicyConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ActiveScanPolicyConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveScanPolicyConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveScanPolicyConfig) ToYaml() IActiveScanPolicy {
	var returns IActiveScanPolicy

	_jsii_.Invoke(
		a,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

