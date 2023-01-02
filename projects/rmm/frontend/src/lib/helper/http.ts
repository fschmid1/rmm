import { apiBase } from "../../vars";
import { toast } from '@zerodevx/svelte-toast';

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

export async function callDeviceFunction<T>(id: string, event: string, data = ''): Promise<{data: T, event: string, id: string}> {
	const response = await fetchWithToken(`${apiBase}/devices/functions`, {
    method: "POST",
    body: JSON.stringify({
      id,
      event,
      data
    }),
	})
	const json = await response.json();
	if (json.error) {
		toast.push(json.error, {});
		throw new Error(json.error);
	}
	return json;
}