import { useQuery } from "react-query";
import { useParams, useNavigate } from "react-router-dom";

import Modal from "../../components/Modal";
import useModal from "../../hooks/useModal";
import FileUploadForm from "../../components/FileUploadForm";
import TemplateConfigForm from "../../components/TemplateConfig/form";
import { getDataTemplate, updateDataTemplate } from "../../api/dataTemplate";

const EditDataTemplatePage = () => {
	const { id } = useParams();
	const navigate = useNavigate();
	const { open, toggle, ref } = useModal();

	if (!id) {
		navigate("/template");
		return null;
	}

	const { data: template, status } = useQuery({
		queryKey: ["template", id],
		queryFn: ({ queryKey }) => getDataTemplate(queryKey[1]),
	});

	return (
		<>
			<header className="flex flex-col gap-2 px-wrap py-wrap">
				<h1>Edit DataTemplate</h1>

				<button className="btn btn__basic" onClick={() => toggle(true)}>
					Upload new data
				</button>
				<pre>{JSON.stringify({ showModal: open }, null, 2)}</pre>
			</header>

			{status === "loading" && <h2>Loading...</h2>}
			{template && (
				<TemplateConfigForm config={template} onSubmitFn={updateDataTemplate} controls={true} />
			)}

			<Modal
				ref={ref}
				title="Upload a file for a document of some type that is known to the system that we know"
				onClose={() => toggle(!open)}
			>
				<FileUploadForm dataTemplateId={id} />
			</Modal>
		</>
	);
};

export default EditDataTemplatePage;
