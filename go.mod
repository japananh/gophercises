module github.com/japananh/gophercises

go 1.19

require (
	github.com/japananh/gophercises/cyoa v0.0.0-20220921080512-d59d70964889
	github.com/japananh/gophercises/link v0.0.0-20220921070518-41e7f2dc3ffe
	github.com/japananh/gophercises/quiz v0.0.0-20220921080512-d59d70964889
	github.com/japananh/gophercises/sitemap v0.0.0-20220921080512-d59d70964889
	github.com/japananh/gophercises/urlshort v0.0.0-20220921080512-d59d70964889
)

require (
	golang.org/x/exp v0.0.0-20220921023135-46d9e7742f1e // indirect
	golang.org/x/net v0.0.0-20220920203100-d0c6ba3f52d9 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/japananh/gophercises/cyoa v0.0.0-20220921080512-d59d70964889 => ./cyoa
	github.com/japananh/gophercises/link v0.0.0-20220921070518-41e7f2dc3ffe => ./link
	github.com/japananh/gophercises/quiz v0.0.0-20220921080512-d59d70964889 => ./quiz
	github.com/japananh/gophercises/sitemap v0.0.0-20220921080512-d59d70964889 => ./sitemap
	github.com/japananh/gophercises/urlshort v0.0.0-20220921080512-d59d70964889 => ./urlshort
)
