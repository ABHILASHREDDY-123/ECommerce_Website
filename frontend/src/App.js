import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/home";
import Navbar from "./components/navbar"
import Products from "./components/products";
import About from "./components/about";
import Cart from "./components/cart";
import { connect } from "react-redux";
function App() { 
  
  return (
    <div className="App">
    <Navbar/>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About/>}/>
          <Route path="/search" element={<Products/>}/>
          <Route path="/cart" element={<Cart/>}/>
        </Routes>
      </Router>
    </div>

  );
}

const mapStateToProps = (state) =>{
  return {
    ...state
  }
}

export default connect(mapStateToProps)(App);
