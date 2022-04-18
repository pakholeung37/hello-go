module example/hello

go 1.18

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/tour v0.1.0
)
