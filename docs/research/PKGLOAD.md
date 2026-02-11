# When gopls Calls packages.Load

This note summarizes when gopls triggers `go/packages.Load` at runtime, and maps those triggers to common LSP requests and user actions.

## The Single Entry Point

`go/packages.Load` is only called from `Snapshot.load`.

Call sites for `Snapshot.load`:
- Initial workspace load in `Snapshot.initialize`.
- Reloads after invalidation in `Snapshot.reloadWorkspace`.
- Per-file loads in `Snapshot.MetadataForFile` when metadata is missing or marked stale (`shouldLoad`).

## High-Level Triggers

`packages.Load` can run multiple times per session, not just once at startup.

- Initial workspace load (IWL) when a view is initialized.
- Reloads after invalidation caused by file edits, saves, closes, or on-disk changes.
- File-scoped loads when a request needs metadata for a file that is missing or marked for reload.

## LSP Requests and User Actions

The items below identify common user actions, the LSP entry point, and the path that can reach `packages.Load`.

1. Open a file in the editor
- LSP: `textDocument/didOpen`
- Path: `didModifyFiles` invalidates state and schedules diagnostics. Initial diagnostics or any request that needs metadata will call `Snapshot.awaitLoaded`, which triggers initial workspace load (`Snapshot.initialize` -> `Snapshot.load` -> `packages.Load`).

2. Edit a file (typing)
- LSP: `textDocument/didChange`
- Path: invalidation marks packages `shouldLoad`. Subsequent diagnostics or any request needing metadata calls `Snapshot.reloadWorkspace` -> `Snapshot.load`.

3. Save a file
- LSP: `textDocument/didSave`
- Path: same as edit, via invalidation and a later reload when metadata is required.

4. Close a file
- LSP: `textDocument/didClose`
- Path: may cause view recomputation or a later reload on demand, which triggers `Snapshot.load`.

5. On-disk file changes (git checkout, file add/delete)
- LSP: `workspace/didChangeWatchedFiles`
- Path: invalidation and later `Snapshot.reloadWorkspace` when diagnostics or a metadata-based feature is requested.

6. Workspace changes (go.mod or go.work changes, workspace folders)
- LSP: `workspace/didChangeWatchedFiles`, `workspace/didChangeWorkspaceFolders`, `workspace/didChangeConfiguration`
- Path: view recomputation or new view creation, which performs the initial workspace load (`Snapshot.initialize` -> `Snapshot.load`).

7. Pull diagnostics
- LSP: `textDocument/diagnostic`
- Path: `golang.DiagnoseFile` -> `NarrowestMetadataForFile` -> `MetadataForFile` -> `Snapshot.load` if metadata is missing or stale.

8. Push diagnostics (background after edits)
- LSP: internal diagnostic pipeline after file changes
- Path: `diagnoseSnapshot` -> `WorkspaceMetadata` -> `awaitLoaded` -> `reloadWorkspace` -> `Snapshot.load`.

9. File-scoped code intelligence (Go files)
- LSP: `textDocument/hover`, `definition`, `typeDefinition`, `references`, `implementation`, `rename`, `codeAction`, `formatting`, `completion`, `signatureHelp`, `documentSymbol`, `inlayHint`, `semanticTokens*`
- Path: these call `golang.*` helpers that use `NarrowestPackageForFile` or `MetadataForFile`; if metadata is missing or stale, `Snapshot.load` runs.

10. Workspace-wide features
- LSP: `workspace/symbol`
- Path: uses `WorkspaceMetadata` or `AllMetadata`, which call `awaitLoaded` and may trigger `reloadWorkspace` -> `Snapshot.load`.

11. gopls commands
- LSP: `workspace/executeCommand` with commands like `gopls.doc`, `gopls.run_tests`, `gopls.gc_details`, `gopls.list_imports`, `gopls.list_known_packages`, `gopls.regenerate_cgo`
- Path: these use `NarrowestPackageForFile` / `NarrowestMetadataForFile` or reset views. This can trigger file-scoped loads or full workspace reloads.

## Why It May Not Run

If metadata is already up-to-date, requests can be satisfied from cache without calling `packages.Load`. The invalidation logic uses `shouldLoad` to mark packages that require reloading, and only then does `Snapshot.load` run.

## Relevant Files

- `gopls/internal/cache/load.go` (calls `packages.Load`)
- `gopls/internal/cache/view.go` (initial workspace load in `Snapshot.initialize`)
- `gopls/internal/cache/snapshot.go` (`MetadataForFile`, `reloadWorkspace`, `awaitLoaded`)
- `gopls/internal/server/text_synchronization.go` (LSP file-change notifications)
- `gopls/internal/server/diagnostics.go` (diagnostic pipeline)
- `gopls/internal/server/*.go` and `gopls/internal/golang/*.go` (code intelligence features)

