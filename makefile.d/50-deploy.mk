##@ Development environment

PROJECT_NAME = go-cruise-control
DEPLOY_DIR ?= $(PROJECT_DIR)/deploy

.PHONY: env-up
env-up: ## Spin up local development environment
	$(info *** Spinning up local development environment...)
	@docker-compose \
		--project-name $(PROJECT_NAME) \
		--project-directory $(DEPLOY_DIR) \
		--file $(DEPLOY_DIR)/docker-compose.yml \
		up \
		--detach \
		--quiet-pull \
		--no-recreate \
		--remove-orphans

.PHONY: env-down
env-down: ## Stop local development environment
	$(info *** Stopping local development environment...)
	@docker-compose \
		--project-name $(PROJECT_NAME) \
		--project-directory $(DEPLOY_DIR) \
		--file $(DEPLOY_DIR)/docker-compose.yml \
		down \
		--remove-orphans

.PHONY: env-clean
env-clean: ## Cleanup local development environment
	$(info *** Tearing down local development environment...)
	@docker-compose \
		--project-name $(PROJECT_NAME) \
		--project-directory $(DEPLOY_DIR) \
		--file $(DEPLOY_DIR)/docker-compose.yml \
		down \
		--volumes \
		--remove-orphans

.PHONY: env-with-ui-up
env-with-ui-up: ## Spin up local development environment
	$(info *** Spinning up local development environment...)
	@docker-compose \
		--project-name $(PROJECT_NAME) \
		--project-directory $(DEPLOY_DIR) \
		--file $(DEPLOY_DIR)/docker-compose.yml \
		--file $(DEPLOY_DIR)/docker-compose.ui.yml \
		up \
		--detach \
		--quiet-pull \
		--no-recreate \
		--remove-orphans