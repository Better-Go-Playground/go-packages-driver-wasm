// NOTE: DriverRequest and DriverResponse types are declared in ./drivertypes.ts
import type { DriverRequest, DriverResponse } from "./drivertypes"

// CallEnvelope is a type that represents single JSON line.
// Contains either package driver or Go command call.
//
// Each calls can be parent-child referenced using SpanInfo in the header.
// Usually any "drv" call is followed by "cmd" call associated by traceId relation.
type CallEnvelope = CallHeader & CallBody

type CallBody = {
  drv: PackageDriverCall
} | {
  cmd: GoCommandCall
}

interface CallHeader extends SpanInfo {
  // ts is call timestamp in milliseconds
  ts: number

  // stack contains stack trace of where call was done in gopls source code.
  // used to reference gopls source code.
  stack: StackFrame[]
}

// GoCommandCall contains "go" command call metadata done by gopls.
interface GoCommandCall {
  // verb is go sub-command. E.g: list, version.
  verb: string

  // args is command-line arguments after the verb.
  args: string[]

  // go command output. May be post-processed or decorated by gopls (see `goListState.createDriverResponse` in `go/packages` package in golang.org/x/tools).
  result: Result<Record<string, any> | string>
}

// PackageDriverCall contains arguments & context for Go package driver call.
interface PackageDriverCall {
  // working directory
  cwd: string

  // Go package driver query. See Go package driver protocol.
  patterns: string[]

  // overlay is contents of overlay file that is passed to downstream "go list" command.
  // Represents the same as overlay in [DriverRequest], but contains file paths mapping only.
  overlay?: OverlayFile

  // Go package driver request.
  req: DriverRequest

  // Go package driver response.
  result: Result<DriverResponse>
}

// Result monad represents operation result or error.
// Either ok or error is present.
interface Result<T> {
  ok?: T
  error?: string
}

interface SpanInfo {
  // Call span ID, can be used to reference children calls.
  spanId?: number
  // parentSpanId is "spanID" of a parent call that caused this call.
  parentSpanId?: number
}

interface StackFrame {
  // func is name of the function
  func: string
  // at is call location in format "filename:line"
  at: string
}

interface OverlayFile {
  // Path to Go overlay file.
  path: string

  // Contents of Go overlay file.
  Content: {
    // key-value pair of source file and actual file that needs to be read.
    replace: Record<string, string>
  }
}
