#!/bin/bash

CONTAINER_NAME = mongon
NETWORK_NAME = mongonn
HELP_FUN = \
	%help; \
    while(<>) { \
        if(/^([a-z0-9_-]+):.*\#\#(?:@(\w+))?\s(.*)$$/) { \
            push(@{$$help{$$2}}, [$$1, $$3]); \
        } \
    }; \
    print "\nusage: make [target]\n\n"; \
    for ( sort keys %help ) { \
        printf("  %-20s %s\n", $$_->[0], $$_->[1]) for @{$$help{$$_}}; \
        print "\n"; \
    }

setup:		## Prepare and run docker environment
	docker network create ${NETWORK_NAME} || true
	cp -n .env.docker .env || true
	docker-compose up -d

restart:		## Restart all containers
	docker-compose restart

bash:  		## Bash in the container
	docker exec -it ${CONTAINER_NAME} sh

sh:  		## Bash in the container
	docker exec -it ${CONTAINER_NAME} sh

help: 		## List commands with help
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)