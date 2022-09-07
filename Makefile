.PHONY: restart
restart:
	docker-compose down
	rm -rf mysql
	rm -rf phpmyadmin
	docker compose up -d --build