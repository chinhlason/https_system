import logo from './logo.svg';
import './App.css';
import {useState} from "react";

function App() {
  const [data, setData] = useState("");
  fetch("https://localhost:3000/api/test")
      .then((response) => {
            return response.json()
      })
      .then((result) => {
          setData(result.message)
      })
      .catch((error) => {
          console.log(error)
      })
  return (
      <h1>{data}</h1>
  );
}

export default App;
