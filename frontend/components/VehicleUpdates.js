import React, { useEffect, useState } from 'react';
import { WebSocket } from 'react-websocket';

const VehicleUpdates = ({ tripId }) => {
  const [vehicle, setVehicle] = useState(null);

  const handleData = (data) => {
    setVehicle(JSON.parse(data));
  };

  return (
    <div>
      <h2>Vehicle Updates for Trip #{tripId}</h2>
      <WebSocket url={`ws://localhost:8080/trips/${tripId}/subscribe`} onMessage={handleData} />
      {vehicle ? (
        <div>
          <p>Current Location: {vehicle.current_location}</p>
          <p>ETA: {vehicle.eta}</p>
          <p>Status: {vehicle.status}</p>
          <p>Driver: {vehicle.driver_info}</p>
          <p>License Plate: {vehicle.license_plate}</p>
        </div>
      ) : (
        <p>Loading vehicle updates...</p>
      )}
    </div>
  );
};

export default VehicleUpdates;
