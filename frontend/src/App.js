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
    "url" : "https://media.licdn.com/dms/image/C5603AQHj2zcj1G-73g/profile-displayphoto-shrink_800_800/0/1661696842947?e=2147483647&v=beta&t=pu6rL-Eh0dTKXw6EfwIRJVGZ0HOvNjcgBo2QnbzqAmo"
  },
  {
    "name" :"tulasi",
    "gender":"female",
    "age":"20",
    "url" :"data:image/jpeg;base64,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"

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
