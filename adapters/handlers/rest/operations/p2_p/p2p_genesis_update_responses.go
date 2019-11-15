//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// P2pGenesisUpdateOKCode is the HTTP code returned for type P2pGenesisUpdateOK
const P2pGenesisUpdateOKCode int = 200

/*P2pGenesisUpdateOK Alive and kicking!

swagger:response p2pGenesisUpdateOK
*/
type P2pGenesisUpdateOK struct {
}

// NewP2pGenesisUpdateOK creates P2pGenesisUpdateOK with default headers values
func NewP2pGenesisUpdateOK() *P2pGenesisUpdateOK {

	return &P2pGenesisUpdateOK{}
}

// WriteResponse to the client
func (o *P2pGenesisUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// P2pGenesisUpdateUnauthorizedCode is the HTTP code returned for type P2pGenesisUpdateUnauthorized
const P2pGenesisUpdateUnauthorizedCode int = 401

/*P2pGenesisUpdateUnauthorized Unauthorized update.

swagger:response p2pGenesisUpdateUnauthorized
*/
type P2pGenesisUpdateUnauthorized struct {
}

// NewP2pGenesisUpdateUnauthorized creates P2pGenesisUpdateUnauthorized with default headers values
func NewP2pGenesisUpdateUnauthorized() *P2pGenesisUpdateUnauthorized {

	return &P2pGenesisUpdateUnauthorized{}
}

// WriteResponse to the client
func (o *P2pGenesisUpdateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// P2pGenesisUpdateInternalServerErrorCode is the HTTP code returned for type P2pGenesisUpdateInternalServerError
const P2pGenesisUpdateInternalServerErrorCode int = 500

/*P2pGenesisUpdateInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response p2pGenesisUpdateInternalServerError
*/
type P2pGenesisUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewP2pGenesisUpdateInternalServerError creates P2pGenesisUpdateInternalServerError with default headers values
func NewP2pGenesisUpdateInternalServerError() *P2pGenesisUpdateInternalServerError {

	return &P2pGenesisUpdateInternalServerError{}
}

// WithPayload adds the payload to the p2p genesis update internal server error response
func (o *P2pGenesisUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *P2pGenesisUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p2p genesis update internal server error response
func (o *P2pGenesisUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *P2pGenesisUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}