build-test:
	docker stop  cintaku-test || true && docker rm  cintaku-test || true
	docker image rm faaawidia/beesocket -f
	docker build -t faaawidia/beesocket:test -f deployment/test/Dockerfile .
	docker run -p 5120:8084 --name cintaku-test -d faaawidia/beesocket:test
