import { Link } from "react-router-dom";
import { getAllDataTemplates } from "../../api/dataTemplate";

import { useQuery } from "react-query";
import DataTemplateGroup from "../../components/DataTemplateGroup";

const TemplateHomePage = () => {
	const { isLoading, isError, data } = useQuery("templates", getAllDataTemplates);

	return (
		<>
			<h1 className="px-wrap py-wrap">Templates</h1>
			<main className="px-wrap">
				<Link to="/template/create">Create a new template</Link>

				{data !== undefined && (
					<DataTemplateGroup templates={data} isLoading={isLoading} isError={isError} />
				)}
			</main>
		</>
	);
};

export default TemplateHomePage;
