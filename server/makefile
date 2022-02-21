app_target = cli/cli.go

fresh-db-re:
	go build -o dist/cli.exe $(app_target)
	dist/cli.exe fresh:db

fresh-db:
	dist/cli.exe fresh:db

run:
	go run ./main.go

.PHONY:fresh-db