LOCALBIN := $(PWD)/.local/bin
# tools
GORELEASER := $(LOCALBIN)/goreleaser
OGEN := $(LOCALBIN)/ogen

$(LOCALBIN):
	mkdir -p "$(LOCALBIN)"

$(GORELEASER): $(LOCALBIN)
	GOBIN="$(LOCALBIN)" go install -v github.com/goreleaser/goreleaser@latest

$(OGEN): $(LOCALBIN)
	GOBIN="$(LOCALBIN)" go install -v github.com/ogen-go/ogen/cmd/...@v0.68.4

gen: $(OGEN)
	go generate ./...
	# generate public client
	$(OGEN) --target api/client --package client --no-server --clean  --no-webhook-client --no-webhook-server --convenient-errors=off openapi.yaml
	# generate internal server
	$(OGEN) --target internal/server/api --no-client --clean  --no-webhook-client --no-webhook-server --convenient-errors=off openapi.yaml