import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "{{.ProjectName}}",
  description: "A web application built with Next.js and Go",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="antialiased">{children}</body>
    </html>
  );
}
