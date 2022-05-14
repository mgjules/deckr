// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Michaël Giovanni Jules",
            "url": "https://mgjules.dev",
            "email": "julesmichaelgiovanni@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "checks if server is running",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/decks": {
            "post": {
                "description": "creates a new full or partial deck of cards given an optional list of codes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "creates a new deck of cards",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "example": "AS, 2C, 3D, 4H, 5S",
                        "description": "list of codes",
                        "name": "codes",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/http.DeckClosed"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/decks/{id}": {
            "get": {
                "description": "opens a deck of cards given an id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "opens a deck of cards",
                "parameters": [
                    {
                        "type": "string",
                        "example": "9302b603-13bb-5275-a3b9-5fcefafa34e0",
                        "description": "id of deck",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.DeckOpened"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/decks/{id}/draw": {
            "get": {
                "description": "draws cards from a deck of cards given an id and the number of cards",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "draws cards from a deck of cards",
                "parameters": [
                    {
                        "type": "string",
                        "example": "9302b603-13bb-5275-a3b9-5fcefafa34e0",
                        "description": "id of deck",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "example": 5,
                        "description": "number of cards",
                        "name": "num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Cards"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/decks/{id}/shuffle": {
            "post": {
                "description": "shuffle a deck of cards given an id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deck"
                ],
                "summary": "shuffle a deck of cards",
                "parameters": [
                    {
                        "type": "string",
                        "example": "9302b603-13bb-5275-a3b9-5fcefafa34e0",
                        "description": "id of deck",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "checks the server's version",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "Version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/build.Info"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "build.Info": {
            "type": "object",
            "properties": {
                "dirty_build": {
                    "type": "boolean"
                },
                "go_version": {
                    "type": "string"
                },
                "last_commit": {
                    "type": "string"
                },
                "revision": {
                    "type": "string"
                }
            }
        },
        "http.Card": {
            "description": "represents a card",
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "AS"
                },
                "suit": {
                    "type": "string",
                    "example": "SPADES"
                },
                "value": {
                    "type": "string",
                    "example": "ACE"
                }
            }
        },
        "http.Cards": {
            "description": "represents a collection of cards",
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.Card"
                    }
                }
            }
        },
        "http.DeckClosed": {
            "description": "represents a closed deck of cards",
            "type": "object",
            "properties": {
                "deck_id": {
                    "type": "string",
                    "example": "f6afe993-9847-508e-b206-2487f1ef5a3c"
                },
                "remaining": {
                    "type": "integer",
                    "example": 1
                },
                "shuffled": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "http.DeckOpened": {
            "description": "represents a opened deck of cards",
            "type": "object",
            "properties": {
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/http.Card"
                    }
                },
                "deck_id": {
                    "type": "string",
                    "example": "f6afe993-9847-508e-b206-2487f1ef5a3c"
                },
                "remaining": {
                    "type": "integer",
                    "example": 1
                },
                "shuffled": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "http.Error": {
            "description": "defines the structure for a failed response",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.1.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Deckr",
	Description:      "A REST API for playing with a deck of cards.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
