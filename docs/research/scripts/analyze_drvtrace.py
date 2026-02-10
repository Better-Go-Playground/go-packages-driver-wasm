#!/usr/bin/env python3
"""Analyze docs/research/drvtrace.jsonl.

Outputs a JSON summary to stdout.
"""
from __future__ import annotations

import argparse
import json
from collections import Counter
from pathlib import Path
from typing import Any

LOAD_MODE_BITS = {
    "NeedName": 1 << 0,
    "NeedFiles": 1 << 1,
    "NeedCompiledGoFiles": 1 << 2,
    "NeedImports": 1 << 3,
    "NeedDeps": 1 << 4,
    "NeedExportFile": 1 << 5,
    "NeedTypes": 1 << 6,
    "NeedSyntax": 1 << 7,
    "NeedTypesInfo": 1 << 8,
    "NeedTypesSizes": 1 << 9,
    "needInternalDepsErrors": 1 << 10,
    "NeedForTest": 1 << 11,
    "typecheckCgo": 1 << 12,
    "NeedModule": 1 << 13,
    "NeedEmbedFiles": 1 << 14,
    "NeedEmbedPatterns": 1 << 15,
    "NeedTarget": 1 << 16,
}


def count_json_objects(payload: str) -> int:
    dec = json.JSONDecoder()
    i = 0
    n = 0
    ln = len(payload)
    while i < ln:
        while i < ln and payload[i].isspace():
            i += 1
        if i >= ln:
            break
        try:
            _, idx = dec.raw_decode(payload, i)
        except json.JSONDecodeError:
            break
        n += 1
        i = idx
    return n


def summarize(path: Path) -> dict[str, Any]:
    totals = Counter()
    cmd_verbs = Counter()
    cmd_args_first = Counter()
    cmd_result_kinds = Counter()
    cmd_result_counts = Counter()

    drv_modes = Counter()
    drv_mode_bits = Counter()
    drv_tests = Counter()
    drv_pattern_lists = Counter()
    drv_patterns = Counter()
    drv_cwds = Counter()
    drv_build_flags = Counter()
    drv_env_keys = Counter()
    drv_env_values = Counter()
    drv_overlay_entries = Counter()
    drv_overlay_paths = Counter()
    drv_overlay_replace_entries = Counter()
    drv_overlay_bytes = Counter()

    drv_result_kinds = Counter()
    drv_not_handled = 0
    drv_go_versions = Counter()
    drv_compiler_arch = Counter()
    drv_packages_count = Counter()
    drv_roots_count = Counter()
    drv_packages_with_module = 0
    drv_packages_total = 0

    stack_funcs = Counter()
    stack_sites = Counter()

    span_type: dict[int, str] = {}
    parent_links: list[tuple[str, int]] = []

    with path.open("r", encoding="utf-8") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            totals["lines"] += 1
            obj = json.loads(line)

            span_id = obj.get("spanId")
            parent_id = obj.get("parentSpanId", 0)

            if "cmd" in obj:
                totals["cmd"] += 1
                span_type[span_id] = "cmd"
                parent_links.append(("cmd", parent_id))

                cmd = obj["cmd"]
                verb = cmd.get("verb")
                if verb:
                    cmd_verbs[verb] += 1
                args = cmd.get("args") or []
                if args:
                    cmd_args_first[args[0]] += 1
                result = cmd.get("result") or {}
                if "ok" in result:
                    cmd_result_kinds["ok"] += 1
                    payload = result["ok"]
                    if isinstance(payload, str):
                        cmd_result_counts["json_objects"] += count_json_objects(payload)
                        cmd_result_counts["bytes"] += len(payload.encode("utf-8"))
                if "error" in result:
                    cmd_result_kinds["error"] += 1

            if "drv" in obj:
                totals["drv"] += 1
                span_type[span_id] = "drv"
                parent_links.append(("drv", parent_id))

                drv = obj["drv"]
                cwd = drv.get("cwd")
                if cwd:
                    drv_cwds[cwd] += 1

                patterns = drv.get("patterns") or []
                if patterns:
                    drv_pattern_lists["|".join(patterns)] += 1
                    for p in patterns:
                        drv_patterns[p] += 1

                req = drv.get("req") or {}
                mode = req.get("mode")
                if mode is not None:
                    drv_modes[str(mode)] += 1
                    for name, bit in LOAD_MODE_BITS.items():
                        if mode & bit:
                            drv_mode_bits[name] += 1
                tests = req.get("tests")
                if tests is not None:
                    drv_tests[str(tests)] += 1

                build_flags = req.get("build_flags") or []
                for flag in build_flags:
                    drv_build_flags[flag] += 1

                env = req.get("env") or []
                for e in env:
                    if "=" in e:
                        k, v = e.split("=", 1)
                        drv_env_keys[k] += 1
                        drv_env_values[k + "=" + v] += 1

                overlay = req.get("overlay") or {}
                if overlay:
                    drv_overlay_entries["overlay_files"] += len(overlay)
                    drv_overlay_entries["overlay_calls"] += 1
                    for _, content in overlay.items():
                        if isinstance(content, str):
                            drv_overlay_bytes["bytes"] += len(content.encode("utf-8"))
                        elif isinstance(content, list):
                            drv_overlay_bytes["bytes"] += len(bytes(content))

                ov = drv.get("overlay")
                if isinstance(ov, dict):
                    path = ov.get("path")
                    if path:
                        drv_overlay_paths[path] += 1
                    content = ov.get("content") or ov.get("Content") or {}
                    replace = content.get("replace") if isinstance(content, dict) else None
                    if isinstance(replace, dict):
                        drv_overlay_replace_entries["replace_entries"] += len(replace)
                        drv_overlay_replace_entries["replace_calls"] += 1

                result = drv.get("result") or {}
                if "ok" in result:
                    drv_result_kinds["ok"] += 1
                    rsp = result["ok"]
                    if isinstance(rsp, dict):
                        if rsp.get("NotHandled"):
                            drv_not_handled += 1
                        if "GoVersion" in rsp:
                            drv_go_versions[str(rsp.get("GoVersion"))] += 1
                        comp = rsp.get("Compiler")
                        arch = rsp.get("Arch")
                        if comp or arch:
                            drv_compiler_arch[f"{comp}|{arch}"] += 1
                        packages = rsp.get("Packages") or []
                        roots = rsp.get("Roots") or []
                        drv_packages_count[str(len(packages))] += 1
                        drv_roots_count[str(len(roots))] += 1
                        for pkg in packages:
                            if isinstance(pkg, dict):
                                drv_packages_total += 1
                                if pkg.get("Module") is not None:
                                    drv_packages_with_module += 1
                if "error" in result:
                    drv_result_kinds["error"] += 1

            stack = obj.get("stack") or []
            for frame in stack:
                if not isinstance(frame, dict):
                    continue
                func = frame.get("func")
                at = frame.get("at")
                if func:
                    stack_funcs[func] += 1
                if at:
                    stack_sites[at] += 1

    parent_counts = Counter()
    root_counts = Counter()
    missing_parent = 0
    for child_type, parent_id in parent_links:
        if parent_id == 0:
            root_counts[child_type] += 1
            continue
        parent_type = span_type.get(parent_id)
        if parent_type is None:
            missing_parent += 1
            continue
        parent_counts[f"{child_type}->{parent_type}"] += 1

    summary = {
        "totals": totals,
        "cmd": {
            "verbs": cmd_verbs,
            "args_first": cmd_args_first,
            "result_kinds": cmd_result_kinds,
            "result_counts": cmd_result_counts,
        },
        "drv": {
            "modes": drv_modes,
            "mode_bits": drv_mode_bits,
            "tests": drv_tests,
            "pattern_lists": drv_pattern_lists,
            "patterns": drv_patterns,
            "cwds": drv_cwds,
            "build_flags": drv_build_flags,
            "env_keys": drv_env_keys,
            "env_values": drv_env_values,
            "overlay_entries": drv_overlay_entries,
            "overlay_paths": drv_overlay_paths,
            "overlay_replace_entries": drv_overlay_replace_entries,
            "overlay_bytes": drv_overlay_bytes,
            "result_kinds": drv_result_kinds,
            "not_handled": drv_not_handled,
            "go_versions": drv_go_versions,
            "compiler_arch": drv_compiler_arch,
            "packages_count": drv_packages_count,
            "roots_count": drv_roots_count,
            "packages_with_module": drv_packages_with_module,
            "packages_total": drv_packages_total,
        },
        "stack": {
            "funcs": stack_funcs,
            "sites": stack_sites,
        },
        "spans": {
            "links": parent_counts,
            "roots": root_counts,
            "missing_parent": missing_parent,
        },
    }

    def to_plain(obj: Any) -> Any:
        if isinstance(obj, Counter):
            return dict(obj)
        if isinstance(obj, dict):
            return {k: to_plain(v) for k, v in obj.items()}
        return obj

    return to_plain(summary)


def main() -> None:
    ap = argparse.ArgumentParser()
    ap.add_argument("path", nargs="?", default="docs/research/drvtrace.jsonl")
    args = ap.parse_args()
    summary = summarize(Path(args.path))
    print(json.dumps(summary, indent=2, sort_keys=True))


if __name__ == "__main__":
    main()
