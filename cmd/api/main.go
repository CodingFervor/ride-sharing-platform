package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"github.com/CodingFervor/ride-sharing-platform/internal/database"
)

func main() {
	r := gin.Default()
	r.Use(CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().Format(time.RFC3339)})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/auth/login", Login)
		api.POST("/auth/register", Register)
		api.POST("/auth/driver/register", DriverRegister)

		auth := api.Group("/")
		auth.Use(AuthMiddleware())
		{
			// Ride booking
			auth.POST("/rides/request", RequestRide)
			auth.GET("/rides/:id", GetRide)
			auth.PUT("/rides/:id/cancel", CancelRide)
			auth.PUT("/rides/:id/accept", AcceptRide)
			auth.PUT("/rides/:id/start", StartRide)
			auth.PUT("/rides/:id/complete", CompleteRide)
			auth.GET("/rides/history", RideHistory)
			auth.GET("/rides/active", ActiveRide)

			// Drivers
			auth.GET("/drivers/profile", GetDriverProfile)
			auth.PUT("/drivers/profile", UpdateDriverProfile)
			auth.PUT("/drivers/location", UpdateDriverLocation)
			auth.PUT("/drivers/online", GoOnline)
			auth.PUT("/drivers/offline", GoOffline)
			auth.GET("/drivers/earnings", DriverEarnings)

			// Vehicles
			auth.GET("/vehicles", ListVehicles)
			auth.POST("/vehicles", RegisterVehicle)
			auth.PUT("/vehicles/:id", UpdateVehicle)

			// Passengers
			auth.GET("/passengers/profile", GetPassengerProfile)
			auth.PUT("/passengers/profile", UpdatePassengerProfile)

			// Payments & Wallet
			auth.GET("/wallet", GetWallet)
			auth.POST("/wallet/topup", TopUpWallet)
			auth.POST("/wallet/withdraw", WithdrawWallet)
			auth.GET("/wallet/transactions", WalletTransactions)
			auth.POST("/payments/process", ProcessPayment)

			// Ratings
			auth.POST("/rides/:id/rate", RateRide)
			auth.GET("/ratings/driver/:id", GetDriverRatings)

			// Fare estimation
			auth.POST("/fare/estimate", EstimateFare)
			auth.GET("/fare/surge", GetSurgePricing)

			// Admin
			auth.GET("/admin/drivers", AdminListDrivers)
			auth.PUT("/admin/drivers/:id/status", AdminUpdateDriverStatus)
			auth.GET("/admin/rides", AdminListRides)
			auth.GET("/admin/analytics", AdminAnalytics)
			auth.GET("/admin/revenue", AdminRevenueReport)
		}
	}
	log.Println("Ride-Sharing Platform starting on :8080")
	addr := ":" + strconv.Itoa(8080)
	srv := &http.Server{Addr: addr, Handler: r}
	go func() {
		logger.Info("server listening", "port", 8080)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("forced shutdown", "error", err)
	}
	logger.Info("server exited")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == "OPTIONS" { c.AbortWithStatus(http.StatusNoContent); return }
		c.Next()
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" { c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"}); return }
		c.Next()
	}
}

func Login(c *gin.Context)              { c.JSON(http.StatusOK, gin.H{"message": "login"}) }
func Register(c *gin.Context)           { c.JSON(http.StatusCreated, gin.H{"message": "registered"}) }
func DriverRegister(c *gin.Context)     { c.JSON(http.StatusCreated, gin.H{"message": "driver registered"}) }
func RequestRide(c *gin.Context)        { c.JSON(http.StatusCreated, gin.H{"message": "ride requested"}) }
func GetRide(c *gin.Context)            { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func CancelRide(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"message": "ride cancelled"}) }
func AcceptRide(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"message": "ride accepted"}) }
func StartRide(c *gin.Context)          { c.JSON(http.StatusOK, gin.H{"message": "ride started"}) }
func CompleteRide(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"message": "ride completed"}) }
func RideHistory(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func ActiveRide(c *gin.Context)         { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func GetDriverProfile(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func UpdateDriverProfile(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "profile updated"}) }
func UpdateDriverLocation(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "location updated"}) }
func GoOnline(c *gin.Context)           { c.JSON(http.StatusOK, gin.H{"message": "driver online"}) }
func GoOffline(c *gin.Context)          { c.JSON(http.StatusOK, gin.H{"message": "driver offline"}) }
func DriverEarnings(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func ListVehicles(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func RegisterVehicle(c *gin.Context)    { c.JSON(http.StatusCreated, gin.H{"message": "vehicle registered"}) }
func UpdateVehicle(c *gin.Context)      { c.JSON(http.StatusOK, gin.H{"message": "vehicle updated"}) }
func GetPassengerProfile(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func UpdatePassengerProfile(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "profile updated"}) }
func GetWallet(c *gin.Context)          { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func TopUpWallet(c *gin.Context)        { c.JSON(http.StatusOK, gin.H{"message": "wallet topped up"}) }
func WithdrawWallet(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"message": "withdrawal initiated"}) }
func WalletTransactions(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func ProcessPayment(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"message": "payment processed"}) }
func RateRide(c *gin.Context)           { c.JSON(http.StatusOK, gin.H{"message": "ride rated"}) }
func GetDriverRatings(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func EstimateFare(c *gin.Context)       { c.JSON(http.StatusOK, gin.H{"data": gin.H{"fare": 0}}) }
func GetSurgePricing(c *gin.Context)    { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func AdminListDrivers(c *gin.Context)   { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func AdminUpdateDriverStatus(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "driver status updated"}) }
func AdminListRides(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
func AdminAnalytics(c *gin.Context)     { c.JSON(http.StatusOK, gin.H{"data": gin.H{}}) }
func AdminRevenueReport(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"data": []gin.H{}}) }
