import { urls } from ".";
import { UploadStatus, ErrorResponse, ErrorResponseSchema } from "./types";

export async function uploadFile(file: File): Promise<UploadStatus | ErrorResponse> {
	const formData = new FormData();
	formData.append("fileUpload", file);
	formData.append("fileName", file.name);

	const config: RequestInit = {
		method: "POST",
		body: formData,
	};

	const response = await fetch(urls.fileUpload, config);

	if (!response.ok) {
		const err = await response.json();
		const result = ErrorResponseSchema.safeParse(err);

		if (result.success) {
			return result.data;
		}

		return {
			message: "There has been a unexpected error uploading the file",
			status: 500,
			statusText: "Internal Server Error",
		};
	}

	return response.json();
}
