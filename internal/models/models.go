package models

import (
	"time"
)

// Client Data struct
type Client struct {
	ID        int
	Name      string
	Email     string
	Phone     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Provider data struct
type Provider struct {
	ID          int
	Name        string
	Username    string
	Description string
	Password    string
	Photo       []byte
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Specialist data struct
type Specialist struct {
	ID        int
	Username  string
	Password  string
	Photo     []byte
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Notification data struct
type Notification struct {
	ID        int
	Condition string
	Name      string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Prescription data struct
type Prescription struct {
	ID          int
	FormID      string
	Name        string
	Institution string
	Physician   string
	Location    string
	Image       []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Forms meta data struct
type FormMetaData struct {
	Message string
	Button  string
	Url     string
	Section string
}

type PageTitle struct {
	Main        string
	Sub         string
	Description string
	PlaceHolder string
}

// Client data struct
type Customer struct {
	ID              int
	CustomerId      string
	CustImage       []byte
	IDType          string
	CardImage       []byte
	FirstName       string
	LastName        string
	Phone           int
	HouseAddress    string
	Location        string
	Landmark        string
	Status          string
	Months          int
	Agreement       string
	UserId          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CustImgString   string
	CardImgString   string
	CreatedAtString string
	UpdatedAtString string
}

// Client Witness data struct
type Witness struct {
	ID              int
	CustomerId      string
	FirstName       string
	LastName        string
	Phone           int
	Terms           string
	Image           []byte
	UserId          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ImageString     string
	CreatedAtString string
	UpdatedAtString string
}

// Client's items credited data struct
type Item struct {
	ID              int       `json:"-"`
	CustomerId      string    `json:"customerId"`
	Serial          string    `json:"serial"`
	Price           float64   `json:"price"`
	Quantity        int       `json:"quantity"`
	Total           float64   `json:"-"`
	Deposit         float64   `json:"deposit"`
	Balance         float64   `json:"balance"`
	UserId          int       `json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
	CreatedAtString string
	UpdatedAtString string
}

type Friends struct {
	Name  string `json:"name"`
	Place string `json:"place"`
	Year  string `json:"year"`
}

// Payment is the model type for payment database
type Payments struct {
	CustomerId      string
	Month           string
	Amount          float64
	Date            time.Time
	UserId          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DateString      string
	CreatedAtString string
	UpdatedAtString string
}

type Purchases struct {
	Serial          string
	Quantity        int
	Amount          float64
	UserId          int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedAtString string
	UpdatedAtString string
}
