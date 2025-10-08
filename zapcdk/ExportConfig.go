package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the export configuration.
type ExportConfig interface {
	constructs.Construct
	Config() IExport
	SetConfig(val IExport)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the export configuration to YAML format.
	//
	// Returns: The export configuration in YAML format.
	ToYaml() IExport
}

// The jsii proxy struct for ExportConfig
type jsiiProxy_ExportConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ExportConfig) Config() IExport {
	var returns IExport
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExportConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ExportConfig.
func NewExportConfig(scope constructs.Construct, id *string, props IExportProps) ExportConfig {
	_init_.Initialize()

	if err := validateNewExportConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExportConfig{}

	_jsii_.Create(
		"zap-cdk.ExportConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ExportConfig.
func NewExportConfig_Override(e ExportConfig, scope constructs.Construct, id *string, props IExportProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ExportConfig",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_ExportConfig)SetConfig(val IExport) {
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
func ExportConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExportConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ExportConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExportConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExportConfig) ToYaml() IExport {
	var returns IExport

	_jsii_.Invoke(
		e,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

