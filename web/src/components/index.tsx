import React from "react"
import { Card, CardContent } from "./ui/card"
import { Label } from "./ui/label"
import { Copy, Link2, Scissors, Unlink2 } from "lucide-react"
import { Input } from "./ui/input"
import { Button } from "./ui/button"
import { Separator } from "./ui/separator"
import { Item, ItemActions, ItemContent, ItemDescription } from "./ui/item"

export function Index() {
    const [shortUrl, setShortUrl] = React.useState("")
    const [inputUrl, setInputUrl] = React.useState("")

    const handleCopy = () => {
        if (!shortUrl) return
        navigator.clipboard.writeText(shortUrl)
    }

    async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        const form = e.currentTarget
        const url = (form.elements.namedItem("url") as HTMLInputElement).value

        const formData = new FormData()
        formData.append('url', url)

        const res = await fetch("/api/shorten", {
            method: "POST",
            body: formData,
        })

        const data = await res.json()

        setShortUrl(data.data?.short_url)
        setInputUrl("")
    }

    return (
        <main>
            <Card className="max-w-md mx-auto mt-10">
                <CardContent>
                    <form onSubmit={handleSubmit}>
                        <div className="flex flex-col gap-6">
                            <div className="grid gap-2">
                                <Label htmlFor="url">
                                    <Unlink2 className="size-5" /> Paste your URL
                                </Label>
                                <Input
                                    id="url"
                                    type="url"
                                    placeholder="https://example.com/my-very-super-duper-long-url"
                                    autoComplete="off"
                                    className="h-14"
                                    value={inputUrl}
                                    onChange={(e) => setInputUrl(e.target.value)}
                                />
                                <Button type="submit" className="w-full h-10">
                                    <Scissors /> Shorten
                                </Button>
                            </div>
                        </div>
                    </form>
                </CardContent>

                <div className="px-6">
                    <Separator className="mt-2" />
                </div>

                <CardContent>
                    <div className="flex flex-col gap-6">
                        <div className="grid gap-2">
                            <Label htmlFor="url-result">
                                <Link2 className="size-5" /> Your shortened URL
                            </Label>
                            <Item variant="outline" size="sm">
                                <ItemContent>
                                    <ItemDescription>
                                        {shortUrl ? shortUrl : "Your shortened URL will appear here"}
                                    </ItemDescription>
                                </ItemContent>
                                <ItemActions>
                                    <Button type="button" size="sm" onClick={handleCopy}>
                                        <Copy /> Copy
                                    </Button>
                                </ItemActions>
                            </Item>
                        </div>
                    </div>
                </CardContent>
            </Card>
        </main>
    )
}