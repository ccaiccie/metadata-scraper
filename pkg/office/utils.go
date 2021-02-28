package office

import (
	"archive/zip"
	"encoding/xml"
)

func NewProperties(r *zip.Reader) (*officeCoreProperty, *officeAppProperty, error) {
	var coreProps officeCoreProperty
	var appProps officeAppProperty

	for _, f := range r.File {
		switch f.Name {
		case "docProps/core.xml":
			if err := process(f, &coreProps); err != nil {
				return nil, nil, err
			}
		case "docProps/app.xml":
			if err := process(f, &appProps); err != nil {
				return nil, nil, err
			}
		default:
			continue
		}
	}
	return &coreProps, &appProps, nil
}

func process(f *zip.File, prop interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}

	defer func() {
		err := rc.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err := xml.NewDecoder(rc).Decode(prop); err != nil {
		return err
	}
	return nil
}
