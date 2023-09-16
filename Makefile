run:
	docker compose up -d

add:
	curl --request POST localhost:8080/api/v1/memo/create \
		--header 'content-type: application/json' \
		--data '{"description": "test todo"}'

get:
	curl --request GET localhost:8080/api/v1/memo


restart:
	docker compose restart

build:
	docker compose build

down:
	docker compose down

logs:
	docker compose logs -f

downall:
	docker-compose down --rmi all --volumes --remove-orphans

watch:
	watch docker compose ps

uml:
	goplantuml -output doc/uml.md -ignore ent -recursive ./

