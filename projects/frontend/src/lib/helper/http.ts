import { toast } from '@zerodevx/svelte-toast';
import { accessToken } from '../../stores';
import { apiBase } from '../../vars';
import { get } from 'svelte/store';

export async function fetchWithToken(url: string, options: any, json = true) {
    if (!options.headers) options.headers = {};
    options.headers['Authorization'] = `Bearer ${get(accessToken)}`;
    if (json) options.headers['Content-Type'] = 'application/json';

    const response = await fetch(url, options);
    if (response.status == 401) {
        location.href = '/login';
    }
	if (response.status != 200) {
		toast.push(await response.text(), {});
	}
    return response;
}

export async function callDeviceFunction<T>(
    id: string,
    event: string,
    data = '',
): Promise<{ data: T; event: string; id: string }> {
    const response = await fetchWithToken(`${apiBase}/devices/functions`, {
        method: 'POST',
        body: JSON.stringify({
            id,
            event,
            data,
        }),
    });
    if (response.status != 200) {
        toast.push(await response.text(), {});
        throw new Error('Error calling function');
    }
    const json = await response.json();
    if (json.error) {
        toast.push(json.error, {});
        throw new Error(json.error);
    }
    return json;
}
