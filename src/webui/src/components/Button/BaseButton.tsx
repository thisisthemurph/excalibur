import { ButtonProps, ButtonType } from "./props";

const BaseButton = ({ text, type }: ButtonProps) => {
  const className = type ? `btn btn__${type}` : "btn btn__basic";

  if (type === ButtonType.Submit) {
    return <input className={className} type="submit" value={text} />;
  }

  return <button className={className}>{text}</button>;
};

export default BaseButton;
