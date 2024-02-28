package models

type DoctorSignup struct {
	DoctorID  string `json:"doctorid" bson:"doctorid"`
	Name        string `json:"name" bson:"name" binding:"required"`
	EmailID     string `json:"emailid" bson:"emailid" binding:"required"`
	Password    string `json:"password" bson:"password" binding:"required"`
	CreatedTime string `json:"createdtime" bson:"createdtime"`
}

type DoctorSignin struct {
	EmailId  string `json:"emailid" bson:"emailid" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DoctorToken struct {
	DoctorID string `json:"doctorid" bson:"doctorid"`
	EmailID    string `json:"emailid" bson:"emailid" binding:"required"`
	Token      string `json:"token" bson:"token" binding:"required"`
}
