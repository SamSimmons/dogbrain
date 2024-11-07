import {
	NavigationMenu,
	NavigationMenuItem,
	NavigationMenuLink,
	NavigationMenuList,
} from "@/components/ui/navigation-menu";
import { useAuth } from "@/hooks/useAuth";
import { NavigationMenuTrigger } from "@radix-ui/react-navigation-menu";
import { Link } from "@tanstack/react-router";

export function AnonHeader() {
	const auth = useAuth();
	console.log("🧠", { auth });
	return (
		<NavigationMenu className="flex justify-between items-center w-full grow-0">
			<NavigationMenuList>
				<Link to="/">
					<NavigationMenuItem title="TODO: logo">🧠</NavigationMenuItem>
				</Link>
				<NavigationMenuItem>
					<NavigationMenuLink>About Us</NavigationMenuLink>
				</NavigationMenuItem>
				<NavigationMenuItem>
					<NavigationMenuLink>Blog</NavigationMenuLink>
				</NavigationMenuItem>
			</NavigationMenuList>
			{!auth.isAuthenticated ? (
				<NavigationMenuList>
					<NavigationMenuItem>
						<Link to="/login">Log In</Link>
					</NavigationMenuItem>
					<NavigationMenuItem>
						<Link to="/sign-up">Sign Up</Link>
					</NavigationMenuItem>
				</NavigationMenuList>
			) : (
				<NavigationMenuList>
					<NavigationMenuItem>
						<NavigationMenuTrigger onClick={auth.logout}>
							Log Out
						</NavigationMenuTrigger>
					</NavigationMenuItem>
				</NavigationMenuList>
			)}
		</NavigationMenu>
	);
}
