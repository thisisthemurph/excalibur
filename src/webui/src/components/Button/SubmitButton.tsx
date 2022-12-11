import BaseButton from "./BaseButton";
import { ButtonProps, ButtonType } from "./props";

const SubmitButton = (props: ButtonProps) => {
  return <BaseButton {...{ ...props, type: ButtonType.Submit }} />;
};

export default SubmitButton;
