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
	Token                   string `json:"token,omitempty" bson:"token,omitempty"`
	CustomerID              string `json:"customerid,omitempty" bson:"customerid,omitempty"`
	Name                    string `json:"name" bson:"name" binding:"required"`
	DateOfBirth             string `json:"dateofbirth" bson:"dateofbirth" binding:"required"`
	Gender                  string `json:"gender" bson:"gender" binding:"required"`
	PhoneNumber             string `json:"phonenumber" bson:"phonenumber" binding:"required"`
	EmailID                 string `json:"emailid" bson:"emailid" binding:"required"`
	Address                 string `json:"address" bson:"address" binding:"required"`
	BriefDescription        string `json:"briefdescription" bson:"briefdescription" binding:"required"`
	Symptoms                string `json:"symptoms" bson:"symptoms" binding:"required"`
	Date                    string `json:"date" bson:"date" binding:"required"`
	FromDateTime            string `json:"fromdatetime" bson:"fromdatetime"`
	ToDateTime              string `json:"todatetime" bson:"todatetime"`
	ExistingConditions      string `json:"existingconditions" bson:"existingconditions" binding:"required"`
	Medications             string `json:"medications" bson:"medications" binding:"required"`
	PastSurgeriesTreatments string `json:"pastsurgeriestreatments" bson:"pastsurgeriestreatments" binding:"required"`
	EmergencyContactName    string `json:"emergencycontactname" bson:"emergencycontactname" binding:"required"`
	EmergencyPhoneNumber    string `json:"emergencycontactphonenumber" bson:"emergencycontactphonenumber" binding:"required"`
	Notes                   string `json:"notes" bson:"notes"`
	DoctorSpecialization    string `json:"doctorspecialization" bson:"doctorspecialization"`
	PreferredDoctorID       string `json:"preferreddoctorid" bson:"preferreddoctorid"`
}
