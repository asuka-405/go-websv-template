# Project Conventions

## Code Style

- Indent: tabs
- Max line length: 100
- No unused vars/imports
- Use `go fmt` always

## Naming

- Snake_case for filenames
- PascalCase for structs/types
- camelCase for variables/functions

## Project Structure

- All source code will be inside `/src`
- All binaries and tooling will be in `/bin`
- Common code resides inside `/src/lib`
- `/src/lib/internal` for private packages
- `/src/lib/pkg` for reusable code
- `/src/web` for frontend endpoints

## Commits

- feat: for new features
- fix: for bug fixes
- chore: for non-functional stuff
- test: for tests

## Dependencies

- Use stdlib unless there's a real reason
- Prefer small focused libs like `chi` over big frameworks

## Testing

- All exported funcs must have tests
- Use `t.Parallel()` where possible
