interface TypedFetchResponse<T> {
	data: T;
	status: number;
	error?: string;
}

export async function typedFetch<T>(
	url: string,
	options?: RequestInit
): Promise<TypedFetchResponse<T>> {
	const response = await fetch(url, options);
	const data = (await response.json()) as T;
	const status = response.status;

	return { data, status };
}
