import { defaultConfig, urls } from ".";
import {
	DataTemplate,
	DataTemplateListModel,
	DataTemplateModel,
	HateoasResponseType,
} from "./types";

export async function createDataTemplate(dt: DataTemplateModel): Promise<HateoasResponseType> {
	// Validate the input data
	const result = DataTemplate.safeParse(dt);
	if (!result.success) {
		throw new Error(result.error.message);
	}

	const config = {
		...defaultConfig,
		method: "POST",
		body: JSON.stringify(result.data),
	};

	const response = await fetch(urls.dataTemplate, config);
	if (!response.ok) {
		throw new Error("There has been an issue creating the new data template.");
	}

	return response.json();
}

export async function getAllDataTemplates(): Promise<DataTemplateListModel> {
	const response = await fetch(urls.dataTemplate, defaultConfig);

	if (!response.ok) {
		throw new Error("There has been an issue fetching the data templates.");
	}

	return await response.json();
}

export async function getDataTemplate(id: string): Promise<DataTemplateModel> {
	const url = `{urls.dataTemplate}/${id}`;

	const response = await fetch(url, defaultConfig);
	if (!response.ok) {
		throw new Error("There has been an issue fetching the specific data template.");
	}

	return await response.json();
}
