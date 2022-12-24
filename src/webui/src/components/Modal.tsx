import { XMarkIcon, ArrowUpIcon, ArrowDownIcon } from "@heroicons/react/24/outline";
import React from "react";
import { useEffect, useRef, useState } from "react";

type Props = {
	title: string;
	fullScreen?: boolean;
	onClose: () => void;
	children: JSX.Element;
};

type Ref = HTMLDialogElement;

// eslint-disable-next-line react/display-name
const Modal = React.forwardRef<Ref, Props>(({ title, fullScreen, onClose, children }, ref) => {
	const [isFullScreen, setIsFullScreen] = useState(fullScreen ?? false);

	const handleClose = () => {
		onClose();
	};

	return (
		<dialog
			ref={ref}
			className="w-max-10/12 overflow-y-auto overflow-x-hidden rounded-lg border border-gray-300 bg-white px-0 py-0 shadow-lg"
		>
			<div className="flex flex-col items-center justify-center "></div>
			<header className="flex w-full justify-between">
				<h5 className="w-full py-8 px-8 text-3xl">{title}</h5>
				<section className="flex w-auto justify-end">
					<button
						onClick={() => setIsFullScreen(!isFullScreen)}
						className="flex aspect-square h-12 w-12 items-center justify-center bg-gray-400 px-1 py-1 text-white opacity-50 hover:opacity-100"
					>
						{isFullScreen ? (
							<ArrowDownIcon className="h-8 w-8" />
						) : (
							<ArrowUpIcon className="h-8 w-8" />
						)}
					</button>
					<button
						onClick={handleClose}
						className="flex aspect-square h-12 w-12 items-center justify-center rounded-tr-sm bg-red-500 px-1 py-1 text-white opacity-50 hover:opacity-100"
					>
						<XMarkIcon className="h-8 w-8" />
					</button>
				</section>
			</header>
			<main className="w-full px-8">{children}</main>
			<footer className="flex w-full items-center justify-end gap-4 py-8 px-8">
				<button className="btn btn__basic" onClick={handleClose}>
					Close
				</button>
				<button className="btn btn__primary">OK</button>
			</footer>
		</dialog>
	);
});

export default Modal;
