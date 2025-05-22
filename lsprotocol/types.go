package lsprotocol


import (
	"encoding/json"
	"fmt"
	"bytes"
)
type DocumentUri string

type URI string

// An identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier string

// The declaration of a symbol representation as one or many {@link Location locations}.
type Declaration Or_Declaration

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
type Definition Or_Definition

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
type DocumentDiagnosticReport Or_DocumentDiagnosticReport

// A document filter describes a top level text document or
// a notebook cell document.
// 
// @since 3.17.0 - support for NotebookCellTextDocumentFilter.
type DocumentFilter Or_DocumentFilter

// A document selector is the combination of one or many document filters.
// 
// @sample `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`;
// 
// The use of a string as a document filter is deprecated @since 3.16.0.
type DocumentSelector []DocumentFilter

// The glob pattern. Either a string pattern or a relative pattern.
// 
// @since 3.17.0
type GlobPattern Or_GlobPattern

// Inline value information can be provided by different means:
// - directly as a text value (class InlineValueText).
// - as a name to use for a variable lookup (class InlineValueVariableLookup)
// - as an evaluatable expression (class InlineValueEvaluatableExpression)
// The InlineValue types combines all inline value types into one type.
// 
// @since 3.17.0
type InlineValue Or_InlineValue

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
type MarkedString Or_MarkedString

// A notebook document filter denotes a notebook document by
// different properties. The properties will be match
// against the notebook's URI (same as with documents)
// 
// @since 3.17.0
type NotebookDocumentFilter Or_NotebookDocumentFilter

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


type PrepareRenameResult Or_PrepareRenameResult


type ProgressToken Or_ProgressToken


type RegularExpressionEngineKind string

// An event describing a change to a text document. If only a text is provided
// it is considered to be the full content of the document.
type TextDocumentContentChangeEvent Or_TextDocumentContentChangeEvent

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
type TextDocumentFilter Or_TextDocumentFilter

// A workspace diagnostic document report.
// 
// @since 3.17.0
type WorkspaceDocumentDiagnosticReport Or_WorkspaceDocumentDiagnosticReport

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
func (t Tuple[T1, T2]) UnmarshalJSON(data []byte) error {
   var arr [2]json.RawMessage
   if err := json.Unmarshal(data, &arr); err != nil {
       return err
   }
   if err := json.Unmarshal(arr[0], &t.V1); err != nil {
       return err
   }
   if err := json.Unmarshal(arr[1], &t.V2); err != nil {
       return err
   }
   return nil
}
type ResponseError struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
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
	Metadata *WorkspaceEditMetadata `json:"metadata,omitempty"`
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
	Tags []SymbolTag `json:"tags,omitempty"`
}
// @since 3.16.0
type CallHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
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
// The parameter of a `callHierarchy/incomingCalls` request.
// 
// @since 3.16.0
type CallHierarchyIncomingCallsParams struct {

	Item CallHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
	Tags []SymbolTag `json:"tags,omitempty"`
	// The resource identifier of this item.
	Uri DocumentUri `json:"uri"`
}
// Call hierarchy options used during static registration.
// 
// @since 3.16.0
type CallHierarchyOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
// The parameter of a `callHierarchy/outgoingCalls` request.
// 
// @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {

	Item CallHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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

type CancelParams struct {
	// The request id to cancel.
	Id Or_CancelParams_Id `json:"id"`
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
// @since 3.18.0
type ChangeAnnotationsSupportOptions struct {
	// Whether the client groups edits with equal labels into tree nodes,
	// for instance all edits labelled with "Changes in Strings" would
	// be a tree node.
	GroupsOnLabel bool `json:"groupsOnLabel,omitempty"`
}
// Defines the capabilities provided by the client.
type ClientCapabilities struct {
	// Experimental client capabilities.
	Experimental any `json:"experimental,omitempty"`
	// General client capabilities.
	// 
	// @since 3.16.0
	General *GeneralClientCapabilities `json:"general,omitempty"`
	// Capabilities specific to the notebook document support.
	// 
	// @since 3.17.0
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitempty"`
	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`
	// Workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitempty"`
}
// @since 3.18.0
type ClientCodeActionKindOptions struct {
	// The code action kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []CodeActionKind `json:"valueSet"`
}
// @since 3.18.0
type ClientCodeActionLiteralOptions struct {
	// The code action kind is support with the following value
	// set.
	CodeActionKind ClientCodeActionKindOptions `json:"codeActionKind"`
}
// @since 3.18.0
type ClientCodeActionResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
// @since 3.18.0
type ClientCodeLensResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
// @since 3.18.0
type ClientCompletionItemInsertTextModeOptions struct {

	ValueSet []InsertTextMode `json:"valueSet"`
}
// @since 3.18.0
type ClientCompletionItemOptions struct {
	// Client supports commit characters on a completion item.
	CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`
	// Client supports the deprecated property on a completion item.
	DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`
	// Client supports the following content formats for the documentation
	// property. The order describes the preferred format of the client.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`
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
	InsertTextModeSupport *ClientCompletionItemInsertTextModeOptions `json:"insertTextModeSupport,omitempty"`
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
	ResolveSupport *ClientCompletionItemResolveOptions `json:"resolveSupport,omitempty"`
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
	TagSupport *CompletionItemTagOptions `json:"tagSupport,omitempty"`
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
	ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
}
// @since 3.18.0
type ClientCompletionItemResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
// @since 3.18.0
type ClientDiagnosticsTagOptions struct {
	// The tags supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet"`
}
// @since 3.18.0
type ClientFoldingRangeKindOptions struct {
	// The folding range kind values the client supports. When this
	// property exists the client also guarantees that it will
	// handle values outside its set gracefully and falls back
	// to a default value when unknown.
	ValueSet []FoldingRangeKind `json:"valueSet,omitempty"`
}
// @since 3.18.0
type ClientFoldingRangeOptions struct {
	// If set, the client signals that it supports setting collapsedText on
	// folding ranges to display custom labels instead of the default text.
	// 
	// @since 3.17.0
	CollapsedText bool `json:"collapsedText,omitempty"`
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
// @since 3.18.0
type ClientInlayHintResolveOptions struct {
	// The properties that a client can resolve lazily.
	Properties []string `json:"properties"`
}
// @since 3.18.0
type ClientSemanticTokensRequestFullDelta struct {
	// The client will send the `textDocument/semanticTokens/full/delta` request if
	// the server provides a corresponding handler.
	Delta bool `json:"delta,omitempty"`
}
// @since 3.18.0
type ClientSemanticTokensRequestOptions struct {
	// The client will send the `textDocument/semanticTokens/full` request if
	// the server provides a corresponding handler.
	Full *Or_ClientSemanticTokensRequestOptions_Full `json:"full,omitempty"`
	// The client will send the `textDocument/semanticTokens/range` request if
	// the server provides a corresponding handler.
	Range *Or_ClientSemanticTokensRequestOptions_Range `json:"range,omitempty"`
}
// @since 3.18.0
type ClientShowMessageActionItemOptions struct {
	// Whether the client supports additional attributes which
	// are preserved and send back to the server in the
	// request's response.
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
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
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`
	// The client supports the `activeParameter` property on
	// `SignatureHelp`/`SignatureInformation` being set to `null` to
	// indicate that no parameter should be active.
	// 
	// @since 3.18.0
	// @proposed
	NoActiveParameterSupport bool `json:"noActiveParameterSupport,omitempty"`
	// Client capabilities specific to parameter information.
	ParameterInformation *ClientSignatureParameterInformationOptions `json:"parameterInformation,omitempty"`
}
// @since 3.18.0
type ClientSignatureParameterInformationOptions struct {
	// The client supports processing label offsets instead of a
	// simple label string.
	// 
	// @since 3.14.0
	LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
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
	ValueSet []SymbolKind `json:"valueSet,omitempty"`
}
// @since 3.18.0
type ClientSymbolResolveOptions struct {
	// The properties that a client can resolve lazily. Usually
	// `location.range`
	Properties []string `json:"properties"`
}
// @since 3.18.0
type ClientSymbolTagOptions struct {
	// The tags supported by the client.
	ValueSet []SymbolTag `json:"valueSet"`
}
// A code action represents a change that can be performed in code, e.g. to fix a problem or
// to refactor code.
// 
// A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {
	// A command this code action executes. If a code action
	// provides an edit and a command, first the edit is
	// executed and then the command.
	Command *Command `json:"command,omitempty"`
	// A data entry field that is preserved on a code action between
	// a `textDocument/codeAction` and a `codeAction/resolve` request.
	// 
	// @since 3.16.0
	Data any `json:"data,omitempty"`
	// The diagnostics that this code action resolves.
	Diagnostics []Diagnostic `json:"diagnostics,omitempty"`
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
	Disabled *CodeActionDisabled `json:"disabled,omitempty"`
	// The workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`
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
	Kind *CodeActionKind `json:"kind,omitempty"`
	// Tags for this code action.
	// 
	// @since 3.18.0 - proposed
	Tags []CodeActionTag `json:"tags,omitempty"`
	// A short, human-readable, title for this code action.
	Title string `json:"title"`
}
// The Client Capabilities of a {@link CodeActionRequest}.
type CodeActionClientCapabilities struct {
	// The client support code action literals of type `CodeAction` as a valid
	// response of the `textDocument/codeAction` request. If the property is not
	// set the request can only return `Command` literals.
	// 
	// @since 3.8.0
	CodeActionLiteralSupport *ClientCodeActionLiteralOptions `json:"codeActionLiteralSupport,omitempty"`
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
	ResolveSupport *ClientCodeActionResolveOptions `json:"resolveSupport,omitempty"`
	// Client supports the tag property on a code action. Clients
	// supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.18.0 - proposed
	TagSupport *CodeActionTagOptions `json:"tagSupport,omitempty"`
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
	Only []CodeActionKind `json:"only,omitempty"`
	// The reason why code actions were requested.
	// 
	// @since 3.17.0
	TriggerKind *CodeActionTriggerKind `json:"triggerKind,omitempty"`
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
// Provider options for a {@link CodeActionRequest}.
type CodeActionOptions struct {
	// CodeActionKinds that this server may return.
	// 
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
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
	Documentation []CodeActionKindDocumentation `json:"documentation,omitempty"`
	// The server provides support to resolve additional
	// information for a code action.
	// 
	// @since 3.16.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// The parameters of a {@link CodeActionRequest}.
type CodeActionParams struct {
	// Context carrying additional information.
	Context CodeActionContext `json:"context"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The range for which the command was invoked.
	Range Range `json:"range"`
	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link CodeActionRequest}.
type CodeActionRegistrationOptions struct {
	// CodeActionKinds that this server may return.
	// 
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
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
	Documentation []CodeActionKindDocumentation `json:"documentation,omitempty"`
	// The server provides support to resolve additional
	// information for a code action.
	// 
	// @since 3.16.0
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// @since 3.18.0 - proposed
type CodeActionTagOptions struct {
	// The tags supported by the client.
	ValueSet []CodeActionTag `json:"valueSet"`
}
// Structure to capture a description for an error code.
// 
// @since 3.16.0
type CodeDescription struct {
	// An URI to open with more information about the diagnostic error.
	Href URI `json:"href"`
}
// A code lens represents a {@link Command command} that should be shown along with
// source text, like the number of references, a way to run tests, etc.
// 
// A code lens is _unresolved_ when no command is associated to it. For performance
// reasons the creation of a code lens and resolving should be done in two stages.
type CodeLens struct {
	// The command this code lens represents.
	Command *Command `json:"command,omitempty"`
	// A data entry field that is preserved on a code lens item between
	// a {@link CodeLensRequest} and a {@link CodeLensResolveRequest}
	Data any `json:"data,omitempty"`
	// The range in which this code lens is valid. Should only span a single line.
	Range Range `json:"range"`
}
// The client capabilities  of a {@link CodeLensRequest}.
type CodeLensClientCapabilities struct {
	// Whether code lens supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Whether the client supports resolving additional code lens
	// properties via a separate `codeLens/resolve` request.
	// 
	// @since 3.18.0
	ResolveSupport *ClientCodeLensResolveOptions `json:"resolveSupport,omitempty"`
}
// Code Lens provider options of a {@link CodeLensRequest}.
type CodeLensOptions struct {
	// Code lens has a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// The parameters of a {@link CodeLensRequest}.
type CodeLensParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The document to request code lens for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Represents a color range from a document.
type ColorInformation struct {
	// The actual color value for this color range.
	Color Color `json:"color"`
	// The range in the document where this color appears.
	Range Range `json:"range"`
}

type ColorPresentation struct {
	// An optional array of additional {@link TextEdit text edits} that are applied when
	// selecting this color presentation. Edits must not overlap with the main {@link ColorPresentation.textEdit edit} nor with themselves.
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
	// The label of this color presentation. It will be shown on the color
	// picker header. By default this is also the text that is inserted when selecting
	// this color presentation.
	Label string `json:"label"`
	// An {@link TextEdit edit} which is applied to a document when selecting
	// this presentation for the color.  When `falsy` the {@link ColorPresentation.label label}
	// is used.
	TextEdit *TextEdit `json:"textEdit,omitempty"`
}
// Parameters for a {@link ColorPresentationRequest}.
type ColorPresentationParams struct {
	// The color to request presentations for.
	Color Color `json:"color"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The range where the color would be inserted. Serves as a context.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Represents a reference to a command. Provides a title which
// will be used to represent a command in the UI and, optionally,
// an array of arguments which will be passed to the command handler
// function when invoked.
type Command struct {
	// Arguments that the command handler should be
	// invoked with.
	Arguments []any `json:"arguments,omitempty"`
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
// Completion client capabilities
type CompletionClientCapabilities struct {
	// The client supports the following `CompletionItem` specific
	// capabilities.
	CompletionItem *ClientCompletionItemOptions `json:"completionItem,omitempty"`

	CompletionItemKind *ClientCompletionItemOptionsKind `json:"completionItemKind,omitempty"`
	// The client supports the following `CompletionList` specific
	// capabilities.
	// 
	// @since 3.17.0
	CompletionList *CompletionListCapabilities `json:"completionList,omitempty"`
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
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`
}
// Contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
	// The trigger character (a single character) that has trigger code complete.
	// Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter string `json:"triggerCharacter,omitempty"`
	// How the completion was triggered.
	TriggerKind CompletionTriggerKind `json:"triggerKind"`
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
	AdditionalTextEdits []TextEdit `json:"additionalTextEdits,omitempty"`
	// An optional {@link Command command} that is executed *after* inserting this completion. *Note* that
	// additional modifications to the current document should be described with the
	// {@link CompletionItem.additionalTextEdits additionalTextEdits}-property.
	Command *Command `json:"command,omitempty"`
	// An optional set of characters that when pressed while this completion is active will accept it first and
	// then type that character. *Note* that all commit characters should have `length=1` and that superfluous
	// characters will be ignored.
	CommitCharacters []string `json:"commitCharacters,omitempty"`
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
	Documentation *Or_CompletionItem_Documentation `json:"documentation,omitempty"`
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
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`
	// How whitespace and indentation is handled during completion
	// item insertion. If not provided the clients default value depends on
	// the `textDocument.completion.insertTextMode` client capability.
	// 
	// @since 3.16.0
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`
	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor.
	Kind *CompletionItemKind `json:"kind,omitempty"`
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
	LabelDetails *CompletionItemLabelDetails `json:"labelDetails,omitempty"`
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
	Tags []CompletionItemTag `json:"tags,omitempty"`
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
	TextEdit *Or_CompletionItem_TextEdit `json:"textEdit,omitempty"`
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
	CommitCharacters *ApplyKind `json:"commitCharacters,omitempty"`
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
	Data *ApplyKind `json:"data,omitempty"`
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
	CommitCharacters []string `json:"commitCharacters,omitempty"`
	// A default data value.
	// 
	// @since 3.17.0
	Data any `json:"data,omitempty"`
	// A default edit range.
	// 
	// @since 3.17.0
	EditRange *Or_CompletionItemDefaults_EditRange `json:"editRange,omitempty"`
	// A default insert text format.
	// 
	// @since 3.17.0
	InsertTextFormat *InsertTextFormat `json:"insertTextFormat,omitempty"`
	// A default insert text mode.
	// 
	// @since 3.17.0
	InsertTextMode *InsertTextMode `json:"insertTextMode,omitempty"`
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
// @since 3.18.0
type CompletionItemTagOptions struct {
	// The tags supported by the client.
	ValueSet []CompletionItemTag `json:"valueSet"`
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
	ApplyKind *CompletionItemApplyKinds `json:"applyKind,omitempty"`
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
	ItemDefaults *CompletionItemDefaults `json:"itemDefaults,omitempty"`
	// The completion items.
	Items []CompletionItem `json:"items"`
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
	ItemDefaults []string `json:"itemDefaults,omitempty"`
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
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`
	// The server supports the following `CompletionItem` specific
	// capabilities.
	// 
	// @since 3.17.0
	CompletionItem *ServerCompletionItemOptions `json:"completionItem,omitempty"`
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
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Completion parameters
type CompletionParams struct {
	// The completion context. This is only available it the client specifies
	// to send this using the client capability `textDocument.completion.contextSupport === true`
	Context *CompletionContext `json:"context,omitempty"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`
	// The server supports the following `CompletionItem` specific
	// capabilities.
	// 
	// @since 3.17.0
	CompletionItem *ServerCompletionItemOptions `json:"completionItem,omitempty"`
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
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type ConfigurationItem struct {
	// The scope to get the configuration section for.
	ScopeUri *URI `json:"scopeUri,omitempty"`
	// The configuration section asked for.
	Section string `json:"section,omitempty"`
}
// The parameters of a configuration request.
type ConfigurationParams struct {

	Items []ConfigurationItem `json:"items"`
}
// Create file operation.
type CreateFile struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
	// A create
	Kind string `json:"kind"`
	// Additional options
	Options *CreateFileOptions `json:"options,omitempty"`
	// The resource to create.
	Uri DocumentUri `json:"uri"`
}
// Options to create a file.
type CreateFileOptions struct {
	// Ignore if exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
	// Overwrite existing file. Overwrite wins over `ignoreIfExists`
	Overwrite bool `json:"overwrite,omitempty"`
}
// The parameters sent in notifications/requests for user-initiated creation of
// files.
// 
// @since 3.16.0
type CreateFilesParams struct {
	// An array of all files/folders created in this operation.
	Files []FileCreate `json:"files"`
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

type DeclarationOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type DeclarationParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Client Capabilities for a {@link DefinitionRequest}.
type DefinitionClientCapabilities struct {
	// Whether definition supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// The client supports additional metadata in the form of definition links.
	// 
	// @since 3.14.0
	LinkSupport bool `json:"linkSupport,omitempty"`
}
// Server Capabilities for a {@link DefinitionRequest}.
type DefinitionOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link DefinitionRequest}.
type DefinitionParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link DefinitionRequest}.
type DefinitionRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Delete file operation
type DeleteFile struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
	// A delete
	Kind string `json:"kind"`
	// Delete options.
	Options *DeleteFileOptions `json:"options,omitempty"`
	// The file to delete.
	Uri DocumentUri `json:"uri"`
}
// Delete file options
type DeleteFileOptions struct {
	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists bool `json:"ignoreIfNotExists,omitempty"`
	// Delete the content recursively if a folder is denoted.
	Recursive bool `json:"recursive,omitempty"`
}
// The parameters sent in notifications/requests for user-initiated deletes of
// files.
// 
// @since 3.16.0
type DeleteFilesParams struct {
	// An array of all files/folders deleted in this operation.
	Files []FileDelete `json:"files"`
}
// Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
// are only valid in the scope of a resource.
type Diagnostic struct {
	// The diagnostic's code, which usually appear in the user interface.
	Code *Or_Diagnostic_Code `json:"code,omitempty"`
	// An optional property to describe the error code.
	// Requires the code field (above) to be present/not null.
	// 
	// @since 3.16.0
	CodeDescription *CodeDescription `json:"codeDescription,omitempty"`
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
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
	// The diagnostic's severity. To avoid interpretation mismatches when a
	// server is used with different clients it is highly recommended that servers
	// always provide a severity value.
	Severity *DiagnosticSeverity `json:"severity,omitempty"`
	// A human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'. It usually
	// appears in the user interface.
	Source string `json:"source,omitempty"`
	// Additional metadata about the diagnostic.
	// 
	// @since 3.15.0
	Tags []DiagnosticTag `json:"tags,omitempty"`
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
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitempty"`
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
// Represents a related message and source code location for a diagnostic. This should be
// used to point to code locations that cause or related to a diagnostics, e.g when duplicating
// a symbol in a scope.
type DiagnosticRelatedInformation struct {
	// The location of this related diagnostic information.
	Location Location `json:"location"`
	// The message of this related diagnostic information.
	Message string `json:"message"`
}
// Cancellation data returned from a diagnostic request.
// 
// @since 3.17.0
type DiagnosticServerCancellationData struct {

	RetriggerRequest bool `json:"retriggerRequest"`
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
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitempty"`
}

type DidChangeConfigurationClientCapabilities struct {
	// Did change configuration notification supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// The parameters of a change configuration notification.
type DidChangeConfigurationParams struct {
	// The actual changed settings
	Settings any `json:"settings"`
}

type DidChangeConfigurationRegistrationOptions struct {

	Section *Or_DidChangeConfigurationRegistrationOptions_Section `json:"section,omitempty"`
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
// The watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {
	// The actual file events.
	Changes []FileEvent `json:"changes"`
}
// Describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {
	// The watchers to register.
	Watchers []FileSystemWatcher `json:"watchers"`
}
// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent `json:"event"`
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
// The parameters sent in a close text document notification
type DidCloseTextDocumentParams struct {
	// The document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
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
// The parameters sent in an open text document notification
type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}
// The params sent in a save notebook document notification.
// 
// @since 3.17.0
type DidSaveNotebookDocumentParams struct {
	// The notebook document that got saved.
	NotebookDocument NotebookDocumentIdentifier `json:"notebookDocument"`
}
// The parameters sent in a save text document notification
type DidSaveTextDocumentParams struct {
	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`
	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type DocumentColorClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `DocumentColorRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type DocumentColorOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link DocumentColorRequest}.
type DocumentColorParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Parameters of the document diagnostic request.
// 
// @since 3.17.0
type DocumentDiagnosticParams struct {
	// The additional identifier  provided during registration.
	Identifier string `json:"identifier,omitempty"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The result id of a previous response if provided.
	PreviousResultId string `json:"previousResultId,omitempty"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// A partial result for a document diagnostic report.
// 
// @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {

	RelatedDocuments map[DocumentUri]Or_DocumentDiagnosticReportPartialResult_RelatedDocuments `json:"relatedDocuments"`
}
// Client capabilities of a {@link DocumentFormattingRequest}.
type DocumentFormattingClientCapabilities struct {
	// Whether formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Provider options for a {@link DocumentFormattingRequest}.
type DocumentFormattingOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// The parameters of a {@link DocumentFormattingRequest}.
type DocumentFormattingParams struct {
	// The format options.
	Options FormattingOptions `json:"options"`
	// The document to format.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link DocumentFormattingRequest}.
type DocumentFormattingRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// A document highlight is a range inside a text document which deserves
// special attention. Usually a document highlight is visualized by changing
// the background color of its range.
type DocumentHighlight struct {
	// The highlight kind, default is {@link DocumentHighlightKind.Text text}.
	Kind *DocumentHighlightKind `json:"kind,omitempty"`
	// The range this highlight applies to.
	Range Range `json:"range"`
}
// Client Capabilities for a {@link DocumentHighlightRequest}.
type DocumentHighlightClientCapabilities struct {
	// Whether document highlight supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Provider options for a {@link DocumentHighlightRequest}.
type DocumentHighlightOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link DocumentHighlightRequest}.
type DocumentHighlightParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link DocumentHighlightRequest}.
type DocumentHighlightRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	Target *URI `json:"target,omitempty"`
	// The tooltip text when you hover over this link.
	// 
	// If a tooltip is provided, is will be displayed in a string that includes instructions on how to
	// trigger the link, such as `{0} (ctrl + click)`. The specific instructions vary depending on OS,
	// user settings, and localization.
	// 
	// @since 3.15.0
	Tooltip string `json:"tooltip,omitempty"`
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
// Provider options for a {@link DocumentLinkRequest}.
type DocumentLinkOptions struct {
	// Document links have a resolve provider as well.
	ResolveProvider bool `json:"resolveProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// The parameters of a {@link DocumentLinkRequest}.
type DocumentLinkParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The document to provide document links for.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Client capabilities of a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingClientCapabilities struct {
	// Whether on type formatting supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Provider options for a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingOptions struct {
	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`
	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
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
// Registration options for a {@link DocumentOnTypeFormattingRequest}.
type DocumentOnTypeFormattingRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`
	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
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
// Provider options for a {@link DocumentRangeFormattingRequest}.
type DocumentRangeFormattingOptions struct {
	// Whether the server supports formatting multiple ranges at once.
	// 
	// @since 3.18.0
	// @proposed
	RangesSupport bool `json:"rangesSupport,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they
// have two ranges: one that encloses its definition and one that points to
// its most interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {
	// Children of this symbol, e.g. properties of a class.
	Children []DocumentSymbol `json:"children,omitempty"`
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
	Tags []SymbolTag `json:"tags,omitempty"`
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
	SymbolKind *ClientSymbolKindOptions `json:"symbolKind,omitempty"`
	// The client supports tags on `SymbolInformation`. Tags are supported on
	// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.16.0
	TagSupport *ClientSymbolTagOptions `json:"tagSupport,omitempty"`
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
// Parameters for a {@link DocumentSymbolRequest}.
type DocumentSymbolParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Edit range variant that includes ranges for insert and replace operations.
// 
// @since 3.18.0
type EditRangeWithInsertReplace struct {

	Insert Range `json:"insert"`

	Replace Range `json:"replace"`
}
// The client capabilities of a {@link ExecuteCommandRequest}.
type ExecuteCommandClientCapabilities struct {
	// Execute command supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// The server capabilities of a {@link ExecuteCommandRequest}.
type ExecuteCommandOptions struct {
	// The commands to be executed on the server
	Commands []string `json:"commands"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// The parameters of a {@link ExecuteCommandRequest}.
type ExecuteCommandParams struct {
	// Arguments that the command should be invoked with.
	Arguments []any `json:"arguments,omitempty"`
	// The identifier of the actual command handler.
	Command string `json:"command"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link ExecuteCommandRequest}.
type ExecuteCommandRegistrationOptions struct {
	// The commands to be executed on the server
	Commands []string `json:"commands"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
// Represents information on a file/folder create.
// 
// @since 3.16.0
type FileCreate struct {
	// A file:// URI for the location of the file/folder being created.
	Uri string `json:"uri"`
}
// Represents information on a file/folder delete.
// 
// @since 3.16.0
type FileDelete struct {
	// A file:// URI for the location of the file/folder being deleted.
	Uri string `json:"uri"`
}
// An event describing a file change.
type FileEvent struct {
	// The change type.
	Type FileChangeType `json:"type"`
	// The file's uri.
	Uri DocumentUri `json:"uri"`
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
// Options for notifications/requests for user operations on files.
// 
// @since 3.16.0
type FileOperationOptions struct {
	// The server is interested in receiving didCreateFiles notifications.
	DidCreate *FileOperationRegistrationOptions `json:"didCreate,omitempty"`
	// The server is interested in receiving didDeleteFiles file notifications.
	DidDelete *FileOperationRegistrationOptions `json:"didDelete,omitempty"`
	// The server is interested in receiving didRenameFiles notifications.
	DidRename *FileOperationRegistrationOptions `json:"didRename,omitempty"`
	// The server is interested in receiving willCreateFiles requests.
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`
	// The server is interested in receiving willDeleteFiles file requests.
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
	// The server is interested in receiving willRenameFiles requests.
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`
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
	Matches *FileOperationPatternKind `json:"matches,omitempty"`
	// Additional options used during matching.
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}
// Matching options for the file operation pattern.
// 
// @since 3.16.0
type FileOperationPatternOptions struct {
	// The pattern should be matched ignoring casing.
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}
// The options to register for file operations.
// 
// @since 3.16.0
type FileOperationRegistrationOptions struct {
	// The actual filters.
	Filters []FileOperationFilter `json:"filters"`
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

type FileSystemWatcher struct {
	// The glob pattern to watch. See {@link GlobPattern glob pattern} for more detail.
	// 
	// @since 3.17.0 support for relative patterns.
	GlobPattern GlobPattern `json:"globPattern"`
	// The kind of events of interest. If omitted it defaults
	// to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	// which is 7.
	Kind *WatchKind `json:"kind,omitempty"`
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
	Kind *FoldingRangeKind `json:"kind,omitempty"`
	// The zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	StartCharacter uint32 `json:"startCharacter,omitempty"`
	// The zero-based start line of the range to fold. The folded area starts after the line's last character.
	// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	StartLine uint32 `json:"startLine"`
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
	FoldingRange *ClientFoldingRangeOptions `json:"foldingRange,omitempty"`
	// Specific options for the folding range kind.
	// 
	// @since 3.17.0
	FoldingRangeKind *ClientFoldingRangeKindOptions `json:"foldingRangeKind,omitempty"`
	// If set, the client signals that it only supports folding complete lines.
	// If set, client will ignore specified `startCharacter` and `endCharacter`
	// properties in a FoldingRange.
	LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
	// The maximum number of folding ranges that the client prefers to receive
	// per document. The value serves as a hint, servers are free to follow the
	// limit.
	RangeLimit uint32 `json:"rangeLimit,omitempty"`
}

type FoldingRangeOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link FoldingRangeRequest}.
type FoldingRangeParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// General client capabilities.
// 
// @since 3.16.0
type GeneralClientCapabilities struct {
	// Client capabilities specific to the client's markdown parser.
	// 
	// @since 3.16.0
	Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`
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
	PositionEncodings []PositionEncodingKind `json:"positionEncodings,omitempty"`
	// Client capabilities specific to regular expressions.
	// 
	// @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`
	// Client capability that signals how the client
	// handles stale requests (e.g. a request
	// for which the client will not process the response
	// anymore since the information is outdated).
	// 
	// @since 3.17.0
	StaleRequestSupport *StaleRequestSupportOptions `json:"staleRequestSupport,omitempty"`
}
// The result of a hover request.
type Hover struct {
	// The hover's content
	Contents Or_Hover_Contents `json:"contents"`
	// An optional range inside the text document that is used to
	// visualize the hover, e.g. by changing the background color.
	Range *Range `json:"range,omitempty"`
}

type HoverClientCapabilities struct {
	// Client supports the following content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
	// Whether hover supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Hover options.
type HoverOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link HoverRequest}.
type HoverParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link HoverRequest}.
type HoverRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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

type ImplementationOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type ImplementationParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// The data type of the ResponseError if the
// initialize request fails.
type InitializeError struct {
	// Indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool `json:"retry"`
}

type InitializeParams struct {
	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`
	// Information about the client
	// 
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`
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
	RootPath **string `json:"rootPath,omitempty"`
	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	// 
	// @deprecated in favour of workspaceFolders.
	RootUri *DocumentUri `json:"rootUri"`
	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitempty"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
	// The workspace folders configured in the client when the server starts.
	// 
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	// 
	// @since 3.6.0
	WorkspaceFolders *Or_InitializeParams_WorkspaceFolders `json:"workspaceFolders,omitempty"`
}
// The result returned from an initialize request.
type InitializeResult struct {
	// The capabilities the language server provides.
	Capabilities ServerCapabilities `json:"capabilities"`
	// Information about the server.
	// 
	// @since 3.15.0
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`
}

type InitializedParams struct {
}
// Inlay hint information.
// 
// @since 3.17.0
type InlayHint struct {
	// A data entry field that is preserved on an inlay hint between
	// a `textDocument/inlayHint` and a `inlayHint/resolve` request.
	Data any `json:"data,omitempty"`
	// The kind of this hint. Can be omitted in which case the client
	// should fall back to a reasonable default.
	Kind *InlayHintKind `json:"kind,omitempty"`
	// The label of this hint. A human readable string or an array of
	// InlayHintLabelPart label parts.
	// 
	// *Note* that neither the string nor the label part can be empty.
	Label Or_InlayHint_Label `json:"label"`
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
	TextEdits []TextEdit `json:"textEdits,omitempty"`
	// The tooltip text when you hover over this item.
	Tooltip *Or_InlayHint_Tooltip `json:"tooltip,omitempty"`
}
// Inlay hint client capabilities.
// 
// @since 3.17.0
type InlayHintClientCapabilities struct {
	// Whether inlay hints support dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Indicates which properties a client can resolve lazily on an inlay
	// hint.
	ResolveSupport *ClientInlayHintResolveOptions `json:"resolveSupport,omitempty"`
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
	Command *Command `json:"command,omitempty"`
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
	Location *Location `json:"location,omitempty"`
	// The tooltip text when you hover over this label part. Depending on
	// the client capability `inlayHint.resolveSupport` clients might resolve
	// this property late using the resolve request.
	Tooltip *Or_InlayHintLabelPart_Tooltip `json:"tooltip,omitempty"`
	// The value of this label part.
	Value string `json:"value"`
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
// A parameter literal used in inlay hint requests.
// 
// @since 3.17.0
type InlayHintParams struct {
	// The document range for which inlay hints should be computed.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Client capabilities specific to inline completions.
// 
// @since 3.18.0
// @proposed
type InlineCompletionClientCapabilities struct {
	// Whether implementation supports dynamic registration for inline completion providers.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Provides information about the context in which an inline completion was requested.
// 
// @since 3.18.0
// @proposed
type InlineCompletionContext struct {
	// Provides information about the currently selected item in the autocomplete widget if it is visible.
	SelectedCompletionInfo *SelectedCompletionInfo `json:"selectedCompletionInfo,omitempty"`
	// Describes how the inline completion was triggered.
	TriggerKind InlineCompletionTriggerKind `json:"triggerKind"`
}
// An inline completion item represents a text snippet that is proposed inline to complete text that is being typed.
// 
// @since 3.18.0
// @proposed
type InlineCompletionItem struct {
	// An optional {@link Command} that is executed *after* inserting this completion.
	Command *Command `json:"command,omitempty"`
	// A text that is used to decide if this inline completion should be shown. When `falsy` the {@link InlineCompletionItem.insertText} is used.
	FilterText string `json:"filterText,omitempty"`
	// The text to replace the range with. Must be set.
	InsertText Or_InlineCompletionItem_InsertText `json:"insertText"`
	// The range to replace. Must begin and end on the same line.
	Range *Range `json:"range,omitempty"`
}
// Represents a collection of {@link InlineCompletionItem inline completion items} to be presented in the editor.
// 
// @since 3.18.0
// @proposed
type InlineCompletionList struct {
	// The inline completion items
	Items []InlineCompletionItem `json:"items"`
}
// Inline completion options used during static registration.
// 
// @since 3.18.0
// @proposed
type InlineCompletionOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Client capabilities specific to inline values.
// 
// @since 3.17.0
type InlineValueClientCapabilities struct {
	// Whether implementation supports dynamic registration for inline value providers.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// @since 3.17.0
type InlineValueContext struct {
	// The stack frame (as a DAP Id) where the execution has stopped.
	FrameId int32 `json:"frameId"`
	// The document range where execution has stopped.
	// Typically the end position of the range denotes the line where the inline values are shown.
	StoppedLocation Range `json:"stoppedLocation"`
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
// Inline value options used during static registration.
// 
// @since 3.17.0
type InlineValueOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Provide inline value as text.
// 
// @since 3.17.0
type InlineValueText struct {
	// The document range for which the inline value applies.
	Range Range `json:"range"`
	// The text of the inline value.
	Text string `json:"text"`
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
// Client capabilities for the linked editing range request.
// 
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type LinkedEditingRangeOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type LinkedEditingRangeParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// Represents a location inside a resource, such as a line
// inside a text file.
type Location struct {

	Range Range `json:"range"`

	Uri DocumentUri `json:"uri"`
}
// Represents the connection of two locations. Provides additional metadata over normal {@link Location locations},
// including an origin range.
type LocationLink struct {
	// Span of the origin of this link.
	// 
	// Used as the underlined span for mouse interaction. Defaults to the word range at
	// the definition position.
	OriginSelectionRange *Range `json:"originSelectionRange,omitempty"`
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
// Location with only uri and does not include range.
// 
// @since 3.18.0
type LocationUriOnly struct {

	Uri DocumentUri `json:"uri"`
}
// The log message parameters.
type LogMessageParams struct {
	// The actual message.
	Message string `json:"message"`
	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`
}

type LogTraceParams struct {

	Message string `json:"message"`

	Verbose string `json:"verbose,omitempty"`
}
// Client capabilities specific to the used markdown parser.
// 
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// A list of HTML tags that the client allows / supports in
	// Markdown.
	// 
	// @since 3.17.0
	AllowedTags []string `json:"allowedTags,omitempty"`
	// The name of the parser.
	Parser string `json:"parser"`
	// The version of the parser.
	Version string `json:"version,omitempty"`
}
// @since 3.18.0
// @deprecated use MarkupContent instead.
type MarkedStringWithLanguage struct {

	Language string `json:"language"`

	Value string `json:"value"`
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

type MessageActionItem struct {
	// A short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}
// Moniker definition to match LSIF 0.5 moniker definition.
// 
// @since 3.16.0
type Moniker struct {
	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier string `json:"identifier"`
	// The moniker kind if known.
	Kind *MonikerKind `json:"kind,omitempty"`
	// The scheme of the moniker. For example tsc or .Net
	Scheme string `json:"scheme"`
	// The scope in which the moniker is unique
	Unique UniquenessLevel `json:"unique"`
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

type MonikerOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type MonikerParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

type MonikerRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	ExecutionSummary *ExecutionSummary `json:"executionSummary,omitempty"`
	// The cell's kind
	Kind NotebookCellKind `json:"kind"`
	// Additional metadata stored with the cell.
	// 
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata,omitempty"`
}
// A change describing how to move a `NotebookCell`
// array from state S to S'.
// 
// @since 3.17.0
type NotebookCellArrayChange struct {
	// The new cells, if any
	Cells []NotebookCell `json:"cells,omitempty"`
	// The deleted cells
	DeleteCount uint32 `json:"deleteCount"`
	// The start oftest of the cell that changed.
	Start uint32 `json:"start"`
}
// @since 3.18.0
type NotebookCellLanguage struct {

	Language string `json:"language"`
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
	Notebook Or_NotebookCellTextDocumentFilter_Notebook `json:"notebook"`
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
	Metadata *LSPObject `json:"metadata,omitempty"`
	// The type of the notebook.
	NotebookType string `json:"notebookType"`
	// The notebook document's uri.
	Uri URI `json:"uri"`
	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int32 `json:"version"`
}
// Structural changes to cells in a notebook document.
// 
// @since 3.18.0
type NotebookDocumentCellChangeStructure struct {
	// The change to the cell array.
	Array NotebookCellArrayChange `json:"array"`
	// Additional closed cell text documents.
	DidClose []TextDocumentIdentifier `json:"didClose,omitempty"`
	// Additional opened cell text documents.
	DidOpen []TextDocumentItem `json:"didOpen,omitempty"`
}
// Cell changes to a notebook document.
// 
// @since 3.18.0
type NotebookDocumentCellChanges struct {
	// Changes to notebook cells properties like its
	// kind, execution summary or metadata.
	Data []NotebookCell `json:"data,omitempty"`
	// Changes to the cell structure to add or
	// remove cells.
	Structure *NotebookDocumentCellChangeStructure `json:"structure,omitempty"`
	// Changes to the text content of notebook cells.
	TextContent []NotebookDocumentCellContentChanges `json:"textContent,omitempty"`
}
// Content changes to a cell in a notebook document.
// 
// @since 3.18.0
type NotebookDocumentCellContentChanges struct {

	Changes []TextDocumentContentChangeEvent `json:"changes"`

	Document VersionedTextDocumentIdentifier `json:"document"`
}
// A change event for a notebook document.
// 
// @since 3.17.0
type NotebookDocumentChangeEvent struct {
	// Changes to cells
	Cells *NotebookDocumentCellChanges `json:"cells,omitempty"`
	// The changed meta data if any.
	// 
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata *LSPObject `json:"metadata,omitempty"`
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
// A notebook document filter where `notebookType` is required field.
// 
// @since 3.18.0
type NotebookDocumentFilterNotebookType struct {
	// The type of the enclosing notebook.
	NotebookType string `json:"notebookType"`
	// A glob pattern.
	Pattern *GlobPattern `json:"pattern,omitempty"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
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
// A notebook document filter where `scheme` is required field.
// 
// @since 3.18.0
type NotebookDocumentFilterScheme struct {
	// The type of the enclosing notebook.
	NotebookType string `json:"notebookType,omitempty"`
	// A glob pattern.
	Pattern *GlobPattern `json:"pattern,omitempty"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme"`
}
// @since 3.18.0
type NotebookDocumentFilterWithCells struct {
	// The cells of the matching notebook to be synced.
	Cells []NotebookCellLanguage `json:"cells"`
	// The notebook to be synced If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	Notebook *Or_NotebookDocumentFilterWithCells_Notebook `json:"notebook,omitempty"`
}
// @since 3.18.0
type NotebookDocumentFilterWithNotebook struct {
	// The cells of the matching notebook to be synced.
	Cells []NotebookCellLanguage `json:"cells,omitempty"`
	// The notebook to be synced If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	Notebook Or_NotebookDocumentFilterWithNotebook_Notebook `json:"notebook"`
}
// A literal to identify a notebook document in the client.
// 
// @since 3.17.0
type NotebookDocumentIdentifier struct {
	// The notebook document's uri.
	Uri URI `json:"uri"`
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
	NotebookSelector []Or_NotebookDocumentSyncOptions_NotebookSelector `json:"notebookSelector"`
	// Whether save notification should be forwarded to
	// the server. Will only be honored if mode === `notebook`.
	Save bool `json:"save,omitempty"`
}
// Registration options specific to a notebook.
// 
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// The notebooks to be synced
	NotebookSelector []Or_NotebookDocumentSyncRegistrationOptions_NotebookSelector `json:"notebookSelector"`
	// Whether save notification should be forwarded to
	// the server. Will only be honored if mode === `notebook`.
	Save bool `json:"save,omitempty"`
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
// Represents a parameter of a callable-signature. A parameter can
// have a label and a doc-comment.
type ParameterInformation struct {
	// The human-readable doc-comment of this parameter. Will be shown
	// in the UI but can be omitted.
	Documentation *Or_ParameterInformation_Documentation `json:"documentation,omitempty"`
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
	Label Or_ParameterInformation_Label `json:"label"`
}

type PartialResultParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
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
// @since 3.18.0
type PrepareRenameDefaultBehavior struct {

	DefaultBehavior bool `json:"defaultBehavior"`
}

type PrepareRenameParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// @since 3.18.0
type PrepareRenamePlaceholder struct {

	Placeholder string `json:"placeholder"`

	Range Range `json:"range"`
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

type ProgressParams struct {
	// The progress token provided by the client or server.
	Token ProgressToken `json:"token"`
	// The progress data.
	Value any `json:"value"`
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
	TagSupport *ClientDiagnosticsTagOptions `json:"tagSupport,omitempty"`
	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.
	// 
	// @since 3.15.0
	VersionSupport bool `json:"versionSupport,omitempty"`
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
// Client Capabilities for a {@link ReferencesRequest}.
type ReferenceClientCapabilities struct {
	// Whether references supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Value-object that contains additional information when
// requesting references.
type ReferenceContext struct {
	// Include the declaration of the current symbol.
	IncludeDeclaration bool `json:"includeDeclaration"`
}
// Reference options.
type ReferenceOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link ReferencesRequest}.
type ReferenceParams struct {

	Context ReferenceContext `json:"context"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Registration options for a {@link ReferencesRequest}.
type ReferenceRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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

type RegistrationParams struct {

	Registrations []Registration `json:"registrations"`
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
	RelatedDocuments map[DocumentUri]Or_RelatedFullDocumentDiagnosticReport_RelatedDocuments `json:"relatedDocuments,omitempty"`
	// An optional result id. If provided it will
	// be sent on the next diagnostic request for the
	// same document.
	ResultId string `json:"resultId,omitempty"`
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
	RelatedDocuments map[DocumentUri]Or_RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments `json:"relatedDocuments,omitempty"`
	// A result id which will be sent on the next
	// diagnostic request for the same document.
	ResultId string `json:"resultId"`
}
// A relative pattern is a helper to construct glob patterns that are matched
// relatively to a base URI. The common value for a `baseUri` is a workspace
// folder root, but it can be another absolute URI as well.
// 
// @since 3.17.0
type RelativePattern struct {
	// A workspace folder or a base URI to which this pattern will be matched
	// against relatively.
	BaseUri Or_RelativePattern_BaseUri `json:"baseUri"`
	// The actual glob pattern;
	Pattern Pattern `json:"pattern"`
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
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior,omitempty"`
}
// Rename file operation
type RenameFile struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
	// A rename
	Kind string `json:"kind"`
	// The new location.
	NewUri DocumentUri `json:"newUri"`
	// The old (existing) location.
	OldUri DocumentUri `json:"oldUri"`
	// Rename options.
	Options *RenameFileOptions `json:"options,omitempty"`
}
// Rename file options
type RenameFileOptions struct {
	// Ignores if target exists.
	IgnoreIfExists bool `json:"ignoreIfExists,omitempty"`
	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	Overwrite bool `json:"overwrite,omitempty"`
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
// Provider options for a {@link RenameRequest}.
type RenameOptions struct {
	// Renames should be checked and tested before being executed.
	// 
	// @since version 3.12.0
	PrepareProvider bool `json:"prepareProvider,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// A generic resource operation.
type ResourceOperation struct {
	// An optional annotation identifier describing the operation.
	// 
	// @since 3.16.0
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
	// The resource operation kind.
	Kind string `json:"kind"`
}
// Save options.
type SaveOptions struct {
	// The client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
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
// A selection range represents a part of a selection hierarchy. A selection range
// may have a parent selection range that contains it.
type SelectionRange struct {
	// The parent selection range containing this range. Therefore `parent.range` must contain `this.range`.
	Parent *SelectionRange `json:"parent,omitempty"`
	// The {@link Range range} of this selection range.
	Range Range `json:"range"`
}

type SelectionRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for selection range providers. If this is set to `true`
	// the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server
	// capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type SelectionRangeOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// A parameter literal used in selection range requests.
type SelectionRangeParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The positions inside the text document.
	Positions []Position `json:"positions"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// @since 3.16.0
type SemanticTokensDelta struct {
	// The semantic token edits to transform a previous result into a new result.
	Edits []SemanticTokensEdit `json:"edits"`

	ResultId string `json:"resultId,omitempty"`
}
// @since 3.16.0
type SemanticTokensDeltaParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The result id of a previous response. The result Id can either point to a full response
	// or a delta response depending on what was received last.
	PreviousResultId string `json:"previousResultId"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {

	Edits []SemanticTokensEdit `json:"edits"`
}
// @since 3.16.0
type SemanticTokensEdit struct {
	// The elements to insert.
	Data []uint32 `json:"data,omitempty"`
	// The count of elements to remove.
	DeleteCount uint32 `json:"deleteCount"`
	// The start offset of the edit.
	Start uint32 `json:"start"`
}
// Semantic tokens options to support deltas for full documents
// 
// @since 3.18.0
type SemanticTokensFullDelta struct {
	// The server supports deltas for full documents.
	Delta bool `json:"delta,omitempty"`
}
// @since 3.16.0
type SemanticTokensLegend struct {
	// The token modifiers a server uses.
	TokenModifiers []string `json:"tokenModifiers"`
	// The token types a server uses.
	TokenTypes []string `json:"tokenTypes"`
}
// @since 3.16.0
type SemanticTokensOptions struct {
	// Server supports providing semantic tokens for a full document.
	Full *Or_SemanticTokensOptions_Full `json:"full,omitempty"`
	// The legend used by the server
	Legend SemanticTokensLegend `json:"legend"`
	// Server supports providing semantic tokens for a specific range
	// of a document.
	Range *Or_SemanticTokensOptions_Range `json:"range,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// @since 3.16.0
type SemanticTokensParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// @since 3.16.0
type SemanticTokensPartialResult struct {

	Data []uint32 `json:"data"`
}
// @since 3.16.0
type SemanticTokensRangeParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The range the semantic tokens are requested for.
	Range Range `json:"range"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// @since 3.16.0
type SemanticTokensRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// Server supports providing semantic tokens for a full document.
	Full *Or_SemanticTokensRegistrationOptions_Full `json:"full,omitempty"`
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
	// The legend used by the server
	Legend SemanticTokensLegend `json:"legend"`
	// Server supports providing semantic tokens for a specific range
	// of a document.
	Range *Or_SemanticTokensRegistrationOptions_Range `json:"range,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
// Defines the capabilities provided by a language
// server.
type ServerCapabilities struct {
	// The server provides call hierarchy support.
	// 
	// @since 3.16.0
	CallHierarchyProvider *Or_ServerCapabilities_CallHierarchyProvider `json:"callHierarchyProvider,omitempty"`
	// The server provides code actions. CodeActionOptions may only be
	// specified if the client states that it supports
	// `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider *Or_ServerCapabilities_CodeActionProvider `json:"codeActionProvider,omitempty"`
	// The server provides code lens.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`
	// The server provides color provider support.
	ColorProvider *Or_ServerCapabilities_ColorProvider `json:"colorProvider,omitempty"`
	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`
	// The server provides Goto Declaration support.
	DeclarationProvider *Or_ServerCapabilities_DeclarationProvider `json:"declarationProvider,omitempty"`
	// The server provides goto definition support.
	DefinitionProvider *Or_ServerCapabilities_DefinitionProvider `json:"definitionProvider,omitempty"`
	// The server has support for pull model diagnostics.
	// 
	// @since 3.17.0
	DiagnosticProvider *Or_ServerCapabilities_DiagnosticProvider `json:"diagnosticProvider,omitempty"`
	// The server provides document formatting.
	DocumentFormattingProvider *Or_ServerCapabilities_DocumentFormattingProvider `json:"documentFormattingProvider,omitempty"`
	// The server provides document highlight support.
	DocumentHighlightProvider *Or_ServerCapabilities_DocumentHighlightProvider `json:"documentHighlightProvider,omitempty"`
	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`
	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
	// The server provides document range formatting.
	DocumentRangeFormattingProvider *Or_ServerCapabilities_DocumentRangeFormattingProvider `json:"documentRangeFormattingProvider,omitempty"`
	// The server provides document symbol support.
	DocumentSymbolProvider *Or_ServerCapabilities_DocumentSymbolProvider `json:"documentSymbolProvider,omitempty"`
	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`
	// Experimental server capabilities.
	Experimental any `json:"experimental,omitempty"`
	// The server provides folding provider support.
	FoldingRangeProvider *Or_ServerCapabilities_FoldingRangeProvider `json:"foldingRangeProvider,omitempty"`
	// The server provides hover support.
	HoverProvider *Or_ServerCapabilities_HoverProvider `json:"hoverProvider,omitempty"`
	// The server provides Goto Implementation support.
	ImplementationProvider *Or_ServerCapabilities_ImplementationProvider `json:"implementationProvider,omitempty"`
	// The server provides inlay hints.
	// 
	// @since 3.17.0
	InlayHintProvider *Or_ServerCapabilities_InlayHintProvider `json:"inlayHintProvider,omitempty"`
	// Inline completion options used during static registration.
	// 
	// @since 3.18.0
	// @proposed
	InlineCompletionProvider *Or_ServerCapabilities_InlineCompletionProvider `json:"inlineCompletionProvider,omitempty"`
	// The server provides inline values.
	// 
	// @since 3.17.0
	InlineValueProvider *Or_ServerCapabilities_InlineValueProvider `json:"inlineValueProvider,omitempty"`
	// The server provides linked editing range support.
	// 
	// @since 3.16.0
	LinkedEditingRangeProvider *Or_ServerCapabilities_LinkedEditingRangeProvider `json:"linkedEditingRangeProvider,omitempty"`
	// The server provides moniker support.
	// 
	// @since 3.16.0
	MonikerProvider *Or_ServerCapabilities_MonikerProvider `json:"monikerProvider,omitempty"`
	// Defines how notebook documents are synced.
	// 
	// @since 3.17.0
	NotebookDocumentSync *Or_ServerCapabilities_NotebookDocumentSync `json:"notebookDocumentSync,omitempty"`
	// The position encoding the server picked from the encodings offered
	// by the client via the client capability `general.positionEncodings`.
	// 
	// If the client didn't provide any position encodings the only valid
	// value that a server can return is 'utf-16'.
	// 
	// If omitted it defaults to 'utf-16'.
	// 
	// @since 3.17.0
	PositionEncoding *PositionEncodingKind `json:"positionEncoding,omitempty"`
	// The server provides find references support.
	ReferencesProvider *Or_ServerCapabilities_ReferencesProvider `json:"referencesProvider,omitempty"`
	// The server provides rename support. RenameOptions may only be
	// specified if the client states that it supports
	// `prepareSupport` in its initial `initialize` request.
	RenameProvider *Or_ServerCapabilities_RenameProvider `json:"renameProvider,omitempty"`
	// The server provides selection range support.
	SelectionRangeProvider *Or_ServerCapabilities_SelectionRangeProvider `json:"selectionRangeProvider,omitempty"`
	// The server provides semantic tokens support.
	// 
	// @since 3.16.0
	SemanticTokensProvider *Or_ServerCapabilities_SemanticTokensProvider `json:"semanticTokensProvider,omitempty"`
	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`
	// Defines how text documents are synced. Is either a detailed structure
	// defining each notification or for backwards compatibility the
	// TextDocumentSyncKind number.
	TextDocumentSync *Or_ServerCapabilities_TextDocumentSync `json:"textDocumentSync,omitempty"`
	// The server provides Goto Type Definition support.
	TypeDefinitionProvider *Or_ServerCapabilities_TypeDefinitionProvider `json:"typeDefinitionProvider,omitempty"`
	// The server provides type hierarchy support.
	// 
	// @since 3.17.0
	TypeHierarchyProvider *Or_ServerCapabilities_TypeHierarchyProvider `json:"typeHierarchyProvider,omitempty"`
	// Workspace specific server capabilities.
	Workspace *WorkspaceOptions `json:"workspace,omitempty"`
	// The server provides workspace symbol support.
	WorkspaceSymbolProvider *Or_ServerCapabilities_WorkspaceSymbolProvider `json:"workspaceSymbolProvider,omitempty"`
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

type SetTraceParams struct {

	Value TraceValue `json:"value"`
}
// Client capabilities for the showDocument request.
// 
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// The client has support for the showDocument
	// request.
	Support bool `json:"support"`
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
	Selection *Range `json:"selection,omitempty"`
	// An optional property to indicate whether the editor
	// showing the document should take focus or not.
	// Clients might ignore this property if an external
	// program is started.
	TakeFocus bool `json:"takeFocus,omitempty"`
	// The uri to show.
	Uri URI `json:"uri"`
}
// The result of a showDocument request.
// 
// @since 3.16.0
type ShowDocumentResult struct {
	// A boolean indicating if the show was successful.
	Success bool `json:"success"`
}
// The parameters of a notification message.
type ShowMessageParams struct {
	// The actual message.
	Message string `json:"message"`
	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`
}
// Show message request client capabilities
type ShowMessageRequestClientCapabilities struct {
	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *ClientShowMessageActionItemOptions `json:"messageActionItem,omitempty"`
}

type ShowMessageRequestParams struct {
	// The message action items to present.
	Actions []MessageActionItem `json:"actions,omitempty"`
	// The actual message.
	Message string `json:"message"`
	// The message type. See {@link MessageType}
	Type MessageType `json:"type"`
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
	ActiveParameter **uint32 `json:"activeParameter,omitempty"`
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
	SignatureInformation *ClientSignatureInformationOptions `json:"signatureInformation,omitempty"`
}
// Additional information about the context in which a signature help request was triggered.
// 
// @since 3.15.0
type SignatureHelpContext struct {
	// The currently active `SignatureHelp`.
	// 
	// The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on
	// the user navigating through available signatures.
	ActiveSignatureHelp *SignatureHelp `json:"activeSignatureHelp,omitempty"`
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
// Server Capabilities for a {@link SignatureHelpRequest}.
type SignatureHelpOptions struct {
	// List of characters that re-trigger signature help.
	// 
	// These trigger characters are only active when signature help is already showing. All trigger characters
	// are also counted as re-trigger characters.
	// 
	// @since 3.15.0
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
	// List of characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}
// Parameters for a {@link SignatureHelpRequest}.
type SignatureHelpParams struct {
	// The signature help context. This is only available if the client specifies
	// to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
	// 
	// @since 3.15.0
	Context *SignatureHelpContext `json:"context,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
	// List of characters that trigger signature help automatically.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	ActiveParameter **uint32 `json:"activeParameter,omitempty"`
	// The human-readable doc-comment of this signature. Will be shown
	// in the UI but can be omitted.
	Documentation *Or_SignatureInformation_Documentation `json:"documentation,omitempty"`
	// The label of this signature. Will be shown in
	// the UI.
	Label string `json:"label"`
	// The parameters of this signature.
	Parameters []ParameterInformation `json:"parameters,omitempty"`
}
// An interactive text edit.
// 
// @since 3.18.0
// @proposed
type SnippetTextEdit struct {
	// The actual identifier of the snippet edit.
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
	// The range of the text document to be manipulated.
	Range Range `json:"range"`
	// The snippet to be inserted.
	Snippet StringValue `json:"snippet"`
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
// Static registration options to be returned in the initialize
// request.
type StaticRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id string `json:"id,omitempty"`
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
	Tags []SymbolTag `json:"tags,omitempty"`
}
// Describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// How documents are synced to the server.
	SyncKind TextDocumentSyncKind `json:"syncKind"`
}
// Text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	// Capabilities specific to the various call hierarchy requests.
	// 
	// @since 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitempty"`
	// Capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`
	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`
	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	// 
	// @since 3.6.0
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitempty"`
	// Capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`
	// Capabilities specific to the `textDocument/declaration` request.
	// 
	// @since 3.14.0
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`
	// Capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`
	// Capabilities specific to the diagnostic pull model.
	// 
	// @since 3.17.0
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitempty"`
	// Capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitempty"`
	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitempty"`
	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitempty"`
	// Defines which filters the client supports.
	// 
	// @since 3.18.0
	Filters *TextDocumentFilterClientCapabilities `json:"filters,omitempty"`
	// Capabilities specific to the `textDocument/foldingRange` request.
	// 
	// @since 3.10.0
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`
	// Capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`
	// Capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitempty"`
	// Capabilities specific to the `textDocument/implementation` request.
	// 
	// @since 3.6.0
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`
	// Capabilities specific to the `textDocument/inlayHint` request.
	// 
	// @since 3.17.0
	InlayHint *InlayHintClientCapabilities `json:"inlayHint,omitempty"`
	// Client capabilities specific to inline completions.
	// 
	// @since 3.18.0
	// @proposed
	InlineCompletion *InlineCompletionClientCapabilities `json:"inlineCompletion,omitempty"`
	// Capabilities specific to the `textDocument/inlineValue` request.
	// 
	// @since 3.17.0
	InlineValue *InlineValueClientCapabilities `json:"inlineValue,omitempty"`
	// Capabilities specific to the `textDocument/linkedEditingRange` request.
	// 
	// @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitempty"`
	// Client capabilities specific to the `textDocument/moniker` request.
	// 
	// @since 3.16.0
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`
	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`
	// Capabilities specific to the `textDocument/publishDiagnostics` notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`
	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`
	// Capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references,omitempty"`
	// Capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitempty"`
	// Capabilities specific to the `textDocument/selectionRange` request.
	// 
	// @since 3.15.0
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`
	// Capabilities specific to the various semantic token request.
	// 
	// @since 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`
	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitempty"`
	// Defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`
	// Capabilities specific to the `textDocument/typeDefinition` request.
	// 
	// @since 3.6.0
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitempty"`
	// Capabilities specific to the various type hierarchy requests.
	// 
	// @since 3.17.0
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"typeHierarchy,omitempty"`
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
// @since 3.18.0
type TextDocumentContentChangeWholeDocument struct {
	// The new text of the whole document.
	Text string `json:"text"`
}
// Client capabilities for a text document content provider.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentClientCapabilities struct {
	// Text document content provider supports dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}
// Text document content provider options.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentOptions struct {
	// The schemes for which the server provides content.
	Schemes []string `json:"schemes"`
}
// Parameters for the `workspace/textDocumentContent` request.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentParams struct {
	// The uri of the text document.
	Uri DocumentUri `json:"uri"`
}
// Parameters for the `workspace/textDocumentContent/refresh` request.
// 
// @since 3.18.0
// @proposed
type TextDocumentContentRefreshParams struct {
	// The uri of the text document to refresh.
	Uri DocumentUri `json:"uri"`
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
	Edits []Or_TextDocumentEdit_Edits `json:"edits"`
	// The text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`
}

type TextDocumentFilterClientCapabilities struct {
	// The client supports Relative Patterns.
	// 
	// @since 3.18.0
	RelativePatternSupport bool `json:"relativePatternSupport,omitempty"`
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
	Pattern *GlobPattern `json:"pattern,omitempty"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme,omitempty"`
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
	Pattern *GlobPattern `json:"pattern,omitempty"`
	// A Uri {@link Uri.scheme scheme}, like `file` or `untitled`.
	Scheme string `json:"scheme"`
}
// A literal to identify a text document in the client.
type TextDocumentIdentifier struct {
	// The text document's uri.
	Uri DocumentUri `json:"uri"`
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
// A parameter literal used in requests to pass a text document and a position inside that
// document.
type TextDocumentPositionParams struct {
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
// General text document registration options.
type TextDocumentRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
}
// Save registration options.
type TextDocumentSaveRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector *DocumentSelector `json:"documentSelector"`
	// The client is supposed to include the content on save.
	IncludeText bool `json:"includeText,omitempty"`
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

type TextDocumentSyncOptions struct {
	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind `json:"change,omitempty"`
	// Open and close notifications are sent to the server. If omitted open close notification should not
	// be sent.
	OpenClose bool `json:"openClose,omitempty"`
	// If present save notifications are sent to the server. If omitted the notification should not be
	// sent.
	Save *Or_TextDocumentSyncOptions_Save `json:"save,omitempty"`
	// If present will save notifications are sent to the server. If omitted the notification should not be
	// sent.
	WillSave bool `json:"willSave,omitempty"`
	// If present will save wait until requests are sent to the server. If omitted the request should not be
	// sent.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`
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

type TypeDefinitionOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type TypeDefinitionParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The position inside the text document.
	Position Position `json:"position"`
	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// @since 3.17.0
type TypeHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
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
	Tags []SymbolTag `json:"tags,omitempty"`
	// The resource identifier of this item.
	Uri DocumentUri `json:"uri"`
}
// Type hierarchy options used during static registration.
// 
// @since 3.17.0
type TypeHierarchyOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// The parameter of a `typeHierarchy/subtypes` request.
// 
// @since 3.17.0
type TypeHierarchySubtypesParams struct {

	Item TypeHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// The parameter of a `typeHierarchy/supertypes` request.
// 
// @since 3.17.0
type TypeHierarchySupertypesParams struct {

	Item TypeHierarchyItem `json:"item"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// General parameters to unregister a request or notification.
type Unregistration struct {
	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	Id string `json:"id"`
	// The method to unregister for.
	Method string `json:"method"`
}

type UnregistrationParams struct {

	Unregisterations []Unregistration `json:"unregisterations"`
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
// A text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	// The text document's uri.
	Uri DocumentUri `json:"uri"`
	// The version number of this document.
	Version int32 `json:"version"`
}
// The parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {
	// The 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason `json:"reason"`
	// The document that will be saved.
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type WindowClientCapabilities struct {
	// Capabilities specific to the showDocument request.
	// 
	// @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
	// Capabilities specific to the showMessage request.
	// 
	// @since 3.16.0
	ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`
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

type WorkDoneProgressCancelParams struct {
	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}

type WorkDoneProgressCreateParams struct {
	// The token to be used to report progress.
	Token ProgressToken `json:"token"`
}

type WorkDoneProgressEnd struct {

	Kind string `json:"kind"`
	// Optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message string `json:"message,omitempty"`
}

type WorkDoneProgressOptions struct {

	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

type WorkDoneProgressParams struct {
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`
	// The client supports `workspace/configuration` requests.
	// 
	// @since 3.6.0
	Configuration bool `json:"configuration,omitempty"`
	// Capabilities specific to the diagnostic requests scoped to the
	// workspace.
	// 
	// @since 3.17.0.
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`
	// Capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`
	// Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`
	// Capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`
	// The client has support for file notifications/requests for user operations on files.
	// 
	// Since 3.16.0
	FileOperations *FileOperationClientCapabilities `json:"fileOperations,omitempty"`
	// Capabilities specific to the folding range requests scoped to the workspace.
	// 
	// @since 3.18.0
	// @proposed
	FoldingRange *FoldingRangeWorkspaceClientCapabilities `json:"foldingRange,omitempty"`
	// Capabilities specific to the inlay hint requests scoped to the
	// workspace.
	// 
	// @since 3.17.0.
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint,omitempty"`
	// Capabilities specific to the inline values requests scoped to the
	// workspace.
	// 
	// @since 3.17.0.
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue,omitempty"`
	// Capabilities specific to the semantic token requests scoped to the
	// workspace.
	// 
	// @since 3.16.0.
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`
	// Capabilities specific to the `workspace/symbol` request.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`
	// Capabilities specific to the `workspace/textDocumentContent` request.
	// 
	// @since 3.18.0
	// @proposed
	TextDocumentContent *TextDocumentContentClientCapabilities `json:"textDocumentContent,omitempty"`
	// Capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`
	// The client has support for workspace folders.
	// 
	// @since 3.6.0
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`
}
// Parameters of the workspace diagnostic request.
// 
// @since 3.17.0
type WorkspaceDiagnosticParams struct {
	// The additional identifier provided during registration.
	Identifier string `json:"identifier,omitempty"`
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
	// The currently known diagnostic reports with their
	// previous result ids.
	PreviousResultIds []PreviousResultId `json:"previousResultIds"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// A workspace diagnostic report.
// 
// @since 3.17.0
type WorkspaceDiagnosticReport struct {

	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
}
// A partial result for a workspace diagnostic report.
// 
// @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {

	Items []WorkspaceDocumentDiagnosticReport `json:"items"`
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
	ChangeAnnotations map[ChangeAnnotationIdentifier]ChangeAnnotation `json:"changeAnnotations,omitempty"`
	// Holds changes to existing resources.
	Changes map[DocumentUri][]TextEdit `json:"changes,omitempty"`
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
	DocumentChanges []Or_WorkspaceEdit_DocumentChanges `json:"documentChanges,omitempty"`
}

type WorkspaceEditClientCapabilities struct {
	// Whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	// 
	// @since 3.16.0
	ChangeAnnotationSupport *ChangeAnnotationsSupportOptions `json:"changeAnnotationSupport,omitempty"`
	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges bool `json:"documentChanges,omitempty"`
	// The failure handling strategy of a client if applying the workspace edit
	// fails.
	// 
	// @since 3.13.0
	FailureHandling *FailureHandlingKind `json:"failureHandling,omitempty"`
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
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`
	// Whether the client supports snippets as text edits.
	// 
	// @since 3.18.0
	// @proposed
	SnippetEditSupport bool `json:"snippetEditSupport,omitempty"`
}
// Additional data about a workspace edit.
// 
// @since 3.18.0
// @proposed
type WorkspaceEditMetadata struct {
	// Signal to the editor that this edit is a refactoring.
	IsRefactoring bool `json:"isRefactoring,omitempty"`
}
// A workspace folder inside a client.
type WorkspaceFolder struct {
	// The name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string `json:"name"`
	// The associated URI for this workspace folder.
	Uri URI `json:"uri"`
}
// The workspace folder change event.
type WorkspaceFoldersChangeEvent struct {
	// The array of added workspace folders
	Added []WorkspaceFolder `json:"added"`
	// The array of the removed workspace folders
	Removed []WorkspaceFolder `json:"removed"`
}

type WorkspaceFoldersInitializeParams struct {
	// The workspace folders configured in the client when the server starts.
	// 
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	// 
	// @since 3.6.0
	WorkspaceFolders *Or_WorkspaceFoldersInitializeParams_WorkspaceFolders `json:"workspaceFolders,omitempty"`
}

type WorkspaceFoldersServerCapabilities struct {
	// Whether the server wants to receive workspace folder
	// change notifications.
	// 
	// If a string is provided the string is treated as an ID
	// under which the notification is registered on the client
	// side. The ID can be used to unregister for these events
	// using the `client/unregisterCapability` request.
	ChangeNotifications *Or_WorkspaceFoldersServerCapabilities_ChangeNotifications `json:"changeNotifications,omitempty"`
	// The server has support for workspace folders
	Supported bool `json:"supported,omitempty"`
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
// Defines workspace specific capabilities of the server.
// 
// @since 3.18.0
type WorkspaceOptions struct {
	// The server is interested in notifications/requests for operations on files.
	// 
	// @since 3.16.0
	FileOperations *FileOperationOptions `json:"fileOperations,omitempty"`
	// The server supports the `workspace/textDocumentContent` request.
	// 
	// @since 3.18.0
	// @proposed
	TextDocumentContent *Or_WorkspaceOptions_TextDocumentContent `json:"textDocumentContent,omitempty"`
	// The server supports workspace folder.
	// 
	// @since 3.6.0
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitempty"`
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
	Location Or_WorkspaceSymbol_Location `json:"location"`
	// The name of this symbol.
	Name string `json:"name"`
	// Tags for this symbol.
	// 
	// @since 3.16.0
	Tags []SymbolTag `json:"tags,omitempty"`
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
	ResolveSupport *ClientSymbolResolveOptions `json:"resolveSupport,omitempty"`
	// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind *ClientSymbolKindOptions `json:"symbolKind,omitempty"`
	// The client supports tags on `SymbolInformation`.
	// Clients supporting tags have to handle unknown tags gracefully.
	// 
	// @since 3.16.0
	TagSupport *ClientSymbolTagOptions `json:"tagSupport,omitempty"`
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
// The parameters of a {@link WorkspaceSymbolRequest}.
type WorkspaceSymbolParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken `json:"partialResultToken,omitempty"`
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
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
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
// The initialize parameters
type _InitializeParams struct {
	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities `json:"capabilities"`
	// Information about the client
	// 
	// @since 3.15.0
	ClientInfo *ClientInfo `json:"clientInfo,omitempty"`
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
	RootPath **string `json:"rootPath,omitempty"`
	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	// 
	// @deprecated in favour of workspaceFolders.
	RootUri *DocumentUri `json:"rootUri"`
	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *TraceValue `json:"trace,omitempty"`
	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}
// Defines how values from a set of defaults and an individual item will be
// merged.
// 
// @since 3.18.0
// Defines how values from a set of defaults and an individual item will be
// merged.
// 
// @since 3.18.0
//go:generate stringer -type=ApplyKind
type ApplyKind uint32
const (
	Merge ApplyKind = 2
	Replace ApplyKind = 1
)

// A set of predefined code action kinds
// A set of predefined code action kinds
type CodeActionKind string
const (
	Empty CodeActionKind = ""
	Notebook CodeActionKind = "notebook"
	QuickFix CodeActionKind = "quickfix"
	Refactor CodeActionKind = "refactor"
	RefactorExtract CodeActionKind = "refactor.extract"
	RefactorInline CodeActionKind = "refactor.inline"
	RefactorMove CodeActionKind = "refactor.move"
	RefactorRewrite CodeActionKind = "refactor.rewrite"
	Source CodeActionKind = "source"
	SourceFixAll CodeActionKind = "source.fixAll"
	SourceOrganizeImports CodeActionKind = "source.organizeImports"
)

// Code action tags are extra annotations that tweak the behavior of a code action.
// 
// @since 3.18.0 - proposed
// Code action tags are extra annotations that tweak the behavior of a code action.
// 
// @since 3.18.0 - proposed
//go:generate stringer -type=CodeActionTag
type CodeActionTag uint32
const (
	LLMGenerated CodeActionTag = 1
)

// The reason why code actions were requested.
// 
// @since 3.17.0
// The reason why code actions were requested.
// 
// @since 3.17.0
//go:generate stringer -type=CodeActionTriggerKind
type CodeActionTriggerKind uint32
const (
	Automatic CodeActionTriggerKind = 2
	Invoked CodeActionTriggerKind = 1
)

// The kind of a completion entry.
// The kind of a completion entry.
//go:generate stringer -type=CompletionItemKind
type CompletionItemKind uint32
const (
	CompletionItemKind_Class CompletionItemKind = 7
	CompletionItemKind_Color CompletionItemKind = 16
	CompletionItemKind_Constant CompletionItemKind = 21
	CompletionItemKind_Constructor CompletionItemKind = 4
	CompletionItemKind_Enum CompletionItemKind = 13
	CompletionItemKind_EnumMember CompletionItemKind = 20
	CompletionItemKind_Event CompletionItemKind = 23
	CompletionItemKind_Field CompletionItemKind = 5
	CompletionItemKind_File CompletionItemKind = 17
	CompletionItemKind_Folder CompletionItemKind = 19
	CompletionItemKind_Function CompletionItemKind = 3
	CompletionItemKind_Interface CompletionItemKind = 8
	CompletionItemKind_Keyword CompletionItemKind = 14
	CompletionItemKind_Method CompletionItemKind = 2
	CompletionItemKind_Module CompletionItemKind = 9
	CompletionItemKind_Operator CompletionItemKind = 24
	CompletionItemKind_Property CompletionItemKind = 10
	CompletionItemKind_Reference CompletionItemKind = 18
	CompletionItemKind_Snippet CompletionItemKind = 15
	CompletionItemKind_Struct CompletionItemKind = 22
	CompletionItemKind_Text CompletionItemKind = 1
	CompletionItemKind_TypeParameter CompletionItemKind = 25
	CompletionItemKind_Unit CompletionItemKind = 11
	CompletionItemKind_Value CompletionItemKind = 12
	CompletionItemKind_Variable CompletionItemKind = 6
)

// Completion item tags are extra annotations that tweak the rendering of a completion
// item.
// 
// @since 3.15.0
// Completion item tags are extra annotations that tweak the rendering of a completion
// item.
// 
// @since 3.15.0
//go:generate stringer -type=CompletionItemTag
type CompletionItemTag uint32
const (
	Deprecated CompletionItemTag = 1
)

// How a completion was triggered
// How a completion was triggered
//go:generate stringer -type=CompletionTriggerKind
type CompletionTriggerKind uint32
const (
	CompletionTriggerKind_Invoked CompletionTriggerKind = 1
	CompletionTriggerKind_TriggerCharacter CompletionTriggerKind = 2
	CompletionTriggerKind_TriggerForIncompleteCompletions CompletionTriggerKind = 3
)

// The diagnostic's severity.
// The diagnostic's severity.
//go:generate stringer -type=DiagnosticSeverity
type DiagnosticSeverity uint32
const (
	Error DiagnosticSeverity = 1
	Hint DiagnosticSeverity = 4
	Information DiagnosticSeverity = 3
	Warning DiagnosticSeverity = 2
)

// The diagnostic tags.
// 
// @since 3.15.0
// The diagnostic tags.
// 
// @since 3.15.0
//go:generate stringer -type=DiagnosticTag
type DiagnosticTag uint32
const (
	DiagnosticTag_Deprecated DiagnosticTag = 2
	DiagnosticTag_Unnecessary DiagnosticTag = 1
)

// The document diagnostic report kinds.
// 
// @since 3.17.0
// The document diagnostic report kinds.
// 
// @since 3.17.0
type DocumentDiagnosticReportKind string
const (
	Full DocumentDiagnosticReportKind = "full"
	Unchanged DocumentDiagnosticReportKind = "unchanged"
)

// A document highlight kind.
// A document highlight kind.
//go:generate stringer -type=DocumentHighlightKind
type DocumentHighlightKind uint32
const (
	DocumentHighlightKind_Read DocumentHighlightKind = 2
	DocumentHighlightKind_Text DocumentHighlightKind = 1
	DocumentHighlightKind_Write DocumentHighlightKind = 3
)

// Predefined error codes.
// Predefined error codes.
//go:generate stringer -type=ErrorCodes
type ErrorCodes int32
const (
	InternalError ErrorCodes = -32603
	InvalidParams ErrorCodes = -32602
	InvalidRequest ErrorCodes = -32600
	MethodNotFound ErrorCodes = -32601
	ParseError ErrorCodes = -32700
	ServerNotInitialized ErrorCodes = -32002
	UnknownErrorCode ErrorCodes = -32001
)

type FailureHandlingKind string
const (
	Abort FailureHandlingKind = "abort"
	TextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"
	Transactional FailureHandlingKind = "transactional"
	Undo FailureHandlingKind = "undo"
)

// The file event type
// The file event type
//go:generate stringer -type=FileChangeType
type FileChangeType uint32
const (
	Changed FileChangeType = 2
	Created FileChangeType = 1
	Deleted FileChangeType = 3
)

// A pattern kind describing if a glob pattern matches a file a folder or
// both.
// 
// @since 3.16.0
// A pattern kind describing if a glob pattern matches a file a folder or
// both.
// 
// @since 3.16.0
type FileOperationPatternKind string
const (
	FileOperationPatternKind_File FileOperationPatternKind = "file"
	FileOperationPatternKind_Folder FileOperationPatternKind = "folder"
)

// A set of predefined range kinds.
// A set of predefined range kinds.
type FoldingRangeKind string
const (
	Comment FoldingRangeKind = "comment"
	Imports FoldingRangeKind = "imports"
	Region FoldingRangeKind = "region"
)

// Inlay hint kinds.
// 
// @since 3.17.0
// Inlay hint kinds.
// 
// @since 3.17.0
//go:generate stringer -type=InlayHintKind
type InlayHintKind uint32
const (
	Parameter InlayHintKind = 2
	Type InlayHintKind = 1
)

// Describes how an {@link InlineCompletionItemProvider inline completion provider} was triggered.
// 
// @since 3.18.0
// @proposed
// Describes how an {@link InlineCompletionItemProvider inline completion provider} was triggered.
// 
// @since 3.18.0
// @proposed
//go:generate stringer -type=InlineCompletionTriggerKind
type InlineCompletionTriggerKind uint32
const (
	InlineCompletionTriggerKind_Automatic InlineCompletionTriggerKind = 2
	InlineCompletionTriggerKind_Invoked InlineCompletionTriggerKind = 1
)

// Defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
// Defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
//go:generate stringer -type=InsertTextFormat
type InsertTextFormat uint32
const (
	InsertTextFormat_PlainText InsertTextFormat = 1
	InsertTextFormat_Snippet InsertTextFormat = 2
)

// How whitespace and indentation is handled during completion
// item insertion.
// 
// @since 3.16.0
// How whitespace and indentation is handled during completion
// item insertion.
// 
// @since 3.16.0
//go:generate stringer -type=InsertTextMode
type InsertTextMode uint32
const (
	AdjustIndentation InsertTextMode = 2
	AsIs InsertTextMode = 1
)

//go:generate stringer -type=LSPErrorCodes
type LSPErrorCodes int32
const (
	ContentModified LSPErrorCodes = -32801
	RequestCancelled LSPErrorCodes = -32800
	RequestFailed LSPErrorCodes = -32803
	ServerCancelled LSPErrorCodes = -32802
)

// Predefined Language kinds
// @since 3.18.0
// Predefined Language kinds
// @since 3.18.0
type LanguageKind string
const (
	ABAP LanguageKind = "abap"
	BibTeX LanguageKind = "bibtex"
	C LanguageKind = "c"
	CPP LanguageKind = "cpp"
	CSS LanguageKind = "css"
	CSharp LanguageKind = "csharp"
	Clojure LanguageKind = "clojure"
	Coffeescript LanguageKind = "coffeescript"
	D LanguageKind = "d"
	Dart LanguageKind = "dart"
	Delphi LanguageKind = "pascal"
	Diff LanguageKind = "diff"
	Dockerfile LanguageKind = "dockerfile"
	Elixir LanguageKind = "elixir"
	Erlang LanguageKind = "erlang"
	FSharp LanguageKind = "fsharp"
	GitCommit LanguageKind = "git-commit"
	GitRebase LanguageKind = "rebase"
	Go LanguageKind = "go"
	Groovy LanguageKind = "groovy"
	HTML LanguageKind = "html"
	Handlebars LanguageKind = "handlebars"
	Haskell LanguageKind = "haskell"
	Ini LanguageKind = "ini"
	JSON LanguageKind = "json"
	Java LanguageKind = "java"
	JavaScript LanguageKind = "javascript"
	JavaScriptReact LanguageKind = "javascriptreact"
	LaTeX LanguageKind = "latex"
	Less LanguageKind = "less"
	Lua LanguageKind = "lua"
	Makefile LanguageKind = "makefile"
	Markdown LanguageKind = "markdown"
	ObjectiveC LanguageKind = "objective-c"
	ObjectiveCPP LanguageKind = "objective-cpp"
	PHP LanguageKind = "php"
	Pascal LanguageKind = "pascal"
	Perl LanguageKind = "perl"
	Perl6 LanguageKind = "perl6"
	Powershell LanguageKind = "powershell"
	Pug LanguageKind = "jade"
	Python LanguageKind = "python"
	R LanguageKind = "r"
	Razor LanguageKind = "razor"
	Ruby LanguageKind = "ruby"
	Rust LanguageKind = "rust"
	SASS LanguageKind = "sass"
	SCSS LanguageKind = "scss"
	SQL LanguageKind = "sql"
	Scala LanguageKind = "scala"
	ShaderLab LanguageKind = "shaderlab"
	ShellScript LanguageKind = "shellscript"
	Swift LanguageKind = "swift"
	TeX LanguageKind = "tex"
	TypeScript LanguageKind = "typescript"
	TypeScriptReact LanguageKind = "typescriptreact"
	VisualBasic LanguageKind = "vb"
	WindowsBat LanguageKind = "bat"
	XML LanguageKind = "xml"
	XSL LanguageKind = "xsl"
	YAML LanguageKind = "yaml"
)

// Describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
// 
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
// Describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
// 
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
type MarkupKind string
const (
	MarkupKind_Markdown MarkupKind = "markdown"
	MarkupKind_PlainText MarkupKind = "plaintext"
)

// The message type
// The message type
//go:generate stringer -type=MessageType
type MessageType uint32
const (
	MessageType_Debug MessageType = 5
	MessageType_Error MessageType = 1
	MessageType_Info MessageType = 3
	MessageType_Log MessageType = 4
	MessageType_Warning MessageType = 2
)

// The moniker kind.
// 
// @since 3.16.0
// The moniker kind.
// 
// @since 3.16.0
type MonikerKind string
const (
	Export MonikerKind = "export"
	Import MonikerKind = "import"
	Local MonikerKind = "local"
)

// A notebook cell kind.
// 
// @since 3.17.0
// A notebook cell kind.
// 
// @since 3.17.0
//go:generate stringer -type=NotebookCellKind
type NotebookCellKind uint32
const (
	Code NotebookCellKind = 2
	Markup NotebookCellKind = 1
)

// A set of predefined position encoding kinds.
// 
// @since 3.17.0
// A set of predefined position encoding kinds.
// 
// @since 3.17.0
type PositionEncodingKind string
const (
	UTF16 PositionEncodingKind = "utf-16"
	UTF32 PositionEncodingKind = "utf-32"
	UTF8 PositionEncodingKind = "utf-8"
)

//go:generate stringer -type=PrepareSupportDefaultBehavior
type PrepareSupportDefaultBehavior uint32
const (
	Identifier PrepareSupportDefaultBehavior = 1
)

type ResourceOperationKind string
const (
	Create ResourceOperationKind = "create"
	Delete ResourceOperationKind = "delete"
	Rename ResourceOperationKind = "rename"
)

// A set of predefined token modifiers. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
// A set of predefined token modifiers. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
type SemanticTokenModifiers string
const (
	SemanticTokenModifiers_Abstract SemanticTokenModifiers = "abstract"
	SemanticTokenModifiers_Async SemanticTokenModifiers = "async"
	SemanticTokenModifiers_Declaration SemanticTokenModifiers = "declaration"
	SemanticTokenModifiers_DefaultLibrary SemanticTokenModifiers = "defaultLibrary"
	SemanticTokenModifiers_Definition SemanticTokenModifiers = "definition"
	SemanticTokenModifiers_Deprecated SemanticTokenModifiers = "deprecated"
	SemanticTokenModifiers_Documentation SemanticTokenModifiers = "documentation"
	SemanticTokenModifiers_Modification SemanticTokenModifiers = "modification"
	SemanticTokenModifiers_Readonly SemanticTokenModifiers = "readonly"
	SemanticTokenModifiers_Static SemanticTokenModifiers = "static"
)

// A set of predefined token types. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
// A set of predefined token types. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
type SemanticTokenTypes string
const (
	SemanticTokenTypes_Class SemanticTokenTypes = "class"
	SemanticTokenTypes_Comment SemanticTokenTypes = "comment"
	SemanticTokenTypes_Decorator SemanticTokenTypes = "decorator"
	SemanticTokenTypes_Enum SemanticTokenTypes = "enum"
	SemanticTokenTypes_EnumMember SemanticTokenTypes = "enumMember"
	SemanticTokenTypes_Event SemanticTokenTypes = "event"
	SemanticTokenTypes_Function SemanticTokenTypes = "function"
	SemanticTokenTypes_Interface SemanticTokenTypes = "interface"
	SemanticTokenTypes_Keyword SemanticTokenTypes = "keyword"
	SemanticTokenTypes_Label SemanticTokenTypes = "label"
	SemanticTokenTypes_Macro SemanticTokenTypes = "macro"
	SemanticTokenTypes_Method SemanticTokenTypes = "method"
	SemanticTokenTypes_Modifier SemanticTokenTypes = "modifier"
	SemanticTokenTypes_Namespace SemanticTokenTypes = "namespace"
	SemanticTokenTypes_Number SemanticTokenTypes = "number"
	SemanticTokenTypes_Operator SemanticTokenTypes = "operator"
	SemanticTokenTypes_Parameter SemanticTokenTypes = "parameter"
	SemanticTokenTypes_Property SemanticTokenTypes = "property"
	SemanticTokenTypes_Regexp SemanticTokenTypes = "regexp"
	SemanticTokenTypes_String SemanticTokenTypes = "string"
	SemanticTokenTypes_Struct SemanticTokenTypes = "struct"
	SemanticTokenTypes_Type SemanticTokenTypes = "type"
	SemanticTokenTypes_TypeParameter SemanticTokenTypes = "typeParameter"
	SemanticTokenTypes_Variable SemanticTokenTypes = "variable"
)

// How a signature help was triggered.
// 
// @since 3.15.0
// How a signature help was triggered.
// 
// @since 3.15.0
//go:generate stringer -type=SignatureHelpTriggerKind
type SignatureHelpTriggerKind uint32
const (
	SignatureHelpTriggerKind_ContentChange SignatureHelpTriggerKind = 3
	SignatureHelpTriggerKind_Invoked SignatureHelpTriggerKind = 1
	SignatureHelpTriggerKind_TriggerCharacter SignatureHelpTriggerKind = 2
)

// A symbol kind.
// A symbol kind.
//go:generate stringer -type=SymbolKind
type SymbolKind uint32
const (
	SymbolKind_Array SymbolKind = 18
	SymbolKind_Boolean SymbolKind = 17
	SymbolKind_Class SymbolKind = 5
	SymbolKind_Constant SymbolKind = 14
	SymbolKind_Constructor SymbolKind = 9
	SymbolKind_Enum SymbolKind = 10
	SymbolKind_EnumMember SymbolKind = 22
	SymbolKind_Event SymbolKind = 24
	SymbolKind_Field SymbolKind = 8
	SymbolKind_File SymbolKind = 1
	SymbolKind_Function SymbolKind = 12
	SymbolKind_Interface SymbolKind = 11
	SymbolKind_Key SymbolKind = 20
	SymbolKind_Method SymbolKind = 6
	SymbolKind_Module SymbolKind = 2
	SymbolKind_Namespace SymbolKind = 3
	SymbolKind_Null SymbolKind = 21
	SymbolKind_Number SymbolKind = 16
	SymbolKind_Object SymbolKind = 19
	SymbolKind_Operator SymbolKind = 25
	SymbolKind_Package SymbolKind = 4
	SymbolKind_Property SymbolKind = 7
	SymbolKind_String SymbolKind = 15
	SymbolKind_Struct SymbolKind = 23
	SymbolKind_TypeParameter SymbolKind = 26
	SymbolKind_Variable SymbolKind = 13
)

// Symbol tags are extra annotations that tweak the rendering of a symbol.
// 
// @since 3.16
// Symbol tags are extra annotations that tweak the rendering of a symbol.
// 
// @since 3.16
//go:generate stringer -type=SymbolTag
type SymbolTag uint32
const (
	SymbolTag_Deprecated SymbolTag = 1
)

// Represents reasons why a text document is saved.
// Represents reasons why a text document is saved.
//go:generate stringer -type=TextDocumentSaveReason
type TextDocumentSaveReason uint32
const (
	AfterDelay TextDocumentSaveReason = 2
	FocusOut TextDocumentSaveReason = 3
	Manual TextDocumentSaveReason = 1
)

// Defines how the host (editor) should sync
// document changes to the language server.
// Defines how the host (editor) should sync
// document changes to the language server.
//go:generate stringer -type=TextDocumentSyncKind
type TextDocumentSyncKind uint32
const (
	TextDocumentSyncKind_Full TextDocumentSyncKind = 1
	TextDocumentSyncKind_Incremental TextDocumentSyncKind = 2
	TextDocumentSyncKind_None TextDocumentSyncKind = 0
)

type TokenFormat string
const (
	Relative TokenFormat = "relative"
)

type TraceValue string
const (
	Messages TraceValue = "messages"
	Off TraceValue = "off"
	Verbose TraceValue = "verbose"
)

// Moniker uniqueness level to define scope of the moniker.
// 
// @since 3.16.0
// Moniker uniqueness level to define scope of the moniker.
// 
// @since 3.16.0
type UniquenessLevel string
const (
	Document UniquenessLevel = "document"
	Global UniquenessLevel = "global"
	Group UniquenessLevel = "group"
	Project UniquenessLevel = "project"
	Scheme UniquenessLevel = "scheme"
)

//go:generate stringer -type=WatchKind
type WatchKind uint32
const (
	WatchKind_Change WatchKind = 2
	WatchKind_Create WatchKind = 1
	WatchKind_Delete WatchKind = 4
)

type RequestMethod string

const (
	UnknownRequestMethod RequestMethod = ""
	CallHierarchyIncomingCallsMethod RequestMethod = "callHierarchy/incomingCalls"
	CallHierarchyOutgoingCallsMethod RequestMethod = "callHierarchy/outgoingCalls"
	ClientRegisterCapabilityMethod RequestMethod = "client/registerCapability"
	ClientUnregisterCapabilityMethod RequestMethod = "client/unregisterCapability"
	CodeActionResolveMethod RequestMethod = "codeAction/resolve"
	CodeLensResolveMethod RequestMethod = "codeLens/resolve"
	CompletionItemResolveMethod RequestMethod = "completionItem/resolve"
	DocumentLinkResolveMethod RequestMethod = "documentLink/resolve"
	InitializeMethod RequestMethod = "initialize"
	InlayHintResolveMethod RequestMethod = "inlayHint/resolve"
	ShutdownMethod RequestMethod = "shutdown"
	TextDocumentCodeActionMethod RequestMethod = "textDocument/codeAction"
	TextDocumentCodeLensMethod RequestMethod = "textDocument/codeLens"
	TextDocumentColorPresentationMethod RequestMethod = "textDocument/colorPresentation"
	TextDocumentCompletionMethod RequestMethod = "textDocument/completion"
	TextDocumentDeclarationMethod RequestMethod = "textDocument/declaration"
	TextDocumentDefinitionMethod RequestMethod = "textDocument/definition"
	TextDocumentDiagnosticMethod RequestMethod = "textDocument/diagnostic"
	TextDocumentDocumentColorMethod RequestMethod = "textDocument/documentColor"
	TextDocumentDocumentHighlightMethod RequestMethod = "textDocument/documentHighlight"
	TextDocumentDocumentLinkMethod RequestMethod = "textDocument/documentLink"
	TextDocumentDocumentSymbolMethod RequestMethod = "textDocument/documentSymbol"
	TextDocumentFoldingRangeMethod RequestMethod = "textDocument/foldingRange"
	TextDocumentFormattingMethod RequestMethod = "textDocument/formatting"
	TextDocumentHoverMethod RequestMethod = "textDocument/hover"
	TextDocumentImplementationMethod RequestMethod = "textDocument/implementation"
	TextDocumentInlayHintMethod RequestMethod = "textDocument/inlayHint"
	TextDocumentInlineCompletionMethod RequestMethod = "textDocument/inlineCompletion"
	TextDocumentInlineValueMethod RequestMethod = "textDocument/inlineValue"
	TextDocumentLinkedEditingRangeMethod RequestMethod = "textDocument/linkedEditingRange"
	TextDocumentMonikerMethod RequestMethod = "textDocument/moniker"
	TextDocumentOnTypeFormattingMethod RequestMethod = "textDocument/onTypeFormatting"
	TextDocumentPrepareCallHierarchyMethod RequestMethod = "textDocument/prepareCallHierarchy"
	TextDocumentPrepareRenameMethod RequestMethod = "textDocument/prepareRename"
	TextDocumentPrepareTypeHierarchyMethod RequestMethod = "textDocument/prepareTypeHierarchy"
	TextDocumentRangeFormattingMethod RequestMethod = "textDocument/rangeFormatting"
	TextDocumentRangesFormattingMethod RequestMethod = "textDocument/rangesFormatting"
	TextDocumentReferencesMethod RequestMethod = "textDocument/references"
	TextDocumentRenameMethod RequestMethod = "textDocument/rename"
	TextDocumentSelectionRangeMethod RequestMethod = "textDocument/selectionRange"
	TextDocumentSemanticTokensFullMethod RequestMethod = "textDocument/semanticTokens/full"
	TextDocumentSemanticTokensFullDeltaMethod RequestMethod = "textDocument/semanticTokens/full/delta"
	TextDocumentSemanticTokensRangeMethod RequestMethod = "textDocument/semanticTokens/range"
	TextDocumentSignatureHelpMethod RequestMethod = "textDocument/signatureHelp"
	TextDocumentTypeDefinitionMethod RequestMethod = "textDocument/typeDefinition"
	TextDocumentWillSaveWaitUntilMethod RequestMethod = "textDocument/willSaveWaitUntil"
	TypeHierarchySubtypesMethod RequestMethod = "typeHierarchy/subtypes"
	TypeHierarchySupertypesMethod RequestMethod = "typeHierarchy/supertypes"
	WindowShowDocumentMethod RequestMethod = "window/showDocument"
	WindowShowMessageRequestMethod RequestMethod = "window/showMessageRequest"
	WindowWorkDoneProgressCreateMethod RequestMethod = "window/workDoneProgress/create"
	WorkspaceApplyEditMethod RequestMethod = "workspace/applyEdit"
	WorkspaceCodeLensRefreshMethod RequestMethod = "workspace/codeLens/refresh"
	WorkspaceConfigurationMethod RequestMethod = "workspace/configuration"
	WorkspaceDiagnosticMethod RequestMethod = "workspace/diagnostic"
	WorkspaceDiagnosticRefreshMethod RequestMethod = "workspace/diagnostic/refresh"
	WorkspaceExecuteCommandMethod RequestMethod = "workspace/executeCommand"
	WorkspaceFoldingRangeRefreshMethod RequestMethod = "workspace/foldingRange/refresh"
	WorkspaceInlayHintRefreshMethod RequestMethod = "workspace/inlayHint/refresh"
	WorkspaceInlineValueRefreshMethod RequestMethod = "workspace/inlineValue/refresh"
	WorkspaceSemanticTokensRefreshMethod RequestMethod = "workspace/semanticTokens/refresh"
	WorkspaceSymbolMethod RequestMethod = "workspace/symbol"
	WorkspaceTextDocumentContentMethod RequestMethod = "workspace/textDocumentContent"
	WorkspaceTextDocumentContentRefreshMethod RequestMethod = "workspace/textDocumentContent/refresh"
	WorkspaceWillCreateFilesMethod RequestMethod = "workspace/willCreateFiles"
	WorkspaceWillDeleteFilesMethod RequestMethod = "workspace/willDeleteFiles"
	WorkspaceWillRenameFilesMethod RequestMethod = "workspace/willRenameFiles"
	WorkspaceWorkspaceFoldersMethod RequestMethod = "workspace/workspaceFolders"
	WorkspaceSymbolResolveMethod RequestMethod = "workspaceSymbol/resolve"
)

var RequestMethodMap = map[string]RequestMethod{
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CallHierarchyIncomingCallsParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CallHierarchyOutgoingCallsParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = CallHierarchyOutgoingCallsRequest(result)
   return nil
}
// The `client/registerCapability` request is sent from the server to the client to register a new capability
// handler on the client side.
type RegistrationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params RegistrationParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = RegistrationRequest(result)
   return nil
}
// The `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability
// handler on the client side.
type UnregistrationRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params UnregistrationParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CodeAction `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = CodeActionResolveRequest(result)
   return nil
}
// A request to resolve a command for a given code lens.
type CodeLensResolveRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CodeLens `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CompletionItem `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentLink `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params InitializeParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params InlayHint `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = ShutdownRequest(result)
   return nil
}
// A request to provide commands for the given text document and range.
type CodeActionRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CodeActionParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = CodeActionRequest(result)
   return nil
}
// A request to provide code lens for the given text document.
type CodeLensRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CodeLensParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ColorPresentationParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CompletionParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DeclarationParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DefinitionParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentDiagnosticParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentColorParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentHighlightParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = DocumentHighlightRequest(result)
   return nil
}
// A request to provide document links
type DocumentLinkRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentLinkParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentSymbolParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params FoldingRangeParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = FoldingRangeRequest(result)
   return nil
}
// A request to format a whole document.
type DocumentFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentFormattingParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params HoverParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ImplementationParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params InlayHintParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params InlineCompletionParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params InlineValueParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params LinkedEditingRangeParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params MonikerParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = MonikerRequest(result)
   return nil
}
// A request to format a document on type.
type DocumentOnTypeFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentOnTypeFormattingParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CallHierarchyPrepareParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params PrepareRenameParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params TypeHierarchyPrepareParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = TypeHierarchyPrepareRequest(result)
   return nil
}
// A request to format a range in a document.
type DocumentRangeFormattingRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentRangeFormattingParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DocumentRangesFormattingParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ReferenceParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = ReferencesRequest(result)
   return nil
}
// A request to rename a symbol.
type RenameRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params RenameParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params SelectionRangeParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = SelectionRangeRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params SemanticTokensParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = SemanticTokensRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensDeltaRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params SemanticTokensDeltaParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = SemanticTokensDeltaRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensRangeRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params SemanticTokensRangeParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = SemanticTokensRangeRequest(result)
   return nil
}

type SignatureHelpRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params SignatureHelpParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params TypeDefinitionParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params WillSaveTextDocumentParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params TypeHierarchySubtypesParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params TypeHierarchySupertypesParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ShowDocumentParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = ShowDocumentRequest(result)
   return nil
}
// The show message request is sent from the server to the client to show a message
// and a set of options actions to the user.
type ShowMessageRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ShowMessageRequestParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = ShowMessageRequest(result)
   return nil
}
// The `window/workDoneProgress/create` request is sent from the server to the client to initiate progress
// reporting from the server.
type WorkDoneProgressCreateRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params WorkDoneProgressCreateParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = WorkDoneProgressCreateRequest(result)
   return nil
}
// A request sent from the server to the client to modified certain resources.
type ApplyWorkspaceEditRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ApplyWorkspaceEditParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ConfigurationParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params WorkspaceDiagnosticParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = DiagnosticRefreshRequest(result)
   return nil
}
// A request send from the client to the server to execute a command. The request might return
// a workspace edit which the client will apply to the workspace.
type ExecuteCommandRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params ExecuteCommandParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = ExecuteCommandRequest(result)
   return nil
}
// @since 3.18.0
// @proposed
type FoldingRangeRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = FoldingRangeRefreshRequest(result)
   return nil
}
// @since 3.17.0
type InlayHintRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = InlayHintRefreshRequest(result)
   return nil
}
// @since 3.17.0
type InlineValueRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = InlineValueRefreshRequest(result)
   return nil
}
// @since 3.16.0
type SemanticTokensRefreshRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params WorkspaceSymbolParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params TextDocumentContentParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params TextDocumentContentRefreshParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params CreateFilesParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params DeleteFilesParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params RenameFilesParams `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = WillRenameFilesRequest(result)
   return nil
}
// The `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
type WorkspaceFoldersRequest struct {
	JsonRPC string `json:"jsonrpc"`
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
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
	ID Or_Request_id `json:"id"`
	Method RequestMethod `json:"method"`
	Params WorkspaceSymbol `json:"params"`
}
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
   if err := json.Unmarshal(x, &result); err != nil {
       return err
   }
   *t = WorkspaceSymbolResolveRequest(result)
   return nil
}
type CallHierarchyIncomingCallsResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_CallHierarchyIncomingCallsResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CallHierarchyIncomingCallsResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CallHierarchyIncomingCallsResponse(temp)
   return nil
}
type CallHierarchyOutgoingCallsResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_CallHierarchyOutgoingCallsResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CallHierarchyOutgoingCallsResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CallHierarchyOutgoingCallsResponse(temp)
   return nil
}
type RegistrationResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias RegistrationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = RegistrationResponse(temp)
   return nil
}
type UnregistrationResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias UnregistrationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = UnregistrationResponse(temp)
   return nil
}
type CodeActionResolveResponse struct {
   Id Or_Request_id `json:"id"`
  Result *CodeAction `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CodeActionResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeActionResolveResponse(temp)
   return nil
}
type CodeLensResolveResponse struct {
   Id Or_Request_id `json:"id"`
  Result *CodeLens `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CodeLensResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeLensResolveResponse(temp)
   return nil
}
type CompletionResolveResponse struct {
   Id Or_Request_id `json:"id"`
  Result *CompletionItem `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CompletionResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CompletionResolveResponse(temp)
   return nil
}
type DocumentLinkResolveResponse struct {
   Id Or_Request_id `json:"id"`
  Result *DocumentLink `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentLinkResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentLinkResolveResponse(temp)
   return nil
}
type InitializeResponse struct {
   Id Or_Request_id `json:"id"`
  Result *InitializeResult `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InitializeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InitializeResponse(temp)
   return nil
}
type InlayHintResolveResponse struct {
   Id Or_Request_id `json:"id"`
  Result *InlayHint `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InlayHintResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlayHintResolveResponse(temp)
   return nil
}
type ShutdownResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ShutdownResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ShutdownResponse(temp)
   return nil
}
type CodeActionResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_CodeActionResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CodeActionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeActionResponse(temp)
   return nil
}
type CodeLensResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_CodeLensResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CodeLensResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeLensResponse(temp)
   return nil
}
type ColorPresentationResponse struct {
   Id Or_Request_id `json:"id"`
  Result *[]ColorPresentation `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ColorPresentationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ColorPresentationResponse(temp)
   return nil
}
type CompletionResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_CompletionResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CompletionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CompletionResponse(temp)
   return nil
}
type DeclarationResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DeclarationResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DeclarationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DeclarationResponse(temp)
   return nil
}
type DefinitionResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DefinitionResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DefinitionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DefinitionResponse(temp)
   return nil
}
type DocumentDiagnosticResponse struct {
   Id Or_Request_id `json:"id"`
  Result *DocumentDiagnosticReport `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentDiagnosticResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentDiagnosticResponse(temp)
   return nil
}
type DocumentColorResponse struct {
   Id Or_Request_id `json:"id"`
  Result *[]ColorInformation `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentColorResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentColorResponse(temp)
   return nil
}
type DocumentHighlightResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentHighlightResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentHighlightResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentHighlightResponse(temp)
   return nil
}
type DocumentLinkResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentLinkResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentLinkResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentLinkResponse(temp)
   return nil
}
type DocumentSymbolResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentSymbolResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentSymbolResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentSymbolResponse(temp)
   return nil
}
type FoldingRangeResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_FoldingRangeResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias FoldingRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = FoldingRangeResponse(temp)
   return nil
}
type DocumentFormattingResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentFormattingResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentFormattingResponse(temp)
   return nil
}
type HoverResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Hover `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias HoverResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = HoverResponse(temp)
   return nil
}
type ImplementationResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_ImplementationResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ImplementationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ImplementationResponse(temp)
   return nil
}
type InlayHintResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_InlayHintResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InlayHintResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlayHintResponse(temp)
   return nil
}
type InlineCompletionResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_InlineCompletionResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InlineCompletionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlineCompletionResponse(temp)
   return nil
}
type InlineValueResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_InlineValueResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InlineValueResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlineValueResponse(temp)
   return nil
}
type LinkedEditingRangeResponse struct {
   Id Or_Request_id `json:"id"`
  Result **LinkedEditingRanges `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias LinkedEditingRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = LinkedEditingRangeResponse(temp)
   return nil
}
type MonikerResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_MonikerResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias MonikerResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = MonikerResponse(temp)
   return nil
}
type DocumentOnTypeFormattingResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentOnTypeFormattingResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentOnTypeFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentOnTypeFormattingResponse(temp)
   return nil
}
type CallHierarchyPrepareResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_CallHierarchyPrepareResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CallHierarchyPrepareResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CallHierarchyPrepareResponse(temp)
   return nil
}
type PrepareRenameResponse struct {
   Id Or_Request_id `json:"id"`
  Result **PrepareRenameResult `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias PrepareRenameResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = PrepareRenameResponse(temp)
   return nil
}
type TypeHierarchyPrepareResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_TypeHierarchyPrepareResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias TypeHierarchyPrepareResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeHierarchyPrepareResponse(temp)
   return nil
}
type DocumentRangeFormattingResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentRangeFormattingResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentRangeFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentRangeFormattingResponse(temp)
   return nil
}
type DocumentRangesFormattingResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_DocumentRangesFormattingResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DocumentRangesFormattingResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DocumentRangesFormattingResponse(temp)
   return nil
}
type ReferencesResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_ReferencesResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ReferencesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ReferencesResponse(temp)
   return nil
}
type RenameResponse struct {
   Id Or_Request_id `json:"id"`
  Result **WorkspaceEdit `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias RenameResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = RenameResponse(temp)
   return nil
}
type SelectionRangeResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_SelectionRangeResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias SelectionRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SelectionRangeResponse(temp)
   return nil
}
type SemanticTokensResponse struct {
   Id Or_Request_id `json:"id"`
  Result **SemanticTokens `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias SemanticTokensResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensResponse(temp)
   return nil
}
type SemanticTokensDeltaResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_SemanticTokensDeltaResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias SemanticTokensDeltaResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensDeltaResponse(temp)
   return nil
}
type SemanticTokensRangeResponse struct {
   Id Or_Request_id `json:"id"`
  Result **SemanticTokens `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias SemanticTokensRangeResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensRangeResponse(temp)
   return nil
}
type SignatureHelpResponse struct {
   Id Or_Request_id `json:"id"`
  Result **SignatureHelp `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias SignatureHelpResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SignatureHelpResponse(temp)
   return nil
}
type TypeDefinitionResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_TypeDefinitionResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias TypeDefinitionResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeDefinitionResponse(temp)
   return nil
}
type WillSaveTextDocumentWaitUntilResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_WillSaveTextDocumentWaitUntilResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WillSaveTextDocumentWaitUntilResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillSaveTextDocumentWaitUntilResponse(temp)
   return nil
}
type TypeHierarchySubtypesResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_TypeHierarchySubtypesResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias TypeHierarchySubtypesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeHierarchySubtypesResponse(temp)
   return nil
}
type TypeHierarchySupertypesResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_TypeHierarchySupertypesResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias TypeHierarchySupertypesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TypeHierarchySupertypesResponse(temp)
   return nil
}
type ShowDocumentResponse struct {
   Id Or_Request_id `json:"id"`
  Result *ShowDocumentResult `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ShowDocumentResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ShowDocumentResponse(temp)
   return nil
}
type ShowMessageResponse struct {
   Id Or_Request_id `json:"id"`
  Result **MessageActionItem `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ShowMessageResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ShowMessageResponse(temp)
   return nil
}
type WorkDoneProgressCreateResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WorkDoneProgressCreateResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkDoneProgressCreateResponse(temp)
   return nil
}
type ApplyWorkspaceEditResponse struct {
   Id Or_Request_id `json:"id"`
  Result *ApplyWorkspaceEditResult `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ApplyWorkspaceEditResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ApplyWorkspaceEditResponse(temp)
   return nil
}
type CodeLensRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias CodeLensRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = CodeLensRefreshResponse(temp)
   return nil
}
type ConfigurationResponse struct {
   Id Or_Request_id `json:"id"`
  Result *[]any `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ConfigurationResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ConfigurationResponse(temp)
   return nil
}
type WorkspaceDiagnosticResponse struct {
   Id Or_Request_id `json:"id"`
  Result *WorkspaceDiagnosticReport `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WorkspaceDiagnosticResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceDiagnosticResponse(temp)
   return nil
}
type DiagnosticRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias DiagnosticRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = DiagnosticRefreshResponse(temp)
   return nil
}
type ExecuteCommandResponse struct {
   Id Or_Request_id `json:"id"`
  Result **any `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias ExecuteCommandResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = ExecuteCommandResponse(temp)
   return nil
}
type FoldingRangeRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias FoldingRangeRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = FoldingRangeRefreshResponse(temp)
   return nil
}
type InlayHintRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InlayHintRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlayHintRefreshResponse(temp)
   return nil
}
type InlineValueRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias InlineValueRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = InlineValueRefreshResponse(temp)
   return nil
}
type SemanticTokensRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias SemanticTokensRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = SemanticTokensRefreshResponse(temp)
   return nil
}
type WorkspaceSymbolResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_WorkspaceSymbolResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WorkspaceSymbolResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceSymbolResponse(temp)
   return nil
}
type TextDocumentContentResponse struct {
   Id Or_Request_id `json:"id"`
  Result *TextDocumentContentResult `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias TextDocumentContentResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TextDocumentContentResponse(temp)
   return nil
}
type TextDocumentContentRefreshResponse struct {
   Id Or_Request_id `json:"id"`

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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias TextDocumentContentRefreshResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = TextDocumentContentRefreshResponse(temp)
   return nil
}
type WillCreateFilesResponse struct {
   Id Or_Request_id `json:"id"`
  Result **WorkspaceEdit `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WillCreateFilesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillCreateFilesResponse(temp)
   return nil
}
type WillDeleteFilesResponse struct {
   Id Or_Request_id `json:"id"`
  Result **WorkspaceEdit `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WillDeleteFilesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillDeleteFilesResponse(temp)
   return nil
}
type WillRenameFilesResponse struct {
   Id Or_Request_id `json:"id"`
  Result **WorkspaceEdit `json:"result,omitempty"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WillRenameFilesResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WillRenameFilesResponse(temp)
   return nil
}
type WorkspaceFoldersResponse struct {
   Id Or_Request_id `json:"id"`
  Result **Or_WorkspaceFoldersResponse `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WorkspaceFoldersResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceFoldersResponse(temp)
   return nil
}
type WorkspaceSymbolResolveResponse struct {
   Id Or_Request_id `json:"id"`
  Result *WorkspaceSymbol `json:"result,omitzero"`
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
       return fmt.Errorf("response must have an id and jsonrpc field.")
   }
  type Alias WorkspaceSymbolResolveResponse
   var temp Alias
   if err := json.Unmarshal(x, &temp); err != nil {
       return err
   }
  *t = WorkspaceSymbolResolveResponse(temp)
   return nil
}


type NotificationMethod string

const (
	UnknownNotificationMethod = ""
	OptionalCancelRequestMethod NotificationMethod = "$/cancelRequest"
	OptionalLogTraceMethod NotificationMethod = "$/logTrace"
	OptionalProgressMethod NotificationMethod = "$/progress"
	OptionalSetTraceMethod NotificationMethod = "$/setTrace"
	ExitMethod NotificationMethod = "exit"
	InitializedMethod NotificationMethod = "initialized"
	NotebookDocumentDidChangeMethod NotificationMethod = "notebookDocument/didChange"
	NotebookDocumentDidCloseMethod NotificationMethod = "notebookDocument/didClose"
	NotebookDocumentDidOpenMethod NotificationMethod = "notebookDocument/didOpen"
	NotebookDocumentDidSaveMethod NotificationMethod = "notebookDocument/didSave"
	TelemetryEventMethod NotificationMethod = "telemetry/event"
	TextDocumentDidChangeMethod NotificationMethod = "textDocument/didChange"
	TextDocumentDidCloseMethod NotificationMethod = "textDocument/didClose"
	TextDocumentDidOpenMethod NotificationMethod = "textDocument/didOpen"
	TextDocumentDidSaveMethod NotificationMethod = "textDocument/didSave"
	TextDocumentPublishDiagnosticsMethod NotificationMethod = "textDocument/publishDiagnostics"
	TextDocumentWillSaveMethod NotificationMethod = "textDocument/willSave"
	WindowLogMessageMethod NotificationMethod = "window/logMessage"
	WindowShowMessageMethod NotificationMethod = "window/showMessage"
	WindowWorkDoneProgressCancelMethod NotificationMethod = "window/workDoneProgress/cancel"
	WorkspaceDidChangeConfigurationMethod NotificationMethod = "workspace/didChangeConfiguration"
	WorkspaceDidChangeWatchedFilesMethod NotificationMethod = "workspace/didChangeWatchedFiles"
	WorkspaceDidChangeWorkspaceFoldersMethod NotificationMethod = "workspace/didChangeWorkspaceFolders"
	WorkspaceDidCreateFilesMethod NotificationMethod = "workspace/didCreateFiles"
	WorkspaceDidDeleteFilesMethod NotificationMethod = "workspace/didDeleteFiles"
	WorkspaceDidRenameFilesMethod NotificationMethod = "workspace/didRenameFiles"
)

var NotificationMethodMap = map[string]NotificationMethod{
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
	Method string `json:"method"`
	Params CancelParams `json:"params"`
}

type LogTraceNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params LogTraceParams `json:"params"`
}

type ProgressNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params ProgressParams `json:"params"`
}

type SetTraceNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params SetTraceParams `json:"params"`
}
// The exit event is sent from the client to the server to
// ask the server to exit its process.
type ExitNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
}
// The initialized notification is sent from the client to the
// server after the client is fully initialized and the server
// is allowed to send requests from the server to the client.
type InitializedNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params InitializedParams `json:"params"`
}

type DidChangeNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidChangeNotebookDocumentParams `json:"params"`
}
// A notification sent when a notebook closes.
// 
// @since 3.17.0
type DidCloseNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidCloseNotebookDocumentParams `json:"params"`
}
// A notification sent when a notebook opens.
// 
// @since 3.17.0
type DidOpenNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidOpenNotebookDocumentParams `json:"params"`
}
// A notification sent when a notebook document is saved.
// 
// @since 3.17.0
type DidSaveNotebookDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidSaveNotebookDocumentParams `json:"params"`
}
// The telemetry event notification is sent from the server to the client to ask
// the client to log telemetry data.
type TelemetryEventNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params any `json:"params"`
}
// The document change notification is sent from the client to the server to signal
// changes to a text document.
type DidChangeTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidChangeTextDocumentParams `json:"params"`
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
	Method string `json:"method"`
	Params DidCloseTextDocumentParams `json:"params"`
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
	Method string `json:"method"`
	Params DidOpenTextDocumentParams `json:"params"`
}
// The document save notification is sent from the client to the server when
// the document got saved in the client.
type DidSaveTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidSaveTextDocumentParams `json:"params"`
}
// Diagnostics notification are sent from the server to the client to signal
// results of validation runs.
type PublishDiagnosticsNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params PublishDiagnosticsParams `json:"params"`
}
// A document will save notification is sent from the client to the server before
// the document is actually saved.
type WillSaveTextDocumentNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params WillSaveTextDocumentParams `json:"params"`
}
// The log message notification is sent from the server to the client to ask
// the client to log a particular message.
type LogMessageNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params LogMessageParams `json:"params"`
}
// The show message notification is sent from a server to a client to ask
// the client to display a particular message in the user interface.
type ShowMessageNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params ShowMessageParams `json:"params"`
}
// The `window/workDoneProgress/cancel` notification is sent from  the client to the server to cancel a progress
// initiated on the server side.
type WorkDoneProgressCancelNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params WorkDoneProgressCancelParams `json:"params"`
}
// The configuration change notification is sent from the client to the server
// when the client's configuration has changed. The notification contains
// the changed configuration as defined by the language client.
type DidChangeConfigurationNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidChangeConfigurationParams `json:"params"`
}
// The watched files notification is sent from the client to the server when
// the client detects changes to file watched by the language client.
type DidChangeWatchedFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidChangeWatchedFilesParams `json:"params"`
}
// The `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace
// folder configuration changes.
type DidChangeWorkspaceFoldersNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DidChangeWorkspaceFoldersParams `json:"params"`
}
// The did create files notification is sent from the client to the server when
// files were created from within the client.
// 
// @since 3.16.0
type DidCreateFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params CreateFilesParams `json:"params"`
}
// The will delete files request is sent from the client to the server before files are actually
// deleted as long as the deletion is triggered from within the client.
// 
// @since 3.16.0
type DidDeleteFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params DeleteFilesParams `json:"params"`
}
// The did rename files notification is sent from the client to the server when
// files were renamed from within the client.
// 
// @since 3.16.0
type DidRenameFilesNotification struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Params RenameFilesParams `json:"params"`
}


// Or type for [Location,[]Location]
type Or_Declaration struct {
	Value any `json:"value"`
}
func (t *Or_Declaration) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Location:
        return json.Marshal(x)
    case []Location:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Location, []Location", t)
}
func (t *Or_Declaration) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of Location, []Location in Or_Declaration for value " + stringValue}
	}
   var h0 Location
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []Location
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Location, []Location in Or_Declaration for value " + stringValue}
}
// Or type for [Location,[]Location]
type Or_Definition struct {
	Value any `json:"value"`
}
func (t *Or_Definition) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Location:
        return json.Marshal(x)
    case []Location:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Location, []Location", t)
}
func (t *Or_Definition) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of Location, []Location in Or_Definition for value " + stringValue}
	}
   var h0 Location
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []Location
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Location, []Location in Or_Definition for value " + stringValue}
}
// Or type for [RelatedFullDocumentDiagnosticReport,RelatedUnchangedDocumentDiagnosticReport]
type Or_DocumentDiagnosticReport struct {
	Value any `json:"value"`
}
func (t *Or_DocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case RelatedFullDocumentDiagnosticReport:
        return json.Marshal(x)
    case RelatedUnchangedDocumentDiagnosticReport:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of RelatedFullDocumentDiagnosticReport, RelatedUnchangedDocumentDiagnosticReport", t)
}
func (t *Or_DocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of RelatedFullDocumentDiagnosticReport, RelatedUnchangedDocumentDiagnosticReport in Or_DocumentDiagnosticReport for value " + stringValue}
	}
   var h0 RelatedFullDocumentDiagnosticReport
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 RelatedUnchangedDocumentDiagnosticReport
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of RelatedFullDocumentDiagnosticReport, RelatedUnchangedDocumentDiagnosticReport in Or_DocumentDiagnosticReport for value " + stringValue}
}
// Or type for [TextDocumentFilter,NotebookCellTextDocumentFilter]
type Or_DocumentFilter struct {
	Value any `json:"value"`
}
func (t *Or_DocumentFilter) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextDocumentFilter:
        return json.Marshal(x)
    case NotebookCellTextDocumentFilter:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextDocumentFilter, NotebookCellTextDocumentFilter", t)
}
func (t *Or_DocumentFilter) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentFilter, NotebookCellTextDocumentFilter in Or_DocumentFilter for value " + stringValue}
	}
   var h0 TextDocumentFilter
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookCellTextDocumentFilter
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentFilter, NotebookCellTextDocumentFilter in Or_DocumentFilter for value " + stringValue}
}
// Or type for [Pattern,RelativePattern]
type Or_GlobPattern struct {
	Value any `json:"value"`
}
func (t *Or_GlobPattern) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Pattern:
        return json.Marshal(x)
    case RelativePattern:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Pattern, RelativePattern", t)
}
func (t *Or_GlobPattern) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of Pattern, RelativePattern in Or_GlobPattern for value " + stringValue}
	}
   var h0 Pattern
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 RelativePattern
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Pattern, RelativePattern in Or_GlobPattern for value " + stringValue}
}
// Or type for [InlineValueText,InlineValueVariableLookup,InlineValueEvaluatableExpression]
type Or_InlineValue struct {
	Value any `json:"value"`
}
func (t *Or_InlineValue) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case InlineValueText:
        return json.Marshal(x)
    case InlineValueVariableLookup:
        return json.Marshal(x)
    case InlineValueEvaluatableExpression:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of InlineValueText, InlineValueVariableLookup, InlineValueEvaluatableExpression", t)
}
func (t *Or_InlineValue) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of InlineValueText, InlineValueVariableLookup, InlineValueEvaluatableExpression in Or_InlineValue for value " + stringValue}
	}
   var h0 InlineValueText
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 InlineValueVariableLookup
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 InlineValueEvaluatableExpression
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of InlineValueText, InlineValueVariableLookup, InlineValueEvaluatableExpression in Or_InlineValue for value " + stringValue}
}
// Or type for [string,MarkedStringWithLanguage]
type Or_MarkedString struct {
	Value any `json:"value"`
}
func (t *Or_MarkedString) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case MarkedStringWithLanguage:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, MarkedStringWithLanguage", t)
}
func (t *Or_MarkedString) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkedStringWithLanguage in Or_MarkedString for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkedStringWithLanguage
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkedStringWithLanguage in Or_MarkedString for value " + stringValue}
}
// Or type for [NotebookDocumentFilterNotebookType,NotebookDocumentFilterScheme,NotebookDocumentFilterPattern]
type Or_NotebookDocumentFilter struct {
	Value any `json:"value"`
}
func (t *Or_NotebookDocumentFilter) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case NotebookDocumentFilterNotebookType:
        return json.Marshal(x)
    case NotebookDocumentFilterScheme:
        return json.Marshal(x)
    case NotebookDocumentFilterPattern:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern", t)
}
func (t *Or_NotebookDocumentFilter) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern in Or_NotebookDocumentFilter for value " + stringValue}
	}
   var h0 NotebookDocumentFilterNotebookType
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentFilterScheme
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 NotebookDocumentFilterPattern
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentFilterNotebookType, NotebookDocumentFilterScheme, NotebookDocumentFilterPattern in Or_NotebookDocumentFilter for value " + stringValue}
}
// Or type for [Range,PrepareRenamePlaceholder,PrepareRenameDefaultBehavior]
type Or_PrepareRenameResult struct {
	Value any `json:"value"`
}
func (t *Or_PrepareRenameResult) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Range:
        return json.Marshal(x)
    case PrepareRenamePlaceholder:
        return json.Marshal(x)
    case PrepareRenameDefaultBehavior:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Range, PrepareRenamePlaceholder, PrepareRenameDefaultBehavior", t)
}
func (t *Or_PrepareRenameResult) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of Range, PrepareRenamePlaceholder, PrepareRenameDefaultBehavior in Or_PrepareRenameResult for value " + stringValue}
	}
   var h0 Range
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 PrepareRenamePlaceholder
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 PrepareRenameDefaultBehavior
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Range, PrepareRenamePlaceholder, PrepareRenameDefaultBehavior in Or_PrepareRenameResult for value " + stringValue}
}
// Or type for [int32,string]
type Or_ProgressToken struct {
	Value any `json:"value"`
}
func (t *Or_ProgressToken) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case int32:
        return json.Marshal(x)
    case string:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of int32, string", t)
}
func (t *Or_ProgressToken) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_ProgressToken for value " + stringValue}
	}
   var h0 int32
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 string
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_ProgressToken for value " + stringValue}
}
// Or type for [TextDocumentContentChangePartial,TextDocumentContentChangeWholeDocument]
type Or_TextDocumentContentChangeEvent struct {
	Value any `json:"value"`
}
func (t *Or_TextDocumentContentChangeEvent) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextDocumentContentChangePartial:
        return json.Marshal(x)
    case TextDocumentContentChangeWholeDocument:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextDocumentContentChangePartial, TextDocumentContentChangeWholeDocument", t)
}
func (t *Or_TextDocumentContentChangeEvent) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentContentChangePartial, TextDocumentContentChangeWholeDocument in Or_TextDocumentContentChangeEvent for value " + stringValue}
	}
   var h0 TextDocumentContentChangePartial
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 TextDocumentContentChangeWholeDocument
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentContentChangePartial, TextDocumentContentChangeWholeDocument in Or_TextDocumentContentChangeEvent for value " + stringValue}
}
// Or type for [TextDocumentFilterLanguage,TextDocumentFilterScheme,TextDocumentFilterPattern]
type Or_TextDocumentFilter struct {
	Value any `json:"value"`
}
func (t *Or_TextDocumentFilter) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextDocumentFilterLanguage:
        return json.Marshal(x)
    case TextDocumentFilterScheme:
        return json.Marshal(x)
    case TextDocumentFilterPattern:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern", t)
}
func (t *Or_TextDocumentFilter) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern in Or_TextDocumentFilter for value " + stringValue}
	}
   var h0 TextDocumentFilterLanguage
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 TextDocumentFilterScheme
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 TextDocumentFilterPattern
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentFilterLanguage, TextDocumentFilterScheme, TextDocumentFilterPattern in Or_TextDocumentFilter for value " + stringValue}
}
// Or type for [WorkspaceFullDocumentDiagnosticReport,WorkspaceUnchangedDocumentDiagnosticReport]
type Or_WorkspaceDocumentDiagnosticReport struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case WorkspaceFullDocumentDiagnosticReport:
        return json.Marshal(x)
    case WorkspaceUnchangedDocumentDiagnosticReport:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of WorkspaceFullDocumentDiagnosticReport, WorkspaceUnchangedDocumentDiagnosticReport", t)
}
func (t *Or_WorkspaceDocumentDiagnosticReport) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of WorkspaceFullDocumentDiagnosticReport, WorkspaceUnchangedDocumentDiagnosticReport in Or_WorkspaceDocumentDiagnosticReport for value " + stringValue}
	}
   var h0 WorkspaceFullDocumentDiagnosticReport
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 WorkspaceUnchangedDocumentDiagnosticReport
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of WorkspaceFullDocumentDiagnosticReport, WorkspaceUnchangedDocumentDiagnosticReport in Or_WorkspaceDocumentDiagnosticReport for value " + stringValue}
}
// Or type for [int32,string]
type Or_CancelParams_Id struct {
	Value any `json:"value"`
}
func (t *Or_CancelParams_Id) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case int32:
        return json.Marshal(x)
    case string:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of int32, string", t)
}
func (t *Or_CancelParams_Id) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_CancelParams_Id for value " + stringValue}
	}
   var h0 int32
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 string
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_CancelParams_Id for value " + stringValue}
}
// Or type for [bool,ClientSemanticTokensRequestFullDelta]
type Or_ClientSemanticTokensRequestOptions_Full struct {
	Value any `json:"value"`
}
func (t *Or_ClientSemanticTokensRequestOptions_Full) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case ClientSemanticTokensRequestFullDelta:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, ClientSemanticTokensRequestFullDelta", t)
}
func (t *Or_ClientSemanticTokensRequestOptions_Full) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, ClientSemanticTokensRequestFullDelta in Or_ClientSemanticTokensRequestOptions_Full for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 ClientSemanticTokensRequestFullDelta
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, ClientSemanticTokensRequestFullDelta in Or_ClientSemanticTokensRequestOptions_Full for value " + stringValue}
}
// Or type for [bool,LSPObject]
type Or_ClientSemanticTokensRequestOptions_Range struct {
	Value any `json:"value"`
}
func (t *Or_ClientSemanticTokensRequestOptions_Range) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case map[string]any:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, map[string]any", t)
}
func (t *Or_ClientSemanticTokensRequestOptions_Range) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, map[string]any in Or_ClientSemanticTokensRequestOptions_Range for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 map[string]any
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, map[string]any in Or_ClientSemanticTokensRequestOptions_Range for value " + stringValue}
}
// Or type for [string,MarkupContent]
type Or_CompletionItem_Documentation struct {
	Value any `json:"value"`
}
func (t *Or_CompletionItem_Documentation) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case MarkupContent:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, MarkupContent", t)
}
func (t *Or_CompletionItem_Documentation) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_CompletionItem_Documentation for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkupContent
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_CompletionItem_Documentation for value " + stringValue}
}
// Or type for [TextEdit,InsertReplaceEdit]
type Or_CompletionItem_TextEdit struct {
	Value any `json:"value"`
}
func (t *Or_CompletionItem_TextEdit) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextEdit:
        return json.Marshal(x)
    case InsertReplaceEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextEdit, InsertReplaceEdit", t)
}
func (t *Or_CompletionItem_TextEdit) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextEdit, InsertReplaceEdit in Or_CompletionItem_TextEdit for value " + stringValue}
	}
   var h0 TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 InsertReplaceEdit
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextEdit, InsertReplaceEdit in Or_CompletionItem_TextEdit for value " + stringValue}
}
// Or type for [Range,EditRangeWithInsertReplace]
type Or_CompletionItemDefaults_EditRange struct {
	Value any `json:"value"`
}
func (t *Or_CompletionItemDefaults_EditRange) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Range:
        return json.Marshal(x)
    case EditRangeWithInsertReplace:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Range, EditRangeWithInsertReplace", t)
}
func (t *Or_CompletionItemDefaults_EditRange) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of Range, EditRangeWithInsertReplace in Or_CompletionItemDefaults_EditRange for value " + stringValue}
	}
   var h0 Range
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 EditRangeWithInsertReplace
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Range, EditRangeWithInsertReplace in Or_CompletionItemDefaults_EditRange for value " + stringValue}
}
// Or type for [int32,string]
type Or_Diagnostic_Code struct {
	Value any `json:"value"`
}
func (t *Or_Diagnostic_Code) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case int32:
        return json.Marshal(x)
    case string:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of int32, string", t)
}
func (t *Or_Diagnostic_Code) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_Diagnostic_Code for value " + stringValue}
	}
   var h0 int32
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 string
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_Diagnostic_Code for value " + stringValue}
}
// Or type for [string,[]string]
type Or_DidChangeConfigurationRegistrationOptions_Section struct {
	Value any `json:"value"`
}
func (t *Or_DidChangeConfigurationRegistrationOptions_Section) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case []string:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, []string", t)
}
func (t *Or_DidChangeConfigurationRegistrationOptions_Section) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, []string in Or_DidChangeConfigurationRegistrationOptions_Section for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []string
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, []string in Or_DidChangeConfigurationRegistrationOptions_Section for value " + stringValue}
}
// Or type for [FullDocumentDiagnosticReport,UnchangedDocumentDiagnosticReport]
type Or_DocumentDiagnosticReportPartialResult_RelatedDocuments struct {
	Value any `json:"value"`
}
func (t *Or_DocumentDiagnosticReportPartialResult_RelatedDocuments) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case FullDocumentDiagnosticReport:
        return json.Marshal(x)
    case UnchangedDocumentDiagnosticReport:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport", t)
}
func (t *Or_DocumentDiagnosticReportPartialResult_RelatedDocuments) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport in Or_DocumentDiagnosticReportPartialResult_RelatedDocuments for value " + stringValue}
	}
   var h0 FullDocumentDiagnosticReport
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 UnchangedDocumentDiagnosticReport
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport in Or_DocumentDiagnosticReportPartialResult_RelatedDocuments for value " + stringValue}
}
// Or type for [MarkupContent,MarkedString,[]MarkedString]
type Or_Hover_Contents struct {
	Value any `json:"value"`
}
func (t *Or_Hover_Contents) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case MarkupContent:
        return json.Marshal(x)
    case MarkedString:
        return json.Marshal(x)
    case []MarkedString:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of MarkupContent, MarkedString, []MarkedString", t)
}
func (t *Or_Hover_Contents) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of MarkupContent, MarkedString, []MarkedString in Or_Hover_Contents for value " + stringValue}
	}
   var h0 MarkupContent
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkedString
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 []MarkedString
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of MarkupContent, MarkedString, []MarkedString in Or_Hover_Contents for value " + stringValue}
}
// Or type for [[]WorkspaceFolder,nil]
type Or_InitializeParams_WorkspaceFolders struct {
	Value any `json:"value"`
}
func (t *Or_InitializeParams_WorkspaceFolders) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []WorkspaceFolder:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []WorkspaceFolder, nil", t)
}
func (t *Or_InitializeParams_WorkspaceFolders) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []WorkspaceFolder
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []WorkspaceFolder, nil in Or_InitializeParams_WorkspaceFolders for value " + stringValue}
}
// Or type for [string,[]InlayHintLabelPart]
type Or_InlayHint_Label struct {
	Value any `json:"value"`
}
func (t *Or_InlayHint_Label) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case []InlayHintLabelPart:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, []InlayHintLabelPart", t)
}
func (t *Or_InlayHint_Label) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, []InlayHintLabelPart in Or_InlayHint_Label for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []InlayHintLabelPart
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, []InlayHintLabelPart in Or_InlayHint_Label for value " + stringValue}
}
// Or type for [string,MarkupContent]
type Or_InlayHint_Tooltip struct {
	Value any `json:"value"`
}
func (t *Or_InlayHint_Tooltip) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case MarkupContent:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, MarkupContent", t)
}
func (t *Or_InlayHint_Tooltip) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_InlayHint_Tooltip for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkupContent
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_InlayHint_Tooltip for value " + stringValue}
}
// Or type for [string,MarkupContent]
type Or_InlayHintLabelPart_Tooltip struct {
	Value any `json:"value"`
}
func (t *Or_InlayHintLabelPart_Tooltip) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case MarkupContent:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, MarkupContent", t)
}
func (t *Or_InlayHintLabelPart_Tooltip) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_InlayHintLabelPart_Tooltip for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkupContent
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_InlayHintLabelPart_Tooltip for value " + stringValue}
}
// Or type for [string,StringValue]
type Or_InlineCompletionItem_InsertText struct {
	Value any `json:"value"`
}
func (t *Or_InlineCompletionItem_InsertText) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case StringValue:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, StringValue", t)
}
func (t *Or_InlineCompletionItem_InsertText) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, StringValue in Or_InlineCompletionItem_InsertText for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 StringValue
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, StringValue in Or_InlineCompletionItem_InsertText for value " + stringValue}
}
// Or type for [string,NotebookDocumentFilter]
type Or_NotebookCellTextDocumentFilter_Notebook struct {
	Value any `json:"value"`
}
func (t *Or_NotebookCellTextDocumentFilter_Notebook) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case NotebookDocumentFilter:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, NotebookDocumentFilter", t)
}
func (t *Or_NotebookCellTextDocumentFilter_Notebook) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, NotebookDocumentFilter in Or_NotebookCellTextDocumentFilter_Notebook for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentFilter
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, NotebookDocumentFilter in Or_NotebookCellTextDocumentFilter_Notebook for value " + stringValue}
}
// Or type for [string,NotebookDocumentFilter]
type Or_NotebookDocumentFilterWithCells_Notebook struct {
	Value any `json:"value"`
}
func (t *Or_NotebookDocumentFilterWithCells_Notebook) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case NotebookDocumentFilter:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, NotebookDocumentFilter", t)
}
func (t *Or_NotebookDocumentFilterWithCells_Notebook) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, NotebookDocumentFilter in Or_NotebookDocumentFilterWithCells_Notebook for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentFilter
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, NotebookDocumentFilter in Or_NotebookDocumentFilterWithCells_Notebook for value " + stringValue}
}
// Or type for [string,NotebookDocumentFilter]
type Or_NotebookDocumentFilterWithNotebook_Notebook struct {
	Value any `json:"value"`
}
func (t *Or_NotebookDocumentFilterWithNotebook_Notebook) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case NotebookDocumentFilter:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, NotebookDocumentFilter", t)
}
func (t *Or_NotebookDocumentFilterWithNotebook_Notebook) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, NotebookDocumentFilter in Or_NotebookDocumentFilterWithNotebook_Notebook for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentFilter
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, NotebookDocumentFilter in Or_NotebookDocumentFilterWithNotebook_Notebook for value " + stringValue}
}
// Or type for [NotebookDocumentFilterWithNotebook,NotebookDocumentFilterWithCells]
type Or_NotebookDocumentSyncOptions_NotebookSelector struct {
	Value any `json:"value"`
}
func (t *Or_NotebookDocumentSyncOptions_NotebookSelector) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case NotebookDocumentFilterWithNotebook:
        return json.Marshal(x)
    case NotebookDocumentFilterWithCells:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells", t)
}
func (t *Or_NotebookDocumentSyncOptions_NotebookSelector) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells in Or_NotebookDocumentSyncOptions_NotebookSelector for value " + stringValue}
	}
   var h0 NotebookDocumentFilterWithNotebook
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentFilterWithCells
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells in Or_NotebookDocumentSyncOptions_NotebookSelector for value " + stringValue}
}
// Or type for [NotebookDocumentFilterWithNotebook,NotebookDocumentFilterWithCells]
type Or_NotebookDocumentSyncRegistrationOptions_NotebookSelector struct {
	Value any `json:"value"`
}
func (t *Or_NotebookDocumentSyncRegistrationOptions_NotebookSelector) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case NotebookDocumentFilterWithNotebook:
        return json.Marshal(x)
    case NotebookDocumentFilterWithCells:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells", t)
}
func (t *Or_NotebookDocumentSyncRegistrationOptions_NotebookSelector) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells in Or_NotebookDocumentSyncRegistrationOptions_NotebookSelector for value " + stringValue}
	}
   var h0 NotebookDocumentFilterWithNotebook
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentFilterWithCells
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentFilterWithNotebook, NotebookDocumentFilterWithCells in Or_NotebookDocumentSyncRegistrationOptions_NotebookSelector for value " + stringValue}
}
// Or type for [string,MarkupContent]
type Or_ParameterInformation_Documentation struct {
	Value any `json:"value"`
}
func (t *Or_ParameterInformation_Documentation) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case MarkupContent:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, MarkupContent", t)
}
func (t *Or_ParameterInformation_Documentation) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_ParameterInformation_Documentation for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkupContent
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_ParameterInformation_Documentation for value " + stringValue}
}
// Or type for [string,Tuple[uint32, uint32]]
type Or_ParameterInformation_Label struct {
	Value any `json:"value"`
}
func (t *Or_ParameterInformation_Label) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case Tuple[uint32, uint32]:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, Tuple[uint32, uint32]", t)
}
func (t *Or_ParameterInformation_Label) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, Tuple[uint32, uint32] in Or_ParameterInformation_Label for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 Tuple[uint32, uint32]
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, Tuple[uint32, uint32] in Or_ParameterInformation_Label for value " + stringValue}
}
// Or type for [FullDocumentDiagnosticReport,UnchangedDocumentDiagnosticReport]
type Or_RelatedFullDocumentDiagnosticReport_RelatedDocuments struct {
	Value any `json:"value"`
}
func (t *Or_RelatedFullDocumentDiagnosticReport_RelatedDocuments) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case FullDocumentDiagnosticReport:
        return json.Marshal(x)
    case UnchangedDocumentDiagnosticReport:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport", t)
}
func (t *Or_RelatedFullDocumentDiagnosticReport_RelatedDocuments) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport in Or_RelatedFullDocumentDiagnosticReport_RelatedDocuments for value " + stringValue}
	}
   var h0 FullDocumentDiagnosticReport
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 UnchangedDocumentDiagnosticReport
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport in Or_RelatedFullDocumentDiagnosticReport_RelatedDocuments for value " + stringValue}
}
// Or type for [FullDocumentDiagnosticReport,UnchangedDocumentDiagnosticReport]
type Or_RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments struct {
	Value any `json:"value"`
}
func (t *Or_RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case FullDocumentDiagnosticReport:
        return json.Marshal(x)
    case UnchangedDocumentDiagnosticReport:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport", t)
}
func (t *Or_RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport in Or_RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments for value " + stringValue}
	}
   var h0 FullDocumentDiagnosticReport
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 UnchangedDocumentDiagnosticReport
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of FullDocumentDiagnosticReport, UnchangedDocumentDiagnosticReport in Or_RelatedUnchangedDocumentDiagnosticReport_RelatedDocuments for value " + stringValue}
}
// Or type for [WorkspaceFolder,URI]
type Or_RelativePattern_BaseUri struct {
	Value any `json:"value"`
}
func (t *Or_RelativePattern_BaseUri) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case WorkspaceFolder:
        return json.Marshal(x)
    case URI:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of WorkspaceFolder, URI", t)
}
func (t *Or_RelativePattern_BaseUri) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of WorkspaceFolder, URI in Or_RelativePattern_BaseUri for value " + stringValue}
	}
   var h0 WorkspaceFolder
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 URI
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of WorkspaceFolder, URI in Or_RelativePattern_BaseUri for value " + stringValue}
}
// Or type for [bool,SemanticTokensFullDelta]
type Or_SemanticTokensOptions_Full struct {
	Value any `json:"value"`
}
func (t *Or_SemanticTokensOptions_Full) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case SemanticTokensFullDelta:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, SemanticTokensFullDelta", t)
}
func (t *Or_SemanticTokensOptions_Full) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, SemanticTokensFullDelta in Or_SemanticTokensOptions_Full for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 SemanticTokensFullDelta
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, SemanticTokensFullDelta in Or_SemanticTokensOptions_Full for value " + stringValue}
}
// Or type for [bool,LSPObject]
type Or_SemanticTokensOptions_Range struct {
	Value any `json:"value"`
}
func (t *Or_SemanticTokensOptions_Range) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case map[string]any:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, map[string]any", t)
}
func (t *Or_SemanticTokensOptions_Range) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, map[string]any in Or_SemanticTokensOptions_Range for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 map[string]any
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, map[string]any in Or_SemanticTokensOptions_Range for value " + stringValue}
}
// Or type for [bool,SemanticTokensFullDelta]
type Or_SemanticTokensRegistrationOptions_Full struct {
	Value any `json:"value"`
}
func (t *Or_SemanticTokensRegistrationOptions_Full) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case SemanticTokensFullDelta:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, SemanticTokensFullDelta", t)
}
func (t *Or_SemanticTokensRegistrationOptions_Full) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, SemanticTokensFullDelta in Or_SemanticTokensRegistrationOptions_Full for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 SemanticTokensFullDelta
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, SemanticTokensFullDelta in Or_SemanticTokensRegistrationOptions_Full for value " + stringValue}
}
// Or type for [bool,LSPObject]
type Or_SemanticTokensRegistrationOptions_Range struct {
	Value any `json:"value"`
}
func (t *Or_SemanticTokensRegistrationOptions_Range) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case map[string]any:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, map[string]any", t)
}
func (t *Or_SemanticTokensRegistrationOptions_Range) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, map[string]any in Or_SemanticTokensRegistrationOptions_Range for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 map[string]any
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, map[string]any in Or_SemanticTokensRegistrationOptions_Range for value " + stringValue}
}
// Or type for [bool,CallHierarchyOptions,CallHierarchyRegistrationOptions]
type Or_ServerCapabilities_CallHierarchyProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_CallHierarchyProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case CallHierarchyOptions:
        return json.Marshal(x)
    case CallHierarchyRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, CallHierarchyOptions, CallHierarchyRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_CallHierarchyProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, CallHierarchyOptions, CallHierarchyRegistrationOptions in Or_ServerCapabilities_CallHierarchyProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 CallHierarchyOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 CallHierarchyRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, CallHierarchyOptions, CallHierarchyRegistrationOptions in Or_ServerCapabilities_CallHierarchyProvider for value " + stringValue}
}
// Or type for [bool,CodeActionOptions]
type Or_ServerCapabilities_CodeActionProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_CodeActionProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case CodeActionOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, CodeActionOptions", t)
}
func (t *Or_ServerCapabilities_CodeActionProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, CodeActionOptions in Or_ServerCapabilities_CodeActionProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 CodeActionOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, CodeActionOptions in Or_ServerCapabilities_CodeActionProvider for value " + stringValue}
}
// Or type for [bool,DocumentColorOptions,DocumentColorRegistrationOptions]
type Or_ServerCapabilities_ColorProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_ColorProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DocumentColorOptions:
        return json.Marshal(x)
    case DocumentColorRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DocumentColorOptions, DocumentColorRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_ColorProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentColorOptions, DocumentColorRegistrationOptions in Or_ServerCapabilities_ColorProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DocumentColorOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 DocumentColorRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentColorOptions, DocumentColorRegistrationOptions in Or_ServerCapabilities_ColorProvider for value " + stringValue}
}
// Or type for [bool,DeclarationOptions,DeclarationRegistrationOptions]
type Or_ServerCapabilities_DeclarationProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DeclarationProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DeclarationOptions:
        return json.Marshal(x)
    case DeclarationRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DeclarationOptions, DeclarationRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_DeclarationProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DeclarationOptions, DeclarationRegistrationOptions in Or_ServerCapabilities_DeclarationProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DeclarationOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 DeclarationRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DeclarationOptions, DeclarationRegistrationOptions in Or_ServerCapabilities_DeclarationProvider for value " + stringValue}
}
// Or type for [bool,DefinitionOptions]
type Or_ServerCapabilities_DefinitionProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DefinitionProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DefinitionOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DefinitionOptions", t)
}
func (t *Or_ServerCapabilities_DefinitionProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DefinitionOptions in Or_ServerCapabilities_DefinitionProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DefinitionOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DefinitionOptions in Or_ServerCapabilities_DefinitionProvider for value " + stringValue}
}
// Or type for [DiagnosticOptions,DiagnosticRegistrationOptions]
type Or_ServerCapabilities_DiagnosticProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DiagnosticProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case DiagnosticOptions:
        return json.Marshal(x)
    case DiagnosticRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of DiagnosticOptions, DiagnosticRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_DiagnosticProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of DiagnosticOptions, DiagnosticRegistrationOptions in Or_ServerCapabilities_DiagnosticProvider for value " + stringValue}
	}
   var h0 DiagnosticOptions
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DiagnosticRegistrationOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of DiagnosticOptions, DiagnosticRegistrationOptions in Or_ServerCapabilities_DiagnosticProvider for value " + stringValue}
}
// Or type for [bool,DocumentFormattingOptions]
type Or_ServerCapabilities_DocumentFormattingProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DocumentFormattingProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DocumentFormattingOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DocumentFormattingOptions", t)
}
func (t *Or_ServerCapabilities_DocumentFormattingProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentFormattingOptions in Or_ServerCapabilities_DocumentFormattingProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DocumentFormattingOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentFormattingOptions in Or_ServerCapabilities_DocumentFormattingProvider for value " + stringValue}
}
// Or type for [bool,DocumentHighlightOptions]
type Or_ServerCapabilities_DocumentHighlightProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DocumentHighlightProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DocumentHighlightOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DocumentHighlightOptions", t)
}
func (t *Or_ServerCapabilities_DocumentHighlightProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentHighlightOptions in Or_ServerCapabilities_DocumentHighlightProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DocumentHighlightOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentHighlightOptions in Or_ServerCapabilities_DocumentHighlightProvider for value " + stringValue}
}
// Or type for [bool,DocumentRangeFormattingOptions]
type Or_ServerCapabilities_DocumentRangeFormattingProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DocumentRangeFormattingProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DocumentRangeFormattingOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DocumentRangeFormattingOptions", t)
}
func (t *Or_ServerCapabilities_DocumentRangeFormattingProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentRangeFormattingOptions in Or_ServerCapabilities_DocumentRangeFormattingProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DocumentRangeFormattingOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentRangeFormattingOptions in Or_ServerCapabilities_DocumentRangeFormattingProvider for value " + stringValue}
}
// Or type for [bool,DocumentSymbolOptions]
type Or_ServerCapabilities_DocumentSymbolProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_DocumentSymbolProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case DocumentSymbolOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, DocumentSymbolOptions", t)
}
func (t *Or_ServerCapabilities_DocumentSymbolProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentSymbolOptions in Or_ServerCapabilities_DocumentSymbolProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 DocumentSymbolOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, DocumentSymbolOptions in Or_ServerCapabilities_DocumentSymbolProvider for value " + stringValue}
}
// Or type for [bool,FoldingRangeOptions,FoldingRangeRegistrationOptions]
type Or_ServerCapabilities_FoldingRangeProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_FoldingRangeProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case FoldingRangeOptions:
        return json.Marshal(x)
    case FoldingRangeRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, FoldingRangeOptions, FoldingRangeRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_FoldingRangeProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, FoldingRangeOptions, FoldingRangeRegistrationOptions in Or_ServerCapabilities_FoldingRangeProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 FoldingRangeOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 FoldingRangeRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, FoldingRangeOptions, FoldingRangeRegistrationOptions in Or_ServerCapabilities_FoldingRangeProvider for value " + stringValue}
}
// Or type for [bool,HoverOptions]
type Or_ServerCapabilities_HoverProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_HoverProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case HoverOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, HoverOptions", t)
}
func (t *Or_ServerCapabilities_HoverProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, HoverOptions in Or_ServerCapabilities_HoverProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 HoverOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, HoverOptions in Or_ServerCapabilities_HoverProvider for value " + stringValue}
}
// Or type for [bool,ImplementationOptions,ImplementationRegistrationOptions]
type Or_ServerCapabilities_ImplementationProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_ImplementationProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case ImplementationOptions:
        return json.Marshal(x)
    case ImplementationRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, ImplementationOptions, ImplementationRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_ImplementationProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, ImplementationOptions, ImplementationRegistrationOptions in Or_ServerCapabilities_ImplementationProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 ImplementationOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 ImplementationRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, ImplementationOptions, ImplementationRegistrationOptions in Or_ServerCapabilities_ImplementationProvider for value " + stringValue}
}
// Or type for [bool,InlayHintOptions,InlayHintRegistrationOptions]
type Or_ServerCapabilities_InlayHintProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_InlayHintProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case InlayHintOptions:
        return json.Marshal(x)
    case InlayHintRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, InlayHintOptions, InlayHintRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_InlayHintProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, InlayHintOptions, InlayHintRegistrationOptions in Or_ServerCapabilities_InlayHintProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 InlayHintOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 InlayHintRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, InlayHintOptions, InlayHintRegistrationOptions in Or_ServerCapabilities_InlayHintProvider for value " + stringValue}
}
// Or type for [bool,InlineCompletionOptions]
type Or_ServerCapabilities_InlineCompletionProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_InlineCompletionProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case InlineCompletionOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, InlineCompletionOptions", t)
}
func (t *Or_ServerCapabilities_InlineCompletionProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, InlineCompletionOptions in Or_ServerCapabilities_InlineCompletionProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 InlineCompletionOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, InlineCompletionOptions in Or_ServerCapabilities_InlineCompletionProvider for value " + stringValue}
}
// Or type for [bool,InlineValueOptions,InlineValueRegistrationOptions]
type Or_ServerCapabilities_InlineValueProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_InlineValueProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case InlineValueOptions:
        return json.Marshal(x)
    case InlineValueRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, InlineValueOptions, InlineValueRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_InlineValueProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, InlineValueOptions, InlineValueRegistrationOptions in Or_ServerCapabilities_InlineValueProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 InlineValueOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 InlineValueRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, InlineValueOptions, InlineValueRegistrationOptions in Or_ServerCapabilities_InlineValueProvider for value " + stringValue}
}
// Or type for [bool,LinkedEditingRangeOptions,LinkedEditingRangeRegistrationOptions]
type Or_ServerCapabilities_LinkedEditingRangeProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_LinkedEditingRangeProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case LinkedEditingRangeOptions:
        return json.Marshal(x)
    case LinkedEditingRangeRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, LinkedEditingRangeOptions, LinkedEditingRangeRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_LinkedEditingRangeProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, LinkedEditingRangeOptions, LinkedEditingRangeRegistrationOptions in Or_ServerCapabilities_LinkedEditingRangeProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 LinkedEditingRangeOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 LinkedEditingRangeRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, LinkedEditingRangeOptions, LinkedEditingRangeRegistrationOptions in Or_ServerCapabilities_LinkedEditingRangeProvider for value " + stringValue}
}
// Or type for [bool,MonikerOptions,MonikerRegistrationOptions]
type Or_ServerCapabilities_MonikerProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_MonikerProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case MonikerOptions:
        return json.Marshal(x)
    case MonikerRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, MonikerOptions, MonikerRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_MonikerProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, MonikerOptions, MonikerRegistrationOptions in Or_ServerCapabilities_MonikerProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MonikerOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 MonikerRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, MonikerOptions, MonikerRegistrationOptions in Or_ServerCapabilities_MonikerProvider for value " + stringValue}
}
// Or type for [NotebookDocumentSyncOptions,NotebookDocumentSyncRegistrationOptions]
type Or_ServerCapabilities_NotebookDocumentSync struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_NotebookDocumentSync) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case NotebookDocumentSyncOptions:
        return json.Marshal(x)
    case NotebookDocumentSyncRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of NotebookDocumentSyncOptions, NotebookDocumentSyncRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_NotebookDocumentSync) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentSyncOptions, NotebookDocumentSyncRegistrationOptions in Or_ServerCapabilities_NotebookDocumentSync for value " + stringValue}
	}
   var h0 NotebookDocumentSyncOptions
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 NotebookDocumentSyncRegistrationOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of NotebookDocumentSyncOptions, NotebookDocumentSyncRegistrationOptions in Or_ServerCapabilities_NotebookDocumentSync for value " + stringValue}
}
// Or type for [bool,ReferenceOptions]
type Or_ServerCapabilities_ReferencesProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_ReferencesProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case ReferenceOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, ReferenceOptions", t)
}
func (t *Or_ServerCapabilities_ReferencesProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, ReferenceOptions in Or_ServerCapabilities_ReferencesProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 ReferenceOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, ReferenceOptions in Or_ServerCapabilities_ReferencesProvider for value " + stringValue}
}
// Or type for [bool,RenameOptions]
type Or_ServerCapabilities_RenameProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_RenameProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case RenameOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, RenameOptions", t)
}
func (t *Or_ServerCapabilities_RenameProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, RenameOptions in Or_ServerCapabilities_RenameProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 RenameOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, RenameOptions in Or_ServerCapabilities_RenameProvider for value " + stringValue}
}
// Or type for [bool,SelectionRangeOptions,SelectionRangeRegistrationOptions]
type Or_ServerCapabilities_SelectionRangeProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_SelectionRangeProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case SelectionRangeOptions:
        return json.Marshal(x)
    case SelectionRangeRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, SelectionRangeOptions, SelectionRangeRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_SelectionRangeProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, SelectionRangeOptions, SelectionRangeRegistrationOptions in Or_ServerCapabilities_SelectionRangeProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 SelectionRangeOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 SelectionRangeRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, SelectionRangeOptions, SelectionRangeRegistrationOptions in Or_ServerCapabilities_SelectionRangeProvider for value " + stringValue}
}
// Or type for [SemanticTokensOptions,SemanticTokensRegistrationOptions]
type Or_ServerCapabilities_SemanticTokensProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_SemanticTokensProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case SemanticTokensOptions:
        return json.Marshal(x)
    case SemanticTokensRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of SemanticTokensOptions, SemanticTokensRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_SemanticTokensProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of SemanticTokensOptions, SemanticTokensRegistrationOptions in Or_ServerCapabilities_SemanticTokensProvider for value " + stringValue}
	}
   var h0 SemanticTokensOptions
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 SemanticTokensRegistrationOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of SemanticTokensOptions, SemanticTokensRegistrationOptions in Or_ServerCapabilities_SemanticTokensProvider for value " + stringValue}
}
// Or type for [TextDocumentSyncOptions,TextDocumentSyncKind]
type Or_ServerCapabilities_TextDocumentSync struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_TextDocumentSync) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextDocumentSyncOptions:
        return json.Marshal(x)
    case TextDocumentSyncKind:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextDocumentSyncOptions, TextDocumentSyncKind", t)
}
func (t *Or_ServerCapabilities_TextDocumentSync) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentSyncOptions, TextDocumentSyncKind in Or_ServerCapabilities_TextDocumentSync for value " + stringValue}
	}
   var h0 TextDocumentSyncOptions
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 TextDocumentSyncKind
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentSyncOptions, TextDocumentSyncKind in Or_ServerCapabilities_TextDocumentSync for value " + stringValue}
}
// Or type for [bool,TypeDefinitionOptions,TypeDefinitionRegistrationOptions]
type Or_ServerCapabilities_TypeDefinitionProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_TypeDefinitionProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case TypeDefinitionOptions:
        return json.Marshal(x)
    case TypeDefinitionRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, TypeDefinitionOptions, TypeDefinitionRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_TypeDefinitionProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, TypeDefinitionOptions, TypeDefinitionRegistrationOptions in Or_ServerCapabilities_TypeDefinitionProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 TypeDefinitionOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 TypeDefinitionRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, TypeDefinitionOptions, TypeDefinitionRegistrationOptions in Or_ServerCapabilities_TypeDefinitionProvider for value " + stringValue}
}
// Or type for [bool,TypeHierarchyOptions,TypeHierarchyRegistrationOptions]
type Or_ServerCapabilities_TypeHierarchyProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_TypeHierarchyProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case TypeHierarchyOptions:
        return json.Marshal(x)
    case TypeHierarchyRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, TypeHierarchyOptions, TypeHierarchyRegistrationOptions", t)
}
func (t *Or_ServerCapabilities_TypeHierarchyProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, TypeHierarchyOptions, TypeHierarchyRegistrationOptions in Or_ServerCapabilities_TypeHierarchyProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 TypeHierarchyOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 TypeHierarchyRegistrationOptions
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, TypeHierarchyOptions, TypeHierarchyRegistrationOptions in Or_ServerCapabilities_TypeHierarchyProvider for value " + stringValue}
}
// Or type for [bool,WorkspaceSymbolOptions]
type Or_ServerCapabilities_WorkspaceSymbolProvider struct {
	Value any `json:"value"`
}
func (t *Or_ServerCapabilities_WorkspaceSymbolProvider) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case WorkspaceSymbolOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, WorkspaceSymbolOptions", t)
}
func (t *Or_ServerCapabilities_WorkspaceSymbolProvider) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, WorkspaceSymbolOptions in Or_ServerCapabilities_WorkspaceSymbolProvider for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 WorkspaceSymbolOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, WorkspaceSymbolOptions in Or_ServerCapabilities_WorkspaceSymbolProvider for value " + stringValue}
}
// Or type for [string,MarkupContent]
type Or_SignatureInformation_Documentation struct {
	Value any `json:"value"`
}
func (t *Or_SignatureInformation_Documentation) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case MarkupContent:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, MarkupContent", t)
}
func (t *Or_SignatureInformation_Documentation) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_SignatureInformation_Documentation for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 MarkupContent
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, MarkupContent in Or_SignatureInformation_Documentation for value " + stringValue}
}
// Or type for [TextEdit,AnnotatedTextEdit,SnippetTextEdit]
type Or_TextDocumentEdit_Edits struct {
	Value any `json:"value"`
}
func (t *Or_TextDocumentEdit_Edits) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextEdit:
        return json.Marshal(x)
    case AnnotatedTextEdit:
        return json.Marshal(x)
    case SnippetTextEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextEdit, AnnotatedTextEdit, SnippetTextEdit", t)
}
func (t *Or_TextDocumentEdit_Edits) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextEdit, AnnotatedTextEdit, SnippetTextEdit in Or_TextDocumentEdit_Edits for value " + stringValue}
	}
   var h0 TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 AnnotatedTextEdit
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 SnippetTextEdit
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextEdit, AnnotatedTextEdit, SnippetTextEdit in Or_TextDocumentEdit_Edits for value " + stringValue}
}
// Or type for [bool,SaveOptions]
type Or_TextDocumentSyncOptions_Save struct {
	Value any `json:"value"`
}
func (t *Or_TextDocumentSyncOptions_Save) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case bool:
        return json.Marshal(x)
    case SaveOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of bool, SaveOptions", t)
}
func (t *Or_TextDocumentSyncOptions_Save) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of bool, SaveOptions in Or_TextDocumentSyncOptions_Save for value " + stringValue}
	}
   var h0 bool
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 SaveOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of bool, SaveOptions in Or_TextDocumentSyncOptions_Save for value " + stringValue}
}
// Or type for [TextDocumentEdit,CreateFile,RenameFile,DeleteFile]
type Or_WorkspaceEdit_DocumentChanges struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceEdit_DocumentChanges) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextDocumentEdit:
        return json.Marshal(x)
    case CreateFile:
        return json.Marshal(x)
    case RenameFile:
        return json.Marshal(x)
    case DeleteFile:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextDocumentEdit, CreateFile, RenameFile, DeleteFile", t)
}
func (t *Or_WorkspaceEdit_DocumentChanges) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentEdit, CreateFile, RenameFile, DeleteFile in Or_WorkspaceEdit_DocumentChanges for value " + stringValue}
	}
   var h0 TextDocumentEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 CreateFile
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h2 RenameFile
   if err := decoder.Decode(&h2); err == nil {
       t.Value = h2
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h3 DeleteFile
   if err := decoder.Decode(&h3); err == nil {
       t.Value = h3
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentEdit, CreateFile, RenameFile, DeleteFile in Or_WorkspaceEdit_DocumentChanges for value " + stringValue}
}
// Or type for [[]WorkspaceFolder,nil]
type Or_WorkspaceFoldersInitializeParams_WorkspaceFolders struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceFoldersInitializeParams_WorkspaceFolders) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []WorkspaceFolder:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []WorkspaceFolder, nil", t)
}
func (t *Or_WorkspaceFoldersInitializeParams_WorkspaceFolders) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []WorkspaceFolder
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []WorkspaceFolder, nil in Or_WorkspaceFoldersInitializeParams_WorkspaceFolders for value " + stringValue}
}
// Or type for [string,bool]
type Or_WorkspaceFoldersServerCapabilities_ChangeNotifications struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceFoldersServerCapabilities_ChangeNotifications) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case string:
        return json.Marshal(x)
    case bool:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of string, bool", t)
}
func (t *Or_WorkspaceFoldersServerCapabilities_ChangeNotifications) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of string, bool in Or_WorkspaceFoldersServerCapabilities_ChangeNotifications for value " + stringValue}
	}
   var h0 string
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 bool
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of string, bool in Or_WorkspaceFoldersServerCapabilities_ChangeNotifications for value " + stringValue}
}
// Or type for [TextDocumentContentOptions,TextDocumentContentRegistrationOptions]
type Or_WorkspaceOptions_TextDocumentContent struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceOptions_TextDocumentContent) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case TextDocumentContentOptions:
        return json.Marshal(x)
    case TextDocumentContentRegistrationOptions:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of TextDocumentContentOptions, TextDocumentContentRegistrationOptions", t)
}
func (t *Or_WorkspaceOptions_TextDocumentContent) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentContentOptions, TextDocumentContentRegistrationOptions in Or_WorkspaceOptions_TextDocumentContent for value " + stringValue}
	}
   var h0 TextDocumentContentOptions
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 TextDocumentContentRegistrationOptions
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of TextDocumentContentOptions, TextDocumentContentRegistrationOptions in Or_WorkspaceOptions_TextDocumentContent for value " + stringValue}
}
// Or type for [Location,LocationUriOnly]
type Or_WorkspaceSymbol_Location struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceSymbol_Location) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Location:
        return json.Marshal(x)
    case LocationUriOnly:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Location, LocationUriOnly", t)
}
func (t *Or_WorkspaceSymbol_Location) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of Location, LocationUriOnly in Or_WorkspaceSymbol_Location for value " + stringValue}
	}
   var h0 Location
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 LocationUriOnly
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Location, LocationUriOnly in Or_WorkspaceSymbol_Location for value " + stringValue}
}
// Or type for [int32,string]
type Or_Request_id struct {
	Value any `json:"value"`
}
func (t *Or_Request_id) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case int32:
        return json.Marshal(x)
    case string:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of int32, string", t)
}
func (t *Or_Request_id) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_Request_id for value " + stringValue}
	}
   var h0 int32
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 string
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_Request_id for value " + stringValue}
}
// Or type for [int32,string]
type Or_Response_id struct {
	Value any `json:"value"`
}
func (t *Or_Response_id) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case int32:
        return json.Marshal(x)
    case string:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of int32, string", t)
}
func (t *Or_Response_id) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
	if string(x) == "null" {
		return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_Response_id for value " + stringValue}
	}
   var h0 int32
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 string
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of int32, string in Or_Response_id for value " + stringValue}
}
// Or type for [[]CallHierarchyIncomingCall,nil]
type Or_CallHierarchyIncomingCallsResponse struct {
	Value any `json:"value"`
}
func (t *Or_CallHierarchyIncomingCallsResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []CallHierarchyIncomingCall:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []CallHierarchyIncomingCall, nil", t)
}
func (t *Or_CallHierarchyIncomingCallsResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []CallHierarchyIncomingCall
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []CallHierarchyIncomingCall, nil in Or_CallHierarchyIncomingCallsResponse for value " + stringValue}
}
// Or type for [[]CallHierarchyOutgoingCall,nil]
type Or_CallHierarchyOutgoingCallsResponse struct {
	Value any `json:"value"`
}
func (t *Or_CallHierarchyOutgoingCallsResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []CallHierarchyOutgoingCall:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []CallHierarchyOutgoingCall, nil", t)
}
func (t *Or_CallHierarchyOutgoingCallsResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []CallHierarchyOutgoingCall
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []CallHierarchyOutgoingCall, nil in Or_CallHierarchyOutgoingCallsResponse for value " + stringValue}
}
// Or type for [[]Or_CodeActionResponse,nil]
type Or_CodeActionResponse struct {
	Value any `json:"value"`
}
func (t *Or_CodeActionResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []Or_CodeActionResponse:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []Or_CodeActionResponse, nil", t)
}
func (t *Or_CodeActionResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []Or_CodeActionResponse
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []Or_CodeActionResponse, nil in Or_CodeActionResponse for value " + stringValue}
}
// Or type for [[]CodeLens,nil]
type Or_CodeLensResponse struct {
	Value any `json:"value"`
}
func (t *Or_CodeLensResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []CodeLens:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []CodeLens, nil", t)
}
func (t *Or_CodeLensResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []CodeLens
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []CodeLens, nil in Or_CodeLensResponse for value " + stringValue}
}
// Or type for [[]CompletionItem,CompletionList,nil]
type Or_CompletionResponse struct {
	Value any `json:"value"`
}
func (t *Or_CompletionResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []CompletionItem:
        return json.Marshal(x)
    case CompletionList:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []CompletionItem, CompletionList, nil", t)
}
func (t *Or_CompletionResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []CompletionItem
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 CompletionList
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []CompletionItem, CompletionList, nil in Or_CompletionResponse for value " + stringValue}
}
// Or type for [Declaration,[]DeclarationLink,nil]
type Or_DeclarationResponse struct {
	Value any `json:"value"`
}
func (t *Or_DeclarationResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Declaration:
        return json.Marshal(x)
    case []DeclarationLink:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Declaration, []DeclarationLink, nil", t)
}
func (t *Or_DeclarationResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 Declaration
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []DeclarationLink
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Declaration, []DeclarationLink, nil in Or_DeclarationResponse for value " + stringValue}
}
// Or type for [Definition,[]DefinitionLink,nil]
type Or_DefinitionResponse struct {
	Value any `json:"value"`
}
func (t *Or_DefinitionResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Definition:
        return json.Marshal(x)
    case []DefinitionLink:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Definition, []DefinitionLink, nil", t)
}
func (t *Or_DefinitionResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 Definition
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []DefinitionLink
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Definition, []DefinitionLink, nil in Or_DefinitionResponse for value " + stringValue}
}
// Or type for [[]DocumentHighlight,nil]
type Or_DocumentHighlightResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentHighlightResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []DocumentHighlight:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []DocumentHighlight, nil", t)
}
func (t *Or_DocumentHighlightResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []DocumentHighlight
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []DocumentHighlight, nil in Or_DocumentHighlightResponse for value " + stringValue}
}
// Or type for [[]DocumentLink,nil]
type Or_DocumentLinkResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentLinkResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []DocumentLink:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []DocumentLink, nil", t)
}
func (t *Or_DocumentLinkResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []DocumentLink
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []DocumentLink, nil in Or_DocumentLinkResponse for value " + stringValue}
}
// Or type for [[]SymbolInformation,[]DocumentSymbol,nil]
type Or_DocumentSymbolResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentSymbolResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []SymbolInformation:
        return json.Marshal(x)
    case []DocumentSymbol:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []SymbolInformation, []DocumentSymbol, nil", t)
}
func (t *Or_DocumentSymbolResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []SymbolInformation
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []DocumentSymbol
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []SymbolInformation, []DocumentSymbol, nil in Or_DocumentSymbolResponse for value " + stringValue}
}
// Or type for [[]FoldingRange,nil]
type Or_FoldingRangeResponse struct {
	Value any `json:"value"`
}
func (t *Or_FoldingRangeResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []FoldingRange:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []FoldingRange, nil", t)
}
func (t *Or_FoldingRangeResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []FoldingRange
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []FoldingRange, nil in Or_FoldingRangeResponse for value " + stringValue}
}
// Or type for [[]TextEdit,nil]
type Or_DocumentFormattingResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentFormattingResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TextEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TextEdit, nil", t)
}
func (t *Or_DocumentFormattingResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TextEdit, nil in Or_DocumentFormattingResponse for value " + stringValue}
}
// Or type for [Definition,[]DefinitionLink,nil]
type Or_ImplementationResponse struct {
	Value any `json:"value"`
}
func (t *Or_ImplementationResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Definition:
        return json.Marshal(x)
    case []DefinitionLink:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Definition, []DefinitionLink, nil", t)
}
func (t *Or_ImplementationResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 Definition
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []DefinitionLink
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Definition, []DefinitionLink, nil in Or_ImplementationResponse for value " + stringValue}
}
// Or type for [[]InlayHint,nil]
type Or_InlayHintResponse struct {
	Value any `json:"value"`
}
func (t *Or_InlayHintResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []InlayHint:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []InlayHint, nil", t)
}
func (t *Or_InlayHintResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []InlayHint
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []InlayHint, nil in Or_InlayHintResponse for value " + stringValue}
}
// Or type for [InlineCompletionList,[]InlineCompletionItem,nil]
type Or_InlineCompletionResponse struct {
	Value any `json:"value"`
}
func (t *Or_InlineCompletionResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case InlineCompletionList:
        return json.Marshal(x)
    case []InlineCompletionItem:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of InlineCompletionList, []InlineCompletionItem, nil", t)
}
func (t *Or_InlineCompletionResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 InlineCompletionList
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []InlineCompletionItem
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of InlineCompletionList, []InlineCompletionItem, nil in Or_InlineCompletionResponse for value " + stringValue}
}
// Or type for [[]InlineValue,nil]
type Or_InlineValueResponse struct {
	Value any `json:"value"`
}
func (t *Or_InlineValueResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []InlineValue:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []InlineValue, nil", t)
}
func (t *Or_InlineValueResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []InlineValue
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []InlineValue, nil in Or_InlineValueResponse for value " + stringValue}
}
// Or type for [[]Moniker,nil]
type Or_MonikerResponse struct {
	Value any `json:"value"`
}
func (t *Or_MonikerResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []Moniker:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []Moniker, nil", t)
}
func (t *Or_MonikerResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []Moniker
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []Moniker, nil in Or_MonikerResponse for value " + stringValue}
}
// Or type for [[]TextEdit,nil]
type Or_DocumentOnTypeFormattingResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentOnTypeFormattingResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TextEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TextEdit, nil", t)
}
func (t *Or_DocumentOnTypeFormattingResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TextEdit, nil in Or_DocumentOnTypeFormattingResponse for value " + stringValue}
}
// Or type for [[]CallHierarchyItem,nil]
type Or_CallHierarchyPrepareResponse struct {
	Value any `json:"value"`
}
func (t *Or_CallHierarchyPrepareResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []CallHierarchyItem:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []CallHierarchyItem, nil", t)
}
func (t *Or_CallHierarchyPrepareResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []CallHierarchyItem
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []CallHierarchyItem, nil in Or_CallHierarchyPrepareResponse for value " + stringValue}
}
// Or type for [[]TypeHierarchyItem,nil]
type Or_TypeHierarchyPrepareResponse struct {
	Value any `json:"value"`
}
func (t *Or_TypeHierarchyPrepareResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TypeHierarchyItem:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TypeHierarchyItem, nil", t)
}
func (t *Or_TypeHierarchyPrepareResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TypeHierarchyItem
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TypeHierarchyItem, nil in Or_TypeHierarchyPrepareResponse for value " + stringValue}
}
// Or type for [[]TextEdit,nil]
type Or_DocumentRangeFormattingResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentRangeFormattingResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TextEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TextEdit, nil", t)
}
func (t *Or_DocumentRangeFormattingResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TextEdit, nil in Or_DocumentRangeFormattingResponse for value " + stringValue}
}
// Or type for [[]TextEdit,nil]
type Or_DocumentRangesFormattingResponse struct {
	Value any `json:"value"`
}
func (t *Or_DocumentRangesFormattingResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TextEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TextEdit, nil", t)
}
func (t *Or_DocumentRangesFormattingResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TextEdit, nil in Or_DocumentRangesFormattingResponse for value " + stringValue}
}
// Or type for [[]Location,nil]
type Or_ReferencesResponse struct {
	Value any `json:"value"`
}
func (t *Or_ReferencesResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []Location:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []Location, nil", t)
}
func (t *Or_ReferencesResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []Location
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []Location, nil in Or_ReferencesResponse for value " + stringValue}
}
// Or type for [[]SelectionRange,nil]
type Or_SelectionRangeResponse struct {
	Value any `json:"value"`
}
func (t *Or_SelectionRangeResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []SelectionRange:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []SelectionRange, nil", t)
}
func (t *Or_SelectionRangeResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []SelectionRange
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []SelectionRange, nil in Or_SelectionRangeResponse for value " + stringValue}
}
// Or type for [SemanticTokens,SemanticTokensDelta,nil]
type Or_SemanticTokensDeltaResponse struct {
	Value any `json:"value"`
}
func (t *Or_SemanticTokensDeltaResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case SemanticTokens:
        return json.Marshal(x)
    case SemanticTokensDelta:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of SemanticTokens, SemanticTokensDelta, nil", t)
}
func (t *Or_SemanticTokensDeltaResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 SemanticTokens
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 SemanticTokensDelta
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of SemanticTokens, SemanticTokensDelta, nil in Or_SemanticTokensDeltaResponse for value " + stringValue}
}
// Or type for [Definition,[]DefinitionLink,nil]
type Or_TypeDefinitionResponse struct {
	Value any `json:"value"`
}
func (t *Or_TypeDefinitionResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case Definition:
        return json.Marshal(x)
    case []DefinitionLink:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of Definition, []DefinitionLink, nil", t)
}
func (t *Or_TypeDefinitionResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 Definition
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []DefinitionLink
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of Definition, []DefinitionLink, nil in Or_TypeDefinitionResponse for value " + stringValue}
}
// Or type for [[]TextEdit,nil]
type Or_WillSaveTextDocumentWaitUntilResponse struct {
	Value any `json:"value"`
}
func (t *Or_WillSaveTextDocumentWaitUntilResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TextEdit:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TextEdit, nil", t)
}
func (t *Or_WillSaveTextDocumentWaitUntilResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TextEdit
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TextEdit, nil in Or_WillSaveTextDocumentWaitUntilResponse for value " + stringValue}
}
// Or type for [[]TypeHierarchyItem,nil]
type Or_TypeHierarchySubtypesResponse struct {
	Value any `json:"value"`
}
func (t *Or_TypeHierarchySubtypesResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TypeHierarchyItem:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TypeHierarchyItem, nil", t)
}
func (t *Or_TypeHierarchySubtypesResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TypeHierarchyItem
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TypeHierarchyItem, nil in Or_TypeHierarchySubtypesResponse for value " + stringValue}
}
// Or type for [[]TypeHierarchyItem,nil]
type Or_TypeHierarchySupertypesResponse struct {
	Value any `json:"value"`
}
func (t *Or_TypeHierarchySupertypesResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []TypeHierarchyItem:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []TypeHierarchyItem, nil", t)
}
func (t *Or_TypeHierarchySupertypesResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []TypeHierarchyItem
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []TypeHierarchyItem, nil in Or_TypeHierarchySupertypesResponse for value " + stringValue}
}
// Or type for [[]SymbolInformation,[]WorkspaceSymbol,nil]
type Or_WorkspaceSymbolResponse struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceSymbolResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []SymbolInformation:
        return json.Marshal(x)
    case []WorkspaceSymbol:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []SymbolInformation, []WorkspaceSymbol, nil", t)
}
func (t *Or_WorkspaceSymbolResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []SymbolInformation
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	decoder = json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   var h1 []WorkspaceSymbol
   if err := decoder.Decode(&h1); err == nil {
       t.Value = h1
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []SymbolInformation, []WorkspaceSymbol, nil in Or_WorkspaceSymbolResponse for value " + stringValue}
}
// Or type for [[]WorkspaceFolder,nil]
type Or_WorkspaceFoldersResponse struct {
	Value any `json:"value"`
}
func (t *Or_WorkspaceFoldersResponse) MarshalJSON() ([]byte, error) {
   switch x := t.Value.(type) {
   case nil:
       return []byte("null"), nil
    case []WorkspaceFolder:
        return json.Marshal(x)
   }
  return nil, fmt.Errorf("unknown type %T is not one of []WorkspaceFolder, nil", t)
}
func (t *Or_WorkspaceFoldersResponse) UnmarshalJSON(x []byte) error {
	stringValue := string(x)
	decoder := json.NewDecoder(bytes.NewReader(x))
	decoder.DisallowUnknownFields()
   if string(x) == "null" {
       t.Value = nil
       return nil
   }
   var h0 []WorkspaceFolder
   if err := decoder.Decode(&h0); err == nil {
       t.Value = h0
       return nil
   }
	return &UnmarshalError{msg: "unmarshal failed to match one of []WorkspaceFolder, nil in Or_WorkspaceFoldersResponse for value " + stringValue}
}