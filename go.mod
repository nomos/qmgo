module github.com/nomos/qmgo

go 1.16

require (
	github.com/go-playground/validator/v10 v10.4.1
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.5.1
)

replace go.mongodb.org/mongo-driver v1.5.1 => github.com/nomos/mongo-go-driver v1.1.5
