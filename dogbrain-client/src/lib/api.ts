const API_BASE = import.meta.env.PROD
	? "https://api.dogbrain.app/api/v1"
	: "/api/v1";

export class ApiError extends Error {
	constructor(
		public status: number,
		message: string,
	) {
		super(message);
	}
}

export async function register(email: string, password: string) {
	const res = await fetch(`${API_BASE}/register`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ email, password }),
	});

	const data = await res.json();

	if (!res.ok) {
		throw new ApiError(res.status, data.error);
	}

	return data;
}

export async function verifyEmail(token: string) {
	const res = await fetch(`${API_BASE}/verify/${token}`, {
		method: "GET",
	});

	const data = await res.json();

	if (!res.ok) {
		throw new ApiError(res.status, data.error);
	}

	return data;
}
