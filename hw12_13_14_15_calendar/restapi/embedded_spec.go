// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API для работы с событиями календаря",
    "title": "Calendar API",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/events": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Создать новое событие",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewEvent"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Событие успешно создано",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "description": "Идентификатор события",
                  "type": "integer"
                }
              },
              "$ref": "#/definitions/EventCreated"
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events-by-day": {
      "get": {
        "summary": "Получить все события за день",
        "operationId": "getEventsByDay",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "description": "Дата",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events-by-month": {
      "get": {
        "summary": "Получить все события за месяц",
        "operationId": "getEventsByMonth",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "description": "Дата начала месяца",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events-by-week": {
      "get": {
        "summary": "Получить все события за неделю",
        "operationId": "getEventsByWeek",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "description": "Дата начала недели",
            "name": "weekStart",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events/{id}": {
      "put": {
        "summary": "Обновить событие",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewEvent"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Событие успешно обновлено"
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "summary": "Удалить событие по идентификатору",
        "parameters": [
          {
            "type": "integer",
            "description": "ID события",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Успех"
          },
          "404": {
            "description": "Событие не найдено"
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Event": {
      "type": "object",
      "properties": {
        "daysAmountTillNotify": {
          "type": "integer",
          "example": 5
        },
        "description": {
          "type": "string",
          "example": "description"
        },
        "end": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "ownerId": {
          "type": "integer",
          "example": 1
        },
        "start": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "title": {
          "type": "string",
          "example": "Event"
        }
      }
    },
    "EventCreated": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "example": 1
        }
      }
    },
    "NewEvent": {
      "type": "object",
      "properties": {
        "daysAmountTillNotify": {
          "type": "integer",
          "example": 5
        },
        "description": {
          "type": "string",
          "example": "description"
        },
        "end": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "ownerId": {
          "type": "integer",
          "example": 1
        },
        "start": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "title": {
          "type": "string",
          "example": "title"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API для работы с событиями календаря",
    "title": "Calendar API",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/events": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Создать новое событие",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewEvent"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Событие успешно создано",
            "schema": {
              "type": "object",
              "properties": {
                "id": {
                  "description": "Идентификатор события",
                  "type": "integer"
                }
              },
              "$ref": "#/definitions/EventCreated"
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events-by-day": {
      "get": {
        "summary": "Получить все события за день",
        "operationId": "getEventsByDay",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "description": "Дата",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events-by-month": {
      "get": {
        "summary": "Получить все события за месяц",
        "operationId": "getEventsByMonth",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "description": "Дата начала месяца",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events-by-week": {
      "get": {
        "summary": "Получить все события за неделю",
        "operationId": "getEventsByWeek",
        "parameters": [
          {
            "type": "string",
            "format": "date",
            "description": "Дата начала недели",
            "name": "weekStart",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успех",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/events/{id}": {
      "put": {
        "summary": "Обновить событие",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewEvent"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Событие успешно обновлено"
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "summary": "Удалить событие по идентификатору",
        "parameters": [
          {
            "type": "integer",
            "description": "ID события",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Успех"
          },
          "404": {
            "description": "Событие не найдено"
          },
          "500": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Event": {
      "type": "object",
      "properties": {
        "daysAmountTillNotify": {
          "type": "integer",
          "example": 5
        },
        "description": {
          "type": "string",
          "example": "description"
        },
        "end": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "id": {
          "type": "integer",
          "example": 1
        },
        "ownerId": {
          "type": "integer",
          "example": 1
        },
        "start": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "title": {
          "type": "string",
          "example": "Event"
        }
      }
    },
    "EventCreated": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "example": 1
        }
      }
    },
    "NewEvent": {
      "type": "object",
      "properties": {
        "daysAmountTillNotify": {
          "type": "integer",
          "example": 5
        },
        "description": {
          "type": "string",
          "example": "description"
        },
        "end": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "ownerId": {
          "type": "integer",
          "example": 1
        },
        "start": {
          "type": "string",
          "format": "date",
          "example": "1900-01-01"
        },
        "title": {
          "type": "string",
          "example": "title"
        }
      }
    }
  }
}`))
}
