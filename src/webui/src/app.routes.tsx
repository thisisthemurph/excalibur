import { createBrowserRouter } from "react-router-dom";

import Layout from "./components/Layout";
import FileUploadPage from "./pages/FileUploadPage";
import HomePage from "./pages/HomePage";
import CreateDataTemplatePage from "./pages/templates/Create";
import EditDataTemplatePage from "./pages/templates/Edit";
import DataTemplateHomePage from "./pages/templates/Home";

export const routes = createBrowserRouter([
	{
		path: "/",
		element: <Layout />,
		children: [
			{
				index: true,
				element: <HomePage />,
			},
			{
				path: "/template",
				element: <DataTemplateHomePage />,
			},
			{
				path: "/template/create",
				element: <CreateDataTemplatePage />,
			},
			{
				path: "/template/:id",
				element: <EditDataTemplatePage />,
			},
			{
				path: "/upload",
				element: <FileUploadPage />,
			},
		],
	},
]);
