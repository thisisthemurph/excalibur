export enum ButtonType {
	Basic = "basic",
	Submit = "submit",
	Primary = "primary",
}

export interface ButtonProps {
	text?: string;
	type?: ButtonType;
}
