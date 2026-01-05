// src/components/HeroCarousel.jsx
import React, { useState, useEffect } from "react";
import hospital1 from "../assets/hospital1.jpg";
import hospital2 from "../assets/hospital2.jpg";
import heroImg from "../assets/fondo_hospital.jpeg";

const images = [hospital1, hospital2];

export default function HeroCarousel() {
	return (
		<section className="hero">
			<div
				className="hero-bg"
				style={{ backgroundImage: `url(${heroImg})` }}
			/>
			<div className="hero-inner container">
				<h1 className="hero-title">Tecnología Médica Avanzada</h1>
				<p className="hero-sub">
					Equipos modernos para diagnóstico y tratamiento preciso.
				</p>
				<a href="/login" className="hero-cta">
					Iniciar sesión
				</a>
			</div>
		</section>
	);
}

export function HeroHeader() {
	return (
		<section style={{ padding: 40, background: "#fff" }}>
			<div style={{ maxWidth: 1100, margin: "0 auto" }}>
				<h1 style={{ color: "#123688", marginBottom: 8 }}>
					Bienvenido al Hospital Pediátrico
				</h1>
				<p style={{ color: "#333" }}>
					Atención de calidad para niños y adolescentes. Usa el botón Iniciar
					Sesión para acceder al panel.
				</p>
			</div>
		</section>
	);
}
