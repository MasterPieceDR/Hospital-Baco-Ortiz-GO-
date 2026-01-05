import React from "react";

export default function Servicios() {
	return (
		<div className="max-w-6xl mx-auto py-16 px-4">
			<h1 className="text-3xl font-bold text-center mb-6">Servicios del Hospital</h1>

			<p className="text-gray-700 text-center mb-10">
				Consulta externa, emergencias, hospitalización y cuidados intensivos.
			</p>

			<div className="grid md:grid-cols-2 gap-6">
				<div className="p-6 shadow rounded bg-white">
					<h2 className="text-xl font-semibold mb-2">Emergencias 24/7</h2>
					<p className="text-gray-600">Atención inmediata para pacientes pediátricos.</p>
				</div>

				<div className="p-6 shadow rounded bg-white">
					<h2 className="text-xl font-semibold mb-2">Hospitalización</h2>
					<p className="text-gray-600">Cuidados continuos para niños y adolescentes.</p>
				</div>

				<div className="p-6 shadow rounded bg-white">
					<h2 className="text-xl font-semibold mb-2">Cuidados Intensivos</h2>
					<p className="text-gray-600">Equipos modernos para tratamientos críticos.</p>
				</div>

				<div className="p-6 shadow rounded bg-white">
					<h2 className="text-xl font-semibold mb-2">Consulta Externa</h2>
					<p className="text-gray-600">Atención en múltiples especialidades pediátricas.</p>
				</div>
			</div>
		</div>
	);
}
