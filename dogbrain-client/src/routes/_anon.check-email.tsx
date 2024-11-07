import { Heading, Text } from "@/components/ui/typography";
import { createFileRoute, Link } from "@tanstack/react-router";

export const Route = createFileRoute("/_anon/check-email")({
	component: CheckEmail,
});

function CheckEmail() {
	return (
		<div className="grid place-items-center grow">
			<div className="flex flex-col gap-4 text-center">
				<Heading>Check your email</Heading>
				<Text>
					We've sent you a verification link. Please check your email to verify
					your account.
				</Text>
				<Text variant="muted">
					Didn't receive the email? Check your spam folder or{" "}
					<Link to="/sign-up">try signing up again</Link>.
				</Text>
			</div>
		</div>
	);
}
