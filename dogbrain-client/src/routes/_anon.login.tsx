import {
	createFileRoute,
	Link,
	useNavigate,
	useRouter,
} from "@tanstack/react-router";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";

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
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Heading, Text } from "@/components/ui/typography";
import { ApiError, logIn } from "@/lib/api";
import { useAuth } from "@/hooks/useAuth";

export const Route = createFileRoute("/_anon/login")({
	validateSearch: z.object({
		redirect: z.string().optional(),
	}),
	component: LogInComponent,
});

const loginSchema = z.object({
	email: z
		.string()
		.trim()
		.min(1, "Email is required")
		.max(255, "Email must not exceed 255 characters")
		.email("Invalid email address"),
	password: z
		.string()
		.min(8, "Password must be at least 8 characters long")
		.max(128, "Password must not exceed 128 characters"),
});

type LogInSchema = z.infer<typeof loginSchema>;

function LogInComponent() {
	const { login } = useAuth();
	const navigate = useNavigate();
	const router = useRouter();
	const { redirect } = Route.useSearch();

	const mutation = useMutation({
		mutationFn: (values: LogInSchema) => logIn(values.email, values.password),
		onSuccess: async (_response, values) => {
			await login(values.email);
			await router.invalidate();
			await new Promise((r) => setTimeout(r, 1));
			navigate({ to: redirect ?? "/dashboard" });
		},
	});

	const form = useForm<LogInSchema>({
		resolver: zodResolver(loginSchema),
		defaultValues: {
			email: "",
			password: "",
		},
	});

	function onSubmit(values: LogInSchema) {
		mutation.mutate(values);
	}

	return (
		<div className="grid place-items-center grow">
			<div className="flex flex-col gap-2">
				<Heading size="lg">Log In</Heading>
				<Text variant="muted">
					Don't have an account? <Link to="/sign-up">Sign up.</Link>
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
											autoComplete="current-password"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>
						<Text variant="muted">
							Forgot your password? <Link to="/forgot-password">Reset</Link>
						</Text>
						<Button type="submit" disabled={mutation.isPending}>
							{mutation.isPending ? "Logging In..." : "Log In"}
						</Button>
					</form>
				</Form>
			</div>
		</div>
	);
}
