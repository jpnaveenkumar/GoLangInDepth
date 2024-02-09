Steps to run the code:

go run main.go


run go program with race check:

go run -race main.go


Visualize trace:

go tool trace concurrentTrace.out

go tool trace parallelTrace.out


Visualize PPROF result:

go tool pprof -http=:8080 <PROFILE_NAME>.pb.gz
