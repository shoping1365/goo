#!/usr/bin/env python3
"""Validate that project text files are UTF-8 encoded."""
from __future__ import annotations

from pathlib import Path
import sys

TEXT_EXTENSIONS = {".ts", ".tsx", ".js", ".jsx", ".vue", ".md", ".json", ".cjs", ".mjs", ".go"}
EXCLUDE_PARTS = {"node_modules", ".nuxt", ".output", "dist", "logs", ".git"}


def should_skip(path: Path) -> bool:
    return any(part in EXCLUDE_PARTS for part in path.parts)


def is_text_candidate(path: Path) -> bool:
    return path.suffix.lower() in TEXT_EXTENSIONS


def main() -> int:
    bad: list[str] = []
    for path in Path(".").rglob("*"):
        if not path.is_file() or should_skip(path) or not is_text_candidate(path):
            continue
        try:
            path.read_text(encoding="utf-8")
        except UnicodeDecodeError:
            bad.append(path.as_posix())

    if bad:
        print("ERROR: Files with non-UTF-8 encoding found:")
        for item in bad:
            print(item)
        return 1

    print("OK: All checked files are UTF-8 encoded")
    return 0


if __name__ == "__main__":
    sys.exit(main())
