const Navbar = () => {
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
        </div>
      </ul>
    </nav>
  );
};
export default Navbar;
