import {
	type ReactNode,
	createContext,
	useState,
	useEffect,
	useCallback,
} from "react";

export interface AuthContext {
	isAuthenticated: boolean;
	login: (username: string) => Promise<void>;
	logout: () => Promise<void>;
	user: string | null;
}

export const AuthContext = createContext<AuthContext | null>(null);

const key = "tanstack.auth.user";

function getStoredUser() {
	return localStorage.getItem(key);
}

function setStoredUser(user: string | null) {
	if (user) {
		localStorage.setItem(key, user);
	} else {
		localStorage.removeItem(key);
	}
}

export function AuthProvider({ children }: { children: ReactNode }) {
	const [user, setUser] = useState<string | null>(getStoredUser());
	const isAuthenticated = !!user;

	const logout = useCallback(async () => {
		setStoredUser(null);
		setUser(null);
	}, []);

	const login = useCallback(async (username: string) => {
		setStoredUser(username);
		setUser(username);
	}, []);

	useEffect(() => {
		setUser(getStoredUser());
	}, []);

	return (
		<AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
			{children}
		</AuthContext.Provider>
	);
}
