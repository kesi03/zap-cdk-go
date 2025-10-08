package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the passive scan wait configuration.
type PassiveScanWaitConfig interface {
	constructs.Construct
	Config() IPassiveScanWait
	SetConfig(val IPassiveScanWait)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the passive scan wait configuration to YAML format.
	//
	// Returns: The passive scan wait configuration in YAML format.
	ToYaml() IPassiveScanWait
}

// The jsii proxy struct for PassiveScanWaitConfig
type jsiiProxy_PassiveScanWaitConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PassiveScanWaitConfig) Config() IPassiveScanWait {
	var returns IPassiveScanWait
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanWaitConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of PassiveScanWaitConfig.
func NewPassiveScanWaitConfig(scope constructs.Construct, id *string, props IPassiveScanWaitProps) PassiveScanWaitConfig {
	_init_.Initialize()

	if err := validateNewPassiveScanWaitConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PassiveScanWaitConfig{}

	_jsii_.Create(
		"zap-cdk.PassiveScanWaitConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of PassiveScanWaitConfig.
func NewPassiveScanWaitConfig_Override(p PassiveScanWaitConfig, scope constructs.Construct, id *string, props IPassiveScanWaitProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PassiveScanWaitConfig",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_PassiveScanWaitConfig)SetConfig(val IPassiveScanWait) {
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
func PassiveScanWaitConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePassiveScanWaitConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.PassiveScanWaitConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PassiveScanWaitConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PassiveScanWaitConfig) ToYaml() IPassiveScanWait {
	var returns IPassiveScanWait

	_jsii_.Invoke(
		p,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

