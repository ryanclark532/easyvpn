type DataStatus = 'loading' | 'ready' | 'error' | 'initial';

export type DataWithStatus<T> = {
	status: DataStatus;
	data: T | undefined;
};
