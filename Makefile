run: clean
	docker compose up -d

stop:
	docker compose down

clean: stop
	docker rmi -f api


	