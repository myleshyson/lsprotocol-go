package protocol

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"bytes"
)

func TestType(t *testing.T) {

	cwd, err := os.Getwd()

	if err != nil {
		t.Error(err)
	}

	files, err := os.ReadDir(cwd + "/testdata")

	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		filepath := cwd + "/testdata/" + file.Name()
		if !strings.Contains(file.Name(), "json") {
			continue
		}

		parts := strings.Split(file.Name(), "-")

		testType := parts[0]
		expectedResult := parts[1]
		content, err := os.ReadFile(filepath)

		if err != nil {
			t.Error(err)
		}

		switch testType {
		
		case "ImplementationRequest":
		validateType[ImplementationRequest](t, content, expectedResult, filepath)
		case "ImplementationResponse":
		validateType[ImplementationResponse](t, content, expectedResult, filepath)
		case "TypeDefinitionRequest":
		validateType[TypeDefinitionRequest](t, content, expectedResult, filepath)
		case "TypeDefinitionResponse":
		validateType[TypeDefinitionResponse](t, content, expectedResult, filepath)
		case "WorkspaceFoldersRequest":
		validateType[WorkspaceFoldersRequest](t, content, expectedResult, filepath)
		case "WorkspaceFoldersResponse":
		validateType[WorkspaceFoldersResponse](t, content, expectedResult, filepath)
		case "ConfigurationRequest":
		validateType[ConfigurationRequest](t, content, expectedResult, filepath)
		case "ConfigurationResponse":
		validateType[ConfigurationResponse](t, content, expectedResult, filepath)
		case "DocumentColorRequest":
		validateType[DocumentColorRequest](t, content, expectedResult, filepath)
		case "DocumentColorResponse":
		validateType[DocumentColorResponse](t, content, expectedResult, filepath)
		case "ColorPresentationRequest":
		validateType[ColorPresentationRequest](t, content, expectedResult, filepath)
		case "ColorPresentationResponse":
		validateType[ColorPresentationResponse](t, content, expectedResult, filepath)
		case "FoldingRangeRequest":
		validateType[FoldingRangeRequest](t, content, expectedResult, filepath)
		case "FoldingRangeResponse":
		validateType[FoldingRangeResponse](t, content, expectedResult, filepath)
		case "FoldingRangeRefreshRequest":
		validateType[FoldingRangeRefreshRequest](t, content, expectedResult, filepath)
		case "FoldingRangeRefreshResponse":
		validateType[FoldingRangeRefreshResponse](t, content, expectedResult, filepath)
		case "DeclarationRequest":
		validateType[DeclarationRequest](t, content, expectedResult, filepath)
		case "DeclarationResponse":
		validateType[DeclarationResponse](t, content, expectedResult, filepath)
		case "SelectionRangeRequest":
		validateType[SelectionRangeRequest](t, content, expectedResult, filepath)
		case "SelectionRangeResponse":
		validateType[SelectionRangeResponse](t, content, expectedResult, filepath)
		case "WorkDoneProgressCreateRequest":
		validateType[WorkDoneProgressCreateRequest](t, content, expectedResult, filepath)
		case "WorkDoneProgressCreateResponse":
		validateType[WorkDoneProgressCreateResponse](t, content, expectedResult, filepath)
		case "CallHierarchyPrepareRequest":
		validateType[CallHierarchyPrepareRequest](t, content, expectedResult, filepath)
		case "CallHierarchyPrepareResponse":
		validateType[CallHierarchyPrepareResponse](t, content, expectedResult, filepath)
		case "CallHierarchyIncomingCallsRequest":
		validateType[CallHierarchyIncomingCallsRequest](t, content, expectedResult, filepath)
		case "CallHierarchyIncomingCallsResponse":
		validateType[CallHierarchyIncomingCallsResponse](t, content, expectedResult, filepath)
		case "CallHierarchyOutgoingCallsRequest":
		validateType[CallHierarchyOutgoingCallsRequest](t, content, expectedResult, filepath)
		case "CallHierarchyOutgoingCallsResponse":
		validateType[CallHierarchyOutgoingCallsResponse](t, content, expectedResult, filepath)
		case "SemanticTokensRequest":
		validateType[SemanticTokensRequest](t, content, expectedResult, filepath)
		case "SemanticTokensResponse":
		validateType[SemanticTokensResponse](t, content, expectedResult, filepath)
		case "SemanticTokensDeltaRequest":
		validateType[SemanticTokensDeltaRequest](t, content, expectedResult, filepath)
		case "SemanticTokensDeltaResponse":
		validateType[SemanticTokensDeltaResponse](t, content, expectedResult, filepath)
		case "SemanticTokensRangeRequest":
		validateType[SemanticTokensRangeRequest](t, content, expectedResult, filepath)
		case "SemanticTokensRangeResponse":
		validateType[SemanticTokensRangeResponse](t, content, expectedResult, filepath)
		case "SemanticTokensRefreshRequest":
		validateType[SemanticTokensRefreshRequest](t, content, expectedResult, filepath)
		case "SemanticTokensRefreshResponse":
		validateType[SemanticTokensRefreshResponse](t, content, expectedResult, filepath)
		case "ShowDocumentRequest":
		validateType[ShowDocumentRequest](t, content, expectedResult, filepath)
		case "ShowDocumentResponse":
		validateType[ShowDocumentResponse](t, content, expectedResult, filepath)
		case "LinkedEditingRangeRequest":
		validateType[LinkedEditingRangeRequest](t, content, expectedResult, filepath)
		case "LinkedEditingRangeResponse":
		validateType[LinkedEditingRangeResponse](t, content, expectedResult, filepath)
		case "WillCreateFilesRequest":
		validateType[WillCreateFilesRequest](t, content, expectedResult, filepath)
		case "WillCreateFilesResponse":
		validateType[WillCreateFilesResponse](t, content, expectedResult, filepath)
		case "WillRenameFilesRequest":
		validateType[WillRenameFilesRequest](t, content, expectedResult, filepath)
		case "WillRenameFilesResponse":
		validateType[WillRenameFilesResponse](t, content, expectedResult, filepath)
		case "WillDeleteFilesRequest":
		validateType[WillDeleteFilesRequest](t, content, expectedResult, filepath)
		case "WillDeleteFilesResponse":
		validateType[WillDeleteFilesResponse](t, content, expectedResult, filepath)
		case "MonikerRequest":
		validateType[MonikerRequest](t, content, expectedResult, filepath)
		case "MonikerResponse":
		validateType[MonikerResponse](t, content, expectedResult, filepath)
		case "TypeHierarchyPrepareRequest":
		validateType[TypeHierarchyPrepareRequest](t, content, expectedResult, filepath)
		case "TypeHierarchyPrepareResponse":
		validateType[TypeHierarchyPrepareResponse](t, content, expectedResult, filepath)
		case "TypeHierarchySupertypesRequest":
		validateType[TypeHierarchySupertypesRequest](t, content, expectedResult, filepath)
		case "TypeHierarchySupertypesResponse":
		validateType[TypeHierarchySupertypesResponse](t, content, expectedResult, filepath)
		case "TypeHierarchySubtypesRequest":
		validateType[TypeHierarchySubtypesRequest](t, content, expectedResult, filepath)
		case "TypeHierarchySubtypesResponse":
		validateType[TypeHierarchySubtypesResponse](t, content, expectedResult, filepath)
		case "InlineValueRequest":
		validateType[InlineValueRequest](t, content, expectedResult, filepath)
		case "InlineValueResponse":
		validateType[InlineValueResponse](t, content, expectedResult, filepath)
		case "InlineValueRefreshRequest":
		validateType[InlineValueRefreshRequest](t, content, expectedResult, filepath)
		case "InlineValueRefreshResponse":
		validateType[InlineValueRefreshResponse](t, content, expectedResult, filepath)
		case "InlayHintRequest":
		validateType[InlayHintRequest](t, content, expectedResult, filepath)
		case "InlayHintResponse":
		validateType[InlayHintResponse](t, content, expectedResult, filepath)
		case "InlayHintResolveRequest":
		validateType[InlayHintResolveRequest](t, content, expectedResult, filepath)
		case "InlayHintResolveResponse":
		validateType[InlayHintResolveResponse](t, content, expectedResult, filepath)
		case "InlayHintRefreshRequest":
		validateType[InlayHintRefreshRequest](t, content, expectedResult, filepath)
		case "InlayHintRefreshResponse":
		validateType[InlayHintRefreshResponse](t, content, expectedResult, filepath)
		case "DocumentDiagnosticRequest":
		validateType[DocumentDiagnosticRequest](t, content, expectedResult, filepath)
		case "DocumentDiagnosticResponse":
		validateType[DocumentDiagnosticResponse](t, content, expectedResult, filepath)
		case "WorkspaceDiagnosticRequest":
		validateType[WorkspaceDiagnosticRequest](t, content, expectedResult, filepath)
		case "WorkspaceDiagnosticResponse":
		validateType[WorkspaceDiagnosticResponse](t, content, expectedResult, filepath)
		case "DiagnosticRefreshRequest":
		validateType[DiagnosticRefreshRequest](t, content, expectedResult, filepath)
		case "DiagnosticRefreshResponse":
		validateType[DiagnosticRefreshResponse](t, content, expectedResult, filepath)
		case "InlineCompletionRequest":
		validateType[InlineCompletionRequest](t, content, expectedResult, filepath)
		case "InlineCompletionResponse":
		validateType[InlineCompletionResponse](t, content, expectedResult, filepath)
		case "TextDocumentContentRequest":
		validateType[TextDocumentContentRequest](t, content, expectedResult, filepath)
		case "TextDocumentContentResponse":
		validateType[TextDocumentContentResponse](t, content, expectedResult, filepath)
		case "TextDocumentContentRefreshRequest":
		validateType[TextDocumentContentRefreshRequest](t, content, expectedResult, filepath)
		case "TextDocumentContentRefreshResponse":
		validateType[TextDocumentContentRefreshResponse](t, content, expectedResult, filepath)
		case "RegistrationRequest":
		validateType[RegistrationRequest](t, content, expectedResult, filepath)
		case "RegistrationResponse":
		validateType[RegistrationResponse](t, content, expectedResult, filepath)
		case "UnregistrationRequest":
		validateType[UnregistrationRequest](t, content, expectedResult, filepath)
		case "UnregistrationResponse":
		validateType[UnregistrationResponse](t, content, expectedResult, filepath)
		case "InitializeRequest":
		validateType[InitializeRequest](t, content, expectedResult, filepath)
		case "InitializeResponse":
		validateType[InitializeResponse](t, content, expectedResult, filepath)
		case "ShutdownRequest":
		validateType[ShutdownRequest](t, content, expectedResult, filepath)
		case "ShutdownResponse":
		validateType[ShutdownResponse](t, content, expectedResult, filepath)
		case "ShowMessageRequest":
		validateType[ShowMessageRequest](t, content, expectedResult, filepath)
		case "ShowMessageResponse":
		validateType[ShowMessageResponse](t, content, expectedResult, filepath)
		case "WillSaveTextDocumentWaitUntilRequest":
		validateType[WillSaveTextDocumentWaitUntilRequest](t, content, expectedResult, filepath)
		case "WillSaveTextDocumentWaitUntilResponse":
		validateType[WillSaveTextDocumentWaitUntilResponse](t, content, expectedResult, filepath)
		case "CompletionRequest":
		validateType[CompletionRequest](t, content, expectedResult, filepath)
		case "CompletionResponse":
		validateType[CompletionResponse](t, content, expectedResult, filepath)
		case "CompletionResolveRequest":
		validateType[CompletionResolveRequest](t, content, expectedResult, filepath)
		case "CompletionResolveResponse":
		validateType[CompletionResolveResponse](t, content, expectedResult, filepath)
		case "HoverRequest":
		validateType[HoverRequest](t, content, expectedResult, filepath)
		case "HoverResponse":
		validateType[HoverResponse](t, content, expectedResult, filepath)
		case "SignatureHelpRequest":
		validateType[SignatureHelpRequest](t, content, expectedResult, filepath)
		case "SignatureHelpResponse":
		validateType[SignatureHelpResponse](t, content, expectedResult, filepath)
		case "DefinitionRequest":
		validateType[DefinitionRequest](t, content, expectedResult, filepath)
		case "DefinitionResponse":
		validateType[DefinitionResponse](t, content, expectedResult, filepath)
		case "ReferencesRequest":
		validateType[ReferencesRequest](t, content, expectedResult, filepath)
		case "ReferencesResponse":
		validateType[ReferencesResponse](t, content, expectedResult, filepath)
		case "DocumentHighlightRequest":
		validateType[DocumentHighlightRequest](t, content, expectedResult, filepath)
		case "DocumentHighlightResponse":
		validateType[DocumentHighlightResponse](t, content, expectedResult, filepath)
		case "DocumentSymbolRequest":
		validateType[DocumentSymbolRequest](t, content, expectedResult, filepath)
		case "DocumentSymbolResponse":
		validateType[DocumentSymbolResponse](t, content, expectedResult, filepath)
		case "CodeActionRequest":
		validateType[CodeActionRequest](t, content, expectedResult, filepath)
		case "CodeActionResponse":
		validateType[CodeActionResponse](t, content, expectedResult, filepath)
		case "CodeActionResolveRequest":
		validateType[CodeActionResolveRequest](t, content, expectedResult, filepath)
		case "CodeActionResolveResponse":
		validateType[CodeActionResolveResponse](t, content, expectedResult, filepath)
		case "WorkspaceSymbolRequest":
		validateType[WorkspaceSymbolRequest](t, content, expectedResult, filepath)
		case "WorkspaceSymbolResponse":
		validateType[WorkspaceSymbolResponse](t, content, expectedResult, filepath)
		case "WorkspaceSymbolResolveRequest":
		validateType[WorkspaceSymbolResolveRequest](t, content, expectedResult, filepath)
		case "WorkspaceSymbolResolveResponse":
		validateType[WorkspaceSymbolResolveResponse](t, content, expectedResult, filepath)
		case "CodeLensRequest":
		validateType[CodeLensRequest](t, content, expectedResult, filepath)
		case "CodeLensResponse":
		validateType[CodeLensResponse](t, content, expectedResult, filepath)
		case "CodeLensResolveRequest":
		validateType[CodeLensResolveRequest](t, content, expectedResult, filepath)
		case "CodeLensResolveResponse":
		validateType[CodeLensResolveResponse](t, content, expectedResult, filepath)
		case "CodeLensRefreshRequest":
		validateType[CodeLensRefreshRequest](t, content, expectedResult, filepath)
		case "CodeLensRefreshResponse":
		validateType[CodeLensRefreshResponse](t, content, expectedResult, filepath)
		case "DocumentLinkRequest":
		validateType[DocumentLinkRequest](t, content, expectedResult, filepath)
		case "DocumentLinkResponse":
		validateType[DocumentLinkResponse](t, content, expectedResult, filepath)
		case "DocumentLinkResolveRequest":
		validateType[DocumentLinkResolveRequest](t, content, expectedResult, filepath)
		case "DocumentLinkResolveResponse":
		validateType[DocumentLinkResolveResponse](t, content, expectedResult, filepath)
		case "DocumentFormattingRequest":
		validateType[DocumentFormattingRequest](t, content, expectedResult, filepath)
		case "DocumentFormattingResponse":
		validateType[DocumentFormattingResponse](t, content, expectedResult, filepath)
		case "DocumentRangeFormattingRequest":
		validateType[DocumentRangeFormattingRequest](t, content, expectedResult, filepath)
		case "DocumentRangeFormattingResponse":
		validateType[DocumentRangeFormattingResponse](t, content, expectedResult, filepath)
		case "DocumentRangesFormattingRequest":
		validateType[DocumentRangesFormattingRequest](t, content, expectedResult, filepath)
		case "DocumentRangesFormattingResponse":
		validateType[DocumentRangesFormattingResponse](t, content, expectedResult, filepath)
		case "DocumentOnTypeFormattingRequest":
		validateType[DocumentOnTypeFormattingRequest](t, content, expectedResult, filepath)
		case "DocumentOnTypeFormattingResponse":
		validateType[DocumentOnTypeFormattingResponse](t, content, expectedResult, filepath)
		case "RenameRequest":
		validateType[RenameRequest](t, content, expectedResult, filepath)
		case "RenameResponse":
		validateType[RenameResponse](t, content, expectedResult, filepath)
		case "PrepareRenameRequest":
		validateType[PrepareRenameRequest](t, content, expectedResult, filepath)
		case "PrepareRenameResponse":
		validateType[PrepareRenameResponse](t, content, expectedResult, filepath)
		case "ExecuteCommandRequest":
		validateType[ExecuteCommandRequest](t, content, expectedResult, filepath)
		case "ExecuteCommandResponse":
		validateType[ExecuteCommandResponse](t, content, expectedResult, filepath)
		case "ApplyWorkspaceEditRequest":
		validateType[ApplyWorkspaceEditRequest](t, content, expectedResult, filepath)
		case "ApplyWorkspaceEditResponse":
		validateType[ApplyWorkspaceEditResponse](t, content, expectedResult, filepath)
		case "DidChangeWorkspaceFoldersNotification":
		validateType[DidChangeWorkspaceFoldersNotification](t, content, expectedResult, filepath)
		case "WorkDoneProgressCancelNotification":
		validateType[WorkDoneProgressCancelNotification](t, content, expectedResult, filepath)
		case "DidCreateFilesNotification":
		validateType[DidCreateFilesNotification](t, content, expectedResult, filepath)
		case "DidRenameFilesNotification":
		validateType[DidRenameFilesNotification](t, content, expectedResult, filepath)
		case "DidDeleteFilesNotification":
		validateType[DidDeleteFilesNotification](t, content, expectedResult, filepath)
		case "DidOpenNotebookDocumentNotification":
		validateType[DidOpenNotebookDocumentNotification](t, content, expectedResult, filepath)
		case "DidChangeNotebookDocumentNotification":
		validateType[DidChangeNotebookDocumentNotification](t, content, expectedResult, filepath)
		case "DidSaveNotebookDocumentNotification":
		validateType[DidSaveNotebookDocumentNotification](t, content, expectedResult, filepath)
		case "DidCloseNotebookDocumentNotification":
		validateType[DidCloseNotebookDocumentNotification](t, content, expectedResult, filepath)
		case "InitializedNotification":
		validateType[InitializedNotification](t, content, expectedResult, filepath)
		case "ExitNotification":
		validateType[ExitNotification](t, content, expectedResult, filepath)
		case "DidChangeConfigurationNotification":
		validateType[DidChangeConfigurationNotification](t, content, expectedResult, filepath)
		case "ShowMessageNotification":
		validateType[ShowMessageNotification](t, content, expectedResult, filepath)
		case "LogMessageNotification":
		validateType[LogMessageNotification](t, content, expectedResult, filepath)
		case "TelemetryEventNotification":
		validateType[TelemetryEventNotification](t, content, expectedResult, filepath)
		case "DidOpenTextDocumentNotification":
		validateType[DidOpenTextDocumentNotification](t, content, expectedResult, filepath)
		case "DidChangeTextDocumentNotification":
		validateType[DidChangeTextDocumentNotification](t, content, expectedResult, filepath)
		case "DidCloseTextDocumentNotification":
		validateType[DidCloseTextDocumentNotification](t, content, expectedResult, filepath)
		case "DidSaveTextDocumentNotification":
		validateType[DidSaveTextDocumentNotification](t, content, expectedResult, filepath)
		case "WillSaveTextDocumentNotification":
		validateType[WillSaveTextDocumentNotification](t, content, expectedResult, filepath)
		case "DidChangeWatchedFilesNotification":
		validateType[DidChangeWatchedFilesNotification](t, content, expectedResult, filepath)
		case "PublishDiagnosticsNotification":
		validateType[PublishDiagnosticsNotification](t, content, expectedResult, filepath)
		case "SetTraceNotification":
		validateType[SetTraceNotification](t, content, expectedResult, filepath)
		case "LogTraceNotification":
		validateType[LogTraceNotification](t, content, expectedResult, filepath)
		case "CancelNotification":
		validateType[CancelNotification](t, content, expectedResult, filepath)
		case "ProgressNotification":
		validateType[ProgressNotification](t, content, expectedResult, filepath)

		}
	}
}

func validateType[T any](t *testing.T, content []byte, expectedResult string, filepath string) {
	// I want to use type here
	var object T
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&object)

	if err != nil && expectedResult == "True" {
		t.Fatalf("Expected pass in file: %s, got error: %s", filepath, err)
	}

	if err == nil && expectedResult == "False" {
		t.Fatalf("Expected fail in file: %s", filepath)
	}
}
