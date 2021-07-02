package controller

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []model.Product

	result := repo.Db.Preload("Category").
		Find(&products)
	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product := repo.FindProductById(id)
	fmt.Println(product)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(product)
}

func GetProductsByCategory(c *fiber.Ctx) error {
	category := c.Params("name_category")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	results := repo.FindProductByCategory(category)
	var answer []*model.Product
	if offset+limit > len(results) {
		answer = results[offset:]
	} else {
		answer = results[offset : offset+limit]
	}
	response := make(map[string]interface{})
	response["listProduct"] = answer
	response["totalProducts"] = len(results)
	if len(results)%limit == 0 {
		response["totalPages"] = len(results) / limit
	} else {
		response["totalPages"] = len(results)/limit + 1
	}
	return c.JSON(response)
}

func GetProductsSale(c *fiber.Ctx) error {
	issale := c.Params("is_sale")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	results := repo.FindProductSale(issale)
	var answer []*model.Product
	if offset+limit > len(results) {
		answer = results[offset:]
	} else {
		answer = results[offset : offset+limit]
	}
	response := make(map[string]interface{})
	response["listProduct"] = answer
	response["totalProducts"] = len(results)
	if len(results)%limit == 0 {
		response["totalPages"] = len(results) / limit
	} else {
		response["totalPages"] = len(results)/limit + 1
	}
	return c.JSON(response)
}

func GetProductByName(c *fiber.Ctx) error {
	name, _ := url.PathUnescape(c.Params("input_name"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	results := repo.FindProductByName(name)
	var answer []*model.Product
	if offset+limit > len(results) {
		answer = results[offset:]
	} else {
		answer = results[offset : offset+limit]
	}
	response := make(map[string]interface{})
	response["listProduct"] = answer
	response["totalProducts"] = len(results)
	if len(results)%limit == 0 {
		response["totalPages"] = len(results) / limit
	} else {
		response["totalPages"] = len(results)/limit + 1
	}
	return c.JSON(response)
}
