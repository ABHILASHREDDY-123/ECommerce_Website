import "./App.css";
import Navbar from "./components/navbar";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/Home";
import About from "./components/about";
import DisplayMovie from "./components/DisplayMovie";

const dummy_data=[
  {
    "name" : "Abhi",
    "gender" : "male",
    "age" : "21",
    "url" : "https://shorturl.at/dfvRY"
  },
  {
    "name" :"tulasi",
    "gender":"female",
    "age":"20",
    "url" :"https://shorturl.at/bcGMQ"

  }]
function App() {
  return (
    <div className="App">
      {Navbar()}
      {/* <Navbar/>   ===> same as above(other syntax) */}
      <Router>
        <Routes>
          <Route path="/home" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/display" element={<DisplayMovie movies={dummy_data}/>}/>
        </Routes>
      </Router>
    </div>

  );
}

export default App;
import "./App.css";
import Navbar from "./components/navbar";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/Home";
import About from "./components/about";
import DisplayMovie from "./components/DisplayMovie";
import { useState } from "react";

const dummy_data=[
  {
    "name" : "Abhi",
    "gender" : "male",
    "age" : "21",
    "url" : "https://shorturl.at/dfvRY"
  },
  {
    "name" :"tulasi",
    "gender":"female",
    "age":"20",
    "url" :"https://shorturl.at/bcGMQ"


  }]
function App() {
    const [movies,setMovies] = useState([]); 

  return (
    <div className="App">
      <Navbar setMovies={setMovies}/>
      {/* <Navbar/>   ===> same as above(other syntax) */}
      <Router>
        <Routes>
          <Route path="/home" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/display" element={<DisplayMovie movies={dummy_data}/>}/>
        </Routes>
      </Router>
    </div>

  );
}

export default App;
