package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupRouterGroup(r *gin.Engine) {
	// Setup route group for the API
	api := r.Group("/api")
	{
		// trip 
		api.GET("/trip/:tripId", s.GetTrip)
		api.GET("/trip/:tripId/subscribe", s.SubscribeTripUpdates)
		api.POST("/trip", s.CreateTrip)
		
		// vehicle
		api.GET("/vehicle/:vehicleId", s.GetVehicle)
		api.GET("/vehicle/:vehicleId/subscribe", s.SubscribeVehicleUpdates)
		api.GET("/vehicles", s.GetAllVehicles)
		
		// user
		api.GET("/user/:userId/trips", s.GetAllTripsForUser)
		api.PUT("/trip/:tripId", s.UpdateTrip)
	}
}