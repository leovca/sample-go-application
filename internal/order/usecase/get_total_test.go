package usecase

import (
	"database/sql"
	"intensive/internal/order/entity"
	"intensive/internal/order/infraestructure/database"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type GetTotalUseCaseTestSuite struct {
	suite.Suite
	OrderRepository entity.OrderRepositoryInterface
	Db              *sql.DB
}

func (suite *GetTotalUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = database.NewOrderRepository(db)
}

func (suite *GetTotalUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuiteGetTotal(t *testing.T) {
	suite.Run(t, new(GetTotalUseCaseTestSuite))
}

func (suite *GetTotalUseCaseTestSuite) TesteGetTotalUseCase() {
	calculatePriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	calculatePriceUseCase.Execute(OrderInputDTO{
		ID:    "1",
		Price: 10,
		Tax:   2,
	})
	calculatePriceUseCase.Execute(OrderInputDTO{
		ID:    "2",
		Price: 10,
		Tax:   2,
	})
	calculatePriceUseCase.Execute(OrderInputDTO{
		ID:    "3",
		Price: 10,
		Tax:   2,
	})

	getTotalUseCase := NewGetTotalUseCase(suite.OrderRepository)

	output, err := getTotalUseCase.Execute()
	suite.NoError(err)

	suite.Equal(output.Total, 3)

}
