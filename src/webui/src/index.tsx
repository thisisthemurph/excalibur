import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";
import { routes } from "./app.routes";
import "./index.css";

const queryClient = new QueryClient();
const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement);

root.render(
	<React.StrictMode>
		<QueryClientProvider client={queryClient}>
			<RouterProvider router={routes} />
		</QueryClientProvider>
	</React.StrictMode>,
);
