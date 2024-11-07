import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_anon/")({
	component: HomeComponent,
});

function HomeComponent() {
	return (
		<div className="grid place-items-center grow">
			<h1>DOGBRAIN</h1>
		</div>
	);
}
