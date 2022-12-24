export type User = {
	id: number;
	email: string;
	name:string;
}


export type Device = {
	id:           number;
	created_at:   string;
	updated_at:   string;
	deleted_at:   null;
	name:         string;
	deviceID:     string;
	connected:    boolean;
	systemInfo:   SystemInfo;
	systemInfoId: number;
}

export type SystemInfo = {
	id:         number;
	os:         string;
	ip:         string;
	macAddress: string;
	hostName:   string;
	cores:      number;
	memory:     string;
	disk:       string;
}
