package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a report configuration.
type IReport interface {
	Confidences() *[]*string
	SetConfidences(c *[]*string)
	DisplayReport() *bool
	SetDisplayReport(d *bool)
	ReportDescription() *string
	SetReportDescription(r *string)
	ReportDir() *string
	SetReportDir(r *string)
	ReportFile() *string
	SetReportFile(r *string)
	ReportTitle() *string
	SetReportTitle(r *string)
	Risks() *[]*string
	SetRisks(r *[]*string)
	Sections() *[]*string
	SetSections(s *[]*string)
	Sites() *[]*string
	SetSites(s *[]*string)
	Template() *string
	SetTemplate(t *string)
	Theme() *string
	SetTheme(t *string)
}

// The jsii proxy for IReport
type jsiiProxy_IReport struct {
	_ byte // padding
}

func (j *jsiiProxy_IReport) Confidences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"confidences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetConfidences(val *[]*string) {
	_jsii_.Set(
		j,
		"confidences",
		val,
	)
}

func (j *jsiiProxy_IReport) DisplayReport() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"displayReport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetDisplayReport(val *bool) {
	_jsii_.Set(
		j,
		"displayReport",
		val,
	)
}

func (j *jsiiProxy_IReport) ReportDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetReportDescription(val *string) {
	_jsii_.Set(
		j,
		"reportDescription",
		val,
	)
}

func (j *jsiiProxy_IReport) ReportDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetReportDir(val *string) {
	_jsii_.Set(
		j,
		"reportDir",
		val,
	)
}

func (j *jsiiProxy_IReport) ReportFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetReportFile(val *string) {
	_jsii_.Set(
		j,
		"reportFile",
		val,
	)
}

func (j *jsiiProxy_IReport) ReportTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetReportTitle(val *string) {
	_jsii_.Set(
		j,
		"reportTitle",
		val,
	)
}

func (j *jsiiProxy_IReport) Risks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"risks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetRisks(val *[]*string) {
	_jsii_.Set(
		j,
		"risks",
		val,
	)
}

func (j *jsiiProxy_IReport) Sections() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetSections(val *[]*string) {
	_jsii_.Set(
		j,
		"sections",
		val,
	)
}

func (j *jsiiProxy_IReport) Sites() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetSites(val *[]*string) {
	_jsii_.Set(
		j,
		"sites",
		val,
	)
}

func (j *jsiiProxy_IReport) Template() *string {
	var returns *string
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetTemplate(val *string) {
	_jsii_.Set(
		j,
		"template",
		val,
	)
}

func (j *jsiiProxy_IReport) Theme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"theme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReport)SetTheme(val *string) {
	_jsii_.Set(
		j,
		"theme",
		val,
	)
}

