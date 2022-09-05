import { tag } from "$lib/content";
import type { EndpointOutput, RequestEvent } from "@sveltejs/kit";

export async function get({ params }: RequestEvent): Promise<EndpointOutput> {
	const t = await tag(params.page).catch((e) => console.log(e));
	if (!t) {
		return { status: 404, body: { message: "Tag not found", status: 404 } };
	}
	return {
		body: JSON.stringify(t),
	};
}
