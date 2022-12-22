export const defaultConfig: RequestInit = {
	method: "GET",
	headers: {
		Accept: "application-json",
		"Content-Type": "application/json",
	},
};

const baseUrl = process.env.REACT_APP_API_SERVER as string;
export const urls = {
	base: baseUrl,
	dataTemplate: `${baseUrl}/datatemplate`,
	fileUpload: `${baseUrl}/file/upload`,
};
