<<<<<<< HEAD
import "../styles/navbar.css"

const Navbar = () => {
    return (
        <>
       <nav className="navbar navbar-expand-lg navbar-light bg-light">
  <a className="navbar-brand" href="/">Ecommerce </a>
  <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarText" aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
    <span className="navbar-toggler-icon"></span>
  </button>
  <div className="collapse navbar-collapse" id="navbarText">
    <ul className="navbar-nav mr-auto">
      
      <li className="nav-item">
        <a className="nav-link" href="/about">About</a>
      </li>
      <li className="nav-item">
        <a className="nav-link" href="/cart">Cart</a>
      </li>
    </ul>
    <span className="navbar-text">
      <input/>
      <button className="btn btn-outline-sucess" >
        <a href="/search">
        Search
        </a>
        </button>
    </span>
  </div>
</nav>
        </>
    )
}

export default Navbar;
=======
import Search from "./search";
const Navbar = (props) => {
  return (
    <nav class="navbar">
      <div class="logo">MUO</div>
      <ul class="nav-links">
        <div class="menu">
          <li>
            <a href="/home">Home</a>
          </li>
          <li>
            <a href="/about">About</a>
          </li>
          <li>
            <a href="/login">Login</a>
          </li>
          <li>
            <a href="/logout">Logout</a>
          </li>
          <li>
            <a href ="/display">Display</a>
          </li>
          <li>
            <Search setMovies={props.setMovies}/> 

            {/* same as calling a function */}
          </li>
        </div>
      </ul>
    </nav>
  );
};
export default Navbar;
>>>>>>> 4b7270c15204e76023b334ddfab845ef867e6cef
