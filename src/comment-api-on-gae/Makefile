init:
	@docker-compose run --rm dep init
install:
	@docker-compose run --rm dep ensure
serve-dev:
	@dev_appserver.py ./app --host 0.0.0.0 --enable_host_checking=false --admin_host 0.0.0.0 --skip_sdk_update_check --enable_watching_go_path --use_mtime_file_watcher
run:
	@docker-compose up app