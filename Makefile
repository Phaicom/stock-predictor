run:
	go run cmd/main.go
test:
	go test -race -coverprofile=bin/cov.out -covermode=atomic ./... 
bench:
	go test -bench=. -coverprofile=bin/cov.out ./...  
sonar:
	docker run --rm --network host --mount type=volume,src="$(shell pwd)",dst=/opt/app,type=bind -w=/opt/app red6/docker-sonar-scanner:latest sonar-scanner -Dsonar.login=def71f87f3ba2a59cf366102d7d73ac3d3f10aca