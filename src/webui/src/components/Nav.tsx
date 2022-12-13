import { Link } from "react-router-dom";

const Nav = () => {
	return (
		<nav className="flex gap-8 items-center">
			<Link to="/" className="no-underline text-lg">
				Home
			</Link>
			<Link to="/template" className="no-underline text-lg">
				Templates
			</Link>
		</nav>
	);
};

export default Nav;
