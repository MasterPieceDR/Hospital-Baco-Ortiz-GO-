import React, { useEffect, useState } from "react";
import api from "../../../api/axios";

export default function Citas() {
	const [list, setList] = useState([]);
	const [loading, setLoading] = useState(true);

	const fetch = () => {
		setLoading(true);
		api.get("/citas")
			.then(r => setList(r.data || []))
			.catch(e => console.error(e))
			.finally(()=>setLoading(false));
	};

	useEffect(()=>fetch(),[]);

	const changeState = async (citaId, estado) => {
		try {
			await api.put(`/citas/${citaId}`, { estado });
			fetch();
		} catch (e) { alert(e.response?.data?.error || e.message); }
	};

	const remove = async (c) => {
		if(!confirm("Eliminar cita?")) return;
		await api.delete(`/citas/${c.id || c.cita_id}`);
		fetch();
	};

	if (loading) return <div style={{padding:24}}>Cargando citas...</div>;

	return (
		<div style={{padding:24}}>
			<div style={{display:"flex",justifyContent:"space-between",alignItems:"center"}}>
				<h2>Control de Citas</h2>
				<button className="btn" onClick={fetch}>Refrescar</button>
			</div>

			{list.map(c => (
				<div className="card" key={c.id || c.cita_id}>
					<div className="list-row">
						<div>
							<div><strong>Fecha:</strong> {new Date(c.fecha_hora || c.fecha_hora || c.fecha).toLocaleString()}</div>
							<div style={{color:"#666"}}>Doctor ID: {c.medico_id || c.medicoId || "-"}</div>
						</div>
						<div style={{display:"flex",gap:8}}>
							<button className="btn secondary" onClick={()=>changeState(c.id || c.cita_id, "Completada")}>Marcar completada</button>
							<button className="btn danger" onClick={()=>remove(c)}>Eliminar</button>
						</div>
					</div>
				</div>
			))}
		</div>
	);
}
