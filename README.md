# hello-go
Hello world in golang example with prometheus metrics, the total number of saying hello.

```
% go run hello.go
Hello, Prometheus #1 
Hello, Prometheus #2 
Hello, Prometheus #3 
Hello, Prometheus #4 
Hello, Prometheus #5 
Hello, Prometheus #6 
Hello, Prometheus #7 
...
```

```
% curl -s localhost:8080/metrics | grep hello
# HELP hello_go_num_say_hello_total The total number of saying hello
# TYPE hello_go_num_say_hello_total counter
hello_go_num_say_hello_total 6

% curl -s localhost:8080/metrics | grep hello
# HELP hello_go_num_say_hello_total The total number of saying hello
# TYPE hello_go_num_say_hello_total counter
hello_go_num_say_hello_total 7

...
```