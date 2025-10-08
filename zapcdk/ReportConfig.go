package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing the report configuration.
type ReportConfig interface {
	constructs.Construct
	Config() IReport
	SetConfig(val IReport)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the report configuration to YAML format.
	//
	// Returns: The report configuration in YAML format.
	ToYaml() IReport
}

// The jsii proxy struct for ReportConfig
type jsiiProxy_ReportConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ReportConfig) Config() IReport {
	var returns IReport
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ReportConfig.
func NewReportConfig(scope constructs.Construct, id *string, props IReportProps) ReportConfig {
	_init_.Initialize()

	if err := validateNewReportConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReportConfig{}

	_jsii_.Create(
		"zap-cdk.ReportConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ReportConfig.
func NewReportConfig_Override(r ReportConfig, scope constructs.Construct, id *string, props IReportProps) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ReportConfig",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_ReportConfig)SetConfig(val IReport) {
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
func ReportConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateReportConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ReportConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReportConfig) ToYaml() IReport {
	var returns IReport

	_jsii_.Invoke(
		r,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

