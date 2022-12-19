import {
	DataTemplate,
	DataTemplateList,
	DataTemplateListModel,
	DataTemplateModel,
	HateoasResponseType,
} from "./types";

export async function createNewDataTemplate(
	dt: DataTemplateModel,
): Promise<DataTemplateModel | string> {
	const url = "http://localhost:8000/datatemplate";

	const result = DataTemplate.safeParse(dt);
	if (!result.success) {
		return result.error.message;
	}

	const config: RequestInit = {
		method: "POST",
		body: JSON.stringify(result.data),
		headers: {
			Accept: "application-json",
			"Content-Type": "application/json",
		},
	};

	return fetch(url, config)
		.then(async (response) => {
			if (response.ok) {
				return response.json();
			}

			throw Error(response.statusText);
		})
		.then(async (data: HateoasResponseType) => {
			console.log({ data });
			return data;
		})
		.catch((e) => {
			console.error(e);
			return e;
		});
}

export async function getAllDataTemplates(): Promise<DataTemplateListModel> {
	const url = "http://localhost:8000/datatemplate";

	const config: RequestInit = {
		method: "GET",
		headers: {
			Accept: "application-json",
			"Content-Type": "application/json",
		},
	};

	const response = await fetch(url, config);
	if (!response.ok) {
		return [];
	}

	const result = DataTemplateList.safeParse(await response.json());
	if (!result.success) {
		console.error("The result is bad");
		console.warn({ result });
		return [];
	}

	return result.data;
}

export async function getDataTemplate(id: string): Promise<DataTemplateModel> {
	const url = `http://localhost:8000/datatemplate/${id}`;

	const config: RequestInit = {
		method: "GET",
		headers: {
			Accept: "application-json",
			"Content-Type": "application/json",
		},
	};

	return fetch(url, config)
		.then((response) => {
			if (response.ok) {
				return response.json();
			}

			throw new Error(response.statusText);
		})
		.then((data: HateoasResponseType) => {
			return data;
		})
		.catch((e) => {
			return e;
		});
}
