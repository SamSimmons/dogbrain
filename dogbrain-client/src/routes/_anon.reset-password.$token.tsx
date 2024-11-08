import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import {
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
	Form,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Heading } from "@/components/ui/typography";
import { ApiError, resetPassword } from "@/lib/api";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { z } from "zod";

export const Route = createFileRoute("/_anon/reset-password/$token")({
	component: NewPasswordComponent,
});

const resetPasswordSchema = z.object({
	password: z
		.string()
		.min(8, "Password must be at least 8 characters long")
		.max(128, "Password must not exceed 128 characters"),
	token: z.string().min(1),
});

type ResetPasswordSchema = z.infer<typeof resetPasswordSchema>;

function NewPasswordComponent() {
	const navigate = useNavigate();
	const { token } = Route.useParams();
	const mutation = useMutation({
		mutationFn: (values: ResetPasswordSchema) =>
			resetPassword(values.password, values.token),
		onSuccess: () => {
			navigate({ to: "/login" });
		},
	});

	const form = useForm<ResetPasswordSchema>({
		resolver: zodResolver(resetPasswordSchema),
		defaultValues: {
			token,
			password: "",
		},
	});

	function onSubmit(values: ResetPasswordSchema) {
		mutation.mutate(values);
	}

	return (
		<div className="grid place-items-center grow">
			<div className="flex flex-col gap-2">
				<Heading size="lg">Reset Passowrd</Heading>

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
							name="password"
							render={({ field }) => (
								<FormItem>
									<FormLabel>New Password</FormLabel>
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
							{mutation.isPending ? "Resetting..." : "Reset Password"}
						</Button>
					</form>
				</Form>
			</div>
		</div>
	);
}
