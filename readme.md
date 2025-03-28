# REST api example

1. Create project directory (student-api)
2. Check GO install or not

```bash
> go version
```

3. Initialize go module

```bash
> go mod init github.com/<username>/<repo-name>
> go mod init github.com/bjayanta/student-api
```

4. Create `cmd` directory
5. Create `<project-name>` directory in `cmd`
6. Create `main.go` file in `<project-name>` directory
7. Create `.gitignore` file

- Install gitignore extension in vscode
- Ctrl + p to select & create gitignore file

8. Initialize git + commit

9. Create `config` directory in root

- Create `local.yaml` file
- Add this directory into `.gitignore` file

10. Create `storage` directory in root

- Create `storage.db` sqlite file
- Add this directory into `.gitignore` file

11. To serialize config directory data we need to make `internal` directory

- Create config/config.go file
- To serialize install `cleanenv` package

Example:

`go get -u <package-name>`

```bash
> go get -u github.com/ilyakaznacheev/cleanenv
```

- Set `struct-tags` to define serialization
- Create `MustLoad` function to load environment

12. Load config
13. Setup router
14. Create server + Graceful Shutdown
15. Create student handlers

- Create student.go file in `/internal/handlers/student`
- Create types.go file in `/internal/types`
- Create response.go file in `/internal/utils/response`

16. Request validation

- Add new validation package `go-playground/validator`

```bash
> go get -u github.com/go-playground/validator/v10
```

- Create validation method in `response.go` named `ValidationError`

17. Database connection

- Create `storage.go` file into `/internal/storage`
- Create `sqlite.go` file into `/internal/storage/sqlite`
- Install `sqlite` driver

```bash
> go get modernc.org/sqlite
```

- Create `students` table
- Add DB connection into `main.go` file

18.

### Tutorial

https://www.youtube.com/watch?v=OGhQhFKvMiM&t=747s
