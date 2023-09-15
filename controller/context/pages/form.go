package pages

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func FormContext(c echo.Context) error {
	productID := c.FormValue("product_id")
	quantity := c.FormValue("quantity")
	cornColor := c.FormValue("corn_color")

	fmt.Println(productID)
	fmt.Println(quantity)
	fmt.Println(cornColor)
	

	return nil
}