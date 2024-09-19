.PHONY: init init-db help check-job-status check-job-logs cancel-job generate-proof

init:
	cd flink && docker-compose up -d

init-db:
	docker-compose -f flink/docker-compose.yml exec -T postgres psql -U postgres -f /schema/node.sql

check-job-status:
	@echo "Checking Flink job status..."
	@docker-compose -f flink/docker-compose.yml exec jobmanager flink list -r

cancel-job:
	@echo "Canceling Flink job..."
	@read -p "Enter the Job ID to cancel: " jobid; \
	docker-compose -f flink/docker-compose.yml exec jobmanager flink cancel $$jobid

generate-proof:
	@read -p "Enter chain (default: zkevm): " chain; \
	chain=$${chain:-zkevm}; \
	read -p "Enter start block (default: 10000): " start_at; \
	start_at=$${start_at:-10000}; \
	read -p "Enter end block (default: 10010): " end_at; \
	end_at=$${end_at:-10010}; \
	read -p "Enter difficulty (default: 22): " difficulty; \
	difficulty=$${difficulty:-22}; \
	docker-compose -f flink/docker-compose.yml exec jobmanager ./bin/flink run -py /opt/proof/streaming.py --chain $$chain --start_at $$start_at --end_at $$end_at --difficulty $$difficulty

help:
	@echo "Available targets:"
	@echo "  init              - Initialize and start Flink containers"
	@echo "  init-db           - Initialize the database"
	@echo "  check-job-status  - Check the status of Flink jobs"
	@echo "  cancel-job        - Cancel a specific Flink job"
	@echo "  generate-proof    - Generate proof of streaming task"
	@echo "  help              - Display this help message"

.DEFAULT_GOAL := help
