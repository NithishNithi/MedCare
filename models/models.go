package models

type MailGunEmail struct {
	RecipientEmail string `json:"recipientemail"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}

type ListBlogs struct {
	BlogID      string `json:"blogid" bson:"blogid"`
	Pictures    string `json:"pictures" bson:"pictures"`
	Title       string `json:"title" bson:"title"`
	Content     string `json:"content" bson:"content"`
	Author      string `json:"author" bson:"author"`
	CreatedTime string `json:"createdtime" bson:"createdtime"`
}

type Count struct {
	PatientCount int64 `json:"patientcount" bson:"patientcount"`
	DoctorCount  int64 `json:"doctorcount" bson:"doctorcount"`
}
