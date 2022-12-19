import { DataTemplateListModel } from "../api/types";
import DataTemplateListItem from "./DataTemplateListItem";

type Props = {
	isLoading: boolean;
	isError: boolean;
	templates: DataTemplateListModel;
};

const DataTemplateList = ({ templates, isLoading, isError }: Props) => {
	return (
		<section className="mt-16">
			<h2>Your data templates</h2>

			{isLoading && <p>Loading</p>}
			{isError && <p>There is an error</p>}

			{!isLoading && templates?.map((dt, i) => <DataTemplateListItem key={i} template={dt} />)}
		</section>
	);
};

export default DataTemplateList;
