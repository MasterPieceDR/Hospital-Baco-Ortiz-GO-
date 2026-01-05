// src/components/Footer.jsx
import React from "react";
import { Link } from "react-router-dom";
import logoImg from "../assets/logo.jpeg";

export default function Footer() {
	return (
		<footer className="site-footer">
			<div className="container footer-inner">
				<div style={{ display: "flex", gap: 12, alignItems: "center" }}>
					<img
						src={logoImg}
						alt="MSP"
						style={{ height: 48 }}
						onError={(e) => { e.currentTarget.style.display = "none"; }}
					/>
					<div>
						<strong>Hospital Pediátrico Baca Ortiz</strong>
						<p>© 2025 — Ministerio de Salud Pública</p>
					</div>
				</div>
				<div className="footer-links">
					<Link to="/servicios">Servicios</Link>
					<Link to="/transparencia">Transparencia</Link>
					<Link to="/contacto">Contacto</Link>
				</div>
			</div>
		</footer>
	);
}
