import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({ variable: "--font-sans", subsets: ["latin"] });
const geistMono = Geist_Mono({ variable: "--font-mono", subsets: ["latin"] });

export const metadata: Metadata = {
  metadataBase: new URL("https://alanthssss.github.io/blossom-router/"),
  title: "Blossom Router — One prompt. The right model.",
  description: "A lightweight, local-first CLI that routes AI prompts across Ollama and configurable cloud models.",
  openGraph: { title: "Blossom Router", description: "Route everyday AI work wisely.", type: "website", images: ["/og.png"] },
  twitter: { card: "summary_large_image", title: "Blossom Router", description: "One prompt. The right model.", images: ["/og.png"] },
};

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return <html lang="en"><body className={`${geistSans.variable} ${geistMono.variable}`}>{children}</body></html>;
}
