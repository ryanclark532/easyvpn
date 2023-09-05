interface TypedFetchResponse<T> {
	data: T;
	status: number;
}

export async function typedFetch<T>(
	url: string,
	options?: RequestInit
): Promise<TypedFetchResponse<T>> {
	const response = await fetch(url, options);

	if (!response.ok) {
		throw new Error(`HTTP error! Status: ${response.status}`);
	}

	const data = (await response.json()) as T;
	const status = response.status;

	return { data, status };
}
