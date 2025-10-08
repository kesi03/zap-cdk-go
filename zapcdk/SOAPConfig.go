package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the SOAP configuration.
type SOAPConfig interface {
	constructs.Construct
	Config() ISoap
	SetConfig(val ISoap)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the SOAP configuration to YAML format.
	//
	// Returns: The SOAP configuration in YAML format.
	ToYaml() ISoap
}

// The jsii proxy struct for SOAPConfig
type jsiiProxy_SOAPConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SOAPConfig) Config() ISoap {
	var returns ISoap
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SOAPConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of SOAPConfig.
func NewSOAPConfig(scope constructs.Construct, id *string, props ISoapProps) SOAPConfig {
	_init_.Initialize()

	if err := validateNewSOAPConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SOAPConfig{}

	_jsii_.Create(
		"zap-cdk.SOAPConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of SOAPConfig.
func NewSOAPConfig_Override(s SOAPConfig, scope constructs.Construct, id *string, props ISoapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SOAPConfig",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SOAPConfig)SetConfig(val ISoap) {
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
func SOAPConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSOAPConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.SOAPConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SOAPConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SOAPConfig) ToYaml() ISoap {
	var returns ISoap

	_jsii_.Invoke(
		s,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

