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
        "/provider": {
            "get": {
                "description": "get Provider by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provider"
                ],
                "summary": "Get Provider by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Provider object",
                        "schema": {
                            "$ref": "#/definitions/dto.Provider"
                        }
                    }
                }
            },
            "post": {
                "description": "create Provider by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provider"
                ],
                "summary": "Create Provider by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Provider object",
                        "schema": {
                            "$ref": "#/definitions/dto.Provider"
                        }
                    }
                }
            },
            "patch": {
                "description": "create Provider by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provider"
                ],
                "summary": "Create Provider by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Provider object",
                        "schema": {
                            "$ref": "#/definitions/dto.Provider"
                        }
                    }
                }
            }
        },
        "/provider/${provider_id}/rates": {
            "get": {
                "description": "get Rates by Provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rates"
                ],
                "summary": "Get Rates by Provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a []dot.Rates object",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Rates"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create Provider by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rates"
                ],
                "summary": "Create Rates by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Provider object",
                        "schema": {
                            "$ref": "#/definitions/dto.Rates"
                        }
                    }
                }
            },
            "patch": {
                "description": "update rates by provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rates"
                ],
                "summary": "update rates by provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Rates object",
                        "schema": {
                            "$ref": "#/definitions/dto.Rates"
                        }
                    }
                }
            }
        },
        "/provider/{provider_id}": {
            "delete": {
                "description": "create Provider by user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Provider"
                ],
                "summary": "Create Provider by user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Provider object",
                        "schema": {
                            "$ref": "#/definitions/dto.Provider"
                        }
                    }
                }
            }
        },
        "/provider/{provider_id}/charger": {
            "get": {
                "description": "get Charger by provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charger"
                ],
                "summary": "Get Charger by provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a list of Charger object",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Charger"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create Charger by provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charger"
                ],
                "summary": "Create Charger by provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Charger object",
                        "schema": {
                            "$ref": "#/definitions/dto.Charger"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Charger by provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charger"
                ],
                "summary": "Update Charger by provider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Charger object",
                        "schema": {
                            "$ref": "#/definitions/dto.Charger"
                        }
                    }
                }
            }
        },
        "/provider/{provider_id}/charger/{charger_id}": {
            "delete": {
                "description": "Delete Charger by charger id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Charger"
                ],
                "summary": "Delete Charger by charger id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Charger id",
                        "name": "charger_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns a Charger object",
                        "schema": {
                            "$ref": "#/definitions/dto.Charger"
                        }
                    }
                }
            }
        },
        "/provider/{provider_id}/rates/{rates_id}": {
            "delete": {
                "description": "delete rates by rates id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rates"
                ],
                "summary": "delete rates by rates id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "jwtToken of the user",
                        "name": "authentication",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "Provider id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "provider id",
                        "name": "provider_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "rates id",
                        "name": "rates_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns true/false",
                        "schema": {
                            "$ref": "#/definitions/dto.Rates"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Charger": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                },
                "provider_id": {
                    "type": "integer"
                },
                "rates_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.Provider": {
            "type": "object",
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "user_email": {
                    "type": "string"
                }
            }
        },
        "dto.Rates": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "no_show_penalty_rate": {
                    "type": "number"
                },
                "normal_rate": {
                    "type": "number"
                },
                "penalty_rate": {
                    "type": "number"
                },
                "provider_id": {
                    "type": "integer"
                },
                "status": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Provider Service API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
