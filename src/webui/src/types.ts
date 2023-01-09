import { z } from "zod";

export const ColumnDataTypeEnum = z.enum(["String", "Number", "Boolean"]);

export const DataTemplateColumnSchema = z.object({
	originalName: z.string(),
	prettyName: z.string().optional(),
	dataType: ColumnDataTypeEnum,
});

export const DataTemplateSchema = z.object({
	id: z.string().optional(),
	name: z.string(),
	columns: z.array(DataTemplateColumnSchema),
});

export const DataTemplateListSchema = z.array(DataTemplateSchema);

export type DataTemplate = z.infer<typeof DataTemplateSchema>;
export type DataTemplateList = z.infer<typeof DataTemplateListSchema>;
export type DataTemplateColumn = z.infer<typeof DataTemplateColumnSchema>;
