### Info

1. Repository structure based on [this!](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
2. Program is located in `/cmd/fareestimate/main.go`
3. program starts by initiating `/pkg/config/config.go` package that sets program wide default vars as well as accepts command line argumens to override defaults. Run `go run cmd/fareestimate/main.go -help` to see all available options.
4. program gets pointer to fare calculator. It's used for calculating fares based on time of the day and distance gained between 2 points. It also keeps track of total ride fare and able to produce csv friendly representation of this data.
5. it then gets data stream window `streamwindow.New()`. This is a 2 slot moving window that will be accepting data from input file stream. It will act as filter and decorator, ensuring invalid data does not pass through it and 2 different rides don't get mixed up. It returns valid pair of 2 enriched points.
6. programs starts reading file line by line using bufio scanner so that host machine does not run out of memory when processing large files. Process starts in separate routine and accepts channel over which it's going to send file data. Process `csv row -> []string{} -> PointInTime{} -> filter -> fare calculator -> report`. One ride can have millions of points in csv file but in fare calculator it's going to be represented by single key/value pair. So it's memory efficient and safe.
8. reading is done, program get's valid fares, transforms them into csv friendly input and generates report file.
9. run tests from root folder `go test ./...`
