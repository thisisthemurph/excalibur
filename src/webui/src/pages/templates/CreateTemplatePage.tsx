import { FormSchemaType } from "../../components/TemplateConfig/types";
import TemplateConfigForm from "../../components/TemplateConfig/form";

const formDefaultConfiguration: FormSchemaType = {
	name: "Company registered vehicles",
	columns: [
		{
			dataType: "String",
			originalName: "first_name",
			prettyName: "First Name",
		},
		{
			dataType: "String",
			originalName: "last_name",
			prettyName: "Surname",
		},
		{
			dataType: "String",
			originalName: "vreg",
			prettyName: "Registration Number",
		},
		{
			dataType: "Number",
			originalName: "value",
			prettyName: "Value",
		},
		{
			dataType: "Boolean",
			originalName: "Vehicle currently in use",
			prettyName: "In use?",
		},
	],
};

const CreateTemplatePage = () => {
	return (
		<>
			<h1 className="px-wrap">Create a new template</h1>
			<TemplateConfigForm config={formDefaultConfiguration} />
		</>
	);
};

export default CreateTemplatePage;
