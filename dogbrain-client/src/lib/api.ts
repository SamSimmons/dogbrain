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

export async function logIn(email: string, password: string) {
	const res = await fetch(`${API_BASE}/login`, {
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

export async function logOut() {
	const res = await fetch(`${API_BASE}/logout`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
	});

	const data = await res.json();

	if (!res.ok) {
		throw new ApiError(res.status, data.error);
	}

	return data;
}

export async function forgotPassword(email: string) {
	const res = await fetch(`${API_BASE}/forgot-password`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ email }),
	});

	const data = await res.json();

	if (!res.ok) {
		throw new ApiError(res.status, data.error);
	}

	return data;
}

export async function resetPassword(password: string, token: string) {
	const res = await fetch(`${API_BASE}/reset-password`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ password, token }),
	});

	const data = await res.json();

	if (!res.ok) {
		throw new ApiError(res.status, data.error);
	}

	return data;
}
