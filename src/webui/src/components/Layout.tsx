import { Link, Outlet } from "react-router-dom";
import Logo from "./Logo";
import Nav from "./Nav";

const Layout = () => {
	return (
		<main>
			<header className="flex justify-between px-wrap py-wrap">
				<Logo />
				<Nav />
			</header>

			<section>
				<Outlet />
			</section>
		</main>
	);
};

export default Layout;
