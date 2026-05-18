# Releasing

Releases are tag-driven. Pushing a `vX.Y.Z` tag runs
`.github/workflows/release.yml`, which publishes:

- **npm** — the TypeScript package
- **PyPI** — the Python package
- **GitHub Release** — generated notes
- **Go** — nothing extra; the module *is* the repository, so the tag
  itself makes `go get github.com/ekkx/yaylib/v2@vX.Y.Z` resolve

## Steps

1. Bump the version on `master` so both packages agree:
   - `packages/typescript/package.json` → `"version"`
   - `packages/python/pyproject.toml` → `version`
2. Push a matching tag:

   ```sh
   git tag v2.0.0-rc.1
   git push origin v2.0.0-rc.1
   ```

The workflow's `guard` job fails the release if the tag does not match
both package versions, so a mismatched tag never publishes.

## Pre-releases

A tag containing a hyphen (`v2.0.0-rc.1`) is a pre-release:

- npm is published under the `next` dist-tag, so `npm install yaylib`
  still resolves the last stable; opt in with `npm install yaylib@next`
- the GitHub Release is marked pre-release
- on PyPI it is a PEP 440 pre-release, which `pip install yaylib` skips
  by default (use `pip install --pre yaylib`)

A tag without a hyphen (`v2.0.0`) publishes to the default `latest`.

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
