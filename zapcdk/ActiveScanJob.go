package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// Class representing an active scan job.
type ActiveScanJob interface {
	constructs.Construct
	// The active scan job properties.
	Job() IActiveScanJob
	SetJob(val IActiveScanJob)
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
	// Converts the active scan job to YAML format.
	//
	// Returns: The active scan job in YAML format.
	ToYaml() IActiveScanJob
}

// The jsii proxy struct for ActiveScanJob
type jsiiProxy_ActiveScanJob struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ActiveScanJob) Job() IActiveScanJob {
	var returns IActiveScanJob
	_jsii_.Get(
		j,
		"job",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanJob.
func NewActiveScanJob(scope constructs.Construct, id *string, props IActiveScanJob) ActiveScanJob {
	_init_.Initialize()

	if err := validateNewActiveScanJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanJob{}

	_jsii_.Create(
		"zap-cdk.ActiveScanJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanJob.
func NewActiveScanJob_Override(a ActiveScanJob, scope constructs.Construct, id *string, props IActiveScanJob) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanJob",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ActiveScanJob)SetJob(val IActiveScanJob) {
	if err := j.validateSetJobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"job",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ActiveScanJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActiveScanJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.ActiveScanJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveScanJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveScanJob) ToYaml() IActiveScanJob {
	var returns IActiveScanJob

	_jsii_.Invoke(
		a,
		"toYaml",
		nil, // no parameters
		&returns,
	)

	return returns
}

