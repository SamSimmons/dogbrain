import { Button } from "@/components/ui/button";
import { Text, Heading } from "@/components/ui/typography";
import { ApiError, verifyEmail } from "@/lib/api";
import { useMutation } from "@tanstack/react-query";
import { createFileRoute, Link, useNavigate } from "@tanstack/react-router";
import { useEffect } from "react";

export const Route = createFileRoute("/_anon/verify/$token")({
	component: VerifyEmail,
});

function VerifyEmail() {
	const { token } = Route.useParams();
	const navigate = useNavigate();

	const mutation = useMutation({
		mutationFn: verifyEmail,
		onSuccess: () => {
			navigate({ to: "/login" });
		},
	});

	useEffect(() => {
		if (mutation.isIdle) {
			mutation.mutate(token);
		}
	}, [token, mutation]);

	if (mutation.isPending) {
		return (
			<div className="grid place-items-center grow">
				<Text>Verifying your email...</Text>
			</div>
		);
	}

	if (mutation.isError) {
		return (
			<div className="grid place-items-center grow">
				<div className="flex flex-col gap-4 text-center">
					<Heading>Verification failed</Heading>
					<Text variant="error">
						{mutation.error instanceof ApiError
							? mutation.error.message
							: "Something went wrong. Please try again."}
					</Text>
					<Button asChild>
						<Link to="/login">Log In</Link>
					</Button>
					<Button asChild>
						<Link to="/sign-up">Sign up again</Link>
					</Button>
				</div>
			</div>
		);
	}

	return (
		<div className="grid place-items-center grow">
			<div className="flex flex-col gap-4 text-center">
				<Heading>Email verified!</Heading>
				<Text>Your email has been verified. You can now sign in.</Text>
				<Button asChild>
					<Link to="/login">Sign in</Link>
				</Button>
			</div>
		</div>
	);
}
