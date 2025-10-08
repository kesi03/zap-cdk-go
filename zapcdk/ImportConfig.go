package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the import configuration.
type ImportConfig interface {
	constructs.Construct
	Config() IImport
	SetConfig(val IImport)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the import configuration to YAML format.
	//
	// Returns: The import configuration in YAML format.
	ToYaml() IImport
}

// The jsii proxy struct for ImportConfig
type jsiiProxy_ImportConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ImportConfig) Config() IImport {
	var returns IImport
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImportConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ImportConfig.
func NewImportConfig(scope constructs.Construct, id *string, props IImportProps) ImportConfig {
	_init_.Initialize()

	if err := validateNewImportConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImportConfig{}

	_jsii_.Create(
		"zap-cdk.ImportConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ImportConfig.
func NewImportConfig_Override(i ImportConfig, scope constructs.Construct, id *string, props IImportProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ImportConfig",
		[]interface{}{scope, id, props},
		i,
	)
}

func (j *jsiiProxy_ImportConfig)SetConfig(val IImport) {
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
func ImportConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImportConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ImportConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImportConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImportConfig) ToYaml() IImport {
	var returns IImport

	_jsii_.Invoke(
		i,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

