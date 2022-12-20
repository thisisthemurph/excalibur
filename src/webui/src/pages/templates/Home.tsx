import { Link } from "react-router-dom";
import { getAllDataTemplates } from "../../api/dataTemplate";

import { useQuery } from "react-query";
import DataTemplateGroup from "../../components/DataTemplateGroup";

const DataTemplateHomePage = () => {
	const { isLoading, isError, data, error } = useQuery("templates", getAllDataTemplates);

	return (
		<>
			<h1 className="px-wrap py-wrap">Templates</h1>
			<pre>{JSON.stringify(error, null, 2)}</pre>
			<main className="px-wrap">
				<Link to="/template/create">Create a new template</Link>

				{data !== undefined && (
					<DataTemplateGroup templates={data} isLoading={isLoading} isError={isError} />
				)}
			</main>
		</>
	);
};

export default DataTemplateHomePage;
