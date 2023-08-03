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