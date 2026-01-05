// src/components/ServicesSection.jsx
import React from "react";
import { Link } from "react-router-dom";
import { FaStethoscope, FaClinicMedical, FaUsers, FaPhone } from "react-icons/fa";
import { MdOutlineFactCheck } from "react-icons/md";

const cards = [
	{ title: "Servicios", desc: "Consulta externa, emergencias, hospitalización, cuidados intensivos." },
	{ title: "Clínicas", desc: "Especialidades pediátricas: neurología, cardiología, neonatología y más." },
	{ title: "Pacientes y Familias", desc: "Información clave para pacientes, requisitos y atención." },
	{ title: "Transparencia", desc: "Acceso público a documentos e informes de gestión." },
	{ title: "Contacto", desc: "Números de emergencia y correo institucional." },
];

export default function ServicesSection() {
	return (
		<section className="services container py-10">
			<h2 className="section-title text-center text-2xl font-bold mb-8">
				Servicios del Hospital
			</h2>
			<div className="cards grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-5">
				{cards.map((c) => (
					<div
						key={c.title}
						className="card bg-white shadow-md rounded-xl p-5 flex flex-col items-center
                                   text-center hover:shadow-lg transition"
					>
						<div className="card-icon text-blue-600" style={{ fontSize: '32px' }}>
							{c.title === "Servicios" && <FaStethoscope />}
							{c.title === "Clínicas" && <FaClinicMedical />}
							{c.title === "Pacientes y Familias" && <FaUsers />}
							{c.title === "Transparencia" && <MdOutlineFactCheck />}
							{c.title === "Contacto" && <FaPhone />}
						</div>
						<h3 className="font-semibold mt-3">{c.title}</h3>
						<p className="text-sm text-gray-600 mt-1">{c.desc}</p>
						<Link to="/servicios" className="card-link mt-3 text-blue-500 hover:underline">
							Ver más
						</Link>
					</div>
				))}
			</div>
		</section>
	);
}
