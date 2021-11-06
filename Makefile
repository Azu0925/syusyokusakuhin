up:
	docker-compose up -d

stop:
	docker-compose stop

restart:
	docker-compose restart

down:
	docker-compose down

down_image:
	docker-compose down --rmi all

down_vol:
	docker-compose down -v