import { useQuery } from "react-query";
import { useParams, useNavigate } from "react-router-dom";
import { getDataTemplate, updateDataTemplate } from "../../api/dataTemplate";
import TemplateConfigForm from "../../components/TemplateConfig/form";

const EditDataTemplatePage = () => {
	const { id } = useParams();
	const navigate = useNavigate();

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
			<h1 className="px-wrap py-wrap">Edit DataTemplate</h1>
			{status === "loading" && <h2>Loading...</h2>}
			{template && (
				<TemplateConfigForm config={template} onSubmitFn={updateDataTemplate} controls={true} />
			)}
		</>
	);
};

export default EditDataTemplatePage;
