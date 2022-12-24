import React, { FormEventHandler, useState } from "react";
import { uploadFile } from "../api/file";
import { ErrorResponseSchema } from "../api/types";

type Props = { dataTemplateId: string };

const FileUploadForm = ({ dataTemplateId }: Props) => {
	const [file, setFile] = useState<File | null>(null);
	const [error, setError] = useState<string | null>(null);

	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const fileList = e.currentTarget.files;
		setFile(fileList ? fileList[0] : null);
	};

	const handleSubmit: FormEventHandler<HTMLFormElement> = async (e) => {
		e.preventDefault();

		if (!file) {
			return;
		}

		const data = await uploadFile(dataTemplateId, file);
		const errResult = ErrorResponseSchema.safeParse(data);

		if (errResult.success) {
			setError(errResult.data.message);
			return;
		}

		setError(null);
		alert("The file has been uploaded");
	};

	return (
		<form encType="multipart/form-data" onSubmit={handleSubmit}>
			{error && (
				<section className="mb-8 rounded border-2 border-red-600 bg-red-400 px-4 py-4">
					<p className="font-semibold text-white">{error}</p>
				</section>
			)}

			<input type="file" name="fileUpload" onChange={handleFileChange} />
			<input className="btn btn__basic" type="submit" value="upload" />
		</form>
	);
};

export default FileUploadForm;
