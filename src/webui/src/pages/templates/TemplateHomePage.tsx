import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { getAllDataTemplates } from "../../api/dataTemplate";
import { DataTemplateListModel } from "../../api/types";

const TemplateHomePage = () => {
	const [dataTemplates, setDataTemplates] = useState<DataTemplateListModel>([]);

	useEffect(() => {
		let mounted = true;

		getAllDataTemplates()
			.then((templates) => {
				if (mounted) {
					setDataTemplates(templates);
				}
			})
			.catch((reason) => console.warn(reason));

		return () => {
			mounted = false;
		};
	}, []);

	return (
		<>
			<h1 className="px-wrap py-wrap">Templates</h1>
			<main className="px-wrap">
				<Link to="/template/create">Create a new template</Link>

				<section className="mt-16">
					<h2>Your data templates</h2>

					{dataTemplates.map((dt, i) => {
						return (
							<div key={i} className="border px-4 py-4 bg-gray-300 rounded">
								<h3 className="mb-2">{dt.name}</h3>
								<p>
									Contains {dt.columns.length} column
									{dt.columns.length === 0 || dt.columns.length > 1 ? "s" : ""}
								</p>
							</div>
						);
					})}
				</section>
			</main>
		</>
	);
};

export default TemplateHomePage;
