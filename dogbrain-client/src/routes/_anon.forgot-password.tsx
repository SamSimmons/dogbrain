import { Alert, AlertDescription } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  Form,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Heading } from '@/components/ui/typography'
import { ApiError, forgotPassword } from '@/lib/api'
import { zodResolver } from '@hookform/resolvers/zod'
import { useMutation } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

export const Route = createFileRoute('/_anon/forgot-password')({
  component: ResetPasswordComponent,
})

const resetPasswordSchema = z.object({
  email: z
    .string()
    .trim()
    .min(1, 'Email is required')
    .max(255, 'Email must not exceed 255 characters')
    .email('Invalid email address'),
})

type ResetPasswordSchema = z.infer<typeof resetPasswordSchema>

function ResetPasswordComponent() {
  const [requestSent, setRequestSent] = useState(false)

  const mutation = useMutation({
    mutationFn: (values: ResetPasswordSchema) => forgotPassword(values.email),
    onSuccess: () => {
      setRequestSent(true)
    },
  })

  const form = useForm<ResetPasswordSchema>({
    resolver: zodResolver(resetPasswordSchema),
    defaultValues: {
      email: '',
    },
  })

  function onSubmit(values: ResetPasswordSchema) {
    mutation.mutate(values)
  }

  return (
    <div className="grid place-items-center grow">
      <div className="flex flex-col gap-2">
        <Heading size="lg">Forgot Passowrd</Heading>

        {mutation.isError && (
          <Alert variant="destructive">
            <AlertDescription>
              {mutation.error instanceof ApiError
                ? mutation.error.message
                : 'Something went wrong. Please try again.'}
            </AlertDescription>
          </Alert>
        )}

        {requestSent ? (
          <Alert>
            <AlertDescription>
              If an account with that email exists, you'll recieve a link to
              reset your password.
            </AlertDescription>
          </Alert>
        ) : (
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
              <Button type="submit" disabled={mutation.isPending}>
                {mutation.isPending ? 'Resetting...' : 'Reset Password'}
              </Button>
            </form>
          </Form>
        )}
      </div>
    </div>
  )
}
