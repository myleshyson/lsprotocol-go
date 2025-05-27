package protocol


import (
	"encoding/json"
	"fmt"
	"bytes"
	"reflect"
	"strings"
)
type DocumentUri string

type URI string

// An identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier string

// The declaration of a symbol representation as one or many {@link Location locations}.
type Declaration Or2[Location, []Location]

func (t *Declaration) UnmarshalJSON(x []byte) error {
	return (*Or2[Location, []Location])(t).UnmarshalJSON(x)
}
func (t *Declaration) MarshalJSON() ([]byte, error) {
	return (*Or2[Location, []Location])(t).MarshalJSON()
}
// Information about where a symbol is declared.
// 
// Provides additional metadata over normal {@link Location location} declarations, including the range of
// the declaring symbol.
// 
// Servers should prefer returning `DeclarationLink` over `Declaration` if supported
// by the client.
type DeclarationLink LocationLink

// The definition of a symbol represented as one or many {@link Location locations}.
// For most programming languages there is only one location at which a symbol is
// defined.
// 
// Servers should prefer returning `DefinitionLink` over `Definition` if supported
// by the client.
type Definition Or2[Location, []Location]

func (t *Definition) UnmarshalJSON(x []byte) error {
	return (*Or2[Location, []Location])(t).UnmarshalJSON(x)
}
func (t *Definition) MarshalJSON() ([]byte, error) {
	return (*Or2[Location, []Location])(t).MarshalJSON()
}
// Information about where a symbol is defined.
// 
// Provides additional metadata over normal {@link Location location} definitions, including the range of
// the defining symbol
type DefinitionLink LocationLink

// The result of a document diagnostic pull request. A report can
// either be a full report containing all diagnostics for the
// requested document or an unchanged report indicating that nothing
// has changed in terms of diagnostics in comparison to the last
// pull request.
// 
// @since 3.17.0
type DocumentDiagnosticReport Or2[RelatedFullDocumentDiagnosticReport, RelatedUnchangedDocumentDiagnosticReport]

func (t *DocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	return (*Or2[RelatedFullDocumentDiagnosticReport, RelatedUnchangedDocumentDiagnosticReport])(t).UnmarshalJSON(x)
}
func (t *DocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	return (*Or2[RelatedFullDocumentDiagnosticReport, RelatedUnchangedDocumentDiagnosticReport])(t).MarshalJSON()
}
// A document filter describes a top level text document or
// a notebook cell document.
// 
// @since 3.17.0 - support for NotebookCellTextDocumentFilter.
type DocumentFilter Or2[TextDocumentFilter, NotebookCellTextDocumentFilter]

func (t *DocumentFilter) UnmarshalJSON(x []byte) error {
	return (*Or2[TextDocumentFilter, NotebookCellTextDocumentFilter])(t).UnmarshalJSON(x)
}
func (t *DocumentFilter) MarshalJSON() ([]byte, error) {
	return (*Or2[TextDocumentFilter, NotebookCellTextDocumentFilter])(t).MarshalJSON()
}
// A document selector is the combination of one or many document filters.
// 
// @sample `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`;
// 
// The use of a string as a document filter is deprecated @since 3.16.0.
type DocumentSelector []DocumentFilter

// The glob pattern. Either a string pattern or a relative pattern.
// 
// @since 3.17.0
type GlobPattern Or2[Pattern, RelativePattern]

func (t *GlobPattern) UnmarshalJSON(x []byte) error {
	return (*Or2[Pattern, RelativePattern])(t).UnmarshalJSON(x)
}
func (t *GlobPattern) MarshalJSON() ([]byte, error) {
	return (*Or2[Pattern, RelativePattern])(t).MarshalJSON()
}
// Inline value information can be provided by different means:
// - directly as a text value (class InlineValueText).
// - as a name to use for a variable lookup (class InlineValueVariableLookup)
// - as an evaluatable expression (class InlineValueEvaluatableExpression)
// The InlineValue types combines all inline value types into one type.
// 
// @since 3.17.0
type InlineValue Or3[InlineValueText, InlineValueVariableLookup, InlineValueEvaluatableExpression]

func (t *InlineValue) UnmarshalJSON(x []byte) error {
	return (*Or3[InlineValueText, InlineValueVariableLookup, InlineValueEvaluatableExpression])(t).UnmarshalJSON(x)
}
func (t *InlineValue) MarshalJSON() ([]byte, error) {
	return (*Or3[InlineValueText, InlineValueVariableLookup, InlineValueEvaluatableExpression])(t).MarshalJSON()
}
// The LSP any type.
// Please note that strictly speaking a property with the value `undefined`
// can't be converted into JSON preserving the property name. However for
// convenience it is allowed and assumed that all these properties are
// optional as well.
// @since 3.17.0
type LSPAny any

// LSP arrays.
// @since 3.17.0
type LSPArray []any

// LSP object definition.
// @since 3.17.0
type LSPObject map[string]any

// MarkedString can be used to render human readable text. It is either a markdown string
// or a code-block that provides a language and a code snippet. The language identifier
// is semantically equal to the optional language identifier in fenced code blocks in GitHub
// issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
// 
// The pair of a language and a value is an equivalent to markdown:
// ```${language}
// ${value}
// ```
// 
// Note that markdown strings will be sanitized - that means html will be escaped.
// @deprecated use MarkupContent instead.
type MarkedString Or2[string, MarkedStringWithLanguage]

func (t *MarkedString) UnmarshalJSON(x []byte) error {
	return (*Or2[string, MarkedStringWithLanguage])(t).UnmarshalJSON(x)
}
func (t *MarkedString) MarshalJSON() ([]byte, error) {
	return (*Or2[string, MarkedStringWithLanguage])(t).MarshalJSON()
}
// A notebook document filter denotes a notebook document by
// different properties. The properties will be match
// against the notebook's URI (same as with documents)
// 
// @since 3.17.0
type NotebookDocumentFilter Or3[NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern]

func (t *NotebookDocumentFilter) UnmarshalJSON(x []byte) error {
	return (*Or3[NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern])(t).UnmarshalJSON(x)
}
func (t *NotebookDocumentFilter) MarshalJSON() ([]byte, error) {
	return (*Or3[NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern])(t).MarshalJSON()
}
// The glob pattern to watch relative to the base path. Glob patterns can have the following syntax:
// - `*` to match one or more characters in a path segment
// - `?` to match on one character in a path segment
// - `**` to match any number of path segments, including none
// - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
// 
// @since 3.17.0
type Pattern string


type PrepareRenameResult Or3[Range, PrepareRenamePlaceholder, PrepareRenameDefaultBehavior]

func (t *PrepareRenameResult) UnmarshalJSON(x []byte) error {
	return (*Or3[Range, PrepareRenamePlaceholder, PrepareRenameDefaultBehavior])(t).UnmarshalJSON(x)
}
func (t *PrepareRenameResult) MarshalJSON() ([]byte, error) {
	return (*Or3[Range, PrepareRenamePlaceholder, PrepareRenameDefaultBehavior])(t).MarshalJSON()
}

type ProgressToken Or2[int32, string]

func (t *ProgressToken) UnmarshalJSON(x []byte) error {
	return (*Or2[int32, string])(t).UnmarshalJSON(x)
}
func (t *ProgressToken) MarshalJSON() ([]byte, error) {
	return (*Or2[int32, string])(t).MarshalJSON()
}

type RegularExpressionEngineKind string

// An event describing a change to a text document. If only a text is provided
// it is considered to be the full content of the document.
type TextDocumentContentChangeEvent Or2[TextDocumentContentChangePartial, TextDocumentContentChangeWholeDocument]

func (t *TextDocumentContentChangeEvent) UnmarshalJSON(x []byte) error {
	return (*Or2[TextDocumentContentChangePartial, TextDocumentContentChangeWholeDocument])(t).UnmarshalJSON(x)
}
func (t *TextDocumentContentChangeEvent) MarshalJSON() ([]byte, error) {
	return (*Or2[TextDocumentContentChangePartial, TextDocumentContentChangeWholeDocument])(t).MarshalJSON()
}
// A document filter denotes a document by different properties like
// the {@link TextDocument.languageId language}, the {@link Uri.scheme scheme} of
// its resource, or a glob-pattern that is applied to the {@link TextDocument.fileName path}.
// 
// Glob patterns can have the following syntax:
// - `*` to match one or more characters in a path segment
// - `?` to match on one character in a path segment
// - `**` to match any number of path segments, including none
// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
// 
// @sample A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }`
// @sample A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
// 
// @since 3.17.0
type TextDocumentFilter Or3[TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern]

func (t *TextDocumentFilter) UnmarshalJSON(x []byte) error {
	return (*Or3[TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern])(t).UnmarshalJSON(x)
}
func (t *TextDocumentFilter) MarshalJSON() ([]byte, error) {
	return (*Or3[TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern])(t).MarshalJSON()
}
// A workspace diagnostic document report.
// 
// @since 3.17.0
type WorkspaceDocumentDiagnosticReport Or2[WorkspaceFullDocumentDiagnosticReport, WorkspaceUnchangedDocumentDiagnosticReport]

func (t *WorkspaceDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	return (*Or2[WorkspaceFullDocumentDiagnosticReport, WorkspaceUnchangedDocumentDiagnosticReport])(t).UnmarshalJSON(x)
}
func (t *WorkspaceDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	return (*Or2[WorkspaceFullDocumentDiagnosticReport, WorkspaceUnchangedDocumentDiagnosticReport])(t).MarshalJSON()
}
type UnmarshalError struct {
   msg string
}
func (e UnmarshalError) Error() string {
   return e.msg
}
type Tuple[T1, T2 any] struct {
   V1 T1
   V2 T2
}
func (t Tuple[T1, T2]) MarshalJSON() ([]byte, error) {
   return json.Marshal([2]any{t.V1, t.V2})
}
func (t *Tuple[T1, T2]) UnmarshalJSON(data []byte) error {
   var arr [2]json.RawMessage
   if err := json.Unmarshal(data, &arr); err != nil {
       return err
   }
   var t1 T1
   var t2 T2
   if err := json.Unmarshal(arr[0], &t1); err != nil {
       return err
   }
   if err := json.Unmarshal(arr[1], &t2); err != nil {
       return err
   }
   t.V1 = t1
   t.V2 = t2
   return nil
}
type ResponseError struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}
type MethodKind string
type Message interface {
	isMessage()
}
type IncomingMessage interface {
	GetMethod() MethodKind
	GetParams() any
}
type Request interface {
	isRequest()
}
type Notification interface {
	isNotification()
}
type Response interface {
	isResponse()
}


// A special text edit with an additional change annotation.
// 
// @since 3.16.0.
type AnnotatedTextEdit struct {
	// The actual identifier of the change annotation
	AnnotationId ChangeAnnotationIdentifier `json:"annotationId"`
	// The string to be inserted. For delete operations use an
	// empty string.
	NewText string `json:"newText"`
	// The range of the text document to be manipulated. To insert
	// text into a document create a range where start === end.
	Range Range `json:"range"`
}
func (t *AnnotatedTextEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["annotationId"]; !exists {
		return fmt.Errorf("missing required field: annotationId")
	}
	if _, exists := m["newText"]; !exists {
		return fmt.Errorf("missing required field: newText")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias AnnotatedTextEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = AnnotatedTextEdit(test)
	
	return nil
}
// The parameters passed via an apply workspace edit request.
type ApplyWorkspaceEditParams struct {
	// The edits to apply.
	Edit WorkspaceEdit `json:"edit"`
	// An optional label of the workspace edit. This label is
	// presented in the user interface for example on an undo
	// stack to undo the workspace edit.
	Label string `json:"label,omitempty"`
	// Additional data about the edit.
	// 
	// @since 3.18.0
	// @proposed
	Metadata *WorkspaceEditMetadata `json:"metadata,omitzero"`
}
func (t *ApplyWorkspaceEditParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["edit"]; !exists {
		return fmt.Errorf("missing required field: edit")
	}
	type Alias ApplyWorkspaceEditParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ApplyWorkspaceEditParams(test)
	
	return nil
}
// The result returned from the apply workspace edit request.
// 
// @since 3.17 renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {
	// Indicates whether the edit was applied or not.
	Applied bool `json:"applied"`
	// Depending on the client's failure handling strategy `failedChange` might
	// contain the index of the change that failed. This property is only available
	// if the client signals a `failureHandlingStrategy` in its client capabilities.
	FailedChange uint32 `json:"failedChange,omitempty"`
	// An optional textual description for why the edit was not applied.
	// This may be used by the server for diagnostic logging or to provide
	// a suitable error for a request that triggered the edit.
	FailureReason string `json:"failureReason,omitempty"`
}
func (t *ApplyWorkspaceEditResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["applied"]; !exists {
		return fmt.Errorf("missing required field: applied")
	}
	type Alias ApplyWorkspaceEditResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ApplyWorkspaceEditResult(test)
	
	return nil
}
// A base for all symbol information.
type BaseSymbolInformation struct {
	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName string `json:"containerName,omitempty"`
	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`
	// The name of this symbol.
	Name string `json:"name"`
	// Tags for this symbol.
	// 
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitzero"`
}
func (t *BaseSymbolInformation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	type Alias BaseSymbolInformation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = BaseSymbolInformation(test)
	
	return nil
}
// @since 3.16.0
type CallHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *CallHierarchyClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CallHierarchyClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyClientCapabilities(test)
	
	return nil
}
// Represents an incoming call, e.g. a caller of a method or constructor.
// 
// @since 3.16.0
type CallHierarchyIncomingCall struct {
	// The item that makes the call.
	From CallHierarchyItem `json:"from"`
	// The ranges at which the calls appear. This is relative to the caller
	// denoted by {@link CallHierarchyIncomingCall.from `this.from`}.
	FromRanges []Range `json:"fromRanges"`
}
func (t *CallHierarchyIncomingCall) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["from"]; !exists {
		return fmt.Errorf("missing required field: from")
	}
	if _, exists := m["fromRanges"]; !exists {
		return fmt.Errorf("missing required field: fromRanges")
	}
	type Alias CallHierarchyIncomingCall
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyIncomingCall(test)
	
	return nil
}
// The parameter of a `callHierarchy/incomingCalls` request.
// 
// @since 3.16.0
type CallHierarchyIncomingCallsParams struct {

	Item CallHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *CallHierarchyIncomingCallsParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["item"]; !exists {
		return fmt.Errorf("missing required field: item")
	}
	type Alias CallHierarchyIncomingCallsParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyIncomingCallsParams(test)
	
	return nil
}
// Represents programming constructs like functions or constructors in the context
// of call hierarchy.
// 
// @since 3.16.0
type CallHierarchyItem struct {
	// A data entry field that is preserved between a call hierarchy prepare and
	// incoming calls or outgoing calls requests.
	Data any `json:"data,omitempty"`
	// More detail for this item, e.g. the signature of a function.
	Detail string `json:"detail,omitempty"`
	// The kind of this item.
	Kind SymbolKind `json:"kind"`
	// The name of this item.
	Name string `json:"name"`
	// The range enclosing this symbol not including leading/trailing whitespace but everything else, e.g. comments and code.
	Range Range `json:"range"`
	// The range that should be selected and revealed when this symbol is being picked, e.g. the name of a function.
	// Must be contained by the {@link CallHierarchyItem.range `range`}.
	SelectionRange Range `json:"selectionRange"`
	// Tags for this item.
	Tags []SymbolTag `json:"tags,omitzero"`
	// The resource identifier of this item.
	Uri DocumentUri `json:"uri"`
}
func (t *CallHierarchyItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["selectionRange"]; !exists {
		return fmt.Errorf("missing required field: selectionRange")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias CallHierarchyItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyItem(test)
	
	return nil
}
// Call hierarchy options used during static registration.
// 
// @since 3.16.0
type CallHierarchyOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CallHierarchyOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CallHierarchyOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyOptions(test)
	
	return nil
}
// Represents an outgoing call, e.g. calling a getter from a method or a method from a constructor etc.
// 
// @since 3.16.0
type CallHierarchyOutgoingCall struct {
	// The range at which this item is called. This is the range relative to the caller, e.g the item
	// passed to {@link CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls `provideCallHierarchyOutgoingCalls`}
	// and not {@link CallHierarchyOutgoingCall.to `this.to`}.
	FromRanges []Range `json:"fromRanges"`
	// The item that is called.
	To CallHierarchyItem `json:"to"`
}
func (t *CallHierarchyOutgoingCall) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["fromRanges"]; !exists {
		return fmt.Errorf("missing required field: fromRanges")
	}
	if _, exists := m["to"]; !exists {
		return fmt.Errorf("missing required field: to")
	}
	type Alias CallHierarchyOutgoingCall
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyOutgoingCall(test)
	
	return nil
}
// The parameter of a `callHierarchy/outgoingCalls` request.
// 
// @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {

	Item CallHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *CallHierarchyOutgoingCallsParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["item"]; !exists {
		return fmt.Errorf("missing required field: item")
	}
	type Alias CallHierarchyOutgoingCallsParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyOutgoingCallsParams(test)
	
	return nil
}
// The parameter of a `textDocument/prepareCallHierarchy` request.
// 
// @since 3.16.0
type CallHierarchyPrepareParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *CallHierarchyPrepareParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias CallHierarchyPrepareParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyPrepareParams(test)
	
	return nil
}
// Call hierarchy options used during static or dynamic registration.
// 
// @since 3.16.0
type CallHierarchyRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CallHierarchyRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias CallHierarchyRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CallHierarchyRegistrationOptions(test)
	
	return nil
}

type CancelParams struct {
	// The request id to cancel.
	Id Or2[int32, string] `json:"id"`
}
func (t *CancelParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["id"]; !exists {
		return fmt.Errorf("missing required field: id")
	}
	type Alias CancelParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CancelParams(test)
	
	return nil
}
// Additional information that describes document changes.
// 
// @since 3.16.0
type ChangeAnnotation struct {
	// A human-readable string which is rendered less prominent in
	// the user interface.
	Description string `json:"description,omitempty"`
	// A human-readable string describing the actual change. The string
	// is rendered prominent in the user interface.
	Label string `json:"label"`
	// A flag which indicates that user confirmation is needed
	// before applying the change.
	NeedsConfirmation bool `json:"needsConfirmation,omitempty"`
}
func (t *ChangeAnnotation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["label"]; !exists {
		return fmt.Errorf("missing required field: label")
	}
	type Alias ChangeAnnotation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ChangeAnnotation(test)
	
	return nil
}
// @since 3.18.0
type ChangeAnnotationsSupportOptions struct {
	// Whether the client groups edits with equal labels into tree nodes,
	// for instance all edits labelled with "Changes in Strings" would
	// be a tree node.
	GroupsOnLabel bool `json:"groupsOnLabel,omitempty"`
}
func (t *ChangeAnnotationsSupportOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ChangeAnnotationsSupportOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ChangeAnnotationsSupportOptions(test)
	
	return nil
}
// Defines the capabilities provided by the client.
type ClientCapabilities struct {
	// Experimental client capabilities.
	Experimental any `json:"experimental,omitempty"`
	// General client capabilities.
	// 
	// @since 3.16.0
	General *GeneralClientCapabilities `json:"general,omitzero"`
	// Capabilities specific to the notebook document support.
	// 
	// @since 3.17.0
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitzero"`
	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitzero"`
	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitzero"`
	// Workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitzero"`
}
func (t *ClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCapabilities(test)
	
	return nil
}
// @since 3.18.0
type ClientCodeActionKindOptions struct {
	// The code action kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []CodeActionKind `json:"valueSet"`
}
func (t *ClientCodeActionKindOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["valueSet"]; !exists {
		return fmt.Errorf("missing required field: valueSet")
	}
	type Alias ClientCodeActionKindOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCodeActionKindOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientCodeActionLiteralOptions struct {
	// The code action kind is support with the following value
	// set.
	CodeActionKind ClientCodeActionKindOptions `json:"codeActionKind"`
}
func (t *ClientCodeActionLiteralOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["codeActionKind"]; !exists {
		return fmt.Errorf("missing required field: codeActionKind")
	}
	type Alias ClientCodeActionLiteralOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCodeActionLiteralOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientCodeActionResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
func (t *ClientCodeActionResolveOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["properties"]; !exists {
		return fmt.Errorf("missing required field: properties")
	}
	type Alias ClientCodeActionResolveOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCodeActionResolveOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientCodeLensResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
func (t *ClientCodeLensResolveOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["properties"]; !exists {
		return fmt.Errorf("missing required field: properties")
	}
	type Alias ClientCodeLensResolveOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCodeLensResolveOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientCompletionItemInsertTextModeOptions struct {

	ValueSet []InsertTextMode `json:"valueSet"`
}
func (t *ClientCompletionItemInsertTextModeOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["valueSet"]; !exists {
		return fmt.Errorf("missing required field: valueSet")
	}
	type Alias ClientCompletionItemInsertTextModeOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCompletionItemInsertTextModeOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientCompletionItemOptions struct {
	// Client supports commit characters on a completion item.
	CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`
	// Client supports the deprecated property on a completion item.
	DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`
	// Client supports the following content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitzero"`
	// Client support insert replace edit to control different behavior if a
	// completion item is inserted in the text or should replace text.
	// 
	// @since 3.16.0
	InsertReplaceSupport bool `json:"insertReplaceSupport,omitempty"`
	// The client supports the `insertTextMode` property on
	// a completion item to override the whitespace handling mode
	// as defined by the client (see `insertTextMode`).
	// 
	// @since 3.16.0
	InsertTextModeSupport *ClientCompletionItemInsertTextModeOptions `json:"insertTextModeSupport,omitzero"`
	// The client has support for completion item label
	// details (see also `CompletionItemLabelDetails`).
	// 
	// @since 3.17.0
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
	// Client supports the preselect property on a completion item.
	PreselectSupport bool `json:"preselectSupport,omitempty"`
	// Indicates which properties a client can resolve lazily on a completion
	// item. Before version 3.16.0 only the predefined properties `documentation`
	// and `details` could be resolved lazily.
	// 
	// @since 3.16.0
	ResolveSupport *ClientCompletionItemResolveOptions `json:"resolveSupport,omitzero"`
	// Client supports snippets as insert text.
	// 
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	SnippetSupport bool `json:"snippetSupport,omitempty"`
	// Client supports the tag property on a completion item. Clients supporting
	// tags have to handle unknown tags gracefully. Clients especially need to
	// preserve unknown tags when sending a completion item back to the server in
	// a resolve call.
	// 
	// @since 3.15.0
	TagSupport *CompletionItemTagOptions `json:"tagSupport,omitzero"`
}
func (t *ClientCompletionItemOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientCompletionItemOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCompletionItemOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientCompletionItemOptionsKind struct {
	// The completion item kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	// 
	// If this property is not present the client only supports
	// the completion items kinds from `Text` to `Reference` as defined in
	// the initial version of the protocol.
	ValueSet []CompletionItemKind `json:"valueSet,omitzero"`
}
func (t *ClientCompletionItemOptionsKind) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientCompletionItemOptionsKind
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCompletionItemOptionsKind(test)
	
	return nil
}
// @since 3.18.0
type ClientCompletionItemResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
func (t *ClientCompletionItemResolveOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["properties"]; !exists {
		return fmt.Errorf("missing required field: properties")
	}
	type Alias ClientCompletionItemResolveOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientCompletionItemResolveOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientDiagnosticsTagOptions struct {
	// The tags supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet"`
}
func (t *ClientDiagnosticsTagOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["valueSet"]; !exists {
		return fmt.Errorf("missing required field: valueSet")
	}
	type Alias ClientDiagnosticsTagOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientDiagnosticsTagOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientFoldingRangeKindOptions struct {
	// The folding range kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []FoldingRangeKind `json:"valueSet,omitzero"`
}
func (t *ClientFoldingRangeKindOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientFoldingRangeKindOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientFoldingRangeKindOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientFoldingRangeOptions struct {
	// If set, the client signals that it supports setting collapsedText on
	// folding ranges to display custom labels instead of the default text.
	// 
	// @since 3.17.0
	CollapsedText bool `json:"collapsedText,omitempty"`
}
func (t *ClientFoldingRangeOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientFoldingRangeOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientFoldingRangeOptions(test)
	
	return nil
}
// Information about the client
// 
// @since 3.15.0
// @since 3.18.0 ClientInfo type name added.
type ClientInfo struct {
	// The name of the client as defined by the client.
	Name string `json:"name"`
	// The client's version as defined by the client.
	Version string `json:"version,omitempty"`
}
func (t *ClientInfo) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	type Alias ClientInfo
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientInfo(test)
	
	return nil
}
// @since 3.18.0
type ClientInlayHintResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
func (t *ClientInlayHintResolveOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["properties"]; !exists {
		return fmt.Errorf("missing required field: properties")
	}
	type Alias ClientInlayHintResolveOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientInlayHintResolveOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientSemanticTokensRequestFullDelta struct {
	// The client will send the `textDocument/semanticTokens/full/delta` request if
	// the server provides a corresponding handler.
	Delta bool `json:"delta,omitempty"`
}
func (t *ClientSemanticTokensRequestFullDelta) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientSemanticTokensRequestFullDelta
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSemanticTokensRequestFullDelta(test)
	
	return nil
}
// @since 3.18.0
type ClientSemanticTokensRequestOptions struct {
	// The client will send the `textDocument/semanticTokens/full` request if
	// the server provides a corresponding handler.
	Full *Or2[bool, ClientSemanticTokensRequestFullDelta] `json:"full,omitzero"`
	// The client will send the `textDocument/semanticTokens/range` request if
	// the server provides a corresponding handler.
	Range *Or2[bool, LSPObject] `json:"range,omitzero"`
}
func (t *ClientSemanticTokensRequestOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientSemanticTokensRequestOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSemanticTokensRequestOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientShowMessageActionItemOptions struct {
	// Whether the client supports additional attributes which
	// are preserved and send back to the server in the
	// request's response.
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
}
func (t *ClientShowMessageActionItemOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientShowMessageActionItemOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientShowMessageActionItemOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientSignatureInformationOptions struct {
	// The client supports the `activeParameter` property on `SignatureInformation`
	// literal.
	// 
	// @since 3.16.0
	ActiveParameterSupport bool `json:"activeParameterSupport,omitempty"`
	// Client supports the following content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitzero"`
	// The client supports the `activeParameter` property on
	// `SignatureHelp`/`SignatureInformation` being set to `null` to
	// indicate that no parameter should be active.
	// 
	// @since 3.18.0
	// @proposed
	NoActiveParameterSupport bool `json:"noActiveParameterSupport,omitempty"`
	// Client capabilities specific to parameter information.
	ParameterInformation *ClientSignatureParameterInformationOptions `json:"parameterInformation,omitzero"`
}
func (t *ClientSignatureInformationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientSignatureInformationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSignatureInformationOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientSignatureParameterInformationOptions struct {
	// The client supports processing label offsets instead of a
	// simple label string.
	// 
	// @since 3.14.0
	LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
}
func (t *ClientSignatureParameterInformationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientSignatureParameterInformationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSignatureParameterInformationOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientSymbolKindOptions struct {
	// The symbol kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	// 
	// If this property is not present the client only supports
	// the symbol kinds from `File` to `Array` as defined in
	// the initial version of the protocol.
	ValueSet []SymbolKind `json:"valueSet,omitzero"`
}
func (t *ClientSymbolKindOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ClientSymbolKindOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSymbolKindOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientSymbolResolveOptions struct {
	// The properties that a client can resolve lazily. Usually
	// `location.range`
	Properties []string `json:"properties"`
}
func (t *ClientSymbolResolveOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["properties"]; !exists {
		return fmt.Errorf("missing required field: properties")
	}
	type Alias ClientSymbolResolveOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSymbolResolveOptions(test)
	
	return nil
}
// @since 3.18.0
type ClientSymbolTagOptions struct {
	// The tags supported by the client.
	ValueSet []SymbolTag `json:"valueSet"`
}
func (t *ClientSymbolTagOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["valueSet"]; !exists {
		return fmt.Errorf("missing required field: valueSet")
	}
	type Alias ClientSymbolTagOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ClientSymbolTagOptions(test)
	
	return nil
}
// A code action represents a change that can be performed in code, e.g. to fix a problem or
// to refactor code.
// 
// A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {
	// A command this code action executes. If a code action
	// provides an edit and a command, first the edit is
	// executed and then the command.
	Command *Command `json:"command,omitzero"`
	// A data entry field that is preserved on a code action between
	// a `textDocument/codeAction` and a `codeAction/resolve` request.
	// 
	// @since 3.16.0
	Data any `json:"data,omitempty"`
	// The diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitzero"`
	// Marks that the code action cannot currently be applied.
	// 
	// Clients should follow the following guidelines regarding disabled code actions:
	// 
	//   - Disabled code actions are not shown in automatic [lightbulbs](https://code.visualstudio.com/docs/editor/editingevolved#_code-action)
	//     code action menus.
	// 
	//   - Disabled actions are shown as faded out in the code action menu when the user requests a more specific type
	//     of code action, such as refactorings.
	// 
	//   - If the user has a [keybinding](https://code.visualstudio.com/docs/editor/refactoring#_keybindings-for-code-actions)
	//     that auto applies a code action and only disabled code actions are returned, the client should show the user an
	//     error message with `reason` in the editor.
	// 
	// @since 3.16.0
	Disabled *CodeActionDisabled `json:"disabled,omitzero"`
	// The workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitzero"`
	// Marks this as a preferred action. Preferred actions are used by the `auto fix` command and can be targeted
	// by keybindings.
	// 
	// A quick fix should be marked preferred if it properly addresses the underlying error.
	// A refactoring should be marked preferred if it is the most reasonable choice of actions to take.
	// 
	// @since 3.15.0
	IsPreferred bool `json:"isPreferred,omitempty"`
	// The kind of the code action.
	// 
	// Used to filter code actions.
	Kind *CodeActionKind `json:"kind,omitzero"`
	// Tags for this code action.
	// 
	// @since 3.18.0 - proposed
	Tags []CodeActionTag `json:"tags,omitzero"`
	// A short, human-readable, title for this code action.
	Title string `json:"title"`
}
func (t *CodeAction) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["title"]; !exists {
		return fmt.Errorf("missing required field: title")
	}
	type Alias CodeAction
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeAction(test)
	
	return nil
}
// The Client Capabilities of a {@link CodeActionRequest}.
type CodeActionClientCapabilities struct {
	// The client support code action literals of type `CodeAction` as a valid
	// response of the `textDocument/codeAction` request. If the property is not
	// set the request can only return `Command` literals.
	// 
	// @since 3.8.0
	CodeActionLiteralSupport *ClientCodeActionLiteralOptions `json:"codeActionLiteralSupport,omitzero"`
	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/codeAction` and a
	// `codeAction/resolve` request.
	// 
	// @since 3.16.0
	DataSupport bool `json:"dataSupport,omitempty"`
	// Whether code action supports the `disabled` property.
	// 
	// @since 3.16.0
	DisabledSupport bool `json:"disabledSupport,omitempty"`
	// Whether the client supports documentation for a class of
	// code actions.
	// 
	// @since 3.18.0
	// @proposed
	DocumentationSupport bool `json:"documentationSupport,omitempty"`
	// Whether code action supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// `CodeAction#edit` property by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	// 
	// @since 3.16.0
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
	// Whether code action supports the `isPreferred` property.
	// 
	// @since 3.15.0
	IsPreferredSupport bool `json:"isPreferredSupport,omitempty"`
	// Whether the client supports resolving additional code action
	// properties via a separate `codeAction/resolve` request.
	// 
	// @since 3.16.0
	ResolveSupport *ClientCodeActionResolveOptions `json:"resolveSupport,omitzero"`
	// Client supports the tag property on a code action. Clients
	// supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.18.0 - proposed
	TagSupport *CodeActionTagOptions `json:"tagSupport,omitzero"`
}
func (t *CodeActionClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CodeActionClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionClientCapabilities(test)
	
	return nil
}
// Contains additional diagnostic information about the context in which
// a {@link CodeActionProvider.provideCodeActions code action} is run.
type CodeActionContext struct {
	// An array of diagnostics known on the client side overlapping the range provided to the
	// `textDocument/codeAction` request. They are provided so that the server knows which
	// errors are currently presented to the user for the given range. There is no guarantee
	// that these accurately reflect the error state of the resource. The primary parameter
	// to compute code actions is the provided range.
	Diagnostics []Diagnostic `json:"diagnostics"`
	// Requested kind of actions to return.
	// 
	// Actions not of this kind are filtered out by the client before being shown. So servers
	// can omit computing them.
	Only []CodeActionKind `json:"only,omitzero"`
	// The reason why code actions were requested.
	// 
	// @since 3.17.0
	TriggerKind *CodeActionTriggerKind `json:"triggerKind,omitzero"`
}
func (t *CodeActionContext) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["diagnostics"]; !exists {
		return fmt.Errorf("missing required field: diagnostics")
	}
	type Alias CodeActionContext
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionContext(test)
	
	return nil
}
// Captures why the code action is currently disabled.
// 
// @since 3.18.0
type CodeActionDisabled struct {
	// Human readable description of why the code action is currently disabled.
	// 
	// This is displayed in the code actions UI.
	Reason string `json:"reason"`
}
func (t *CodeActionDisabled) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["reason"]; !exists {
		return fmt.Errorf("missing required field: reason")
	}
	type Alias CodeActionDisabled
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionDisabled(test)
	
	return nil
}
// Documentation for a class of code actions.
// 
// @since 3.18.0
// @proposed
type CodeActionKindDocumentation struct {
	// Command that is ued to display the documentation to the user.
	// 
	// The title of this documentation code action is taken from {@linkcode Command.title}
	Command Command `json:"command"`
	// The kind of the code action being documented.
	// 
	// If the kind is generic, such as `CodeActionKind.Refactor`, the documentation will be shown whenever any
	// refactorings are returned. If the kind if more specific, such as `CodeActionKind.RefactorExtract`, the
	// documentation will only be shown when extract refactoring code actions are returned.
	Kind CodeActionKind `json:"kind"`
}
func (t *CodeActionKindDocumentation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["command"]; !exists {
		return fmt.Errorf("missing required field: command")
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias CodeActionKindDocumentation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionKindDocumentation(test)
	
	return nil
}
// Provider options for a {@link CodeActionRequest}.
type CodeActionOptions struct {
	// CodeActionKinds that this server may return.
	// 
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitzero"`
	// Static documentation for a class of code actions.
	// 
	// Documentation from the provider should be shown in the code actions menu if either:
	// 
	// - Code actions of `kind` are requested by the editor. In this case, the editor will show the documentation that
	//   most closely matches the requested code action kind. For example, if a provider has documentation for
	//   both `Refactor` and `RefactorExtract`, when the user requests code actions for `RefactorExtract`,
	//   the editor will use the documentation for `RefactorExtract` instead of the documentation for `Refactor`.
	// 
	// - Any code actions of `kind` are returned by the provider.
	// 
	// At most one documentation entry should be shown per provider.
	// 
	// @since 3.18.0
	// @proposed
	Documentation []CodeActionKindDocumentation `json:"documentation,omitzero"`
	// The server provides support to resolve additional
	// information for a code action.
	// 
	// @since 3.16.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CodeActionOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CodeActionOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionOptions(test)
	
	return nil
}
// The parameters of a {@link CodeActionRequest}.
type CodeActionParams struct {
	// Context carrying additional information.
	Context CodeActionContext `json:"context"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The range for which the command was invoked.
	Range Range `json:"range"`
	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *CodeActionParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["context"]; !exists {
		return fmt.Errorf("missing required field: context")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias CodeActionParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionParams(test)
	
	return nil
}
// Registration options for a {@link CodeActionRequest}.
type CodeActionRegistrationOptions struct {
	// CodeActionKinds that this server may return.
	// 
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitzero"`
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Static documentation for a class of code actions.
	// 
	// Documentation from the provider should be shown in the code actions menu if either:
	// 
	// - Code actions of `kind` are requested by the editor. In this case, the editor will show the documentation that
	//   most closely matches the requested code action kind. For example, if a provider has documentation for
	//   both `Refactor` and `RefactorExtract`, when the user requests code actions for `RefactorExtract`,
	//   the editor will use the documentation for `RefactorExtract` instead of the documentation for `Refactor`.
	// 
	// - Any code actions of `kind` are returned by the provider.
	// 
	// At most one documentation entry should be shown per provider.
	// 
	// @since 3.18.0
	// @proposed
	Documentation []CodeActionKindDocumentation `json:"documentation,omitzero"`
	// The server provides support to resolve additional
	// information for a code action.
	// 
	// @since 3.16.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CodeActionRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias CodeActionRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionRegistrationOptions(test)
	
	return nil
}
// @since 3.18.0 - proposed
type CodeActionTagOptions struct {
	// The tags supported by the client.
	ValueSet []CodeActionTag `json:"valueSet"`
}
func (t *CodeActionTagOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["valueSet"]; !exists {
		return fmt.Errorf("missing required field: valueSet")
	}
	type Alias CodeActionTagOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeActionTagOptions(test)
	
	return nil
}
// Structure to capture a description for an error code.
// 
// @since 3.16.0
type CodeDescription struct {
	// An URI to open with more information about the diagnostic error.
	Href URI `json:"href"`
}
func (t *CodeDescription) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["href"]; !exists {
		return fmt.Errorf("missing required field: href")
	}
	type Alias CodeDescription
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeDescription(test)
	
	return nil
}
// A code lens represents a {@link Command command} that should be shown along with
// source text, like the number of references, a way to run tests, etc.
// 
// A code lens is _unresolved_ when no command is associated to it. For performance
// reasons the creation of a code lens and resolving should be done in two stages.
type CodeLens struct {
	// The command this code lens represents.
	Command *Command `json:"command,omitzero"`
	// A data entry field that is preserved on a code lens item between
	// a {@link CodeLensRequest} and a {@link CodeLensResolveRequest}
	Data any `json:"data,omitempty"`
	// The range in which this code lens is valid. Should only span a single line.
	Range Range `json:"range"`
}
func (t *CodeLens) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias CodeLens
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeLens(test)
	
	return nil
}
// The client capabilities  of a {@link CodeLensRequest}.
type CodeLensClientCapabilities struct {
	// Whether code lens supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client supports resolving additional code lens
	// properties via a separate `codeLens/resolve` request.
	// 
	// @since 3.18.0
	ResolveSupport *ClientCodeLensResolveOptions `json:"resolveSupport,omitzero"`
}
func (t *CodeLensClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CodeLensClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeLensClientCapabilities(test)
	
	return nil
}
// Code Lens provider options of a {@link CodeLensRequest}.
type CodeLensOptions struct {
	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CodeLensOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CodeLensOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeLensOptions(test)
	
	return nil
}
// The parameters of a {@link CodeLensRequest}.
type CodeLensParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *CodeLensParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias CodeLensParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeLensParams(test)
	
	return nil
}
// Registration options for a {@link CodeLensRequest}.
type CodeLensRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CodeLensRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias CodeLensRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeLensRegistrationOptions(test)
	
	return nil
}
// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	// 
	// Note that this event is global and will force the client to refresh all
	// code lenses currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}
func (t *CodeLensWorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CodeLensWorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CodeLensWorkspaceClientCapabilities(test)
	
	return nil
}
// Represents a color in RGBA space.
type Color struct {
	// The alpha component of this color in the range [0-1].
	Alpha float64 `json:"alpha"`
	// The blue component of this color in the range [0-1].
	Blue float64 `json:"blue"`
	// The green component of this color in the range [0-1].
	Green float64 `json:"green"`
	// The red component of this color in the range [0-1].
	Red float64 `json:"red"`
}
func (t *Color) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["alpha"]; !exists {
		return fmt.Errorf("missing required field: alpha")
	}
	if _, exists := m["blue"]; !exists {
		return fmt.Errorf("missing required field: blue")
	}
	if _, exists := m["green"]; !exists {
		return fmt.Errorf("missing required field: green")
	}
	if _, exists := m["red"]; !exists {
		return fmt.Errorf("missing required field: red")
	}
	type Alias Color
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Color(test)
	
	return nil
}
// Represents a color range from a document.
type ColorInformation struct {
	// The actual color value for this color range.
	Color Color `json:"color"`
	// The range in the document where this color appears.
	Range Range `json:"range"`
}
func (t *ColorInformation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["color"]; !exists {
		return fmt.Errorf("missing required field: color")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias ColorInformation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ColorInformation(test)
	
	return nil
}

type ColorPresentation struct {
	// An optional array of additional {@link TextEdit text edits} that are applied when
	// selecting this color presentation. Edits must not overlap with the main {@link ColorPresentation.textEdit edit} nor with themselves.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitzero"`
	// The label of this color presentation. It will be shown on the color
	// picker header. By default this is also the text that is inserted when selecting
	// this color presentation.
	Label string `json:"label"`
	// An {@link TextEdit edit} which is applied to a document when selecting
	// this presentation for the color.  When `falsy` the {@link ColorPresentation.label label}
	// is used.
	TextEdit *TextEdit `json:"textEdit,omitzero"`
}
func (t *ColorPresentation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["label"]; !exists {
		return fmt.Errorf("missing required field: label")
	}
	type Alias ColorPresentation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ColorPresentation(test)
	
	return nil
}
// Parameters for a {@link ColorPresentationRequest}.
type ColorPresentationParams struct {
	// The color to request presentations for.
	Color Color `json:"color"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *ColorPresentationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["color"]; !exists {
		return fmt.Errorf("missing required field: color")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias ColorPresentationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ColorPresentationParams(test)
	
	return nil
}
// Represents a reference to a command. Provides a title which
// will be used to represent a command in the UI and, optionally,
// an array of arguments which will be passed to the command handler
// function when invoked.
type Command struct {
	// Arguments that the command handler should be
	// invoked with.
	Arguments []any `json:"arguments,omitzero"`
	// The identifier of the actual command handler.
	Command string `json:"command"`
	// Title of the command, like `save`.
	Title string `json:"title"`
	// An optional tooltip.
	// 
	// @since 3.18.0
	// @proposed
	Tooltip string `json:"tooltip,omitempty"`
}
func (t *Command) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["command"]; !exists {
		return fmt.Errorf("missing required field: command")
	}
	if _, exists := m["title"]; !exists {
		return fmt.Errorf("missing required field: title")
	}
	type Alias Command
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Command(test)
	
	return nil
}
// Completion client capabilities
type CompletionClientCapabilities struct {
	// The client supports the following `CompletionItem` specific
	// capabilities.
	CompletionItem *ClientCompletionItemOptions `json:"completionItem,omitzero"`

	CompletionItemKind *ClientCompletionItemOptionsKind `json:"completionItemKind,omitzero"`
	// The client supports the following `CompletionList` specific
	// capabilities.
	// 
	// @since 3.17.0
	CompletionList *CompletionListCapabilities `json:"completionList,omitzero"`
	// The client supports to send additional context information for a
	// `textDocument/completion` request.
	ContextSupport bool `json:"contextSupport,omitempty"`
	// Whether completion supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Defines how the client handles whitespace and indentation
	// when accepting a completion item that uses multi line
	// text in either `insertText` or `textEdit`.
	// 
	// @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitzero"`
}
func (t *CompletionClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CompletionClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionClientCapabilities(test)
	
	return nil
}
// Contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
	// The trigger character (a single character) that has trigger code complete.
	// Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
	// How the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`
}
func (t *CompletionContext) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["triggerKind"]; !exists {
		return fmt.Errorf("missing required field: triggerKind")
	}
	type Alias CompletionContext
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionContext(test)
	
	return nil
}
// A completion item represents a text snippet that is
// proposed to complete text that is being typed.
type CompletionItem struct {
	// An optional array of additional {@link TextEdit text edits} that are applied when
	// selecting this completion. Edits must not overlap (including the same insert position)
	// with the main {@link CompletionItem.textEdit edit} nor with themselves.
	// 
	// Additional text edits should be used to change text unrelated to the current cursor position
	// (for example adding an import statement at the top of the file if the completion item will
	// insert an unqualified type).
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitzero"`
	// An optional {@link Command command} that is executed *after* inserting this completion. *Note* that
	// additional modifications to the current document should be described with the
	// {@link CompletionItem.additionalTextEdits additionalTextEdits}-property.
	Command *Command `json:"command,omitzero"`
	// An optional set of characters that when pressed while this completion is active will accept it first and
	// then type that character. *Note* that all commit characters should have `length=1` and that superfluous
	// characters will be ignored.
	CommitCharacters []string `json:"commitCharacters,omitzero"`
	// A data entry field that is preserved on a completion item between a
	// {@link CompletionRequest} and a {@link CompletionResolveRequest}.
	Data any `json:"data,omitempty"`
	// Indicates if this item is deprecated.
	// @deprecated Use `tags` instead.
	Deprecated bool `json:"deprecated,omitempty"`
	// A human-readable string with additional information
	// about this item, like type or symbol information.
	Detail string `json:"detail,omitempty"`
	// A human-readable string that represents a doc-comment.
	Documentation *Or2[string, MarkupContent] `json:"documentation,omitzero"`
	// A string that should be used when filtering a set of
	// completion items. When `falsy` the {@link CompletionItem.label label}
	// is used.
	FilterText string `json:"filterText,omitempty"`
	// A string that should be inserted into a document when selecting
	// this completion. When `falsy` the {@link CompletionItem.label label}
	// is used.
	// 
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally. For example
	// VS Code when code complete is requested in this example
	// `con<cursor position>` and a completion item with an `insertText` of
	// `console` is provided it will only insert `sole`. Therefore it is
	// recommended to use `textEdit` instead since it avoids additional client
	// side interpretation.
	InsertText string `json:"insertText,omitempty"`
	// The format of the insert text. The format applies to both the
	// `insertText` property and the `newText` property of a provided
	// `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
	// 
	// Please note that the insertTextFormat doesn't apply to
	// `additionalTextEdits`.
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitzero"`
	// How whitespace and indentation is handled during completion
	// item insertion. If not provided the clients default value depends on
	// the `textDocument.completion.insertTextMode` client capability.
	// 
	// @since 3.16.0
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitzero"`
	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor.
	Kind *CompletionItemKind `json:"kind,omitzero"`
	// The label of this completion item.
	// 
	// The label property is also by default the text that
	// is inserted when selecting this completion.
	// 
	// If label details are provided the label itself should
	// be an unqualified name of the completion item.
	Label string `json:"label"`
	// Additional details for the label
	// 
	// @since 3.17.0
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitzero"`
	// Select this item when showing.
	// 
	// *Note* that only one completion item can be selected and that the
	// tool / client decides which item that is. The rule is that the *first*
	// item of those that match best is selected.
	Preselect bool `json:"preselect,omitempty"`
	// A string that should be used when comparing this item
	// with other items. When `falsy` the {@link CompletionItem.label label}
	// is used.
	SortText string `json:"sortText,omitempty"`
	// Tags for this completion item.
	// 
	// @since 3.15.0
	Tags []CompletionItemTag `json:"tags,omitzero"`
	// An {@link TextEdit edit} which is applied to a document when selecting
	// this completion. When an edit is provided the value of
	// {@link CompletionItem.insertText insertText} is ignored.
	// 
	// Most editors support two different operations when accepting a completion
	// item. One is to insert a completion text and the other is to replace an
	// existing text with a completion text. Since this can usually not be
	// predetermined by a server it can report both ranges. Clients need to
	// signal support for `InsertReplaceEdits` via the
	// `textDocument.completion.insertReplaceSupport` client capability
	// property.
	// 
	// *Note 1:* The text edit's range as well as both ranges from an insert
	// replace edit must be a [single line] and they must contain the position
	// at which completion has been requested.
	// *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range
	// must be a prefix of the edit's replace range, that means it must be
	// contained and starting at the same position.
	// 
	// @since 3.16.0 additional type `InsertReplaceEdit`
	TextEdit *Or2[TextEdit, InsertReplaceEdit] `json:"textEdit,omitzero"`
	// The edit text used if the completion item is part of a CompletionList and
	// CompletionList defines an item default for the text edit range.
	// 
	// Clients will only honor this property if they opt into completion list
	// item defaults using the capability `completionList.itemDefaults`.
	// 
	// If not provided and a list's default range is provided the label
	// property is used as a text.
	// 
	// @since 3.17.0
	TextEditText string `json:"textEditText,omitempty"`
}
func (t *CompletionItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["label"]; !exists {
		return fmt.Errorf("missing required field: label")
	}
	type Alias CompletionItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionItem(test)
	
	return nil
}
// Specifies how fields from a completion item should be combined with those
// from `completionList.itemDefaults`.
// 
// If unspecified, all fields will be treated as ApplyKind.Replace.
// 
// If a field's value is ApplyKind.Replace, the value from a completion item (if
// provided and not `null`) will always be used instead of the value from
// `completionItem.itemDefaults`.
// 
// If a field's value is ApplyKind.Merge, the values will be merged using the rules
// defined against each field below.
// 
// Servers are only allowed to return `applyKind` if the client
// signals support for this via the `completionList.applyKindSupport`
// capability.
// 
// @since 3.18.0
type CompletionItemApplyKinds struct {
	// Specifies whether commitCharacters on a completion will replace or be
	// merged with those in `completionList.itemDefaults.commitCharacters`.
	// 
	// If ApplyKind.Replace, the commit characters from the completion item will
	// always be used unless not provided, in which case those from
	// `completionList.itemDefaults.commitCharacters` will be used. An
	// empty list can be used if a completion item does not have any commit
	// characters and also should not use those from
	// `completionList.itemDefaults.commitCharacters`.
	// 
	// If ApplyKind.Merge the commitCharacters for the completion will be the
	// union of all values in both `completionList.itemDefaults.commitCharacters`
	// and the completion's own `commitCharacters`.
	// 
	// @since 3.18.0
	CommitCharacters *ApplyKind `json:"commitCharacters,omitzero"`
	// Specifies whether the `data` field on a completion will replace or
	// be merged with data from `completionList.itemDefaults.data`.
	// 
	// If ApplyKind.Replace, the data from the completion item will be used if
	// provided (and not `null`), otherwise
	// `completionList.itemDefaults.data` will be used. An empty object can
	// be used if a completion item does not have any data but also should
	// not use the value from `completionList.itemDefaults.data`.
	// 
	// If ApplyKind.Merge, a shallow merge will be performed between
	// `completionList.itemDefaults.data` and the completion's own data
	// using the following rules:
	// 
	// - If a completion's `data` field is not provided (or `null`), the
	//   entire `data` field from `completionList.itemDefaults.data` will be
	//   used as-is.
	// - If a completion's `data` field is provided, each field will
	//   overwrite the field of the same name in
	//   `completionList.itemDefaults.data` but no merging of nested fields
	//   within that value will occur.
	// 
	// @since 3.18.0
	Data *ApplyKind `json:"data,omitzero"`
}
func (t *CompletionItemApplyKinds) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CompletionItemApplyKinds
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionItemApplyKinds(test)
	
	return nil
}
// In many cases the items of an actual completion result share the same
// value for properties like `commitCharacters` or the range of a text
// edit. A completion list can therefore define item defaults which will
// be used if a completion item itself doesn't specify the value.
// 
// If a completion list specifies a default value and a completion item
// also specifies a corresponding value, the rules for combining these are
// defined by `applyKinds` (if the client supports it), defaulting to
// ApplyKind.Replace.
// 
// Servers are only allowed to return default values if the client
// signals support for this via the `completionList.itemDefaults`
// capability.
// 
// @since 3.17.0
type CompletionItemDefaults struct {
	// A default commit character set.
	// 
	// @since 3.17.0
	CommitCharacters []string `json:"commitCharacters,omitzero"`
	// A default data value.
	// 
	// @since 3.17.0
	Data any `json:"data,omitempty"`
	// A default edit range.
	// 
	// @since 3.17.0
	EditRange *Or2[Range, EditRangeWithInsertReplace] `json:"editRange,omitzero"`
	// A default insert text format.
	// 
	// @since 3.17.0
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitzero"`
	// A default insert text mode.
	// 
	// @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitzero"`
}
func (t *CompletionItemDefaults) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CompletionItemDefaults
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionItemDefaults(test)
	
	return nil
}
// Additional details for a completion item label.
// 
// @since 3.17.0
type CompletionItemLabelDetails struct {
	// An optional string which is rendered less prominently after {@link CompletionItem.detail}. Should be used
	// for fully qualified names and file paths.
	Description string `json:"description,omitempty"`
	// An optional string which is rendered less prominently directly after {@link CompletionItem.label label},
	// without any spacing. Should be used for function signatures and type annotations.
	Detail string `json:"detail,omitempty"`
}
func (t *CompletionItemLabelDetails) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CompletionItemLabelDetails
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionItemLabelDetails(test)
	
	return nil
}
// @since 3.18.0
type CompletionItemTagOptions struct {
	// The tags supported by the client.
	ValueSet []CompletionItemTag `json:"valueSet"`
}
func (t *CompletionItemTagOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["valueSet"]; !exists {
		return fmt.Errorf("missing required field: valueSet")
	}
	type Alias CompletionItemTagOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionItemTagOptions(test)
	
	return nil
}
// Represents a collection of {@link CompletionItem completion items} to be presented
// in the editor.
type CompletionList struct {
	// Specifies how fields from a completion item should be combined with those
	// from `completionList.itemDefaults`.
	// 
	// If unspecified, all fields will be treated as ApplyKind.Replace.
	// 
	// If a field's value is ApplyKind.Replace, the value from a completion item
	// (if provided and not `null`) will always be used instead of the value
	// from `completionItem.itemDefaults`.
	// 
	// If a field's value is ApplyKind.Merge, the values will be merged using
	// the rules defined against each field below.
	// 
	// Servers are only allowed to return `applyKind` if the client
	// signals support for this via the `completionList.applyKindSupport`
	// capability.
	// 
	// @since 3.18.0
	ApplyKind *CompletionItemApplyKinds `json:"applyKind,omitzero"`
	// This list it not complete. Further typing results in recomputing this list.
	// 
	// Recomputed lists have all their items replaced (not appended) in the
	// incomplete completion sessions.
	IsIncomplete bool `json:"isIncomplete"`
	// In many cases the items of an actual completion result share the same
	// value for properties like `commitCharacters` or the range of a text
	// edit. A completion list can therefore define item defaults which will
	// be used if a completion item itself doesn't specify the value.
	// 
	// If a completion list specifies a default value and a completion item
	// also specifies a corresponding value, the rules for combining these are
	// defined by `applyKinds` (if the client supports it), defaulting to
	// ApplyKind.Replace.
	// 
	// Servers are only allowed to return default values if the client
	// signals support for this via the `completionList.itemDefaults`
	// capability.
	// 
	// @since 3.17.0
	ItemDefaults *CompletionItemDefaults `json:"itemDefaults,omitzero"`
	// The completion items.
	Items []CompletionItem `json:"items"`
}
func (t *CompletionList) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["isIncomplete"]; !exists {
		return fmt.Errorf("missing required field: isIncomplete")
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	type Alias CompletionList
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionList(test)
	
	return nil
}
// The client supports the following `CompletionList` specific
// capabilities.
// 
// @since 3.17.0
type CompletionListCapabilities struct {
	// Specifies whether the client supports `CompletionList.applyKind` to
	// indicate how supported values from `completionList.itemDefaults`
	// and `completion` will be combined.
	// 
	// If a client supports `applyKind` it must support it for all fields
	// that it supports that are listed in `CompletionList.applyKind`. This
	// means when clients add support for new/future fields in completion
	// items the MUST also support merge for them if those fields are
	// defined in `CompletionList.applyKind`.
	// 
	// @since 3.18.0
	ApplyKindSupport bool `json:"applyKindSupport,omitempty"`
	// The client supports the following itemDefaults on
	// a completion list.
	// 
	// The value lists the supported property names of the
	// `CompletionList.itemDefaults` object. If omitted
	// no properties are supported.
	// 
	// @since 3.17.0
	ItemDefaults []string `json:"itemDefaults,omitzero"`
}
func (t *CompletionListCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CompletionListCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionListCapabilities(test)
	
	return nil
}
// Completion options.
type CompletionOptions struct {
	// The list of all possible characters that commit a completion. This field can be used
	// if clients don't support individual commit characters per completion item. See
	// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	// 
	// If a server provides both `allCommitCharacters` and commit characters on an individual
	// completion item the ones on the completion item win.
	// 
	// @since 3.2.0
	AllCommitCharacters []string `json:"allCommitCharacters,omitzero"`
	// The server supports the following `CompletionItem` specific
	// capabilities.
	// 
	// @since 3.17.0
	CompletionItem *ServerCompletionItemOptions `json:"completionItem,omitzero"`
	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
	// Most tools trigger completion request automatically without explicitly requesting
	// it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
	// starts to type an identifier. For example if the user types `c` in a JavaScript file
	// code complete will automatically pop up present `console` besides others as a
	// completion item. Characters that make up identifiers don't need to be listed here.
	// 
	// If code complete should automatically be trigger on characters not being valid inside
	// an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters []string `json:"triggerCharacters,omitzero"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CompletionOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CompletionOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionOptions(test)
	
	return nil
}
// Completion parameters
type CompletionParams struct {
	// The completion context. This is only available it the client specifies
	// to send this using the client capability `textDocument.completion.contextSupport === true`
	Context *CompletionContext `json:"context,omitzero"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *CompletionParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias CompletionParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionParams(test)
	
	return nil
}
// Registration options for a {@link CompletionRequest}.
type CompletionRegistrationOptions struct {
	// The list of all possible characters that commit a completion. This field can be used
	// if clients don't support individual commit characters per completion item. See
	// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	// 
	// If a server provides both `allCommitCharacters` and commit characters on an individual
	// completion item the ones on the completion item win.
	// 
	// @since 3.2.0
	AllCommitCharacters []string `json:"allCommitCharacters,omitzero"`
	// The server supports the following `CompletionItem` specific
	// capabilities.
	// 
	// @since 3.17.0
	CompletionItem *ServerCompletionItemOptions `json:"completionItem,omitzero"`
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`
	// Most tools trigger completion request automatically without explicitly requesting
	// it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
	// starts to type an identifier. For example if the user types `c` in a JavaScript file
	// code complete will automatically pop up present `console` besides others as a
	// completion item. Characters that make up identifiers don't need to be listed here.
	// 
	// If code complete should automatically be trigger on characters not being valid inside
	// an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters []string `json:"triggerCharacters,omitzero"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *CompletionRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias CompletionRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CompletionRegistrationOptions(test)
	
	return nil
}

type ConfigurationItem struct {
	// The scope to get the configuration section for.
	ScopeUri *URI `json:"scopeUri,omitzero"`
	// The configuration section asked for.
	Section string `json:"section,omitempty"`
}
func (t *ConfigurationItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ConfigurationItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ConfigurationItem(test)
	
	return nil
}
// The parameters of a configuration request.
type ConfigurationParams struct {

	Items []ConfigurationItem `json:"items"`
}
func (t *ConfigurationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	type Alias ConfigurationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ConfigurationParams(test)
	
	return nil
}
// Create file operation.
type CreateFile struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitzero"`
	// A create
	Kind string `json:"kind"`
	// Additional options
	Options *CreateFileOptions `json:"options,omitzero"`
	// The resource to create.
	Uri DocumentUri `json:"uri"`
}
func (t *CreateFile) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias CreateFile
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CreateFile(test)
	
	return nil
}
// Options to create a file.
type CreateFileOptions struct {
	// Ignore if exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
	// Overwrite existing file. Overwrite wins over `ignoreIfExists`
	Overwrite bool `json:"overwrite,omitempty"`
}
func (t *CreateFileOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias CreateFileOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CreateFileOptions(test)
	
	return nil
}
// The parameters sent in notifications/requests for user-initiated creation of
// files.
// 
// @since 3.16.0
type CreateFilesParams struct {
	// An array of all files/folders created in this operation.
	Files []FileCreate `json:"files"`
}
func (t *CreateFilesParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["files"]; !exists {
		return fmt.Errorf("missing required field: files")
	}
	type Alias CreateFilesParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = CreateFilesParams(test)
	
	return nil
}
// @since 3.14.0
type DeclarationClientCapabilities struct {
	// Whether declaration supports dynamic registration. If this is set to `true`
	// the client supports the new `DeclarationRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports additional metadata in the form of declaration links.
	LinkSupport bool `json:"linkSupport,omitempty"`
}
func (t *DeclarationClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DeclarationClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeclarationClientCapabilities(test)
	
	return nil
}

type DeclarationOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DeclarationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DeclarationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeclarationOptions(test)
	
	return nil
}

type DeclarationParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DeclarationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DeclarationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeclarationParams(test)
	
	return nil
}

type DeclarationRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DeclarationRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DeclarationRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeclarationRegistrationOptions(test)
	
	return nil
}
// Client Capabilities for a {@link DefinitionRequest}.
type DefinitionClientCapabilities struct {
	// Whether definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports additional metadata in the form of definition links.
	// 
	// @since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}
func (t *DefinitionClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DefinitionClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DefinitionClientCapabilities(test)
	
	return nil
}
// Server Capabilities for a {@link DefinitionRequest}.
type DefinitionOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DefinitionOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DefinitionOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DefinitionOptions(test)
	
	return nil
}
// Parameters for a {@link DefinitionRequest}.
type DefinitionParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DefinitionParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DefinitionParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DefinitionParams(test)
	
	return nil
}
// Registration options for a {@link DefinitionRequest}.
type DefinitionRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DefinitionRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DefinitionRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DefinitionRegistrationOptions(test)
	
	return nil
}
// Delete file operation
type DeleteFile struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitzero"`
	// A delete
	Kind string `json:"kind"`
	// Delete options.
	Options *DeleteFileOptions `json:"options,omitzero"`
	// The file to delete.
	Uri DocumentUri `json:"uri"`
}
func (t *DeleteFile) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias DeleteFile
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeleteFile(test)
	
	return nil
}
// Delete file options
type DeleteFileOptions struct {
	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
	// Delete the content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`
}
func (t *DeleteFileOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DeleteFileOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeleteFileOptions(test)
	
	return nil
}
// The parameters sent in notifications/requests for user-initiated deletes of
// files.
// 
// @since 3.16.0
type DeleteFilesParams struct {
	// An array of all files/folders deleted in this operation.
	Files []FileDelete `json:"files"`
}
func (t *DeleteFilesParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["files"]; !exists {
		return fmt.Errorf("missing required field: files")
	}
	type Alias DeleteFilesParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DeleteFilesParams(test)
	
	return nil
}
// Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
// are only valid in the scope of a resource.
type Diagnostic struct {
	// The diagnostic's code, which usually appear in the user interface.
	Code *Or2[int32, string] `json:"code,omitzero"`
	// An optional property to describe the error code.
	// Requires the code field (above) to be present/not null.
	// 
	// @since 3.16.0
	CodeDescription *CodeDescription `json:"codeDescription,omitzero"`
	// A data entry field that is preserved between a `textDocument/publishDiagnostics`
	// notification and `textDocument/codeAction` request.
	// 
	// @since 3.16.0
	Data any `json:"data,omitempty"`
	// The diagnostic's message. It usually appears in the user interface
	Message string `json:"message"`
	// The range at which the message applies
	Range Range `json:"range"`
	// An array of related diagnostic information, e.g. when symbol-names within
	// a scope collide all definitions can be marked via this property.
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitzero"`
	// The diagnostic's severity. To avoid interpretation mismatches when a
	// server is used with different clients it is highly recommended that servers
	// always provide a severity value.
	Severity *DiagnosticSeverity `json:"severity,omitzero"`
	// A human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'. It usually
	// appears in the user interface.
	Source string `json:"source,omitempty"`
	// Additional metadata about the diagnostic.
	// 
	// @since 3.15.0
	Tags []DiagnosticTag `json:"tags,omitzero"`
}
func (t *Diagnostic) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["message"]; !exists {
		return fmt.Errorf("missing required field: message")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias Diagnostic
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Diagnostic(test)
	
	return nil
}
// Client capabilities specific to diagnostic pull requests.
// 
// @since 3.17.0
type DiagnosticClientCapabilities struct {
	// Client supports a codeDescription property
	// 
	// @since 3.16.0
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`
	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	// 
	// @since 3.16.0
	DataSupport bool `json:"dataSupport,omitempty"`
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the clients supports related documents for document diagnostic pulls.
	RelatedDocumentSupport bool `json:"relatedDocumentSupport,omitempty"`
	// Whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"relatedInformation,omitempty"`
	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.15.0
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitzero"`
}
func (t *DiagnosticClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DiagnosticClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticClientCapabilities(test)
	
	return nil
}
// Diagnostic options.
// 
// @since 3.17.0
type DiagnosticOptions struct {
	// An optional identifier under which the diagnostics are
	// managed by the client.
	Identifier string `json:"identifier,omitempty"`
	// Whether the language has inter file dependencies meaning that
	// editing code in one file can result in a different diagnostic
	// set in another file. Inter file dependencies are common for
	// most programming languages and typically uncommon for linters.
	InterFileDependencies bool `json:"interFileDependencies"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}
func (t *DiagnosticOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["interFileDependencies"]; !exists {
		return fmt.Errorf("missing required field: interFileDependencies")
	}
	if _, exists := m["workspaceDiagnostics"]; !exists {
		return fmt.Errorf("missing required field: workspaceDiagnostics")
	}
	type Alias DiagnosticOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticOptions(test)
	
	return nil
}
// Diagnostic registration options.
// 
// @since 3.17.0
type DiagnosticRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// An optional identifier under which the diagnostics are
	// managed by the client.
	Identifier string `json:"identifier,omitempty"`
	// Whether the language has inter file dependencies meaning that
	// editing code in one file can result in a different diagnostic
	// set in another file. Inter file dependencies are common for
	// most programming languages and typically uncommon for linters.
	InterFileDependencies bool `json:"interFileDependencies"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics bool `json:"workspaceDiagnostics"`
}
func (t *DiagnosticRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	if _, exists := m["interFileDependencies"]; !exists {
		return fmt.Errorf("missing required field: interFileDependencies")
	}
	if _, exists := m["workspaceDiagnostics"]; !exists {
		return fmt.Errorf("missing required field: workspaceDiagnostics")
	}
	type Alias DiagnosticRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticRegistrationOptions(test)
	
	return nil
}
// Represents a related message and source code location for a diagnostic. This should be
// used to point to code locations that cause or related to a diagnostics, e.g when duplicating
// a symbol in a scope.
type DiagnosticRelatedInformation struct {
	// The location of this related diagnostic information.
	Location Location `json:"location"`
	// The message of this related diagnostic information.
	Message string `json:"message"`
}
func (t *DiagnosticRelatedInformation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["location"]; !exists {
		return fmt.Errorf("missing required field: location")
	}
	if _, exists := m["message"]; !exists {
		return fmt.Errorf("missing required field: message")
	}
	type Alias DiagnosticRelatedInformation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticRelatedInformation(test)
	
	return nil
}
// Cancellation data returned from a diagnostic request.
// 
// @since 3.17.0
type DiagnosticServerCancellationData struct {

	RetriggerRequest bool `json:"retriggerRequest"`
}
func (t *DiagnosticServerCancellationData) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["retriggerRequest"]; !exists {
		return fmt.Errorf("missing required field: retriggerRequest")
	}
	type Alias DiagnosticServerCancellationData
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticServerCancellationData(test)
	
	return nil
}
// Workspace client capabilities specific to diagnostic pull requests.
// 
// @since 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	// 
	// Note that this event is global and will force the client to refresh all
	// pulled diagnostics currently shown. It should be used with absolute care and
	// is useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}
func (t *DiagnosticWorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DiagnosticWorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticWorkspaceClientCapabilities(test)
	
	return nil
}
// General diagnostics capabilities for pull and push model.
type DiagnosticsCapabilities struct {
	// Client supports a codeDescription property
	// 
	// @since 3.16.0
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`
	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	// 
	// @since 3.16.0
	DataSupport bool `json:"dataSupport,omitempty"`
	// Whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"relatedInformation,omitempty"`
	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.15.0
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitzero"`
}
func (t *DiagnosticsCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DiagnosticsCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DiagnosticsCapabilities(test)
	
	return nil
}

type DidChangeConfigurationClientCapabilities struct {
	// Did change configuration notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *DidChangeConfigurationClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DidChangeConfigurationClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeConfigurationClientCapabilities(test)
	
	return nil
}
// The parameters of a change configuration notification.
type DidChangeConfigurationParams struct {
	// The actual changed settings
	Settings any `json:"settings"`
}
func (t *DidChangeConfigurationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["settings"]; !exists {
		return fmt.Errorf("missing required field: settings")
	}
	type Alias DidChangeConfigurationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeConfigurationParams(test)
	
	return nil
}

type DidChangeConfigurationRegistrationOptions struct {

	Section Or2[string, []string] `json:"section,omitzero"`
}
func (t *DidChangeConfigurationRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DidChangeConfigurationRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeConfigurationRegistrationOptions(test)
	
	return nil
}
// The params sent in a change notebook document notification.
// 
// @since 3.17.0
type DidChangeNotebookDocumentParams struct {
	// The actual changes to the notebook document.
	// 
	// The changes describe single state changes to the notebook document.
	// So if there are two changes c1 (at array index 0) and c2 (at array
	// index 1) for a notebook in state S then c1 moves the notebook from
	// S to S' and c2 from S' to S''. So c1 is computed on the state S and
	// c2 is computed on the state S'.
	// 
	// To mirror the content of a notebook using change events use the following approach:
	// - start with the same initial content
	// - apply the 'notebookDocument/didChange' notifications in the order you receive them.
	// - apply the `NotebookChangeEvent`s in a single notification in the order
	//   you receive them.
	Change NotebookDocumentChangeEvent `json:"change"`
	// The notebook document that did change. The version number points
	// to the version after all provided changes have been applied. If
	// only the text document content of a cell changes the notebook version
	// doesn't necessarily have to change.
	NotebookDocument VersionedNotebookDocumentIdentifier `json:"notebookDocument"`
}
func (t *DidChangeNotebookDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["change"]; !exists {
		return fmt.Errorf("missing required field: change")
	}
	if _, exists := m["notebookDocument"]; !exists {
		return fmt.Errorf("missing required field: notebookDocument")
	}
	type Alias DidChangeNotebookDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeNotebookDocumentParams(test)
	
	return nil
}
// The change text document notification's parameters.
type DidChangeTextDocumentParams struct {
	// The actual content changes. The content changes describe single state changes
	// to the document. So if there are two content changes c1 (at array index 0) and
	// c2 (at array index 1) for a document in state S then c1 moves the document from
	// S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed
	// on the state S'.
	// 
	// To mirror the content of a document using change events use the following approach:
	// - start with the same initial content
	// - apply the 'textDocument/didChange' notifications in the order you receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single notification in the order
	//   you receive them.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`
}
func (t *DidChangeTextDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["contentChanges"]; !exists {
		return fmt.Errorf("missing required field: contentChanges")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DidChangeTextDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeTextDocumentParams(test)
	
	return nil
}

type DidChangeWatchedFilesClientCapabilities struct {
	// Did change watched files notification supports dynamic registration. Please note
	// that the current protocol doesn't support static configuration for file changes
	// from the server side.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client has support for {@link  RelativePattern relative pattern}
	// or not.
	// 
	// @since 3.17.0
	RelativePatternSupport bool `json:"relativePatternSupport,omitempty"`
}
func (t *DidChangeWatchedFilesClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DidChangeWatchedFilesClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeWatchedFilesClientCapabilities(test)
	
	return nil
}
// The watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {
	// The actual file events.
	Changes []FileEvent `json:"changes"`
}
func (t *DidChangeWatchedFilesParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["changes"]; !exists {
		return fmt.Errorf("missing required field: changes")
	}
	type Alias DidChangeWatchedFilesParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeWatchedFilesParams(test)
	
	return nil
}
// Describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {
	// The watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}
func (t *DidChangeWatchedFilesRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["watchers"]; !exists {
		return fmt.Errorf("missing required field: watchers")
	}
	type Alias DidChangeWatchedFilesRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeWatchedFilesRegistrationOptions(test)
	
	return nil
}
// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
}
func (t *DidChangeWorkspaceFoldersParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["event"]; !exists {
		return fmt.Errorf("missing required field: event")
	}
	type Alias DidChangeWorkspaceFoldersParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidChangeWorkspaceFoldersParams(test)
	
	return nil
}
// The params sent in a close notebook document notification.
// 
// @since 3.17.0
type DidCloseNotebookDocumentParams struct {
	// The text documents that represent the content
	// of a notebook cell that got closed.
	CellTextDocuments []TextDocumentIdentifier `json:"cellTextDocuments"`
	// The notebook document that got closed.
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`
}
func (t *DidCloseNotebookDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["cellTextDocuments"]; !exists {
		return fmt.Errorf("missing required field: cellTextDocuments")
	}
	if _, exists := m["notebookDocument"]; !exists {
		return fmt.Errorf("missing required field: notebookDocument")
	}
	type Alias DidCloseNotebookDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidCloseNotebookDocumentParams(test)
	
	return nil
}
// The parameters sent in a close text document notification
type DidCloseTextDocumentParams struct {
	// The document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
func (t *DidCloseTextDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DidCloseTextDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidCloseTextDocumentParams(test)
	
	return nil
}
// The params sent in an open notebook document notification.
// 
// @since 3.17.0
type DidOpenNotebookDocumentParams struct {
	// The text documents that represent the content
	// of a notebook cell.
	CellTextDocuments []TextDocumentItem `json:"cellTextDocuments"`
	// The notebook document that got opened.
	NotebookDocument NotebookDocument `json:"notebookDocument"`
}
func (t *DidOpenNotebookDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["cellTextDocuments"]; !exists {
		return fmt.Errorf("missing required field: cellTextDocuments")
	}
	if _, exists := m["notebookDocument"]; !exists {
		return fmt.Errorf("missing required field: notebookDocument")
	}
	type Alias DidOpenNotebookDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidOpenNotebookDocumentParams(test)
	
	return nil
}
// The parameters sent in an open text document notification
type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}
func (t *DidOpenTextDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DidOpenTextDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidOpenTextDocumentParams(test)
	
	return nil
}
// The params sent in a save notebook document notification.
// 
// @since 3.17.0
type DidSaveNotebookDocumentParams struct {
	// The notebook document that got saved.
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`
}
func (t *DidSaveNotebookDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["notebookDocument"]; !exists {
		return fmt.Errorf("missing required field: notebookDocument")
	}
	type Alias DidSaveNotebookDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidSaveNotebookDocumentParams(test)
	
	return nil
}
// The parameters sent in a save text document notification
type DidSaveTextDocumentParams struct {
	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`
	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
func (t *DidSaveTextDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DidSaveTextDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DidSaveTextDocumentParams(test)
	
	return nil
}

type DocumentColorClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `DocumentColorRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *DocumentColorClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentColorClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentColorClientCapabilities(test)
	
	return nil
}

type DocumentColorOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentColorOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentColorOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentColorOptions(test)
	
	return nil
}
// Parameters for a {@link DocumentColorRequest}.
type DocumentColorParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentColorParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentColorParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentColorParams(test)
	
	return nil
}

type DocumentColorRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentColorRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DocumentColorRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentColorRegistrationOptions(test)
	
	return nil
}
// Parameters of the document diagnostic request.
// 
// @since 3.17.0
type DocumentDiagnosticParams struct {
	// The additional identifier  provided during registration.
	Identifier string `json:"identifier,omitempty"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The result id of a previous response if provided.
	PreviousResultId string `json:"previousResultId,omitempty"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentDiagnosticParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentDiagnosticParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentDiagnosticParams(test)
	
	return nil
}
// A partial result for a document diagnostic report.
// 
// @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {

	RelatedDocuments map[DocumentUri]Or2[FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport] `json:"relatedDocuments"`
}
func (t *DocumentDiagnosticReportPartialResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["relatedDocuments"]; !exists {
		return fmt.Errorf("missing required field: relatedDocuments")
	}
	type Alias DocumentDiagnosticReportPartialResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentDiagnosticReportPartialResult(test)
	
	return nil
}
// Client capabilities of a {@link DocumentFormattingRequest}.
type DocumentFormattingClientCapabilities struct {
	// Whether formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *DocumentFormattingClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentFormattingClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentFormattingClientCapabilities(test)
	
	return nil
}
// Provider options for a {@link DocumentFormattingRequest}.
type DocumentFormattingOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentFormattingOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentFormattingOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentFormattingOptions(test)
	
	return nil
}
// The parameters of a {@link DocumentFormattingRequest}.
type DocumentFormattingParams struct {
	// The format options.
	Options FormattingOptions `json:"options"`
	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentFormattingParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["options"]; !exists {
		return fmt.Errorf("missing required field: options")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentFormattingParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentFormattingParams(test)
	
	return nil
}
// Registration options for a {@link DocumentFormattingRequest}.
type DocumentFormattingRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentFormattingRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DocumentFormattingRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentFormattingRegistrationOptions(test)
	
	return nil
}
// A document highlight is a range inside a text document which deserves
// special attention. Usually a document highlight is visualized by changing
// the background color of its range.
type DocumentHighlight struct {
	// The highlight kind, default is {@link DocumentHighlightKind.Text text}.
	Kind *DocumentHighlightKind `json:"kind,omitzero"`
	// The range this highlight applies to.
	Range Range `json:"range"`
}
func (t *DocumentHighlight) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias DocumentHighlight
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentHighlight(test)
	
	return nil
}
// Client Capabilities for a {@link DocumentHighlightRequest}.
type DocumentHighlightClientCapabilities struct {
	// Whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *DocumentHighlightClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentHighlightClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentHighlightClientCapabilities(test)
	
	return nil
}
// Provider options for a {@link DocumentHighlightRequest}.
type DocumentHighlightOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentHighlightOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentHighlightOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentHighlightOptions(test)
	
	return nil
}
// Parameters for a {@link DocumentHighlightRequest}.
type DocumentHighlightParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentHighlightParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentHighlightParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentHighlightParams(test)
	
	return nil
}
// Registration options for a {@link DocumentHighlightRequest}.
type DocumentHighlightRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentHighlightRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DocumentHighlightRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentHighlightRegistrationOptions(test)
	
	return nil
}
// A document link is a range in a text document that links to an internal or external resource, like another
// text document or a web site.
type DocumentLink struct {
	// A data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data any `json:"data,omitempty"`
	// The range this link applies to.
	Range Range `json:"range"`
	// The uri this link points to. If missing a resolve request is sent later.
	Target *URI `json:"target,omitzero"`
	// The tooltip text when you hover over this link.
	// 
	// If a tooltip is provided, is will be displayed in a string that includes instructions on how to
	// trigger the link, such as `{0} (ctrl + click)`. The specific instructions vary depending on OS,
	// user settings, and localization.
	// 
	// @since 3.15.0
	Tooltip string `json:"tooltip,omitempty"`
}
func (t *DocumentLink) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias DocumentLink
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentLink(test)
	
	return nil
}
// The client capabilities of a {@link DocumentLinkRequest}.
type DocumentLinkClientCapabilities struct {
	// Whether document link supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client supports the `tooltip` property on `DocumentLink`.
	// 
	// @since 3.15.0
	TooltipSupport bool `json:"tooltipSupport,omitempty"`
}
func (t *DocumentLinkClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentLinkClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentLinkClientCapabilities(test)
	
	return nil
}
// Provider options for a {@link DocumentLinkRequest}.
type DocumentLinkOptions struct {
	// Document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentLinkOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentLinkOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentLinkOptions(test)
	
	return nil
}
// The parameters of a {@link DocumentLinkRequest}.
type DocumentLinkParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentLinkParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentLinkParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentLinkParams(test)
	
	return nil
}
// Registration options for a {@link DocumentLinkRequest}.
type DocumentLinkRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentLinkRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DocumentLinkRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentLinkRegistrationOptions(test)
	
	return nil
}
// Client capabilities of a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingClientCapabilities struct {
	// Whether on type formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *DocumentOnTypeFormattingClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentOnTypeFormattingClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentOnTypeFormattingClientCapabilities(test)
	
	return nil
}
// Provider options for a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingOptions struct {
	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`
	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitzero"`
}
func (t *DocumentOnTypeFormattingOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["firstTriggerCharacter"]; !exists {
		return fmt.Errorf("missing required field: firstTriggerCharacter")
	}
	type Alias DocumentOnTypeFormattingOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentOnTypeFormattingOptions(test)
	
	return nil
}
// The parameters of a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingParams struct {
	// The character that has been typed that triggered the formatting
	// on type request. That is not necessarily the last character that
	// got inserted into the document since the client could auto insert
	// characters as well (e.g. like automatic brace completion).
	Ch string `json:"ch"`
	// The formatting options.
	Options FormattingOptions `json:"options"`
	// The position around which the on type formatting should happen.
	// This is not necessarily the exact position where the character denoted
	// by the property `ch` got typed.
	Position Position `json:"position"`
	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
func (t *DocumentOnTypeFormattingParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["ch"]; !exists {
		return fmt.Errorf("missing required field: ch")
	}
	if _, exists := m["options"]; !exists {
		return fmt.Errorf("missing required field: options")
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentOnTypeFormattingParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentOnTypeFormattingParams(test)
	
	return nil
}
// Registration options for a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`
	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitzero"`
}
func (t *DocumentOnTypeFormattingRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	if _, exists := m["firstTriggerCharacter"]; !exists {
		return fmt.Errorf("missing required field: firstTriggerCharacter")
	}
	type Alias DocumentOnTypeFormattingRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentOnTypeFormattingRegistrationOptions(test)
	
	return nil
}
// Client capabilities of a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingClientCapabilities struct {
	// Whether range formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client supports formatting multiple ranges at once.
	// 
	// @since 3.18.0
	// @proposed
	RangesSupport bool `json:"rangesSupport,omitempty"`
}
func (t *DocumentRangeFormattingClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentRangeFormattingClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentRangeFormattingClientCapabilities(test)
	
	return nil
}
// Provider options for a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingOptions struct {
	// Whether the server supports formatting multiple ranges at once.
	// 
	// @since 3.18.0
	// @proposed
	RangesSupport bool `json:"rangesSupport,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentRangeFormattingOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentRangeFormattingOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentRangeFormattingOptions(test)
	
	return nil
}
// The parameters of a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingParams struct {
	// The format options
	Options FormattingOptions `json:"options"`
	// The range to format
	Range Range `json:"range"`
	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentRangeFormattingParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["options"]; !exists {
		return fmt.Errorf("missing required field: options")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentRangeFormattingParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentRangeFormattingParams(test)
	
	return nil
}
// Registration options for a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Whether the server supports formatting multiple ranges at once.
	// 
	// @since 3.18.0
	// @proposed
	RangesSupport bool `json:"rangesSupport,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentRangeFormattingRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DocumentRangeFormattingRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentRangeFormattingRegistrationOptions(test)
	
	return nil
}
// The parameters of a {@link DocumentRangesFormattingRequest}.
// 
// @since 3.18.0
// @proposed
type DocumentRangesFormattingParams struct {
	// The format options
	Options FormattingOptions `json:"options"`
	// The ranges to format
	Ranges []Range `json:"ranges"`
	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentRangesFormattingParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["options"]; !exists {
		return fmt.Errorf("missing required field: options")
	}
	if _, exists := m["ranges"]; !exists {
		return fmt.Errorf("missing required field: ranges")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentRangesFormattingParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentRangesFormattingParams(test)
	
	return nil
}
// Represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they
// have two ranges: one that encloses its definition and one that points to
// its most interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {
	// Children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children,omitzero"`
	// Indicates if this symbol is deprecated.
	// 
	// @deprecated Use tags instead
	Deprecated bool `json:"deprecated,omitempty"`
	// More detail for this symbol, e.g the signature of a function.
	Detail string `json:"detail,omitempty"`
	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`
	// The name of this symbol. Will be displayed in the user interface and therefore must not be
	// an empty string or a string only consisting of white spaces.
	Name string `json:"name"`
	// The range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to determine if the clients cursor is
	// inside the symbol to reveal in the symbol in the UI.
	Range Range `json:"range"`
	// The range that should be selected and revealed when this symbol is being picked, e.g the name of a function.
	// Must be contained by the `range`.
	SelectionRange Range `json:"selectionRange"`
	// Tags for this document symbol.
	// 
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitzero"`
}
func (t *DocumentSymbol) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["selectionRange"]; !exists {
		return fmt.Errorf("missing required field: selectionRange")
	}
	type Alias DocumentSymbol
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentSymbol(test)
	
	return nil
}
// Client Capabilities for a {@link DocumentSymbolRequest}.
type DocumentSymbolClientCapabilities struct {
	// Whether document symbol supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`
	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.
	// 
	// @since 3.16.0
	LabelSupport bool `json:"labelSupport,omitempty"`
	// Specific capabilities for the `SymbolKind` in the
	// `textDocument/documentSymbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbolKind,omitzero"`
	// The client supports tags on `SymbolInformation`. Tags are supported on
	// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.16.0
	TagSupport *ClientSymbolTagOptions `json:"tagSupport,omitzero"`
}
func (t *DocumentSymbolClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentSymbolClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentSymbolClientCapabilities(test)
	
	return nil
}
// Provider options for a {@link DocumentSymbolRequest}.
type DocumentSymbolOptions struct {
	// A human-readable string that is shown when multiple outlines trees
	// are shown for the same document.
	// 
	// @since 3.16.0
	Label string `json:"label,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentSymbolOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias DocumentSymbolOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentSymbolOptions(test)
	
	return nil
}
// Parameters for a {@link DocumentSymbolRequest}.
type DocumentSymbolParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *DocumentSymbolParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias DocumentSymbolParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentSymbolParams(test)
	
	return nil
}
// Registration options for a {@link DocumentSymbolRequest}.
type DocumentSymbolRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// A human-readable string that is shown when multiple outlines trees
	// are shown for the same document.
	// 
	// @since 3.16.0
	Label string `json:"label,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *DocumentSymbolRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias DocumentSymbolRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = DocumentSymbolRegistrationOptions(test)
	
	return nil
}
// Edit range variant that includes ranges for insert and replace operations.
// 
// @since 3.18.0
type EditRangeWithInsertReplace struct {

	Insert Range `json:"insert"`

	Replace Range `json:"replace"`
}
func (t *EditRangeWithInsertReplace) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["insert"]; !exists {
		return fmt.Errorf("missing required field: insert")
	}
	if _, exists := m["replace"]; !exists {
		return fmt.Errorf("missing required field: replace")
	}
	type Alias EditRangeWithInsertReplace
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = EditRangeWithInsertReplace(test)
	
	return nil
}
// The client capabilities of a {@link ExecuteCommandRequest}.
type ExecuteCommandClientCapabilities struct {
	// Execute command supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *ExecuteCommandClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ExecuteCommandClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ExecuteCommandClientCapabilities(test)
	
	return nil
}
// The server capabilities of a {@link ExecuteCommandRequest}.
type ExecuteCommandOptions struct {
	// The commands to be executed on the server
	Commands []string `json:"commands"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *ExecuteCommandOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["commands"]; !exists {
		return fmt.Errorf("missing required field: commands")
	}
	type Alias ExecuteCommandOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ExecuteCommandOptions(test)
	
	return nil
}
// The parameters of a {@link ExecuteCommandRequest}.
type ExecuteCommandParams struct {
	// Arguments that the command should be invoked with.
	Arguments []any `json:"arguments,omitzero"`
	// The identifier of the actual command handler.
	Command string `json:"command"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *ExecuteCommandParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["command"]; !exists {
		return fmt.Errorf("missing required field: command")
	}
	type Alias ExecuteCommandParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ExecuteCommandParams(test)
	
	return nil
}
// Registration options for a {@link ExecuteCommandRequest}.
type ExecuteCommandRegistrationOptions struct {
	// The commands to be executed on the server
	Commands []string `json:"commands"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *ExecuteCommandRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["commands"]; !exists {
		return fmt.Errorf("missing required field: commands")
	}
	type Alias ExecuteCommandRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ExecuteCommandRegistrationOptions(test)
	
	return nil
}

type ExecutionSummary struct {
	// A strict monotonically increasing value
	// indicating the execution order of a cell
	// inside a notebook.
	ExecutionOrder uint32 `json:"executionOrder"`
	// Whether the execution was successful or
	// not if known by the client.
	Success bool `json:"success,omitempty"`
}
func (t *ExecutionSummary) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["executionOrder"]; !exists {
		return fmt.Errorf("missing required field: executionOrder")
	}
	type Alias ExecutionSummary
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ExecutionSummary(test)
	
	return nil
}
// Represents information on a file/folder create.
// 
// @since 3.16.0
type FileCreate struct {
	// A file:// URI for the location of the file/folder being created.
	Uri string `json:"uri"`
}
func (t *FileCreate) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias FileCreate
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileCreate(test)
	
	return nil
}
// Represents information on a file/folder delete.
// 
// @since 3.16.0
type FileDelete struct {
	// A file:// URI for the location of the file/folder being deleted.
	Uri string `json:"uri"`
}
func (t *FileDelete) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias FileDelete
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileDelete(test)
	
	return nil
}
// An event describing a file change.
type FileEvent struct {
	// The change type.
	Type FileChangeType `json:"type"`
	// The file's uri.
	Uri DocumentUri `json:"uri"`
}
func (t *FileEvent) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["type"]; !exists {
		return fmt.Errorf("missing required field: type")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias FileEvent
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileEvent(test)
	
	return nil
}
// Capabilities relating to events from file operations by the user in the client.
// 
// These events do not come from the file system, they come from user operations
// like renaming a file in the UI.
// 
// @since 3.16.0
type FileOperationClientCapabilities struct {
	// The client has support for sending didCreateFiles notifications.
	DidCreate bool `json:"didCreate,omitempty"`
	// The client has support for sending didDeleteFiles notifications.
	DidDelete bool `json:"didDelete,omitempty"`
	// The client has support for sending didRenameFiles notifications.
	DidRename bool `json:"didRename,omitempty"`
	// Whether the client supports dynamic registration for file requests/notifications.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client has support for sending willCreateFiles requests.
	WillCreate bool `json:"willCreate,omitempty"`
	// The client has support for sending willDeleteFiles requests.
	WillDelete bool `json:"willDelete,omitempty"`
	// The client has support for sending willRenameFiles requests.
	WillRename bool `json:"willRename,omitempty"`
}
func (t *FileOperationClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias FileOperationClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileOperationClientCapabilities(test)
	
	return nil
}
// A filter to describe in which file operation requests or notifications
// the server is interested in receiving.
// 
// @since 3.16.0
type FileOperationFilter struct {
	// The actual file operation pattern.
	Pattern FileOperationPattern `json:"pattern"`
	// A Uri scheme like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
}
func (t *FileOperationFilter) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["pattern"]; !exists {
		return fmt.Errorf("missing required field: pattern")
	}
	type Alias FileOperationFilter
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileOperationFilter(test)
	
	return nil
}
// Options for notifications/requests for user operations on files.
// 
// @since 3.16.0
type FileOperationOptions struct {
	// The server is interested in receiving didCreateFiles notifications.
	DidCreate *FileOperationRegistrationOptions `json:"didCreate,omitzero"`
	// The server is interested in receiving didDeleteFiles file notifications.
	DidDelete *FileOperationRegistrationOptions `json:"didDelete,omitzero"`
	// The server is interested in receiving didRenameFiles notifications.
	DidRename *FileOperationRegistrationOptions `json:"didRename,omitzero"`
	// The server is interested in receiving willCreateFiles requests.
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitzero"`
	// The server is interested in receiving willDeleteFiles file requests.
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitzero"`
	// The server is interested in receiving willRenameFiles requests.
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitzero"`
}
func (t *FileOperationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias FileOperationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileOperationOptions(test)
	
	return nil
}
// A pattern to describe in which file operation requests or notifications
// the server is interested in receiving.
// 
// @since 3.16.0
type FileOperationPattern struct {
	// The glob pattern to match. Glob patterns can have the following syntax:
	// - `*` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `**` to match any number of path segments, including none
	// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
	Glob string `json:"glob"`
	// Whether to match files or folders with this pattern.
	// 
	// Matches both if undefined.
	Matches *FileOperationPatternKind `json:"matches,omitzero"`
	// Additional options used during matching.
	Options *FileOperationPatternOptions `json:"options,omitzero"`
}
func (t *FileOperationPattern) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["glob"]; !exists {
		return fmt.Errorf("missing required field: glob")
	}
	type Alias FileOperationPattern
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileOperationPattern(test)
	
	return nil
}
// Matching options for the file operation pattern.
// 
// @since 3.16.0
type FileOperationPatternOptions struct {
	// The pattern should be matched ignoring casing.
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}
func (t *FileOperationPatternOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias FileOperationPatternOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileOperationPatternOptions(test)
	
	return nil
}
// The options to register for file operations.
// 
// @since 3.16.0
type FileOperationRegistrationOptions struct {
	// The actual filters.
	Filters []FileOperationFilter `json:"filters"`
}
func (t *FileOperationRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["filters"]; !exists {
		return fmt.Errorf("missing required field: filters")
	}
	type Alias FileOperationRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileOperationRegistrationOptions(test)
	
	return nil
}
// Represents information on a file/folder rename.
// 
// @since 3.16.0
type FileRename struct {
	// A file:// URI for the new location of the file/folder being renamed.
	NewUri string `json:"newUri"`
	// A file:// URI for the original location of the file/folder being renamed.
	OldUri string `json:"oldUri"`
}
func (t *FileRename) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["newUri"]; !exists {
		return fmt.Errorf("missing required field: newUri")
	}
	if _, exists := m["oldUri"]; !exists {
		return fmt.Errorf("missing required field: oldUri")
	}
	type Alias FileRename
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileRename(test)
	
	return nil
}

type FileSystemWatcher struct {
	// The glob pattern to watch. See {@link GlobPattern glob pattern} for more detail.
	// 
	// @since 3.17.0 support for relative patterns.
	GlobPattern GlobPattern `json:"globPattern"`
	// The kind of events of interest. If omitted it defaults
	// to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	// which is 7.
	Kind *WatchKind `json:"kind,omitzero"`
}
func (t *FileSystemWatcher) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["globPattern"]; !exists {
		return fmt.Errorf("missing required field: globPattern")
	}
	type Alias FileSystemWatcher
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FileSystemWatcher(test)
	
	return nil
}
// Represents a folding range. To be valid, start and end line must be bigger than zero and smaller
// than the number of lines in the document. Clients are free to ignore invalid ranges.
type FoldingRange struct {
	// The text that the client should show when the specified range is
	// collapsed. If not defined or not supported by the client, a default
	// will be chosen by the client.
	// 
	// @since 3.17.0
	CollapsedText string `json:"collapsedText,omitempty"`
	// The zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
	EndCharacter uint32 `json:"endCharacter,omitempty"`
	// The zero-based end line of the range to fold. The folded area ends with the line's last character.
	// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	EndLine uint32 `json:"endLine"`
	// Describes the kind of the folding range such as 'comment' or 'region'. The kind
	// is used to categorize folding ranges and used by commands like 'Fold all comments'.
	// See {@link FoldingRangeKind} for an enumeration of standardized kinds.
	Kind *FoldingRangeKind `json:"kind,omitzero"`
	// The zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	StartCharacter uint32 `json:"startCharacter,omitempty"`
	// The zero-based start line of the range to fold. The folded area starts after the line's last character.
	// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	StartLine uint32 `json:"startLine"`
}
func (t *FoldingRange) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["endLine"]; !exists {
		return fmt.Errorf("missing required field: endLine")
	}
	if _, exists := m["startLine"]; !exists {
		return fmt.Errorf("missing required field: startLine")
	}
	type Alias FoldingRange
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FoldingRange(test)
	
	return nil
}

type FoldingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for folding range
	// providers. If this is set to `true` the client supports the new
	// `FoldingRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Specific options for the folding range.
	// 
	// @since 3.17.0
	FoldingRange *ClientFoldingRangeOptions `json:"foldingRange,omitzero"`
	// Specific options for the folding range kind.
	// 
	// @since 3.17.0
	FoldingRangeKind *ClientFoldingRangeKindOptions `json:"foldingRangeKind,omitzero"`
	// If set, the client signals that it only supports folding complete lines.
	// If set, client will ignore specified `startCharacter` and `endCharacter`
	// properties in a FoldingRange.
	LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
	// The maximum number of folding ranges that the client prefers to receive
	// per document. The value serves as a hint, servers are free to follow the
	// limit.
	RangeLimit uint32 `json:"rangeLimit,omitempty"`
}
func (t *FoldingRangeClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias FoldingRangeClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FoldingRangeClientCapabilities(test)
	
	return nil
}

type FoldingRangeOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *FoldingRangeOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias FoldingRangeOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FoldingRangeOptions(test)
	
	return nil
}
// Parameters for a {@link FoldingRangeRequest}.
type FoldingRangeParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *FoldingRangeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias FoldingRangeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FoldingRangeParams(test)
	
	return nil
}

type FoldingRangeRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *FoldingRangeRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias FoldingRangeRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FoldingRangeRegistrationOptions(test)
	
	return nil
}
// Client workspace capabilities specific to folding ranges
// 
// @since 3.18.0
// @proposed
type FoldingRangeWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	// 
	// Note that this event is global and will force the client to refresh all
	// folding ranges currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	// 
	// @since 3.18.0
	// @proposed
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}
func (t *FoldingRangeWorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias FoldingRangeWorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FoldingRangeWorkspaceClientCapabilities(test)
	
	return nil
}
// Value-object describing what options formatting should use.
type FormattingOptions struct {
	// Insert a newline character at the end of the file if one does not exist.
	// 
	// @since 3.15.0
	InsertFinalNewline bool `json:"insertFinalNewline,omitempty"`
	// Prefer spaces over tabs.
	InsertSpaces bool `json:"insertSpaces"`
	// Size of a tab in spaces.
	TabSize uint32 `json:"tabSize"`
	// Trim all newlines after the final newline at the end of the file.
	// 
	// @since 3.15.0
	TrimFinalNewlines bool `json:"trimFinalNewlines,omitempty"`
	// Trim trailing whitespace on a line.
	// 
	// @since 3.15.0
	TrimTrailingWhitespace bool `json:"trimTrailingWhitespace,omitempty"`
}
func (t *FormattingOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["insertSpaces"]; !exists {
		return fmt.Errorf("missing required field: insertSpaces")
	}
	if _, exists := m["tabSize"]; !exists {
		return fmt.Errorf("missing required field: tabSize")
	}
	type Alias FormattingOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FormattingOptions(test)
	
	return nil
}
// A diagnostic report with a full set of problems.
// 
// @since 3.17.0
type FullDocumentDiagnosticReport struct {
	// The actual items.
	Items []Diagnostic `json:"items"`
	// A full document diagnostic report.
	Kind string `json:"kind"`
	// An optional result id. If provided it will
	// be sent on the next diagnostic request for the
	// same document.
	ResultId string `json:"resultId,omitempty"`
}
func (t *FullDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias FullDocumentDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = FullDocumentDiagnosticReport(test)
	
	return nil
}
// General client capabilities.
// 
// @since 3.16.0
type GeneralClientCapabilities struct {
	// Client capabilities specific to the client's markdown parser.
	// 
	// @since 3.16.0
	Markdown *MarkdownClientCapabilities `json:"markdown,omitzero"`
	// The position encodings supported by the client. Client and server
	// have to agree on the same position encoding to ensure that offsets
	// (e.g. character position in a line) are interpreted the same on both
	// sides.
	// 
	// To keep the protocol backwards compatible the following applies: if
	// the value 'utf-16' is missing from the array of position encodings
	// servers can assume that the client supports UTF-16. UTF-16 is
	// therefore a mandatory encoding.
	// 
	// If omitted it defaults to ['utf-16'].
	// 
	// Implementation considerations: since the conversion from one encoding
	// into another requires the content of the file / line the conversion
	// is best done where the file is read which is usually on the server
	// side.
	// 
	// @since 3.17.0
	PositionEncodings []PositionEncodingKind `json:"positionEncodings,omitzero"`
	// Client capabilities specific to regular expressions.
	// 
	// @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitzero"`
	// Client capability that signals how the client
	// handles stale requests (e.g. a request
	// for which the client will not process the response
	// anymore since the information is outdated).
	// 
	// @since 3.17.0
	StaleRequestSupport *StaleRequestSupportOptions `json:"staleRequestSupport,omitzero"`
}
func (t *GeneralClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias GeneralClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = GeneralClientCapabilities(test)
	
	return nil
}
// The result of a hover request.
type Hover struct {
	// The hover's content
	Contents Or3[MarkupContent, MarkedString, []MarkedString] `json:"contents"`
	// An optional range inside the text document that is used to
	// visualize the hover, e.g. by changing the background color.
	Range *Range `json:"range,omitzero"`
}
func (t *Hover) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["contents"]; !exists {
		return fmt.Errorf("missing required field: contents")
	}
	type Alias Hover
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Hover(test)
	
	return nil
}

type HoverClientCapabilities struct {
	// Client supports the following content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"contentFormat,omitzero"`
	// Whether hover supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *HoverClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias HoverClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = HoverClientCapabilities(test)
	
	return nil
}
// Hover options.
type HoverOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *HoverOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias HoverOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = HoverOptions(test)
	
	return nil
}
// Parameters for a {@link HoverRequest}.
type HoverParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *HoverParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias HoverParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = HoverParams(test)
	
	return nil
}
// Registration options for a {@link HoverRequest}.
type HoverRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *HoverRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias HoverRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = HoverRegistrationOptions(test)
	
	return nil
}
// @since 3.6.0
type ImplementationClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `ImplementationRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports additional metadata in the form of definition links.
	// 
	// @since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}
func (t *ImplementationClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ImplementationClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ImplementationClientCapabilities(test)
	
	return nil
}

type ImplementationOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *ImplementationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ImplementationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ImplementationOptions(test)
	
	return nil
}

type ImplementationParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *ImplementationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias ImplementationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ImplementationParams(test)
	
	return nil
}

type ImplementationRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *ImplementationRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias ImplementationRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ImplementationRegistrationOptions(test)
	
	return nil
}
// The data type of the ResponseError if the
// initialize request fails.
type InitializeError struct {
	// Indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry"`
}
func (t *InitializeError) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["retry"]; !exists {
		return fmt.Errorf("missing required field: retry")
	}
	type Alias InitializeError
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InitializeError(test)
	
	return nil
}

type InitializeParams struct {
	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`
	// Information about the client
	// 
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitzero"`
	// User provided initialization options.
	InitializationOptions any `json:"initializationOptions,omitempty"`
	// The locale the client is currently showing the user interface
	// in. This must not necessarily be the locale of the operating
	// system.
	// 
	// Uses IETF language tags as the value's syntax
	// (See https://en.wikipedia.org/wiki/IETF_language_tag)
	// 
	// @since 3.16.0
	Locale string `json:"locale,omitempty"`
	// The process Id of the parent process that started
	// the server.
	// 
	// Is `null` if the process has not been started by another process.
	// If the parent process is not alive then the server should exit.
	ProcessId *int32 `json:"processId"`
	// The rootPath of the workspace. Is null
	// if no folder is open.
	// 
	// @deprecated in favour of rootUri.
	RootPath **string `json:"rootPath,omitzero"`
	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	// 
	// @deprecated in favour of workspaceFolders.
	RootUri *DocumentUri `json:"rootUri"`
	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitzero"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
	// The workspace folders configured in the client when the server starts.
	// 
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	// 
	// @since 3.6.0
	WorkspaceFolders *[]WorkspaceFolder `json:"workspaceFolders,omitzero"`
}
func (t *InitializeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["capabilities"]; !exists {
		return fmt.Errorf("missing required field: capabilities")
	}
	if _, exists := m["processId"]; !exists {
		return fmt.Errorf("missing required field: processId")
	}
	if _, exists := m["rootUri"]; !exists {
		return fmt.Errorf("missing required field: rootUri")
	}
	type Alias InitializeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InitializeParams(test)
	
	return nil
}
// The result returned from an initialize request.
type InitializeResult struct {
	// The capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`
	// Information about the server.
	// 
	// @since 3.15.0
	ServerInfo *ServerInfo `json:"serverInfo,omitzero"`
}
func (t *InitializeResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["capabilities"]; !exists {
		return fmt.Errorf("missing required field: capabilities")
	}
	type Alias InitializeResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InitializeResult(test)
	
	return nil
}

type InitializedParams LSPObject
// Inlay hint information.
// 
// @since 3.17.0
type InlayHint struct {
	// A data entry field that is preserved on an inlay hint between
	// a `textDocument/inlayHint` and a `inlayHint/resolve` request.
	Data any `json:"data,omitempty"`
	// The kind of this hint. Can be omitted in which case the client
	// should fall back to a reasonable default.
	Kind *InlayHintKind `json:"kind,omitzero"`
	// The label of this hint. A human readable string or an array of
	// InlayHintLabelPart label parts.
	// 
	// *Note* that neither the string nor the label part can be empty.
	Label Or2[string, []InlayHintLabelPart] `json:"label"`
	// Render padding before the hint.
	// 
	// Note: Padding should use the editor's background color, not the
	// background color of the hint itself. That means padding can be used
	// to visually align/separate an inlay hint.
	PaddingLeft bool `json:"paddingLeft,omitempty"`
	// Render padding after the hint.
	// 
	// Note: Padding should use the editor's background color, not the
	// background color of the hint itself. That means padding can be used
	// to visually align/separate an inlay hint.
	PaddingRight bool `json:"paddingRight,omitempty"`
	// The position of this hint.
	// 
	// If multiple hints have the same position, they will be shown in the order
	// they appear in the response.
	Position Position `json:"position"`
	// Optional text edits that are performed when accepting this inlay hint.
	// 
	// *Note* that edits are expected to change the document so that the inlay
	// hint (or its nearest variant) is now part of the document and the inlay
	// hint itself is now obsolete.
	TextEdits []TextEdit `json:"textEdits,omitzero"`
	// The tooltip text when you hover over this item.
	Tooltip *Or2[string, MarkupContent] `json:"tooltip,omitzero"`
}
func (t *InlayHint) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["label"]; !exists {
		return fmt.Errorf("missing required field: label")
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	type Alias InlayHint
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHint(test)
	
	return nil
}
// Inlay hint client capabilities.
// 
// @since 3.17.0
type InlayHintClientCapabilities struct {
	// Whether inlay hints support dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Indicates which properties a client can resolve lazily on an inlay
	// hint.
	ResolveSupport *ClientInlayHintResolveOptions `json:"resolveSupport,omitzero"`
}
func (t *InlayHintClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlayHintClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHintClientCapabilities(test)
	
	return nil
}
// An inlay hint label part allows for interactive and composite labels
// of inlay hints.
// 
// @since 3.17.0
type InlayHintLabelPart struct {
	// An optional command for this label part.
	// 
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	Command *Command `json:"command,omitzero"`
	// An optional source code location that represents this
	// label part.
	// 
	// The editor will use this location for the hover and for code navigation
	// features: This part will become a clickable link that resolves to the
	// definition of the symbol at the given location (not necessarily the
	// location itself), it shows the hover that shows at the given location,
	// and it shows a context menu with further code navigation commands.
	// 
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	Location *Location `json:"location,omitzero"`
	// The tooltip text when you hover over this label part. Depending on
	// the client capability `inlayHint.resolveSupport` clients might resolve
	// this property late using the resolve request.
	Tooltip *Or2[string, MarkupContent] `json:"tooltip,omitzero"`
	// The value of this label part.
	Value string `json:"value"`
}
func (t *InlayHintLabelPart) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias InlayHintLabelPart
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHintLabelPart(test)
	
	return nil
}
// Inlay hint options used during static registration.
// 
// @since 3.17.0
type InlayHintOptions struct {
	// The server provides support to resolve additional
	// information for an inlay hint item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *InlayHintOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlayHintOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHintOptions(test)
	
	return nil
}
// A parameter literal used in inlay hint requests.
// 
// @since 3.17.0
type InlayHintParams struct {
	// The document range for which inlay hints should be computed.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *InlayHintParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias InlayHintParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHintParams(test)
	
	return nil
}
// Inlay hint options used during static or dynamic registration.
// 
// @since 3.17.0
type InlayHintRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// The server provides support to resolve additional
	// information for an inlay hint item.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *InlayHintRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias InlayHintRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHintRegistrationOptions(test)
	
	return nil
}
// Client workspace capabilities specific to inlay hints.
// 
// @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	// 
	// Note that this event is global and will force the client to refresh all
	// inlay hints currently shown. It should be used with absolute care and
	// is useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}
func (t *InlayHintWorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlayHintWorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlayHintWorkspaceClientCapabilities(test)
	
	return nil
}
// Client capabilities specific to inline completions.
// 
// @since 3.18.0
// @proposed
type InlineCompletionClientCapabilities struct {
	// Whether implementation supports dynamic registration for inline completion providers.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *InlineCompletionClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlineCompletionClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionClientCapabilities(test)
	
	return nil
}
// Provides information about the context in which an inline completion was requested.
// 
// @since 3.18.0
// @proposed
type InlineCompletionContext struct {
	// Provides information about the currently selected item in the autocomplete widget if it is visible.
	SelectedCompletionInfo *SelectedCompletionInfo `json:"selectedCompletionInfo,omitzero"`
	// Describes how the inline completion was triggered.
	TriggerKind InlineCompletionTriggerKind `json:"triggerKind"`
}
func (t *InlineCompletionContext) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["triggerKind"]; !exists {
		return fmt.Errorf("missing required field: triggerKind")
	}
	type Alias InlineCompletionContext
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionContext(test)
	
	return nil
}
// An inline completion item represents a text snippet that is proposed inline to complete text that is being typed.
// 
// @since 3.18.0
// @proposed
type InlineCompletionItem struct {
	// An optional {@link Command} that is executed *after* inserting this completion.
	Command *Command `json:"command,omitzero"`
	// A text that is used to decide if this inline completion should be shown. When `falsy` the {@link InlineCompletionItem.insertText} is used.
	FilterText string `json:"filterText,omitempty"`
	// The text to replace the range with. Must be set.
	InsertText Or2[string, StringValue] `json:"insertText"`
	// The range to replace. Must begin and end on the same line.
	Range *Range `json:"range,omitzero"`
}
func (t *InlineCompletionItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["insertText"]; !exists {
		return fmt.Errorf("missing required field: insertText")
	}
	type Alias InlineCompletionItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionItem(test)
	
	return nil
}
// Represents a collection of {@link InlineCompletionItem inline completion items} to be presented in the editor.
// 
// @since 3.18.0
// @proposed
type InlineCompletionList struct {
	// The inline completion items
	Items []InlineCompletionItem `json:"items"`
}
func (t *InlineCompletionList) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	type Alias InlineCompletionList
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionList(test)
	
	return nil
}
// Inline completion options used during static registration.
// 
// @since 3.18.0
// @proposed
type InlineCompletionOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *InlineCompletionOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlineCompletionOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionOptions(test)
	
	return nil
}
// A parameter literal used in inline completion requests.
// 
// @since 3.18.0
// @proposed
type InlineCompletionParams struct {
	// Additional information about the context in which inline completions were
	// requested.
	Context InlineCompletionContext `json:"context"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *InlineCompletionParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["context"]; !exists {
		return fmt.Errorf("missing required field: context")
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias InlineCompletionParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionParams(test)
	
	return nil
}
// Inline completion options used during static or dynamic registration.
// 
// @since 3.18.0
// @proposed
type InlineCompletionRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *InlineCompletionRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias InlineCompletionRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineCompletionRegistrationOptions(test)
	
	return nil
}
// Client capabilities specific to inline values.
// 
// @since 3.17.0
type InlineValueClientCapabilities struct {
	// Whether implementation supports dynamic registration for inline value providers.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *InlineValueClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlineValueClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueClientCapabilities(test)
	
	return nil
}
// @since 3.17.0
type InlineValueContext struct {
	// The stack frame (as a DAP Id) where the execution has stopped.
	FrameId int32 `json:"frameId"`
	// The document range where execution has stopped.
	// Typically the end position of the range denotes the line where the inline values are shown.
	StoppedLocation Range `json:"stoppedLocation"`
}
func (t *InlineValueContext) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["frameId"]; !exists {
		return fmt.Errorf("missing required field: frameId")
	}
	if _, exists := m["stoppedLocation"]; !exists {
		return fmt.Errorf("missing required field: stoppedLocation")
	}
	type Alias InlineValueContext
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueContext(test)
	
	return nil
}
// Provide an inline value through an expression evaluation.
// If only a range is specified, the expression will be extracted from the underlying document.
// An optional expression can be used to override the extracted expression.
// 
// @since 3.17.0
type InlineValueEvaluatableExpression struct {
	// If specified the expression overrides the extracted expression.
	Expression string `json:"expression,omitempty"`
	// The document range for which the inline value applies.
	// The range is used to extract the evaluatable expression from the underlying document.
	Range Range `json:"range"`
}
func (t *InlineValueEvaluatableExpression) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias InlineValueEvaluatableExpression
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueEvaluatableExpression(test)
	
	return nil
}
// Inline value options used during static registration.
// 
// @since 3.17.0
type InlineValueOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *InlineValueOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlineValueOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueOptions(test)
	
	return nil
}
// A parameter literal used in inline value requests.
// 
// @since 3.17.0
type InlineValueParams struct {
	// Additional information about the context in which inline values were
	// requested.
	Context InlineValueContext `json:"context"`
	// The document range for which inline values should be computed.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *InlineValueParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["context"]; !exists {
		return fmt.Errorf("missing required field: context")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias InlineValueParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueParams(test)
	
	return nil
}
// Inline value options used during static or dynamic registration.
// 
// @since 3.17.0
type InlineValueRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *InlineValueRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias InlineValueRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueRegistrationOptions(test)
	
	return nil
}
// Provide inline value as text.
// 
// @since 3.17.0
type InlineValueText struct {
	// The document range for which the inline value applies.
	Range Range `json:"range"`
	// The text of the inline value.
	Text string `json:"text"`
}
func (t *InlineValueText) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["text"]; !exists {
		return fmt.Errorf("missing required field: text")
	}
	type Alias InlineValueText
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueText(test)
	
	return nil
}
// Provide inline value through a variable lookup.
// If only a range is specified, the variable name will be extracted from the underlying document.
// An optional variable name can be used to override the extracted name.
// 
// @since 3.17.0
type InlineValueVariableLookup struct {
	// How to perform the lookup.
	CaseSensitiveLookup bool `json:"caseSensitiveLookup"`
	// The document range for which the inline value applies.
	// The range is used to extract the variable name from the underlying document.
	Range Range `json:"range"`
	// If specified the name of the variable to look up.
	VariableName string `json:"variableName,omitempty"`
}
func (t *InlineValueVariableLookup) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["caseSensitiveLookup"]; !exists {
		return fmt.Errorf("missing required field: caseSensitiveLookup")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias InlineValueVariableLookup
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueVariableLookup(test)
	
	return nil
}
// Client workspace capabilities specific to inline values.
// 
// @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	// 
	// Note that this event is global and will force the client to refresh all
	// inline values currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}
func (t *InlineValueWorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias InlineValueWorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InlineValueWorkspaceClientCapabilities(test)
	
	return nil
}
// A special text edit to provide an insert and a replace operation.
// 
// @since 3.16.0
type InsertReplaceEdit struct {
	// The range if the insert is requested
	Insert Range `json:"insert"`
	// The string to be inserted.
	NewText string `json:"newText"`
	// The range if the replace is requested.
	Replace Range `json:"replace"`
}
func (t *InsertReplaceEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["insert"]; !exists {
		return fmt.Errorf("missing required field: insert")
	}
	if _, exists := m["newText"]; !exists {
		return fmt.Errorf("missing required field: newText")
	}
	if _, exists := m["replace"]; !exists {
		return fmt.Errorf("missing required field: replace")
	}
	type Alias InsertReplaceEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = InsertReplaceEdit(test)
	
	return nil
}
// Client capabilities for the linked editing range request.
// 
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *LinkedEditingRangeClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias LinkedEditingRangeClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LinkedEditingRangeClientCapabilities(test)
	
	return nil
}

type LinkedEditingRangeOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *LinkedEditingRangeOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias LinkedEditingRangeOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LinkedEditingRangeOptions(test)
	
	return nil
}

type LinkedEditingRangeParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *LinkedEditingRangeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias LinkedEditingRangeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LinkedEditingRangeParams(test)
	
	return nil
}

type LinkedEditingRangeRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *LinkedEditingRangeRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias LinkedEditingRangeRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LinkedEditingRangeRegistrationOptions(test)
	
	return nil
}
// The result of a linked editing range request.
// 
// @since 3.16.0
type LinkedEditingRanges struct {
	// A list of ranges that can be edited together. The ranges must have
	// identical length and contain identical text content. The ranges cannot overlap.
	Ranges []Range `json:"ranges"`
	// An optional word pattern (regular expression) that describes valid contents for
	// the given ranges. If no pattern is provided, the client configuration's word
	// pattern will be used.
	WordPattern string `json:"wordPattern,omitempty"`
}
func (t *LinkedEditingRanges) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["ranges"]; !exists {
		return fmt.Errorf("missing required field: ranges")
	}
	type Alias LinkedEditingRanges
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LinkedEditingRanges(test)
	
	return nil
}
// Represents a location inside a resource, such as a line
// inside a text file.
type Location struct {

	Range Range `json:"range"`

	Uri DocumentUri `json:"uri"`
}
func (t *Location) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias Location
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Location(test)
	
	return nil
}
// Represents the connection of two locations. Provides additional metadata over normal {@link Location locations},
// including an origin range.
type LocationLink struct {
	// Span of the origin of this link.
	// 
	// Used as the underlined span for mouse interaction. Defaults to the word range at
	// the definition position.
	OriginSelectionRange *Range `json:"originSelectionRange,omitzero"`
	// The full target range of this link. If the target for example is a symbol then target range is the
	// range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range `json:"targetRange"`
	// The range that should be selected and revealed when this link is being followed, e.g the name of a function.
	// Must be contained by the `targetRange`. See also `DocumentSymbol#range`
	TargetSelectionRange Range `json:"targetSelectionRange"`
	// The target resource identifier of this link.
	TargetUri DocumentUri `json:"targetUri"`
}
func (t *LocationLink) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["targetRange"]; !exists {
		return fmt.Errorf("missing required field: targetRange")
	}
	if _, exists := m["targetSelectionRange"]; !exists {
		return fmt.Errorf("missing required field: targetSelectionRange")
	}
	if _, exists := m["targetUri"]; !exists {
		return fmt.Errorf("missing required field: targetUri")
	}
	type Alias LocationLink
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LocationLink(test)
	
	return nil
}
// Location with only uri and does not include range.
// 
// @since 3.18.0
type LocationUriOnly struct {

	Uri DocumentUri `json:"uri"`
}
func (t *LocationUriOnly) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias LocationUriOnly
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LocationUriOnly(test)
	
	return nil
}
// The log message parameters.
type LogMessageParams struct {
	// The actual message.
	Message string `json:"message"`
	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`
}
func (t *LogMessageParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["message"]; !exists {
		return fmt.Errorf("missing required field: message")
	}
	if _, exists := m["type"]; !exists {
		return fmt.Errorf("missing required field: type")
	}
	type Alias LogMessageParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LogMessageParams(test)
	
	return nil
}

type LogTraceParams struct {

	Message string `json:"message"`

	Verbose string `json:"verbose,omitempty"`
}
func (t *LogTraceParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["message"]; !exists {
		return fmt.Errorf("missing required field: message")
	}
	type Alias LogTraceParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = LogTraceParams(test)
	
	return nil
}
// Client capabilities specific to the used markdown parser.
// 
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// A list of HTML tags that the client allows / supports in
	// Markdown.
	// 
	// @since 3.17.0
	AllowedTags []string `json:"allowedTags,omitzero"`
	// The name of the parser.
	Parser string `json:"parser"`
	// The version of the parser.
	Version string `json:"version,omitempty"`
}
func (t *MarkdownClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["parser"]; !exists {
		return fmt.Errorf("missing required field: parser")
	}
	type Alias MarkdownClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MarkdownClientCapabilities(test)
	
	return nil
}
// @since 3.18.0
// @deprecated use MarkupContent instead.
type MarkedStringWithLanguage struct {

	Language string `json:"language"`

	Value string `json:"value"`
}
func (t *MarkedStringWithLanguage) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["language"]; !exists {
		return fmt.Errorf("missing required field: language")
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias MarkedStringWithLanguage
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MarkedStringWithLanguage(test)
	
	return nil
}
// A `MarkupContent` literal represents a string value which content is interpreted base on its
// kind flag. Currently the protocol supports `plaintext` and `markdown` as markup kinds.
// 
// If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues.
// See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
// 
// Here is an example how such a string can be constructed using JavaScript / TypeScript:
// ```ts
// let markdown: MarkdownContent = {
//  kind: MarkupKind.Markdown,
//  value: [
//    '# Header',
//    'Some text',
//    '```typescript',
//    'someCode();',
//    '```'
//  ].join('\n')
// };
// ```
// 
// *Please Note* that clients might sanitize the return markdown. A client could decide to
// remove HTML from the markdown to avoid script execution.
type MarkupContent struct {
	// The type of the Markup
	Kind MarkupKind `json:"kind"`
	// The content itself
	Value string `json:"value"`
}
func (t *MarkupContent) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias MarkupContent
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MarkupContent(test)
	
	return nil
}

type MessageActionItem struct {
	// A short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}
func (t *MessageActionItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["title"]; !exists {
		return fmt.Errorf("missing required field: title")
	}
	type Alias MessageActionItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MessageActionItem(test)
	
	return nil
}
// Moniker definition to match LSIF 0.5 moniker definition.
// 
// @since 3.16.0
type Moniker struct {
	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`
	// The moniker kind if known.
	Kind *MonikerKind `json:"kind,omitzero"`
	// The scheme of the moniker. For example tsc or .Net
	Scheme string `json:"scheme"`
	// The scope in which the moniker is unique
	Unique UniquenessLevel `json:"unique"`
}
func (t *Moniker) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["identifier"]; !exists {
		return fmt.Errorf("missing required field: identifier")
	}
	if _, exists := m["scheme"]; !exists {
		return fmt.Errorf("missing required field: scheme")
	}
	if _, exists := m["unique"]; !exists {
		return fmt.Errorf("missing required field: unique")
	}
	type Alias Moniker
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Moniker(test)
	
	return nil
}
// Client capabilities specific to the moniker request.
// 
// @since 3.16.0
type MonikerClientCapabilities struct {
	// Whether moniker supports dynamic registration. If this is set to `true`
	// the client supports the new `MonikerRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *MonikerClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias MonikerClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MonikerClientCapabilities(test)
	
	return nil
}

type MonikerOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *MonikerOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias MonikerOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MonikerOptions(test)
	
	return nil
}

type MonikerParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *MonikerParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias MonikerParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MonikerParams(test)
	
	return nil
}

type MonikerRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *MonikerRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias MonikerRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = MonikerRegistrationOptions(test)
	
	return nil
}
// A notebook cell.
// 
// A cell's document URI must be unique across ALL notebook
// cells and can therefore be used to uniquely identify a
// notebook cell or the cell's text document.
// 
// @since 3.17.0
type NotebookCell struct {
	// The URI of the cell's text document
	// content.
	Document DocumentUri `json:"document"`
	// Additional execution summary information
	// if supported by the client.
	ExecutionSummary *ExecutionSummary `json:"executionSummary,omitzero"`
	// The cell's kind
	Kind NotebookCellKind `json:"kind"`
	// Additional metadata stored with the cell.
	// 
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata,omitzero"`
}
func (t *NotebookCell) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["document"]; !exists {
		return fmt.Errorf("missing required field: document")
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias NotebookCell
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookCell(test)
	
	return nil
}
// A change describing how to move a `NotebookCell`
// array from state S to S'.
// 
// @since 3.17.0
type NotebookCellArrayChange struct {
	// The new cells, if any
	Cells []NotebookCell `json:"cells,omitzero"`
	// The deleted cells
	DeleteCount uint32 `json:"deleteCount"`
	// The start oftest of the cell that changed.
	Start uint32 `json:"start"`
}
func (t *NotebookCellArrayChange) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["deleteCount"]; !exists {
		return fmt.Errorf("missing required field: deleteCount")
	}
	if _, exists := m["start"]; !exists {
		return fmt.Errorf("missing required field: start")
	}
	type Alias NotebookCellArrayChange
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookCellArrayChange(test)
	
	return nil
}
// @since 3.18.0
type NotebookCellLanguage struct {

	Language string `json:"language"`
}
func (t *NotebookCellLanguage) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["language"]; !exists {
		return fmt.Errorf("missing required field: language")
	}
	type Alias NotebookCellLanguage
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookCellLanguage(test)
	
	return nil
}
// A notebook cell text document filter denotes a cell text
// document by different properties.
// 
// @since 3.17.0
type NotebookCellTextDocumentFilter struct {
	// A language id like `python`.
	// 
	// Will be matched against the language id of the
	// notebook cell document. '*' matches every language.
	Language string `json:"language,omitempty"`
	// A filter that matches against the notebook
	// containing the notebook cell. If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	Notebook Or2[string, NotebookDocumentFilter] `json:"notebook"`
}
func (t *NotebookCellTextDocumentFilter) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["notebook"]; !exists {
		return fmt.Errorf("missing required field: notebook")
	}
	type Alias NotebookCellTextDocumentFilter
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookCellTextDocumentFilter(test)
	
	return nil
}
// A notebook document.
// 
// @since 3.17.0
type NotebookDocument struct {
	// The cells of a notebook.
	Cells []NotebookCell `json:"cells"`
	// Additional metadata stored with the notebook
	// document.
	// 
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata,omitzero"`
	// The type of the notebook.
	NotebookType string `json:"notebookType"`
	// The notebook document's uri.
	Uri URI `json:"uri"`
	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int32 `json:"version"`
}
func (t *NotebookDocument) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["cells"]; !exists {
		return fmt.Errorf("missing required field: cells")
	}
	if _, exists := m["notebookType"]; !exists {
		return fmt.Errorf("missing required field: notebookType")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias NotebookDocument
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocument(test)
	
	return nil
}
// Structural changes to cells in a notebook document.
// 
// @since 3.18.0
type NotebookDocumentCellChangeStructure struct {
	// The change to the cell array.
	Array NotebookCellArrayChange `json:"array"`
	// Additional closed cell text documents.
	DidClose []TextDocumentIdentifier `json:"didClose,omitzero"`
	// Additional opened cell text documents.
	DidOpen []TextDocumentItem `json:"didOpen,omitzero"`
}
func (t *NotebookDocumentCellChangeStructure) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["array"]; !exists {
		return fmt.Errorf("missing required field: array")
	}
	type Alias NotebookDocumentCellChangeStructure
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentCellChangeStructure(test)
	
	return nil
}
// Cell changes to a notebook document.
// 
// @since 3.18.0
type NotebookDocumentCellChanges struct {
	// Changes to notebook cells properties like its
	// kind, execution summary or metadata.
	Data []NotebookCell `json:"data,omitzero"`
	// Changes to the cell structure to add or
	// remove cells.
	Structure *NotebookDocumentCellChangeStructure `json:"structure,omitzero"`
	// Changes to the text content of notebook cells.
	TextContent []NotebookDocumentCellContentChanges `json:"textContent,omitzero"`
}
func (t *NotebookDocumentCellChanges) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias NotebookDocumentCellChanges
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentCellChanges(test)
	
	return nil
}
// Content changes to a cell in a notebook document.
// 
// @since 3.18.0
type NotebookDocumentCellContentChanges struct {

	Changes []TextDocumentContentChangeEvent `json:"changes"`

	Document VersionedTextDocumentIdentifier `json:"document"`
}
func (t *NotebookDocumentCellContentChanges) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["changes"]; !exists {
		return fmt.Errorf("missing required field: changes")
	}
	if _, exists := m["document"]; !exists {
		return fmt.Errorf("missing required field: document")
	}
	type Alias NotebookDocumentCellContentChanges
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentCellContentChanges(test)
	
	return nil
}
// A change event for a notebook document.
// 
// @since 3.17.0
type NotebookDocumentChangeEvent struct {
	// Changes to cells
	Cells *NotebookDocumentCellChanges `json:"cells,omitzero"`
	// The changed meta data if any.
	// 
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata,omitzero"`
}
func (t *NotebookDocumentChangeEvent) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias NotebookDocumentChangeEvent
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentChangeEvent(test)
	
	return nil
}
// Capabilities specific to the notebook document support.
// 
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
	// Capabilities specific to notebook document synchronization
	// 
	// @since 3.17.0
	Synchronization NotebookDocumentSyncClientCapabilities `json:"synchronization"`
}
func (t *NotebookDocumentClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["synchronization"]; !exists {
		return fmt.Errorf("missing required field: synchronization")
	}
	type Alias NotebookDocumentClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentClientCapabilities(test)
	
	return nil
}
// A notebook document filter where `notebookType` is required field.
// 
// @since 3.18.0
type NotebookDocumentFilterNotebookType struct {
	// The type of the enclosing notebook.
	NotebookType string `json:"notebookType"`
	// A glob pattern.
	Pattern *GlobPattern `json:"pattern,omitzero"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
}
func (t *NotebookDocumentFilterNotebookType) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["notebookType"]; !exists {
		return fmt.Errorf("missing required field: notebookType")
	}
	type Alias NotebookDocumentFilterNotebookType
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentFilterNotebookType(test)
	
	return nil
}
// A notebook document filter where `pattern` is required field.
// 
// @since 3.18.0
type NotebookDocumentFilterPattern struct {
	// The type of the enclosing notebook.
	NotebookType string `json:"notebookType,omitempty"`
	// A glob pattern.
	Pattern GlobPattern `json:"pattern"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
}
func (t *NotebookDocumentFilterPattern) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["pattern"]; !exists {
		return fmt.Errorf("missing required field: pattern")
	}
	type Alias NotebookDocumentFilterPattern
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentFilterPattern(test)
	
	return nil
}
// A notebook document filter where `scheme` is required field.
// 
// @since 3.18.0
type NotebookDocumentFilterScheme struct {
	// The type of the enclosing notebook.
	NotebookType string `json:"notebookType,omitempty"`
	// A glob pattern.
	Pattern *GlobPattern `json:"pattern,omitzero"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme"`
}
func (t *NotebookDocumentFilterScheme) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["scheme"]; !exists {
		return fmt.Errorf("missing required field: scheme")
	}
	type Alias NotebookDocumentFilterScheme
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentFilterScheme(test)
	
	return nil
}
// @since 3.18.0
type NotebookDocumentFilterWithCells struct {
	// The cells of the matching notebook to be synced.
	Cells []NotebookCellLanguage `json:"cells"`
	// The notebook to be synced If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	Notebook *Or2[string, NotebookDocumentFilter] `json:"notebook,omitzero"`
}
func (t *NotebookDocumentFilterWithCells) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["cells"]; !exists {
		return fmt.Errorf("missing required field: cells")
	}
	type Alias NotebookDocumentFilterWithCells
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentFilterWithCells(test)
	
	return nil
}
// @since 3.18.0
type NotebookDocumentFilterWithNotebook struct {
	// The cells of the matching notebook to be synced.
	Cells []NotebookCellLanguage `json:"cells,omitzero"`
	// The notebook to be synced If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	Notebook Or2[string, NotebookDocumentFilter] `json:"notebook"`
}
func (t *NotebookDocumentFilterWithNotebook) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["notebook"]; !exists {
		return fmt.Errorf("missing required field: notebook")
	}
	type Alias NotebookDocumentFilterWithNotebook
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentFilterWithNotebook(test)
	
	return nil
}
// A literal to identify a notebook document in the client.
// 
// @since 3.17.0
type NotebookDocumentIdentifier struct {
	// The notebook document's uri.
	Uri URI `json:"uri"`
}
func (t *NotebookDocumentIdentifier) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias NotebookDocumentIdentifier
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentIdentifier(test)
	
	return nil
}
// Notebook specific client capabilities.
// 
// @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is
	// set to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports sending execution summary data per cell.
	ExecutionSummarySupport bool `json:"executionSummarySupport,omitempty"`
}
func (t *NotebookDocumentSyncClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias NotebookDocumentSyncClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentSyncClientCapabilities(test)
	
	return nil
}
// Options specific to a notebook plus its cells
// to be synced to the server.
// 
// If a selector provides a notebook document
// filter but no cell selector all cells of a
// matching notebook document will be synced.
// 
// If a selector provides no notebook document
// filter but only a cell selector all notebook
// document that contain at least one matching
// cell will be synced.
// 
// @since 3.17.0
type NotebookDocumentSyncOptions struct {
	// The notebooks to be synced
	NotebookSelector []Or2[NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells] `json:"notebookSelector"`
	// Whether save notification should be forwarded to
	// the server. Will only be honored if mode === `notebook`.
	Save bool `json:"save,omitempty"`
}
func (t *NotebookDocumentSyncOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["notebookSelector"]; !exists {
		return fmt.Errorf("missing required field: notebookSelector")
	}
	type Alias NotebookDocumentSyncOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentSyncOptions(test)
	
	return nil
}
// Registration options specific to a notebook.
// 
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// The notebooks to be synced
	NotebookSelector []Or2[NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells] `json:"notebookSelector"`
	// Whether save notification should be forwarded to
	// the server. Will only be honored if mode === `notebook`.
	Save bool `json:"save,omitempty"`
}
func (t *NotebookDocumentSyncRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["notebookSelector"]; !exists {
		return fmt.Errorf("missing required field: notebookSelector")
	}
	type Alias NotebookDocumentSyncRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = NotebookDocumentSyncRegistrationOptions(test)
	
	return nil
}
// A text document identifier to optionally denote a specific version of a text document.
type OptionalVersionedTextDocumentIdentifier struct {
	// The text document's uri.
	Uri DocumentUri `json:"uri"`
	// The version number of this document. If a versioned text document identifier
	// is sent from the server to the client and the file is not open in the editor
	// (the server has not received an open notification before) the server can send
	// `null` to indicate that the version is unknown and the content on disk is the
	// truth (as specified with document content ownership).
	Version *int32 `json:"version"`
}
func (t *OptionalVersionedTextDocumentIdentifier) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias OptionalVersionedTextDocumentIdentifier
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = OptionalVersionedTextDocumentIdentifier(test)
	
	return nil
}
// Represents a parameter of a callable-signature. A parameter can
// have a label and a doc-comment.
type ParameterInformation struct {
	// The human-readable doc-comment of this parameter. Will be shown
	// in the UI but can be omitted.
	Documentation *Or2[string, MarkupContent] `json:"documentation,omitzero"`
	// The label of this parameter information.
	// 
	// Either a string or an inclusive start and exclusive end offsets within its containing
	// signature label. (see SignatureInformation.label). The offsets are based on a UTF-16
	// string representation as `Position` and `Range` does.
	// 
	// To avoid ambiguities a server should use the [start, end] offset value instead of using
	// a substring. Whether a client support this is controlled via `labelOffsetSupport` client
	// capability.
	// 
	// *Note*: a label of type string should be a substring of its containing signature label.
	// Its intended use case is to highlight the parameter label part in the `SignatureInformation.label`.
	Label Or2[string, Tuple[uint32, uint32]] `json:"label"`
}
func (t *ParameterInformation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["label"]; !exists {
		return fmt.Errorf("missing required field: label")
	}
	type Alias ParameterInformation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ParameterInformation(test)
	
	return nil
}

type PartialResultParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
}
func (t *PartialResultParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias PartialResultParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PartialResultParams(test)
	
	return nil
}
// Position in a text document expressed as zero-based line and character
// offset. Prior to 3.17 the offsets were always based on a UTF-16 string
// representation. So a string of the form `a𐐀b` the character offset of the
// character `a` is 0, the character offset of `𐐀` is 1 and the character
// offset of b is 3 since `𐐀` is represented using two code units in UTF-16.
// Since 3.17 clients and servers can agree on a different string encoding
// representation (e.g. UTF-8). The client announces it's supported encoding
// via the client capability [`general.positionEncodings`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#clientCapabilities).
// The value is an array of position encodings the client supports, with
// decreasing preference (e.g. the encoding at index `0` is the most preferred
// one). To stay backwards compatible the only mandatory encoding is UTF-16
// represented via the string `utf-16`. The server can pick one of the
// encodings offered by the client and signals that encoding back to the
// client via the initialize result's property
// [`capabilities.positionEncoding`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#serverCapabilities). If the string value
// `utf-16` is missing from the client's capability `general.positionEncodings`
// servers can safely assume that the client supports UTF-16. If the server
// omits the position encoding in its initialize result the encoding defaults
// to the string value `utf-16`. Implementation considerations: since the
// conversion from one encoding into another requires the content of the
// file / line the conversion is best done where the file is read which is
// usually on the server side.
// 
// Positions are line end character agnostic. So you can not specify a position
// that denotes `\r|\n` or `\n|` where `|` represents the character offset.
// 
// @since 3.17.0 - support for negotiated position encoding.
type Position struct {
	// Character offset on a line in a document (zero-based).
	// 
	// The meaning of this offset is determined by the negotiated
	// `PositionEncodingKind`.
	Character uint32 `json:"character"`
	// Line position in a document (zero-based).
	Line uint32 `json:"line"`
}
func (t *Position) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["character"]; !exists {
		return fmt.Errorf("missing required field: character")
	}
	if _, exists := m["line"]; !exists {
		return fmt.Errorf("missing required field: line")
	}
	type Alias Position
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Position(test)
	
	return nil
}
// @since 3.18.0
type PrepareRenameDefaultBehavior struct {

	DefaultBehavior bool `json:"defaultBehavior"`
}
func (t *PrepareRenameDefaultBehavior) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["defaultBehavior"]; !exists {
		return fmt.Errorf("missing required field: defaultBehavior")
	}
	type Alias PrepareRenameDefaultBehavior
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PrepareRenameDefaultBehavior(test)
	
	return nil
}

type PrepareRenameParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *PrepareRenameParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias PrepareRenameParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PrepareRenameParams(test)
	
	return nil
}
// @since 3.18.0
type PrepareRenamePlaceholder struct {

	Placeholder string `json:"placeholder"`

	Range Range `json:"range"`
}
func (t *PrepareRenamePlaceholder) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["placeholder"]; !exists {
		return fmt.Errorf("missing required field: placeholder")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias PrepareRenamePlaceholder
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PrepareRenamePlaceholder(test)
	
	return nil
}
// A previous result id in a workspace pull request.
// 
// @since 3.17.0
type PreviousResultId struct {
	// The URI for which the client knowns a
	// result id.
	Uri DocumentUri `json:"uri"`
	// The value of the previous result id.
	Value string `json:"value"`
}
func (t *PreviousResultId) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias PreviousResultId
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PreviousResultId(test)
	
	return nil
}

type ProgressParams struct {
	// The progress token provided by the client or server.
	Token ProgressToken `json:"token"`
	// The progress data.
	Value any `json:"value"`
}
func (t *ProgressParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["token"]; !exists {
		return fmt.Errorf("missing required field: token")
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias ProgressParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ProgressParams(test)
	
	return nil
}
// The publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {
	// Client supports a codeDescription property
	// 
	// @since 3.16.0
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`
	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	// 
	// @since 3.16.0
	DataSupport bool `json:"dataSupport,omitempty"`
	// Whether the clients accepts diagnostics with related information.
	RelatedInformation bool `json:"relatedInformation,omitempty"`
	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.15.0
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitzero"`
	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.
	// 
	// @since 3.15.0
	VersionSupport bool `json:"versionSupport,omitempty"`
}
func (t *PublishDiagnosticsClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias PublishDiagnosticsClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PublishDiagnosticsClientCapabilities(test)
	
	return nil
}
// The publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {
	// An array of diagnostic information items.
	Diagnostics []Diagnostic `json:"diagnostics"`
	// The URI for which diagnostic information is reported.
	Uri DocumentUri `json:"uri"`
	// Optional the version number of the document the diagnostics are published for.
	// 
	// @since 3.15.0
	Version int32 `json:"version,omitempty"`
}
func (t *PublishDiagnosticsParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["diagnostics"]; !exists {
		return fmt.Errorf("missing required field: diagnostics")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias PublishDiagnosticsParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = PublishDiagnosticsParams(test)
	
	return nil
}
// A range in a text document expressed as (zero-based) start and end positions.
// 
// If you want to specify a range that contains a line including the line ending
// character(s) then use an end position denoting the start of the next line.
// For example:
// ```ts
// {
//     start: { line: 5, character: 23 }
//     end : { line 6, character : 0 }
// }
// ```
type Range struct {
	// The range's end position.
	End Position `json:"end"`
	// The range's start position.
	Start Position `json:"start"`
}
func (t *Range) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["end"]; !exists {
		return fmt.Errorf("missing required field: end")
	}
	if _, exists := m["start"]; !exists {
		return fmt.Errorf("missing required field: start")
	}
	type Alias Range
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Range(test)
	
	return nil
}
// Client Capabilities for a {@link ReferencesRequest}.
type ReferenceClientCapabilities struct {
	// Whether references supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *ReferenceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ReferenceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ReferenceClientCapabilities(test)
	
	return nil
}
// Value-object that contains additional information when
// requesting references.
type ReferenceContext struct {
	// Include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration"`
}
func (t *ReferenceContext) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["includeDeclaration"]; !exists {
		return fmt.Errorf("missing required field: includeDeclaration")
	}
	type Alias ReferenceContext
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ReferenceContext(test)
	
	return nil
}
// Reference options.
type ReferenceOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *ReferenceOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ReferenceOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ReferenceOptions(test)
	
	return nil
}
// Parameters for a {@link ReferencesRequest}.
type ReferenceParams struct {

	Context ReferenceContext `json:"context"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *ReferenceParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["context"]; !exists {
		return fmt.Errorf("missing required field: context")
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias ReferenceParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ReferenceParams(test)
	
	return nil
}
// Registration options for a {@link ReferencesRequest}.
type ReferenceRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *ReferenceRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias ReferenceRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ReferenceRegistrationOptions(test)
	
	return nil
}
// General parameters to register for a notification or to register a provider.
type Registration struct {
	// The id used to register the request. The id can be used to deregister
	// the request again.
	Id string `json:"id"`
	// The method / capability to register for.
	Method string `json:"method"`
	// Options necessary for the registration.
	RegisterOptions any `json:"registerOptions,omitempty"`
}
func (t *Registration) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["id"]; !exists {
		return fmt.Errorf("missing required field: id")
	}
	if _, exists := m["method"]; !exists {
		return fmt.Errorf("missing required field: method")
	}
	type Alias Registration
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Registration(test)
	
	return nil
}

type RegistrationParams struct {

	Registrations []Registration `json:"registrations"`
}
func (t *RegistrationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["registrations"]; !exists {
		return fmt.Errorf("missing required field: registrations")
	}
	type Alias RegistrationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RegistrationParams(test)
	
	return nil
}
// Client capabilities specific to regular expressions.
// 
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
	// The engine's name.
	Engine RegularExpressionEngineKind `json:"engine"`
	// The engine's version.
	Version string `json:"version,omitempty"`
}
func (t *RegularExpressionsClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["engine"]; !exists {
		return fmt.Errorf("missing required field: engine")
	}
	type Alias RegularExpressionsClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RegularExpressionsClientCapabilities(test)
	
	return nil
}
// A full diagnostic report with a set of related documents.
// 
// @since 3.17.0
type RelatedFullDocumentDiagnosticReport struct {
	// The actual items.
	Items []Diagnostic `json:"items"`
	// A full document diagnostic report.
	Kind string `json:"kind"`
	// Diagnostics of related documents. This information is useful
	// in programming languages where code in a file A can generate
	// diagnostics in a file B which A depends on. An example of
	// such a language is C/C++ where marco definitions in a file
	// a.cpp and result in errors in a header file b.hpp.
	// 
	// @since 3.17.0
	RelatedDocuments map[DocumentUri]Or2[FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport] `json:"relatedDocuments,omitzero"`
	// An optional result id. If provided it will
	// be sent on the next diagnostic request for the
	// same document.
	ResultId string `json:"resultId,omitempty"`
}
func (t *RelatedFullDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias RelatedFullDocumentDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RelatedFullDocumentDiagnosticReport(test)
	
	return nil
}
// An unchanged diagnostic report with a set of related documents.
// 
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReport struct {
	// A document diagnostic report indicating
	// no changes to the last result. A server can
	// only return `unchanged` if result ids are
	// provided.
	Kind string `json:"kind"`
	// Diagnostics of related documents. This information is useful
	// in programming languages where code in a file A can generate
	// diagnostics in a file B which A depends on. An example of
	// such a language is C/C++ where marco definitions in a file
	// a.cpp and result in errors in a header file b.hpp.
	// 
	// @since 3.17.0
	RelatedDocuments map[DocumentUri]Or2[FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport] `json:"relatedDocuments,omitzero"`
	// A result id which will be sent on the next
	// diagnostic request for the same document.
	ResultId string `json:"resultId"`
}
func (t *RelatedUnchangedDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["resultId"]; !exists {
		return fmt.Errorf("missing required field: resultId")
	}
	type Alias RelatedUnchangedDocumentDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RelatedUnchangedDocumentDiagnosticReport(test)
	
	return nil
}
// A relative pattern is a helper to construct glob patterns that are matched
// relatively to a base URI. The common value for a `baseUri` is a workspace
// folder root, but it can be another absolute URI as well.
// 
// @since 3.17.0
type RelativePattern struct {
	// A workspace folder or a base URI to which this pattern will be matched
	// against relatively.
	BaseUri Or2[WorkspaceFolder, URI] `json:"baseUri"`
	// The actual glob pattern;
	Pattern Pattern `json:"pattern"`
}
func (t *RelativePattern) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["baseUri"]; !exists {
		return fmt.Errorf("missing required field: baseUri")
	}
	if _, exists := m["pattern"]; !exists {
		return fmt.Errorf("missing required field: pattern")
	}
	type Alias RelativePattern
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RelativePattern(test)
	
	return nil
}

type RenameClientCapabilities struct {
	// Whether rename supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// rename request's workspace edit by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	// 
	// @since 3.16.0
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
	// Client supports testing for validity of rename operations
	// before execution.
	// 
	// @since 3.12.0
	PrepareSupport bool `json:"prepareSupport,omitempty"`
	// Client supports the default behavior result.
	// 
	// The value indicates the default behavior used by the
	// client.
	// 
	// @since 3.16.0
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitzero"`
}
func (t *RenameClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias RenameClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameClientCapabilities(test)
	
	return nil
}
// Rename file operation
type RenameFile struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitzero"`
	// A rename
	Kind string `json:"kind"`
	// The new location.
	NewUri DocumentUri `json:"newUri"`
	// The old (existing) location.
	OldUri DocumentUri `json:"oldUri"`
	// Rename options.
	Options *RenameFileOptions `json:"options,omitzero"`
}
func (t *RenameFile) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["newUri"]; !exists {
		return fmt.Errorf("missing required field: newUri")
	}
	if _, exists := m["oldUri"]; !exists {
		return fmt.Errorf("missing required field: oldUri")
	}
	type Alias RenameFile
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameFile(test)
	
	return nil
}
// Rename file options
type RenameFileOptions struct {
	// Ignores if target exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	Overwrite bool `json:"overwrite,omitempty"`
}
func (t *RenameFileOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias RenameFileOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameFileOptions(test)
	
	return nil
}
// The parameters sent in notifications/requests for user-initiated renames of
// files.
// 
// @since 3.16.0
type RenameFilesParams struct {
	// An array of all files/folders renamed in this operation. When a folder is renamed, only
	// the folder will be included, and not its children.
	Files []FileRename `json:"files"`
}
func (t *RenameFilesParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["files"]; !exists {
		return fmt.Errorf("missing required field: files")
	}
	type Alias RenameFilesParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameFilesParams(test)
	
	return nil
}
// Provider options for a {@link RenameRequest}.
type RenameOptions struct {
	// Renames should be checked and tested before being executed.
	// 
	// @since version 3.12.0
	PrepareProvider bool `json:"prepareProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *RenameOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias RenameOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameOptions(test)
	
	return nil
}
// The parameters of a {@link RenameRequest}.
type RenameParams struct {
	// The new name of the symbol. If the given name is not valid the
	// request must return a {@link ResponseError} with an
	// appropriate message set.
	NewName string `json:"newName"`
	// The position at which this request was sent.
	Position Position `json:"position"`
	// The document to rename.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *RenameParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["newName"]; !exists {
		return fmt.Errorf("missing required field: newName")
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias RenameParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameParams(test)
	
	return nil
}
// Registration options for a {@link RenameRequest}.
type RenameRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Renames should be checked and tested before being executed.
	// 
	// @since version 3.12.0
	PrepareProvider bool `json:"prepareProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *RenameRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias RenameRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = RenameRegistrationOptions(test)
	
	return nil
}
// A generic resource operation.
type ResourceOperation struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitzero"`
	// The resource operation kind.
	Kind string `json:"kind"`
}
func (t *ResourceOperation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias ResourceOperation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ResourceOperation(test)
	
	return nil
}
// Save options.
type SaveOptions struct {
	// The client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}
func (t *SaveOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SaveOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SaveOptions(test)
	
	return nil
}
// Describes the currently selected completion item.
// 
// @since 3.18.0
// @proposed
type SelectedCompletionInfo struct {
	// The range that will be replaced if this completion item is accepted.
	Range Range `json:"range"`
	// The text the range will be replaced with if this completion is accepted.
	Text string `json:"text"`
}
func (t *SelectedCompletionInfo) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["text"]; !exists {
		return fmt.Errorf("missing required field: text")
	}
	type Alias SelectedCompletionInfo
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SelectedCompletionInfo(test)
	
	return nil
}
// A selection range represents a part of a selection hierarchy. A selection range
// may have a parent selection range that contains it.
type SelectionRange struct {
	// The parent selection range containing this range. Therefore `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent,omitzero"`
	// The {@link Range range} of this selection range.
	Range Range `json:"range"`
}
func (t *SelectionRange) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias SelectionRange
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SelectionRange(test)
	
	return nil
}

type SelectionRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for selection range providers. If this is set to `true`
	// the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *SelectionRangeClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SelectionRangeClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SelectionRangeClientCapabilities(test)
	
	return nil
}

type SelectionRangeOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *SelectionRangeOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SelectionRangeOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SelectionRangeOptions(test)
	
	return nil
}
// A parameter literal used in selection range requests.
type SelectionRangeParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The positions inside the text document.
	Positions []Position `json:"positions"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *SelectionRangeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["positions"]; !exists {
		return fmt.Errorf("missing required field: positions")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias SelectionRangeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SelectionRangeParams(test)
	
	return nil
}

type SelectionRangeRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *SelectionRangeRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias SelectionRangeRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SelectionRangeRegistrationOptions(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokens struct {
	// The actual tokens.
	Data []uint32 `json:"data"`
	// An optional result id. If provided and clients support delta updating
	// the client will include the result id in the next semantic token request.
	// A server can then instead of computing all semantic tokens again simply
	// send a delta.
	ResultId string `json:"resultId,omitempty"`
}
func (t *SemanticTokens) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["data"]; !exists {
		return fmt.Errorf("missing required field: data")
	}
	type Alias SemanticTokens
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokens(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensClientCapabilities struct {
	// Whether the client uses semantic tokens to augment existing
	// syntax tokens. If set to `true` client side created syntax
	// tokens and semantic tokens are both used for colorization. If
	// set to `false` the client only uses the returned semantic tokens
	// for colorization.
	// 
	// If the value is `undefined` then the client behavior is not
	// specified.
	// 
	// @since 3.17.0
	AugmentsSyntaxTokens bool `json:"augmentsSyntaxTokens,omitempty"`
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The token formats the clients supports.
	Formats []TokenFormat `json:"formats"`
	// Whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport bool `json:"multilineTokenSupport,omitempty"`
	// Whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport bool `json:"overlappingTokenSupport,omitempty"`
	// Which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might not
	// show semantic tokens or degrade some of the user experience if a range
	// or full request is advertised by the client but not provided by the
	// server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or might
	// even decide to not show any semantic tokens at all.
	Requests ClientSemanticTokensRequestOptions `json:"requests"`
	// Whether the client allows the server to actively cancel a
	// semantic token request, e.g. supports returning
	// LSPErrorCodes.ServerCancelled. If a server does the client
	// needs to retrigger the request.
	// 
	// @since 3.17.0
	ServerCancelSupport bool `json:"serverCancelSupport,omitempty"`
	// The token modifiers that the client supports.
	TokenModifiers []string `json:"tokenModifiers"`
	// The token types that the client supports.
	TokenTypes []string `json:"tokenTypes"`
}
func (t *SemanticTokensClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["formats"]; !exists {
		return fmt.Errorf("missing required field: formats")
	}
	if _, exists := m["requests"]; !exists {
		return fmt.Errorf("missing required field: requests")
	}
	if _, exists := m["tokenModifiers"]; !exists {
		return fmt.Errorf("missing required field: tokenModifiers")
	}
	if _, exists := m["tokenTypes"]; !exists {
		return fmt.Errorf("missing required field: tokenTypes")
	}
	type Alias SemanticTokensClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensClientCapabilities(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensDelta struct {
	// The semantic token edits to transform a previous result into a new result.
	Edits []SemanticTokensEdit `json:"edits"`

	ResultId string `json:"resultId,omitempty"`
}
func (t *SemanticTokensDelta) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["edits"]; !exists {
		return fmt.Errorf("missing required field: edits")
	}
	type Alias SemanticTokensDelta
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensDelta(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensDeltaParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The result id of a previous response. The result Id can either point to a full response
	// or a delta response depending on what was received last.
	PreviousResultId string `json:"previousResultId"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *SemanticTokensDeltaParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["previousResultId"]; !exists {
		return fmt.Errorf("missing required field: previousResultId")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias SemanticTokensDeltaParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensDeltaParams(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {

	Edits []SemanticTokensEdit `json:"edits"`
}
func (t *SemanticTokensDeltaPartialResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["edits"]; !exists {
		return fmt.Errorf("missing required field: edits")
	}
	type Alias SemanticTokensDeltaPartialResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensDeltaPartialResult(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensEdit struct {
	// The elements to insert.
	Data []uint32 `json:"data,omitzero"`
	// The count of elements to remove.
	DeleteCount uint32 `json:"deleteCount"`
	// The start offset of the edit.
	Start uint32 `json:"start"`
}
func (t *SemanticTokensEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["deleteCount"]; !exists {
		return fmt.Errorf("missing required field: deleteCount")
	}
	if _, exists := m["start"]; !exists {
		return fmt.Errorf("missing required field: start")
	}
	type Alias SemanticTokensEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensEdit(test)
	
	return nil
}
// Semantic tokens options to support deltas for full documents
// 
// @since 3.18.0
type SemanticTokensFullDelta struct {
	// The server supports deltas for full documents.
	Delta bool `json:"delta,omitempty"`
}
func (t *SemanticTokensFullDelta) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SemanticTokensFullDelta
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensFullDelta(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensLegend struct {
	// The token modifiers a server uses.
	TokenModifiers []string `json:"tokenModifiers"`
	// The token types a server uses.
	TokenTypes []string `json:"tokenTypes"`
}
func (t *SemanticTokensLegend) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["tokenModifiers"]; !exists {
		return fmt.Errorf("missing required field: tokenModifiers")
	}
	if _, exists := m["tokenTypes"]; !exists {
		return fmt.Errorf("missing required field: tokenTypes")
	}
	type Alias SemanticTokensLegend
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensLegend(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensOptions struct {
	// Server supports providing semantic tokens for a full document.
	Full *Or2[bool, SemanticTokensFullDelta] `json:"full,omitzero"`
	// The legend used by the server
	Legend SemanticTokensLegend `json:"legend"`
	// Server supports providing semantic tokens for a specific range
	// of a document.
	Range *Or2[bool, LSPObject] `json:"range,omitzero"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *SemanticTokensOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["legend"]; !exists {
		return fmt.Errorf("missing required field: legend")
	}
	type Alias SemanticTokensOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensOptions(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *SemanticTokensParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias SemanticTokensParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensParams(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensPartialResult struct {

	Data []uint32 `json:"data"`
}
func (t *SemanticTokensPartialResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["data"]; !exists {
		return fmt.Errorf("missing required field: data")
	}
	type Alias SemanticTokensPartialResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensPartialResult(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensRangeParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The range the semantic tokens are requested for.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *SemanticTokensRangeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias SemanticTokensRangeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensRangeParams(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Server supports providing semantic tokens for a full document.
	Full *Or2[bool, SemanticTokensFullDelta] `json:"full,omitzero"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// The legend used by the server
	Legend SemanticTokensLegend `json:"legend"`
	// Server supports providing semantic tokens for a specific range
	// of a document.
	Range *Or2[bool, LSPObject] `json:"range,omitzero"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *SemanticTokensRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	if _, exists := m["legend"]; !exists {
		return fmt.Errorf("missing required field: legend")
	}
	type Alias SemanticTokensRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensRegistrationOptions(test)
	
	return nil
}
// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	// 
	// Note that this event is global and will force the client to refresh all
	// semantic tokens currently shown. It should be used with absolute care
	// and is useful for situation where a server for example detects a project
	// wide change that requires such a calculation.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}
func (t *SemanticTokensWorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SemanticTokensWorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SemanticTokensWorkspaceClientCapabilities(test)
	
	return nil
}
// Defines the capabilities provided by a language
// server.
type ServerCapabilities struct {
	// The server provides call hierarchy support.
	// 
	// @since 3.16.0
	CallHierarchyProvider *Or3[bool, CallHierarchyOptions, CallHierarchyRegistrationOptions] `json:"callHierarchyProvider,omitzero"`
	// The server provides code actions. CodeActionOptions may only be
	// specified if the client states that it supports
	// `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider *Or2[bool, CodeActionOptions] `json:"codeActionProvider,omitzero"`
	// The server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitzero"`
	// The server provides color provider support.
	ColorProvider *Or3[bool, DocumentColorOptions, DocumentColorRegistrationOptions] `json:"colorProvider,omitzero"`
	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitzero"`
	// The server provides Goto Declaration support.
	DeclarationProvider *Or3[bool, DeclarationOptions, DeclarationRegistrationOptions] `json:"declarationProvider,omitzero"`
	// The server provides goto definition support.
	DefinitionProvider *Or2[bool, DefinitionOptions] `json:"definitionProvider,omitzero"`
	// The server has support for pull model diagnostics.
	// 
	// @since 3.17.0
	DiagnosticProvider *Or2[DiagnosticOptions, DiagnosticRegistrationOptions] `json:"diagnosticProvider,omitzero"`
	// The server provides document formatting.
	DocumentFormattingProvider *Or2[bool, DocumentFormattingOptions] `json:"documentFormattingProvider,omitzero"`
	// The server provides document highlight support.
	DocumentHighlightProvider *Or2[bool, DocumentHighlightOptions] `json:"documentHighlightProvider,omitzero"`
	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitzero"`
	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitzero"`
	// The server provides document range formatting.
	DocumentRangeFormattingProvider *Or2[bool, DocumentRangeFormattingOptions] `json:"documentRangeFormattingProvider,omitzero"`
	// The server provides document symbol support.
	DocumentSymbolProvider *Or2[bool, DocumentSymbolOptions] `json:"documentSymbolProvider,omitzero"`
	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitzero"`
	// Experimental server capabilities.
	Experimental any `json:"experimental,omitempty"`
	// The server provides folding provider support.
	FoldingRangeProvider *Or3[bool, FoldingRangeOptions, FoldingRangeRegistrationOptions] `json:"foldingRangeProvider,omitzero"`
	// The server provides hover support.
	HoverProvider *Or2[bool, HoverOptions] `json:"hoverProvider,omitzero"`
	// The server provides Goto Implementation support.
	ImplementationProvider *Or3[bool, ImplementationOptions, ImplementationRegistrationOptions] `json:"implementationProvider,omitzero"`
	// The server provides inlay hints.
	// 
	// @since 3.17.0
	InlayHintProvider *Or3[bool, InlayHintOptions, InlayHintRegistrationOptions] `json:"inlayHintProvider,omitzero"`
	// Inline completion options used during static registration.
	// 
	// @since 3.18.0
	// @proposed
	InlineCompletionProvider *Or2[bool, InlineCompletionOptions] `json:"inlineCompletionProvider,omitzero"`
	// The server provides inline values.
	// 
	// @since 3.17.0
	InlineValueProvider *Or3[bool, InlineValueOptions, InlineValueRegistrationOptions] `json:"inlineValueProvider,omitzero"`
	// The server provides linked editing range support.
	// 
	// @since 3.16.0
	LinkedEditingRangeProvider *Or3[bool, LinkedEditingRangeOptions, LinkedEditingRangeRegistrationOptions] `json:"linkedEditingRangeProvider,omitzero"`
	// The server provides moniker support.
	// 
	// @since 3.16.0
	MonikerProvider *Or3[bool, MonikerOptions, MonikerRegistrationOptions] `json:"monikerProvider,omitzero"`
	// Defines how notebook documents are synced.
	// 
	// @since 3.17.0
	NotebookDocumentSync *Or2[NotebookDocumentSyncOptions, NotebookDocumentSyncRegistrationOptions] `json:"notebookDocumentSync,omitzero"`
	// The position encoding the server picked from the encodings offered
	// by the client via the client capability `general.positionEncodings`.
	// 
	// If the client didn't provide any position encodings the only valid
	// value that a server can return is 'utf-16'.
	// 
	// If omitted it defaults to 'utf-16'.
	// 
	// @since 3.17.0
	PositionEncoding *PositionEncodingKind `json:"positionEncoding,omitzero"`
	// The server provides find references support.
	ReferencesProvider *Or2[bool, ReferenceOptions] `json:"referencesProvider,omitzero"`
	// The server provides rename support. RenameOptions may only be
	// specified if the client states that it supports
	// `prepareSupport` in its initial `initialize` request.
	RenameProvider *Or2[bool, RenameOptions] `json:"renameProvider,omitzero"`
	// The server provides selection range support.
	SelectionRangeProvider *Or3[bool, SelectionRangeOptions, SelectionRangeRegistrationOptions] `json:"selectionRangeProvider,omitzero"`
	// The server provides semantic tokens support.
	// 
	// @since 3.16.0
	SemanticTokensProvider *Or2[SemanticTokensOptions, SemanticTokensRegistrationOptions] `json:"semanticTokensProvider,omitzero"`
	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitzero"`
	// Defines how text documents are synced. Is either a detailed structure
	// defining each notification or for backwards compatibility the
	// TextDocumentSyncKind number.
	TextDocumentSync *Or2[TextDocumentSyncOptions, TextDocumentSyncKind] `json:"textDocumentSync,omitzero"`
	// The server provides Goto Type Definition support.
	TypeDefinitionProvider *Or3[bool, TypeDefinitionOptions, TypeDefinitionRegistrationOptions] `json:"typeDefinitionProvider,omitzero"`
	// The server provides type hierarchy support.
	// 
	// @since 3.17.0
	TypeHierarchyProvider *Or3[bool, TypeHierarchyOptions, TypeHierarchyRegistrationOptions] `json:"typeHierarchyProvider,omitzero"`
	// Workspace specific server capabilities.
	Workspace *WorkspaceOptions `json:"workspace,omitzero"`
	// The server provides workspace symbol support.
	WorkspaceSymbolProvider *Or2[bool, WorkspaceSymbolOptions] `json:"workspaceSymbolProvider,omitzero"`
}
func (t *ServerCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ServerCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ServerCapabilities(test)
	
	return nil
}
// @since 3.18.0
type ServerCompletionItemOptions struct {
	// The server has support for completion item label
	// details (see also `CompletionItemLabelDetails`) when
	// receiving a completion item in a resolve call.
	// 
	// @since 3.17.0
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}
func (t *ServerCompletionItemOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ServerCompletionItemOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ServerCompletionItemOptions(test)
	
	return nil
}
// Information about the server
// 
// @since 3.15.0
// @since 3.18.0 ServerInfo type name added.
type ServerInfo struct {
	// The name of the server as defined by the server.
	Name string `json:"name"`
	// The server's version as defined by the server.
	Version string `json:"version,omitempty"`
}
func (t *ServerInfo) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	type Alias ServerInfo
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ServerInfo(test)
	
	return nil
}

type SetTraceParams struct {

	Value TraceValue `json:"value"`
}
func (t *SetTraceParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias SetTraceParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SetTraceParams(test)
	
	return nil
}
// Client capabilities for the showDocument request.
// 
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// The client has support for the showDocument
	// request.
	Support bool `json:"support"`
}
func (t *ShowDocumentClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["support"]; !exists {
		return fmt.Errorf("missing required field: support")
	}
	type Alias ShowDocumentClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ShowDocumentClientCapabilities(test)
	
	return nil
}
// Params to show a resource in the UI.
// 
// @since 3.16.0
type ShowDocumentParams struct {
	// Indicates to show the resource in an external program.
	// To show, for example, `https://code.visualstudio.com/`
	// in the default WEB browser set `external` to `true`.
	External bool `json:"external,omitempty"`
	// An optional selection range if the document is a text
	// document. Clients might ignore the property if an
	// external program is started or the file is not a text
	// file.
	Selection *Range `json:"selection,omitzero"`
	// An optional property to indicate whether the editor
	// showing the document should take focus or not.
	// Clients might ignore this property if an external
	// program is started.
	TakeFocus bool `json:"takeFocus,omitempty"`
	// The uri to show.
	Uri URI `json:"uri"`
}
func (t *ShowDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias ShowDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ShowDocumentParams(test)
	
	return nil
}
// The result of a showDocument request.
// 
// @since 3.16.0
type ShowDocumentResult struct {
	// A boolean indicating if the show was successful.
	Success bool `json:"success"`
}
func (t *ShowDocumentResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["success"]; !exists {
		return fmt.Errorf("missing required field: success")
	}
	type Alias ShowDocumentResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ShowDocumentResult(test)
	
	return nil
}
// The parameters of a notification message.
type ShowMessageParams struct {
	// The actual message.
	Message string `json:"message"`
	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`
}
func (t *ShowMessageParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["message"]; !exists {
		return fmt.Errorf("missing required field: message")
	}
	if _, exists := m["type"]; !exists {
		return fmt.Errorf("missing required field: type")
	}
	type Alias ShowMessageParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ShowMessageParams(test)
	
	return nil
}
// Show message request client capabilities
type ShowMessageRequestClientCapabilities struct {
	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *ClientShowMessageActionItemOptions `json:"messageActionItem,omitzero"`
}
func (t *ShowMessageRequestClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias ShowMessageRequestClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ShowMessageRequestClientCapabilities(test)
	
	return nil
}

type ShowMessageRequestParams struct {
	// The message action items to present.
	Actions []MessageActionItem `json:"actions,omitzero"`
	// The actual message.
	Message string `json:"message"`
	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`
}
func (t *ShowMessageRequestParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["message"]; !exists {
		return fmt.Errorf("missing required field: message")
	}
	if _, exists := m["type"]; !exists {
		return fmt.Errorf("missing required field: type")
	}
	type Alias ShowMessageRequestParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = ShowMessageRequestParams(test)
	
	return nil
}
// Signature help represents the signature of something
// callable. There can be multiple signature but only one
// active and only one active parameter.
type SignatureHelp struct {
	// The active parameter of the active signature.
	// 
	// If `null`, no parameter of the signature is active (for example a named
	// argument that does not match any declared parameters). This is only valid
	// if the client specifies the client capability
	// `textDocument.signatureHelp.noActiveParameterSupport === true`
	// 
	// If omitted or the value lies outside the range of
	// `signatures[activeSignature].parameters` defaults to 0 if the active
	// signature has parameters.
	// 
	// If the active signature has no parameters it is ignored.
	// 
	// In future version of the protocol this property might become
	// mandatory (but still nullable) to better express the active parameter if
	// the active signature does have any.
	ActiveParameter **uint32 `json:"activeParameter,omitzero"`
	// The active signature. If omitted or the value lies outside the
	// range of `signatures` the value defaults to zero or is ignored if
	// the `SignatureHelp` has no signatures.
	// 
	// Whenever possible implementors should make an active decision about
	// the active signature and shouldn't rely on a default value.
	// 
	// In future version of the protocol this property might become
	// mandatory to better express this.
	ActiveSignature uint32 `json:"activeSignature,omitempty"`
	// One or more signatures.
	Signatures []SignatureInformation `json:"signatures"`
}
func (t *SignatureHelp) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["signatures"]; !exists {
		return fmt.Errorf("missing required field: signatures")
	}
	type Alias SignatureHelp
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureHelp(test)
	
	return nil
}
// Client Capabilities for a {@link SignatureHelpRequest}.
type SignatureHelpClientCapabilities struct {
	// The client supports to send additional context information for a
	// `textDocument/signatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.
	// 
	// @since 3.15.0
	ContextSupport bool `json:"contextSupport,omitempty"`
	// Whether signature help supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports the following `SignatureInformation`
	// specific properties.
	SignatureInformation *ClientSignatureInformationOptions `json:"signatureInformation,omitzero"`
}
func (t *SignatureHelpClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SignatureHelpClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureHelpClientCapabilities(test)
	
	return nil
}
// Additional information about the context in which a signature help request was triggered.
// 
// @since 3.15.0
type SignatureHelpContext struct {
	// The currently active `SignatureHelp`.
	// 
	// The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on
	// the user navigating through available signatures.
	ActiveSignatureHelp *SignatureHelp `json:"activeSignatureHelp,omitzero"`
	// `true` if signature help was already showing when it was triggered.
	// 
	// Retriggers occurs when the signature help is already active and can be caused by actions such as
	// typing a trigger character, a cursor move, or document content changes.
	IsRetrigger bool `json:"isRetrigger"`
	// Character that caused signature help to be triggered.
	// 
	// This is undefined when `triggerKind !== SignatureHelpTriggerKind.TriggerCharacter`
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
	// Action that caused signature help to be triggered.
	TriggerKind SignatureHelpTriggerKind `json:"triggerKind"`
}
func (t *SignatureHelpContext) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["isRetrigger"]; !exists {
		return fmt.Errorf("missing required field: isRetrigger")
	}
	if _, exists := m["triggerKind"]; !exists {
		return fmt.Errorf("missing required field: triggerKind")
	}
	type Alias SignatureHelpContext
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureHelpContext(test)
	
	return nil
}
// Server Capabilities for a {@link SignatureHelpRequest}.
type SignatureHelpOptions struct {
	// List of characters that re-trigger signature help.
	// 
	// These trigger characters are only active when signature help is already showing. All trigger characters
	// are also counted as re-trigger characters.
	// 
	// @since 3.15.0
	RetriggerCharacters []string `json:"retriggerCharacters,omitzero"`
	// List of characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitzero"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *SignatureHelpOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias SignatureHelpOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureHelpOptions(test)
	
	return nil
}
// Parameters for a {@link SignatureHelpRequest}.
type SignatureHelpParams struct {
	// The signature help context. This is only available if the client specifies
	// to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
	// 
	// @since 3.15.0
	Context *SignatureHelpContext `json:"context,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *SignatureHelpParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias SignatureHelpParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureHelpParams(test)
	
	return nil
}
// Registration options for a {@link SignatureHelpRequest}.
type SignatureHelpRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// List of characters that re-trigger signature help.
	// 
	// These trigger characters are only active when signature help is already showing. All trigger characters
	// are also counted as re-trigger characters.
	// 
	// @since 3.15.0
	RetriggerCharacters []string `json:"retriggerCharacters,omitzero"`
	// List of characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitzero"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *SignatureHelpRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias SignatureHelpRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureHelpRegistrationOptions(test)
	
	return nil
}
// Represents the signature of something callable. A signature
// can have a label, like a function-name, a doc-comment, and
// a set of parameters.
type SignatureInformation struct {
	// The index of the active parameter.
	// 
	// If `null`, no parameter of the signature is active (for example a named
	// argument that does not match any declared parameters). This is only valid
	// if the client specifies the client capability
	// `textDocument.signatureHelp.noActiveParameterSupport === true`
	// 
	// If provided (or `null`), this is used in place of
	// `SignatureHelp.activeParameter`.
	// 
	// @since 3.16.0
	ActiveParameter **uint32 `json:"activeParameter,omitzero"`
	// The human-readable doc-comment of this signature. Will be shown
	// in the UI but can be omitted.
	Documentation *Or2[string, MarkupContent] `json:"documentation,omitzero"`
	// The label of this signature. Will be shown in
	// the UI.
	Label string `json:"label"`
	// The parameters of this signature.
	Parameters []ParameterInformation `json:"parameters,omitzero"`
}
func (t *SignatureInformation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["label"]; !exists {
		return fmt.Errorf("missing required field: label")
	}
	type Alias SignatureInformation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SignatureInformation(test)
	
	return nil
}
// An interactive text edit.
// 
// @since 3.18.0
// @proposed
type SnippetTextEdit struct {
	// The actual identifier of the snippet edit.
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitzero"`
	// The range of the text document to be manipulated.
	Range Range `json:"range"`
	// The snippet to be inserted.
	Snippet StringValue `json:"snippet"`
}
func (t *SnippetTextEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["snippet"]; !exists {
		return fmt.Errorf("missing required field: snippet")
	}
	type Alias SnippetTextEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SnippetTextEdit(test)
	
	return nil
}
// @since 3.18.0
type StaleRequestSupportOptions struct {
	// The client will actively cancel the request.
	Cancel bool `json:"cancel"`
	// The list of requests for which the client
	// will retry the request if it receives a
	// response with error code `ContentModified`
	RetryOnContentModified []string `json:"retryOnContentModified"`
}
func (t *StaleRequestSupportOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["cancel"]; !exists {
		return fmt.Errorf("missing required field: cancel")
	}
	if _, exists := m["retryOnContentModified"]; !exists {
		return fmt.Errorf("missing required field: retryOnContentModified")
	}
	type Alias StaleRequestSupportOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = StaleRequestSupportOptions(test)
	
	return nil
}
// Static registration options to be returned in the initialize
// request.
type StaticRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
}
func (t *StaticRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias StaticRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = StaticRegistrationOptions(test)
	
	return nil
}
// A string value used as a snippet is a template which allows to insert text
// and to control the editor cursor when insertion happens.
// 
// A snippet can define tab stops and placeholders with `$1`, `$2`
// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
// the end of the snippet. Variables are defined with `$name` and
// `${name:default value}`.
// 
// @since 3.18.0
// @proposed
type StringValue struct {
	// The kind of string value.
	Kind string `json:"kind"`
	// The snippet string.
	Value string `json:"value"`
}
func (t *StringValue) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["value"]; !exists {
		return fmt.Errorf("missing required field: value")
	}
	type Alias StringValue
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = StringValue(test)
	
	return nil
}
// Represents information about programming constructs like variables, classes,
// interfaces etc.
type SymbolInformation struct {
	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName string `json:"containerName,omitempty"`
	// Indicates if this symbol is deprecated.
	// 
	// @deprecated Use tags instead
	Deprecated bool `json:"deprecated,omitempty"`
	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`
	// The location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in the
	// tool the range's start information is used to position the cursor. So
	// the range usually spans more than the actual symbol's name and does
	// normally include things like visibility modifiers.
	// 
	// The range doesn't have to denote a node range in the sense of an abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy of
	// the symbols.
	Location Location `json:"location"`
	// The name of this symbol.
	Name string `json:"name"`
	// Tags for this symbol.
	// 
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitzero"`
}
func (t *SymbolInformation) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["location"]; !exists {
		return fmt.Errorf("missing required field: location")
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	type Alias SymbolInformation
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = SymbolInformation(test)
	
	return nil
}
// Describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// How documents are synced to the server.
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}
func (t *TextDocumentChangeRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	if _, exists := m["syncKind"]; !exists {
		return fmt.Errorf("missing required field: syncKind")
	}
	type Alias TextDocumentChangeRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentChangeRegistrationOptions(test)
	
	return nil
}
// Text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	// Capabilities specific to the various call hierarchy requests.
	// 
	// @since 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitzero"`
	// Capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitzero"`
	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitzero"`
	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	// 
	// @since 3.6.0
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitzero"`
	// Capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitzero"`
	// Capabilities specific to the `textDocument/declaration` request.
	// 
	// @since 3.14.0
	Declaration *DeclarationClientCapabilities `json:"declaration,omitzero"`
	// Capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitzero"`
	// Capabilities specific to the diagnostic pull model.
	// 
	// @since 3.17.0
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitzero"`
	// Capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitzero"`
	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitzero"`
	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitzero"`
	// Defines which filters the client supports.
	// 
	// @since 3.18.0
	Filters *TextDocumentFilterClientCapabilities `json:"filters,omitzero"`
	// Capabilities specific to the `textDocument/foldingRange` request.
	// 
	// @since 3.10.0
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitzero"`
	// Capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitzero"`
	// Capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitzero"`
	// Capabilities specific to the `textDocument/implementation` request.
	// 
	// @since 3.6.0
	Implementation *ImplementationClientCapabilities `json:"implementation,omitzero"`
	// Capabilities specific to the `textDocument/inlayHint` request.
	// 
	// @since 3.17.0
	InlayHint *InlayHintClientCapabilities `json:"inlayHint,omitzero"`
	// Client capabilities specific to inline completions.
	// 
	// @since 3.18.0
	// @proposed
	InlineCompletion *InlineCompletionClientCapabilities `json:"inlineCompletion,omitzero"`
	// Capabilities specific to the `textDocument/inlineValue` request.
	// 
	// @since 3.17.0
	InlineValue *InlineValueClientCapabilities `json:"inlineValue,omitzero"`
	// Capabilities specific to the `textDocument/linkedEditingRange` request.
	// 
	// @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitzero"`
	// Client capabilities specific to the `textDocument/moniker` request.
	// 
	// @since 3.16.0
	Moniker *MonikerClientCapabilities `json:"moniker,omitzero"`
	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitzero"`
	// Capabilities specific to the `textDocument/publishDiagnostics` notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitzero"`
	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitzero"`
	// Capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references,omitzero"`
	// Capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitzero"`
	// Capabilities specific to the `textDocument/selectionRange` request.
	// 
	// @since 3.15.0
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitzero"`
	// Capabilities specific to the various semantic token request.
	// 
	// @since 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitzero"`
	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitzero"`
	// Defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitzero"`
	// Capabilities specific to the `textDocument/typeDefinition` request.
	// 
	// @since 3.6.0
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitzero"`
	// Capabilities specific to the various type hierarchy requests.
	// 
	// @since 3.17.0
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"typeHierarchy,omitzero"`
}
func (t *TextDocumentClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TextDocumentClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentClientCapabilities(test)
	
	return nil
}
// @since 3.18.0
type TextDocumentContentChangePartial struct {
	// The range of the document that changed.
	Range Range `json:"range"`
	// The optional length of the range that got replaced.
	// 
	// @deprecated use range instead.
	RangeLength uint32 `json:"rangeLength,omitempty"`
	// The new text for the provided range.
	Text string `json:"text"`
}
func (t *TextDocumentContentChangePartial) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["text"]; !exists {
		return fmt.Errorf("missing required field: text")
	}
	type Alias TextDocumentContentChangePartial
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentChangePartial(test)
	
	return nil
}
// @since 3.18.0
type TextDocumentContentChangeWholeDocument struct {
	// The new text of the whole document.
	Text string `json:"text"`
}
func (t *TextDocumentContentChangeWholeDocument) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["text"]; !exists {
		return fmt.Errorf("missing required field: text")
	}
	type Alias TextDocumentContentChangeWholeDocument
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentChangeWholeDocument(test)
	
	return nil
}
// Client capabilities for a text document content provider.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentClientCapabilities struct {
	// Text document content provider supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *TextDocumentContentClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TextDocumentContentClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentClientCapabilities(test)
	
	return nil
}
// Text document content provider options.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentOptions struct {
	// The schemes for which the server provides content.
	Schemes []string `json:"schemes"`
}
func (t *TextDocumentContentOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["schemes"]; !exists {
		return fmt.Errorf("missing required field: schemes")
	}
	type Alias TextDocumentContentOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentOptions(test)
	
	return nil
}
// Parameters for the `workspace/textDocumentContent` request.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentParams struct {
	// The uri of the text document.
	Uri DocumentUri `json:"uri"`
}
func (t *TextDocumentContentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias TextDocumentContentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentParams(test)
	
	return nil
}
// Parameters for the `workspace/textDocumentContent/refresh` request.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentRefreshParams struct {
	// The uri of the text document to refresh.
	Uri DocumentUri `json:"uri"`
}
func (t *TextDocumentContentRefreshParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias TextDocumentContentRefreshParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentRefreshParams(test)
	
	return nil
}
// Text document content provider registration options.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// The schemes for which the server provides content.
	Schemes []string `json:"schemes"`
}
func (t *TextDocumentContentRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["schemes"]; !exists {
		return fmt.Errorf("missing required field: schemes")
	}
	type Alias TextDocumentContentRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentRegistrationOptions(test)
	
	return nil
}
// Result of the `workspace/textDocumentContent` request.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentResult struct {
	// The text content of the text document. Please note, that the content of
	// any subsequent open notifications for the text document might differ
	// from the returned content due to whitespace and line ending
	// normalizations done on the client
	Text string `json:"text"`
}
func (t *TextDocumentContentResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["text"]; !exists {
		return fmt.Errorf("missing required field: text")
	}
	type Alias TextDocumentContentResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentContentResult(test)
	
	return nil
}
// Describes textual changes on a text document. A TextDocumentEdit describes all changes
// on a document version Si and after they are applied move the document to version Si+1.
// So the creator of a TextDocumentEdit doesn't need to sort the array of edits or do any
// kind of ordering. However the edits must be non overlapping.
type TextDocumentEdit struct {
	// The edits to be applied.
	// 
	// @since 3.16.0 - support for AnnotatedTextEdit. This is guarded using a
	// client capability.
	// 
	// @since 3.18.0 - support for SnippetTextEdit. This is guarded using a
	// client capability.
	Edits []Or3[TextEdit, AnnotatedTextEdit, SnippetTextEdit] `json:"edits"`
	// The text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`
}
func (t *TextDocumentEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["edits"]; !exists {
		return fmt.Errorf("missing required field: edits")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias TextDocumentEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentEdit(test)
	
	return nil
}

type TextDocumentFilterClientCapabilities struct {
	// The client supports Relative Patterns.
	// 
	// @since 3.18.0
	RelativePatternSupport bool `json:"relativePatternSupport,omitempty"`
}
func (t *TextDocumentFilterClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TextDocumentFilterClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentFilterClientCapabilities(test)
	
	return nil
}
// A document filter where `language` is required field.
// 
// @since 3.18.0
type TextDocumentFilterLanguage struct {
	// A language id, like `typescript`.
	Language string `json:"language"`
	// A glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	// 
	// @since 3.18.0 - support for relative patterns. Whether clients support
	// relative patterns depends on the client capability
	// `textDocuments.filters.relativePatternSupport`.
	Pattern *GlobPattern `json:"pattern,omitzero"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
}
func (t *TextDocumentFilterLanguage) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["language"]; !exists {
		return fmt.Errorf("missing required field: language")
	}
	type Alias TextDocumentFilterLanguage
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentFilterLanguage(test)
	
	return nil
}
// A document filter where `pattern` is required field.
// 
// @since 3.18.0
type TextDocumentFilterPattern struct {
	// A language id, like `typescript`.
	Language string `json:"language,omitempty"`
	// A glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	// 
	// @since 3.18.0 - support for relative patterns. Whether clients support
	// relative patterns depends on the client capability
	// `textDocuments.filters.relativePatternSupport`.
	Pattern GlobPattern `json:"pattern"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
}
func (t *TextDocumentFilterPattern) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["pattern"]; !exists {
		return fmt.Errorf("missing required field: pattern")
	}
	type Alias TextDocumentFilterPattern
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentFilterPattern(test)
	
	return nil
}
// A document filter where `scheme` is required field.
// 
// @since 3.18.0
type TextDocumentFilterScheme struct {
	// A language id, like `typescript`.
	Language string `json:"language,omitempty"`
	// A glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	// 
	// @since 3.18.0 - support for relative patterns. Whether clients support
	// relative patterns depends on the client capability
	// `textDocuments.filters.relativePatternSupport`.
	Pattern *GlobPattern `json:"pattern,omitzero"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme"`
}
func (t *TextDocumentFilterScheme) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["scheme"]; !exists {
		return fmt.Errorf("missing required field: scheme")
	}
	type Alias TextDocumentFilterScheme
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentFilterScheme(test)
	
	return nil
}
// A literal to identify a text document in the client.
type TextDocumentIdentifier struct {
	// The text document's uri.
	Uri DocumentUri `json:"uri"`
}
func (t *TextDocumentIdentifier) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias TextDocumentIdentifier
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentIdentifier(test)
	
	return nil
}
// An item to transfer a text document from the client to the
// server.
type TextDocumentItem struct {
	// The text document's language identifier.
	LanguageId LanguageKind `json:"languageId"`
	// The content of the opened text document.
	Text string `json:"text"`
	// The text document's uri.
	Uri DocumentUri `json:"uri"`
	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int32 `json:"version"`
}
func (t *TextDocumentItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["languageId"]; !exists {
		return fmt.Errorf("missing required field: languageId")
	}
	if _, exists := m["text"]; !exists {
		return fmt.Errorf("missing required field: text")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias TextDocumentItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentItem(test)
	
	return nil
}
// A parameter literal used in requests to pass a text document and a position inside that
// document.
type TextDocumentPositionParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
func (t *TextDocumentPositionParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias TextDocumentPositionParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentPositionParams(test)
	
	return nil
}
// General text document registration options.
type TextDocumentRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
}
func (t *TextDocumentRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias TextDocumentRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentRegistrationOptions(test)
	
	return nil
}
// Save registration options.
type TextDocumentSaveRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
}
func (t *TextDocumentSaveRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias TextDocumentSaveRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentSaveRegistrationOptions(test)
	
	return nil
}

type TextDocumentSyncClientCapabilities struct {
	// The client supports did save notifications.
	DidSave bool `json:"didSave,omitempty"`
	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports sending will save notifications.
	WillSave bool `json:"willSave,omitempty"`
	// The client supports sending a will save request and
	// waits for a response providing text edits which will
	// be applied to the document before it is saved.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`
}
func (t *TextDocumentSyncClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TextDocumentSyncClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentSyncClientCapabilities(test)
	
	return nil
}

type TextDocumentSyncOptions struct {
	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind `json:"change,omitzero"`
	// Open and close notifications are sent to the server. If omitted open close notification should not
	// be sent.
	OpenClose bool `json:"openClose,omitempty"`
	// If present save notifications are sent to the server. If omitted the notification should not be
	// sent.
	Save *Or2[bool, SaveOptions] `json:"save,omitzero"`
	// If present will save notifications are sent to the server. If omitted the notification should not be
	// sent.
	WillSave bool `json:"willSave,omitempty"`
	// If present will save wait until requests are sent to the server. If omitted the request should not be
	// sent.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`
}
func (t *TextDocumentSyncOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TextDocumentSyncOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextDocumentSyncOptions(test)
	
	return nil
}
// A text edit applicable to a text document.
type TextEdit struct {
	// The string to be inserted. For delete operations use an
	// empty string.
	NewText string `json:"newText"`
	// The range of the text document to be manipulated. To insert
	// text into a document create a range where start === end.
	Range Range `json:"range"`
}
func (t *TextEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["newText"]; !exists {
		return fmt.Errorf("missing required field: newText")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	type Alias TextEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TextEdit(test)
	
	return nil
}
// Since 3.6.0
type TypeDefinitionClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `TypeDefinitionRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports additional metadata in the form of definition links.
	// 
	// Since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}
func (t *TypeDefinitionClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TypeDefinitionClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeDefinitionClientCapabilities(test)
	
	return nil
}

type TypeDefinitionOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *TypeDefinitionOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TypeDefinitionOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeDefinitionOptions(test)
	
	return nil
}

type TypeDefinitionParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *TypeDefinitionParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias TypeDefinitionParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeDefinitionParams(test)
	
	return nil
}

type TypeDefinitionRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *TypeDefinitionRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias TypeDefinitionRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeDefinitionRegistrationOptions(test)
	
	return nil
}
// @since 3.17.0
type TypeHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
func (t *TypeHierarchyClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TypeHierarchyClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchyClientCapabilities(test)
	
	return nil
}
// @since 3.17.0
type TypeHierarchyItem struct {
	// A data entry field that is preserved between a type hierarchy prepare and
	// supertypes or subtypes requests. It could also be used to identify the
	// type hierarchy in the server, helping improve the performance on
	// resolving supertypes and subtypes.
	Data any `json:"data,omitempty"`
	// More detail for this item, e.g. the signature of a function.
	Detail string `json:"detail,omitempty"`
	// The kind of this item.
	Kind SymbolKind `json:"kind"`
	// The name of this item.
	Name string `json:"name"`
	// The range enclosing this symbol not including leading/trailing whitespace
	// but everything else, e.g. comments and code.
	Range Range `json:"range"`
	// The range that should be selected and revealed when this symbol is being
	// picked, e.g. the name of a function. Must be contained by the
	// {@link TypeHierarchyItem.range `range`}.
	SelectionRange Range `json:"selectionRange"`
	// Tags for this item.
	Tags []SymbolTag `json:"tags,omitzero"`
	// The resource identifier of this item.
	Uri DocumentUri `json:"uri"`
}
func (t *TypeHierarchyItem) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	if _, exists := m["range"]; !exists {
		return fmt.Errorf("missing required field: range")
	}
	if _, exists := m["selectionRange"]; !exists {
		return fmt.Errorf("missing required field: selectionRange")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias TypeHierarchyItem
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchyItem(test)
	
	return nil
}
// Type hierarchy options used during static registration.
// 
// @since 3.17.0
type TypeHierarchyOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *TypeHierarchyOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias TypeHierarchyOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchyOptions(test)
	
	return nil
}
// The parameter of a `textDocument/prepareTypeHierarchy` request.
// 
// @since 3.17.0
type TypeHierarchyPrepareParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *TypeHierarchyPrepareParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["position"]; !exists {
		return fmt.Errorf("missing required field: position")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias TypeHierarchyPrepareParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchyPrepareParams(test)
	
	return nil
}
// Type hierarchy options used during static or dynamic registration.
// 
// @since 3.17.0
type TypeHierarchyRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *TypeHierarchyRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["documentSelector"]; !exists {
		return fmt.Errorf("missing required field: documentSelector")
	}
	type Alias TypeHierarchyRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchyRegistrationOptions(test)
	
	return nil
}
// The parameter of a `typeHierarchy/subtypes` request.
// 
// @since 3.17.0
type TypeHierarchySubtypesParams struct {

	Item TypeHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *TypeHierarchySubtypesParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["item"]; !exists {
		return fmt.Errorf("missing required field: item")
	}
	type Alias TypeHierarchySubtypesParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchySubtypesParams(test)
	
	return nil
}
// The parameter of a `typeHierarchy/supertypes` request.
// 
// @since 3.17.0
type TypeHierarchySupertypesParams struct {

	Item TypeHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *TypeHierarchySupertypesParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["item"]; !exists {
		return fmt.Errorf("missing required field: item")
	}
	type Alias TypeHierarchySupertypesParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = TypeHierarchySupertypesParams(test)
	
	return nil
}
// A diagnostic report indicating that the last returned
// report is still accurate.
// 
// @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {
	// A document diagnostic report indicating
	// no changes to the last result. A server can
	// only return `unchanged` if result ids are
	// provided.
	Kind string `json:"kind"`
	// A result id which will be sent on the next
	// diagnostic request for the same document.
	ResultId string `json:"resultId"`
}
func (t *UnchangedDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["resultId"]; !exists {
		return fmt.Errorf("missing required field: resultId")
	}
	type Alias UnchangedDocumentDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = UnchangedDocumentDiagnosticReport(test)
	
	return nil
}
// General parameters to unregister a request or notification.
type Unregistration struct {
	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	Id string `json:"id"`
	// The method to unregister for.
	Method string `json:"method"`
}
func (t *Unregistration) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["id"]; !exists {
		return fmt.Errorf("missing required field: id")
	}
	if _, exists := m["method"]; !exists {
		return fmt.Errorf("missing required field: method")
	}
	type Alias Unregistration
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = Unregistration(test)
	
	return nil
}

type UnregistrationParams struct {

	Unregisterations []Unregistration `json:"unregisterations"`
}
func (t *UnregistrationParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["unregisterations"]; !exists {
		return fmt.Errorf("missing required field: unregisterations")
	}
	type Alias UnregistrationParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = UnregistrationParams(test)
	
	return nil
}
// A versioned notebook document identifier.
// 
// @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {
	// The notebook document's uri.
	Uri URI `json:"uri"`
	// The version number of this notebook document.
	Version int32 `json:"version"`
}
func (t *VersionedNotebookDocumentIdentifier) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias VersionedNotebookDocumentIdentifier
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = VersionedNotebookDocumentIdentifier(test)
	
	return nil
}
// A text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	// The text document's uri.
	Uri DocumentUri `json:"uri"`
	// The version number of this document.
	Version int32 `json:"version"`
}
func (t *VersionedTextDocumentIdentifier) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias VersionedTextDocumentIdentifier
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = VersionedTextDocumentIdentifier(test)
	
	return nil
}
// The parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {
	// The 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason"`
	// The document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
func (t *WillSaveTextDocumentParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["reason"]; !exists {
		return fmt.Errorf("missing required field: reason")
	}
	if _, exists := m["textDocument"]; !exists {
		return fmt.Errorf("missing required field: textDocument")
	}
	type Alias WillSaveTextDocumentParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WillSaveTextDocumentParams(test)
	
	return nil
}

type WindowClientCapabilities struct {
	// Capabilities specific to the showDocument request.
	// 
	// @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitzero"`
	// Capabilities specific to the showMessage request.
	// 
	// @since 3.16.0
	ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitzero"`
	// It indicates whether the client supports server initiated
	// progress using the `window/workDoneProgress/create` request.
	// 
	// The capability also controls Whether client supports handling
	// of progress notifications. If set servers are allowed to report a
	// `workDoneProgress` property in the request specific server
	// capabilities.
	// 
	// @since 3.15.0
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *WindowClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WindowClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WindowClientCapabilities(test)
	
	return nil
}

type WorkDoneProgressBegin struct {
	// Controls if a cancel button should show to allow the user to cancel the
	// long running operation. Clients that don't support cancellation are allowed
	// to ignore the setting.
	Cancellable bool `json:"cancellable,omitempty"`

	Kind string `json:"kind"`
	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	// 
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`
	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	// 
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule. The value range is [0, 100].
	Percentage uint32 `json:"percentage,omitempty"`
	// Mandatory title of the progress operation. Used to briefly inform about
	// the kind of operation being performed.
	// 
	// Examples: "Indexing" or "Linking dependencies".
	Title string `json:"title"`
}
func (t *WorkDoneProgressBegin) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["title"]; !exists {
		return fmt.Errorf("missing required field: title")
	}
	type Alias WorkDoneProgressBegin
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressBegin(test)
	
	return nil
}

type WorkDoneProgressCancelParams struct {
	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}
func (t *WorkDoneProgressCancelParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["token"]; !exists {
		return fmt.Errorf("missing required field: token")
	}
	type Alias WorkDoneProgressCancelParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressCancelParams(test)
	
	return nil
}

type WorkDoneProgressCreateParams struct {
	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}
func (t *WorkDoneProgressCreateParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["token"]; !exists {
		return fmt.Errorf("missing required field: token")
	}
	type Alias WorkDoneProgressCreateParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressCreateParams(test)
	
	return nil
}

type WorkDoneProgressEnd struct {

	Kind string `json:"kind"`
	// Optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message string `json:"message,omitempty"`
}
func (t *WorkDoneProgressEnd) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias WorkDoneProgressEnd
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressEnd(test)
	
	return nil
}

type WorkDoneProgressOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *WorkDoneProgressOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkDoneProgressOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressOptions(test)
	
	return nil
}

type WorkDoneProgressParams struct {
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *WorkDoneProgressParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkDoneProgressParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressParams(test)
	
	return nil
}

type WorkDoneProgressReport struct {
	// Controls enablement state of a cancel button.
	// 
	// Clients that don't support cancellation or don't support controlling the button's
	// enablement state are allowed to ignore the property.
	Cancellable bool `json:"cancellable,omitempty"`

	Kind string `json:"kind"`
	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	// 
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message string `json:"message,omitempty"`
	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	// 
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule. The value range is [0, 100]
	Percentage uint32 `json:"percentage,omitempty"`
}
func (t *WorkDoneProgressReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	type Alias WorkDoneProgressReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkDoneProgressReport(test)
	
	return nil
}
// Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {
	// The client supports applying batch edits
	// to the workspace by supporting the request
	// 'workspace/applyEdit'
	ApplyEdit bool `json:"applyEdit,omitempty"`
	// Capabilities specific to the code lens requests scoped to the
	// workspace.
	// 
	// @since 3.16.0.
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitzero"`
	// The client supports `workspace/configuration` requests.
	// 
	// @since 3.6.0
	Configuration bool `json:"configuration,omitempty"`
	// Capabilities specific to the diagnostic requests scoped to the
	// workspace.
	// 
	// @since 3.17.0.
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitzero"`
	// Capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitzero"`
	// Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitzero"`
	// Capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitzero"`
	// The client has support for file notifications/requests for user operations on files.
	// 
	// Since 3.16.0
	FileOperations *FileOperationClientCapabilities `json:"fileOperations,omitzero"`
	// Capabilities specific to the folding range requests scoped to the workspace.
	// 
	// @since 3.18.0
	// @proposed
	FoldingRange *FoldingRangeWorkspaceClientCapabilities `json:"foldingRange,omitzero"`
	// Capabilities specific to the inlay hint requests scoped to the
	// workspace.
	// 
	// @since 3.17.0.
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint,omitzero"`
	// Capabilities specific to the inline values requests scoped to the
	// workspace.
	// 
	// @since 3.17.0.
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue,omitzero"`
	// Capabilities specific to the semantic token requests scoped to the
	// workspace.
	// 
	// @since 3.16.0.
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitzero"`
	// Capabilities specific to the `workspace/symbol` request.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitzero"`
	// Capabilities specific to the `workspace/textDocumentContent` request.
	// 
	// @since 3.18.0
	// @proposed
	TextDocumentContent *TextDocumentContentClientCapabilities `json:"textDocumentContent,omitzero"`
	// Capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitzero"`
	// The client has support for workspace folders.
	// 
	// @since 3.6.0
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`
}
func (t *WorkspaceClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceClientCapabilities(test)
	
	return nil
}
// Parameters of the workspace diagnostic request.
// 
// @since 3.17.0
type WorkspaceDiagnosticParams struct {
	// The additional identifier provided during registration.
	Identifier string `json:"identifier,omitempty"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// The currently known diagnostic reports with their
	// previous result ids.
	PreviousResultIds []PreviousResultId `json:"previousResultIds"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *WorkspaceDiagnosticParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["previousResultIds"]; !exists {
		return fmt.Errorf("missing required field: previousResultIds")
	}
	type Alias WorkspaceDiagnosticParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceDiagnosticParams(test)
	
	return nil
}
// A workspace diagnostic report.
// 
// @since 3.17.0
type WorkspaceDiagnosticReport struct {

	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}
func (t *WorkspaceDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	type Alias WorkspaceDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceDiagnosticReport(test)
	
	return nil
}
// A partial result for a workspace diagnostic report.
// 
// @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {

	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}
func (t *WorkspaceDiagnosticReportPartialResult) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	type Alias WorkspaceDiagnosticReportPartialResult
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceDiagnosticReportPartialResult(test)
	
	return nil
}
// A workspace edit represents changes to many resources managed in the workspace. The edit
// should either provide `changes` or `documentChanges`. If documentChanges are present
// they are preferred over `changes` if the client can handle versioned document edits.
// 
// Since version 3.13.0 a workspace edit can contain resource operations as well. If resource
// operations are present clients need to execute the operations in the order in which they
// are provided. So a workspace edit for example can consist of the following two changes:
// (1) a create file a.txt and (2) a text document edit which insert text into file a.txt.
// 
// An invalid sequence (e.g. (1) delete file a.txt and (2) insert text into file a.txt) will
// cause failure of the operation. How the client recovers from the failure is described by
// the client capability: `workspace.workspaceEdit.failureHandling`
type WorkspaceEdit struct {
	// A map of change annotations that can be referenced in `AnnotatedTextEdit`s or create, rename and
	// delete file / folder operations.
	// 
	// Whether clients honor this property depends on the client capability `workspace.changeAnnotationSupport`.
	// 
	// @since 3.16.0
	ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations,omitzero"`
	// Holds changes to existing resources.
	Changes map[DocumentUri][]TextEdit `json:"changes,omitzero"`
	// Depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes
	// are either an array of `TextDocumentEdit`s to express changes to n different text documents
	// where each text document edit addresses a specific version of a text document. Or it can contain
	// above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations.
	// 
	// Whether a client supports versioned document edits is expressed via
	// `workspace.workspaceEdit.documentChanges` client capability.
	// 
	// If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then
	// only plain `TextEdit`s using the `changes` property are supported.
	DocumentChanges []Or4[TextDocumentEdit, CreateFile, RenameFile, DeleteFile] `json:"documentChanges,omitzero"`
}
func (t *WorkspaceEdit) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceEdit
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceEdit(test)
	
	return nil
}

type WorkspaceEditClientCapabilities struct {
	// Whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	// 
	// @since 3.16.0
	ChangeAnnotationSupport *ChangeAnnotationsSupportOptions `json:"changeAnnotationSupport,omitzero"`
	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges bool `json:"documentChanges,omitempty"`
	// The failure handling strategy of a client if applying the workspace edit
	// fails.
	// 
	// @since 3.13.0
	FailureHandling *FailureHandlingKind `json:"failureHandling,omitzero"`
	// Whether the client supports `WorkspaceEditMetadata` in `WorkspaceEdit`s.
	// 
	// @since 3.18.0
	// @proposed
	MetadataSupport bool `json:"metadataSupport,omitempty"`
	// Whether the client normalizes line endings to the client specific
	// setting.
	// If set to `true` the client will normalize line ending characters
	// in a workspace edit to the client-specified new line
	// character.
	// 
	// @since 3.16.0
	NormalizesLineEndings bool `json:"normalizesLineEndings,omitempty"`
	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.
	// 
	// @since 3.13.0
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitzero"`
	// Whether the client supports snippets as text edits.
	// 
	// @since 3.18.0
	// @proposed
	SnippetEditSupport bool `json:"snippetEditSupport,omitempty"`
}
func (t *WorkspaceEditClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceEditClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceEditClientCapabilities(test)
	
	return nil
}
// Additional data about a workspace edit.
// 
// @since 3.18.0
// @proposed
type WorkspaceEditMetadata struct {
	// Signal to the editor that this edit is a refactoring.
	IsRefactoring bool `json:"isRefactoring,omitempty"`
}
func (t *WorkspaceEditMetadata) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceEditMetadata
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceEditMetadata(test)
	
	return nil
}
// A workspace folder inside a client.
type WorkspaceFolder struct {
	// The name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string `json:"name"`
	// The associated URI for this workspace folder.
	Uri URI `json:"uri"`
}
func (t *WorkspaceFolder) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	type Alias WorkspaceFolder
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceFolder(test)
	
	return nil
}
// The workspace folder change event.
type WorkspaceFoldersChangeEvent struct {
	// The array of added workspace folders
	Added []WorkspaceFolder `json:"added"`
	// The array of the removed workspace folders
	Removed []WorkspaceFolder `json:"removed"`
}
func (t *WorkspaceFoldersChangeEvent) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["added"]; !exists {
		return fmt.Errorf("missing required field: added")
	}
	if _, exists := m["removed"]; !exists {
		return fmt.Errorf("missing required field: removed")
	}
	type Alias WorkspaceFoldersChangeEvent
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceFoldersChangeEvent(test)
	
	return nil
}

type WorkspaceFoldersInitializeParams struct {
	// The workspace folders configured in the client when the server starts.
	// 
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	// 
	// @since 3.6.0
	WorkspaceFolders *[]WorkspaceFolder `json:"workspaceFolders,omitzero"`
}
func (t *WorkspaceFoldersInitializeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceFoldersInitializeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceFoldersInitializeParams(test)
	
	return nil
}

type WorkspaceFoldersServerCapabilities struct {
	// Whether the server wants to receive workspace folder
	// change notifications.
	// 
	// If a string is provided the string is treated as an ID
	// under which the notification is registered on the client
	// side. The ID can be used to unregister for these events
	// using the `client/unregisterCapability` request.
	ChangeNotifications *Or2[string, bool] `json:"changeNotifications,omitzero"`
	// The server has support for workspace folders
	Supported bool `json:"supported,omitempty"`
}
func (t *WorkspaceFoldersServerCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceFoldersServerCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceFoldersServerCapabilities(test)
	
	return nil
}
// A full document diagnostic report for a workspace diagnostic result.
// 
// @since 3.17.0
type WorkspaceFullDocumentDiagnosticReport struct {
	// The actual items.
	Items []Diagnostic `json:"items"`
	// A full document diagnostic report.
	Kind string `json:"kind"`
	// An optional result id. If provided it will
	// be sent on the next diagnostic request for the
	// same document.
	ResultId string `json:"resultId,omitempty"`
	// The URI for which diagnostic information is reported.
	Uri DocumentUri `json:"uri"`
	// The version number for which the diagnostics are reported.
	// If the document is not marked as open `null` can be provided.
	Version *int32 `json:"version"`
}
func (t *WorkspaceFullDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["items"]; !exists {
		return fmt.Errorf("missing required field: items")
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias WorkspaceFullDocumentDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceFullDocumentDiagnosticReport(test)
	
	return nil
}
// Defines workspace specific capabilities of the server.
// 
// @since 3.18.0
type WorkspaceOptions struct {
	// The server is interested in notifications/requests for operations on files.
	// 
	// @since 3.16.0
	FileOperations *FileOperationOptions `json:"fileOperations,omitzero"`
	// The server supports the `workspace/textDocumentContent` request.
	// 
	// @since 3.18.0
	// @proposed
	TextDocumentContent *Or2[TextDocumentContentOptions, TextDocumentContentRegistrationOptions] `json:"textDocumentContent,omitzero"`
	// The server supports workspace folder.
	// 
	// @since 3.6.0
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitzero"`
}
func (t *WorkspaceOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceOptions(test)
	
	return nil
}
// A special workspace symbol that supports locations without a range.
// 
// See also SymbolInformation.
// 
// @since 3.17.0
type WorkspaceSymbol struct {
	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName string `json:"containerName,omitempty"`
	// A data entry field that is preserved on a workspace symbol between a
	// workspace symbol request and a workspace symbol resolve request.
	Data any `json:"data,omitempty"`
	// The kind of this symbol.
	Kind SymbolKind `json:"kind"`
	// The location of the symbol. Whether a server is allowed to
	// return a location without a range depends on the client
	// capability `workspace.symbol.resolveSupport`.
	// 
	// See SymbolInformation#location for more details.
	Location Or2[Location, LocationUriOnly] `json:"location"`
	// The name of this symbol.
	Name string `json:"name"`
	// Tags for this symbol.
	// 
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitzero"`
}
func (t *WorkspaceSymbol) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["location"]; !exists {
		return fmt.Errorf("missing required field: location")
	}
	if _, exists := m["name"]; !exists {
		return fmt.Errorf("missing required field: name")
	}
	type Alias WorkspaceSymbol
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceSymbol(test)
	
	return nil
}
// Client capabilities for a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolClientCapabilities struct {
	// Symbol request supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client support partial workspace symbols. The client will send the
	// request `workspaceSymbol/resolve` to the server to resolve additional
	// properties.
	// 
	// @since 3.17.0
	ResolveSupport *ClientSymbolResolveOptions `json:"resolveSupport,omitzero"`
	// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbolKind,omitzero"`
	// The client supports tags on `SymbolInformation`.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.16.0
	TagSupport *ClientSymbolTagOptions `json:"tagSupport,omitzero"`
}
func (t *WorkspaceSymbolClientCapabilities) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceSymbolClientCapabilities
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceSymbolClientCapabilities(test)
	
	return nil
}
// Server capabilities for a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolOptions struct {
	// The server provides support to resolve additional
	// information for a workspace symbol.
	// 
	// @since 3.17.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *WorkspaceSymbolOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceSymbolOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceSymbolOptions(test)
	
	return nil
}
// The parameters of a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitzero"`
	// A query string to filter symbols by. Clients may send an empty
	// string here to request all symbols.
	// 
	// The `query`-parameter should be interpreted in a *relaxed way* as editors
	// will apply their own highlighting and scoring on the results. A good rule
	// of thumb is to match case-insensitive and to simply check that the
	// characters of *query* appear in their order in a candidate symbol.
	// Servers shouldn't use prefix, substring, or similar strict matching.
	Query string `json:"query"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *WorkspaceSymbolParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["query"]; !exists {
		return fmt.Errorf("missing required field: query")
	}
	type Alias WorkspaceSymbolParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceSymbolParams(test)
	
	return nil
}
// Registration options for a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolRegistrationOptions struct {
	// The server provides support to resolve additional
	// information for a workspace symbol.
	// 
	// @since 3.17.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
func (t *WorkspaceSymbolRegistrationOptions) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}

	type Alias WorkspaceSymbolRegistrationOptions
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceSymbolRegistrationOptions(test)
	
	return nil
}
// An unchanged document diagnostic report for a workspace diagnostic result.
// 
// @since 3.17.0
type WorkspaceUnchangedDocumentDiagnosticReport struct {
	// A document diagnostic report indicating
	// no changes to the last result. A server can
	// only return `unchanged` if result ids are
	// provided.
	Kind string `json:"kind"`
	// A result id which will be sent on the next
	// diagnostic request for the same document.
	ResultId string `json:"resultId"`
	// The URI for which diagnostic information is reported.
	Uri DocumentUri `json:"uri"`
	// The version number for which the diagnostics are reported.
	// If the document is not marked as open `null` can be provided.
	Version *int32 `json:"version"`
}
func (t *WorkspaceUnchangedDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["kind"]; !exists {
		return fmt.Errorf("missing required field: kind")
	}
	if _, exists := m["resultId"]; !exists {
		return fmt.Errorf("missing required field: resultId")
	}
	if _, exists := m["uri"]; !exists {
		return fmt.Errorf("missing required field: uri")
	}
	if _, exists := m["version"]; !exists {
		return fmt.Errorf("missing required field: version")
	}
	type Alias WorkspaceUnchangedDocumentDiagnosticReport
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = WorkspaceUnchangedDocumentDiagnosticReport(test)
	
	return nil
}
// The initialize parameters
type _InitializeParams struct {
	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`
	// Information about the client
	// 
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitzero"`
	// User provided initialization options.
	InitializationOptions any `json:"initializationOptions,omitempty"`
	// The locale the client is currently showing the user interface
	// in. This must not necessarily be the locale of the operating
	// system.
	// 
	// Uses IETF language tags as the value's syntax
	// (See https://en.wikipedia.org/wiki/IETF_language_tag)
	// 
	// @since 3.16.0
	Locale string `json:"locale,omitempty"`
	// The process Id of the parent process that started
	// the server.
	// 
	// Is `null` if the process has not been started by another process.
	// If the parent process is not alive then the server should exit.
	ProcessId *int32 `json:"processId"`
	// The rootPath of the workspace. Is null
	// if no folder is open.
	// 
	// @deprecated in favour of rootUri.
	RootPath **string `json:"rootPath,omitzero"`
	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	// 
	// @deprecated in favour of workspaceFolders.
	RootUri *DocumentUri `json:"rootUri"`
	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitzero"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitzero"`
}
func (t *_InitializeParams) UnmarshalJSON(x []byte) error {
	var m map[string]any
	if err := json.Unmarshal(x, &m); err != nil {
		return err
	}
	if _, exists := m["capabilities"]; !exists {
		return fmt.Errorf("missing required field: capabilities")
	}
	if _, exists := m["processId"]; !exists {
		return fmt.Errorf("missing required field: processId")
	}
	if _, exists := m["rootUri"]; !exists {
		return fmt.Errorf("missing required field: rootUri")
	}
	type Alias _InitializeParams
	var test Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&test); err != nil {
		return err
	}
	*t = _InitializeParams(test)
	
	return nil
}
const (
	UnknownRequestMethod MethodKind = ""
	CallHierarchyIncomingCallsMethod MethodKind = "callHierarchy/incomingCalls"
	CallHierarchyOutgoingCallsMethod MethodKind = "callHierarchy/outgoingCalls"
	ClientRegisterCapabilityMethod MethodKind = "client/registerCapability"
	ClientUnregisterCapabilityMethod MethodKind = "client/unregisterCapability"
	CodeActionResolveMethod MethodKind = "codeAction/resolve"
	CodeLensResolveMethod MethodKind = "codeLens/resolve"
	CompletionItemResolveMethod MethodKind = "completionItem/resolve"
	DocumentLinkResolveMethod MethodKind = "documentLink/resolve"
	InitializeMethod MethodKind = "initialize"
	InlayHintResolveMethod MethodKind = "inlayHint/resolve"
	ShutdownMethod MethodKind = "shutdown"
	TextDocumentCodeActionMethod MethodKind = "textDocument/codeAction"
	TextDocumentCodeLensMethod MethodKind = "textDocument/codeLens"
	TextDocumentColorPresentationMethod MethodKind = "textDocument/colorPresentation"
	TextDocumentCompletionMethod MethodKind = "textDocument/completion"
	TextDocumentDeclarationMethod MethodKind = "textDocument/declaration"
	TextDocumentDefinitionMethod MethodKind = "textDocument/definition"
	TextDocumentDiagnosticMethod MethodKind = "textDocument/diagnostic"
	TextDocumentDocumentColorMethod MethodKind = "textDocument/documentColor"
	TextDocumentDocumentHighlightMethod MethodKind = "textDocument/documentHighlight"
	TextDocumentDocumentLinkMethod MethodKind = "textDocument/documentLink"
	TextDocumentDocumentSymbolMethod MethodKind = "textDocument/documentSymbol"
	TextDocumentFoldingRangeMethod MethodKind = "textDocument/foldingRange"
	TextDocumentFormattingMethod MethodKind = "textDocument/formatting"
	TextDocumentHoverMethod MethodKind = "textDocument/hover"
	TextDocumentImplementationMethod MethodKind = "textDocument/implementation"
	TextDocumentInlayHintMethod MethodKind = "textDocument/inlayHint"
	TextDocumentInlineCompletionMethod MethodKind = "textDocument/inlineCompletion"
	TextDocumentInlineValueMethod MethodKind = "textDocument/inlineValue"
	TextDocumentLinkedEditingRangeMethod MethodKind = "textDocument/linkedEditingRange"
	TextDocumentMonikerMethod MethodKind = "textDocument/moniker"
	TextDocumentOnTypeFormattingMethod MethodKind = "textDocument/onTypeFormatting"
	TextDocumentPrepareCallHierarchyMethod MethodKind = "textDocument/prepareCallHierarchy"
	TextDocumentPrepareRenameMethod MethodKind = "textDocument/prepareRename"
	TextDocumentPrepareTypeHierarchyMethod MethodKind = "textDocument/prepareTypeHierarchy"
	TextDocumentRangeFormattingMethod MethodKind = "textDocument/rangeFormatting"
	TextDocumentRangesFormattingMethod MethodKind = "textDocument/rangesFormatting"
	TextDocumentReferencesMethod MethodKind = "textDocument/references"
	TextDocumentRenameMethod MethodKind = "textDocument/rename"
	TextDocumentSelectionRangeMethod MethodKind = "textDocument/selectionRange"
	TextDocumentSemanticTokensFullMethod MethodKind = "textDocument/semanticTokens/full"
	TextDocumentSemanticTokensFullDeltaMethod MethodKind = "textDocument/semanticTokens/full/delta"
	TextDocumentSemanticTokensRangeMethod MethodKind = "textDocument/semanticTokens/range"
	TextDocumentSignatureHelpMethod MethodKind = "textDocument/signatureHelp"
	TextDocumentTypeDefinitionMethod MethodKind = "textDocument/typeDefinition"
	TextDocumentWillSaveWaitUntilMethod MethodKind = "textDocument/willSaveWaitUntil"
	TypeHierarchySubtypesMethod MethodKind = "typeHierarchy/subtypes"
	TypeHierarchySupertypesMethod MethodKind = "typeHierarchy/supertypes"
	WindowShowDocumentMethod MethodKind = "window/showDocument"
	WindowShowMessageRequestMethod MethodKind = "window/showMessageRequest"
	WindowWorkDoneProgressCreateMethod MethodKind = "window/workDoneProgress/create"
	WorkspaceApplyEditMethod MethodKind = "workspace/applyEdit"
	WorkspaceCodeLensRefreshMethod MethodKind = "workspace/codeLens/refresh"
	WorkspaceConfigurationMethod MethodKind = "workspace/configuration"
	WorkspaceDiagnosticMethod MethodKind = "workspace/diagnostic"
	WorkspaceDiagnosticRefreshMethod MethodKind = "workspace/diagnostic/refresh"
	WorkspaceExecuteCommandMethod MethodKind = "workspace/executeCommand"
	WorkspaceFoldingRangeRefreshMethod MethodKind = "workspace/foldingRange/refresh"
	WorkspaceInlayHintRefreshMethod MethodKind = "workspace/inlayHint/refresh"
	WorkspaceInlineValueRefreshMethod MethodKind = "workspace/inlineValue/refresh"
	WorkspaceSemanticTokensRefreshMethod MethodKind = "workspace/semanticTokens/refresh"
	WorkspaceSymbolMethod MethodKind = "workspace/symbol"
	WorkspaceTextDocumentContentMethod MethodKind = "workspace/textDocumentContent"
	WorkspaceTextDocumentContentRefreshMethod MethodKind = "workspace/textDocumentContent/refresh"
	WorkspaceWillCreateFilesMethod MethodKind = "workspace/willCreateFiles"
	WorkspaceWillDeleteFilesMethod MethodKind = "workspace/willDeleteFiles"
	WorkspaceWillRenameFilesMethod MethodKind = "workspace/willRenameFiles"
	WorkspaceWorkspaceFoldersMethod MethodKind = "workspace/workspaceFolders"
	WorkspaceSymbolResolveMethod MethodKind = "workspaceSymbol/resolve"
)

var RequestMethodMap = map[string]MethodKind{
	"callHierarchy/incomingCalls": CallHierarchyIncomingCallsMethod,
	"callHierarchy/outgoingCalls": CallHierarchyOutgoingCallsMethod,
	"client/registerCapability": ClientRegisterCapabilityMethod,
	"client/unregisterCapability": ClientUnregisterCapabilityMethod,
	"codeAction/resolve": CodeActionResolveMethod,
	"codeLens/resolve": CodeLensResolveMethod,
	"completionItem/resolve": CompletionItemResolveMethod,
	"documentLink/resolve": DocumentLinkResolveMethod,
	"initialize": InitializeMethod,
	"inlayHint/resolve": InlayHintResolveMethod,
	"shutdown": ShutdownMethod,
	"textDocument/codeAction": TextDocumentCodeActionMethod,
	"textDocument/codeLens": TextDocumentCodeLensMethod,
	"textDocument/colorPresentation": TextDocumentColorPresentationMethod,
	"textDocument/completion": TextDocumentCompletionMethod,
	"textDocument/declaration": TextDocumentDeclarationMethod,
	"textDocument/definition": TextDocumentDefinitionMethod,
	"textDocument/diagnostic": TextDocumentDiagnosticMethod,
	"textDocument/documentColor": TextDocumentDocumentColorMethod,
	"textDocument/documentHighlight": TextDocumentDocumentHighlightMethod,
	"textDocument/documentLink": TextDocumentDocumentLinkMethod,
	"textDocument/documentSymbol": TextDocumentDocumentSymbolMethod,
	"textDocument/foldingRange": TextDocumentFoldingRangeMethod,
	"textDocument/formatting": TextDocumentFormattingMethod,
	"textDocument/hover": TextDocumentHoverMethod,
	"textDocument/implementation": TextDocumentImplementationMethod,
	"textDocument/inlayHint": TextDocumentInlayHintMethod,
	"textDocument/inlineCompletion": TextDocumentInlineCompletionMethod,
	"textDocument/inlineValue": TextDocumentInlineValueMethod,
	"textDocument/linkedEditingRange": TextDocumentLinkedEditingRangeMethod,
	"textDocument/moniker": TextDocumentMonikerMethod,
	"textDocument/onTypeFormatting": TextDocumentOnTypeFormattingMethod,
	"textDocument/prepareCallHierarchy": TextDocumentPrepareCallHierarchyMethod,
	"textDocument/prepareRename": TextDocumentPrepareRenameMethod,
	"textDocument/prepareTypeHierarchy": TextDocumentPrepareTypeHierarchyMethod,
	"textDocument/rangeFormatting": TextDocumentRangeFormattingMethod,
	"textDocument/rangesFormatting": TextDocumentRangesFormattingMethod,
	"textDocument/references": TextDocumentReferencesMethod,
	"textDocument/rename": TextDocumentRenameMethod,
	"textDocument/selectionRange": TextDocumentSelectionRangeMethod,
	"textDocument/semanticTokens/full": TextDocumentSemanticTokensFullMethod,
	"textDocument/semanticTokens/full/delta": TextDocumentSemanticTokensFullDeltaMethod,
	"textDocument/semanticTokens/range": TextDocumentSemanticTokensRangeMethod,
	"textDocument/signatureHelp": TextDocumentSignatureHelpMethod,
	"textDocument/typeDefinition": TextDocumentTypeDefinitionMethod,
	"textDocument/willSaveWaitUntil": TextDocumentWillSaveWaitUntilMethod,
	"typeHierarchy/subtypes": TypeHierarchySubtypesMethod,
	"typeHierarchy/supertypes": TypeHierarchySupertypesMethod,
	"window/showDocument": WindowShowDocumentMethod,
	"window/showMessageRequest": WindowShowMessageRequestMethod,
	"window/workDoneProgress/create": WindowWorkDoneProgressCreateMethod,
	"workspace/applyEdit": WorkspaceApplyEditMethod,
	"workspace/codeLens/refresh": WorkspaceCodeLensRefreshMethod,
	"workspace/configuration": WorkspaceConfigurationMethod,
	"workspace/diagnostic": WorkspaceDiagnosticMethod,
	"workspace/diagnostic/refresh": WorkspaceDiagnosticRefreshMethod,
	"workspace/executeCommand": WorkspaceExecuteCommandMethod,
	"workspace/foldingRange/refresh": WorkspaceFoldingRangeRefreshMethod,
	"workspace/inlayHint/refresh": WorkspaceInlayHintRefreshMethod,
	"workspace/inlineValue/refresh": WorkspaceInlineValueRefreshMethod,
	"workspace/semanticTokens/refresh": WorkspaceSemanticTokensRefreshMethod,
	"workspace/symbol": WorkspaceSymbolMethod,
	"workspace/textDocumentContent": WorkspaceTextDocumentContentMethod,
	"workspace/textDocumentContent/refresh": WorkspaceTextDocumentContentRefreshMethod,
	"workspace/willCreateFiles": WorkspaceWillCreateFilesMethod,
	"workspace/willDeleteFiles": WorkspaceWillDeleteFilesMethod,
	"workspace/willRenameFiles": WorkspaceWillRenameFilesMethod,
	"workspace/workspaceFolders": WorkspaceWorkspaceFoldersMethod,
	"workspaceSymbol/resolve": WorkspaceSymbolResolveMethod,
}

// A request to resolve the incoming calls for a given `CallHierarchyItem`.
// 
// @since 3.16.0
type CallHierarchyIncomingCallsRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CallHierarchyIncomingCallsParams `json:"params"`
}
func (t CallHierarchyIncomingCallsRequest) isMessage() {}
func (t CallHierarchyIncomingCallsRequest) isRequest() {}
func (t CallHierarchyIncomingCallsRequest) GetMethod() MethodKind { return t.Method }
func (t CallHierarchyIncomingCallsRequest) GetParams() CallHierarchyIncomingCallsParams { return t.Params }
func (t *CallHierarchyIncomingCallsRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CallHierarchyIncomingCallsRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CallHierarchyIncomingCallsRequest(result)
   return nil
}
// A request to resolve the outgoing calls for a given `CallHierarchyItem`.
// 
// @since 3.16.0
type CallHierarchyOutgoingCallsRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CallHierarchyOutgoingCallsParams `json:"params"`
}
func (t CallHierarchyOutgoingCallsRequest) isMessage() {}
func (t CallHierarchyOutgoingCallsRequest) isRequest() {}
func (t CallHierarchyOutgoingCallsRequest) GetMethod() MethodKind { return t.Method }
func (t CallHierarchyOutgoingCallsRequest) GetParams() CallHierarchyOutgoingCallsParams { return t.Params }
func (t *CallHierarchyOutgoingCallsRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CallHierarchyOutgoingCallsRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CallHierarchyOutgoingCallsRequest(result)
   return nil
}
// The `client/registerCapability` request is sent from the server to the client to register a new capability
// handler on the client side.
type RegistrationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params RegistrationParams `json:"params"`
}
func (t RegistrationRequest) isMessage() {}
func (t RegistrationRequest) isRequest() {}
func (t RegistrationRequest) GetMethod() MethodKind { return t.Method }
func (t RegistrationRequest) GetParams() RegistrationParams { return t.Params }
func (t *RegistrationRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias RegistrationRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = RegistrationRequest(result)
   return nil
}
// The `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability
// handler on the client side.
type UnregistrationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params UnregistrationParams `json:"params"`
}
func (t UnregistrationRequest) isMessage() {}
func (t UnregistrationRequest) isRequest() {}
func (t UnregistrationRequest) GetMethod() MethodKind { return t.Method }
func (t UnregistrationRequest) GetParams() UnregistrationParams { return t.Params }
func (t *UnregistrationRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias UnregistrationRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = UnregistrationRequest(result)
   return nil
}
// Request to resolve additional information for a given code action.The request's
// parameter is of type {@link CodeAction} the response
// is of type {@link CodeAction} or a Thenable that resolves to such.
type CodeActionResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CodeAction `json:"params"`
}
func (t CodeActionResolveRequest) isMessage() {}
func (t CodeActionResolveRequest) isRequest() {}
func (t CodeActionResolveRequest) GetMethod() MethodKind { return t.Method }
func (t CodeActionResolveRequest) GetParams() CodeAction { return t.Params }
func (t *CodeActionResolveRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CodeActionResolveRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CodeActionResolveRequest(result)
   return nil
}
// A request to resolve a command for a given code lens.
type CodeLensResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CodeLens `json:"params"`
}
func (t CodeLensResolveRequest) isMessage() {}
func (t CodeLensResolveRequest) isRequest() {}
func (t CodeLensResolveRequest) GetMethod() MethodKind { return t.Method }
func (t CodeLensResolveRequest) GetParams() CodeLens { return t.Params }
func (t *CodeLensResolveRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CodeLensResolveRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CodeLensResolveRequest(result)
   return nil
}
// Request to resolve additional information for a given completion item.The request's
// parameter is of type {@link CompletionItem} the response
// is of type {@link CompletionItem} or a Thenable that resolves to such.
type CompletionResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CompletionItem `json:"params"`
}
func (t CompletionResolveRequest) isMessage() {}
func (t CompletionResolveRequest) isRequest() {}
func (t CompletionResolveRequest) GetMethod() MethodKind { return t.Method }
func (t CompletionResolveRequest) GetParams() CompletionItem { return t.Params }
func (t *CompletionResolveRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CompletionResolveRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CompletionResolveRequest(result)
   return nil
}
// Request to resolve additional information for a given document link. The request's
// parameter is of type {@link DocumentLink} the response
// is of type {@link DocumentLink} or a Thenable that resolves to such.
type DocumentLinkResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentLink `json:"params"`
}
func (t DocumentLinkResolveRequest) isMessage() {}
func (t DocumentLinkResolveRequest) isRequest() {}
func (t DocumentLinkResolveRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentLinkResolveRequest) GetParams() DocumentLink { return t.Params }
func (t *DocumentLinkResolveRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentLinkResolveRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentLinkResolveRequest(result)
   return nil
}
// The initialize request is sent from the client to the server.
// It is sent once as the request after starting up the server.
// The requests parameter is of type {@link InitializeParams}
// the response if of type {@link InitializeResult} of a Thenable that
// resolves to such.
type InitializeRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params InitializeParams `json:"params"`
}
func (t InitializeRequest) isMessage() {}
func (t InitializeRequest) isRequest() {}
func (t InitializeRequest) GetMethod() MethodKind { return t.Method }
func (t InitializeRequest) GetParams() InitializeParams { return t.Params }
func (t *InitializeRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InitializeRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InitializeRequest(result)
   return nil
}
// A request to resolve additional properties for an inlay hint.
// The request's parameter is of type {@link InlayHint}, the response is
// of type {@link InlayHint} or a Thenable that resolves to such.
// 
// @since 3.17.0
type InlayHintResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params InlayHint `json:"params"`
}
func (t InlayHintResolveRequest) isMessage() {}
func (t InlayHintResolveRequest) isRequest() {}
func (t InlayHintResolveRequest) GetMethod() MethodKind { return t.Method }
func (t InlayHintResolveRequest) GetParams() InlayHint { return t.Params }
func (t *InlayHintResolveRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InlayHintResolveRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InlayHintResolveRequest(result)
   return nil
}
// A shutdown request is sent from the client to the server.
// It is sent once when the client decides to shutdown the
// server. The only notification that is sent after a shutdown request
// is the exit event.
type ShutdownRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t ShutdownRequest) isMessage() {}
func (t ShutdownRequest) isRequest() {}
func (t ShutdownRequest) GetMethod() MethodKind { return t.Method }
func (t ShutdownRequest) GetParams() any { return t.Params }
func (t *ShutdownRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ShutdownRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ShutdownRequest(result)
   return nil
}
// A request to provide commands for the given text document and range.
type CodeActionRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CodeActionParams `json:"params"`
}
func (t CodeActionRequest) isMessage() {}
func (t CodeActionRequest) isRequest() {}
func (t CodeActionRequest) GetMethod() MethodKind { return t.Method }
func (t CodeActionRequest) GetParams() CodeActionParams { return t.Params }
func (t *CodeActionRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CodeActionRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CodeActionRequest(result)
   return nil
}
// A request to provide code lens for the given text document.
type CodeLensRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CodeLensParams `json:"params"`
}
func (t CodeLensRequest) isMessage() {}
func (t CodeLensRequest) isRequest() {}
func (t CodeLensRequest) GetMethod() MethodKind { return t.Method }
func (t CodeLensRequest) GetParams() CodeLensParams { return t.Params }
func (t *CodeLensRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CodeLensRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CodeLensRequest(result)
   return nil
}
// A request to list all presentation for a color. The request's
// parameter is of type {@link ColorPresentationParams} the
// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
// that resolves to such.
type ColorPresentationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ColorPresentationParams `json:"params"`
}
func (t ColorPresentationRequest) isMessage() {}
func (t ColorPresentationRequest) isRequest() {}
func (t ColorPresentationRequest) GetMethod() MethodKind { return t.Method }
func (t ColorPresentationRequest) GetParams() ColorPresentationParams { return t.Params }
func (t *ColorPresentationRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ColorPresentationRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ColorPresentationRequest(result)
   return nil
}
// Request to request completion at a given text document position. The request's
// parameter is of type {@link TextDocumentPosition} the response
// is of type {@link CompletionItem CompletionItem[]} or {@link CompletionList}
// or a Thenable that resolves to such.
// 
// The request can delay the computation of the {@link CompletionItem.detail `detail`}
// and {@link CompletionItem.documentation `documentation`} properties to the `completionItem/resolve`
// request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
type CompletionRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CompletionParams `json:"params"`
}
func (t CompletionRequest) isMessage() {}
func (t CompletionRequest) isRequest() {}
func (t CompletionRequest) GetMethod() MethodKind { return t.Method }
func (t CompletionRequest) GetParams() CompletionParams { return t.Params }
func (t *CompletionRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CompletionRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CompletionRequest(result)
   return nil
}
// A request to resolve the type definition locations of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPositionParams}
// the response is of type {@link Declaration} or a typed array of {@link DeclarationLink}
// or a Thenable that resolves to such.
type DeclarationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DeclarationParams `json:"params"`
}
func (t DeclarationRequest) isMessage() {}
func (t DeclarationRequest) isRequest() {}
func (t DeclarationRequest) GetMethod() MethodKind { return t.Method }
func (t DeclarationRequest) GetParams() DeclarationParams { return t.Params }
func (t *DeclarationRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DeclarationRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DeclarationRequest(result)
   return nil
}
// A request to resolve the definition location of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPosition}
// the response is of either type {@link Definition} or a typed array of
// {@link DefinitionLink} or a Thenable that resolves to such.
type DefinitionRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DefinitionParams `json:"params"`
}
func (t DefinitionRequest) isMessage() {}
func (t DefinitionRequest) isRequest() {}
func (t DefinitionRequest) GetMethod() MethodKind { return t.Method }
func (t DefinitionRequest) GetParams() DefinitionParams { return t.Params }
func (t *DefinitionRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DefinitionRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DefinitionRequest(result)
   return nil
}
// The document diagnostic request definition.
// 
// @since 3.17.0
type DocumentDiagnosticRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentDiagnosticParams `json:"params"`
}
func (t DocumentDiagnosticRequest) isMessage() {}
func (t DocumentDiagnosticRequest) isRequest() {}
func (t DocumentDiagnosticRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentDiagnosticRequest) GetParams() DocumentDiagnosticParams { return t.Params }
func (t *DocumentDiagnosticRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentDiagnosticRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentDiagnosticRequest(result)
   return nil
}
// A request to list all color symbols found in a given text document. The request's
// parameter is of type {@link DocumentColorParams} the
// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
// that resolves to such.
type DocumentColorRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentColorParams `json:"params"`
}
func (t DocumentColorRequest) isMessage() {}
func (t DocumentColorRequest) isRequest() {}
func (t DocumentColorRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentColorRequest) GetParams() DocumentColorParams { return t.Params }
func (t *DocumentColorRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentColorRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentColorRequest(result)
   return nil
}
// Request to resolve a {@link DocumentHighlight} for a given
// text document position. The request's parameter is of type {@link TextDocumentPosition}
// the request response is an array of type {@link DocumentHighlight}
// or a Thenable that resolves to such.
type DocumentHighlightRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentHighlightParams `json:"params"`
}
func (t DocumentHighlightRequest) isMessage() {}
func (t DocumentHighlightRequest) isRequest() {}
func (t DocumentHighlightRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentHighlightRequest) GetParams() DocumentHighlightParams { return t.Params }
func (t *DocumentHighlightRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentHighlightRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentHighlightRequest(result)
   return nil
}
// A request to provide document links
type DocumentLinkRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentLinkParams `json:"params"`
}
func (t DocumentLinkRequest) isMessage() {}
func (t DocumentLinkRequest) isRequest() {}
func (t DocumentLinkRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentLinkRequest) GetParams() DocumentLinkParams { return t.Params }
func (t *DocumentLinkRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentLinkRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentLinkRequest(result)
   return nil
}
// A request to list all symbols found in a given text document. The request's
// parameter is of type {@link TextDocumentIdentifier} the
// response is of type {@link SymbolInformation SymbolInformation[]} or a Thenable
// that resolves to such.
type DocumentSymbolRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentSymbolParams `json:"params"`
}
func (t DocumentSymbolRequest) isMessage() {}
func (t DocumentSymbolRequest) isRequest() {}
func (t DocumentSymbolRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentSymbolRequest) GetParams() DocumentSymbolParams { return t.Params }
func (t *DocumentSymbolRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentSymbolRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentSymbolRequest(result)
   return nil
}
// A request to provide folding ranges in a document. The request's
// parameter is of type {@link FoldingRangeParams}, the
// response is of type {@link FoldingRangeList} or a Thenable
// that resolves to such.
type FoldingRangeRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params FoldingRangeParams `json:"params"`
}
func (t FoldingRangeRequest) isMessage() {}
func (t FoldingRangeRequest) isRequest() {}
func (t FoldingRangeRequest) GetMethod() MethodKind { return t.Method }
func (t FoldingRangeRequest) GetParams() FoldingRangeParams { return t.Params }
func (t *FoldingRangeRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias FoldingRangeRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = FoldingRangeRequest(result)
   return nil
}
// A request to format a whole document.
type DocumentFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentFormattingParams `json:"params"`
}
func (t DocumentFormattingRequest) isMessage() {}
func (t DocumentFormattingRequest) isRequest() {}
func (t DocumentFormattingRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentFormattingRequest) GetParams() DocumentFormattingParams { return t.Params }
func (t *DocumentFormattingRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentFormattingRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentFormattingRequest(result)
   return nil
}
// Request to request hover information at a given text document position. The request's
// parameter is of type {@link TextDocumentPosition} the response is of
// type {@link Hover} or a Thenable that resolves to such.
type HoverRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params HoverParams `json:"params"`
}
func (t HoverRequest) isMessage() {}
func (t HoverRequest) isRequest() {}
func (t HoverRequest) GetMethod() MethodKind { return t.Method }
func (t HoverRequest) GetParams() HoverParams { return t.Params }
func (t *HoverRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias HoverRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = HoverRequest(result)
   return nil
}
// A request to resolve the implementation locations of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPositionParams}
// the response is of type {@link Definition} or a Thenable that resolves to such.
type ImplementationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ImplementationParams `json:"params"`
}
func (t ImplementationRequest) isMessage() {}
func (t ImplementationRequest) isRequest() {}
func (t ImplementationRequest) GetMethod() MethodKind { return t.Method }
func (t ImplementationRequest) GetParams() ImplementationParams { return t.Params }
func (t *ImplementationRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ImplementationRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ImplementationRequest(result)
   return nil
}
// A request to provide inlay hints in a document. The request's parameter is of
// type {@link InlayHintsParams}, the response is of type
// {@link InlayHint InlayHint[]} or a Thenable that resolves to such.
// 
// @since 3.17.0
type InlayHintRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params InlayHintParams `json:"params"`
}
func (t InlayHintRequest) isMessage() {}
func (t InlayHintRequest) isRequest() {}
func (t InlayHintRequest) GetMethod() MethodKind { return t.Method }
func (t InlayHintRequest) GetParams() InlayHintParams { return t.Params }
func (t *InlayHintRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InlayHintRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InlayHintRequest(result)
   return nil
}
// A request to provide inline completions in a document. The request's parameter is of
// type {@link InlineCompletionParams}, the response is of type
// {@link InlineCompletion InlineCompletion[]} or a Thenable that resolves to such.
// 
// @since 3.18.0
// @proposed
type InlineCompletionRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params InlineCompletionParams `json:"params"`
}
func (t InlineCompletionRequest) isMessage() {}
func (t InlineCompletionRequest) isRequest() {}
func (t InlineCompletionRequest) GetMethod() MethodKind { return t.Method }
func (t InlineCompletionRequest) GetParams() InlineCompletionParams { return t.Params }
func (t *InlineCompletionRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InlineCompletionRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InlineCompletionRequest(result)
   return nil
}
// A request to provide inline values in a document. The request's parameter is of
// type {@link InlineValueParams}, the response is of type
// {@link InlineValue InlineValue[]} or a Thenable that resolves to such.
// 
// @since 3.17.0
type InlineValueRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params InlineValueParams `json:"params"`
}
func (t InlineValueRequest) isMessage() {}
func (t InlineValueRequest) isRequest() {}
func (t InlineValueRequest) GetMethod() MethodKind { return t.Method }
func (t InlineValueRequest) GetParams() InlineValueParams { return t.Params }
func (t *InlineValueRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InlineValueRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InlineValueRequest(result)
   return nil
}
// A request to provide ranges that can be edited together.
// 
// @since 3.16.0
type LinkedEditingRangeRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params LinkedEditingRangeParams `json:"params"`
}
func (t LinkedEditingRangeRequest) isMessage() {}
func (t LinkedEditingRangeRequest) isRequest() {}
func (t LinkedEditingRangeRequest) GetMethod() MethodKind { return t.Method }
func (t LinkedEditingRangeRequest) GetParams() LinkedEditingRangeParams { return t.Params }
func (t *LinkedEditingRangeRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias LinkedEditingRangeRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = LinkedEditingRangeRequest(result)
   return nil
}
// A request to get the moniker of a symbol at a given text document position.
// The request parameter is of type {@link TextDocumentPositionParams}.
// The response is of type {@link Moniker Moniker[]} or `null`.
type MonikerRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params MonikerParams `json:"params"`
}
func (t MonikerRequest) isMessage() {}
func (t MonikerRequest) isRequest() {}
func (t MonikerRequest) GetMethod() MethodKind { return t.Method }
func (t MonikerRequest) GetParams() MonikerParams { return t.Params }
func (t *MonikerRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias MonikerRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = MonikerRequest(result)
   return nil
}
// A request to format a document on type.
type DocumentOnTypeFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentOnTypeFormattingParams `json:"params"`
}
func (t DocumentOnTypeFormattingRequest) isMessage() {}
func (t DocumentOnTypeFormattingRequest) isRequest() {}
func (t DocumentOnTypeFormattingRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentOnTypeFormattingRequest) GetParams() DocumentOnTypeFormattingParams { return t.Params }
func (t *DocumentOnTypeFormattingRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentOnTypeFormattingRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentOnTypeFormattingRequest(result)
   return nil
}
// A request to result a `CallHierarchyItem` in a document at a given position.
// Can be used as an input to an incoming or outgoing call hierarchy.
// 
// @since 3.16.0
type CallHierarchyPrepareRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CallHierarchyPrepareParams `json:"params"`
}
func (t CallHierarchyPrepareRequest) isMessage() {}
func (t CallHierarchyPrepareRequest) isRequest() {}
func (t CallHierarchyPrepareRequest) GetMethod() MethodKind { return t.Method }
func (t CallHierarchyPrepareRequest) GetParams() CallHierarchyPrepareParams { return t.Params }
func (t *CallHierarchyPrepareRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CallHierarchyPrepareRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CallHierarchyPrepareRequest(result)
   return nil
}
// A request to test and perform the setup necessary for a rename.
// 
// @since 3.16 - support for default behavior
type PrepareRenameRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params PrepareRenameParams `json:"params"`
}
func (t PrepareRenameRequest) isMessage() {}
func (t PrepareRenameRequest) isRequest() {}
func (t PrepareRenameRequest) GetMethod() MethodKind { return t.Method }
func (t PrepareRenameRequest) GetParams() PrepareRenameParams { return t.Params }
func (t *PrepareRenameRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias PrepareRenameRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = PrepareRenameRequest(result)
   return nil
}
// A request to result a `TypeHierarchyItem` in a document at a given position.
// Can be used as an input to a subtypes or supertypes type hierarchy.
// 
// @since 3.17.0
type TypeHierarchyPrepareRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params TypeHierarchyPrepareParams `json:"params"`
}
func (t TypeHierarchyPrepareRequest) isMessage() {}
func (t TypeHierarchyPrepareRequest) isRequest() {}
func (t TypeHierarchyPrepareRequest) GetMethod() MethodKind { return t.Method }
func (t TypeHierarchyPrepareRequest) GetParams() TypeHierarchyPrepareParams { return t.Params }
func (t *TypeHierarchyPrepareRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TypeHierarchyPrepareRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TypeHierarchyPrepareRequest(result)
   return nil
}
// A request to format a range in a document.
type DocumentRangeFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentRangeFormattingParams `json:"params"`
}
func (t DocumentRangeFormattingRequest) isMessage() {}
func (t DocumentRangeFormattingRequest) isRequest() {}
func (t DocumentRangeFormattingRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentRangeFormattingRequest) GetParams() DocumentRangeFormattingParams { return t.Params }
func (t *DocumentRangeFormattingRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentRangeFormattingRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentRangeFormattingRequest(result)
   return nil
}
// A request to format ranges in a document.
// 
// @since 3.18.0
// @proposed
type DocumentRangesFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DocumentRangesFormattingParams `json:"params"`
}
func (t DocumentRangesFormattingRequest) isMessage() {}
func (t DocumentRangesFormattingRequest) isRequest() {}
func (t DocumentRangesFormattingRequest) GetMethod() MethodKind { return t.Method }
func (t DocumentRangesFormattingRequest) GetParams() DocumentRangesFormattingParams { return t.Params }
func (t *DocumentRangesFormattingRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DocumentRangesFormattingRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DocumentRangesFormattingRequest(result)
   return nil
}
// A request to resolve project-wide references for the symbol denoted
// by the given text document position. The request's parameter is of
// type {@link ReferenceParams} the response is of type
// {@link Location Location[]} or a Thenable that resolves to such.
type ReferencesRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ReferenceParams `json:"params"`
}
func (t ReferencesRequest) isMessage() {}
func (t ReferencesRequest) isRequest() {}
func (t ReferencesRequest) GetMethod() MethodKind { return t.Method }
func (t ReferencesRequest) GetParams() ReferenceParams { return t.Params }
func (t *ReferencesRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ReferencesRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ReferencesRequest(result)
   return nil
}
// A request to rename a symbol.
type RenameRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params RenameParams `json:"params"`
}
func (t RenameRequest) isMessage() {}
func (t RenameRequest) isRequest() {}
func (t RenameRequest) GetMethod() MethodKind { return t.Method }
func (t RenameRequest) GetParams() RenameParams { return t.Params }
func (t *RenameRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias RenameRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = RenameRequest(result)
   return nil
}
// A request to provide selection ranges in a document. The request's
// parameter is of type {@link SelectionRangeParams}, the
// response is of type {@link SelectionRange SelectionRange[]} or a Thenable
// that resolves to such.
type SelectionRangeRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params SelectionRangeParams `json:"params"`
}
func (t SelectionRangeRequest) isMessage() {}
func (t SelectionRangeRequest) isRequest() {}
func (t SelectionRangeRequest) GetMethod() MethodKind { return t.Method }
func (t SelectionRangeRequest) GetParams() SelectionRangeParams { return t.Params }
func (t *SelectionRangeRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SelectionRangeRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SelectionRangeRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params SemanticTokensParams `json:"params"`
}
func (t SemanticTokensRequest) isMessage() {}
func (t SemanticTokensRequest) isRequest() {}
func (t SemanticTokensRequest) GetMethod() MethodKind { return t.Method }
func (t SemanticTokensRequest) GetParams() SemanticTokensParams { return t.Params }
func (t *SemanticTokensRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SemanticTokensRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SemanticTokensRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensDeltaRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params SemanticTokensDeltaParams `json:"params"`
}
func (t SemanticTokensDeltaRequest) isMessage() {}
func (t SemanticTokensDeltaRequest) isRequest() {}
func (t SemanticTokensDeltaRequest) GetMethod() MethodKind { return t.Method }
func (t SemanticTokensDeltaRequest) GetParams() SemanticTokensDeltaParams { return t.Params }
func (t *SemanticTokensDeltaRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SemanticTokensDeltaRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SemanticTokensDeltaRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensRangeRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params SemanticTokensRangeParams `json:"params"`
}
func (t SemanticTokensRangeRequest) isMessage() {}
func (t SemanticTokensRangeRequest) isRequest() {}
func (t SemanticTokensRangeRequest) GetMethod() MethodKind { return t.Method }
func (t SemanticTokensRangeRequest) GetParams() SemanticTokensRangeParams { return t.Params }
func (t *SemanticTokensRangeRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SemanticTokensRangeRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SemanticTokensRangeRequest(result)
   return nil
}

type SignatureHelpRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params SignatureHelpParams `json:"params"`
}
func (t SignatureHelpRequest) isMessage() {}
func (t SignatureHelpRequest) isRequest() {}
func (t SignatureHelpRequest) GetMethod() MethodKind { return t.Method }
func (t SignatureHelpRequest) GetParams() SignatureHelpParams { return t.Params }
func (t *SignatureHelpRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SignatureHelpRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SignatureHelpRequest(result)
   return nil
}
// A request to resolve the type definition locations of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPositionParams}
// the response is of type {@link Definition} or a Thenable that resolves to such.
type TypeDefinitionRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params TypeDefinitionParams `json:"params"`
}
func (t TypeDefinitionRequest) isMessage() {}
func (t TypeDefinitionRequest) isRequest() {}
func (t TypeDefinitionRequest) GetMethod() MethodKind { return t.Method }
func (t TypeDefinitionRequest) GetParams() TypeDefinitionParams { return t.Params }
func (t *TypeDefinitionRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TypeDefinitionRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TypeDefinitionRequest(result)
   return nil
}
// A document will save request is sent from the client to the server before
// the document is actually saved. The request can return an array of TextEdits
// which will be applied to the text document before it is saved. Please note that
// clients might drop results if computing the text edits took too long or if a
// server constantly fails on this request. This is done to keep the save fast and
// reliable.
type WillSaveTextDocumentWaitUntilRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params WillSaveTextDocumentParams `json:"params"`
}
func (t WillSaveTextDocumentWaitUntilRequest) isMessage() {}
func (t WillSaveTextDocumentWaitUntilRequest) isRequest() {}
func (t WillSaveTextDocumentWaitUntilRequest) GetMethod() MethodKind { return t.Method }
func (t WillSaveTextDocumentWaitUntilRequest) GetParams() WillSaveTextDocumentParams { return t.Params }
func (t *WillSaveTextDocumentWaitUntilRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WillSaveTextDocumentWaitUntilRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WillSaveTextDocumentWaitUntilRequest(result)
   return nil
}
// A request to resolve the subtypes for a given `TypeHierarchyItem`.
// 
// @since 3.17.0
type TypeHierarchySubtypesRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params TypeHierarchySubtypesParams `json:"params"`
}
func (t TypeHierarchySubtypesRequest) isMessage() {}
func (t TypeHierarchySubtypesRequest) isRequest() {}
func (t TypeHierarchySubtypesRequest) GetMethod() MethodKind { return t.Method }
func (t TypeHierarchySubtypesRequest) GetParams() TypeHierarchySubtypesParams { return t.Params }
func (t *TypeHierarchySubtypesRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TypeHierarchySubtypesRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TypeHierarchySubtypesRequest(result)
   return nil
}
// A request to resolve the supertypes for a given `TypeHierarchyItem`.
// 
// @since 3.17.0
type TypeHierarchySupertypesRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params TypeHierarchySupertypesParams `json:"params"`
}
func (t TypeHierarchySupertypesRequest) isMessage() {}
func (t TypeHierarchySupertypesRequest) isRequest() {}
func (t TypeHierarchySupertypesRequest) GetMethod() MethodKind { return t.Method }
func (t TypeHierarchySupertypesRequest) GetParams() TypeHierarchySupertypesParams { return t.Params }
func (t *TypeHierarchySupertypesRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TypeHierarchySupertypesRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TypeHierarchySupertypesRequest(result)
   return nil
}
// A request to show a document. This request might open an
// external program depending on the value of the URI to open.
// For example a request to open `https://code.visualstudio.com/`
// will very likely open the URI in a WEB browser.
// 
// @since 3.16.0
type ShowDocumentRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ShowDocumentParams `json:"params"`
}
func (t ShowDocumentRequest) isMessage() {}
func (t ShowDocumentRequest) isRequest() {}
func (t ShowDocumentRequest) GetMethod() MethodKind { return t.Method }
func (t ShowDocumentRequest) GetParams() ShowDocumentParams { return t.Params }
func (t *ShowDocumentRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ShowDocumentRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ShowDocumentRequest(result)
   return nil
}
// The show message request is sent from the server to the client to show a message
// and a set of options actions to the user.
type ShowMessageRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ShowMessageRequestParams `json:"params"`
}
func (t ShowMessageRequest) isMessage() {}
func (t ShowMessageRequest) isRequest() {}
func (t ShowMessageRequest) GetMethod() MethodKind { return t.Method }
func (t ShowMessageRequest) GetParams() ShowMessageRequestParams { return t.Params }
func (t *ShowMessageRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ShowMessageRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ShowMessageRequest(result)
   return nil
}
// The `window/workDoneProgress/create` request is sent from the server to the client to initiate progress
// reporting from the server.
type WorkDoneProgressCreateRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params WorkDoneProgressCreateParams `json:"params"`
}
func (t WorkDoneProgressCreateRequest) isMessage() {}
func (t WorkDoneProgressCreateRequest) isRequest() {}
func (t WorkDoneProgressCreateRequest) GetMethod() MethodKind { return t.Method }
func (t WorkDoneProgressCreateRequest) GetParams() WorkDoneProgressCreateParams { return t.Params }
func (t *WorkDoneProgressCreateRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WorkDoneProgressCreateRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WorkDoneProgressCreateRequest(result)
   return nil
}
// A request sent from the server to the client to modified certain resources.
type ApplyWorkspaceEditRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ApplyWorkspaceEditParams `json:"params"`
}
func (t ApplyWorkspaceEditRequest) isMessage() {}
func (t ApplyWorkspaceEditRequest) isRequest() {}
func (t ApplyWorkspaceEditRequest) GetMethod() MethodKind { return t.Method }
func (t ApplyWorkspaceEditRequest) GetParams() ApplyWorkspaceEditParams { return t.Params }
func (t *ApplyWorkspaceEditRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ApplyWorkspaceEditRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ApplyWorkspaceEditRequest(result)
   return nil
}
// A request to refresh all code actions
// 
// @since 3.16.0
type CodeLensRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t CodeLensRefreshRequest) isMessage() {}
func (t CodeLensRefreshRequest) isRequest() {}
func (t CodeLensRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t CodeLensRefreshRequest) GetParams() any { return t.Params }
func (t *CodeLensRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CodeLensRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CodeLensRefreshRequest(result)
   return nil
}
// The 'workspace/configuration' request is sent from the server to the client to fetch a certain
// configuration setting.
// 
// This pull model replaces the old push model were the client signaled configuration change via an
// event. If the server still needs to react to configuration changes (since the server caches the
// result of `workspace/configuration` requests) the server should register for an empty configuration
// change event and empty the cache if such an event is received.
type ConfigurationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ConfigurationParams `json:"params"`
}
func (t ConfigurationRequest) isMessage() {}
func (t ConfigurationRequest) isRequest() {}
func (t ConfigurationRequest) GetMethod() MethodKind { return t.Method }
func (t ConfigurationRequest) GetParams() ConfigurationParams { return t.Params }
func (t *ConfigurationRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ConfigurationRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ConfigurationRequest(result)
   return nil
}
// The workspace diagnostic request definition.
// 
// @since 3.17.0
type WorkspaceDiagnosticRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params WorkspaceDiagnosticParams `json:"params"`
}
func (t WorkspaceDiagnosticRequest) isMessage() {}
func (t WorkspaceDiagnosticRequest) isRequest() {}
func (t WorkspaceDiagnosticRequest) GetMethod() MethodKind { return t.Method }
func (t WorkspaceDiagnosticRequest) GetParams() WorkspaceDiagnosticParams { return t.Params }
func (t *WorkspaceDiagnosticRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WorkspaceDiagnosticRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WorkspaceDiagnosticRequest(result)
   return nil
}
// The diagnostic refresh request definition.
// 
// @since 3.17.0
type DiagnosticRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t DiagnosticRefreshRequest) isMessage() {}
func (t DiagnosticRefreshRequest) isRequest() {}
func (t DiagnosticRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t DiagnosticRefreshRequest) GetParams() any { return t.Params }
func (t *DiagnosticRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DiagnosticRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DiagnosticRefreshRequest(result)
   return nil
}
// A request send from the client to the server to execute a command. The request might return
// a workspace edit which the client will apply to the workspace.
type ExecuteCommandRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params ExecuteCommandParams `json:"params"`
}
func (t ExecuteCommandRequest) isMessage() {}
func (t ExecuteCommandRequest) isRequest() {}
func (t ExecuteCommandRequest) GetMethod() MethodKind { return t.Method }
func (t ExecuteCommandRequest) GetParams() ExecuteCommandParams { return t.Params }
func (t *ExecuteCommandRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ExecuteCommandRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ExecuteCommandRequest(result)
   return nil
}
// @since 3.18.0
// @proposed
type FoldingRangeRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t FoldingRangeRefreshRequest) isMessage() {}
func (t FoldingRangeRefreshRequest) isRequest() {}
func (t FoldingRangeRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t FoldingRangeRefreshRequest) GetParams() any { return t.Params }
func (t *FoldingRangeRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias FoldingRangeRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = FoldingRangeRefreshRequest(result)
   return nil
}
// @since 3.17.0
type InlayHintRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t InlayHintRefreshRequest) isMessage() {}
func (t InlayHintRefreshRequest) isRequest() {}
func (t InlayHintRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t InlayHintRefreshRequest) GetParams() any { return t.Params }
func (t *InlayHintRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InlayHintRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InlayHintRefreshRequest(result)
   return nil
}
// @since 3.17.0
type InlineValueRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t InlineValueRefreshRequest) isMessage() {}
func (t InlineValueRefreshRequest) isRequest() {}
func (t InlineValueRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t InlineValueRefreshRequest) GetParams() any { return t.Params }
func (t *InlineValueRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InlineValueRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InlineValueRefreshRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t SemanticTokensRefreshRequest) isMessage() {}
func (t SemanticTokensRefreshRequest) isRequest() {}
func (t SemanticTokensRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t SemanticTokensRefreshRequest) GetParams() any { return t.Params }
func (t *SemanticTokensRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SemanticTokensRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SemanticTokensRefreshRequest(result)
   return nil
}
// A request to list project-wide symbols matching the query string given
// by the {@link WorkspaceSymbolParams}. The response is
// of type {@link SymbolInformation SymbolInformation[]} or a Thenable that
// resolves to such.
// 
// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients
//  need to advertise support for WorkspaceSymbols via the client capability
//  `workspace.symbol.resolveSupport`.
type WorkspaceSymbolRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params WorkspaceSymbolParams `json:"params"`
}
func (t WorkspaceSymbolRequest) isMessage() {}
func (t WorkspaceSymbolRequest) isRequest() {}
func (t WorkspaceSymbolRequest) GetMethod() MethodKind { return t.Method }
func (t WorkspaceSymbolRequest) GetParams() WorkspaceSymbolParams { return t.Params }
func (t *WorkspaceSymbolRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WorkspaceSymbolRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WorkspaceSymbolRequest(result)
   return nil
}
// The `workspace/textDocumentContent` request is sent from the client to the
// server to request the content of a text document.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params TextDocumentContentParams `json:"params"`
}
func (t TextDocumentContentRequest) isMessage() {}
func (t TextDocumentContentRequest) isRequest() {}
func (t TextDocumentContentRequest) GetMethod() MethodKind { return t.Method }
func (t TextDocumentContentRequest) GetParams() TextDocumentContentParams { return t.Params }
func (t *TextDocumentContentRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TextDocumentContentRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TextDocumentContentRequest(result)
   return nil
}
// The `workspace/textDocumentContent` request is sent from the server to the client to refresh
// the content of a specific text document.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params TextDocumentContentRefreshParams `json:"params"`
}
func (t TextDocumentContentRefreshRequest) isMessage() {}
func (t TextDocumentContentRefreshRequest) isRequest() {}
func (t TextDocumentContentRefreshRequest) GetMethod() MethodKind { return t.Method }
func (t TextDocumentContentRefreshRequest) GetParams() TextDocumentContentRefreshParams { return t.Params }
func (t *TextDocumentContentRefreshRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TextDocumentContentRefreshRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TextDocumentContentRefreshRequest(result)
   return nil
}
// The will create files request is sent from the client to the server before files are actually
// created as long as the creation is triggered from within the client.
// 
// The request can return a `WorkspaceEdit` which will be applied to workspace before the
// files are created. Hence the `WorkspaceEdit` can not manipulate the content of the file
// to be created.
// 
// @since 3.16.0
type WillCreateFilesRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params CreateFilesParams `json:"params"`
}
func (t WillCreateFilesRequest) isMessage() {}
func (t WillCreateFilesRequest) isRequest() {}
func (t WillCreateFilesRequest) GetMethod() MethodKind { return t.Method }
func (t WillCreateFilesRequest) GetParams() CreateFilesParams { return t.Params }
func (t *WillCreateFilesRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WillCreateFilesRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WillCreateFilesRequest(result)
   return nil
}
// The did delete files notification is sent from the client to the server when
// files were deleted from within the client.
// 
// @since 3.16.0
type WillDeleteFilesRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params DeleteFilesParams `json:"params"`
}
func (t WillDeleteFilesRequest) isMessage() {}
func (t WillDeleteFilesRequest) isRequest() {}
func (t WillDeleteFilesRequest) GetMethod() MethodKind { return t.Method }
func (t WillDeleteFilesRequest) GetParams() DeleteFilesParams { return t.Params }
func (t *WillDeleteFilesRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WillDeleteFilesRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WillDeleteFilesRequest(result)
   return nil
}
// The will rename files request is sent from the client to the server before files are actually
// renamed as long as the rename is triggered from within the client.
// 
// @since 3.16.0
type WillRenameFilesRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params RenameFilesParams `json:"params"`
}
func (t WillRenameFilesRequest) isMessage() {}
func (t WillRenameFilesRequest) isRequest() {}
func (t WillRenameFilesRequest) GetMethod() MethodKind { return t.Method }
func (t WillRenameFilesRequest) GetParams() RenameFilesParams { return t.Params }
func (t *WillRenameFilesRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WillRenameFilesRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WillRenameFilesRequest(result)
   return nil
}
// The `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
type WorkspaceFoldersRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t WorkspaceFoldersRequest) isMessage() {}
func (t WorkspaceFoldersRequest) isRequest() {}
func (t WorkspaceFoldersRequest) GetMethod() MethodKind { return t.Method }
func (t WorkspaceFoldersRequest) GetParams() any { return t.Params }
func (t *WorkspaceFoldersRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WorkspaceFoldersRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WorkspaceFoldersRequest(result)
   return nil
}
// A request to resolve the range inside the workspace
// symbol's location.
// 
// @since 3.17.0
type WorkspaceSymbolResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or2[string, int32] `json:"id"`
	Method MethodKind `json:"method"`
	Params WorkspaceSymbol `json:"params"`
}
func (t WorkspaceSymbolResolveRequest) isMessage() {}
func (t WorkspaceSymbolResolveRequest) isRequest() {}
func (t WorkspaceSymbolResolveRequest) GetMethod() MethodKind { return t.Method }
func (t WorkspaceSymbolResolveRequest) GetParams() WorkspaceSymbol { return t.Params }
func (t *WorkspaceSymbolResolveRequest) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["id"]; !exists {
       return fmt.Errorf("Missing required request field: id")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WorkspaceSymbolResolveRequest
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WorkspaceSymbolResolveRequest(result)
   return nil
}
type CallHierarchyIncomingCallsResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]CallHierarchyIncomingCall `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CallHierarchyIncomingCallsResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CallHierarchyIncomingCallsResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CallHierarchyIncomingCallsResponse(temp)
   return nil
}
func (t CallHierarchyIncomingCallsResponse) isMessage() {}
func (t CallHierarchyIncomingCallsResponse) isResponse() {}
type CallHierarchyOutgoingCallsResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]CallHierarchyOutgoingCall `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CallHierarchyOutgoingCallsResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CallHierarchyOutgoingCallsResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CallHierarchyOutgoingCallsResponse(temp)
   return nil
}
func (t CallHierarchyOutgoingCallsResponse) isMessage() {}
func (t CallHierarchyOutgoingCallsResponse) isResponse() {}
type RegistrationResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *RegistrationResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias RegistrationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = RegistrationResponse(temp)
   return nil
}
func (t RegistrationResponse) isMessage() {}
func (t RegistrationResponse) isResponse() {}
type UnregistrationResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *UnregistrationResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias UnregistrationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = UnregistrationResponse(temp)
   return nil
}
func (t UnregistrationResponse) isMessage() {}
func (t UnregistrationResponse) isResponse() {}
type CodeActionResolveResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result CodeAction `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CodeActionResolveResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CodeActionResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeActionResolveResponse(temp)
   return nil
}
func (t CodeActionResolveResponse) isMessage() {}
func (t CodeActionResolveResponse) isResponse() {}
type CodeLensResolveResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result CodeLens `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CodeLensResolveResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CodeLensResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeLensResolveResponse(temp)
   return nil
}
func (t CodeLensResolveResponse) isMessage() {}
func (t CodeLensResolveResponse) isResponse() {}
type CompletionResolveResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result CompletionItem `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CompletionResolveResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CompletionResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CompletionResolveResponse(temp)
   return nil
}
func (t CompletionResolveResponse) isMessage() {}
func (t CompletionResolveResponse) isResponse() {}
type DocumentLinkResolveResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result DocumentLink `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentLinkResolveResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentLinkResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentLinkResolveResponse(temp)
   return nil
}
func (t DocumentLinkResolveResponse) isMessage() {}
func (t DocumentLinkResolveResponse) isResponse() {}
type InitializeResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result InitializeResult `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *InitializeResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InitializeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InitializeResponse(temp)
   return nil
}
func (t InitializeResponse) isMessage() {}
func (t InitializeResponse) isResponse() {}
type InlayHintResolveResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result InlayHint `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *InlayHintResolveResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InlayHintResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlayHintResolveResponse(temp)
   return nil
}
func (t InlayHintResolveResponse) isMessage() {}
func (t InlayHintResolveResponse) isResponse() {}
type ShutdownResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *ShutdownResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ShutdownResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ShutdownResponse(temp)
   return nil
}
func (t ShutdownResponse) isMessage() {}
func (t ShutdownResponse) isResponse() {}
type CodeActionResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]Or2[Command, CodeAction] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CodeActionResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CodeActionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeActionResponse(temp)
   return nil
}
func (t CodeActionResponse) isMessage() {}
func (t CodeActionResponse) isResponse() {}
type CodeLensResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]CodeLens `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CodeLensResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CodeLensResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeLensResponse(temp)
   return nil
}
func (t CodeLensResponse) isMessage() {}
func (t CodeLensResponse) isResponse() {}
type ColorPresentationResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result []ColorPresentation `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ColorPresentationResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ColorPresentationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ColorPresentationResponse(temp)
   return nil
}
func (t ColorPresentationResponse) isMessage() {}
func (t ColorPresentationResponse) isResponse() {}
type CompletionResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[[]CompletionItem, CompletionList] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CompletionResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CompletionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CompletionResponse(temp)
   return nil
}
func (t CompletionResponse) isMessage() {}
func (t CompletionResponse) isResponse() {}
type DeclarationResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[Declaration, []DeclarationLink] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DeclarationResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DeclarationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DeclarationResponse(temp)
   return nil
}
func (t DeclarationResponse) isMessage() {}
func (t DeclarationResponse) isResponse() {}
type DefinitionResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[Definition, []DefinitionLink] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DefinitionResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DefinitionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DefinitionResponse(temp)
   return nil
}
func (t DefinitionResponse) isMessage() {}
func (t DefinitionResponse) isResponse() {}
type DocumentDiagnosticResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result DocumentDiagnosticReport `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentDiagnosticResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentDiagnosticResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentDiagnosticResponse(temp)
   return nil
}
func (t DocumentDiagnosticResponse) isMessage() {}
func (t DocumentDiagnosticResponse) isResponse() {}
type DocumentColorResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result []ColorInformation `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentColorResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentColorResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentColorResponse(temp)
   return nil
}
func (t DocumentColorResponse) isMessage() {}
func (t DocumentColorResponse) isResponse() {}
type DocumentHighlightResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]DocumentHighlight `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentHighlightResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentHighlightResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentHighlightResponse(temp)
   return nil
}
func (t DocumentHighlightResponse) isMessage() {}
func (t DocumentHighlightResponse) isResponse() {}
type DocumentLinkResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]DocumentLink `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentLinkResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentLinkResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentLinkResponse(temp)
   return nil
}
func (t DocumentLinkResponse) isMessage() {}
func (t DocumentLinkResponse) isResponse() {}
type DocumentSymbolResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[[]SymbolInformation, []DocumentSymbol] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentSymbolResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentSymbolResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentSymbolResponse(temp)
   return nil
}
func (t DocumentSymbolResponse) isMessage() {}
func (t DocumentSymbolResponse) isResponse() {}
type FoldingRangeResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]FoldingRange `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *FoldingRangeResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias FoldingRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = FoldingRangeResponse(temp)
   return nil
}
func (t FoldingRangeResponse) isMessage() {}
func (t FoldingRangeResponse) isResponse() {}
type DocumentFormattingResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TextEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentFormattingResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentFormattingResponse(temp)
   return nil
}
func (t DocumentFormattingResponse) isMessage() {}
func (t DocumentFormattingResponse) isResponse() {}
type HoverResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *Hover `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *HoverResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias HoverResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = HoverResponse(temp)
   return nil
}
func (t HoverResponse) isMessage() {}
func (t HoverResponse) isResponse() {}
type ImplementationResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[Definition, []DefinitionLink] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ImplementationResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ImplementationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ImplementationResponse(temp)
   return nil
}
func (t ImplementationResponse) isMessage() {}
func (t ImplementationResponse) isResponse() {}
type InlayHintResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]InlayHint `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *InlayHintResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InlayHintResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlayHintResponse(temp)
   return nil
}
func (t InlayHintResponse) isMessage() {}
func (t InlayHintResponse) isResponse() {}
type InlineCompletionResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[InlineCompletionList, []InlineCompletionItem] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *InlineCompletionResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InlineCompletionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlineCompletionResponse(temp)
   return nil
}
func (t InlineCompletionResponse) isMessage() {}
func (t InlineCompletionResponse) isResponse() {}
type InlineValueResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]InlineValue `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *InlineValueResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InlineValueResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlineValueResponse(temp)
   return nil
}
func (t InlineValueResponse) isMessage() {}
func (t InlineValueResponse) isResponse() {}
type LinkedEditingRangeResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *LinkedEditingRanges `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *LinkedEditingRangeResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias LinkedEditingRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = LinkedEditingRangeResponse(temp)
   return nil
}
func (t LinkedEditingRangeResponse) isMessage() {}
func (t LinkedEditingRangeResponse) isResponse() {}
type MonikerResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]Moniker `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *MonikerResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias MonikerResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = MonikerResponse(temp)
   return nil
}
func (t MonikerResponse) isMessage() {}
func (t MonikerResponse) isResponse() {}
type DocumentOnTypeFormattingResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TextEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentOnTypeFormattingResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentOnTypeFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentOnTypeFormattingResponse(temp)
   return nil
}
func (t DocumentOnTypeFormattingResponse) isMessage() {}
func (t DocumentOnTypeFormattingResponse) isResponse() {}
type CallHierarchyPrepareResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]CallHierarchyItem `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *CallHierarchyPrepareResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CallHierarchyPrepareResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CallHierarchyPrepareResponse(temp)
   return nil
}
func (t CallHierarchyPrepareResponse) isMessage() {}
func (t CallHierarchyPrepareResponse) isResponse() {}
type PrepareRenameResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *PrepareRenameResult `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *PrepareRenameResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias PrepareRenameResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = PrepareRenameResponse(temp)
   return nil
}
func (t PrepareRenameResponse) isMessage() {}
func (t PrepareRenameResponse) isResponse() {}
type TypeHierarchyPrepareResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TypeHierarchyItem `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *TypeHierarchyPrepareResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias TypeHierarchyPrepareResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeHierarchyPrepareResponse(temp)
   return nil
}
func (t TypeHierarchyPrepareResponse) isMessage() {}
func (t TypeHierarchyPrepareResponse) isResponse() {}
type DocumentRangeFormattingResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TextEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentRangeFormattingResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentRangeFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentRangeFormattingResponse(temp)
   return nil
}
func (t DocumentRangeFormattingResponse) isMessage() {}
func (t DocumentRangeFormattingResponse) isResponse() {}
type DocumentRangesFormattingResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TextEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *DocumentRangesFormattingResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DocumentRangesFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentRangesFormattingResponse(temp)
   return nil
}
func (t DocumentRangesFormattingResponse) isMessage() {}
func (t DocumentRangesFormattingResponse) isResponse() {}
type ReferencesResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]Location `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ReferencesResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ReferencesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ReferencesResponse(temp)
   return nil
}
func (t ReferencesResponse) isMessage() {}
func (t ReferencesResponse) isResponse() {}
type RenameResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *WorkspaceEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *RenameResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias RenameResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = RenameResponse(temp)
   return nil
}
func (t RenameResponse) isMessage() {}
func (t RenameResponse) isResponse() {}
type SelectionRangeResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]SelectionRange `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *SelectionRangeResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias SelectionRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SelectionRangeResponse(temp)
   return nil
}
func (t SelectionRangeResponse) isMessage() {}
func (t SelectionRangeResponse) isResponse() {}
type SemanticTokensResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *SemanticTokens `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *SemanticTokensResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias SemanticTokensResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensResponse(temp)
   return nil
}
func (t SemanticTokensResponse) isMessage() {}
func (t SemanticTokensResponse) isResponse() {}
type SemanticTokensDeltaResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[SemanticTokens, SemanticTokensDelta] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *SemanticTokensDeltaResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias SemanticTokensDeltaResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensDeltaResponse(temp)
   return nil
}
func (t SemanticTokensDeltaResponse) isMessage() {}
func (t SemanticTokensDeltaResponse) isResponse() {}
type SemanticTokensRangeResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *SemanticTokens `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *SemanticTokensRangeResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias SemanticTokensRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensRangeResponse(temp)
   return nil
}
func (t SemanticTokensRangeResponse) isMessage() {}
func (t SemanticTokensRangeResponse) isResponse() {}
type SignatureHelpResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *SignatureHelp `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *SignatureHelpResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias SignatureHelpResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SignatureHelpResponse(temp)
   return nil
}
func (t SignatureHelpResponse) isMessage() {}
func (t SignatureHelpResponse) isResponse() {}
type TypeDefinitionResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[Definition, []DefinitionLink] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *TypeDefinitionResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias TypeDefinitionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeDefinitionResponse(temp)
   return nil
}
func (t TypeDefinitionResponse) isMessage() {}
func (t TypeDefinitionResponse) isResponse() {}
type WillSaveTextDocumentWaitUntilResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TextEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WillSaveTextDocumentWaitUntilResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WillSaveTextDocumentWaitUntilResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillSaveTextDocumentWaitUntilResponse(temp)
   return nil
}
func (t WillSaveTextDocumentWaitUntilResponse) isMessage() {}
func (t WillSaveTextDocumentWaitUntilResponse) isResponse() {}
type TypeHierarchySubtypesResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TypeHierarchyItem `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *TypeHierarchySubtypesResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias TypeHierarchySubtypesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeHierarchySubtypesResponse(temp)
   return nil
}
func (t TypeHierarchySubtypesResponse) isMessage() {}
func (t TypeHierarchySubtypesResponse) isResponse() {}
type TypeHierarchySupertypesResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]TypeHierarchyItem `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *TypeHierarchySupertypesResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias TypeHierarchySupertypesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeHierarchySupertypesResponse(temp)
   return nil
}
func (t TypeHierarchySupertypesResponse) isMessage() {}
func (t TypeHierarchySupertypesResponse) isResponse() {}
type ShowDocumentResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result ShowDocumentResult `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ShowDocumentResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ShowDocumentResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ShowDocumentResponse(temp)
   return nil
}
func (t ShowDocumentResponse) isMessage() {}
func (t ShowDocumentResponse) isResponse() {}
type ShowMessageResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *MessageActionItem `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ShowMessageResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ShowMessageResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ShowMessageResponse(temp)
   return nil
}
func (t ShowMessageResponse) isMessage() {}
func (t ShowMessageResponse) isResponse() {}
type WorkDoneProgressCreateResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *WorkDoneProgressCreateResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WorkDoneProgressCreateResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkDoneProgressCreateResponse(temp)
   return nil
}
func (t WorkDoneProgressCreateResponse) isMessage() {}
func (t WorkDoneProgressCreateResponse) isResponse() {}
type ApplyWorkspaceEditResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result ApplyWorkspaceEditResult `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ApplyWorkspaceEditResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ApplyWorkspaceEditResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ApplyWorkspaceEditResponse(temp)
   return nil
}
func (t ApplyWorkspaceEditResponse) isMessage() {}
func (t ApplyWorkspaceEditResponse) isResponse() {}
type CodeLensRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *CodeLensRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias CodeLensRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeLensRefreshResponse(temp)
   return nil
}
func (t CodeLensRefreshResponse) isMessage() {}
func (t CodeLensRefreshResponse) isResponse() {}
type ConfigurationResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result []any `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ConfigurationResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ConfigurationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ConfigurationResponse(temp)
   return nil
}
func (t ConfigurationResponse) isMessage() {}
func (t ConfigurationResponse) isResponse() {}
type WorkspaceDiagnosticResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result WorkspaceDiagnosticReport `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WorkspaceDiagnosticResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WorkspaceDiagnosticResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceDiagnosticResponse(temp)
   return nil
}
func (t WorkspaceDiagnosticResponse) isMessage() {}
func (t WorkspaceDiagnosticResponse) isResponse() {}
type DiagnosticRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *DiagnosticRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias DiagnosticRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DiagnosticRefreshResponse(temp)
   return nil
}
func (t DiagnosticRefreshResponse) isMessage() {}
func (t DiagnosticRefreshResponse) isResponse() {}
type ExecuteCommandResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *any `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *ExecuteCommandResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias ExecuteCommandResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ExecuteCommandResponse(temp)
   return nil
}
func (t ExecuteCommandResponse) isMessage() {}
func (t ExecuteCommandResponse) isResponse() {}
type FoldingRangeRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *FoldingRangeRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias FoldingRangeRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = FoldingRangeRefreshResponse(temp)
   return nil
}
func (t FoldingRangeRefreshResponse) isMessage() {}
func (t FoldingRangeRefreshResponse) isResponse() {}
type InlayHintRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *InlayHintRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InlayHintRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlayHintRefreshResponse(temp)
   return nil
}
func (t InlayHintRefreshResponse) isMessage() {}
func (t InlayHintRefreshResponse) isResponse() {}
type InlineValueRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *InlineValueRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias InlineValueRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlineValueRefreshResponse(temp)
   return nil
}
func (t InlineValueRefreshResponse) isMessage() {}
func (t InlineValueRefreshResponse) isResponse() {}
type SemanticTokensRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *SemanticTokensRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias SemanticTokensRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensRefreshResponse(temp)
   return nil
}
func (t SemanticTokensRefreshResponse) isMessage() {}
func (t SemanticTokensRefreshResponse) isResponse() {}
type WorkspaceSymbolResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result NullableOr2[[]SymbolInformation, []WorkspaceSymbol] `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WorkspaceSymbolResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WorkspaceSymbolResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceSymbolResponse(temp)
   return nil
}
func (t WorkspaceSymbolResponse) isMessage() {}
func (t WorkspaceSymbolResponse) isResponse() {}
type TextDocumentContentResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result TextDocumentContentResult `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *TextDocumentContentResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias TextDocumentContentResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TextDocumentContentResponse(temp)
   return nil
}
func (t TextDocumentContentResponse) isMessage() {}
func (t TextDocumentContentResponse) isResponse() {}
type TextDocumentContentRefreshResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`

	Error *ResponseError `json:"error,omitzero"`
}
func (t *TextDocumentContentRefreshResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias TextDocumentContentRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TextDocumentContentRefreshResponse(temp)
   return nil
}
func (t TextDocumentContentRefreshResponse) isMessage() {}
func (t TextDocumentContentRefreshResponse) isResponse() {}
type WillCreateFilesResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *WorkspaceEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WillCreateFilesResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WillCreateFilesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillCreateFilesResponse(temp)
   return nil
}
func (t WillCreateFilesResponse) isMessage() {}
func (t WillCreateFilesResponse) isResponse() {}
type WillDeleteFilesResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *WorkspaceEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WillDeleteFilesResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WillDeleteFilesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillDeleteFilesResponse(temp)
   return nil
}
func (t WillDeleteFilesResponse) isMessage() {}
func (t WillDeleteFilesResponse) isResponse() {}
type WillRenameFilesResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *WorkspaceEdit `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WillRenameFilesResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WillRenameFilesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillRenameFilesResponse(temp)
   return nil
}
func (t WillRenameFilesResponse) isMessage() {}
func (t WillRenameFilesResponse) isResponse() {}
type WorkspaceFoldersResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result *[]WorkspaceFolder `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WorkspaceFoldersResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WorkspaceFoldersResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceFoldersResponse(temp)
   return nil
}
func (t WorkspaceFoldersResponse) isMessage() {}
func (t WorkspaceFoldersResponse) isResponse() {}
type WorkspaceSymbolResolveResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id Or2[int32, string] `json:"id"`
	Result WorkspaceSymbol `json:"result,omitzero"`
	Error *ResponseError `json:"error,omitzero"`
}
func (t *WorkspaceSymbolResolveResponse) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   _, idExists := m["id"]
   _, jsonrpcExists := m["jsonrpc"]
   if !idExists || !jsonrpcExists {
       return fmt.Errorf("response must have an id and jsonrpc field")
   }
  type Alias WorkspaceSymbolResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceSymbolResolveResponse(temp)
   return nil
}
func (t WorkspaceSymbolResolveResponse) isMessage() {}
func (t WorkspaceSymbolResolveResponse) isResponse() {}
const (
	UnknownNotificationMethod MethodKind = ""
	OptionalCancelRequestMethod MethodKind = "$/cancelRequest"
	OptionalLogTraceMethod MethodKind = "$/logTrace"
	OptionalProgressMethod MethodKind = "$/progress"
	OptionalSetTraceMethod MethodKind = "$/setTrace"
	ExitMethod MethodKind = "exit"
	InitializedMethod MethodKind = "initialized"
	NotebookDocumentDidChangeMethod MethodKind = "notebookDocument/didChange"
	NotebookDocumentDidCloseMethod MethodKind = "notebookDocument/didClose"
	NotebookDocumentDidOpenMethod MethodKind = "notebookDocument/didOpen"
	NotebookDocumentDidSaveMethod MethodKind = "notebookDocument/didSave"
	TelemetryEventMethod MethodKind = "telemetry/event"
	TextDocumentDidChangeMethod MethodKind = "textDocument/didChange"
	TextDocumentDidCloseMethod MethodKind = "textDocument/didClose"
	TextDocumentDidOpenMethod MethodKind = "textDocument/didOpen"
	TextDocumentDidSaveMethod MethodKind = "textDocument/didSave"
	TextDocumentPublishDiagnosticsMethod MethodKind = "textDocument/publishDiagnostics"
	TextDocumentWillSaveMethod MethodKind = "textDocument/willSave"
	WindowLogMessageMethod MethodKind = "window/logMessage"
	WindowShowMessageMethod MethodKind = "window/showMessage"
	WindowWorkDoneProgressCancelMethod MethodKind = "window/workDoneProgress/cancel"
	WorkspaceDidChangeConfigurationMethod MethodKind = "workspace/didChangeConfiguration"
	WorkspaceDidChangeWatchedFilesMethod MethodKind = "workspace/didChangeWatchedFiles"
	WorkspaceDidChangeWorkspaceFoldersMethod MethodKind = "workspace/didChangeWorkspaceFolders"
	WorkspaceDidCreateFilesMethod MethodKind = "workspace/didCreateFiles"
	WorkspaceDidDeleteFilesMethod MethodKind = "workspace/didDeleteFiles"
	WorkspaceDidRenameFilesMethod MethodKind = "workspace/didRenameFiles"
)

var NotificationMethodMap = map[string]MethodKind{
	"$/cancelRequest": OptionalCancelRequestMethod,
	"$/logTrace": OptionalLogTraceMethod,
	"$/progress": OptionalProgressMethod,
	"$/setTrace": OptionalSetTraceMethod,
	"exit": ExitMethod,
	"initialized": InitializedMethod,
	"notebookDocument/didChange": NotebookDocumentDidChangeMethod,
	"notebookDocument/didClose": NotebookDocumentDidCloseMethod,
	"notebookDocument/didOpen": NotebookDocumentDidOpenMethod,
	"notebookDocument/didSave": NotebookDocumentDidSaveMethod,
	"telemetry/event": TelemetryEventMethod,
	"textDocument/didChange": TextDocumentDidChangeMethod,
	"textDocument/didClose": TextDocumentDidCloseMethod,
	"textDocument/didOpen": TextDocumentDidOpenMethod,
	"textDocument/didSave": TextDocumentDidSaveMethod,
	"textDocument/publishDiagnostics": TextDocumentPublishDiagnosticsMethod,
	"textDocument/willSave": TextDocumentWillSaveMethod,
	"window/logMessage": WindowLogMessageMethod,
	"window/showMessage": WindowShowMessageMethod,
	"window/workDoneProgress/cancel": WindowWorkDoneProgressCancelMethod,
	"workspace/didChangeConfiguration": WorkspaceDidChangeConfigurationMethod,
	"workspace/didChangeWatchedFiles": WorkspaceDidChangeWatchedFilesMethod,
	"workspace/didChangeWorkspaceFolders": WorkspaceDidChangeWorkspaceFoldersMethod,
	"workspace/didCreateFiles": WorkspaceDidCreateFilesMethod,
	"workspace/didDeleteFiles": WorkspaceDidDeleteFilesMethod,
	"workspace/didRenameFiles": WorkspaceDidRenameFilesMethod,
}


type CancelNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params CancelParams `json:"params"`
}
func (t CancelNotification) isMessage() {}
func (t CancelNotification) isNotification() {}
func (t CancelNotification) GetMethod() MethodKind { return t.Method }
func (t CancelNotification) GetParams() CancelParams { return t.Params }
func (t *CancelNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias CancelNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = CancelNotification(result)
   return nil
}

type LogTraceNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params LogTraceParams `json:"params"`
}
func (t LogTraceNotification) isMessage() {}
func (t LogTraceNotification) isNotification() {}
func (t LogTraceNotification) GetMethod() MethodKind { return t.Method }
func (t LogTraceNotification) GetParams() LogTraceParams { return t.Params }
func (t *LogTraceNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias LogTraceNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = LogTraceNotification(result)
   return nil
}

type ProgressNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params ProgressParams `json:"params"`
}
func (t ProgressNotification) isMessage() {}
func (t ProgressNotification) isNotification() {}
func (t ProgressNotification) GetMethod() MethodKind { return t.Method }
func (t ProgressNotification) GetParams() ProgressParams { return t.Params }
func (t *ProgressNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ProgressNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ProgressNotification(result)
   return nil
}

type SetTraceNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params SetTraceParams `json:"params"`
}
func (t SetTraceNotification) isMessage() {}
func (t SetTraceNotification) isNotification() {}
func (t SetTraceNotification) GetMethod() MethodKind { return t.Method }
func (t SetTraceNotification) GetParams() SetTraceParams { return t.Params }
func (t *SetTraceNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias SetTraceNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = SetTraceNotification(result)
   return nil
}
// The exit event is sent from the client to the server to
// ask the server to exit its process.
type ExitNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t ExitNotification) isMessage() {}
func (t ExitNotification) isNotification() {}
func (t ExitNotification) GetMethod() MethodKind { return t.Method }
func (t ExitNotification) GetParams() any { return t.Params }
func (t *ExitNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ExitNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ExitNotification(result)
   return nil
}
// The initialized notification is sent from the client to the
// server after the client is fully initialized and the server
// is allowed to send requests from the server to the client.
type InitializedNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params InitializedParams `json:"params"`
}
func (t InitializedNotification) isMessage() {}
func (t InitializedNotification) isNotification() {}
func (t InitializedNotification) GetMethod() MethodKind { return t.Method }
func (t InitializedNotification) GetParams() InitializedParams { return t.Params }
func (t *InitializedNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias InitializedNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = InitializedNotification(result)
   return nil
}

type DidChangeNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidChangeNotebookDocumentParams `json:"params"`
}
func (t DidChangeNotebookDocumentNotification) isMessage() {}
func (t DidChangeNotebookDocumentNotification) isNotification() {}
func (t DidChangeNotebookDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidChangeNotebookDocumentNotification) GetParams() DidChangeNotebookDocumentParams { return t.Params }
func (t *DidChangeNotebookDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidChangeNotebookDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidChangeNotebookDocumentNotification(result)
   return nil
}
// A notification sent when a notebook closes.
// 
// @since 3.17.0
type DidCloseNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidCloseNotebookDocumentParams `json:"params"`
}
func (t DidCloseNotebookDocumentNotification) isMessage() {}
func (t DidCloseNotebookDocumentNotification) isNotification() {}
func (t DidCloseNotebookDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidCloseNotebookDocumentNotification) GetParams() DidCloseNotebookDocumentParams { return t.Params }
func (t *DidCloseNotebookDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidCloseNotebookDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidCloseNotebookDocumentNotification(result)
   return nil
}
// A notification sent when a notebook opens.
// 
// @since 3.17.0
type DidOpenNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidOpenNotebookDocumentParams `json:"params"`
}
func (t DidOpenNotebookDocumentNotification) isMessage() {}
func (t DidOpenNotebookDocumentNotification) isNotification() {}
func (t DidOpenNotebookDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidOpenNotebookDocumentNotification) GetParams() DidOpenNotebookDocumentParams { return t.Params }
func (t *DidOpenNotebookDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidOpenNotebookDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidOpenNotebookDocumentNotification(result)
   return nil
}
// A notification sent when a notebook document is saved.
// 
// @since 3.17.0
type DidSaveNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidSaveNotebookDocumentParams `json:"params"`
}
func (t DidSaveNotebookDocumentNotification) isMessage() {}
func (t DidSaveNotebookDocumentNotification) isNotification() {}
func (t DidSaveNotebookDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidSaveNotebookDocumentNotification) GetParams() DidSaveNotebookDocumentParams { return t.Params }
func (t *DidSaveNotebookDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidSaveNotebookDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidSaveNotebookDocumentNotification(result)
   return nil
}
// The telemetry event notification is sent from the server to the client to ask
// the client to log telemetry data.
type TelemetryEventNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params any `json:"params"`
}
func (t TelemetryEventNotification) isMessage() {}
func (t TelemetryEventNotification) isNotification() {}
func (t TelemetryEventNotification) GetMethod() MethodKind { return t.Method }
func (t TelemetryEventNotification) GetParams() any { return t.Params }
func (t *TelemetryEventNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias TelemetryEventNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = TelemetryEventNotification(result)
   return nil
}
// The document change notification is sent from the client to the server to signal
// changes to a text document.
type DidChangeTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidChangeTextDocumentParams `json:"params"`
}
func (t DidChangeTextDocumentNotification) isMessage() {}
func (t DidChangeTextDocumentNotification) isNotification() {}
func (t DidChangeTextDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidChangeTextDocumentNotification) GetParams() DidChangeTextDocumentParams { return t.Params }
func (t *DidChangeTextDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidChangeTextDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidChangeTextDocumentNotification(result)
   return nil
}
// The document close notification is sent from the client to the server when
// the document got closed in the client. The document's truth now exists where
// the document's uri points to (e.g. if the document's uri is a file uri the
// truth now exists on disk). As with the open notification the close notification
// is about managing the document's content. Receiving a close notification
// doesn't mean that the document was open in an editor before. A close
// notification requires a previous open notification to be sent.
type DidCloseTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidCloseTextDocumentParams `json:"params"`
}
func (t DidCloseTextDocumentNotification) isMessage() {}
func (t DidCloseTextDocumentNotification) isNotification() {}
func (t DidCloseTextDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidCloseTextDocumentNotification) GetParams() DidCloseTextDocumentParams { return t.Params }
func (t *DidCloseTextDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidCloseTextDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidCloseTextDocumentNotification(result)
   return nil
}
// The document open notification is sent from the client to the server to signal
// newly opened text documents. The document's truth is now managed by the client
// and the server must not try to read the document's truth using the document's
// uri. Open in this sense means it is managed by the client. It doesn't necessarily
// mean that its content is presented in an editor. An open notification must not
// be sent more than once without a corresponding close notification send before.
// This means open and close notification must be balanced and the max open count
// is one.
type DidOpenTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidOpenTextDocumentParams `json:"params"`
}
func (t DidOpenTextDocumentNotification) isMessage() {}
func (t DidOpenTextDocumentNotification) isNotification() {}
func (t DidOpenTextDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidOpenTextDocumentNotification) GetParams() DidOpenTextDocumentParams { return t.Params }
func (t *DidOpenTextDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidOpenTextDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidOpenTextDocumentNotification(result)
   return nil
}
// The document save notification is sent from the client to the server when
// the document got saved in the client.
type DidSaveTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidSaveTextDocumentParams `json:"params"`
}
func (t DidSaveTextDocumentNotification) isMessage() {}
func (t DidSaveTextDocumentNotification) isNotification() {}
func (t DidSaveTextDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t DidSaveTextDocumentNotification) GetParams() DidSaveTextDocumentParams { return t.Params }
func (t *DidSaveTextDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidSaveTextDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidSaveTextDocumentNotification(result)
   return nil
}
// Diagnostics notification are sent from the server to the client to signal
// results of validation runs.
type PublishDiagnosticsNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params PublishDiagnosticsParams `json:"params"`
}
func (t PublishDiagnosticsNotification) isMessage() {}
func (t PublishDiagnosticsNotification) isNotification() {}
func (t PublishDiagnosticsNotification) GetMethod() MethodKind { return t.Method }
func (t PublishDiagnosticsNotification) GetParams() PublishDiagnosticsParams { return t.Params }
func (t *PublishDiagnosticsNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias PublishDiagnosticsNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = PublishDiagnosticsNotification(result)
   return nil
}
// A document will save notification is sent from the client to the server before
// the document is actually saved.
type WillSaveTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params WillSaveTextDocumentParams `json:"params"`
}
func (t WillSaveTextDocumentNotification) isMessage() {}
func (t WillSaveTextDocumentNotification) isNotification() {}
func (t WillSaveTextDocumentNotification) GetMethod() MethodKind { return t.Method }
func (t WillSaveTextDocumentNotification) GetParams() WillSaveTextDocumentParams { return t.Params }
func (t *WillSaveTextDocumentNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WillSaveTextDocumentNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WillSaveTextDocumentNotification(result)
   return nil
}
// The log message notification is sent from the server to the client to ask
// the client to log a particular message.
type LogMessageNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params LogMessageParams `json:"params"`
}
func (t LogMessageNotification) isMessage() {}
func (t LogMessageNotification) isNotification() {}
func (t LogMessageNotification) GetMethod() MethodKind { return t.Method }
func (t LogMessageNotification) GetParams() LogMessageParams { return t.Params }
func (t *LogMessageNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias LogMessageNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = LogMessageNotification(result)
   return nil
}
// The show message notification is sent from a server to a client to ask
// the client to display a particular message in the user interface.
type ShowMessageNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params ShowMessageParams `json:"params"`
}
func (t ShowMessageNotification) isMessage() {}
func (t ShowMessageNotification) isNotification() {}
func (t ShowMessageNotification) GetMethod() MethodKind { return t.Method }
func (t ShowMessageNotification) GetParams() ShowMessageParams { return t.Params }
func (t *ShowMessageNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias ShowMessageNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = ShowMessageNotification(result)
   return nil
}
// The `window/workDoneProgress/cancel` notification is sent from  the client to the server to cancel a progress
// initiated on the server side.
type WorkDoneProgressCancelNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params WorkDoneProgressCancelParams `json:"params"`
}
func (t WorkDoneProgressCancelNotification) isMessage() {}
func (t WorkDoneProgressCancelNotification) isNotification() {}
func (t WorkDoneProgressCancelNotification) GetMethod() MethodKind { return t.Method }
func (t WorkDoneProgressCancelNotification) GetParams() WorkDoneProgressCancelParams { return t.Params }
func (t *WorkDoneProgressCancelNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias WorkDoneProgressCancelNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = WorkDoneProgressCancelNotification(result)
   return nil
}
// The configuration change notification is sent from the client to the server
// when the client's configuration has changed. The notification contains
// the changed configuration as defined by the language client.
type DidChangeConfigurationNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidChangeConfigurationParams `json:"params"`
}
func (t DidChangeConfigurationNotification) isMessage() {}
func (t DidChangeConfigurationNotification) isNotification() {}
func (t DidChangeConfigurationNotification) GetMethod() MethodKind { return t.Method }
func (t DidChangeConfigurationNotification) GetParams() DidChangeConfigurationParams { return t.Params }
func (t *DidChangeConfigurationNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidChangeConfigurationNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidChangeConfigurationNotification(result)
   return nil
}
// The watched files notification is sent from the client to the server when
// the client detects changes to file watched by the language client.
type DidChangeWatchedFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidChangeWatchedFilesParams `json:"params"`
}
func (t DidChangeWatchedFilesNotification) isMessage() {}
func (t DidChangeWatchedFilesNotification) isNotification() {}
func (t DidChangeWatchedFilesNotification) GetMethod() MethodKind { return t.Method }
func (t DidChangeWatchedFilesNotification) GetParams() DidChangeWatchedFilesParams { return t.Params }
func (t *DidChangeWatchedFilesNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidChangeWatchedFilesNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidChangeWatchedFilesNotification(result)
   return nil
}
// The `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace
// folder configuration changes.
type DidChangeWorkspaceFoldersNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DidChangeWorkspaceFoldersParams `json:"params"`
}
func (t DidChangeWorkspaceFoldersNotification) isMessage() {}
func (t DidChangeWorkspaceFoldersNotification) isNotification() {}
func (t DidChangeWorkspaceFoldersNotification) GetMethod() MethodKind { return t.Method }
func (t DidChangeWorkspaceFoldersNotification) GetParams() DidChangeWorkspaceFoldersParams { return t.Params }
func (t *DidChangeWorkspaceFoldersNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidChangeWorkspaceFoldersNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidChangeWorkspaceFoldersNotification(result)
   return nil
}
// The did create files notification is sent from the client to the server when
// files were created from within the client.
// 
// @since 3.16.0
type DidCreateFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params CreateFilesParams `json:"params"`
}
func (t DidCreateFilesNotification) isMessage() {}
func (t DidCreateFilesNotification) isNotification() {}
func (t DidCreateFilesNotification) GetMethod() MethodKind { return t.Method }
func (t DidCreateFilesNotification) GetParams() CreateFilesParams { return t.Params }
func (t *DidCreateFilesNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidCreateFilesNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidCreateFilesNotification(result)
   return nil
}
// The will delete files request is sent from the client to the server before files are actually
// deleted as long as the deletion is triggered from within the client.
// 
// @since 3.16.0
type DidDeleteFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params DeleteFilesParams `json:"params"`
}
func (t DidDeleteFilesNotification) isMessage() {}
func (t DidDeleteFilesNotification) isNotification() {}
func (t DidDeleteFilesNotification) GetMethod() MethodKind { return t.Method }
func (t DidDeleteFilesNotification) GetParams() DeleteFilesParams { return t.Params }
func (t *DidDeleteFilesNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidDeleteFilesNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidDeleteFilesNotification(result)
   return nil
}
// The did rename files notification is sent from the client to the server when
// files were renamed from within the client.
// 
// @since 3.16.0
type DidRenameFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method MethodKind `json:"method"`
	Params RenameFilesParams `json:"params"`
}
func (t DidRenameFilesNotification) isMessage() {}
func (t DidRenameFilesNotification) isNotification() {}
func (t DidRenameFilesNotification) GetMethod() MethodKind { return t.Method }
func (t DidRenameFilesNotification) GetParams() RenameFilesParams { return t.Params }
func (t *DidRenameFilesNotification) UnmarshalJSON(x []byte) error {
   var m map[string]any
   if err := json.Unmarshal(x, &m); err != nil {
       return err
   }
   if _, exists := m["method"]; !exists {
       return fmt.Errorf("Missing required request field: method")
   }
   if _, exists := m["jsonrpc"]; !exists {
       return fmt.Errorf("Missing required request field: jsonrpc")
   }
  type Alias DidRenameFilesNotification
   var result Alias
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if err := decoder.Decode(&result); err != nil {
       return err
	}
	*t = DidRenameFilesNotification(result)
   return nil
}


// Defines how values from a set of defaults and an individual item will be
// merged.
// 
// @since 3.18.0
type ApplyKind uint32
const (
	ApplyKindMerge ApplyKind = 2
	ApplyKindReplace ApplyKind = 1
)
func (t ApplyKind) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid ApplyKind: %v", t)
}

func (t *ApplyKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := ApplyKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *ApplyKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A set of predefined code action kinds
type CodeActionKind string
const (
	CodeActionKindEmpty CodeActionKind = ""
	CodeActionKindNotebook CodeActionKind = "notebook"
	CodeActionKindQuickFix CodeActionKind = "quickfix"
	CodeActionKindRefactor CodeActionKind = "refactor"
	CodeActionKindRefactorExtract CodeActionKind = "refactor.extract"
	CodeActionKindRefactorInline CodeActionKind = "refactor.inline"
	CodeActionKindRefactorMove CodeActionKind = "refactor.move"
	CodeActionKindRefactorRewrite CodeActionKind = "refactor.rewrite"
	CodeActionKindSource CodeActionKind = "source"
	CodeActionKindSourceFixAll CodeActionKind = "source.fixAll"
	CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"
)

// Code action tags are extra annotations that tweak the behavior of a code action.
// 
// @since 3.18.0 - proposed
type CodeActionTag uint32
const (
	CodeActionTagLLMGenerated CodeActionTag = 1
)
func (t CodeActionTag) validate() error {
	switch t {
	case 1:
	return nil
	}
	return fmt.Errorf("invalid CodeActionTag: %v", t)
}

func (t *CodeActionTag) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := CodeActionTag(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *CodeActionTag) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The reason why code actions were requested.
// 
// @since 3.17.0
type CodeActionTriggerKind uint32
const (
	CodeActionTriggerKindAutomatic CodeActionTriggerKind = 2
	CodeActionTriggerKindInvoked CodeActionTriggerKind = 1
)
func (t CodeActionTriggerKind) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid CodeActionTriggerKind: %v", t)
}

func (t *CodeActionTriggerKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := CodeActionTriggerKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *CodeActionTriggerKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The kind of a completion entry.
type CompletionItemKind uint32
const (
	CompletionItemKindClass CompletionItemKind = 7
	CompletionItemKindColor CompletionItemKind = 16
	CompletionItemKindConstant CompletionItemKind = 21
	CompletionItemKindConstructor CompletionItemKind = 4
	CompletionItemKindEnum CompletionItemKind = 13
	CompletionItemKindEnumMember CompletionItemKind = 20
	CompletionItemKindEvent CompletionItemKind = 23
	CompletionItemKindField CompletionItemKind = 5
	CompletionItemKindFile CompletionItemKind = 17
	CompletionItemKindFolder CompletionItemKind = 19
	CompletionItemKindFunction CompletionItemKind = 3
	CompletionItemKindInterface CompletionItemKind = 8
	CompletionItemKindKeyword CompletionItemKind = 14
	CompletionItemKindMethod CompletionItemKind = 2
	CompletionItemKindModule CompletionItemKind = 9
	CompletionItemKindOperator CompletionItemKind = 24
	CompletionItemKindProperty CompletionItemKind = 10
	CompletionItemKindReference CompletionItemKind = 18
	CompletionItemKindSnippet CompletionItemKind = 15
	CompletionItemKindStruct CompletionItemKind = 22
	CompletionItemKindText CompletionItemKind = 1
	CompletionItemKindTypeParameter CompletionItemKind = 25
	CompletionItemKindUnit CompletionItemKind = 11
	CompletionItemKindValue CompletionItemKind = 12
	CompletionItemKindVariable CompletionItemKind = 6
)
func (t CompletionItemKind) validate() error {
	switch t {
	case 8,4,10,5,21,1,16,17,9,14,12,23,18,15,24,6,22,13,2,3,25,19,20,7,11:
	return nil
	}
	return fmt.Errorf("invalid CompletionItemKind: %v", t)
}

func (t *CompletionItemKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := CompletionItemKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *CompletionItemKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Completion item tags are extra annotations that tweak the rendering of a completion
// item.
// 
// @since 3.15.0
type CompletionItemTag uint32
const (
	CompletionItemTagDeprecated CompletionItemTag = 1
)
func (t CompletionItemTag) validate() error {
	switch t {
	case 1:
	return nil
	}
	return fmt.Errorf("invalid CompletionItemTag: %v", t)
}

func (t *CompletionItemTag) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := CompletionItemTag(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *CompletionItemTag) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// How a completion was triggered
type CompletionTriggerKind uint32
const (
	CompletionTriggerKindInvoked CompletionTriggerKind = 1
	CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2
	CompletionTriggerKindTriggerForIncompleteCompletions CompletionTriggerKind = 3
)
func (t CompletionTriggerKind) validate() error {
	switch t {
	case 1,2,3:
	return nil
	}
	return fmt.Errorf("invalid CompletionTriggerKind: %v", t)
}

func (t *CompletionTriggerKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := CompletionTriggerKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *CompletionTriggerKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The diagnostic's severity.
type DiagnosticSeverity uint32
const (
	DiagnosticSeverityError DiagnosticSeverity = 1
	DiagnosticSeverityHint DiagnosticSeverity = 4
	DiagnosticSeverityInformation DiagnosticSeverity = 3
	DiagnosticSeverityWarning DiagnosticSeverity = 2
)
func (t DiagnosticSeverity) validate() error {
	switch t {
	case 1,4,2,3:
	return nil
	}
	return fmt.Errorf("invalid DiagnosticSeverity: %v", t)
}

func (t *DiagnosticSeverity) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := DiagnosticSeverity(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *DiagnosticSeverity) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The diagnostic tags.
// 
// @since 3.15.0
type DiagnosticTag uint32
const (
	DiagnosticTagDeprecated DiagnosticTag = 2
	DiagnosticTagUnnecessary DiagnosticTag = 1
)
func (t DiagnosticTag) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid DiagnosticTag: %v", t)
}

func (t *DiagnosticTag) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := DiagnosticTag(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *DiagnosticTag) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The document diagnostic report kinds.
// 
// @since 3.17.0
type DocumentDiagnosticReportKind string
const (
	DocumentDiagnosticReportKindFull DocumentDiagnosticReportKind = "full"
	DocumentDiagnosticReportKindUnchanged DocumentDiagnosticReportKind = "unchanged"
)
func (t DocumentDiagnosticReportKind) validate() error {
	switch t {
	case "unchanged","full":
	return nil
	}
	return fmt.Errorf("invalid DocumentDiagnosticReportKind: %v", t)
}

func (t *DocumentDiagnosticReportKind) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := DocumentDiagnosticReportKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *DocumentDiagnosticReportKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A document highlight kind.
type DocumentHighlightKind uint32
const (
	DocumentHighlightKindRead DocumentHighlightKind = 2
	DocumentHighlightKindText DocumentHighlightKind = 1
	DocumentHighlightKindWrite DocumentHighlightKind = 3
)
func (t DocumentHighlightKind) validate() error {
	switch t {
	case 2,3,1:
	return nil
	}
	return fmt.Errorf("invalid DocumentHighlightKind: %v", t)
}

func (t *DocumentHighlightKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := DocumentHighlightKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *DocumentHighlightKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Predefined error codes.
type ErrorCodes int32
const (
	ErrorCodesInternalError ErrorCodes = -32603
	ErrorCodesInvalidParams ErrorCodes = -32602
	ErrorCodesInvalidRequest ErrorCodes = -32600
	ErrorCodesMethodNotFound ErrorCodes = -32601
	ErrorCodesParseError ErrorCodes = -32700
	ErrorCodesServerNotInitialized ErrorCodes = -32002
	ErrorCodesUnknownErrorCode ErrorCodes = -32001
)

type FailureHandlingKind string
const (
	FailureHandlingKindAbort FailureHandlingKind = "abort"
	FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"
	FailureHandlingKindTransactional FailureHandlingKind = "transactional"
	FailureHandlingKindUndo FailureHandlingKind = "undo"
)
func (t FailureHandlingKind) validate() error {
	switch t {
	case "undo","textOnlyTransactional","transactional","abort":
	return nil
	}
	return fmt.Errorf("invalid FailureHandlingKind: %v", t)
}

func (t *FailureHandlingKind) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := FailureHandlingKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *FailureHandlingKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The file event type
type FileChangeType uint32
const (
	FileChangeTypeChanged FileChangeType = 2
	FileChangeTypeCreated FileChangeType = 1
	FileChangeTypeDeleted FileChangeType = 3
)
func (t FileChangeType) validate() error {
	switch t {
	case 2,3,1:
	return nil
	}
	return fmt.Errorf("invalid FileChangeType: %v", t)
}

func (t *FileChangeType) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := FileChangeType(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *FileChangeType) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A pattern kind describing if a glob pattern matches a file a folder or
// both.
// 
// @since 3.16.0
type FileOperationPatternKind string
const (
	FileOperationPatternKindFile FileOperationPatternKind = "file"
	FileOperationPatternKindFolder FileOperationPatternKind = "folder"
)
func (t FileOperationPatternKind) validate() error {
	switch t {
	case "folder","file":
	return nil
	}
	return fmt.Errorf("invalid FileOperationPatternKind: %v", t)
}

func (t *FileOperationPatternKind) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := FileOperationPatternKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *FileOperationPatternKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A set of predefined range kinds.
type FoldingRangeKind string
const (
	FoldingRangeKindComment FoldingRangeKind = "comment"
	FoldingRangeKindImports FoldingRangeKind = "imports"
	FoldingRangeKindRegion FoldingRangeKind = "region"
)

// Inlay hint kinds.
// 
// @since 3.17.0
type InlayHintKind uint32
const (
	InlayHintKindParameter InlayHintKind = 2
	InlayHintKindType InlayHintKind = 1
)
func (t InlayHintKind) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid InlayHintKind: %v", t)
}

func (t *InlayHintKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := InlayHintKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *InlayHintKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Describes how an {@link InlineCompletionItemProvider inline completion provider} was triggered.
// 
// @since 3.18.0
// @proposed
type InlineCompletionTriggerKind uint32
const (
	InlineCompletionTriggerKindAutomatic InlineCompletionTriggerKind = 2
	InlineCompletionTriggerKindInvoked InlineCompletionTriggerKind = 1
)
func (t InlineCompletionTriggerKind) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid InlineCompletionTriggerKind: %v", t)
}

func (t *InlineCompletionTriggerKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := InlineCompletionTriggerKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *InlineCompletionTriggerKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
type InsertTextFormat uint32
const (
	InsertTextFormatPlainText InsertTextFormat = 1
	InsertTextFormatSnippet InsertTextFormat = 2
)
func (t InsertTextFormat) validate() error {
	switch t {
	case 1,2:
	return nil
	}
	return fmt.Errorf("invalid InsertTextFormat: %v", t)
}

func (t *InsertTextFormat) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := InsertTextFormat(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *InsertTextFormat) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// How whitespace and indentation is handled during completion
// item insertion.
// 
// @since 3.16.0
type InsertTextMode uint32
const (
	InsertTextModeAdjustIndentation InsertTextMode = 2
	InsertTextModeAsIs InsertTextMode = 1
)
func (t InsertTextMode) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid InsertTextMode: %v", t)
}

func (t *InsertTextMode) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := InsertTextMode(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *InsertTextMode) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
type LSPErrorCodes int32
const (
	LSPErrorCodesContentModified LSPErrorCodes = -32801
	LSPErrorCodesRequestCancelled LSPErrorCodes = -32800
	LSPErrorCodesRequestFailed LSPErrorCodes = -32803
	LSPErrorCodesServerCancelled LSPErrorCodes = -32802
)

// Predefined Language kinds
// @since 3.18.0
type LanguageKind string
const (
	LanguageKindABAP LanguageKind = "abap"
	LanguageKindBibTeX LanguageKind = "bibtex"
	LanguageKindC LanguageKind = "c"
	LanguageKindCPP LanguageKind = "cpp"
	LanguageKindCSS LanguageKind = "css"
	LanguageKindCSharp LanguageKind = "csharp"
	LanguageKindClojure LanguageKind = "clojure"
	LanguageKindCoffeescript LanguageKind = "coffeescript"
	LanguageKindD LanguageKind = "d"
	LanguageKindDart LanguageKind = "dart"
	LanguageKindDelphi LanguageKind = "pascal"
	LanguageKindDiff LanguageKind = "diff"
	LanguageKindDockerfile LanguageKind = "dockerfile"
	LanguageKindElixir LanguageKind = "elixir"
	LanguageKindErlang LanguageKind = "erlang"
	LanguageKindFSharp LanguageKind = "fsharp"
	LanguageKindGitCommit LanguageKind = "git-commit"
	LanguageKindGitRebase LanguageKind = "rebase"
	LanguageKindGo LanguageKind = "go"
	LanguageKindGroovy LanguageKind = "groovy"
	LanguageKindHTML LanguageKind = "html"
	LanguageKindHandlebars LanguageKind = "handlebars"
	LanguageKindHaskell LanguageKind = "haskell"
	LanguageKindIni LanguageKind = "ini"
	LanguageKindJSON LanguageKind = "json"
	LanguageKindJava LanguageKind = "java"
	LanguageKindJavaScript LanguageKind = "javascript"
	LanguageKindJavaScriptReact LanguageKind = "javascriptreact"
	LanguageKindLaTeX LanguageKind = "latex"
	LanguageKindLess LanguageKind = "less"
	LanguageKindLua LanguageKind = "lua"
	LanguageKindMakefile LanguageKind = "makefile"
	LanguageKindMarkdown LanguageKind = "markdown"
	LanguageKindObjectiveC LanguageKind = "objective-c"
	LanguageKindObjectiveCPP LanguageKind = "objective-cpp"
	LanguageKindPHP LanguageKind = "php"
	LanguageKindPascal LanguageKind = "pascal"
	LanguageKindPerl LanguageKind = "perl"
	LanguageKindPerl6 LanguageKind = "perl6"
	LanguageKindPowershell LanguageKind = "powershell"
	LanguageKindPug LanguageKind = "jade"
	LanguageKindPython LanguageKind = "python"
	LanguageKindR LanguageKind = "r"
	LanguageKindRazor LanguageKind = "razor"
	LanguageKindRuby LanguageKind = "ruby"
	LanguageKindRust LanguageKind = "rust"
	LanguageKindSASS LanguageKind = "sass"
	LanguageKindSCSS LanguageKind = "scss"
	LanguageKindSQL LanguageKind = "sql"
	LanguageKindScala LanguageKind = "scala"
	LanguageKindShaderLab LanguageKind = "shaderlab"
	LanguageKindShellScript LanguageKind = "shellscript"
	LanguageKindSwift LanguageKind = "swift"
	LanguageKindTeX LanguageKind = "tex"
	LanguageKindTypeScript LanguageKind = "typescript"
	LanguageKindTypeScriptReact LanguageKind = "typescriptreact"
	LanguageKindVisualBasic LanguageKind = "vb"
	LanguageKindWindowsBat LanguageKind = "bat"
	LanguageKindXML LanguageKind = "xml"
	LanguageKindXSL LanguageKind = "xsl"
	LanguageKindYAML LanguageKind = "yaml"
)

// Describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
// 
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
type MarkupKind string
const (
	MarkupKindMarkdown MarkupKind = "markdown"
	MarkupKindPlainText MarkupKind = "plaintext"
)
func (t MarkupKind) validate() error {
	switch t {
	case "markdown","plaintext":
	return nil
	}
	return fmt.Errorf("invalid MarkupKind: %v", t)
}

func (t *MarkupKind) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := MarkupKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *MarkupKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The message type
type MessageType uint32
const (
	MessageTypeDebug MessageType = 5
	MessageTypeError MessageType = 1
	MessageTypeInfo MessageType = 3
	MessageTypeLog MessageType = 4
	MessageTypeWarning MessageType = 2
)
func (t MessageType) validate() error {
	switch t {
	case 2,3,1,4,5:
	return nil
	}
	return fmt.Errorf("invalid MessageType: %v", t)
}

func (t *MessageType) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := MessageType(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *MessageType) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// The moniker kind.
// 
// @since 3.16.0
type MonikerKind string
const (
	MonikerKindExport MonikerKind = "export"
	MonikerKindImport MonikerKind = "import"
	MonikerKindLocal MonikerKind = "local"
)
func (t MonikerKind) validate() error {
	switch t {
	case "export","local","import":
	return nil
	}
	return fmt.Errorf("invalid MonikerKind: %v", t)
}

func (t *MonikerKind) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := MonikerKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *MonikerKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A notebook cell kind.
// 
// @since 3.17.0
type NotebookCellKind uint32
const (
	NotebookCellKindCode NotebookCellKind = 2
	NotebookCellKindMarkup NotebookCellKind = 1
)
func (t NotebookCellKind) validate() error {
	switch t {
	case 2,1:
	return nil
	}
	return fmt.Errorf("invalid NotebookCellKind: %v", t)
}

func (t *NotebookCellKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := NotebookCellKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *NotebookCellKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A set of predefined position encoding kinds.
// 
// @since 3.17.0
type PositionEncodingKind string
const (
	PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"
	PositionEncodingKindUTF32 PositionEncodingKind = "utf-32"
	PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"
)

type PrepareSupportDefaultBehavior uint32
const (
	PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1
)
func (t PrepareSupportDefaultBehavior) validate() error {
	switch t {
	case 1:
	return nil
	}
	return fmt.Errorf("invalid PrepareSupportDefaultBehavior: %v", t)
}

func (t *PrepareSupportDefaultBehavior) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := PrepareSupportDefaultBehavior(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *PrepareSupportDefaultBehavior) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
type ResourceOperationKind string
const (
	ResourceOperationKindCreate ResourceOperationKind = "create"
	ResourceOperationKindDelete ResourceOperationKind = "delete"
	ResourceOperationKindRename ResourceOperationKind = "rename"
)
func (t ResourceOperationKind) validate() error {
	switch t {
	case "create","delete","rename":
	return nil
	}
	return fmt.Errorf("invalid ResourceOperationKind: %v", t)
}

func (t *ResourceOperationKind) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := ResourceOperationKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *ResourceOperationKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A set of predefined token modifiers. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
type SemanticTokenModifiers string
const (
	SemanticTokenModifiersAbstract SemanticTokenModifiers = "abstract"
	SemanticTokenModifiersAsync SemanticTokenModifiers = "async"
	SemanticTokenModifiersDeclaration SemanticTokenModifiers = "declaration"
	SemanticTokenModifiersDefaultLibrary SemanticTokenModifiers = "defaultLibrary"
	SemanticTokenModifiersDefinition SemanticTokenModifiers = "definition"
	SemanticTokenModifiersDeprecated SemanticTokenModifiers = "deprecated"
	SemanticTokenModifiersDocumentation SemanticTokenModifiers = "documentation"
	SemanticTokenModifiersModification SemanticTokenModifiers = "modification"
	SemanticTokenModifiersReadonly SemanticTokenModifiers = "readonly"
	SemanticTokenModifiersStatic SemanticTokenModifiers = "static"
)

// A set of predefined token types. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
type SemanticTokenTypes string
const (
	SemanticTokenTypesClass SemanticTokenTypes = "class"
	SemanticTokenTypesComment SemanticTokenTypes = "comment"
	SemanticTokenTypesDecorator SemanticTokenTypes = "decorator"
	SemanticTokenTypesEnum SemanticTokenTypes = "enum"
	SemanticTokenTypesEnumMember SemanticTokenTypes = "enumMember"
	SemanticTokenTypesEvent SemanticTokenTypes = "event"
	SemanticTokenTypesFunction SemanticTokenTypes = "function"
	SemanticTokenTypesInterface SemanticTokenTypes = "interface"
	SemanticTokenTypesKeyword SemanticTokenTypes = "keyword"
	SemanticTokenTypesLabel SemanticTokenTypes = "label"
	SemanticTokenTypesMacro SemanticTokenTypes = "macro"
	SemanticTokenTypesMethod SemanticTokenTypes = "method"
	SemanticTokenTypesModifier SemanticTokenTypes = "modifier"
	SemanticTokenTypesNamespace SemanticTokenTypes = "namespace"
	SemanticTokenTypesNumber SemanticTokenTypes = "number"
	SemanticTokenTypesOperator SemanticTokenTypes = "operator"
	SemanticTokenTypesParameter SemanticTokenTypes = "parameter"
	SemanticTokenTypesProperty SemanticTokenTypes = "property"
	SemanticTokenTypesRegexp SemanticTokenTypes = "regexp"
	SemanticTokenTypesString SemanticTokenTypes = "string"
	SemanticTokenTypesStruct SemanticTokenTypes = "struct"
	SemanticTokenTypesType SemanticTokenTypes = "type"
	SemanticTokenTypesTypeParameter SemanticTokenTypes = "typeParameter"
	SemanticTokenTypesVariable SemanticTokenTypes = "variable"
)

// How a signature help was triggered.
// 
// @since 3.15.0
type SignatureHelpTriggerKind uint32
const (
	SignatureHelpTriggerKindContentChange SignatureHelpTriggerKind = 3
	SignatureHelpTriggerKindInvoked SignatureHelpTriggerKind = 1
	SignatureHelpTriggerKindTriggerCharacter SignatureHelpTriggerKind = 2
)
func (t SignatureHelpTriggerKind) validate() error {
	switch t {
	case 1,2,3:
	return nil
	}
	return fmt.Errorf("invalid SignatureHelpTriggerKind: %v", t)
}

func (t *SignatureHelpTriggerKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := SignatureHelpTriggerKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *SignatureHelpTriggerKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// A symbol kind.
type SymbolKind uint32
const (
	SymbolKindArray SymbolKind = 18
	SymbolKindBoolean SymbolKind = 17
	SymbolKindClass SymbolKind = 5
	SymbolKindConstant SymbolKind = 14
	SymbolKindConstructor SymbolKind = 9
	SymbolKindEnum SymbolKind = 10
	SymbolKindEnumMember SymbolKind = 22
	SymbolKindEvent SymbolKind = 24
	SymbolKindField SymbolKind = 8
	SymbolKindFile SymbolKind = 1
	SymbolKindFunction SymbolKind = 12
	SymbolKindInterface SymbolKind = 11
	SymbolKindKey SymbolKind = 20
	SymbolKindMethod SymbolKind = 6
	SymbolKindModule SymbolKind = 2
	SymbolKindNamespace SymbolKind = 3
	SymbolKindNull SymbolKind = 21
	SymbolKindNumber SymbolKind = 16
	SymbolKindObject SymbolKind = 19
	SymbolKindOperator SymbolKind = 25
	SymbolKindPackage SymbolKind = 4
	SymbolKindProperty SymbolKind = 7
	SymbolKindString SymbolKind = 15
	SymbolKindStruct SymbolKind = 23
	SymbolKindTypeParameter SymbolKind = 26
	SymbolKindVariable SymbolKind = 13
)
func (t SymbolKind) validate() error {
	switch t {
	case 8,4,10,5,21,26,1,16,17,9,14,12,18,6,15,24,23,22,13,2,3,25,19,20,7,11:
	return nil
	}
	return fmt.Errorf("invalid SymbolKind: %v", t)
}

func (t *SymbolKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := SymbolKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *SymbolKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Symbol tags are extra annotations that tweak the rendering of a symbol.
// 
// @since 3.16
type SymbolTag uint32
const (
	SymbolTagDeprecated SymbolTag = 1
)
func (t SymbolTag) validate() error {
	switch t {
	case 1:
	return nil
	}
	return fmt.Errorf("invalid SymbolTag: %v", t)
}

func (t *SymbolTag) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := SymbolTag(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *SymbolTag) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Represents reasons why a text document is saved.
type TextDocumentSaveReason uint32
const (
	TextDocumentSaveReasonAfterDelay TextDocumentSaveReason = 2
	TextDocumentSaveReasonFocusOut TextDocumentSaveReason = 3
	TextDocumentSaveReasonManual TextDocumentSaveReason = 1
)
func (t TextDocumentSaveReason) validate() error {
	switch t {
	case 2,1,3:
	return nil
	}
	return fmt.Errorf("invalid TextDocumentSaveReason: %v", t)
}

func (t *TextDocumentSaveReason) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := TextDocumentSaveReason(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *TextDocumentSaveReason) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Defines how the host (editor) should sync
// document changes to the language server.
type TextDocumentSyncKind uint32
const (
	TextDocumentSyncKindFull TextDocumentSyncKind = 1
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
	TextDocumentSyncKindNone TextDocumentSyncKind = 0
)
func (t TextDocumentSyncKind) validate() error {
	switch t {
	case 1,2,0:
	return nil
	}
	return fmt.Errorf("invalid TextDocumentSyncKind: %v", t)
}

func (t *TextDocumentSyncKind) UnmarshalJSON(x []byte) error {
	var test uint32
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := TextDocumentSyncKind(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *TextDocumentSyncKind) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
type TokenFormat string
const (
	TokenFormatRelative TokenFormat = "relative"
)
func (t TokenFormat) validate() error {
	switch t {
	case "relative":
	return nil
	}
	return fmt.Errorf("invalid TokenFormat: %v", t)
}

func (t *TokenFormat) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := TokenFormat(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *TokenFormat) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
type TraceValue string
const (
	TraceValueMessages TraceValue = "messages"
	TraceValueOff TraceValue = "off"
	TraceValueVerbose TraceValue = "verbose"
)
func (t TraceValue) validate() error {
	switch t {
	case "off","messages","verbose":
	return nil
	}
	return fmt.Errorf("invalid TraceValue: %v", t)
}

func (t *TraceValue) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := TraceValue(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *TraceValue) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
// Moniker uniqueness level to define scope of the moniker.
// 
// @since 3.16.0
type UniquenessLevel string
const (
	UniquenessLevelDocument UniquenessLevel = "document"
	UniquenessLevelGlobal UniquenessLevel = "global"
	UniquenessLevelGroup UniquenessLevel = "group"
	UniquenessLevelProject UniquenessLevel = "project"
	UniquenessLevelScheme UniquenessLevel = "scheme"
)
func (t UniquenessLevel) validate() error {
	switch t {
	case "group","scheme","project","document","global":
	return nil
	}
	return fmt.Errorf("invalid UniquenessLevel: %v", t)
}

func (t *UniquenessLevel) UnmarshalJSON(x []byte) error {
	var test string
	if err := json.Unmarshal(x, &test); err != nil {
		return err
	}
	kind := UniquenessLevel(test)
	if err := kind.validate(); err != nil {
		return err
	}
	*t = kind
	return nil
}

func (t *UniquenessLevel) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(t)
}
type WatchKind uint32
const (
	WatchKindChange WatchKind = 2
	WatchKindCreate WatchKind = 1
	WatchKindDelete WatchKind = 4
)

type Or2[A any, B any] struct {
	Value any
}
func (t *Or2[A,B]) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		return &UnmarshalError{"Or2[A,B] cannot be null"}
	}
	var h0 A
	decoder0 := json.NewDecoder(bytes.NewReader(x))
	decoder0.DisallowUnknownFields()
	if err := decoder0.Decode(&h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 B
	decoder1 := json.NewDecoder(bytes.NewReader(x))
	decoder1.DisallowUnknownFields()
	if err := decoder1.Decode(&h1); err == nil {
		t.Value = h1
		return nil
	}
	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))
}
func (t Or2[A,B]) MarshalJSON() ([]byte, error) {

	switch x := t.Value.(type) {
	case A:
		return json.Marshal(x)
	case B:
		return json.Marshal(x)
	}

	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())
}
func (t Or2[A,B]) ConcreteTypes() string {
	var types []string
		var h0 [0]A
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h0).Elem()))
	var h1 [0]B
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h1).Elem()))
	return strings.Join(types, ",")
}
type NullableOr2[A any, B any] struct {
	Value any
}
func (t *NullableOr2[A,B]) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 A
	decoder0 := json.NewDecoder(bytes.NewReader(x))
	decoder0.DisallowUnknownFields()
	if err := decoder0.Decode(&h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 B
	decoder1 := json.NewDecoder(bytes.NewReader(x))
	decoder1.DisallowUnknownFields()
	if err := decoder1.Decode(&h1); err == nil {
		t.Value = h1
		return nil
	}
	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))
}
func (t NullableOr2[A,B]) MarshalJSON() ([]byte, error) {
	if t.Value == nil {
		return []byte("null"), nil
	}
	switch x := t.Value.(type) {
	case A:
		return json.Marshal(x)
	case B:
		return json.Marshal(x)
	}

	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())
}
func (t NullableOr2[A,B]) ConcreteTypes() string {
	var types []string
		var h0 [0]A
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h0).Elem()))
	var h1 [0]B
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h1).Elem()))
	return strings.Join(types, ",")
}
type Or3[A any, B any, C any] struct {
	Value any
}
func (t *Or3[A,B,C]) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		return &UnmarshalError{"Or3[A,B,C] cannot be null"}
	}
	var h0 A
	decoder0 := json.NewDecoder(bytes.NewReader(x))
	decoder0.DisallowUnknownFields()
	if err := decoder0.Decode(&h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 B
	decoder1 := json.NewDecoder(bytes.NewReader(x))
	decoder1.DisallowUnknownFields()
	if err := decoder1.Decode(&h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 C
	decoder2 := json.NewDecoder(bytes.NewReader(x))
	decoder2.DisallowUnknownFields()
	if err := decoder2.Decode(&h2); err == nil {
		t.Value = h2
		return nil
	}
	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))
}
func (t Or3[A,B,C]) MarshalJSON() ([]byte, error) {

	switch x := t.Value.(type) {
	case A:
		return json.Marshal(x)
	case B:
		return json.Marshal(x)
	case C:
		return json.Marshal(x)
	}

	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())
}
func (t Or3[A,B,C]) ConcreteTypes() string {
	var types []string
		var h0 [0]A
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h0).Elem()))
	var h1 [0]B
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h1).Elem()))
	var h2 [0]C
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h2).Elem()))
	return strings.Join(types, ",")
}
type NullableOr3[A any, B any, C any] struct {
	Value any
}
func (t *NullableOr3[A,B,C]) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 A
	decoder0 := json.NewDecoder(bytes.NewReader(x))
	decoder0.DisallowUnknownFields()
	if err := decoder0.Decode(&h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 B
	decoder1 := json.NewDecoder(bytes.NewReader(x))
	decoder1.DisallowUnknownFields()
	if err := decoder1.Decode(&h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 C
	decoder2 := json.NewDecoder(bytes.NewReader(x))
	decoder2.DisallowUnknownFields()
	if err := decoder2.Decode(&h2); err == nil {
		t.Value = h2
		return nil
	}
	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))
}
func (t NullableOr3[A,B,C]) MarshalJSON() ([]byte, error) {
	if t.Value == nil {
		return []byte("null"), nil
	}
	switch x := t.Value.(type) {
	case A:
		return json.Marshal(x)
	case B:
		return json.Marshal(x)
	case C:
		return json.Marshal(x)
	}

	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())
}
func (t NullableOr3[A,B,C]) ConcreteTypes() string {
	var types []string
		var h0 [0]A
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h0).Elem()))
	var h1 [0]B
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h1).Elem()))
	var h2 [0]C
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h2).Elem()))
	return strings.Join(types, ",")
}
type Or4[A any, B any, C any, D any] struct {
	Value any
}
func (t *Or4[A,B,C,D]) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		return &UnmarshalError{"Or4[A,B,C,D] cannot be null"}
	}
	var h0 A
	decoder0 := json.NewDecoder(bytes.NewReader(x))
	decoder0.DisallowUnknownFields()
	if err := decoder0.Decode(&h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 B
	decoder1 := json.NewDecoder(bytes.NewReader(x))
	decoder1.DisallowUnknownFields()
	if err := decoder1.Decode(&h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 C
	decoder2 := json.NewDecoder(bytes.NewReader(x))
	decoder2.DisallowUnknownFields()
	if err := decoder2.Decode(&h2); err == nil {
		t.Value = h2
		return nil
	}
	var h3 D
	decoder3 := json.NewDecoder(bytes.NewReader(x))
	decoder3.DisallowUnknownFields()
	if err := decoder3.Decode(&h3); err == nil {
		t.Value = h3
		return nil
	}
	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))
}
func (t Or4[A,B,C,D]) MarshalJSON() ([]byte, error) {

	switch x := t.Value.(type) {
	case A:
		return json.Marshal(x)
	case B:
		return json.Marshal(x)
	case C:
		return json.Marshal(x)
	case D:
		return json.Marshal(x)
	}

	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())
}
func (t Or4[A,B,C,D]) ConcreteTypes() string {
	var types []string
		var h0 [0]A
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h0).Elem()))
	var h1 [0]B
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h1).Elem()))
	var h2 [0]C
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h2).Elem()))
	var h3 [0]D
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h3).Elem()))
	return strings.Join(types, ",")
}
type NullableOr4[A any, B any, C any, D any] struct {
	Value any
}
func (t *NullableOr4[A,B,C,D]) UnmarshalJSON(x []byte) error {
	if string(x) == "null" {
		t.Value = nil
		return nil
	}
	var h0 A
	decoder0 := json.NewDecoder(bytes.NewReader(x))
	decoder0.DisallowUnknownFields()
	if err := decoder0.Decode(&h0); err == nil {
		t.Value = h0
		return nil
	}
	var h1 B
	decoder1 := json.NewDecoder(bytes.NewReader(x))
	decoder1.DisallowUnknownFields()
	if err := decoder1.Decode(&h1); err == nil {
		t.Value = h1
		return nil
	}
	var h2 C
	decoder2 := json.NewDecoder(bytes.NewReader(x))
	decoder2.DisallowUnknownFields()
	if err := decoder2.Decode(&h2); err == nil {
		t.Value = h2
		return nil
	}
	var h3 D
	decoder3 := json.NewDecoder(bytes.NewReader(x))
	decoder3.DisallowUnknownFields()
	if err := decoder3.Decode(&h3); err == nil {
		t.Value = h3
		return nil
	}
	return fmt.Errorf("expected one of [%s], got %s", t.ConcreteTypes(), string(x))
}
func (t NullableOr4[A,B,C,D]) MarshalJSON() ([]byte, error) {
	if t.Value == nil {
		return []byte("null"), nil
	}
	switch x := t.Value.(type) {
	case A:
		return json.Marshal(x)
	case B:
		return json.Marshal(x)
	case C:
		return json.Marshal(x)
	case D:
		return json.Marshal(x)
	}

	return nil, fmt.Errorf("type %T not one of [%s]", t.Value, t.ConcreteTypes())
}
func (t NullableOr4[A,B,C,D]) ConcreteTypes() string {
	var types []string
		var h0 [0]A
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h0).Elem()))
	var h1 [0]B
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h1).Elem()))
	var h2 [0]C
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h2).Elem()))
	var h3 [0]D
	types = append(types, fmt.Sprintf("%v", reflect.TypeOf(h3).Elem()))
	return strings.Join(types, ",")
}