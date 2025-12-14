module github.com/rmocchy/convinient_wire

go 1.25.1

require (
	github.com/rmocchy/convinient_wire/sample/basic v0.0.0-00010101000000-000000000000
	golang.org/x/tools v0.40.0
)

require (
	github.com/google/wire v0.6.0 // indirect
	golang.org/x/mod v0.31.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
)

replace github.com/rmocchy/convinient_wire/sample/basic => ./sample/basic
