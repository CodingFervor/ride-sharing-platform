package model

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Phone     string    `json:"phone" db:"phone"`
	Password  string    `json:"-" db:"password"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Avatar    string    `json:"avatar" db:"avatar"`
	Role      string    `json:"role" db:"role"` // passenger, driver, admin
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Driver struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	LicenseNo    string    `json:"license_no" db:"license_no"`
	LicensePhoto string    `json:"license_photo" db:"license_photo"`
	Rating       float64   `json:"rating" db:"rating"`
	TotalTrips   int       `json:"total_trips" db:"total_trips"`
	Earnings     float64   `json:"earnings" db:"earnings"`
	IsOnline     bool      `json:"is_online" db:"is_online"`
	Latitude     float64   `json:"latitude" db:"latitude"`
	Longitude    float64   `json:"longitude" db:"longitude"`
	Verified     bool      `json:"verified" db:"verified"`
	Status       string    `json:"status" db:"status"` // active, suspended, rejected
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type Vehicle struct {
	ID           int64     `json:"id" db:"id"`
	DriverID     int64     `json:"driver_id" db:"driver_id"`
	PlateNumber  string    `json:"plate_number" db:"plate_number"`
	Brand        string    `json:"brand" db:"brand"`
	Model        string    `json:"model" db:"model"`
	Year         int       `json:"year" db:"year"`
	Color        string    `json:"color" db:"color"`
	Type         string    `json:"type" db:"type"` // sedan, suv, van, luxury
	Capacity     int       `json:"capacity" db:"capacity"`
	Verified     bool      `json:"verified" db:"verified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type Ride struct {
	ID            int64      `json:"id" db:"id"`
	PassengerID   int64      `json:"passenger_id" db:"passenger_id"`
	DriverID      *int64     `json:"driver_id" db:"driver_id"`
	VehicleID     *int64     `json:"vehicle_id" db:"vehicle_id"`
	PickupLat     float64    `json:"pickup_lat" db:"pickup_lat"`
	PickupLng     float64    `json:"pickup_lng" db:"pickup_lng"`
	PickupAddr    string     `json:"pickup_addr" db:"pickup_addr"`
	DestLat       float64    `json:"dest_lat" db:"dest_lat"`
	DestLng       float64    `json:"dest_lng" db:"dest_lng"`
	DestAddr      string     `json:"dest_addr" db:"dest_addr"`
	Distance      float64    `json:"distance" db:"distance"` // km
	Duration      int        `json:"duration" db:"duration"` // minutes
	Fare          float64    `json:"fare" db:"fare"`
	SurgeFactor   float64    `json:"surge_factor" db:"surge_factor"`
	Status        string     `json:"status" db:"status"` // requested, accepted, arrived, in_progress, completed, cancelled
	VehicleType   string     `json:"vehicle_type" db:"vehicle_type"`
	PaymentMethod string     `json:"payment_method" db:"payment_method"`
	RequestedAt   time.Time  `json:"requested_at" db:"requested_at"`
	AcceptedAt    *time.Time `json:"accepted_at" db:"accepted_at"`
	StartedAt     *time.Time `json:"started_at" db:"started_at"`
	CompletedAt   *time.Time `json:"completed_at" db:"completed_at"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
}

type Wallet struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Balance   float64   `json:"balance" db:"balance"`
	Currency  string    `json:"currency" db:"currency"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Transaction struct {
	ID          int64     `json:"id" db:"id"`
	WalletID    int64     `json:"wallet_id" db:"wallet_id"`
	Type        string    `json:"type" db:"type"` // topup, payment, withdrawal, refund
	Amount      float64   `json:"amount" db:"amount"`
	Description string    `json:"description" db:"description"`
	RideID      *int64    `json:"ride_id" db:"ride_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Rating struct {
	ID        int64     `json:"id" db:"id"`
	RideID    int64     `json:"ride_id" db:"ride_id"`
	RaterID   int64     `json:"rater_id" db:"rater_id"`
	RateeID   int64     `json:"ratee_id" db:"ratee_id"`
	Score     int       `json:"score" db:"score"` // 1-5
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type SurgeZone struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Factor    float64   `json:"factor" db:"factor"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	Radius    float64   `json:"radius" db:"radius"` // km
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
