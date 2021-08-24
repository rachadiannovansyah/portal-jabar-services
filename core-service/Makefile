BINARY									:= engine
ROOT                    := $(PWD)
GO_HTML_COV             := ./coverage.html
GO_TEST_OUTFILE         := ./c.out
GOLANG_DOCKER_IMAGE     := golang:1.15
CC_TEST_REPORTER_ID			:= ${CC_TEST_REPORTER_ID}
CC_PREFIX								:= github.com/jabardigitalservice/portal-jabar-services/core-service

# ========================== database ==========================	
# export env first: export $(xargs <.env)
migrate:
	docker run --rm -v "${ROOT}/src/database/migrations:/migrations" \
		--network host migrate/migrate -path ./migrations/ \
		-database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:3306)/${DB_NAME}" up

# ========================== build ==========================
engine:
	cd src && go build -o ${BINARY} cmd/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -f ./docker/Dockerfile -t portal-jabar_core-service --no-cache .

run:
	docker-compose up --build -d

stop:
	docker-compose down

# ========================== test ==========================
# custom logic for code climate, gross but necessary
coverage:
	# download CC test reported
	docker run  --rm -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} \
		/bin/bash -c \
		"curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter"
	
	# update perms
	docker run --rm -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} chmod +x ./cc-test-reporter

	# run before build
	docker run --rm -w /app -v ${ROOT}:/app \
		 -e CC_TEST_REPORTER_ID=${CC_TEST_REPORTER_ID} \
		${GOLANG_DOCKER_IMAGE} ./cc-test-reporter before-build

	# run testing
	docker run --rm -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go test ./... -coverprofile=${GO_TEST_OUTFILE}
	docker run --rm -w /app -v ${ROOT}:/app ${GOLANG_DOCKER_IMAGE} go tool cover -html=${GO_TEST_OUTFILE} -o ${GO_HTML_COV}

	#upload coverage result
	$(eval PREFIX=${CC_PREFIX})
ifdef prefix
	$(eval PREFIX=${prefix})
endif
	# upload data to CC
	docker run -w /app -v ${ROOT}:/app \
		-e CC_TEST_REPORTER_ID=${CC_TEST_REPORTER_ID} \
		${GOLANG_DOCKER_IMAGE} ./cc-test-reporter after-build --prefix ${PREFIX}

test:
	@go test ./... -coverprofile=./coverage.out & go tool cover -html=./coverage.out

