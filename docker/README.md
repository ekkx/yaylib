# Containerized development

Build, test and run everything without installing Go, Node or Python on
the host. Source is bind-mounted; toolchains, caches and `node_modules`
live in named Docker volumes, so the host tree stays clean.

Requires Docker with Compose v2.

## Run a command

```sh
docker compose run --rm dev go test ./.                 # Go hermetic suite
docker compose run --rm dev sh -c 'cd packages/typescript && npm ci && npm test'
docker compose run --rm dev sh -c 'cd packages/python && pip install -r requirements.txt -r requirements-dev.txt && PYTHONPATH="$PWD" python -m pytest -q'
docker compose run --rm dev bash scripts/mock-parity.sh  # all three vs the mock server
docker compose run --rm dev bash                         # interactive shell
```

## Mock server as a service

```sh
docker compose up mockd          # serves on http://localhost:8080
```

Point the suites at a running instance instead of letting the parity
script boot its own:

```sh
docker compose run --rm \
  -e YAYLIB_MOCK_BASE_URL=http://mockd:8080 \
  -e YAYLIB_MOCK_WS_URL=ws://mockd:8080 \
  dev go test ./.
```

## Reset

```sh
docker compose down -v           # remove the cache/node_modules volumes
```
