package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*Item item

swagger:model item
*/
type Item struct {

	/* Completed completed
	 */
	Completed bool `json:"completed,omitempty"`

	/* Description description

	Required: true
	*/
	Description string `json:"description,omitempty"`

	/* ID id

	Read Only: true
	*/
	ID int64 `json:"id,omitempty"`
}

// Validate validates this item
func (m *Item) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Item) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}
