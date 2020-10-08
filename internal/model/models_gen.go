// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
)

// User is either an Admin or a Participant.
type User interface {
	IsUser()
}

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResp struct {
	Token string `json:"token"`
}

type CancelRunInput struct {
	RunID string `json:"runID"`
}

// Possible Condition values. Only one of the fields in a CompValue should be
// defined.
type CompValue struct {
	Int     *int     `json:"int"`
	Float   *float64 `json:"float"`
	String  *string  `json:"string"`
	Boolean *bool    `json:"boolean"`
}

// Possible Condition values. Only one of the fields in a CompValue should be
// defined.
type CompValueInput struct {
	Int     *int     `json:"int"`
	Float   *float64 `json:"float"`
	String  *string  `json:"string"`
	Boolean *bool    `json:"boolean"`
}

// Condition for a filter. A condition must **either**:
// - have one or more `and` Conditions
// - have one or more `or` Conditions
// - have a comparator and one or more values
// When comparing against values, the first value is used for operands comparing
// single values (LessThan, LessThanOrEqualTo, GreaterThan, GreaterThanOrEqualTo,
// EqualTo, NotEqualTo). When testing for existence an empty array is DoesNotExist
// and an array with one or more values Exists. For In and NotIn all values in the
// values array are used.
// If the  condition is empty (no and, no or, no key), it will match any record.
type Condition struct {
	And        []*Condition `json:"and"`
	Or         []*Condition `json:"or"`
	Key        *string      `json:"key"`
	Comparator *Comparator  `json:"comparator"`
	Values     []*CompValue `json:"values"`
}

// Condition for a filter. A condition must **either**:
// - have one or more `and` Conditions
// - have one or more `or` Conditions
// - have a comparator and one or more values
// When comparing against values, the first value is used for operands comparing
// single values (LessThan, LessThanOrEqualTo, GreaterThan, GreaterThanOrEqualTo,
// EqualTo, NotEqualTo). When testing for existence an empty array is DoesNotExist
// and an array with one or more values Exists. For In and NotIn all values in the
// values array are used.
type ConditionInput struct {
	And        []*ConditionInput `json:"and"`
	Or         []*ConditionInput `json:"or"`
	Key        *string           `json:"key"`
	Comparator *Comparator       `json:"comparator"`
	Values     []*CompValueInput `json:"values"`
}

type CreateProcedureInput struct {
	// Project in which to create the Procedure.
	ProjectID string          `json:"projectID"`
	Procedure *ProcedureInput `json:"procedure"`
}

type CreateProjectInput struct {
	ProjectID string `json:"projectID"`
	Name      string `json:"name"`
}

type CreateRunInput struct {
	ProjectID string          `json:"projectID"`
	Procedure *ProcedureInput `json:"procedure"`
}

// Datum is a single piece of custom data.
type Datum struct {
	// id is the unique globally identifier for the record.
	ID string `json:"id"`
	// createdAt is the time of creation of the record.
	CreatedAt time.Time `json:"createdAt"`
	// updatedAt is the time of last update of the record.
	UpdatedAt time.Time `json:"updatedAt"`
	// Creator is the user or participant that created the Datum.
	Creator User `json:"creator"`
	// deletedAt is the time when the Datum was deleted. If null, the Datum was not
	// deleted.
	DeletedAt *time.Time `json:"deletedAt"`
	// key identifies the unique key of the Datum.
	Key string `json:"key"`
	// val is the value of the Datum. It can be any JSON encodable value.
	// Passing null will delete the Datum.
	Val *string `json:"val"`
	// next returns the Datum ID of the next value in an array of Datum values.
	Next *string `json:"next"`
	// root returns the Datum ID of the root value in an array of Datum values.
	Root *string `json:"root"`
	// versions returns previous versions for the Datum (they all have the same ID).
	Versions []*Datum `json:"versions"`
}

type DuplicateProcedureInput struct {
	ProcedureID string  `json:"procedureID"`
	ProjectID   string  `json:"projectID"`
	Name        *string `json:"name"`
}

// FilterStepArgs are arguments passed to a Pariticipant Filter Step.
// It must contains **either** JS code or the name of pre-defined filtering function.
// This is only valid for an PARTICIPANT_FILTER Step.
type FilterStepArgs struct {
	// Type is whether to use a predefined filter, JS code, or the Condition filter
	// mechanism.
	Type ParticipantFilterType `json:"type"`
	// Filter should be the name of pre-defined filtering function.
	Filter *string `json:"filter"`
	// Javascript to execute as a participant filter step.
	// The code must contain a functinon exported using a default ES6 export.
	// The function should accept a single argument object. This object contains the
	// following fields:
	// - `participants`: the participants entering this step
	// - `step`: this step (contains the definition of this step: duration, etc.)
	// - `stepRun`: instance of this step (contains the execution of this step: start time, etc.)
	// - `procedure`: parent procedure of step (contains the definition of the Procedure)
	// - `run`: run this step is part of (contains the instance of the Procedure)
	// The functions should return an array of participants.
	// If the functions returns null or undefined, the participants are not filtered.
	// If the function throws an exception, the run will fail.
	Js *string `json:"js"`
	// Condition set the participant must meet to be allowed to participate.
	Condition *Condition `json:"condition"`
}

// FilterStepArgs are arguments passed to a Pariticipant Filter Step.
// It must contains **either** JS code or the name of pre-defined filtering function.
// This is only valid for an PARTICIPANT_FILTER Step.
type FilterStepArgsInput struct {
	// Type is whether to use a predefined filter, JS code, or the Condition filter
	// mechanism.
	Type ParticipantFilterType `json:"type"`
	// Javascript to execute as a participant filter step.
	// The code must contain a functinon exported using a default ES6 export.
	// The function should accept a single argument object. This object contains the
	// following fields:
	// - `participants`: the participants entering this step
	// - `step`: this step (contains the definition of this step: duration, etc.)
	// - `stepRun`: instance of this step (contains the execution of this step: start time, etc.)
	// - `procedure`: parent procedure of step (contains the definition of the Procedure)
	// - `run`: run this step is part of (contains the instance of the Procedure)
	// The functions should return an array of participants.
	// If the functions returns null or undefined, the participants are not filtered.
	// If the function throws an exception, the run will fail.
	Js *string `json:"js"`
	// Filter should be the name of pre-defined filtering function.
	Filter *string `json:"filter"`
	// Condition set the participant must meet to be allowed to participate.
	Condition *ConditionInput `json:"condition"`
}

// HITStepArgs are arguments passed to a HIT Step.
// This is only valid for an MTURK_HIT Step.
type HITStepArgs struct {
	// Title of HIT.
	// From MTurk: Describe the task to Workers. Be as specific as possible,
	// e.g. "answer a survey about movies", instead of "short survey", so Workers
	// know what to expect.
	// Tasks that contain adult content are required to include the following phrase
	// in your task title: (WARNING: This HIT may contain adult content. Worker
	// discretion is advised.)
	Title string `json:"title"`
	// Description of HIT.
	// From MTurk: Give more detail about this task. This gives Workers a bit more
	// information before they decide to view your task.
	Description string `json:"description"`
	// Keywords of HIT. Comma-seratred.
	// From MTurk: Provide keywords that will help Workers search for your tasks.
	Keywords string `json:"keywords"`
	// DISABLED - Micro-batching is still TBD, probably needs more args.
	Microbatch bool `json:"microbatch"`
	// MTurk HIT reward for task in USD.
	Reward float64 `json:"reward"`
	// Timeout of a single accepted HIT in seconds.
	Timeout int `json:"timeout"`
	// Duration in seconds from start of Step before expiration of unconsumed HITs.
	Duration int `json:"duration"`
	// Number of HIT workers to accept.
	// Note: is this needed? The count is determined by the selection in the first
	// Step, then by the number of participants remaining at each Step.
	WorkersCount int `json:"workersCount"`
}

// HITStepArgs are arguments passed to a HIT Step.
// This is only valid for an MTURK_HIT Step.
type HITStepArgsInput struct {
	// Title of HIT.
	// From MTurk: Describe the task to Workers. Be as specific as possible,
	// e.g. "answer a survey about movies", instead of "short survey", so Workers
	// know what to expect.
	// Tasks that contain adult content are required to include the following phrase
	// in your task title: (WARNING: This HIT may contain adult content. Worker
	// discretion is advised.)
	Title *string `json:"title"`
	// Description of HIT.
	// From MTurk: Give more detail about this task. This gives Workers a bit more
	// information before they decide to view your task.
	Description *string `json:"description"`
	// Keywords of HIT. Comma-seratred.
	// From MTurk: Provide keywords that will help Workers search for your tasks.
	Keywords *string `json:"keywords"`
	// DISABLED - Micro-batching is still TBD, probably needs more args.
	Microbatch *bool `json:"microbatch"`
	// MTurk HIT reward for task in USD.
	Reward *float64 `json:"reward"`
	// Timeout of a single accepted HIT in seconds.
	Timeout *int `json:"timeout"`
	// Duration in seconds from start of Step before expiration of unconsumed HITs.
	Duration int `json:"duration"`
	// Number of HIT workers to accept.
	// Note: is this needed? The count is determined by the selection in the first
	// Step, then by the number of participants remaining at each Step.
	WorkersCount *int `json:"workersCount"`
}

// InternalCriteria is the criteria for internal database participant selection.
type InternalCriteria struct {
	// All means use all participants and ignore the condition field below.
	All bool `json:"all"`
	// Condition set the participant must meet to be allowed to participate.
	Condition *Condition `json:"condition"`
}

// InternalCriteria is the criteria for internal database participant selection.
type InternalCriteriaInput struct {
	// All means use all participants and ignore the condition field below.
	All bool `json:"all"`
	// Condition set the participant must meet to be allowed to participate.
	Condition *ConditionInput `json:"condition"`
}

// MTurkCriteria is the criteria for MTurk Qualifications participant selection.
type MTurkCriteria struct {
	// MTurk Qualifications a Worker must meet before the Worker is allowed to accept
	// and complete the HIT.
	Qualifications []*MTurkQualificationCriteria `json:"qualifications"`
}

// MTurkCriteria is the criteria for MTurk Qualifications participant selection.
type MTurkCriteriaInput struct {
	// MTurk Qualifications a Worker must meet before the Worker is allowed to accept
	// and complete the HIT.
	Qualifications []*MTurkQualificationCriteriaInput `json:"qualifications"`
}

// The Locale data structure represents a geographical region or location in MTurk.
type MTurkLocale struct {
	// The country of the locale.
	// Type: A valid ISO 3166 country code. For example, the code US refers to the
	// United States of America.
	Country string `json:"country"`
	// The state or subdivision of the locale.
	// Type: Type: A valid ISO 3166-2 subdivision code. For example, the code CA
	// refers to the state of California.
	// Subdivisions or states are only available for the United States of America.
	Subdivision *string `json:"subdivision"`
}

// The Locale data structure represents a geographical region or location in MTurk.
type MTurkLocaleInput struct {
	// The country of the locale.
	// Type: A valid ISO 3166 country code. For example, the code US refers to the
	// United States of America.
	Country string `json:"country"`
	// The state or subdivision of the locale.
	// Type: Type: A valid ISO 3166-2 subdivision code. For example, the code CA
	// refers to the state of California.
	// Subdivisions or states are only available for the United States of America.
	Subdivision *string `json:"subdivision"`
}

// MTurkQualificationCriteria is an MTurk Qualification requirement. It is an
// MTurk Qualification that a Worker must have before the Worker is allowed to
// accept a HIT.
// See https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_QualificationRequirementDataStructureArticle.html
type MTurkQualificationCriteria struct {
	// The ID of the MTurk Qualification Type.
	ID string `json:"id"`
	// The kind of comparison to make against a Qualification's value.
	// You can compare a Qualification's value:
	// - To an IntegerValue to see if it is LessThan, LessThanOrEqualTo, GreaterThan,
	//   GreaterThanOrEqualTo, EqualTo, or NotEqualTo the IntegerValue.
	// - To a LocaleValue to see if it is EqualTo, or NotEqualTo the LocaleValue.
	// - To see if the value is In or NotIn a set of IntegerValue or LocaleValue
	//   values.
	// A Qualification requirement can also test if a Qualification Exists or
	// DoesNotExist in the user's profile, regardless of its value.
	Comparator Comparator `json:"comparator"`
	// Array of integer values to compare against the Qualification's value.
	// IntegerValue must not be present if Comparator is Exists or DoesNotExist.
	// IntegerValue can only be used if the Qualification type has an integer value;
	// it cannot be used with the Worker_Locale QualificationType ID, see
	// Qualification Type IDs.
	// When performing a set comparison by using the In or the NotIn comparator, you
	// can use up to 15 elements in this list.
	Values []int `json:"values"`
	// The locale value to compare against the Qualification's value. The local value
	// must be a valid ISO 3166 country code or supports ISO 3166-2 subdivisions.
	// LocaleValue can only be used with a Worker_Locale QualificationType ID, see
	// Qualification Type IDs.
	// LocaleValue can only be used with the EqualTo, NotEqualTo, In, and NotIn
	// comparators.
	// You must only use a single LocaleValue element when using the EqualTo or
	// NotEqualTo comparators.
	// When performing a set comparison by using the In or the NotIn comparator, you
	// can use up to 30 LocaleValue elements in a QualificationRequirement data
	// structure.
	Locales []*MTurkLocale `json:"locales"`
}

// MTurkQualificationCriteria is an MTurk Qualification requirement. It is an
// MTurk Qualification that a Worker must have before the Worker is allowed to
// accept a HIT.
// See https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_QualificationRequirementDataStructureArticle.html
type MTurkQualificationCriteriaInput struct {
	// The ID of the MTurk Qualification Type.
	ID string `json:"id"`
	// The kind of comparison to make against a Qualification's value.
	// You can compare a Qualification's value:
	// - To an IntegerValue to see if it is LessThan, LessThanOrEqualTo, GreaterThan,
	//   GreaterThanOrEqualTo, EqualTo, or NotEqualTo the IntegerValue.
	// - To a LocaleValue to see if it is EqualTo, or NotEqualTo the LocaleValue.
	// - To see if the value is In or NotIn a set of IntegerValue or LocaleValue
	//   values.
	// A Qualification requirement can also test if a Qualification Exists or
	// DoesNotExist in the user's profile, regardless of its value.
	Comparator Comparator `json:"comparator"`
	// Array of integer values to compare against the Qualification's value.
	// IntegerValue must not be present if Comparator is Exists or DoesNotExist.
	// IntegerValue can only be used if the Qualification type has an integer value;
	// it cannot be used with the Worker_Locale QualificationType ID, see
	// Qualification Type IDs.
	// When performing a set comparison by using the In or the NotIn comparator, you
	// can use up to 15 elements in this list.
	Values []int `json:"values"`
	// The locale value to compare against the Qualification's value. The local value
	// must be a valid ISO 3166 country code or supports ISO 3166-2 subdivisions.
	// LocaleValue can only be used with a Worker_Locale QualificationType ID, see
	// Qualification Type IDs.
	// LocaleValue can only be used with the EqualTo, NotEqualTo, In, and NotIn
	// comparators.
	// You must only use a single LocaleValue element when using the EqualTo or
	// NotEqualTo comparators.
	// When performing a set comparison by using the In or the NotIn comparator, you
	// can use up to 30 LocaleValue elements in a QualificationRequirement data
	// structure.
	Locales []*MTurkLocaleInput `json:"locales"`
}

// The QualificationType data structure represents a Qualification type, a description of a property of a Worker that must
// match the requirements of a HIT for the Worker to be able to accept the HIT. The type also describes how a Worker can obtain
// a Qualification of that type, such as through a Qualification test.
// See https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_QualificationTypeDataStructureArticle.html
type MTurkQulificationType struct {
	// A unique identifier for the Qualification type. A Qualification type is given a Qualification type ID when you call
	// the CreateQualificationType operation operation, and it retains that ID forever. Can be up to 255 bytes in length.
	ID string `json:"id"`
	// The name of the Qualification type. The type name is used to identify the type, and to find the type using a Qualification type search.
	Name string `json:"name"`
	// A long description for the Qualification type.
	Description string `json:"description"`
	// A type that is used to define the comparator.
	Type QualType `json:"type"`
}

// MessageStepArgs are arguments passed to a Step that has a message.
// This is only valid for MTURK_HIT and MTURK_MESSAGE Steps.
type MessageStepArgs struct {
	// URL that will be transformed into a redirect (proxy URL) through the Empirica
	// Recruitment website and passed to the Message template. This URL is the final
	// destination the worker will land on. Empirica Recruitment redirects
	// through the app so we can add parameters to the proxy URL and hide the final
	// URL (to limit sharing of URLs).
	URL *string `json:"url"`
	// Message the content to display to the user.
	// Template variables:
	// - `url`: proxy URL if `url` exist on Step.
	// - `step`: this step (contains the definition of this step: duration, etc.)
	// - `stepRun`: instance of this step (contains the execution of this step: start time, etc.)
	// - `procedure`: parent Procedure of step (contains the definition of the Procedure)
	// - `run`: run this step is part of (contains the instance of the Procedure)
	// - `participant`: current participant
	Message string `json:"message"`
	// MessageType indicates the rendering language of the Message.
	MessageType ContentType `json:"messageType"`
	// Lobby enables to showing a lobby, and rich-text message to put in the lobby
	// Lobby can either expire (see expiration below) to produce the effect of a
	// precise start time, or must have a submit button.
	// Only available if URL is present.
	// Template variables are identical to message.
	Lobby *string `json:"lobby"`
	// LobbyType indicates the rendering language of the Lobby.
	LobbyType *ContentType `json:"lobbyType"`
	// LobbyExpirtation in seconds from the beginning of the step.
	LobbyExpiration *int `json:"lobbyExpiration"`
}

// MessageStepArgs are arguments passed to a Step that has a message.
// This is only valid for MTURK_HIT and MTURK_MESSAGE Steps.
type MessageStepArgsInput struct {
	// URL that will be transformed into a redirect (proxy URL) through the Empirica
	// Recruitment website and passed to the Message template. This URL is the final
	// destination the worker will land on. Empirica Recruitment redirects
	// through the app so we can add parameters to the proxy URL and hide the final
	// URL (to limit sharing of URLs).
	URL *string `json:"url"`
	// Message the content to display to the user.
	// Template variables:
	// - `url`: proxy URL if `url` exist on Step.
	// - `step`: this step (contains the definition of this step: duration, etc.)
	// - `stepRun`: instance of this step (contains the execution of this step: start time, etc.)
	// - `procedure`: parent Procedure of step (contains the definition of the Procedure)
	// - `run`: run this step is part of (contains the instance of the Procedure)
	// - `participant`: current participant
	Message *string `json:"message"`
	// MessageType indicates the rendering language of the Message.
	MessageType *ContentType `json:"messageType"`
	// Lobby enables to showing a lobby, and rich-text message to put in the lobby
	// Lobby can either expire (see expiration below) to produce the effect of a
	// precise start time, or must have a submit button.
	// Only available if URL is present.
	// Template variables are identical to message.
	Lobby *string `json:"lobby"`
	// LobbyType indicates the rendering language of the Lobby.
	LobbyType *ContentType `json:"lobbyType"`
	// LobbyExpirtation in seconds from the beginning of the step.
	LobbyExpiration *int `json:"lobbyExpiration"`
}

// MutateDatumInput adds/appends/updates/deletes Data to a Node.
type MutateDatumInput struct {
	// Operation to perform on Datum.
	Operation *DatumOp `json:"operation"`
	// key identifies the unique key of the Datum.
	Key string `json:"key"`
	// val is the value of the Datum. It can be any JSON encodable value.
	// If Delete op, value is ignored.
	Val *string `json:"val"`
	// ID of object on which to set the value.
	NodeID string `json:"nodeID"`
}

// Page returns everything needed to display a page to Participants or redirect
// them, depending on the Step, whether it's a message, or a lobby, and the token
// used to acquire the page.
type Page struct {
	// Redirect is URL. If present, the page should just redirect to the given URL.
	// No other field should be present if the redirect is non-null.
	Redirect *string `json:"redirect"`
	// Props to be passed to the Content (if the Content is dynamic).
	Props map[string]interface{} `json:"props"`
	// The content to display to the Participant.
	Content *string `json:"content"`
	// The rendered for the content to display.
	ContentType ContentType `json:"contentType"`
}

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

// Participant is a worker in the system.
type Participant struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// Step during which the Participant was created.
	CreatedBy *ent.StepRun `json:"createdBy"`
	// All StepRuns the Participant participated in.
	Steps []*ent.StepRun `json:"steps"`
	// ProviderIDs contains the IDs from 3rd providers corresponding the participant.
	// A single participant could potentially be referenced in different in multiple
	// providers.
	ProviderIDs []*ProviderID `json:"providerIDs"`
	// Data returns the custom data that has been set on the Participant.
	Data []*Datum `json:"data"`
}

func (Participant) IsUser() {}

type ParticipantsConnection struct {
	TotalCount   int                 `json:"totalCount"`
	Edges        []*ParticipantsEdge `json:"edges"`
	Participants []*Participant      `json:"participants"`
	PageInfo     *PageInfo           `json:"pageInfo"`
}

type ParticipantsEdge struct {
	Cursor string       `json:"cursor"`
	Node   *Participant `json:"node"`
}

type ProcedureInput struct {
	// ID of procedure to update, if updating
	ID *string `json:"id"`
	// Friendly name.
	Name string `json:"name"`
	// Ordered list of Steps for Procedure.
	Steps []*StepInput `json:"steps"`
	// Determines participant selection type.
	SelectionType SelectionType `json:"selectionType"`
	// Internal Selection criteria for participants
	InternalCriteria *InternalCriteriaInput `json:"internalCriteria"`
	// Mturk Selection criteria for participants
	MturkCriteria *MTurkCriteriaInput `json:"mturkCriteria"`
	// Number of participants desired.
	ParticipantCount int `json:"participantCount"`
	// Contains adult content.
	// From MTurk: This project may contain potentially explicit or offensive
	// content, for example, nudity.
	Adult bool `json:"adult"`
}

// ProviderID contains the identifier for a 3rd party provider.
type ProviderID struct {
	// createdAt is the time of creation of the record.
	CreatedAt time.Time `json:"createdAt"`
	// providerID is the ID of the 3rd party Provider.
	ProviderID string `json:"providerID"`
	// ID is the ID of the 3rd party Provider.
	Provider *Provider `json:"provider"`
}

type RegisterParticipantInput struct {
	// ID from provider.
	ID string `json:"id"`
	// Provider of ID.
	Provider *Provider `json:"provider"`
	// Initial Data to attach to Participant.
	Data map[string]interface{} `json:"data"`
}

type ScheduleRunInput struct {
	RunID   string    `json:"runID"`
	StartAt time.Time `json:"startAt"`
}

type StartRunInput struct {
	RunID string `json:"runID"`
}

type StepInput struct {
	// ID of the step to update, if updating
	ID *string `json:"id"`
	// Index is the position of the step in the Procedure.
	Index int `json:"index"`
	// The Type defines what kind of action this step represents.
	Type StepType `json:"type"`
	// Duration of Step in seconds. At the end of the duration, the next Step will
	// execute.
	// If set to 0, the Step executes and immediately moves onto the next Step. This
	// mostly works for PARTICIPANT_FILTER Steps and the last Step in a Procedure.
	Duration int `json:"duration"`
	// Arguments for Message type Step.
	MsgArgs *MessageStepArgsInput `json:"msgArgs"`
	// Arguments for HIT type Step.
	HitArgs *HITStepArgsInput `json:"hitArgs"`
	// Arguments for Filter type Step.
	FilterArgs *FilterStepArgsInput `json:"filterArgs"`
}

type UnscheduleRunInput struct {
	RunID string `json:"runID"`
}

type UpdateProcedureInput struct {
	ProjectID string          `json:"projectID"`
	RunID     string          `json:"runID"`
	Procedure *ProcedureInput `json:"procedure"`
}

type UpdateRunInput struct {
	ID        string `json:"ID"`
	ProjectID string `json:"projectID"`
	Name      string `json:"name"`
}

// The kind of comparison to make against a value.
type Comparator string

const (
	ComparatorLessThan             Comparator = "LESS_THAN"
	ComparatorLessThanOrEqualTo    Comparator = "LESS_THAN_OR_EQUAL_TO"
	ComparatorGreaterThan          Comparator = "GREATER_THAN"
	ComparatorGreaterThanOrEqualTo Comparator = "GREATER_THAN_OR_EQUAL_TO"
	ComparatorEqualTo              Comparator = "EQUAL_TO"
	ComparatorNotEqualTo           Comparator = "NOT_EQUAL_TO"
	ComparatorExists               Comparator = "EXISTS"
	ComparatorDoesNotExist         Comparator = "DOES_NOT_EXIST"
	ComparatorIn                   Comparator = "IN"
	ComparatorNotIn                Comparator = "NOT_IN"
)

var AllComparator = []Comparator{
	ComparatorLessThan,
	ComparatorLessThanOrEqualTo,
	ComparatorGreaterThan,
	ComparatorGreaterThanOrEqualTo,
	ComparatorEqualTo,
	ComparatorNotEqualTo,
	ComparatorExists,
	ComparatorDoesNotExist,
	ComparatorIn,
	ComparatorNotIn,
}

func (e Comparator) IsValid() bool {
	switch e {
	case ComparatorLessThan, ComparatorLessThanOrEqualTo, ComparatorGreaterThan, ComparatorGreaterThanOrEqualTo, ComparatorEqualTo, ComparatorNotEqualTo, ComparatorExists, ComparatorDoesNotExist, ComparatorIn, ComparatorNotIn:
		return true
	}
	return false
}

func (e Comparator) String() string {
	return string(e)
}

func (e *Comparator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Comparator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Comparator", str)
	}
	return nil
}

func (e Comparator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// ContentType is the type rendering used for a field.
type ContentType string

const (
	// MARKDOWN uses a Markdown renderer. Templating uses Mustache-style
	// interpolation (i.e. {{url}}).
	ContentTypeMarkdown ContentType = "MARKDOWN"
	// HTML uses an HTML rendered. Templating uses Mustache-style
	// interpolation (i.e. {{url}}).
	ContentTypeHTML ContentType = "HTML"
	// REACT uses a React renderer. Templating passes template arguments as props.
	// The root component should be the default ES6 export.
	ContentTypeReact ContentType = "REACT"
	// SVELTE uses a Svelte renderer. Templating passes template arguments as props.
	// The root component should be the default ES6 export.
	ContentTypeSvelte ContentType = "SVELTE"
)

var AllContentType = []ContentType{
	ContentTypeMarkdown,
	ContentTypeHTML,
	ContentTypeReact,
	ContentTypeSvelte,
}

func (e ContentType) IsValid() bool {
	switch e {
	case ContentTypeMarkdown, ContentTypeHTML, ContentTypeReact, ContentTypeSvelte:
		return true
	}
	return false
}

func (e ContentType) String() string {
	return string(e)
}

func (e *ContentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentType", str)
	}
	return nil
}

func (e ContentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Operation to perform on a Datum
type DatumOp string

const (
	// Set the datum to given value.
	DatumOpSet DatumOp = "SET"
	// Append the datum to given value.
	DatumOpAppend DatumOp = "APPEND"
	// Delete the datum.
	DatumOpDelete DatumOp = "DELETE"
)

var AllDatumOp = []DatumOp{
	DatumOpSet,
	DatumOpAppend,
	DatumOpDelete,
}

func (e DatumOp) IsValid() bool {
	switch e {
	case DatumOpSet, DatumOpAppend, DatumOpDelete:
		return true
	}
	return false
}

func (e DatumOp) String() string {
	return string(e)
}

func (e *DatumOp) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DatumOp(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DatumOp", str)
	}
	return nil
}

func (e DatumOp) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Supported 3rd party providers.
type Provider string

const (
	// MTURK represents AWS Mechanical Turk
	ProviderMturk Provider = "MTURK"
)

var AllProvider = []Provider{
	ProviderMturk,
}

func (e Provider) IsValid() bool {
	switch e {
	case ProviderMturk:
		return true
	}
	return false
}

func (e Provider) String() string {
	return string(e)
}

func (e *Provider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Provider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PROVIDER", str)
	}
	return nil
}

func (e Provider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// ParticipantFilterType is the type of user filtering to use.
type ParticipantFilterType string

const (
	// Predefined Filter is a server side defined filter.
	ParticipantFilterTypePredefinedFilter ParticipantFilterType = "PREDEFINED_FILTER"
	// JS is a piece of JAvascript code doing the filtering.
	ParticipantFilterTypeJs ParticipantFilterType = "JS"
	// CONDITION uses the Condition object to do the filtering.
	ParticipantFilterTypeCondition ParticipantFilterType = "CONDITION"
)

var AllParticipantFilterType = []ParticipantFilterType{
	ParticipantFilterTypePredefinedFilter,
	ParticipantFilterTypeJs,
	ParticipantFilterTypeCondition,
}

func (e ParticipantFilterType) IsValid() bool {
	switch e {
	case ParticipantFilterTypePredefinedFilter, ParticipantFilterTypeJs, ParticipantFilterTypeCondition:
		return true
	}
	return false
}

func (e ParticipantFilterType) String() string {
	return string(e)
}

func (e *ParticipantFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ParticipantFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ParticipantFilterType", str)
	}
	return nil
}

func (e ParticipantFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The kind of type used on MTurkQualificationType that later we be used to decide the comparator.
type QualType string

const (
	QualTypeBool       QualType = "BOOL"
	QualTypeComparison QualType = "COMPARISON"
	QualTypeLocation   QualType = "LOCATION"
	QualTypeCustom     QualType = "CUSTOM"
)

var AllQualType = []QualType{
	QualTypeBool,
	QualTypeComparison,
	QualTypeLocation,
	QualTypeCustom,
}

func (e QualType) IsValid() bool {
	switch e {
	case QualTypeBool, QualTypeComparison, QualTypeLocation, QualTypeCustom:
		return true
	}
	return false
}

func (e QualType) String() string {
	return string(e)
}

func (e *QualType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = QualType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QualType", str)
	}
	return nil
}

func (e QualType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin       Role = "ADMIN"
	RoleParticipant Role = "PARTICIPANT"
)

var AllRole = []Role{
	RoleAdmin,
	RoleParticipant,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleParticipant:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Participant selection type.
type SelectionType string

const (
	// INTERNAL_DB uses local participants database.
	SelectionTypeInternalDb SelectionType = "INTERNAL_DB"
	// MTURK_QUALIFICATIONS uses MTurk Qualitifications to select participants.
	SelectionTypeMturkQualifications SelectionType = "MTURK_QUALIFICATIONS"
)

var AllSelectionType = []SelectionType{
	SelectionTypeInternalDb,
	SelectionTypeMturkQualifications,
}

func (e SelectionType) IsValid() bool {
	switch e {
	case SelectionTypeInternalDb, SelectionTypeMturkQualifications:
		return true
	}
	return false
}

func (e SelectionType) String() string {
	return string(e)
}

func (e *SelectionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SelectionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SelectionType", str)
	}
	return nil
}

func (e SelectionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Status of Runs.
type Status string

const (
	// CREATED means the run has been created but hasn't started yet
	StatusCreated Status = "CREATED"
	// RUNNING means the run is currently in progress
	StatusRunning Status = "RUNNING"
	// PAUSED means the run has been paused by an admin
	StatusPaused Status = "PAUSED"
	// DONE means the run has finished `naturally`
	StatusDone Status = "DONE"
	// TERMINATED means the run has been manually cancelled
	StatusTerminated Status = "TERMINATED"
	// FAILED means the run has failed (due to an unrecoverable error)
	StatusFailed Status = "FAILED"
)

var AllStatus = []Status{
	StatusCreated,
	StatusRunning,
	StatusPaused,
	StatusDone,
	StatusTerminated,
	StatusFailed,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusCreated, StatusRunning, StatusPaused, StatusDone, StatusTerminated, StatusFailed:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Type of a Step.
type StepType string

const (
	// MTURK_HIT is a step where AWS Mechanical Turk HIT is published.
	StepTypeMturkHit StepType = "MTURK_HIT"
	// MTURK_MESSAGE is a step where a message sent to AWS Mechanical Turk Workers.
	StepTypeMturkMessage StepType = "MTURK_MESSAGE"
	// PARTICIPANT_FILTER is a participant filtering step.
	StepTypeParticipantFilter StepType = "PARTICIPANT_FILTER"
)

var AllStepType = []StepType{
	StepTypeMturkHit,
	StepTypeMturkMessage,
	StepTypeParticipantFilter,
}

func (e StepType) IsValid() bool {
	switch e {
	case StepTypeMturkHit, StepTypeMturkMessage, StepTypeParticipantFilter:
		return true
	}
	return false
}

func (e StepType) String() string {
	return string(e)
}

func (e *StepType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StepType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StepType", str)
	}
	return nil
}

func (e StepType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
