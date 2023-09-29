import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

const backendBaseUrl = 'http://localhost:8080';

function App() {
  const [students, setStudents] = useState([]);

  useEffect(() => {
    const fetchStudents = async () => {
      try {
        const response = await axios.get(`${backendBaseUrl}/students`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        console.log(response)
        const data = response.data;

        setStudents(data);
      } catch (error) {
        console.error('Error fetching students:', error);
      }
    };
    fetchStudents();
  });

  return (
    <div className="App">
      <div className="student-list-container">
        <h2>Student List</h2>
        <ul>
          {console.log(students.length)}
          {students.map((student, index) => (
            <li key={index}>{`${student.name} ${student.surname}`}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
