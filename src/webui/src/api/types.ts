import { z } from "zod";

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

export type HateoasResponse = z.infer<typeof HateoasResponseSchema>;
