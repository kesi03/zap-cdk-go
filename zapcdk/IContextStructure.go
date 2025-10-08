package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IContextStructure interface {
	DataDrivenNodes() *[]IDataDrivenNode
	SetDataDrivenNodes(d *[]IDataDrivenNode)
	StructuralParameters() *[]*string
	SetStructuralParameters(s *[]*string)
}

// The jsii proxy for IContextStructure
type jsiiProxy_IContextStructure struct {
	_ byte // padding
}

func (j *jsiiProxy_IContextStructure) DataDrivenNodes() *[]IDataDrivenNode {
	var returns *[]IDataDrivenNode
	_jsii_.Get(
		j,
		"dataDrivenNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextStructure)SetDataDrivenNodes(val *[]IDataDrivenNode) {
	_jsii_.Set(
		j,
		"dataDrivenNodes",
		val,
	)
}

func (j *jsiiProxy_IContextStructure) StructuralParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuralParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextStructure)SetStructuralParameters(val *[]*string) {
	_jsii_.Set(
		j,
		"structuralParameters",
		val,
	)
}

