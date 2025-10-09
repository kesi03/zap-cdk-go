package zapcdk

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the policy definition for an active scan.
type ActiveScanPolicyDefinition interface {
	IActiveScanPolicyDefinition
	CreatedAt() *time.Time
	SetCreatedAt(val *time.Time)
	Description() *string
	SetDescription(val *string)
	Id() *float64
	SetId(val *float64)
	Name() *string
	SetName(val *string)
	UpdatedAt() *time.Time
	SetUpdatedAt(val *time.Time)
}

// The jsii proxy struct for ActiveScanPolicyDefinition
type jsiiProxy_ActiveScanPolicyDefinition struct {
	jsiiProxy_IActiveScanPolicyDefinition
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) CreatedAt() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) UpdatedAt() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanPolicyDefinition.
func NewActiveScanPolicyDefinition(options IActiveScanPolicyDefinition) ActiveScanPolicyDefinition {
	_init_.Initialize()

	if err := validateNewActiveScanPolicyDefinitionParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanPolicyDefinition{}

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicyDefinition",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanPolicyDefinition.
func NewActiveScanPolicyDefinition_Override(a ActiveScanPolicyDefinition, options IActiveScanPolicyDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicyDefinition",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScanPolicyDefinition)SetCreatedAt(val *time.Time) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicyDefinition)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicyDefinition)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicyDefinition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicyDefinition)SetUpdatedAt(val *time.Time) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

