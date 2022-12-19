import { DataTemplate } from "../types";

type Props = { template: DataTemplate };

const DataTemplateGroupItem = ({ template }: Props) => {
	return (
		<div className="my-6 rounded border bg-gray-300 px-4 py-4">
			<h3 className="mb-2">{template.name}</h3>
			<p>
				Contains {template.columns.length} column
				{template.columns.length === 0 || template.columns.length > 1 ? "s" : ""}
			</p>
		</div>
	);
};

export default DataTemplateGroupItem;
