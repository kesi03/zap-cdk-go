package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents additional headers for poll request in authentication verification.
type PollAdditionalHeaders interface {
	IPollAdditionalHeaders
	Header() *string
	SetHeader(val *string)
	Value() *string
	SetValue(val *string)
}

// The jsii proxy struct for PollAdditionalHeaders
type jsiiProxy_PollAdditionalHeaders struct {
	jsiiProxy_IPollAdditionalHeaders
}

func (j *jsiiProxy_PollAdditionalHeaders) Header() *string {
	var returns *string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PollAdditionalHeaders) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates an instance of PollAdditionalHeaders.
func NewPollAdditionalHeaders(options IPollAdditionalHeaders) PollAdditionalHeaders {
	_init_.Initialize()

	if err := validateNewPollAdditionalHeadersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PollAdditionalHeaders{}

	_jsii_.Create(
		"zap-cdk.PollAdditionalHeaders",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of PollAdditionalHeaders.
func NewPollAdditionalHeaders_Override(p PollAdditionalHeaders, options IPollAdditionalHeaders) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PollAdditionalHeaders",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PollAdditionalHeaders)SetHeader(val *string) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_PollAdditionalHeaders)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

