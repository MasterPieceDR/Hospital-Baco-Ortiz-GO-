import React from "react";
import { Link } from "react-router-dom";
import logoImg from "../assets/logo.jpeg";

export default function Navbar() {
	return (
		<header className="site-navbar">
			<div className="container nav-inner">
				<div className="brand">
					<Link to="/" className="brand-link">
						<img
							src={logoImg}
							alt="Hospital"
							className="brand-logo"
							onError={(e) => { e.currentTarget.onerror = null; e.currentTarget.src = "/logo192.png"; }}
						/>
						<span className="brand-text">Hospital Pediátrico Baca Ortiz</span>
					</Link>
				</div>
				<nav className="nav-links">
					<Link to="/" className="nav-link">Inicio</Link>
					<Link to="/servicios" className="nav-link">Servicios</Link>
					<Link to="/clinicas" className="nav-link">Clínicas</Link>
					<Link to="/transparencia" className="nav-link">Transparencia</Link>
					<Link to="/contacto" className="nav-link">Contacto</Link>
					<Link to="/login" className="login-btn">Iniciar Sesión</Link>
					<Link to="/" className="back-btn">Volver al sitio</Link>
				</nav>
			</div>
		</header>
	);
}
