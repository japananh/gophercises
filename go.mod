module github.com/japananh/gophercises

go 1.19

require (
	github.com/japananh/gophercises/cyoa v0.0.0-20220921080512-d59d70964889
	github.com/japananh/gophercises/link v0.0.0-20220921070518-41e7f2dc3ffe
	github.com/japananh/gophercises/phone v0.0.0-20221005112944-3312ccf07da8
	github.com/japananh/gophercises/quiz v0.0.0-20220921080512-d59d70964889
	github.com/japananh/gophercises/sitemap v0.0.0-20220921080512-d59d70964889
	github.com/japananh/gophercises/task v0.0.0-20220922081813-dc578e53e680
	github.com/japananh/gophercises/urlshort v0.0.0-20220921080512-d59d70964889
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/exp v0.0.0-20220921023135-46d9e7742f1e // indirect
	golang.org/x/net v0.0.0-20220920203100-d0c6ba3f52d9 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/postgres v1.3.10 // indirect
	gorm.io/gorm v1.23.10 // indirect
)

replace (
	github.com/japananh/gophercises/cyoa v0.0.0-20220921080512-d59d70964889 => ./cyoa
	github.com/japananh/gophercises/link v0.0.0-20220921070518-41e7f2dc3ffe => ./link
	github.com/japananh/gophercises/phone v0.0.0-20221005112944-3312ccf07da8 => ./phone
	github.com/japananh/gophercises/quiz v0.0.0-20220921080512-d59d70964889 => ./quiz
	github.com/japananh/gophercises/sitemap v0.0.0-20220921080512-d59d70964889 => ./sitemap
	github.com/japananh/gophercises/task v0.0.0-20220922081813-dc578e53e680 => ./task
	github.com/japananh/gophercises/urlshort v0.0.0-20220921080512-d59d70964889 => ./urlshort
)
