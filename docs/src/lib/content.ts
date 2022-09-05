import fs from "fs";
import path from "path";
import type { Tag } from "./packages";

const contentPath: string =
	(import.meta.env.VITE_CONTENT_PATH as string) || "content";

export async function allTags(): Promise<Tag[]> {
	return fs.promises
		.readdir(contentPath)
		.then((files) =>
			files.map(async (f) => [
				f,
				await fs.promises.readFile(path.join(contentPath, f)),
			]),
		)
		.then((e) => Promise.all(e))
		.then((r) => {
			const tags: Tag[] = [];
			for (const [f, data] of r) {
				const tag: Tag = JSON.parse(data.toString());
				tag.path = f.toString().replace(/\.json/, "");
				tags.push(tag);
			}
			return tags.sort((a, b) => a.weight - b.weight);
		});
}

export async function tag(tag: string): Promise<Tag | null> {
	if (tag === "all") {
		return (await allTags()).reduce((t1, t2) => {
			return {
				path: "/all",
				weight: 0,
				name: "All Packages",
				preamble: t1.preamble,
				packages: [...t1.packages, ...t2.packages],
			};
		});
	}
	return fs.promises
		.readFile(path.join(contentPath, tag + ".json"))
		.then((buf) => JSON.parse(buf.toString()))
		.then((t) => {
			return {
				path: tag,
				name: t.name,
				preamble: t.preamble,
				packages: t.packages,
			};
		})
		.catch(() => null);
}
