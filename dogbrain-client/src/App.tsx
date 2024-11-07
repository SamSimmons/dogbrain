import { RouterProvider, createRouter } from "@tanstack/react-router";

import { useAuth } from "./hooks/useAuth";
import { AuthProvider } from "./providers/AuthProvider";
import { routeTree } from "./routeTree.gen";
import { QueryProvider } from "./providers/QueryProvider";

const router = createRouter({
	routeTree,
	defaultPreload: "intent",
	context: {
		// biome-ignore lint/style/noNonNullAssertion: gets set by AuthProvider
		auth: undefined!,
	},
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
	interface Register {
		router: typeof router;
	}
}

function InnerApp() {
	const auth = useAuth();
	return <RouterProvider router={router} context={{ auth }} />;
}

export function App() {
	return (
		<AuthProvider>
			<QueryProvider>
				<InnerApp />
			</QueryProvider>
		</AuthProvider>
	);
}
