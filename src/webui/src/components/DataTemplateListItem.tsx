import { Link } from "react-router-dom";
import { DataTemplate } from "../types";

type Props = { template: DataTemplate };

const DataTemplateGroupItem = ({ template }: Props) => {
	const { _id: templateId } = template;

	return (
		<div className="my-6 rounded border bg-gray-300 px-4 py-4">
			<Link to={templateId || ""} className="mb-2">
				{template.name}
			</Link>
			<p>
				Contains {template.columns.length} column
				{template.columns.length === 0 || template.columns.length > 1 ? "s" : ""}
			</p>
		</div>
	);
};

export default DataTemplateGroupItem;
