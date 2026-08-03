#!/usr/bin/env python3
# SPDX-License-Identifier: Apache-2.0
# SPDX-FileCopyrightText: 2026 The plugin-template Authors

from pathlib import Path


workflow = Path(".github/workflows/release.yml").read_text(encoding="utf-8")


def require(fragment: str) -> None:
    if fragment not in workflow:
        raise SystemExit(f"release workflow is missing required invariant: {fragment}")


require("verify_exact_image()")
require('digest="${existing_digest}"')
require('verify_exact_image "${exact_tag}"')
require('org.opencontainers.image.revision')
require('org.opencontainers.image.version')
require('--arg amd64 "${AMD64_IMAGE_DIGEST}"')
require('--arg arm64 "${ARM64_IMAGE_DIGEST}"')

existing = workflow.index('if existing_digest=$(resolve_registry_ref "${exact_tag}"); then')
candidate = workflow.index('candidate_tag="${IMAGE}:manifest-', existing)
missing_tag_branch = workflow.index("else", existing, candidate)
if not existing < missing_tag_branch < candidate:
    raise SystemExit("candidate manifest creation must occur only when the exact tag is absent")

release_step = workflow.index("uses: softprops/action-gh-release@")
next_step = workflow.find("\n      - name:", release_step)
release_config = workflow[release_step : next_step if next_step >= 0 else None]
if "make_latest:" in release_config:
    raise SystemExit("GitHub must select Latest; manual reruns must not force make_latest")

print("release workflow retry and Latest-selection invariants verified")
