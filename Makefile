.PHONY: build

mig:
	migrate -path ./shema -database postgres://kofjeegdcybwla:b39566238fc7bf3518d3a3474cc4fa1e8b6a7b91183d4e6163cbe432fab95d76@ec2-54-75-184-144.eu-west-1.compute.amazonaws.com:5432/d9lvk2najfs6jb up
	
migd:
	migrate -path ./shema -database postgres://postgres:postgrespw@localhost:49153 up

down:
	migrate -path ./shema -database postgres://kofjeegdcybwla:b39566238fc7bf3518d3a3474cc4fa1e8b6a7b91183d4e6163cbe432fab95d76@ec2-54-75-184-144.eu-west-1.compute.amazonaws.com:5432/d9lvk2najfs6jb down

create:
	docker run --name=todo-db -e POSTGRES_PASSWORD='postal' -p 7777:5432 -d --rm postgres    