module crane_game

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/protobuf v1.26.0

	crane_lib v0.0.1
)

replace (
	crane_lib v0.0.1 => ./libs/crane_lib
)