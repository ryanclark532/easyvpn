interface TypedFetchResponse<T> {
	data: T;
	status: number;
	error?: string;
}

export async function typedFetch<T>(
	url: string,
	options?: RequestInit
): Promise<TypedFetchResponse<T> | Error> {
	try {
		const response = await fetch(url, options);
		const data = (await response.json()) as T;
		const status = response.status;

		return { data, status };
	} catch (e: any) {
		return new Error(e);
	}
}
