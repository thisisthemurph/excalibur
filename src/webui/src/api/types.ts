import { z } from "zod";

export const HateoasResponse = z.object({
	id: z.string(),
	links: z.array(
		z.object({
			href: z.string(),
			rel: z.string(),
			type: z.string(),
		}),
	),
});

export const ColumnDataTypeEnum = z.enum(["String", "Number", "Boolean"]);

export const DataTemplateColumn = z.object({
	originalName: z.string(),
	prettyName: z.string().optional(),
	dataType: ColumnDataTypeEnum,
});

export const DataTemplate = z.object({
	_id: z.string().optional(),
	name: z.string(),
	columns: z.array(DataTemplateColumn),
});

export const DataTemplateList = z.array(DataTemplate);

export type HateoasResponseType = z.infer<typeof HateoasResponse>;
export type DataTemplateModel = z.infer<typeof DataTemplate>;
export type DataTemplateListModel = z.infer<typeof DataTemplateList>;
