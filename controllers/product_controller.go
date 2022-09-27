package controllers

import (
    "ghn-test/configs"
    "ghn-test/models"
    "net/http"
    "time"
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"fmt"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"go.mongodb.org/mongo-driver/bson"
	. "github.com/gobeam/mongo-go-pagination"
)

var productCollection *mongo.Collection = configs.GetCollection(configs.DB, "products")
var validate = validator.New()

type ApiError struct {
    Field string
    Msg   string
}

func msgForTag(tag string) string {
    switch tag {
    case "required":
        return "This field is required"
    }
    return ""
}

func CreateProduct(c echo.Context) error {
	if VerifyToken(c) == false {
		return c.JSON(http.StatusUnauthorized, &echo.Map{"error": "Unauthorized."})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var product models.Product
	defer cancel()
    
    if err := c.Bind(&product); err != nil {	
        var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]ApiError, len(ve))
            for i, fe := range ve {
                out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
            }
            return c.JSON(http.StatusBadRequest, gin.H{"errors": out})
        }
    }

    if validationErr := validate.Struct(&product); validationErr != nil {
        var ve validator.ValidationErrors
        if errors.As(validationErr, &ve) {
            out := make([]ApiError, len(ve))
            for i, fe := range ve {
                out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
            }
            return c.JSON(http.StatusBadRequest, gin.H{"errors": out})
        }
    }

    newProduct := models.Product{
        Id:       		primitive.NewObjectID(),
        Name:     		product.Name,
        Description: 	product.Description,
        Price:    		product.Price,
    }
    result, err := productCollection.InsertOne(ctx, newProduct)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
    }
	fmt.Println(result)
    return c.JSON(http.StatusOK, newProduct)
}

func EditProduct(c echo.Context) error {
	if VerifyToken(c) == false {
		return c.JSON(http.StatusUnauthorized, &echo.Map{"error": "Unauthorized."})
	}
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    productId := c.Param("productId")
    var product models.Product
    defer cancel()

    objId, _ := primitive.ObjectIDFromHex(productId)

    if err := c.Bind(&product); err != nil {	
        var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]ApiError, len(ve))
            for i, fe := range ve {
                out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
            }
            return c.JSON(http.StatusBadRequest, gin.H{"errors": out})
        }
    }

    if validationErr := validate.Struct(&product); validationErr != nil {
        var ve validator.ValidationErrors
        if errors.As(validationErr, &ve) {
            out := make([]ApiError, len(ve))
            for i, fe := range ve {
                out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
            }
            return c.JSON(http.StatusBadRequest, gin.H{"errors": out})
        }
    }

    update := bson.M{"name": product.Name, "description": product.Description, "price": product.Price}

    result, err := productCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

    if err != nil {
        return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
    }

    var updatedProduct models.Product
    if result.MatchedCount == 1 {
        err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedProduct)

        if err != nil {
            return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
        }
    } else {
		return c.JSON(http.StatusNotFound, &echo.Map{"error": "Product with specified ID='" + productId + "' not found."})
	}

    return c.JSON(http.StatusOK, updatedProduct)
}

func GetProduct(c echo.Context) error {
	if VerifyToken(c) == false {
		return c.JSON(http.StatusUnauthorized, &echo.Map{"error": "Unauthorized."})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    productId := c.Param("productId")
    var product models.Product
	defer cancel()
    objId, _ := primitive.ObjectIDFromHex(productId)
	err := productCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&product)
    if err != nil {
        return c.JSON(http.StatusNotFound, &echo.Map{"error": "Product with specified ID='" + productId + "' not found."})
    }
    return c.JSON(http.StatusOK, product)
}

func GetAllProducts(c echo.Context) error {
	if VerifyToken(c) == false {
		return c.JSON(http.StatusUnauthorized, &echo.Map{"error": "Unauthorized."})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var products []models.Product
	defer cancel()
    results, err := productCollection.Find(ctx, bson.M{})

    if err != nil {
        return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
    }

    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleProduct models.Product
        if err = results.Decode(&singleProduct); err != nil {
            return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
        }

        products = append(products, singleProduct)
    }
	
	var max int64 = 10
	if c.QueryParam("max") != "" {
		fmt.Sscan(c.QueryParam("max"), &max)
	}
	var offset int64 = 1
	if c.QueryParam("offset") != "" {
		fmt.Sscan(c.QueryParam("offset"), &offset)
	}
	projection := bson.D{}
	filter := bson.M{}
	var productList []models.Product
	paginatedData, err := New(productCollection).Context(ctx).Limit(max).Page(offset).Select(projection).Filter(filter).Decode(&productList).Find()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
	}

	payload := struct {
		Data       []models.Product	`json:"data"`
		Pagination PaginationData 	`json:"pagination"`
	}{
		Pagination: paginatedData.Pagination,
		Data:       productList,
	}

    return c.JSON(http.StatusOK, payload)
}

func DeleteProduct(c echo.Context) error {
	if VerifyToken(c) == false {
		return c.JSON(http.StatusUnauthorized, &echo.Map{"error": "Unauthorized."})
	}
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    productId := c.Param("productId")
    defer cancel()
    objId, _ := primitive.ObjectIDFromHex(productId)
    result, err := productCollection.DeleteOne(ctx, bson.M{"id": objId})
    if err != nil {
        return c.JSON(http.StatusInternalServerError, &echo.Map{"error": err.Error()})
    }
    if result.DeletedCount < 1 {
        return c.JSON(http.StatusNotFound, &echo.Map{"error": "Product with specified ID='" + productId + "' not found."})
    }
    return c.JSON(http.StatusNoContent, &echo.Map{})
}
