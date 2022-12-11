import { Link, Outlet } from "react-router-dom";

const Layout = () => {
  return (
    <main>
      <header>
        <nav>
          <Link to="/">Home</Link>
          <Link to="/template">Templates</Link>
        </nav>
      </header>

      <section>
        <Outlet />
      </section>
    </main>
  );
};

export default Layout;
