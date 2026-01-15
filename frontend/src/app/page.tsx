import { checkHealth } from "@/lib/api";

export default async function Home() {
  let healthStatus = "unknown";

  try {
    const health = await checkHealth();
    healthStatus = health.status;
  } catch (error) {
    healthStatus = "Backend not reachable";
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">
          Welcome to {"{{.ProjectName}}"}
        </h1>
        <p className="text-gray-600 mb-8">
          Your Next.js + Go + Supabase application
        </p>

        <div className="bg-gray-100 rounded-lg p-4 mb-8">
          <p className="text-sm text-gray-500">Backend Health</p>
          <p className={`font-mono ${healthStatus === "ok" ? "text-green-600" : "text-red-600"}`}>
            {healthStatus}
          </p>
        </div>

        <div className="space-y-2 text-sm text-gray-500">
          <p>Edit <code className="bg-gray-100 px-1 rounded">src/app/page.tsx</code> to get started</p>
          <p>Backend API: <code className="bg-gray-100 px-1 rounded">{process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"}</code></p>
        </div>
      </div>
    </main>
  );
}
