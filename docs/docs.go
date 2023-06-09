// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/orders": {
            "get": {
                "description": "Get details of all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get All Orders",
                "operationId": "get-all-orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AllOrderResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Create Order",
                "operationId": "create-new-order",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponse"
                        }
                    }
                }
            }
        },
        "/orders/{OrderId}": {
            "get": {
                "description": "Get details an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get Details Order",
                "operationId": "get-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the order",
                        "name": "OrderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Update Order",
                "operationId": "update-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the order",
                        "name": "OrderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Delete Order",
                "operationId": "delete-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of the order",
                        "name": "OrderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.DeleteOrderResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AllOrderResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.OrderResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.DeleteOrderResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.ItemResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "itemcode": {
                    "type": "string"
                },
                "orderId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.NewItemRequest": {
            "type": "object",
            "required": [
                "description",
                "itemCode",
                "quantity"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "itemCode": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "dto.NewOrderRequest": {
            "type": "object",
            "required": [
                "customerName",
                "items",
                "orderedAt"
            ],
            "properties": {
                "customerName": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.NewItemRequest"
                    }
                },
                "orderedAt": {
                    "type": "string"
                }
            }
        },
        "dto.NewOrderResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.OrderResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.OrderResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "customerName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemResponse"
                    }
                },
                "orderedAt": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
