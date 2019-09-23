all: run
run:
	go run server.go auth.go controllers.go router.go entities.go --env=local