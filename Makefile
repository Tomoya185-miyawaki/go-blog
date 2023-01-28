up:
	docker-compose up
upd:
	docker-compose up -d
build:
	docker-compose build
down:
	docker-compose down
sh:
	docker-compose exec app sh
db:
	mysql -ugo_blog -h 127.0.0.1 --port 13306 -p