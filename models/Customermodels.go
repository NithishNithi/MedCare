package models

type CustomerSignUp struct {
	CustomerID  string `json:"customerid" bson:"customerid"`
	Name        string `json:"name" bson:"name" binding:"required"`
	EmailID     string `json:"emailid" bson:"emailid" binding:"required"`
	Password    string `json:"password" bson:"password" binding:"required"`
	CreatedTime string `json:"createdtime" bson:"createdtime"`
}

type CustomerSignIn struct {
	EmailId  string `json:"emailid" bson:"emailid" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type PatientToken struct {
	CustomerID string `json:"customerid" bson:"customerid"`
	EmailID    string `json:"emailid" bson:"emailid" binding:"required"`
	Token      string `json:"token" bson:"token" binding:"required"`
}

type BookAppointment struct {
	Token                string               `json:"token,omitempty" bson:"token,omitempty"`
	CustomerID           string               `json:"customerid,omitempty" bson:"customerid,omitempty"`
	Name                 string               `json:"name" bson:"name" binding:"required"`
	DateOfBirth          string               `json:"date_of_birth" bson:"date_of_birth" binding:"required"`
	Gender               string               `json:"gender" bson:"gender" binding:"required"`
	PhoneNumber          string               `json:"phone_number" bson:"phone_number" binding:"required"`
	EmailID              string               `json:"emailid" bson:"emailid" binding:"required"`
	Address              string               `json:"address" bson:"address" binding:"required"`
	ReasonForAppointment ReasonForAppointment `json:"reason_for_appointment" bson:"reason_for_appointment"`
	PreferredDateTime    PreferredDateTime    `json:"preferred_datetime" bson:"preferred_datetime"`
	MedicalHistory       MedicalHistory       `json:"medical_history,omitempty" bson:"medical_history,omitempty"`
	EmergencyContact     EmergencyContact     `json:"emergency_contact,omitempty" bson:"emergency_contact,omitempty"`
	AdditionalNotes      AdditionalNotes      `json:"additional_notes,omitempty" bson:"additional_notes,omitempty"`
}

type ReasonForAppointment struct {
	BriefDescription string `json:"brief_description" bson:"brief_description" binding:"required"`
	Symptoms         string `json:"symptoms" bson:"symptoms" binding:"required"`
}

type PreferredDateTime struct {
	Time string `json:"time" bson:"time"`
}

type MedicalHistory struct {
	ExistingConditions      string `json:"existing_conditions" bson:"existing_conditions" binding:"required"`
	Medications             string `json:"medications" bson:"medications" binding:"required"`
	PastSurgeriesTreatments string `json:"past_surgeries_treatments" bson:"past_surgeries_treatments" binding:"required"`
}

type EmergencyContact struct {
	EmergencyContactName string `json:"emergency_contact_name" bson:"emergency_contact_name" binding:"required"`
	Relationship         string `json:"relationship" bson:"relationship" binding:"required"`
	EmergencyPhoneNumber string `json:"emergency_phone_number" bson:"emergency_phone_number" binding:"required"`
}

type AdditionalNotes struct {
	Notes string `json:"notes" bson:"notes"`
}
