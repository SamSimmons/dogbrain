import { AnonHeader } from "@/components/AnonHeader";
import { createFileRoute, Outlet } from "@tanstack/react-router";

export const Route = createFileRoute("/_anon")({
	component: AnonLayout,
});

function AnonLayout() {
	return (
		<div className="p-2 min-h-screen flex flex-col w-full bg-zinc-200 text-zinc-800">
			<AnonHeader />
			<Outlet />
		</div>
	);
}
