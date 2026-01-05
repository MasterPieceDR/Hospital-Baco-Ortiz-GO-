import { Outlet, Link } from "react-router-dom";
import Sidebar from "../dashboard/Sidebar";

export default function DashboardLayout() {
  return (
    <div className="flex min-h-screen bg-gray-100">
      <Sidebar />

      <div className="flex-1">
        {/* Header */}
        <header className="flex items-center justify-between bg-white px-6 py-4 shadow">
          <h1 className="text-lg font-semibold">Panel Administrativo</h1>

          <Link
            to="/"
            style={{
              borderRadius: 8,
              background: "linear-gradient(90deg, var(--accent), var(--accent-dark))",
              color: "#fff",
              padding: "10px 22px",
              fontWeight: 700,
              fontSize: 16,
              boxShadow: "0 2px 8px rgba(255,122,24,0.10)",
              border: "none",
              transition: "background 0.2s"
            }}
            onMouseOver={e => e.currentTarget.style.background = "linear-gradient(90deg, var(--accent-dark), var(--accent))"}
            onMouseOut={e => e.currentTarget.style.background = "linear-gradient(90deg, var(--accent), var(--accent-dark))"}
          >
            Volver al sitio
          </Link>
        </header>

        {/* Contenido */}
        <main className="p-6">
          <Outlet />
        </main>
      </div>
    </div>
  );
}
