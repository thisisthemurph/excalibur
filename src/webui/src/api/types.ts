import { z } from "zod";

export const ErrorResponseSchema = z.object({
	message: z.string(),
	status: z.number(),
	statusText: z.string(),
});

export const HateoasResponseSchema = z.object({
	id: z.string(),
	links: z.array(
		z.object({
			href: z.string(),
			rel: z.string(),
			type: z.enum([
				"GET",
				"HEAD",
				"POST",
				"PUT",
				"DELETE",
				"CONNECT",
				"OPTIONS",
				"TRACE",
				"PATCH",
			]),
		}),
	),
});

export const UploadStatusSchema = z.object({ status: z.string() });

export type ErrorResponse = z.infer<typeof ErrorResponseSchema>;
export type HateoasResponse = z.infer<typeof HateoasResponseSchema>;
export type UploadStatus = z.infer<typeof UploadStatusSchema>;
