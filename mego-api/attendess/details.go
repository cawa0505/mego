package attendess

import (
	"errors"
	"github.com/mhewedy/ews"
	"github.com/mhewedy/ews/ewsutil"
)

func getAttendeeDetails(c ews.Client, e string) (*AttendeeDetails, error) {

	if len(attendeesDB) == 0 {
		return nil, errors.New("attendees db is empty, build the index first by invoking the search api")
	}

	attendee := attendeesDB[email(e)]
	// check cache
	if attendee.details != nil {
		return attendee.details, nil
	}

	persona, err := ewsutil.GetPersona(c, attendee.PersonaId)
	if err != nil {
		return nil, err
	}

	base64, err := ewsutil.GetUserPhotoBase64(c, e)
	if err != nil {
		base64 = "NA"
	}

	attendee.Image = base64
	attendee.Department = persona.Department // make sure attendee continue to return even if not returned at index time
	attendee.details = &AttendeeDetails{
		Attendee:            attendee,
		BusinessPhoneNumber: persona.BusinessPhoneNumbers.PhoneNumberAttributedValue.Value.Number,
		MobilePhone:         persona.MobilePhones.PhoneNumberAttributedValue.Value.Number,
		OfficeLocation:      persona.OfficeLocations.StringAttributedValue.Value,
	}
	attendeesDB[email(e)] = attendee // cache

	return attendee.details, nil
}
