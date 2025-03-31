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
        "/admin/create-doctor-profile": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new doctor profile and update the user's role to \"doctor\"",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create a new doctor profile",
                "parameters": [
                    {
                        "description": "Doctor profile details",
                        "name": "doctorProfile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DoctorProfileCreateRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/admin/create-speciality": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new speciality with a name, description, and auto-generated slug",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create a new speciality",
                "parameters": [
                    {
                        "description": "Speciality details",
                        "name": "speciality",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SpecialityCreateRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/admin/get-all-users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "retrieve all users with pagination and search",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "retrieve all users with pagination and search",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search query (name or phone starts with)",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "admin",
                            "doctor",
                            "patient"
                        ],
                        "type": "string",
                        "description": "Search query (based on user_role)",
                        "name": "role",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/admin/get-specialities": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve specialities with pagination and search on the name of the speciality",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Retrieve specialities with pagination and search",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search query (name starts with)",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/otp/generate": {
            "post": {
                "description": "Generates an OTP for phone number verification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Generate OTP",
                "parameters": [
                    {
                        "description": "Phone Number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OTPRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/otp/verify": {
            "post": {
                "description": "Verifies the OTP entered by the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "OTP"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "Phone Number and OTP",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VerifyOTPRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/forget-password": {
            "post": {
                "description": "get forget password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get forget password",
                "parameters": [
                    {
                        "description": "Phone number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ForgetPasswordRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/get-access-token-by-refresh-token": {
            "post": {
                "description": "get access token by refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get access token by refresh token",
                "parameters": [
                    {
                        "description": "Refresh Token",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/login": {
            "post": {
                "description": "login users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login user",
                "parameters": [
                    {
                        "description": "Phone Number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/register": {
            "post": {
                "description": "creating users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "registering user",
                "parameters": [
                    {
                        "description": "Phone Number",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/reset-password": {
            "post": {
                "description": "get reset password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get reset password",
                "parameters": [
                    {
                        "description": "password and password_retype",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PasswordResetRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "The encrypted key for password reset",
                        "name": "key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/users/verify-access-token": {
            "post": {
                "description": "verify access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "verify access token",
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.DoctorProfileCreateRequest": {
            "type": "object",
            "required": [
                "experience_years",
                "speciality_id",
                "user_id"
            ],
            "properties": {
                "bio": {
                    "type": "string"
                },
                "experience_years": {
                    "type": "integer"
                },
                "speciality_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string",
                    "example": "123e4567-e89b-12d3-a456-426614174000"
                }
            }
        },
        "dto.ForgetPasswordRequest": {
            "type": "object",
            "properties": {
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "dto.OTPRequest": {
            "type": "object",
            "properties": {
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "dto.PasswordResetRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "password_retype": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.SpecialityCreateRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.UserLoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "dto.UserRegisterRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "verified_token": {
                    "type": "string"
                }
            }
        },
        "dto.VerifyOTPRequest": {
            "type": "object",
            "properties": {
                "otp_code": {
                    "type": "string"
                },
                "phone_number": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
