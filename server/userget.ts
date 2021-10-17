// this file was automatically generated, DO NOT EDIT

// helpers
const maxUnixTSInSeconds = 9999999999;

function ParseDate(d: Date | number | string): Date {
	if (d instanceof Date) return d;
	if (typeof d === 'number') {
		if (d > maxUnixTSInSeconds) return new Date(d);
		return new Date(d * 1000); // go ts
	}
	return new Date(d);
}

function ParseNumber(v: number | string, isInt = false): number {
	if (!v) return 0;
	if (typeof v === 'number') return v;
	return (isInt ? parseInt(v) : parseFloat(v)) || 0;
}

function FromArray<T>(Ctor: { new (v: any): T }, data?: any[] | any, def = null): T[] | null {
	if (!data || !Object.keys(data).length) return def;
	const d = Array.isArray(data) ? data : [data];
	return d.map((v: any) => new Ctor(v));
}

function ToObject(o: any, typeOrCfg: any = {}, child = false): any {
	if (o == null) return null;
	if (typeof o.toObject === 'function' && child) return o.toObject();

	switch (typeof o) {
		case 'string':
			return typeOrCfg === 'number' ? ParseNumber(o) : o;
		case 'boolean':
		case 'number':
			return o;
	}

	if (o instanceof Date) {
		return typeOrCfg === 'string' ? o.toISOString() : Math.floor(o.getTime() / 1000);
	}

	if (Array.isArray(o)) return o.map((v: any) => ToObject(v, typeOrCfg, true));

	const d: any = {};

	for (const k of Object.keys(o)) {
		const v: any = o[k];
		if (v === undefined) continue;
		if (v === null) continue;
		d[k] = ToObject(v, typeOrCfg[k] || {}, true);
	}

	return d;
}

// structs
// struct2ts:server/api.UserGet
class UserGet {
	email: string;
	name: string;
	systemTags: string[] | null;
	userTags: string[] | null;
	lastModified: string;
	userId: string;

	constructor(data?: any) {
		const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
		this.email = ('email' in d) ? d.email as string : '';
		this.name = ('name' in d) ? d.name as string : '';
		this.systemTags = ('systemTags' in d) ? d.systemTags as string[] : null;
		this.userTags = ('userTags' in d) ? d.userTags as string[] : null;
		this.lastModified = ('lastModified' in d) ? d.lastModified as string : '';
		this.userId = ('userId' in d) ? d.userId as string : '';
	}

	toObject(): any {
		const cfg: any = {};
		return ToObject(this, cfg);
	}
}

// struct2ts:server/api.User
class User {
	email: string;
	name: string;
	systemTags: string[] | null;
	userTags: string[] | null;

	constructor(data?: any) {
		const d: any = (data && typeof data === 'object') ? ToObject(data) : {};
		this.email = ('email' in d) ? d.email as string : '';
		this.name = ('name' in d) ? d.name as string : '';
		this.systemTags = ('systemTags' in d) ? d.systemTags as string[] : null;
		this.userTags = ('userTags' in d) ? d.userTags as string[] : null;
	}

	toObject(): any {
		const cfg: any = {};
		return ToObject(this, cfg);
	}
}

// exports
export {
	UserGet,
	User,
	ParseDate,
	ParseNumber,
	FromArray,
	ToObject,
};
