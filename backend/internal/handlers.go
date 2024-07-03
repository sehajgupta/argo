package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) GetTrip(c *gin.Context) {
	var trip Trip
	tripId := c.Param("tripId")
	err := s.DB.QueryRow("SELECT * FROM trips WHERE id = $1", tripId).Scan(
		&trip.ID, &trip.UserID, &trip.StartLocation, &trip.EndLocation,
		&trip.StartTime, &trip.EndTime, &trip.DriverInfo, &trip.LicensePlate, &trip.Status,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trip not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (s *Server) GetVehicle(c *gin.Context) {
	var vehicle Vehicle
	vehicleId := c.Param("vehicleId")
	err := s.DB.QueryRow("SELECT * FROM vehicles WHERE id = $1", vehicleId).Scan(
		&vehicle.ID, &vehicle.TripID, &vehicle.CurrentLocation, &vehicle.ETA,
		&vehicle.Status, &vehicle.DriverInfo, &vehicle.LicensePlate,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func (s *Server) SubscribeTripUpdates(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	tripId := c.Param("tripId")
	for {
		// Fetch updated trip information from the database
		var trip Trip
		err := s.DB.QueryRow("SELECT * FROM trips WHERE id = $1", tripId).Scan(
			&trip.ID, &trip.UserID, &trip.StartLocation, &trip.EndLocation,
			&trip.StartTime, &trip.EndTime, &trip.DriverInfo, &trip.LicensePlate, &trip.Status,
		)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Error fetching trip updates"))
			return
		}

		// Send updated trip information to the client
		err = conn.WriteJSON(trip)
		if err != nil {
			return
		}

		// Poll every 5 seconds
		time.Sleep(5 * time.Second)
	}
}

func (s *Server) SubscribeVehicleUpdates(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	vehicleId := c.Param("vehicleId")
	for {
		// Fetch updated vehicle information from the database
		var vehicle Vehicle
		err := s.DB.QueryRow("SELECT * FROM vehicles WHERE id = $1", vehicleId).Scan(
			&vehicle.ID, &vehicle.TripID, &vehicle.CurrentLocation, &vehicle.ETA,
			&vehicle.Status, &vehicle.DriverInfo, &vehicle.LicensePlate,
		)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Error fetching vehicle updates"))
			return
		}

		// Send updated vehicle information to the client
		err = conn.WriteJSON(vehicle)
		if err != nil {
			return
		}

		// Poll every 5 seconds
		time.Sleep(5 * time.Second)
	}
}

func (s *Server) GetAllTripsForUser(c *gin.Context) {
	userId := c.Param("userId")
	rows, err := s.DB.Query("SELECT * FROM trips WHERE user_id = $1", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var trips []Trip
	for rows.Next() {
		var trip Trip
		if err := rows.Scan(
			&trip.ID, &trip.UserID, &trip.StartLocation, &trip.EndLocation,
			&trip.StartTime, &trip.EndTime, &trip.DriverInfo, &trip.LicensePlate, &trip.Status,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		trips = append(trips, trip)
	}

	c.JSON(http.StatusOK, trips)
}

func (s *Server) CreateTrip(c *gin.Context) {
	var trip Trip
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.DB.QueryRow(
		"INSERT INTO trips (user_id, start_location, end_location, start_time, end_time, driver_info, license_plate, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		trip.UserID, trip.StartLocation, trip.EndLocation, trip.StartTime, trip.EndTime, trip.DriverInfo, trip.LicensePlate, trip.Status,
	).Scan(&trip.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trip)
}

func (s *Server) UpdateTrip(c *gin.Context) {
	var trip Trip
	tripId := c.Param("tripId")
	if err := c.ShouldBindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := s.DB.Exec(
		"UPDATE trips SET user_id = $1, start_location = $2, end_location = $3, start_time = $4, end_time = $5, driver_info = $6, license_plate = $7, status = $8 WHERE id = $9",
		trip.UserID, trip.StartLocation, trip.EndLocation, trip.StartTime, trip.EndTime, trip.DriverInfo, trip.LicensePlate, trip.Status, tripId,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, trip)
}

func (s *Server) GetAllVehicles(c *gin.Context) {
	rows, err := s.DB.Query("SELECT * FROM vehicles")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var vehicles []Vehicle
	for rows.Next() {
		var vehicle Vehicle
		if err := rows.Scan(
			&vehicle.ID, &vehicle.TripID, &vehicle.CurrentLocation, &vehicle.ETA,
			&vehicle.Status, &vehicle.DriverInfo, &vehicle.LicensePlate,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		vehicles = append(vehicles, vehicle)
	}

	c.JSON(http.StatusOK, vehicles)
}
