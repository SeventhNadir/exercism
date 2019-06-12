package erratum

import "fmt"

//Use opens a resource, calls Frob(input) on the result resource and then closes that resource
func Use(o ResourceOpener, input string) (err error) {
	var resource Resource

	resource, err = o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("panic error: %v", r)
			}
			if err, ok := r.(FrobError); ok {
				resource.Defrob(err.defrobTag)
			}
		}
		resource.Close()
	}()

	resource.Frob(input)
	return err
}
