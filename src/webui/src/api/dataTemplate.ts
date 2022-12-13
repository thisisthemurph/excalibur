import { z } from "zod";

const DataTemplateColumn = z.object({
	name: z.string(),
});

const DataTemplate = z.object({
	name: z.string(),
	columns: z.array(DataTemplateColumn),
});

type DataTemplateModel = z.infer<typeof DataTemplate>;

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

	return fetch(url, config).then(async (response) => {
		if (response.ok) {
			const result = DataTemplate.safeParse(await response.json());
			if (!result.success) {
				return "Unexpected result returned";
			}

			return result.data;
		} else {
			return response.statusText;
		}
	});
}

export async function getDataTemplate(id: string): Promise<DataTemplateModel | null> {
	const url = `http://localhost:8000/datatemplate/${id}`;

	const config: RequestInit = {
		method: "GET",
		headers: {
			Accept: "application-json",
			"Content-Type": "application/json",
		},
	};

	const response = await fetch(url, config);
	if (!response.ok) {
		return null;
	}

	const result = DataTemplate.safeParse(await response.json());
	if (!result.success) {
		return null;
	}

	return result.data;
}
