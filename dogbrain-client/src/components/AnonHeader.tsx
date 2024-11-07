import {
	NavigationMenu,
	NavigationMenuItem,
	NavigationMenuLink,
	NavigationMenuList,
} from "@/components/ui/navigation-menu";
import { Link } from "@tanstack/react-router";

export function AnonHeader() {
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
			<NavigationMenuList>
				<NavigationMenuItem>
					<NavigationMenuLink>Log In</NavigationMenuLink>
				</NavigationMenuItem>
				<NavigationMenuItem>
					<Link to="/sign-up">Sign Up</Link>
				</NavigationMenuItem>
			</NavigationMenuList>
		</NavigationMenu>
	);
}
