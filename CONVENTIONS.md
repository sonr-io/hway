# CLAUDE.md - Agent Guidelines for Hway Codebase

## Build Commands
- `make deps` - Install dependencies
- `make build` - Build all binaries
- `make test` - Run all tests
- `make image` - Build Docker image
- `go test -v ./path/to/package -run TestName` - Run single test

## Linting & Formatting
- Use standard Go formatting (`go fmt`)
- Follow standard import ordering: stdlib first, then external packages
- Run `go vet ./...` before committing

## Coding Conventions
- Use CamelCase for exported functions/variables, camelCase for private
- Create interfaces before implementations
- Type aliases should have clear documentation
- Use descriptive variable names that indicate purpose
- Constants should use PascalCase

## Error Handling
- Always check errors and return them up the call stack
- Use `return nil, err` pattern consistently
- Only use panic for unrecoverable startup errors
- Include error descriptions in function documentation

## Testing
- Write tests for new functionality
- Use table-driven tests where appropriate
- Mock external dependencies