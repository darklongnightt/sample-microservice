module github.com/darklongnightt/microservice

go 1.15

replace github.com/darklongnightt/microservice/homepage => ./homepage

replace github.com/darklongnightt/microservice/server => ./server

require (
	github.com/darklongnightt/microservice/homepage v0.0.0-00010101000000-000000000000
	github.com/darklongnightt/microservice/server v0.0.0-00010101000000-000000000000
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v2 v2.3.0
	mellium.im/sasl v0.2.1 // indirect
)
