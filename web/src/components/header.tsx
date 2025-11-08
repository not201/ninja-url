import { ModeToggle } from "./mode-toggle";

export function Header() {
    return (
        <header>
            <div className="flex justify-end">
                <ModeToggle />
            </div>
            <h1 className="text-3xl sm:text-4xl text-center">
                <span className="font-extrabold text-primary">Ninja</span>URL
            </h1>

            <p
                className="text-center text-base sm:text-md italic tracking-wide mt-2 sm:mt-4"
            >
                Shorten them fast and easy — like a ninja
            </p>
        </header>)
}