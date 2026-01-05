import React, { useEffect, useState } from "react";
import api from "../../../api/axios";

export default function Consultas() {
	const [list, setList] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);

	// modal / detalle
	const [modalOpen, setModalOpen] = useState(false);
	const [modalLoading, setModalLoading] = useState(false);
	const [modalError, setModalError] = useState(null);
	const [selected, setSelected] = useState(null);

	const fetch = () => {
		setLoading(true); setError(null);
		// pedir un rango amplio para recuperar todas las consultas
		api.get("/consultas?desde=2000-01-01&hasta=2099-12-31")
			.then(r => setList(r.data || []))
			.catch(e => setError(e.response?.data?.error || e.message))
			.finally(()=>setLoading(false));
	};

	useEffect(()=>fetch(),[]);

	const openDetail = async (consulta) => {
		setModalOpen(true);
		setModalLoading(true);
		setModalError(null);
		setSelected(null);
		try {
			const res = await api.get(`/consultas/${consulta.id || consulta.consulta_id}`);
			setSelected(res.data);
		} catch (e) {
			setModalError(e.response?.data?.error || e.message || "Error cargando detalle");
		} finally {
			setModalLoading(false);
		}
	};

	if (loading) return <div style={{padding:24}}>Cargando historial de consultas...</div>;
	if (error) return <div style={{padding:24,color:"red"}}>{error}</div>;

	return (
		<div style={{padding:24}}>
			<h2>Historial de Consultas</h2>

			{list.length === 0 ? <div className="card">(sin consultas)</div> : (
				<ul style={{listStyle:"none",padding:0}}>
					{list.map(c => {
						const fechaRaw = c.fecha || c.fecha_consulta || c.fecha_consulta;
						const fechaStr = fechaRaw ? new Date(fechaRaw).toLocaleString() : "-";
						const diag = (c.diagnosticos && c.diagnosticos.length > 0 && (c.diagnosticos[0].descripcion || c.diagnosticos[0].Descripcion))
							? (c.diagnosticos[0].descripcion || c.diagnosticos[0].Descripcion)
							: (c.diagnostico || c.Diagnostico || "-");
						return (
							<li key={c.id || c.consulta_id} className="card" style={{marginBottom:12}}>
								<div style={{display:"flex",justifyContent:"space-between",alignItems:"center",gap:16}}>
									<div style={{flex:1}}>
										<div style={{fontWeight:700}}>{fechaStr} - Paciente {c.paciente_id || c.PacienteID || "-"}</div>
										<div style={{fontWeight:600, marginTop:6}}>Diagnóstico: <span style={{fontWeight:500}}>{diag}</span></div>
										<div style={{color:"#666",marginTop:8}}>{c.notas_medicas || c.notas || ""}</div>
									</div>
									<div style={{minWidth:120,textAlign:"right"}}>
										<button className="btn secondary" onClick={()=>openDetail(c)}>Ver</button>
									</div>
								</div>
							</li>
						);
					})}
				</ul>
			)}

			{/* Modal detalle */}
			{modalOpen && (
				<div className="modal-backdrop" onClick={() => setModalOpen(false)} style={{zIndex:1200}}>
					<div className="modal" onClick={(e)=>e.stopPropagation()} style={{maxWidth:760}}>
						{modalLoading ? (
							<div style={{padding:24}}>Cargando detalle...</div>
						) : modalError ? (
							<div style={{padding:24,color:"red"}}>{modalError}</div>
						) : selected ? (
							<div>
								<div style={{display:"flex",justifyContent:"space-between",alignItems:"center",marginBottom:12}}>
									<h3 style={{margin:0}}>Consulta - {new Date(selected.fecha || selected.fecha_consulta).toLocaleString()}</h3>
									<button className="btn secondary" onClick={()=>setModalOpen(false)}>Cerrar</button>
								</div>

								<div style={{marginBottom:10}}>
									<strong>Paciente:</strong> {selected.paciente_id || selected.PacienteID || "-"}
								</div>

								{selected.motivo && <div style={{marginBottom:10}}><strong>Motivo:</strong> {selected.motivo}</div>}
								{(selected.notas_medicas || selected.notas) && <div style={{marginBottom:10}}><strong>Notas:</strong> <div style={{marginTop:6}}>{selected.notas_medicas || selected.notas}</div></div>}

								<div style={{marginTop:8}}>
									<strong>Diagnósticos:</strong>
									{selected.diagnosticos && selected.diagnosticos.length > 0 ? (
										<ul style={{marginTop:8}}>
											{selected.diagnosticos.map(d => (
												<li key={d.diagnostico_id || d.DiagnosticoID} style={{marginBottom:6}}>
													<div style={{fontWeight:700}}>{d.descripcion || d.Descripcion}</div>
													<div style={{color:"#555"}}>{d.fecha ? new Date(d.fecha).toLocaleString() : ""}</div>
												</li>
											))}
										</ul>
									) : (
										<div style={{marginTop:8}}>Sin diagnósticos registrados.</div>
									)}
								</div>

								<div style={{display:"flex",justifyContent:"flex-end",gap:8, marginTop:16}}>
									<button className="btn secondary" onClick={()=>setModalOpen(false)}>Cerrar</button>
								</div>
							</div>
						) : (
							<div style={{padding:24}}>Sin datos</div>
						)}
					</div>
				</div>
			)}
		</div>
	);
}
