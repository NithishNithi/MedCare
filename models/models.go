package models

type CustomerSignUp struct {
	CustomerID      string `json:"customerid" bson:"customerid"`
	Name            string `json:"name" bson:"name" binding:"required"`
	EmailID           string `json:"emailid" bson:"emailid" binding:"required,email"`
	Password        string `json:"password" bson:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" bson:"confirmpassword" binding:"required"`
	CreatedTime     string `json:"createdtime" bson:"createdtime"`
}

type CustomerSignIn struct {
	EmailId  string `json:"emailid" bson:"emailid" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
