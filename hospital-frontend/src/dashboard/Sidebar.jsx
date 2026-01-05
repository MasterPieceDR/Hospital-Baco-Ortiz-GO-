import React, { useContext } from "react";
import { NavLink } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";

// Define qué secciones puede ver cada rol
const sidebarConfig = {
	admin: [
		{ to: "/dashboard", label: "Inicio", end: true },
		{ to: "/dashboard/doctores", label: "Doctores" },
		{ to: "/dashboard/pacientes", label: "Pacientes" },
		{ to: "/dashboard/citas", label: "Citas" },
		{ to: "/dashboard/consultas", label: "Consultas" },
	],
	doctor: [
		{ to: "/dashboard", label: "Inicio", end: true },
		{ to: "/dashboard/pacientes", label: "Pacientes" },
		{ to: "/dashboard/consultas", label: "Consultas" },
	],
	recepcion: [
		{ to: "/dashboard", label: "Inicio", end: true },
		{ to: "/dashboard/citas", label: "Citas" },
		{ to: "/dashboard/pacientes", label: "Pacientes" },
	],
	paciente: [
		{ to: "/dashboard", label: "Inicio", end: true },
		{ to: "/dashboard/consultas", label: "Consultas" },
		{ to: "/dashboard/citas", label: "Citas" },
	],
};

export default function Sidebar() {
	const { user } = useContext(AuthContext);
	// Determina el rol principal (puedes ajustar si hay múltiples roles)
	const role = user?.roles?.[0] || "admin";
	const links = sidebarConfig[role] || sidebarConfig.admin;

	const linkStyle = ({ isActive }) => ({
		display: "block",
		padding: "12px 16px",
		color: isActive ? "#fff" : "var(--primary)",
		background: isActive ? "linear-gradient(90deg, var(--primary), var(--primary-dark))" : "#eaf1fb",
		textDecoration: "none",
		borderRadius: 8,
		margin: "6px 8px",
		fontWeight: isActive ? 700 : 500,
		boxShadow: isActive ? "0 2px 8px rgba(11,102,209,0.10)" : "none",
		transition: "all 0.2s",
	});

	return (
		<aside style={{ width: 220, background: "#f4f8fc", minHeight: "100vh", color: "#063970", paddingTop: 20, borderRight: "1px solid #e0e7ef" }}>
			<div style={{ padding: "0 12px", fontWeight: 700, marginBottom: 12, color: "var(--primary)" }}>Panel {role.charAt(0).toUpperCase() + role.slice(1)}</div>
			<nav>
				{links.map(link => (
					<NavLink key={link.to} to={link.to} style={linkStyle} end={link.end}>
						{link.label}
					</NavLink>
				))}
			</nav>
		</aside>
	);
}
