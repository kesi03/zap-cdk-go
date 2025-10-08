package zapcdk

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScanPolicyDefinition interface {
	CreatedAt() *time.Time
	SetCreatedAt(c *time.Time)
	Description() *string
	SetDescription(d *string)
	Id() *float64
	SetId(i *float64)
	Name() *string
	SetName(n *string)
	UpdatedAt() *time.Time
	SetUpdatedAt(u *time.Time)
}

// The jsii proxy for IActiveScanPolicyDefinition
type jsiiProxy_IActiveScanPolicyDefinition struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) CreatedAt() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyDefinition)SetCreatedAt(val *time.Time) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyDefinition)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyDefinition)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyDefinition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) UpdatedAt() *time.Time {
	var returns *time.Time
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyDefinition)SetUpdatedAt(val *time.Time) {
	if err := j.validateSetUpdatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedAt",
		val,
	)
}

