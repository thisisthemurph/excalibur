import * as z from "zod";

const ButtonVariantEnum = z.enum(["Danger", "Outlined", "Primary", "Standard", "Submit"]);

const P = z.object({
  text: z.string(),
  variant: ButtonVariantEnum,
});

type Props = z.infer<typeof P>;
type ButtonVariantType = z.infer<typeof ButtonVariantEnum>;

const determineClassName = (v: ButtonVariantType) => {
  switch (v) {
    case ButtonVariantEnum.Values.Standard:
      return "btn";
    case ButtonVariantEnum.Values.Submit:
      return `btn btn__${ButtonVariantEnum.Values.Primary.toLowerCase()}`;
    default:
      return `btn btn__${v.toLowerCase()}`;
  }
};

const Button = ({ text, variant }: Props) => {
  const className = determineClassName(variant);

  if (variant === ButtonVariantEnum.Values.Submit) {
    return <input type="submit" className={className} value={text} />;
  }

  return <button className={className}>{text}</button>;
};

export default Button;
