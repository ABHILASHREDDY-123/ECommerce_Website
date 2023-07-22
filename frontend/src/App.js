import NavBar from "./components/navbar";
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import About from "./components/about";
const Home = () => {
  return (
    <h1>HOME</h1>
  )
}
function App() {
  return (
    <div className="hello">
      {NavBar()}
      <Router>
        <Routes>
          <Route path="/home" element={<Home/>} />
          <Route path="/about" element={<About/>} />
        </Routes>
      </Router>
    </div>

  );
}

export default App;
