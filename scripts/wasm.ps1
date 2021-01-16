$env:GOOS = "js"
$env:GOARCH = "wasm"

go build -o ./web/public/wasm ./cmd/wasm
