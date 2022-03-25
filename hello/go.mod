module kanoemon/go-tutorials/hello

go 1.18

require (
	golang.org/x/example v0.0.0-20220304235025-ad95e7f791d8
	kanoemon/go-tutorials/greetings v0.0.0-00010101000000-000000000000
)

replace kanoemon/go-tutorials/greetings => ../greetings
