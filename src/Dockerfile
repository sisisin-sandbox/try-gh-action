ARG GO_VERSION=1.22.6
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS base

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x


FROM --platform=$BUILDPLATFORM base AS build

ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server

FROM --platform=$BUILDPLATFORM base AS vet

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=0 GOARCH=$TARGETARCH go vet

FROM --platform=$BUILDPLATFORM base AS test

RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,target=. \
  CGO_ENABLED=0 GOARCH=$TARGETARCH go test

FROM debian:bookworm-slim AS final

ARG UID=10001
RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  appuser
USER appuser

COPY --from=build /bin/server /bin/

EXPOSE 8080

ENTRYPOINT [ "/bin/server" ]