export async function fetchWithToken(url: string, options: any, json = true) {
	if (!options.headers) options.headers = {};
	options.headers.Authorization= `Bearer ${localStorage.getItem('token')}`;
	if (json)
	options.headers["Content-Type"] = "application/json";

	const response = await fetch(url, options);
	if (response.status == 401) {
		localStorage.removeItem('token');
		location.href = '/login';
	}
	return response;
}