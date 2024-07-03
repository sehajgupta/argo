import React from 'react';

const Trip = ({ trip }) => {
  return (
    <div>
      <h2>Trip #{trip.id}</h2>
      <p>Start Location: {trip.start_location}</p>
      <p>End Location: {trip.end_location}</p>
      <p>Driver: {trip.driver_info}</p>
      <p>License Plate: {trip.license_plate}</p>
      <p>Status: {trip.status}</p>
    </div>
  );
};

export default Trip;
