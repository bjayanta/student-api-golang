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
15.
