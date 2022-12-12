import { Link, Outlet } from "react-router-dom";
import Logo from "./Logo";

const Layout = () => {
  return (
    <main>
      <header className="flex justify-between px-wrap py-wrap">
        <Logo />
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
