package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents a data-driven node in the scanning process.
type DataDrivenNode interface {
	IDataDrivenNode
	Name() *string
	SetName(val *string)
	Regex() *string
	SetRegex(val *string)
}

// The jsii proxy struct for DataDrivenNode
type jsiiProxy_DataDrivenNode struct {
	jsiiProxy_IDataDrivenNode
}

func (j *jsiiProxy_DataDrivenNode) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDrivenNode) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}


// Creates an instance of DataDrivenNode.
func NewDataDrivenNode(options IDataDrivenNode) DataDrivenNode {
	_init_.Initialize()

	if err := validateNewDataDrivenNodeParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDrivenNode{}

	_jsii_.Create(
		"zap-cdk.DataDrivenNode",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of DataDrivenNode.
func NewDataDrivenNode_Override(d DataDrivenNode, options IDataDrivenNode) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.DataDrivenNode",
		[]interface{}{options},
		d,
	)
}

func (j *jsiiProxy_DataDrivenNode)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataDrivenNode)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

