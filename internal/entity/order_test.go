package entity

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T){
	order := Order{}
	assert.Error(t, order.Validate(), "invalid ID")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Zero_Or_Less_Than_Zero(t *testing.T){
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Zero_Or_Less_Than_Zero(t *testing.T){
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_With_All_Valid_Params(t *testing.T){
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}

	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}