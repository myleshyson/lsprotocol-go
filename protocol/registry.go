package protocol
import (
	"encoding/json"
)
var MessageRegistry = map[string]func([]byte) (Message, error) {
	"textDocument/implementation": func(data []byte) (Message, error) {
		var message ImplementationRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/typeDefinition": func(data []byte) (Message, error) {
		var message TypeDefinitionRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/workspaceFolders": func(data []byte) (Message, error) {
		var message WorkspaceFoldersRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/configuration": func(data []byte) (Message, error) {
		var message ConfigurationRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/documentColor": func(data []byte) (Message, error) {
		var message DocumentColorRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/colorPresentation": func(data []byte) (Message, error) {
		var message ColorPresentationRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/foldingRange": func(data []byte) (Message, error) {
		var message FoldingRangeRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/foldingRange/refresh": func(data []byte) (Message, error) {
		var message FoldingRangeRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/declaration": func(data []byte) (Message, error) {
		var message DeclarationRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/selectionRange": func(data []byte) (Message, error) {
		var message SelectionRangeRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"window/workDoneProgress/create": func(data []byte) (Message, error) {
		var message WorkDoneProgressCreateRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/prepareCallHierarchy": func(data []byte) (Message, error) {
		var message CallHierarchyPrepareRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"callHierarchy/incomingCalls": func(data []byte) (Message, error) {
		var message CallHierarchyIncomingCallsRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"callHierarchy/outgoingCalls": func(data []byte) (Message, error) {
		var message CallHierarchyOutgoingCallsRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/semanticTokens/full": func(data []byte) (Message, error) {
		var message SemanticTokensRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/semanticTokens/full/delta": func(data []byte) (Message, error) {
		var message SemanticTokensDeltaRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/semanticTokens/range": func(data []byte) (Message, error) {
		var message SemanticTokensRangeRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/semanticTokens/refresh": func(data []byte) (Message, error) {
		var message SemanticTokensRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"window/showDocument": func(data []byte) (Message, error) {
		var message ShowDocumentRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/linkedEditingRange": func(data []byte) (Message, error) {
		var message LinkedEditingRangeRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/willCreateFiles": func(data []byte) (Message, error) {
		var message WillCreateFilesRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/willRenameFiles": func(data []byte) (Message, error) {
		var message WillRenameFilesRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/willDeleteFiles": func(data []byte) (Message, error) {
		var message WillDeleteFilesRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/moniker": func(data []byte) (Message, error) {
		var message MonikerRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/prepareTypeHierarchy": func(data []byte) (Message, error) {
		var message TypeHierarchyPrepareRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"typeHierarchy/supertypes": func(data []byte) (Message, error) {
		var message TypeHierarchySupertypesRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"typeHierarchy/subtypes": func(data []byte) (Message, error) {
		var message TypeHierarchySubtypesRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/inlineValue": func(data []byte) (Message, error) {
		var message InlineValueRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/inlineValue/refresh": func(data []byte) (Message, error) {
		var message InlineValueRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/inlayHint": func(data []byte) (Message, error) {
		var message InlayHintRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"inlayHint/resolve": func(data []byte) (Message, error) {
		var message InlayHintResolveRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/inlayHint/refresh": func(data []byte) (Message, error) {
		var message InlayHintRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/diagnostic": func(data []byte) (Message, error) {
		var message DocumentDiagnosticRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/diagnostic": func(data []byte) (Message, error) {
		var message WorkspaceDiagnosticRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/diagnostic/refresh": func(data []byte) (Message, error) {
		var message DiagnosticRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/inlineCompletion": func(data []byte) (Message, error) {
		var message InlineCompletionRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/textDocumentContent": func(data []byte) (Message, error) {
		var message TextDocumentContentRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/textDocumentContent/refresh": func(data []byte) (Message, error) {
		var message TextDocumentContentRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"client/registerCapability": func(data []byte) (Message, error) {
		var message RegistrationRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"client/unregisterCapability": func(data []byte) (Message, error) {
		var message UnregistrationRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"initialize": func(data []byte) (Message, error) {
		var message InitializeRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"shutdown": func(data []byte) (Message, error) {
		var message ShutdownRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"window/showMessageRequest": func(data []byte) (Message, error) {
		var message ShowMessageRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/willSaveWaitUntil": func(data []byte) (Message, error) {
		var message WillSaveTextDocumentWaitUntilRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/completion": func(data []byte) (Message, error) {
		var message CompletionRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"completionItem/resolve": func(data []byte) (Message, error) {
		var message CompletionResolveRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/hover": func(data []byte) (Message, error) {
		var message HoverRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/signatureHelp": func(data []byte) (Message, error) {
		var message SignatureHelpRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/definition": func(data []byte) (Message, error) {
		var message DefinitionRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/references": func(data []byte) (Message, error) {
		var message ReferencesRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/documentHighlight": func(data []byte) (Message, error) {
		var message DocumentHighlightRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/documentSymbol": func(data []byte) (Message, error) {
		var message DocumentSymbolRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/codeAction": func(data []byte) (Message, error) {
		var message CodeActionRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"codeAction/resolve": func(data []byte) (Message, error) {
		var message CodeActionResolveRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/symbol": func(data []byte) (Message, error) {
		var message WorkspaceSymbolRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspaceSymbol/resolve": func(data []byte) (Message, error) {
		var message WorkspaceSymbolResolveRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/codeLens": func(data []byte) (Message, error) {
		var message CodeLensRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"codeLens/resolve": func(data []byte) (Message, error) {
		var message CodeLensResolveRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/codeLens/refresh": func(data []byte) (Message, error) {
		var message CodeLensRefreshRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/documentLink": func(data []byte) (Message, error) {
		var message DocumentLinkRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"documentLink/resolve": func(data []byte) (Message, error) {
		var message DocumentLinkResolveRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/formatting": func(data []byte) (Message, error) {
		var message DocumentFormattingRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/rangeFormatting": func(data []byte) (Message, error) {
		var message DocumentRangeFormattingRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/rangesFormatting": func(data []byte) (Message, error) {
		var message DocumentRangesFormattingRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/onTypeFormatting": func(data []byte) (Message, error) {
		var message DocumentOnTypeFormattingRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/rename": func(data []byte) (Message, error) {
		var message RenameRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/prepareRename": func(data []byte) (Message, error) {
		var message PrepareRenameRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/executeCommand": func(data []byte) (Message, error) {
		var message ExecuteCommandRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/applyEdit": func(data []byte) (Message, error) {
		var message ApplyWorkspaceEditRequest
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/didChangeWorkspaceFolders": func(data []byte) (Message, error) {
		var message DidChangeWorkspaceFoldersNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"window/workDoneProgress/cancel": func(data []byte) (Message, error) {
		var message WorkDoneProgressCancelNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/didCreateFiles": func(data []byte) (Message, error) {
		var message DidCreateFilesNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/didRenameFiles": func(data []byte) (Message, error) {
		var message DidRenameFilesNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/didDeleteFiles": func(data []byte) (Message, error) {
		var message DidDeleteFilesNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"notebookDocument/didOpen": func(data []byte) (Message, error) {
		var message DidOpenNotebookDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"notebookDocument/didChange": func(data []byte) (Message, error) {
		var message DidChangeNotebookDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"notebookDocument/didSave": func(data []byte) (Message, error) {
		var message DidSaveNotebookDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"notebookDocument/didClose": func(data []byte) (Message, error) {
		var message DidCloseNotebookDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"initialized": func(data []byte) (Message, error) {
		var message InitializedNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"exit": func(data []byte) (Message, error) {
		var message ExitNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/didChangeConfiguration": func(data []byte) (Message, error) {
		var message DidChangeConfigurationNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"window/showMessage": func(data []byte) (Message, error) {
		var message ShowMessageNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"window/logMessage": func(data []byte) (Message, error) {
		var message LogMessageNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"telemetry/event": func(data []byte) (Message, error) {
		var message TelemetryEventNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/didOpen": func(data []byte) (Message, error) {
		var message DidOpenTextDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/didChange": func(data []byte) (Message, error) {
		var message DidChangeTextDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/didClose": func(data []byte) (Message, error) {
		var message DidCloseTextDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/didSave": func(data []byte) (Message, error) {
		var message DidSaveTextDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/willSave": func(data []byte) (Message, error) {
		var message WillSaveTextDocumentNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"workspace/didChangeWatchedFiles": func(data []byte) (Message, error) {
		var message DidChangeWatchedFilesNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"textDocument/publishDiagnostics": func(data []byte) (Message, error) {
		var message PublishDiagnosticsNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"$/setTrace": func(data []byte) (Message, error) {
		var message SetTraceNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"$/logTrace": func(data []byte) (Message, error) {
		var message LogTraceNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"$/cancelRequest": func(data []byte) (Message, error) {
		var message CancelNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
	"$/progress": func(data []byte) (Message, error) {
		var message ProgressNotification
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	},
}