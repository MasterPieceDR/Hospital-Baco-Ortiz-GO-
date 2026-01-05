import React from "react";
import { Outlet } from "react-router-dom";
import Sidebar from "../../dashboard/Sidebar";

export default function DashboardLayout() {
	return (
		<div style={{ display: "flex", minHeight: "100vh" }}>
			<Sidebar />
			<div style={{ flex: 1 }}>
				<header style={{ height: 64, background: "#fff", borderBottom: "1px solid #eee", padding: "12px 24px" }}>
					<h3 style={{ margin: 0 }}>Panel Administrativo</h3>
				</header>
				<main style={{ padding: 20, background: "#f7f8fa", minHeight: "calc(100vh - 64px)" }}>
					<Outlet />
				</main>
			</div>
		</div>
	);
}
