import { createBrowserRouter } from "react-router-dom";

import Layout from "./components/Layout";
import HomePage from "./pages/HomePage";
import CreateTemplatePage from "./pages/templates/CreateTemplatePage";
import TemplateHomePage from "./pages/templates/TemplateHomePage";

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
        element: <TemplateHomePage />,
      },
      {
        path: "/template/create",
        element: <CreateTemplatePage />,
      },
    ],
  },
]);
