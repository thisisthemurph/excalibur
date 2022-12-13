import BaseButton from "./BaseButton";
import { ButtonProps, ButtonType } from "./props";

const PrimaryButton = (props: ButtonProps) => {
	return <BaseButton {...{ ...props, type: ButtonType.Primary }} />;
};

export default PrimaryButton;
