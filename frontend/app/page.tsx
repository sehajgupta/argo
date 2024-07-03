"use client"; 

import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Link from 'next/link';

interface Trip {
  id: number;
}

const Home: React.FC = () => {
  const [trips, setTrips] = useState<Trip[]>([]);

  useEffect(() => {
    const fetchTrips = async () => {
      try {
        const response = await axios.get('http://localhost:8080/users/1/trips');
        setTrips(response.data);
      } catch (error) {
        console.error('Error fetching trips:', error);
      }
    };

    fetchTrips();
  }, []);

  return (
    <div>
      <h1>Trips</h1>
      <ul>
        {trips.map(trip => (
          <li key={trip.id}>
            <Link href={`/trips/${trip.id}`}>
              <a>Trip #{trip.id}</a>
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Home;
