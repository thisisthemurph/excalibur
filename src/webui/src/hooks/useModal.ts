import { useRef, useState } from "react";

const useModal = () => {
	const [open, setOpen] = useState(false);
	const ref = useRef<HTMLDialogElement>(null);

	const toggle = (togglingOpen: boolean): void => {
		setOpen(togglingOpen);

		if (togglingOpen) {
			ref.current?.showModal();
			return;
		}

		ref.current?.close();
	};

	return { open, toggle, ref };
};

export default useModal;
