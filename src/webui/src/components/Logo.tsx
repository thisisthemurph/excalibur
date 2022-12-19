import { CubeTransparentIcon } from "@heroicons/react/24/outline";
import { Link } from "react-router-dom";

const Logo = () => {
	return (
		<Link to="/" className="group flex cursor-pointer items-center gap-4 no-underline">
			<CubeTransparentIcon className="h-14 w-14 flex-none text-indigo-500 group-hover:text-indigo-700" />
			<span className="w-28 flex-none align-text-bottom text-lg font-extrabold tracking-widest text-indigo-500 group-hover:text-indigo-700">
				Excali<span className="text-gray-400 group-hover:text-gray-500">bur</span>
			</span>
		</Link>
	);
};

export default Logo;
