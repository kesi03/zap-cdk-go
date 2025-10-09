package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the structure details of the context.
type ContextStructure interface {
	IContextStructure
	DataDrivenNodes() *[]IDataDrivenNode
	SetDataDrivenNodes(val *[]IDataDrivenNode)
	StructuralParameters() *[]*string
	SetStructuralParameters(val *[]*string)
}

// The jsii proxy struct for ContextStructure
type jsiiProxy_ContextStructure struct {
	jsiiProxy_IContextStructure
}

func (j *jsiiProxy_ContextStructure) DataDrivenNodes() *[]IDataDrivenNode {
	var returns *[]IDataDrivenNode
	_jsii_.Get(
		j,
		"dataDrivenNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContextStructure) StructuralParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuralParameters",
		&returns,
	)
	return returns
}


// Creates an instance of ContextStructure.
func NewContextStructure(options IContextStructure) ContextStructure {
	_init_.Initialize()

	if err := validateNewContextStructureParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContextStructure{}

	_jsii_.Create(
		"zap-cdk.ContextStructure",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ContextStructure.
func NewContextStructure_Override(c ContextStructure, options IContextStructure) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ContextStructure",
		[]interface{}{options},
		c,
	)
}

func (j *jsiiProxy_ContextStructure)SetDataDrivenNodes(val *[]IDataDrivenNode) {
	_jsii_.Set(
		j,
		"dataDrivenNodes",
		val,
	)
}

func (j *jsiiProxy_ContextStructure)SetStructuralParameters(val *[]*string) {
	_jsii_.Set(
		j,
		"structuralParameters",
		val,
	)
}

