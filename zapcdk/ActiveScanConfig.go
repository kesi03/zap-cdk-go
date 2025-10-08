package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the active scan configuration.
type ActiveScanConfig interface {
	constructs.Construct
	// The active scan configuration properties.
	Config() IActiveScanConfig
	SetConfig(val IActiveScanConfig)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the active scan configuration to YAML format.
	//
	// Returns: The active scan configuration in YAML format.
	ToYaml() IActiveScanConfig
}

// The jsii proxy struct for ActiveScanConfig
type jsiiProxy_ActiveScanConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ActiveScanConfig) Config() IActiveScanConfig {
	var returns IActiveScanConfig
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanConfig.
func NewActiveScanConfig(scope constructs.Construct, id *string, props IActiveScanConfigProps) ActiveScanConfig {
	_init_.Initialize()

	if err := validateNewActiveScanConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanConfig{}

	_jsii_.Create(
		"zap-cdk.ActiveScanConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanConfig.
func NewActiveScanConfig_Override(a ActiveScanConfig, scope constructs.Construct, id *string, props IActiveScanConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanConfig",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ActiveScanConfig)SetConfig(val IActiveScanConfig) {
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
func ActiveScanConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveScanConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ActiveScanConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveScanConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveScanConfig) ToYaml() IActiveScanConfig {
	var returns IActiveScanConfig

	_jsii_.Invoke(
		a,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

