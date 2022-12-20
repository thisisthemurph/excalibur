import { defaultConfig, urls } from ".";
import { HateoasResponse } from "./types";
import { DataTemplate, DataTemplateList, DataTemplateSchema } from "../types";

export async function getAllDataTemplates(): Promise<DataTemplateList> {
	const response = await fetch(urls.dataTemplate, defaultConfig);

	if (!response.ok) {
		throw new Error("There has been an issue fetching the data templates.");
	}

	return await response.json();
}

export async function getDataTemplate(id: string): Promise<DataTemplate> {
	const url = `${urls.dataTemplate}/${id}`;

	const response = await fetch(url, defaultConfig);
	if (!response.ok) {
		throw new Error("There has been an issue fetching the specific data template.");
	}

	return await response.json();
}

export async function createDataTemplate(dt: DataTemplate): Promise<HateoasResponse> {
	// Validate the input data
	const result = DataTemplateSchema.safeParse(dt);
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

export async function updateDataTemplate(dt: DataTemplate): Promise<HateoasResponse> {
	const result = DataTemplateSchema.safeParse(dt);
	if (!result.success) {
		throw new Error(result.error.message);
	}

	const { _id: id, ...data } = result.data;
	if (!id) {
		throw new Error("Could not determine DataTemplate to update");
	}

	const config = {
		...defaultConfig,
		method: "PUT",
		body: JSON.stringify(data),
	};

	const response = await fetch(`${urls.dataTemplate}/${id}`, config);
	if (!response.ok) {
		throw new Error("There has been an issue updating the data template.");
	}

	return response.json();
}
