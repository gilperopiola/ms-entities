all: run
run:
	go run server.go controllers.go router.go entities.go --env=$(env)