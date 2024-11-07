import { Heading, Text } from "@/components/ui/typography";
import {
	createFileRoute,
	Link,
	redirect,
	useNavigate,
} from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@/components/ui/button";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useMutation } from "@tanstack/react-query";
import { ApiError, register } from "@/lib/api";
import { Alert, AlertDescription } from "@/components/ui/alert";

// eslint-disable-next-line @typescript-eslint/no-unnecessary-type-assertion
const fallback = "/dashboard" as const;

export const Route = createFileRoute("/_anon/sign-up")({
	validateSearch: z.object({
		redirect: z.string().optional().catch(""),
	}),
	beforeLoad: ({ context, search }) => {
		if (context.auth.isAuthenticated) {
			throw redirect({ to: search.redirect || fallback });
		}
	},
	component: SignUpComponent,
});

const registrationSchema = z.object({
	email: z
		.string()
		.trim() // matches strings.TrimSpace check
		.min(1, "Email is required")
		.max(255, "Email must not exceed 255 characters")
		.email("Invalid email address"),
	password: z
		.string()
		.min(8, "Password must be at least 8 characters long")
		.max(128, "Password must not exceed 128 characters"),
});

type RegistrationSchema = z.infer<typeof registrationSchema>;

function SignUpComponent() {
	const form = useForm<RegistrationSchema>({
		resolver: zodResolver(registrationSchema),
		defaultValues: {
			email: "",
			password: "",
		},
	});

	const navigate = useNavigate();
	const mutation = useMutation({
		mutationFn: (values: RegistrationSchema) =>
			register(values.email, values.password),
		onSuccess: () => {
			// Show success message and optionally redirect
			navigate({ to: "/check-email" });
		},
	});

	function onSubmit(values: RegistrationSchema) {
		mutation.mutate(values);
	}
	return (
		<div className="grid place-items-center grow">
			<div className="flex flex-col gap-2">
				<Heading size="lg">Sign up</Heading>
				<Text variant="muted">
					Already have an account? <Link to="/login">Sign in.</Link>
				</Text>

				{mutation.isError && (
					<Alert variant="destructive">
						<AlertDescription>
							{mutation.error instanceof ApiError
								? mutation.error.message
								: "Something went wrong. Please try again."}
						</AlertDescription>
					</Alert>
				)}

				<Form {...form}>
					<form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
						<FormField
							control={form.control}
							name="email"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Email</FormLabel>
									<FormControl>
										<Input
											placeholder="email"
											autoComplete="email"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>
						<FormField
							control={form.control}
							name="password"
							render={({ field }) => (
								<FormItem>
									<FormLabel>Password</FormLabel>
									<FormControl>
										<Input
											type="password"
											autoComplete="new-password"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>
						<Button type="submit" disabled={mutation.isPending}>
							{mutation.isPending ? "Registering..." : "Register"}
						</Button>
					</form>
				</Form>
			</div>
		</div>
	);
}
