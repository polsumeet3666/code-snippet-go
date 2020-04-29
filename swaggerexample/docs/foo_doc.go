package docs

// swagger:route GET /pets pets users listPets
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Consumes:
//     - application/json
//     - multipart/formdata
//
//     Produces:
//     - application/json
//     - application/x-protobuf
//
//     Schemes: http, https, ws, wss

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response
// The error message
// in: body
// The validation message
//
// Required: true
// Example: Expected type int
// An optional field name to which this validation applies
