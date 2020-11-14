package handlers

import (
	"greedy-game-test/node"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Initialising root node of the tree
var (
	tree = node.New(nil)
)

//Insert handler
func Insert(c echo.Context) error {
	var (
		request  node.WebTraffic
		response node.Response
		country  string
		device   string
	)
	if err := c.Bind(&request); err != nil {
		return err
	}
	//Reading dimensions such as country and device values from request body
	for _, dim := range request.Dimensions {
		switch dim.Key {
		case "country":
			country = dim.Value
		case "device":
			device = dim.Value
		}
	}
	//Validation
	if country == "" {
		response.Errors = append(response.Errors, "Country is required")
	}
	if device == "" {
		response.Errors = append(response.Errors, "Device is required")
	}
	if len(response.Errors) > 0 {
		return c.JSON(http.StatusBadRequest, response)
	}
	// Valid data, reading through the metrics values
	for _, metric := range request.Metrics {
		tree.UpdateMetric(country, device, metric.Key, metric.Value)
	}
	response.Message = "success"
	return c.JSON(http.StatusOK, response)
}

//Query handler
func Query(c echo.Context) error {
	var (
		request  node.WebTraffic
		response node.WebTraffic
		country  string
	)
	if err := c.Bind(&request); err != nil {
		return err
	}
	for _, dim := range request.Dimensions {
		if dim.Key == "country" {
			country = dim.Value
			break
		}
	}
	if country == "" {
		return c.JSON(http.StatusBadRequest, "Country is required")
	}
	//Vaild data
	response, err := tree.GetMetricByCountry(request, country)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, response)
}
