# Custom Linter

## Build

`go build -buildmode=plugin ./cmd/cyclop`

## Notes

All versions of libraries that overlap `golangci-lint` (including replaced libraries) **MUST** be set to the same version as `golangci-lint`.

## Reference

- <https://developer20.com/custom-go-linter/>
- <https://golangci-lint.run/contributing/new-linters/>
- <https://disaev.me/p/writing-useful-go-analysis-linter/>
- <https://github.com/golangci/golangci-lint/issues/1276>
- <https://golangci-lint.run/usage/install/>
- <https://tech.devoted.com/custom-linter-plugins-for-golangci-lint-cf56b9091491>
- <https://play-with-go.dev/tools-as-dependencies_go119_en/>
- <https://github.com/golangci/golangci-lint/discussions/2215>
