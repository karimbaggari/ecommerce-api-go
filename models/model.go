package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID  `json:" _id" bson:"_id"`
	First_Name *string `json:"first_name" bson:"first_name" validate:"required,min:1,max:100"`
	Last_Name *string `json:"last_name" bson:"last_name" validate:"required,min:1,max:100"`
	Password *string `json:"password" bson:"password" validate:"required,min:1,max:100"`
	Email *string `json:"email" bson:"email" validate:"required,min:1,max:100"`
	Phone *string `json:"phone" bson:"phone" validate:"required,min:1,max:100"`
	Token *string `json:"token" bson:"token"`
	Refresh_Token *string `json:"refresh_token" bson:"refresh_token"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
	Uptaded_At time.Time `json:"uptaded_at" bson:"uptaded_at"`
	User_ID string `json:"user_id" bson:"user_id"`
	UserCart []ProductUser `json:"userCart" bson:"userCart"`
	Address_details []Address `json:"addressDetails" bson:"addressDetails"`
	Order_Status	[]Order `json:"orderStatus" bson:"orderStatus"`
}

type Product struct {
	Product_ID primitive.ObjectID `json:" _id" bson:"_id"`
	Product_Name *string `json:"product_name" bson:"product_name"`
	Price *uint64 `json:"price"`
	Rating *uint8 `json:"rating"`
	Image *string `json:"image"`
}

type ProductUser struct {
	Product_ID primitive.ObjectID `json:" _id" bson:"_id"`
	Product_Name *string `json:"productName" bson:"productName"`
	Price *uint64	`json:"price" bson:"price"`
	Rating *uint8	`json:"rating" bson:"rating"`
	Image *string	`json:"image" bson:"image"`
}


type Address struct {
	Product_ID primitive.ObjectID `json:" _id" bson:"_id"`
	House    *string            `json:"house" bson:"house"`
	Street   *string            `json:"street" bson:"street"`
	City     *string            `json:"city" bson:"city"`
	Pincode  *string            `json:"pincode" bson:"pincode"`
}


type Order struct {
	Order_ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	Order_At       time.Time          `json:"orderedAt" bson:"orderedAt"`
	Price          int                `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}


type Payment struct {
	Digital bool
	COD bool
}