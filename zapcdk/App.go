package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/kesi03/zap-cdk-go/zapcdk/internal"
)

// The main application construct that aggregates all child constructs and synthesizes them into a single YAML file.
type App interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Synthesizes all child constructs into a single YAML file.
	//
	// Each construct must implement a `synth()` method.
	Synth(outputDir *string, fileName *string)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for App
type jsiiProxy_App struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_App) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Initializes the App construct.
func NewApp() App {
	_init_.Initialize()

	j := jsiiProxy_App{}

	_jsii_.Create(
		"zap-cdk.App",
		nil, // no parameters
		&j,
	)

	return &j
}

// Initializes the App construct.
func NewApp_Override(a App) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.App",
		nil, // no parameters
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func App_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"zap-cdk.App",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_App) Synth(outputDir *string, fileName *string) {
	_jsii_.InvokeVoid(
		a,
		"synth",
		[]interface{}{outputDir, fileName},
	)
}

func (a *jsiiProxy_App) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

