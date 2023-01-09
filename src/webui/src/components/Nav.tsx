import { Link } from "react-router-dom";

const Nav = () => {
	return (
		<nav className="flex items-center gap-8">
			<Link to="/" className="text-lg no-underline">
				Home
			</Link>
			<Link to="/template" className="text-lg no-underline">
				Templates
			</Link>
		</nav>
	);
};

export default Nav;
