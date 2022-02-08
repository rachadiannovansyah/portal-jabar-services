BINARY							:= worker

# ========================== build ==========================
worker:
	cd src && go build -o ${BINARY} cmd/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -f ./docker/Dockerfile -t portal-jabar_core-service --no-cache .

run:
	docker-compose up --build -d

stop:
	docker-compose down
