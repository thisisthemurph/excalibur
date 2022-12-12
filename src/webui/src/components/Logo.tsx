import { CubeTransparentIcon } from "@heroicons/react/24/outline";
import { Link } from "react-router-dom";

const Logo = () => {
  return (
    <Link to="/" className="group flex items-center gap-4 cursor-pointer no-underline">
      <CubeTransparentIcon className="flex-none w-14 h-14 text-indigo-500 group-hover:text-indigo-700" />
      <span className="flex-none w-28 text-lg font-extrabold tracking-widest text-indigo-500 align-text-bottom group-hover:text-indigo-700">
        Excali<span className="text-gray-400 group-hover:text-gray-500">bur</span>
      </span>
    </Link>
  );
};

export default Logo;
