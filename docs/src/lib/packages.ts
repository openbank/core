export interface Page {
	name: string;
	path: string;
}

export interface Tag {
	path: string;
	name: string;
	preamble: string;
	weight: number;
	packages?: Package[];
}

export class Packages {
	public raw: Package[];
	public packages: { [key: string]: Package };
	constructor(pkgs: Package[]) {
		this.raw = pkgs;
		this.packages = pkgs.reduce(
			(pkgs, pkg) => ({
				...pkgs,
				[pkg.id]: pkg,
			}),
			{},
		);
	}

	public typ(t: Type): string {
		if (!t) {
			return "/";
		}

		switch (t.type) {
			case "ref":
				return this.ref(t);
			case "basic":
				return t.name;
			case "map":
				return `Map<${this.typ(t.key)}, ${this.typ(t.value)}>`;
			case "array":
				return `Array[${this.typ(t.value)}]`;
		}
		return "?" + t.type;
	}

	obj(t: Ref): Message {
		const rs = this.resolve(t);
		if (rs.type.type !== "message") {
			throw new Error(`${t.name} is not an message`);
		}
		return rs.type;
	}

	ref(t: Ref): string {
		const rs = this.resolve(t);
		return `[${rs.type.name}](/${rs.package.tag}#${t.name})`;
	}

	resolve(t: Ref): ResolvedRef {
		const i = t.name.lastIndexOf(".");
		const pkg = this.packages[t.name.substring(0, i)];
		const typeName = t.name.substring(i + 1);
		const typ = pkg.types[typeName];
		return {
			...t,
			package: pkg,
			type: typ,
		};
	}
}

export interface Package {
	name: string;
	id: string;
	tag: string;
	description: string;
	services: Service[];
	types: { [index: string]: Message | Enum };
}

export interface TypeMap {
	[index: string]: Message | Enum;
}

export interface Service {
	name: string;
	description: string;
	endpoints: Endpoint[];
}

export interface Endpoint {
	name: string;
	description: string;
	method: string;
	path: string;
	body_field: string;
	request?: Message;
	response?: Message | Ref;
	streaming_request: boolean;
	streaming_response: boolean;
}

export type Type = Message | Enum | Ref | Map | Array | Basic;

export interface Message {
	type: "message";
	name: string;
	description: string;
	fields: Field[];
}

export interface Field {
	name: string;
	description: string;
	type: Type;
}

export interface Enum {
	type: "enum";
	name: string;
	description: string;
	values: { value: string; description: string }[];
}

export interface Ref {
	type: "ref";
	name: string;
}

export interface ResolvedRef {
	name: string;
	package: Package;
	type: Message | Enum;
}

export interface Map {
	type: "map";
	key: Type;
	value: Type;
}

export interface Array {
	type: "array";
	value: Type;
}

export interface Basic {
	type: "basic";
	name: string;
	example: string;
}
