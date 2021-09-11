
get-metrics:
	curl -s localhost:8080/metrics | grep hello
