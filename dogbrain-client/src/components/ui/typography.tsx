import { cn } from "@/lib/utils";
import { type VariantProps, cva } from "class-variance-authority";

const textVariants = cva(
	"font-mono leading-normal  [&_a]:text-indigo-700 [&_a:visited]:text-indigo-400",
	{
		variants: {
			size: {
				sm: "text-sm",
				default: "text-base",
				lg: "text-lg",
				xl: "text-xl",
				"2xl": "text-2xl",
			},
			variant: {
				default: "text-zinc-900",
				muted: "text-zinc-700",
				error: "text-red-500",
			},
			bold: {
				true: "font-bold",
				false: "font-normal",
			},
		},
		defaultVariants: {
			size: "default",
			variant: "default",
			bold: false,
		},
	},
);

export interface TextProps
	extends React.HTMLAttributes<HTMLParagraphElement>,
		VariantProps<typeof textVariants> {}

export function Text({
	className,
	size,
	variant,
	bold,
	children,
	...props
}: TextProps) {
	return (
		<p
			className={cn(textVariants({ size, variant, bold }), className)}
			{...props}
		>
			{children}
		</p>
	);
}

const headingVariants = cva("font-mono font-bold tracking-tight", {
	variants: {
		size: {
			default: "text-3xl",
			lg: "text-4xl",
			xl: "text-5xl",
		},
		variant: {
			default: "text-zinc-900",
			muted: "text-zinc-500",
			error: "text-red-500",
		},
	},
	defaultVariants: {
		size: "default",
		variant: "default",
	},
});

export interface HeadingProps
	extends React.HTMLAttributes<HTMLHeadingElement>,
		VariantProps<typeof headingVariants> {}

export function Heading({
	className,
	size,
	variant,
	children,
	...props
}: HeadingProps) {
	return (
		<h1
			className={cn(headingVariants({ size, variant }), className)}
			{...props}
		>
			{children}
		</h1>
	);
}
