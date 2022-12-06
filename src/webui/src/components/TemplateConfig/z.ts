import * as z from "zod";

export const DataTypeEnum = z.enum(["String", "Number", "Boolean"]);

export const FormColumnSchema = z.object({
  originalName: z.string(),
  prettyName: z.string(),
  dataType: DataTypeEnum,
});

export const defaultColumnObject: FormColumnSchemaType = {
  dataType: DataTypeEnum.Values.String,
  originalName: "",
  prettyName: "",
};

export const FormSchema = z.object({
  tableName: z.string(),
  columns: z.array(FormColumnSchema),
});

export type FormSchemaType = z.infer<typeof FormSchema>;
export type FormColumnSchemaType = z.infer<typeof FormColumnSchema>;
