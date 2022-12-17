import { ButtonProps, ButtonType } from "./props";

const BaseButton = ({ text, type, onClick }: ButtonProps) => {
	const className = type ? `btn btn__${type}` : "btn btn__basic";

	if (type === ButtonType.Submit) {
		return <input className={className} type="submit" value={text} />;
	}

	return (
		<button className={className} onClick={onClick}>
			{text}
		</button>
	);
};

export default BaseButton;
