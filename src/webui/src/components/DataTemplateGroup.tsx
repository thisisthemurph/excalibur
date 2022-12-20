import { DataTemplateList } from "../types";
import DataTemplateGroupItem from "./DataTemplateListItem";

type Props = {
	isLoading: boolean;
	isError: boolean;
	templates: DataTemplateList;
};

const DataTemplateGroup = ({ templates, isLoading, isError }: Props) => {
	return (
		<section className="mt-16">
			<h2>Your data templates</h2>

			{isLoading && <p>Loading</p>}
			{isError && <p>There is an error</p>}

			{!isLoading && templates?.map((dt, i) => <DataTemplateGroupItem key={i} template={dt} />)}
		</section>
	);
};

export default DataTemplateGroup;
