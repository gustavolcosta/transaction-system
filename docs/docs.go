// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/accounts": {
            "post": {
                "description": "create an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Create an Account",
                "parameters": [
                    {
                        "description": "Account info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateAccountInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateAccountOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "description": "get account by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get Account by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.GetAccountByIdOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "create an Transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Create an Transaction",
                "parameters": [
                    {
                        "description": "Transactions info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateTransactionInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateTransactionOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.ExceptionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateAccountInputDTO": {
            "type": "object",
            "required": [
                "document_number"
            ],
            "properties": {
                "document_number": {
                    "type": "string",
                    "example": "1234567890"
                }
            }
        },
        "dtos.CreateAccountOutputDTO": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer",
                    "example": 1
                },
                "document_number": {
                    "type": "string",
                    "example": "1234567890"
                }
            }
        },
        "dtos.CreateTransactionInputDTO": {
            "type": "object",
            "required": [
                "account_id",
                "amount",
                "operation_type_id"
            ],
            "properties": {
                "account_id": {
                    "type": "integer",
                    "example": 1
                },
                "amount": {
                    "type": "number",
                    "example": 123.4
                },
                "operation_type_id": {
                    "type": "integer",
                    "example": 4
                }
            }
        },
        "dtos.CreateTransactionOutputDTO": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer",
                    "example": 1
                },
                "amount": {
                    "type": "number",
                    "example": 123.4
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "operation_type_id": {
                    "type": "integer",
                    "example": 4
                }
            }
        },
        "dtos.GetAccountByIdOutputDTO": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer",
                    "example": 1
                },
                "document_number": {
                    "type": "string",
                    "example": "1234567890"
                }
            }
        },
        "response.ExceptionResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Transaction System API",
	Description:      "App to handle with transactions routine",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
