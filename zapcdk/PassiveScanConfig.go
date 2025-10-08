package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the passive scan configuration.
type PassiveScanConfig interface {
	constructs.Construct
	Config() IPassiveScanConfig
	SetConfig(val IPassiveScanConfig)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the passive scan configuration to YAML format.
	//
	// Returns: The passive scan configuration in YAML format.
	ToYaml() IPassiveScanConfig
}

// The jsii proxy struct for PassiveScanConfig
type jsiiProxy_PassiveScanConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PassiveScanConfig) Config() IPassiveScanConfig {
	var returns IPassiveScanConfig
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of PassiveScanConfig.
func NewPassiveScanConfig(scope constructs.Construct, id *string, props IPassiveScanConfigProps) PassiveScanConfig {
	_init_.Initialize()

	if err := validateNewPassiveScanConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PassiveScanConfig{}

	_jsii_.Create(
		"zap-cdk.PassiveScanConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of PassiveScanConfig.
func NewPassiveScanConfig_Override(p PassiveScanConfig, scope constructs.Construct, id *string, props IPassiveScanConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PassiveScanConfig",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_PassiveScanConfig)SetConfig(val IPassiveScanConfig) {
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
func PassiveScanConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePassiveScanConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.PassiveScanConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PassiveScanConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PassiveScanConfig) ToYaml() IPassiveScanConfig {
	var returns IPassiveScanConfig

	_jsii_.Invoke(
		p,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

