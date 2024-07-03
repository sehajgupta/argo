"use client"; 
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';
import io from 'socket.io-client';

interface Trip {
  id: number;
  start_location: string;
  end_location: string;
  driver_info: string;
  license_plate: string;
  status: string;
}

interface Vehicle {
  current_location: string;
  eta: string;
  status: string;
  driver_info: string;
  license_plate: string;
}

const TripDetails: React.FC = () => {
  const router = useRouter();
  const { id } = router.query;
  const [trip, setTrip] = useState<Trip | null>(null);
  const [vehicle, setVehicle] = useState<Vehicle | null>(null);

  useEffect(() => {
    const fetchTrip = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/trips/${id}`);
        setTrip(response.data);
      } catch (error) {
        console.error('Error fetching trip:', error);
      }
    };

    if (id) {
      fetchTrip();
      
      const socket = io(`http://localhost:8080`, {
        path: `/trips/${id}/subscribe`
      });

      socket.on('vehicleUpdate', data => {
        setVehicle(data);
      });

      return () => {
        socket.disconnect();
      };
    }
  }, [id]);

  if (!trip) return <div>Loading...</div>;

  return (
    <div>
      <h1>Trip #{trip.id}</h1>
      <p>Start Location: {trip.start_location}</p>
      <p>End Location: {trip.end_location}</p>
      <p>Driver: {trip.driver_info}</p>
      <p>License Plate: {trip.license_plate}</p>
      <p>Status: {trip.status}</p>
      {vehicle && (
        <div>
          <h2>Vehicle Updates</h2>
          <p>Current Location: {vehicle.current_location}</p>
          <p>ETA: {vehicle.eta}</p>
          <p>Status: {vehicle.status}</p>
          <p>Driver: {vehicle.driver_info}</p>
          <p>License Plate: {vehicle.license_plate}</p>
        </div>
      )}
    </div>
  );
};

export default TripDetails;
