// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// InventoryGetReader is a Reader for the InventoryGet structure.
type InventoryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInventoryGetOK creates a InventoryGetOK with default headers values
func NewInventoryGetOK() *InventoryGetOK {
	return &InventoryGetOK{}
}

/*InventoryGetOK handles this case with default header values.

successful operation
*/
type InventoryGetOK struct {
	Payload map[string]int32
}

func (o *InventoryGetOK) Error() string {
	return fmt.Sprintf("[GET /store/inventory][%d] inventoryGetOK  %+v", 200, o.Payload)
}

func (o *InventoryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
