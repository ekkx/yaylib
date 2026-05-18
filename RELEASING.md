# Releasing

Releases are tag-driven, and each language versions independently
(majors are aligned by convention; minor/patch drift freely). Pushing a
tag runs `.github/workflows/release.yml`:

| tag | publishes |
|---|---|
| `ts-vX.Y.Z` | **npm** (TypeScript package) + a GitHub Release |
| `py-vX.Y.Z` | **PyPI** (Python package) + a GitHub Release |
| `vX.Y.Z` | **Go** module version + a GitHub Release. Nothing extra to publish — the module *is* the repository, so the tag itself makes `go get github.com/ekkx/yaylib/v2@vX.Y.Z` resolve |

Only the language named by the tag prefix is touched, so a Python-only
fix never republishes the TypeScript package (or vice versa), and the
two can sit at different versions.

## Steps

Release the language that changed:

```sh
# TypeScript — after bumping packages/typescript/package.json
git tag ts-v2.0.1 && git push origin ts-v2.0.1

# Python — after bumping packages/python/pyproject.toml
git tag py-v2.0.1 && git push origin py-v2.0.1

# Go — no manifest; the tag is the version
git tag v2.1.0 && git push origin v2.1.0
```

For `ts-`/`py-` tags the `guard` job fails the release if the tag's
version does not match that package's manifest, so a mismatched tag
never publishes. Keep majors aligned across languages by convention.

Release notes come from the matching `## vX.Y.Z` section in
`CHANGELOG.md` (shared by the `ts-`/`py-` tags for the same version),
with auto-generated commit notes appended.

## Pre-releases

A hyphen in the version part (`ts-v2.1.0-rc.1`) is a pre-release:

- npm is published under the `next` dist-tag, so `npm install yaylib`
  still resolves the last stable; opt in with `npm install yaylib@next`
- the GitHub Release is marked pre-release
- on PyPI use the PEP 440 hyphen form (`py-v2.1.0-rc.1` ↔ pyproject
  `2.1.0-rc.1`); `pip install yaylib` skips it by default
  (use `pip install --pre yaylib`)

A version without a hyphen publishes to the default `latest`.

## One-time setup

**npm — OIDC trusted publishing (no token).** On the npm package page →
Settings → Trusted Publisher, add a GitHub Actions publisher:

- Repository: `ekkx/yaylib`
- Workflow filename: `release.yml`

A trusted publisher can only be added to an existing package, so the
very first publish must bootstrap the package: publish once with
`npm publish` locally (or a one-off automation token), then add the
trusted publisher — every release after that is tokenless.

**PyPI — API token.** Add this repository secret:

| secret | what |
|---|---|
| `PYPI_API_TOKEN` | PyPI API token scoped to the `yaylib` project |
